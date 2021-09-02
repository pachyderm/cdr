// package cdr provides Common Data Refs
package cdr

func IsImmutable(ref *Ref) bool {
	panic("not implemented")
}

func MinSize(ref *Ref) int64 {
	panic("not implemented")
}

func MaxSize(ref *Ref) int64 {
	panic("not implemented")
}

func CreateConcatRef(refs []*Ref) *Ref {
	return &Ref{
		Body: &Ref_Concat{Concat: &Concat{Refs: refs}},
	}
}

func CreateSliceRef(x *Ref, start, end uint64) *Ref {
	if end < start {
		panic("end must be >= start")
	}
	return &Ref{
		Body: &Ref_Slice{Slice: &Slice{Inner: x, Start: start, End: end}},
	}
}
