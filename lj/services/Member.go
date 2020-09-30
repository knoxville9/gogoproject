package services

import (
	. "awesomeProject/lj/domain"
	. "awesomeProject/lj/init"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//查询单个用户
func Query(c *gin.Context) {
	n := Member{}

	id := c.Param("id")
	Db.First(&n, id)

	c.JSON(200, n)
}

func Stats(c *gin.Context) {

	stats := SqlDB.Stats()

	marshal, err := json.Marshal(stats)

	if err != nil {
		panic(err)
	}

	c.String(200, string(marshal))
}
