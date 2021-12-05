package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 与服务端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed,err:", err)
		return
	}

	// 利用该链接获取数据，处理后，然后进行发送
	// 新建标准输入对象
	input := bufio.NewReader(os.Stdin)

	// 通信循环
	for {
		// 客户端输入
		s, _ := input.ReadString('\n')
		// 输入数据处理
		s = strings.TrimSpace(s)
		// 通信退出
		if strings.ToUpper(s) == "Q" {
			return
		}
		// 给服务端发送消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Println("send failed err", err)
			return
		}

		// 从服务端接收返回的消息
		var buf [1024]byte
		n,err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed err",err)
			return
		}
		fmt.Println("收到服务端回复:",string(buf[:n]))
	}
}
