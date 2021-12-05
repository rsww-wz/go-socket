package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
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

		// 在服务端就解决掉编码问题
		output,_ = encoding(output)
		header := newHeader(string(output))
		header.completeHeader()
		headerBytes := header.headerBytes()
		if err != nil {
			fmt.Println(err)
			conn.Write([]byte(err.Error()))
			conn.Write(output)
		} else {
			fmt.Println(len(output))
			// 发送头
			conn.Write(headerBytes)
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

// 自定义头大小60bytes
type header struct {
	BodySize int
	TotalSize int
	Complete string
}

func newHeader(content string) *header{
	return &header {
		BodySize: len(content),
		TotalSize: len(content) + 60,
	}
}

func (h *header)completeHeader() {
	headerJson,_ := json.Marshal(*h)
	// 单位bytes
	if len(headerJson) > 60 {
		fmt.Println("自定义头已经超过60bytes了")
	} else {
		sub := 60 - len(headerJson)
		// 填充的是字符串，一个字符0刚好就对应一个byte
		for i:=0;i<sub;i++ {
			h.Complete += "0"
		}
	}
}

func (h *header) headerBytes() []byte{
	headerBytes,err := json.Marshal(*h)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return headerBytes
}
