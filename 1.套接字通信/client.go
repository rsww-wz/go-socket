/*
客户端发起请求:拨号
	func Dial(network, address string) (Conn, error)

客户端实现与服务端的通信很简单
主要就是三个函数
	拨号:func Dial(network, address string) (Conn, error)
	接收数据 : conn.read()
	发送数据 : conn.write()

主要还是对reader对象的操作，操作的都是字节切片，要熟悉类型转换

客户端流程:发送请求——输入数据——发送数据(转换类型)——接收数据(类型转换)——退出
	拨号——发送——接收——退出
*/

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
	// 客户端输入
	s, _ := input.ReadString('\n')

	// 通信退出
	if strings.ToUpper(s) == "Q" {
		return
	}

	// 转换成byte类型，给服务端发送消息
	_, err = conn.Write([]byte(s))
	if err != nil {
		fmt.Println("send failed err", err)
		return
	}

	// 定义接收器大小
	var buf [1024]byte

	// 从服务端接收返回的消息
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("read failed err", err)
		return
	}
	fmt.Println("收到服务端回复:", string(buf[:n]))
}
