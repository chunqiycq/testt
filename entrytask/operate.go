package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GlobalConn *gorm.DB

type Student struct {
	ID    uint
	Name  string
	Score int
}

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
	InsertData()
	SearchData()
	UpdateData()

	GlobalConn.Unscoped().Where("name = ?", "lisi").Delete(new(Student)) //硬删除方法
	// 借助 gorm 创建数据库表.
	//fmt.Println(db.AutoMigrate(new(Student)).Error)

}

func InsertData() {
	// 先创建数据 --- 创建对象
	var stu Student
	stu.Name = "zhangsan"
	stu.Score = 100

	// 插入(创建)数据
	fmt.Println(GlobalConn.Create(&stu).Error)
}

func SearchData() {
	var stu Student
	GlobalConn.First(&stu)
	fmt.Println(stu)
	//	var stu []Student                         // 改为切片
	//	GlobalConn.Select("name, age").Find(&stu) // Find() 查询多条
	//	fmt.Println(stu[1].Name)
}

func UpdateData() {
	var stu Student
	stu.Name = "wangwu"
	stu.Score = 77
	stu.ID = 4 //指定 id -- 更新操作!
	fmt.Println(GlobalConn.Save(&stu).Error)
}

//表的结构如下
//CREATE TABLE IF NOT EXISTS `BI`(
//	`Id` INT UNSIGNED AUTO_INCREMENT,
//	`Name` VARCHAR(100) NOT NULL,
//	`code` VARCHAR(40) NOT NULL,
//	PRIMARY KEY ( `Id` )
//)ENGINE=INNODB DEFAULT CHARSET=utf8;
