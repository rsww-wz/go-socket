package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	client, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})

	if err != nil {
		fmt.Println("dail failed err:", err)
		return
	}
	defer client.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = client.Write([]byte(s))
		if err != nil {
			fmt.Println("send to server failed,err", err)
			return
		}

		// 接收数据
		var buf [1024]byte
		n, addr, err := client.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("recv from udp failed,err", err)
			return
		}
		fmt.Printf("read from %v,msg:%v\n", addr, string(buf[:n]))
	}
}
