/*
客户端如果直接输入回车，要发送的数据是空字符串，操作系统不会发送数据
但是程序已经过了输入数据的那段程序，所以客户端会一直卡在发送数据那段，但是又没有数据发送
所以需要判断，如果是空字符串，则不让程序进行下去

而服务端由于没有接收到数据，不会有任何操作
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
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

	// 通信循环
	for {
		var s string
		// 判断是否是空字符串
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
		str,_ := encoding(buf[:n])
		fmt.Println("收到服务端回复:", string(str))
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
