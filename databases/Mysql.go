package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	//用户名:密码@tcp(数据库ip或域名:端口)/数据库名称?charset=数据库编码&parseTime=True&loc=Local
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_xyx?charset=utf8mb4&parseTime=True&loc=Local"
	Eloquent, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
