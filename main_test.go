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
