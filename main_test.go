package main

import "testing"

func TestAbc(t *testing.T) {
	r := one()
	exp := 1
	if r != exp {
		t.Error() // to indicate test failed
	}
}
