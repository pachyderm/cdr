package main

import (
	"context"
	"io"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/pachyderm/cdr"
	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(derefCmd)
}

var rootCmd = &cobra.Command{
	Use:   "cdr",
	Short: "Common Data Ref Tool",
}

var derefCmd = &cobra.Command{
	Use:   "deref",
	Short: "dereferences a CDR and writes it's content to standard output",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		r := cdr.NewResolver()
		ref, err := readRef(cmd.InOrStdin())
		if err != nil {
			return err
		}
		rc, err := r.Deref(ctx, ref)
		if err != nil {
			return err
		}
		defer rc.Close()
		_, err = io.Copy(cmd.OutOrStdout(), rc)
		return err
	},
}

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "inspects a CDR from standard input",
	RunE: func(cmd *cobra.Command, args []string) error {
		ref, err := readRef(cmd.InOrStdin())
		if err != nil {
			return err
		}
		marsh := jsonpb.Marshaler{
			Indent:       " ",
			EmitDefaults: true,
		}
		return marsh.Marshal(cmd.OutOrStdout(), ref)
	},
}

func readRef(r io.Reader) (*cdr.Ref, error) {
	input, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var ref cdr.Ref
	if err := ref.UnmarshalBase64(input); err != nil {
		return nil, err
	}
	return &ref, nil
}
