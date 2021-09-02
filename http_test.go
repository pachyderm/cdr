package cdr

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"net/url"
	"testing"
	"time"

	docker "github.com/docker/docker/client"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/require"
)

func TestCreateMinioRef(t *testing.T) {
	ctx := context.Background()
	client := NewTestMinioClient(t)
	bucket := NewTestMinioBucket(t, client)
	key := "my_test_object"
	data := []byte([]byte("test data goes here"))
	_, err := client.PutObject(ctx, bucket, key, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{})
	require.NoError(t, err)
	ref, err := CreateMinioRef(ctx, client, bucket, key, time.Hour)
	require.NoError(t, err)
	res := NewTestResolver(t)
	rc, err := res.Deref(ctx, ref)
	require.NoError(t, err)
	actual, err := ioutil.ReadAll(rc)
	require.NoError(t, err)
	require.Equal(t, data, actual)
}

func NewTestMinioClient(t testing.TB) *minio.Client {
	endpoint := getMinioEndpoint()
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	require.NoError(t, err)
	return client
}

func NewTestMinioBucket(t testing.TB, client *minio.Client) string {
	ctx := context.Background()
	bucketName := testName(t)
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := client.RemoveBucketWithOptions(ctx, bucketName, minio.BucketOptions{ForceDelete: true})
		require.NoError(t, err)
	})
	return bucketName
}

func getDockerHost() string {
	client, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	host := client.DaemonHost()
	u, err := url.Parse(host)
	if err != nil {
		panic(err)
	}
	if u.Scheme == "unix" {
		return "127.0.0.1"
	}
	return u.Hostname()
}

func getMinioEndpoint() string {
	host := getDockerHost()
	port := 9000
	return fmt.Sprintf("%s:%d", host, port)
}

func testName(t testing.TB) string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf[:])
	require.NoError(t, err)
	return fmt.Sprintf("test-%x", buf)
}

func NewTestResolver(t testing.TB) *Resolver {
	return NewResolver()
}
