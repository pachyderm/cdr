package cdr

import (
	"context"
	"io"
)

type Resolver struct {
}

func (r *Resolver) Deref(ctx context.Context, ref *Ref) (io.ReadCloser, error) {
	panic("not implemented")
}

func IsImmutable(ref *Ref) bool {
	panic("not implemented")
}

func MinSize(ref *Ref) int64 {
	panic("not implemented")
}

func MaxSize(ref *Ref) int64 {
	panic("not implemented")
}
