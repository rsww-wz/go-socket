package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
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

		//接收数据大小，不能多收，否则会下面的数据会不完整
		var headerSize [4]byte

		n, err := conn.Read(headerSize[:])
		fmt.Println(n)
		if err != nil {
			fmt.Println("read failed err", err)
			return
		}
		str := string(headerSize[:n])
		fmt.Println("收到服务端回复:", str)

		// 解析真实数据
		contentSize, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("str不是数据类型", err)
			return
		}

		// 接收真实数据，用什么容器去接收，数组的大小不能是变量,直接字节切片即可
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

