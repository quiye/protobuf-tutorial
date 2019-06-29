package main

import "testing"

func TestOne(t *testing.T) {
	r := one()
	exp := 1
	if r != exp {
		t.Error()
	}
}
