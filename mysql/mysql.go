package main

import (
	. "awesomeProject/domain"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

func main() {
	SqlInstance()

	for i := 10000; i < 20000; i++ {
		sign := MemberLoginSign{ID: 0, Limittime: 2595496147, Logintime: 1595496147, Mid: i, Ticket: strconv.Itoa(i)}
		Db.Create(&sign)

	}

}

var Db *gorm.DB
var db *sql.DB
var Dsn = "ljyh_test:1q2w3e4r!@!@tcp(rm-wz9x70o6s2im669d0o.mysql.rds.aliyuncs.com:3306)/ljyh_test?charset=utf8mb4&parseTime=True&loc=Local"

func SqlInstance() {
	var err1 error

	Db, err1 = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sq_",
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	db, err2 := Db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)

	if err1 != nil {
		panic(err1)

	}
	if err2 != nil {
		panic(err2)

	}

}
