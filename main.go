package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	if len(b) != 0 {
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := cal(rune(string(b)[0]), []int{1, 2, 3, 4, 5})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "%d\n", res)
}

func cal(symbol rune, values []int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("values must be a non empty list")
	} else if len(values) == 1 {
		return values[0], nil
	} else if symbol == '+' {
		r, e := cal(symbol, values[1:])
		if e != nil {
			return 0, e
		}
		return values[0] + r, nil
	} else if symbol == '*' {
		r, e := cal(symbol, values[1:])
		if e != nil {
			return 0, e
		}
		return values[0] * r, nil
	}
	return 0, errors.New("something happen")
}
