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

		// 由于处理函数本身就是并发的，即使上一个请求没有退出，也不会被某一个连接一直占用
		// 连接都是以参数的方式传递给处理函数的
		// 只有处理完当前请求，就能处理下一条请求
		// 如果同一时间有多个请求，相当于有一个请求队列，这个并发携程只能选择一个连接进行处理，处理完之后在处理其他请求
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