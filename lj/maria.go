package main

import (
	. "awesomeProject/lj/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn string = "dev_ljyh:823j#8S9&33b48!@tcp(rm-wz9x70o6s2im669d0o.mysql.rds.aliyuncs.com:3306)/dev_ljyh?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB
var err error

func main() {
	//dsn := "root:123123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sq_",
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/users/:id", Query) //根据id获取用户
	r.GET("/stats", Stats)     //根据id获取用户

	_ = r.Run()

}
