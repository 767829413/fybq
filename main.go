package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type People struct {
	Name string
	Age  int
}

func main() {
	var (
		p   People = People{Name: "xiaoming", Age: 15}
		buf bytes.Buffer
	)
	// 编码
	encode := gob.NewEncoder(&buf)
	err := encode.Encode(&p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("编码后的xiaoming %v\n", buf.Bytes())
	// 解码
	decode := gob.NewDecoder(bytes.NewReader(buf.Bytes()))

	var pp People
	err = decode.Decode(&pp)
	if err != nil {
		log.Fatal(err)
	}
	pp.Age = 100
	fmt.Printf("解码后的xiaoming %v\n", pp)
}
