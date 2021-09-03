package cdr

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/chacha20"
)

type PostFunc = func(ctx context.Context, data []byte) (*Ref, error)

type Middleware = func(ctx context.Context, in []byte, next PostFunc) (*Ref, error)

type Creator struct {
	layers   []Middleware
	postFunc PostFunc
}

func NewCreator(layers []Middleware, postFunc PostFunc) *Creator {
	return &Creator{
		layers:   layers,
		postFunc: postFunc,
	}
}

func (c *Creator) MakeRef(ctx context.Context, data []byte) (*Ref, error) {
	var getPostFunc func(i int) PostFunc
	getPostFunc = func(i int) PostFunc {
		if i >= len(c.layers) {
			return c.postFunc
		}
		return func(ctx context.Context, in []byte) (*Ref, error) {
			return c.layers[i](ctx, in, getPostFunc(i+1))
		}
	}
	return getPostFunc(0)(ctx, data)
}

func CompressGzip(ctx context.Context, in []byte, next PostFunc) (*Ref, error) {
	buf := bytes.NewBuffer(nil)
	gw := gzip.NewWriter(buf)
	if _, err := gw.Write(in); err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}
	ref, err := next(ctx, buf.Bytes())
	if err != nil {
		return nil, err
	}
	return &Ref{
		Body: &Ref_Compress{Compress: &Compress{
			Inner: ref,
			Algo:  CompressAlgo_GZIP,
		}},
	}, nil
}

func HashBlake2b256(ctx context.Context, in []byte, next PostFunc) (*Ref, error) {
	sum := blake2b.Sum256(in)
	ref, err := next(ctx, in)
	if err != nil {
		return nil, err
	}
	return &Ref{
		Body: &Ref_ContentHash{ContentHash: &ContentHash{
			Inner: ref,
			Algo:  HashAlgo_BLAKE2b_256,
			Hash:  sum[:],
		}},
	}, nil
}

// EncryptChaCha20 encrypts the data with a random key (stored in the ref)
// Does not protect the data from tampering.
func EncryptChaCha20(ctx context.Context, data []byte, next PostFunc) (*Ref, error) {
	nonce := make([]byte, chacha20.NonceSize)
	key := make([]byte, 32)
	if _, err := rand.Read(key[:]); err != nil {
		return nil, err
	}
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}
	ctext := make([]byte, len(data))
	cipher.XORKeyStream(ctext, data)
	ref, err := next(ctx, ctext)
	if err != nil {
		return nil, err
	}
	return &Ref{
		Body: &Ref_Cipher{Cipher: &Cipher{
			Inner: ref,
			Algo:  CipherAlgo_CHACHA20,
			Key:   key,
			Nonce: nonce,
		}},
	}, nil
}
