package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("start...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err", err)
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
		n, err := recv.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed err", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发送来数据", recvStr)
		str := strings.Split(recvStr, " ")
		fmt.Println(str)

		// 指令windows指令 : cmd := exec.Command("cmd", "/c", "command", "路径")
		var cmd *exec.Cmd
		if len(str) == 1 {
			cmd = exec.Command("cmd", "/c", str[0])
		} else {
			cmd = exec.Command("cmd", "/c", str[0], str[1])
		}

		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte(err.Error()))
			conn.Write(output)
		} else {
			conn.Write(output)
		}
	}
}
