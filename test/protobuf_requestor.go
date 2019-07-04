package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
	fi, err := os.Create("dat")
	defer fi.Close()
	if err != nil {
		return
	}
	err = binary.Write(fi, binary.BigEndian, b)
	if err != nil {
		return
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
