package main

import (
	. "awesomeProject/app/service"
	. "awesomeProject/init"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	SqlInstance()

	engine := gin.Default()

	engine.Any("/query", FindRecommended)
	engine.Any("/team", Team)
	engine.POST("/link", Link)

	err := engine.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
