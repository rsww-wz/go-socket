/*
不用建立链接，没有粘包现象
简单快捷，效率高，但是不安全，是不可靠的传输协议，一般用在实时通信，视频直播等领域

由于没有建立链接，所以客户端发来的请求需要保留链接地址发送消息

API
	建立链接：net.ListenUDP("udp", &net.UDPAddr{})
	读取数据：listen.ReadFromUDP(data[:n])
	发送数据：listen.WriteToUDP(data[:n], addr)
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("listen failed err:", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed err:", err)
			continue
		}
	}
}
