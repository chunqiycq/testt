package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/widuu/gojson"
)

var (
	dbhostsip  = "young"
	dbusername = "yangc"
	dbpassowrd = "021541"
	dbname     = "sms"
)

type mysql_db struct {
	db *sql.DB //定义结构体
}

func (f *mysql_db) mysql_open() { //打开
	Odb, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		fmt.Println("链接失败")
	}
	fmt.Println("链接数据库成功...........已经打开1")
	f.db = Odb
}

func (f *mysql_db) mysql_close() { //关闭
	defer f.db.Close()
	fmt.Println("链接数据库成功...........已经关闭")
}

func (f *mysql_db) mysql_select(sql_data string) {
	rows, err := f.db.Query(sql_data)
	if err != nil {
		println(err)
	}
	for rows.Next() {
		var in_param string

		err = rows.Scan(&in_param)
		if err != nil {
			panic(err)
		}

		user_id := gojson.Json(in_param).Get("user_id").Tostring()
		fmt.Println(user_id)

	}
}

func main() {
	db := &mysql_db{}
	db.mysql_open()
	//db.mysql_select("SELECT in_param FROM t_rong_credit_bank a WHERE a.method_code IN ('010402')  limit 10 ")
	db.mysql_close() //关闭
}
