package init

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var SqlDB *sql.DB
var Dsn = "root:123123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func SqlInstance() {
	var err1 error
	var err2 error

	Db, err1 = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sq_",
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	SqlDB, err2 = Db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	SqlDB.SetMaxIdleConns(100)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDB.SetMaxOpenConns(5000)

	if err1 != nil && err2 != nil {
		panic(err1)
		panic(err2)

	}

	defer SqlDB.Close()
}
