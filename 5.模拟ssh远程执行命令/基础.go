package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
)

func main() {
	// 通过exec.Command函数执行命令或者shell
	// 第一个参数是命令路径，当然如果PATH路径可以搜索到命令，可以不用输入完整的路径
	// 第二到第N个参数是命令的参数
	// 下面语句等价于执行命令: ls -l /var/
	//cmd := exec.Command("/bin/ls", "-l", "/var/")

	//func Command(name string, arg ...string) *Cmd
	cmd := exec.Command("cmd", "/c", "ping")

	// 执行命令，并返回结果
	//func (c *Cmd) Output() ([]byte, error)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	// 因为结果是字节数组，需要转换成string
	//因为Go的编码时UTF-8，而cmd的活动页是cp936(GBK)
	str, _ := GbkToUtf8(output)
	fmt.Println(string(str))
}

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	//func NewReader(r io.Reader, t Transformer) *Reader
	//Transformer 是一个接口
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func run() {
	cmd := exec.Command("ipconfig")

	// 执行命令，返回命令是否执行成功
	err := cmd.Run()

	if err != nil {
		// 命令执行失败
		panic(err)
	}
}
