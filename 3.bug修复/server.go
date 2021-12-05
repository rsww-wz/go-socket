package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err",err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	// 通信循环
	for {
		recv := bufio.NewReader(conn)

		var buf [128]byte
		n,err := recv.Read(buf[:])
		if err != nil {
			// go语言能单独处理错误，即使客户端中断连接，也不会导致服务端卡死报错
			// 退出循环，会自动结束这个线程
			fmt.Println("read from client failed err",err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发送来数据",recvStr)
		recvStr = strings.ToUpper(recvStr)

		conn.Write([]byte(recvStr))
	}
}