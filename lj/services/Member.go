package services

import (
	. "awesomeProject/lj/domain"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dsn string = "root:123123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB
var err error

func Query(c *gin.Context) {

	//dsn := "root:123123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := mysqlInstance()
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	n := Member{}

	id := c.Param("id")
	db.First(&n, id)

	c.JSON(200, n)
}

func mysqlInstance() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sq_",
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	return db, err
}
func Stats(c *gin.Context) {

	//dsn := "root:123123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := mysqlInstance()
	sqlDB, err := db.DB()

	defer sqlDB.Close()

	if err != nil {
		panic(err)
	}

	stats := sqlDB.Stats()

	marshal, err := json.Marshal(stats)

	if err != nil {
		panic(err)
	}

	c.String(200, string(marshal))
}
