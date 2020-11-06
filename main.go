package main

import (
	_ "awesomeProject/docs"
	. "awesomeProject/init"
	. "awesomeProject/services"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	SqlInstance()

	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.Any("/query", FindRecommended)
	engine.Any("/team", Team)
	engine.POST("/link", Link)

	engine.Run()

}
