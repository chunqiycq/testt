package main

import (
	"fmt"
	"net"
	"os" //这里可以考虑使用缓冲写入的bufio？
)

func clientaddget(name, code string) {
	// 客户端连接服务器
	dial, err := net.Dial("tcp", "127.0.0.1:8060") //网络连接函数
	errFunction("net.Dial err：", err)
	defer dial.Close()
	if err != nil {
		fmt.Println("客户端未连接服务器")
	}
	fmt.Println("客户端成功连接服务器")
	// 模拟浏览器
	//fmt.Println("姓名密码：", "姓名密码：", name, code)
	requstHttpHeader := "POST /chunqi?nameyangc HTTP/1.1\r\nHost:127.0.0.1:8060\r\nContent-Type:application/x-www-form-urlencoded\r\nContent-Length: 125\r\n\r\n吕大强" //声明并且赋值
	fmt.Println(requstHttpHeader)

	// 给服务器发送请求报文
	dial.Write([]byte(requstHttpHeader))
	fmt.Fprintln(dial)
	buf := make([]byte, 1024) //这就是赋值空间
	fmt.Println("连接服务器1111")
	// 读取服务器的回复
	read, err := dial.Read(buf)
	errFunction("dial.Read err：", err)
	fmt.Println(string(buf[:read]))
	fmt.Println("连接服务器222")
}

func errFunction(describe string, err error) {
	if err != nil {
		fmt.Println(describe, err)
		os.Exit(001)
	}
}
