package formula

import (
	"testing"
)

func TestFormulaProtoBuf(t *testing.T) {

	p := Formula{
		OpString: "*",
		Values:   []int32{1, 3, 5, 7, 9},
	}

	s := p.String()

	exp := `op_string:"*" values:1 values:3 values:5 values:7 values:9 `
	if s != exp {
		t.Error(s)
	}
}
