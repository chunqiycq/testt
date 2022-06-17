package main

import (
	"fmt"
	"net/http"
)

func main() {
	/**
	注册回调函数，该回调函数会在服务器被访问时自动被调用
	func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
		pattern：访问服务器文件位置
		handler：回调函数名，函数必须是ResponseWriter, *Request类型作为参数
	*/
	http.HandleFunc("/chunqi", myHandlerFunc)

	/**
	绑定服务器监听地址
	func ListenAndServe(addr string, handler Handler) error
		addr：要监听的地址
		handler：回调函数，为空则调用系统默认的回调函数
	*/
	http.ListenAndServe("127.0.0.1:8060", nil)
}

/**
ResponseWriter：写给客户端的数据内容
Request：从客户端读取到的数据内容
*/
func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("需要4Test需要2Test"))
	fmt.Println("Header：", r.Header)
	fmt.Println("URL：", r.URL)
	fmt.Println("Method：", r.Method)
	fmt.Println("Host：", r.Host)
	fmt.Println("RemoteAddr：", r.RemoteAddr)
	fmt.Println("Body：", r.Body)
	fmt.Println("报文内容", r)
}
