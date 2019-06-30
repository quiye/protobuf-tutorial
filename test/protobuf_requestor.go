package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
	pb "github.com/quiye/protobuf-tutorial/api"
)

func main() {
	f := &pb.Formula{
		OpString: "*",
		Values:   []int32{1, 2, 3, 6},
	}

	b, err := proto.Marshal(f)
	if err != nil {
		fmt.Printf("%v", err)
	}
	reqdata := bytes.NewReader(b)

	resp, err := http.Post("http://localhost:8080/calcpb", "", reqdata)
	if err != nil {
		fmt.Printf("%v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%s", string(body))
}
