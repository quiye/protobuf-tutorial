package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type formula struct {
	OpString string `json:"op"`
	Args     []int  `json:"args"`
	Op       rune
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("debug %v\n", r)
	fmt.Fprint(w, "Hello, World\n")
}

func main() {
	http.HandleFunc("/", calcHandler)
	http.ListenAndServe(":8080", nil)
}

func one() int {
	return 1
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	f := formula{}
	if err := json.Unmarshal(b, &f); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(f.OpString) != 1 {
		http.Error(w, err.Error(), 500)
		return
	}

	f.Op = rune(f.OpString[0])

	res, err := cal(f.Op, f.Args)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "%d\n", res)
}

func cal(op rune, values []int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("values must be a non empty list")
	} else if len(values) == 1 {
		return values[0], nil
	} else if op == '+' {
		r, e := cal(op, values[1:])
		if e != nil {
			return 0, e
		}
		return values[0] + r, nil
	} else if op == '*' {
		r, e := cal(op, values[1:])
		if e != nil {
			return 0, e
		}
		return values[0] * r, nil
	}
	return 0, errors.New("something happen")
}
