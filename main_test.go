package main

import (
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()
	handler(rec, req)

	if rec.Body.String() != "Hello, World\n" {
		t.Error()
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
		values []int
		exp    int
		err    string
	}{
		{
			desc:   "plus",
			op:     '+',
			values: []int{1, 3, 5},
			exp:    9,
		},
		{
			desc:   "prod",
			op:     '*',
			values: []int{1, 3, 5},
			exp:    15,
		},
		{
			desc:   "empty",
			values: []int{},
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
