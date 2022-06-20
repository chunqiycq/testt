package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GlobalConn *gorm.DB

func main() {
	//初始化
	//var err error
	//db, err = gorm.Open("mysql", "root:021541@tcp(127.0.0.1:3306)/sms?charset=utf8&parseTime=True&loc=Local")
	connArgs := "root:021541@tcp(127.0.0.1:3306)/sms?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connArgs)
	GlobalConn = db
	// 初始数
	GlobalConn.DB().SetMaxIdleConns(10)
	// 最大数
	GlobalConn.DB().SetMaxOpenConns(100)
	if err != nil {
		fmt.Println("err")
	}
	defer db.Close()

	db.SingularTable(true)
	// 借助 gorm 创建数据库表.
	err3 := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin").AutoMigrate(new(manage))
	if err3 == nil {
		fmt.Println("数据表创建失败")
	}
	var flagnum = 4
	switch flagnum {
	case 1:
		var name = "gyfgc"
		var code = "G123张"
		flag := InsertData(name, code)
		if flag == 0 {
			fmt.Println("缺少姓名密码")
		}
		if flag == 1 {
			fmt.Println("插入完成")
		}
		if flag == 2 {
			fmt.Println("名字重复，不能插入")
		}

	case 2:
		var ID = 1
		var name = "a"
		flag := Deletedata(ID, name) //0-未输入删除的数据 1-用户删除完成 2-用户不存在，请重新输入，3-存在删除错误，请检查
		if flag == 0 {
			fmt.Println("未输入删除的数据")
		}
		if flag == 1 {
			fmt.Println("用户删除完成")
		}
		if flag == 2 {
			fmt.Println("用户不存在，请重新输入")
		}
		if flag == 3 {
			fmt.Println("存在删除错误，请检查")
		}

	case 3:
		var ID = 1
		var namec = "c"
		var nameo = "A"
		var code = ""
		flag := Updatedata(ID, namec, nameo, code) //0-未输入删除的数据 1-用户删除完成 2-用户不存在，请重新输入，3-存在删除错误，请检查
		if flag == 0 {
			fmt.Println("未输入更新的数据")
		}
		if flag == 6 {
			fmt.Println("名字存在，无法更新")
		}
		if flag == 1 {
			fmt.Println("更新名字密码")
		}
		if flag == 2 {
			fmt.Println("更新名字")
		}
		if flag == 3 {
			fmt.Println("更新密码")
		}
		if flag == 4 {
			fmt.Println("用户不存在，请重新输入")
		}
		if flag == 5 {
			fmt.Println("存在更新错误，请检查")
		}

	case 4:
		var serchflag = 1
		var name = "杨c"
		var pagenum = 3
		//flag := 1
		if serchflag == 1 {
			mage, flag := SearchData(name) //0-未输入删除的数据 1-用户删除完成 2-用户不存在，请重新输入，3-存在删除错误，请检查
			if flag == 0 {
				fmt.Println("未输入查询的名字")
			}
			if flag == 1 {
				fmt.Println("查询的用户信息为：")
				fmt.Println(mage)
			}
			if flag == 2 {
				fmt.Println("用户不存在，请重新输入")
			}
		}
		if serchflag == 2 {
			mage, flag := PageSearch(pagenum)
			if flag == 0 {
				fmt.Println("未输入查询的页码")
			}
			if flag == 1 {
				fmt.Println("查询的用户信息为：")
				fmt.Println(mage)
			}
			if flag == 2 {
				fmt.Println("查询出错")
			}
		}

	}
}
