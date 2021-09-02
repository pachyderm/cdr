package cdr

import (
	"context"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

func CreateHTTPRef(url string, headers map[string]string) *Ref {
	return &Ref{
		Body: &Ref_Http{Http: &HTTP{
			Url:     url,
			Headers: headers,
		}},
	}
}

func CreateMinioRef(ctx context.Context, client *minio.Client, bucket, key string, ttl time.Duration) (*Ref, error) {
	reqParams := make(url.Values)
	psURL, err := client.PresignedGetObject(ctx, bucket, key, ttl, reqParams)
	if err != nil {
		return nil, err
	}
	return CreateHTTPRef(psURL.String(), nil), nil
}

func CreateGCSRef(ctx context.Context, client interface{}, bucket, key string, ttl time.Duration) (*Ref, error) {
	panic("not implemented")
}

func CreateAWSRef(ctx context.Context, client interface{}, bucket, key string, ttl time.Duration) (*Ref, error) {
	panic("not implemented")
}

func CreateAzureRef(ctx context.Context, client interface{}, bucket, key string, ttl time.Duration) (*Ref, error) {
	panic("not implemented")
}
