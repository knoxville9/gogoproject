package main

import (
	. "awesomeProject/lj/init"
	. "awesomeProject/lj/services"
	"github.com/gin-gonic/gin"
)

func main() {

	SqlInstance()
	r := gin.Default()
	r.GET("/users/:id", Query) //根据id获取用户
	r.GET("/stats", Stats)     //根据id获取用户

	_ = r.Run()

}
