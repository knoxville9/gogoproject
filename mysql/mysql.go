package main

import (
	. "awesomeProject/domain"
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	SqlInstance()

	for i := 10000; i <= 20000; i++ {
		sign := MemberAddress{ID: 0, MemberID: i, Realname: strconv.Itoa(i), Mobile: strconv.Itoa(i), Province: "湖南省", City: "长沙市", Area: "雨花区", Address: "金鸿宇", IsDefault: 1, AddCode: 0, Addtime: time.Now(), Lastupdatetime: time.Now()}
		Db.Create(&sign)

	}

}

var Db *gorm.DB
var db *sql.DB
var Dsn = "ljyh_test:1q2w3e4r!@!@tcp(rm-wz9x70o6s2im669d0o.mysql.rds.aliyuncs.com:3306)/ljyh_test?charset=utf8mb4&parseTime=True&loc=Local"

func SqlInstance() {
	var err1 error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond * 50, // 慢 SQL 阈值
			LogLevel:      logger.Silent,         // Log level
			Colorful:      false,                 // 禁用彩色打印
		},
	)

	Db, err1 = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		Logger:      newLogger,
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
