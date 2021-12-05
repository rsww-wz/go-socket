package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(len([]byte("0")))
	a := rune(100)
	b := rune(1000)
	fmt.Printf("%T\t%v\t%d\n", a, a, len(string(a)))
	fmt.Println([]byte(string(a)), len([]byte(string(a))))
	fmt.Println([]byte(string(b)), len([]byte(string(b))))

	var c, d int64
	c = 1
	d = 100000
	fmt.Println([]byte(string(c)), len([]byte(string(c))))
	fmt.Println([]byte(string(d)), len([]byte(string(d))))

	header := newHead("hello world")
	data := structToByte(header)
	fmt.Println(data)
}

type head struct {
	length int
}

func newHead(str string) head{
	return head{
		length: len(str),
	}
}

func structToByte(header head) []byte{
	jsonHeader,_ := json.Marshal(header)
	fmt.Printf("%T\t%v\n",jsonHeader,jsonHeader)
	return jsonHeader
}
