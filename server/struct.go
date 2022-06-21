package main

type manage struct {
	//gorm.Model
	ID   uint
	Name string
	Code string
}

//表的结构如下
//CREATE TABLE IF NOT EXISTS `BI`(
//	`Id` INT UNSIGNED AUTO_INCREMENT,
//	`Name` VARCHAR(100) NOT NULL,
//	`code` VARCHAR(40) NOT NULL,
//	PRIMARY KEY ( `Id` )
//)ENGINE=INNODB DEFAULT CHARSET=utf8;
