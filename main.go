package main

import (
	_ "awesomeProject/docs"
	. "awesomeProject/init"
	. "awesomeProject/services"
	"github.com/gin-contrib/gzip"
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

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestSpeed))
	r.GET("/tid", FindRecommended)

	r.GET("/team", Team)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":80")

}
