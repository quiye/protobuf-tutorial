package api

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

func TestNestFormula(t *testing.T) {

	f1 := &NestFormula{
		Type: NestFormula_MONO_FORMULA,
		Mono: 11,
	}
	f2 := &NestFormula{
		Type:     NestFormula_POLY_FORMULA,
		OpString: "*",
		Poly:     []*NestFormula{f1, f1},
	}
	f3 := &NestFormula{
		Type: NestFormula_MONO_FORMULA,
		Mono: 44,
	}
	f4 := &NestFormula{
		Type:     NestFormula_POLY_FORMULA,
		OpString: "+",
		Poly:     []*NestFormula{f2, f3},
	}

	s := f4.String()

	exp := `op_string:"+" poly:<op_string:"*" poly:<mono:11 > poly:<mono:11 > type:POLY_FORMULA > poly:<mono:44 > type:POLY_FORMULA `
	if s != exp {
		t.Error(s)
	}

	b, err := proto.Marshal(f4)
	if err != nil {
		t.Error(err)
	}

	ans := &NestFormula{}
	if err := proto.Unmarshal(b, ans); err != nil {
		t.Error(err)
	}

	ans.Poly[1].XXX_sizecache = 2 // need
	if diff := cmp.Diff(f4.Poly[1], ans.Poly[1]); diff != "" {
		t.Errorf(diff)
	}
}
