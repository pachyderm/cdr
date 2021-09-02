package cdr

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/chacha20"
)

type PostFunc = func(ctx context.Context, r io.Reader) (*Ref, error)

type RefMapper = func(*Ref) *Ref

type Middleware = func(out, data []byte) ([]byte, RefMapper)

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
	in := data
	var out []byte
	refMaps := make([]RefMapper, len(c.layers))
	for i, layer := range c.layers {
		out, refMaps[i] = layer(out, in)
		in = out
		out = out[:0]
	}
	ref, err := c.postFunc(ctx, bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	for i := len(refMaps) - 1; i >= 0; i-- {
		ref = refMaps[i](ref)
	}
	return ref, nil
}

func HashBlake2b256(out, in []byte) ([]byte, RefMapper) {
	h, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}
	h.Write(in)
	out = append(out, in...)
	fn := func(x *Ref) *Ref {
		return &Ref{
			Body: &Ref_ContentHash{ContentHash: &ContentHash{
				Inner: x,
				Algo:  HashAlgo_BLAKE2b_256,
				Hash:  h.Sum(nil),
			}},
		}
	}
	return out, fn
}

func CompressGzip(out, in []byte) ([]byte, RefMapper) {
	buf := bytes.NewBuffer(out)
	gw := gzip.NewWriter(buf)
	if _, err := gw.Write(in); err != nil {
		panic(err)
	}
	if err := gw.Close(); err != nil {
		panic(err)
	}
	fn := func(x *Ref) *Ref {
		return &Ref{
			Body: &Ref_Compress{Compress: &Compress{
				Inner: x,
				Algo:  CompressAlgo_GZIP,
			}},
		}
	}
	return buf.Bytes(), fn
}

func EncryptChaCha20(key []byte) func(out, in []byte) ([]byte, RefMapper) {
	return func(out, in []byte) ([]byte, RefMapper) {
		nonce := make([]byte, chacha20.NonceSizeX)
		if _, err := rand.Read(nonce); err != nil {
			panic(err)
		}
		cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
		if err != nil {
			panic(err)
		}
		out = append(out, in...)
		cipher.XORKeyStream(out[len(out)-len(in):], in)
		fn := func(x *Ref) *Ref {
			return &Ref{
				Body: &Ref_Cipher{Cipher: &Cipher{
					Inner: x,
					Algo:  EncAlgo_CHACHA20,
					Key:   key,
					Nonce: nonce,
				}},
			}
		}
		return out, fn
	}
}
