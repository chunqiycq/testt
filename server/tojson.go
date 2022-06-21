package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json
//1，把go中的结构体变量--->json字符串
//2。把json字符串--->go语言中能够识别的结构体
type client struct {
	Flagnum int    `json:"flagnum"` //1-增加，2-查询，3-更新，4-根据用户查询，5-指定页面查询
	ID      string `json:"id"`
	Namec   string `json:"namec"` //仅仅更新用到的更新名字
	Nameo   string `json:"nameo"` //需要更改删除或者新增查询的名字
	Code    string `json:"code"`
	Page    int    `json:"page"`
	Flag    int    `json:"flag"`
}

func tojson(p1 client) (rec string) {
	//将p1转json
	b, err := json.Marshal(p1) //序列号
	if err != nil {
		//	fmt.Printf("marshal fail,err:%v", err)
		return
	}
	rec = string(b)
	fmt.Printf(string(b), err)
	return
}

func toencoding(str string) (rec client) {
	//反序列话
	json.Unmarshal([]byte(str), &rec) //传指针是为了能在函数内部修改p2
	//fmt.Printf("%#v\n", rec)
	return
}
