package cdr

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/cipher"
	"hash"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/chacha20"
)

type Resolver struct {
	httpClient *http.Client
}

func NewResolver() *Resolver {
	return &Resolver{
		httpClient: http.DefaultClient,
	}
}

func (r *Resolver) Deref(ctx context.Context, ref *Ref) (io.ReadCloser, error) {
	switch x := ref.Body.(type) {
	case *Ref_Http:
		return r.derefHTTP(ctx, x.Http)
	case *Ref_ContentHash:
		return r.derefContentHash(ctx, x.ContentHash)
	case *Ref_Cipher:
		return r.derefCipher(ctx, x.Cipher)
	case *Ref_Compress:
		return r.derefCompress(ctx, x.Compress)
	default:
		return nil, errors.Errorf("unsupported Ref variant %v: %T", ref.Body, ref.Body)
	}
}

func (r *Resolver) derefHTTP(ctx context.Context, ref *HTTP) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ref.Url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range ref.Headers {
		req.Header.Add(k, v)
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (r *Resolver) derefContentHash(ctx context.Context, ref *ContentHash) (io.ReadCloser, error) {
	innerRc, err := r.Deref(ctx, ref.Inner)
	if err != nil {
		return nil, err
	}
	// TODO: add a maximum limit here
	data, err := io.ReadAll(innerRc)
	if err != nil {
		return nil, err
	}
	var h hash.Hash
	switch ref.Algo {
	case HashAlgo_BLAKE2b_256:
		h, err = blake2b.New256(nil)
		if err != nil {
			panic(err)
		}
	default:
		return nil, errors.Errorf("unrecognized hash algo %v", ref.Algo)
	}
	h.Write(data)
	sum := h.Sum(nil)
	if !bytes.Equal(sum, ref.Hash) {
		return nil, errors.Errorf("content failed hash check HAVE: %x WANT: %x", sum, ref.Hash)
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}

func (r *Resolver) derefCipher(ctx context.Context, ref *Cipher) (io.ReadCloser, error) {
	innerRc, err := r.Deref(ctx, ref.Inner)
	if err != nil {
		return nil, err
	}
	var rd io.Reader
	switch ref.Algo {
	case CipherAlgo_CHACHA20:
		ciph, err := chacha20.NewUnauthenticatedCipher(ref.Key, ref.Nonce)
		if err != nil {
			return nil, err
		}
		rd = &cipher.StreamReader{R: innerRc, S: ciph}
	default:
		return nil, errors.Errorf("unrecognized cipher algo %v", ref.Algo)
	}
	rc := readCloser{r: rd, closes: []func() error{
		innerRc.Close,
	}}
	return rc, nil
}

func (r *Resolver) derefCompress(ctx context.Context, ref *Compress) (io.ReadCloser, error) {
	innerRc, err := r.Deref(ctx, ref.Inner)
	if err != nil {
		return nil, err
	}
	switch ref.Algo {
	case CompressAlgo_GZIP:
		gr, err := gzip.NewReader(innerRc)
		if err != nil {
			return nil, err
		}
		return readCloser{r: gr, closes: []func() error{
			innerRc.Close,
			gr.Close,
		}}, nil
	default:
		return nil, errors.Errorf("unrecognized compress algo %v", ref.Algo)
	}
}

type readCloser struct {
	r      io.Reader
	closes []func() error
}

func (c readCloser) Read(buf []byte) (int, error) {
	return c.r.Read(buf)
}

func (c readCloser) Close() error {
	var retErr error
	for i := range c.closes {
		if err := c.closes[i](); retErr == nil {
			retErr = err
		}
	}
	return retErr
}
