package main

import (
	"fmt"
	orm "xyx/databases"
	"xyx/operation"
	"xyx/routes"
)

func main() {
	fmt.Println("Hello golang")
	/*r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello,gin")
		redis.SetCache("dd", "oo", time.Second*5)
		val, _ := redis.GetCache("dd")
		fmt.Println("rrr", val)
	})
	r.GET("/dd", func(c *gin.Context) {
		val, _ := redis.GetCache("dd")
		fmt.Println("rrr", val)
		c.String(200, val)
	})
	r.Run()*/
	//SetMaxIdleConns 是设置空闲时的最大连接数
	//SetMaxOpenConns 设置与数据库的最大打开连接数
	//SetConnMaxLifetime 每一个连接的生命周期等信息
	sqlDB, err := orm.Eloquent.DB()
	if err != nil {
		fmt.Println("数据库连接失败!!!")
		panic(err) //数据库
	}
	//延时调用函数
	defer sqlDB.Close()
	//go redis.ConnRedis()
	/*sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(-1)*/
	go operation.GetList()
	r := routes.OpenApi()
	r.Run(":8080")
}
