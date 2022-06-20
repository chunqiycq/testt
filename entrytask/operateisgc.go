package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func InsertData(name, code string) (flag int) {
	// 先创建数据 --- 创建对象
	var stu manage
	stu.Name = name
	stu.Code = code
	if name == "" || code == "" {
		return 0 //缺少姓名密码
	}
	// 插入(创建)数据,重复就不插入
	results := GlobalConn.Table("manage").Where("name = ?", name).First(&stu)
	if results.Error != nil {
		results := GlobalConn.Create(&stu)
		if results.Error == nil {
			flag = 1 //插入数据成功
		}
	}
	if results.Error == nil {
		flag = 2 //数据重复，无法插入
	}
	return
}

func Deletedata(ID int, name string) (flag int) {
	if ID == 0 && name == "" {
		return 0 //未输入删除的数据
	}
	stu := new(manage)
	if ID != 0 {
		results := GlobalConn.Where("ID = ?", ID).First(&stu) //查找存在需要删除的数据
		if results.Error != nil {
			return 2 //用户不存在，请重新输入
		}
		GlobalConn.Unscoped().Where("id = ?", ID).Delete(new(manage))
		return 1 //用户删除完成
	}
	if ID == 0 && name != "" {
		results := GlobalConn.Table("manage").Where("Name = ?", name).First(&stu)
		if results.Error != nil {
			return 2 //用户不存在，请重新输入
		}
		GlobalConn.Unscoped().Where("id = ?", stu.ID).Delete(new(manage))
		return 1 //用户删除完成
	}

	return 3 //存在删除错误，请检查
}

func Updatedata(ID int, namec, nameo, code string) (flag int) {
	if ID == 0 && namec == "" {
		return 0 //未输入更新的数据
	}
	stu := new(manage)
	if ID != 0 {
		results := GlobalConn.Where("ID = ?", ID).First(&stu) //查找存在需要更新的数据
		if results.Error != nil {
			return 4 //用户不存在，请重新输入
		}
		if code != "" && nameo != "" {
			results = GlobalConn.Table("manage").Where("name = ?", nameo).First(&stu)
			if results.Error != nil {
				GlobalConn.Model(&stu).Where("ID = ?", ID).Updates(manage{Name: nameo, Code: code})
				return 1 //更新名字密码
			}
			return 6 //名字存在，无法更新
		}
		if code == "" && nameo != "" {
			results = GlobalConn.Table("manage").Where("name = ?", nameo).First(&stu)
			if results.Error != nil {
				GlobalConn.Model(&stu).Where("ID = ?", ID).Updates(manage{Name: nameo})
				return 2 //更新名字
			}
			return 6 //名字存在，无法更新
		}
		if code != "" && nameo == "" {
			GlobalConn.Model(&stu).Where("ID = ?", ID).Updates(manage{Code: code})
			return 3 //更新密码
		}
		return 5 //存在更新错误，请检查
	}
	if ID == 0 && namec != "" {
		results := GlobalConn.Where("Name = ?", namec).First(&stu) //查找存在需要删除的数据
		if results.Error != nil {
			return 4 //用户不存在，请重新输入
		}
		if code != "" && nameo != "" {
			results = GlobalConn.Table("manage").Where("name = ?", nameo).First(&stu)
			if results.Error != nil {
				GlobalConn.Model(&stu).Where("ID = ?", stu.ID).Updates(manage{Name: nameo, Code: code})
				return 1 //更新名字密码
			}
			return 6 //名字存在，无法更新
		}
		if code == "" && nameo != "" {
			results := GlobalConn.Table("manage").Where("name = ?", nameo).First(&stu)
			if results.Error != nil {
				GlobalConn.Model(&stu).Where("ID = ?", stu.ID).Updates(manage{Name: nameo})
				return 2 //更新名字
			}
			return 6 //名字存在，无法更新
		}
		if code != "" && nameo == "" {
			GlobalConn.Model(&stu).Where("ID = ?", stu.ID).Updates(manage{Code: code})
			return 3 //更新密码
		}
		return 5 //存在更新错误，请检查
	}
	return 5 //存在删除错误，请检查
}

func SearchData(name string) (stu manage, flag int) {
	//	var stu manage
	if name == "" {
		flag = 0
		return
	}
	err := GlobalConn.Where("name = ?", name).First(&stu).Error
	if err != nil {
		flag = 2 //不存在该用户
		fmt.Println(stu)
		return
	}
	flag = 1 //该用户存在，已经输出
	return
	//	var stu []Student                         // 改为切片
	//	GlobalConn.Select("name, age").Find(&stu) // Find() 查询多条
	//	fmt.Println(stu[1].Name)
}

func PageSearch(pagenum int) (stu []manage, flag int) {
	//var stu []manage
	if pagenum == 0 {
		flag = 0
		return
	}
	err := GlobalConn.Scopes(Paginate(pagenum, 2)).Find(&stu).Error //可以设置查询的每个页面的数据数目
	if err != nil {
		flag = 2 //查询超出页面范围，无结果
		//fmt.Println(stu)
		return
	}
	flag = 1
	return
}

//进行了分页分装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
