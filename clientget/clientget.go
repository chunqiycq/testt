package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 客户端连接服务器
	dial, err := net.Dial("tcp", "127.0.0.1:8000")
	errFunction("net.Dial err：", err)
	defer dial.Close()

	// 模拟浏览器
	requstHttpHeader := "GET /itzhuzhu HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"

	// 给服务器发送请求报文
	dial.Write([]byte(requstHttpHeader))

	buf := make([]byte, 1024)

	// 读取服务器的回复
	read, err := dial.Read(buf)
	errFunction("dial.Read err：", err)
	fmt.Println(string(buf[:read]))
}

func errFunction(describe string, err error) {
	if err != nil {
		fmt.Println(describe, err)
		os.Exit(1)
	}
}
