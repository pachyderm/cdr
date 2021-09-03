// package cdr provides Common Data Refs
package cdr

func IsImmutable(ref *Ref) bool {
	switch x := ref.Body.(type) {
	case *Ref_ContentHash:
		return true
	case nil:
		return true
	case *Ref_Concat:
		var ret bool
		for _, ref := range x.Concat.Refs {
			ret = ret && IsImmutable(ref)
		}
		return ret
	case *Ref_Cipher:
		return IsImmutable(x.Cipher.Inner)
	case *Ref_Compress:
		return IsImmutable(x.Compress.Inner)
	case *Ref_SizeLimits:
		return IsImmutable(x.SizeLimits.Inner)
	case *Ref_Slice:
		return IsImmutable(x.Slice.Inner)
	default:
		return false
	}
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
