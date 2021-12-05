package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net"
	"os/exec"
	"strconv"
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

		// 在服务端就解决掉编码问题
		output,_ = encoding(output)
		if err != nil {
			fmt.Println(err)
			fmt.Println(len([]byte(err.Error())))
			fmt.Println(len(output))
			conn.Write([]byte(err.Error()))
			conn.Write(output)
		} else {
			// 数据大小
			fmt.Println(len(output))
			// 因为间隔时间短，数据会一起发送
			// 发送数据大小
			size := strconv.Itoa(len(output))
			fmt.Println(size)
			conn.Write([]byte(size))
			// 发送真实数据
			conn.Write(output)
		}
	}
}

func encoding(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
