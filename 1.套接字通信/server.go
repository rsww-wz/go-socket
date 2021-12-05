/*
TCP服务程序的处理流程
	监听端口
	接收客户端请求，建立链接
	创建goroutine处理链接

API
	监听端口:func Listen(network, address string) (Listener, error)

TCP 传输数据的时候是已字节的方式传输，所以需要转换成byte类型
无论是服务端还是客户端，接收数据需要确定的大小的容器，也只能是字节数组了
但是reader和writer对象，都是需要字节切片作为参数
字节数组转换成字节切片:的字节数组切片即可，从头且到尾
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// 开启服务，规定使用的协议和监听的端口
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 链接循环
	for {
		// 等待客户端的请求，建立链接，返回请求对象和错误
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err", err)
			continue
		}
		// 处理客户端请求对象
		go process(conn)
	}
}

// 处理函数
func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()

	// conn对象已经包含了客户端发送过来的数据
	fmt.Printf("%v\n%T\n", conn, conn)

	// 读取数据，返回reader对象，默认读4K大小，也就是4096byte
	recv := bufio.NewReader(conn)

	// 定义接收器容量大小
	var buf [128]byte

	// 读取数据
	n, err := recv.Read(buf[:])
	if err != nil {
		fmt.Println("read from client failed err", err)
	}

	// 转换成字符串类型，处理数据
	recvStr := string(buf[:n])
	fmt.Println("收到client端发送来数据", recvStr)
	recvStr = strings.ToUpper(recvStr)

	// 转换成byte类型，发送数据
	conn.Write([]byte(recvStr))
}
