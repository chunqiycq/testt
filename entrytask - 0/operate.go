package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:021541@tcp(127.0.0.1:3306)/sms?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed to connect database,", err)
		panic("failed to connect database")
	}
	defer db.Close()
	//设置全局表名禁用复数
	db.SingularTable(true)
}

type Test struct {
	//gorm.Model
	Id   int64  `gorm:"type:int(20);column:id;primary_key"`
	Name string `gorm:"type:varchar(100);column:name"`
	Code string `gorm:"type:varchar(40);column:code"`
}

func main() {
	//db.AutoMigrate(&Test{})
	bi := &Test{
		Id:   24,
		Name: "jackdie",
		Code: "1802d0",
	}
	db.Create(bi)
	biq := &Test{
		Id:   12,
		Name: "jackie",
		Code: "18020",
	}
	fmt.Println("open mysql failed2456,")
	db.Create(biq)
	var testResult Test
	db.Where("Id = ?", "12").First(&testResult)
	fmt.Println("result: ", testResult)
}

//表的结构如下
//CREATE TABLE IF NOT EXISTS `BI`(
//	`Id` INT UNSIGNED AUTO_INCREMENT,
//	`Name` VARCHAR(100) NOT NULL,
//	`code` VARCHAR(40) NOT NULL,
//	PRIMARY KEY ( `Id` )
//)ENGINE=INNODB DEFAULT CHARSET=utf8;
