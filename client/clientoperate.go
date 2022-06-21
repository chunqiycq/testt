package main

import (
	//处理输入高并发用，可以将数据保存在缓冲区，之后再考虑？
	"fmt"
)

var opera client

func operate() {
	var model int
	fmt.Println("输入操作的功能：")
	//	fmt.Scan(&model)
	model = 1 //选择插入数据
	switch model {
	case 1:
		var Inname, Incode = input()
		opera.Nameo = Inname
		opera.Code = Incode
		opera.Flagnum = 1
		clientaddget(opera)
	}
}

func input() (name, passcode string) {
	fmt.Println("请输入你的用户名、用户密码：")
	//_, _ = fmt.Scan(&name, &passcode) // 输入一个后回车再输入下一个
	name = "有fv"
	passcode = "56三"
	fmt.Printf("用户名：%v\n用户密码：%v\n", name, passcode)
	return
}
