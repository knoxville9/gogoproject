package main

import (
	. "awesomeProject/lj/init"
	. "awesomeProject/lj/services"
	"github.com/gin-gonic/gin"
)

func main() {

	SqlInstance()

	r := gin.Default()

	r.GET("/query", Query)

	r.Run()

}
