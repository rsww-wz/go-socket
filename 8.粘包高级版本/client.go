package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}

	input := bufio.NewReader(os.Stdin)

	// 通信循环
	for {
		var s string
		// 判断是否是空字符串
		for {
			fmt.Print(">>>:")
			s, _ = input.ReadString('\n')
			s = strings.TrimSpace(s)
			if s != "" {
				break
			}
		}
		if strings.ToUpper(s) == "Q" {
			return
		}
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("send failed err", err)
			return
		}

		// 接收头
		var headerSize [60]byte

		n, err := conn.Read(headerSize[:])
		fmt.Println(n)
		if err != nil {
			fmt.Println("read failed err", err)
			return
		}
		headerBytes := headerSize[:n]

		// 解析真实数据
		var header data
		err = json.Unmarshal(headerBytes,&header)
		if err != nil{
			fmt.Println(err)
			return
		}
		contentSize := header.BodySize

		var contentBuf = make([]byte,contentSize)
		n, err = conn.Read(contentBuf[:])
		if err != nil {
			fmt.Println("read failed err", err)
			return
		}
		content:= string(contentBuf[:n])
		fmt.Println("收到服务端回复:", content)
	}
}

type data struct {
	BodySize int
	TotalSize int
	Complete string
}

