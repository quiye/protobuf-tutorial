package api

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
)

func TestFormulaProtoBuf(t *testing.T) {

	f1 := &Formula{
		OpString: "*",
		Values:   []int32{1, 3, 5, 7, 9},
	}

	s := f1.String()

	exp := `op_string:"*" values:1 values:3 values:5 values:7 values:9 `
	if s != exp {
		t.Error(s)
	}

	b, err := proto.Marshal(f1)
	if err != nil {
		t.Error(err)
	}

	f2 := &Formula{}
	if err := proto.Unmarshal(b, f2); err != nil {
		t.Error(err)
	}

	f2.XXX_sizecache = 10 // need
	if diff := cmp.Diff(f1, f2); diff != "" {
		t.Errorf(diff)
	}
}
