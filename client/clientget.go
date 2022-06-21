package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os" //这里可以考虑使用缓冲写入的bufio？
)

type userInfo struct {
	Name string
	Code string
}

func clientaddget(opera client) {
	// 客户端连接服务器
	dial, err := net.Dial("tcp", "127.0.0.1:8060") //网络连接函数,tcp连接

	errFunction("net.Dial err:", err)
	defer dial.Close()
	if err != nil {
		fmt.Println("客户端未连接服务器", err)
	}
	fmt.Println("客户端成功连接服务器")
	// 模拟浏览器
	url2 := "http://127.0.0.1:8060/chunqi"
	//var user userInfo
	//user.Name = name
	//user.Code = code
	//body := "{" + "\"flagnum\":1," + "\"name\":" + "\"" + name + "\"" + ",\"code\":" + "\"" + code + "\"" + "}"
	body := tojson(opera)
	resp, err2 := http.Post(url2, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	fmt.Println(body)
	fmt.Println(bytes.NewBuffer([]byte(body)))
	if err2 != nil {
		fmt.Println("请求报文出错", err2)
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))

	fmt.Println("连接服务器222")
}

func errFunction(describe string, err error) {
	if err != nil {
		fmt.Println(describe, err)
		os.Exit(001)
	}
}
