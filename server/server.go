package main

import (
	"fmt"
	"io/ioutil"
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

	//fmt.Println("Header：", r.Header)
	//fmt.Println("URL：", r.URL)
	//fmt.Println("Method：", r.Method)
	//fmt.Println("Host：", r.Host)
	//fmt.Println("RemoteAddr：", r.RemoteAddr)
	//fmt.Println("报文内容", r)
	s, _ := ioutil.ReadAll(r.Body) //把	body 内容读入字符串 s
	inrequire := string(s)
	fmt.Println(inrequire) //在返回页面中显示内容。
	opera := toencoding(inrequire)
	result := operamysql(opera)
	rec := mtojson(result)
	w.Write([]byte(rec)) //这是响应
	//fmt.Println(opera)
	//fmt.Println(opera.Name)
	//operatojson := tojson(opera)
	//fmt.Println(operatojson)
	//fmt.Println(inrequire) //在返回页面中显示内容。
	//rec := toencoding(string(s))
	//fmt.Println("json:", string(s))
	//fmt.Println("json2:", rec)
	//ress := tojson(rec)
	//fmt.Println("json3:", ress)

	//lenth := r.ContentLength
	//body := make([]byte, lenth)
	//r.Body.Read(body)
	//inrequire := bytes.NewBuffer([]byte(body)) //获取的post请求
	//fmt.Println("Body:", inrequire)
	//
}
