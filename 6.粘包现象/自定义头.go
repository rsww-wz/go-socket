package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	head := newHeader("hello world")
	jsonByte := head.headerBytes()
	fmt.Println(jsonByte)
	fmt.Println(head.byteToJson(jsonByte))
	head.completeHeader()
	jsonByte = head.headerBytes()
	fmt.Println(head.byteToJson(jsonByte))

}

// 自定义头，规定头的大小和格式,规定头大小=60bytes
// 注意，成员需要对外可见才能读写
type header struct {
	BodySize int
	TotalSize int
	Complete string
}

func newHeader(content string) *header{
	return &header {
		BodySize: len(content),
		TotalSize: len(content) + 60,
	}
}

func (h *header)completeHeader() {
	headerJson,_ := json.Marshal(*h)
	if len(headerJson) > 60 {
		fmt.Println("自定义头已经超过60bytes了")
	} else {
		sub := 60 - len(headerJson)
		for i:=0;i<sub;i++ {
			h.Complete += "0"
		}
	}
}

func (h *header) headerBytes() []byte{
	headerBytes,err := json.Marshal(*h)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return headerBytes
}

func (h *header) byteToJson(jsonByte []byte) string {
	return string(jsonByte)
}
