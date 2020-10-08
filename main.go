package main

import (
	_ "awesomeProject/docs"
	. "awesomeProject/handler"
	. "awesomeProject/init"
	. "awesomeProject/services"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath /

func main() {

	SqlInstance()

	r := gin.New()
	//r.NoMethod(HandleNotFound)
	//r.NoRoute(HandleNotFound)
	r.Use(Recover)
	r.GET("/tid", FindRecommended)

	r.GET("/team", Team)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/test", Test)

	//测试时下面两个中间件选择一个，注释一个
	r.Use(CustomRouterMiddle1)

	r.Run()

}

func CustomRouterMiddle1(c *gin.Context) {
	fmt.Println("abc")

}
