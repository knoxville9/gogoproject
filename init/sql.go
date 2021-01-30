package init

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var db *sql.DB
var Dsn = "root:123123@tcp(localhost:3306)/dev_crm?charset=utf8mb4&parseTime=True&loc=Local"

func SqlInstance() {
	var err1 error

	Db, err1 = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "wd_",
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
