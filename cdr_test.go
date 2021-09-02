package cdr

import (
	"encoding/hex"
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestExample(t *testing.T) {
	catURL := "https://i.imgur.com/QnkFrG3.gif"
	ref := &Ref{
		Body: &Ref_Http{Http: &HTTP{
			Url: catURL,
		}},
	}
	ref2 := &Ref{
		Body: &Ref_Concat{Concat: &Concat{
			Refs: []*Ref{ref, ref, ref},
		}},
	}
	data, _ := proto.Marshal(ref2)
	fmt.Println(hex.Dump(data))
}
