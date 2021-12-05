package main

import (
	"bufio"
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

	for {
		var s string
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

		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed err", err)
			return
		}
		fmt.Println("收到服务端回复:", string(buf[:n]))
	}
}
