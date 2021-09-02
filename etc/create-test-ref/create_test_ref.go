package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pachyderm/cdr"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	if len(os.Args) < 2 {
		return errors.Errorf("must specifiy minio endpoint as first arg e.g. 127.0.0.1:9000")
	}
	endpoint := os.Args[1]
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		return err
	}
	bucketName := "test-bucket"
	if exists, err := client.BucketExists(ctx, bucketName); err != nil {
		return err
	} else if !exists {
		if err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}
	creator := cdr.NewCreator([]cdr.Middleware{
		cdr.CompressGzip,
		cdr.EncryptChaCha20(make([]byte, 32)),
		cdr.HashBlake2b256,
	},
		func(ctx context.Context, r io.Reader) (*cdr.Ref, error) {
			key := fmt.Sprintf("test-object-key-%d", time.Now().UnixNano())
			if _, err := client.PutObject(ctx, bucketName, key, r, -1, minio.PutObjectOptions{}); err != nil {
				return nil, err
			}
			return cdr.CreateMinioRef(ctx, client, bucketName, key, time.Hour)
		})
	log.Println("reading from stdin...")
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	ref, err := creator.MakeRef(ctx, data)
	if err != nil {
		return err
	}
	refData, err := proto.Marshal(ref)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write([]byte(base64.URLEncoding.EncodeToString(refData)))
	return err
}
