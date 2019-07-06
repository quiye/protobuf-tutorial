package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/golang/protobuf/proto"
	pb "github.com/quiye/protobuf-tutorial/api"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Body.String() != "Hello, World\n" {
		t.Error()
	}
}

func TestCalcPBHandler(t *testing.T) {
	rf := &pb.Formula{
		OpString: "*",
		Values:   []int32{9, 9},
	}

	b, err := proto.Marshal(rf)
	if err != nil {
		t.Error(err)
	}
	reqdata := bytes.NewReader(b)

	req := httptest.NewRequest("GET", "/hello", reqdata)
	rec := httptest.NewRecorder()
	calcPBHandler(false)(rec, req)

	if rec.Body.String() != "81\n" {
		t.Error(rec.Body.String())
	}
}

func TestOne(t *testing.T) {
	r := one()
	exp := 1
	if r != exp {
		t.Error()
	}
}

func TestCal(t *testing.T) {
	testCases := []struct {
		desc   string
		op     rune
		values []int32
		exp    int32
		err    string
	}{
		{
			desc:   "plus",
			op:     '+',
			values: []int32{1, 3, 5},
			exp:    9,
		},
		{
			desc:   "prod",
			op:     '*',
			values: []int32{1, 3, 5},
			exp:    15,
		},
		{
			desc:   "empty",
			values: []int32{},
			err:    "values must be a non empty list",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := cal(tC.op, tC.values)
			if err != nil && err.Error() != tC.err {
				t.Errorf("got %v, expected %v", err, tC.err)
			}
			if result != tC.exp {
				t.Errorf("got %v, expected %v", result, tC.exp)
			}
		})
	}
}

func BenchmarkCalcPBHandler(b *testing.B) {

	rf := &pb.Formula{
		OpString: "*",
		Values:   []int32{9, 9, 9, 9, 9},
	}

	bn, _ := proto.Marshal(rf)

	rec := httptest.NewRecorder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reqdata := bytes.NewReader(bn)
		req := httptest.NewRequest("GET", "/hello", reqdata)
		calcPBHandler(false)(rec, req)
	}
}
func BenchmarkCalcHandler(b *testing.B) {

	rf := &formula{
		OpString: "*",
		Args:     []int32{9, 9, 9, 9, 9},
	}

	bn, _ := json.Marshal(rf)

	rec := httptest.NewRecorder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reqdata := bytes.NewReader(bn)
		req := httptest.NewRequest("GET", "/hello", reqdata)
		calcHandler(false)(rec, req)
	}
}
