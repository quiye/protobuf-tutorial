package api

import (
	"testing"

	proto "github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

func TestNestFormula(t *testing.T) {

	f1 := &NestFormula{
		Elements: &NestFormula_Mono{11},
	}

	f2 := &NestFormula{
		Elements: &NestFormula_Poly{
			&PolyFormula{
				OpString:     "*",
				NestFormulas: []*NestFormula{f1, f1, f1},
			},
		},
	}
	f3 := &NestFormula{
		Elements: &NestFormula_Mono{44},
	}
	f4 := &NestFormula{
		Elements: &NestFormula_Poly{
			&PolyFormula{
				OpString:     "+",
				NestFormulas: []*NestFormula{f2, f3},
			},
		},
	}

	s := f4.String()

	exp := `poly:<op_string:"+" nest_formulas:<poly:<op_string:"*" nest_formulas:<mono:11 > nest_formulas:<mono:11 > nest_formulas:<mono:11 > > > nest_formulas:<mono:44 > > `
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

	removeSizecache(ans)
	removeSizecache(f4)
	if diff := cmp.Diff(f4, ans); diff != "" {
		t.Errorf(diff)
	}
}

func removeSizecache(p *NestFormula) {
	p.XXX_sizecache = 0
	switch x := p.Elements.(type) {
	case *NestFormula_Poly:
		x.Poly.XXX_sizecache = 999
		for _, f := range x.Poly.NestFormulas {
			removeSizecache(f)
		}
	}
}
