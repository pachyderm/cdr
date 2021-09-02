package cdr

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/blake2b"
	"google.golang.org/protobuf/proto"
)

func TestExample(t *testing.T) {
	catURL := "https://i.imgur.com/QnkFrG3.gif"
	ref := CreateHTTPRef(catURL, nil)
	ref2 := CreateConcatRef([]*Ref{ref, ref, ref})
	data, _ := proto.Marshal(ref2)
	fmt.Println(hex.Dump(data))
}

func TestCreateResolve(t *testing.T) {
	client := NewTestMinioClient(t)
	bucketName := NewTestMinioBucket(t, client)

	secret := blake2b.Sum256([]byte("not a very good key stretched in not a very good way"))
	creator := NewCreator([]Middleware{
		CompressGzip,
		EncryptChaCha20(secret[:]),
		HashBlake2b256,
	},
		func(ctx context.Context, r io.Reader) (*Ref, error) {
			key := "test-object-key"
			if _, err := client.PutObject(ctx, bucketName, key, r, -1, minio.PutObjectOptions{}); err != nil {
				return nil, err
			}
			return CreateMinioRef(ctx, client, bucketName, key, time.Hour)
		})
	resolver := NewTestResolver(t)

	testData := []byte("test data goes here")
	ref := makeRef(t, creator, testData)
	actual := deref(t, resolver, ref)
	require.Equal(t, testData, actual)
}

func makeRef(t testing.TB, creator *Creator, data []byte) *Ref {
	ctx := context.Background()
	ref, err := creator.MakeRef(ctx, data)
	require.NoError(t, err)
	return ref
}

func deref(t testing.TB, resolver *Resolver, ref *Ref) []byte {
	ctx := context.Background()
	rc, err := resolver.Deref(ctx, ref)
	require.NoError(t, err)
	data, err := ioutil.ReadAll(rc)
	require.NoError(t, err)
	return data
}
