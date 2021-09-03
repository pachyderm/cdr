package main

import (
	"bytes"
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
	if err := ensureBucket(ctx, client, bucketName); err != nil {
		return err
	}
	creator := cdr.NewCreator([]cdr.Middleware{
		cdr.CompressGzip,
		cdr.EncryptChaCha20,
		cdr.HashBlake2b256,
	}, func(ctx context.Context, data []byte) (*cdr.Ref, error) {
		key := fmt.Sprintf("test-object-key-%d", time.Now().UnixNano())
		r := bytes.NewReader(data)
		if _, err := client.PutObject(ctx, bucketName, key, r, int64(len(data)), minio.PutObjectOptions{}); err != nil {
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

func ensureBucket(ctx context.Context, client *minio.Client, bucketName string) error {
	if exists, err := client.BucketExists(ctx, bucketName); err != nil {
		return err
	} else if !exists {
		if err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			return err
		}
	}
	return nil
}
