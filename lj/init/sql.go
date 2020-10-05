package init

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var Dsn = "dev_crm:1q2w3e4R@tcp(rm-wz94wke3403mitc8mo.mysql.rds.aliyuncs.com:3306)/dev_crm?charset=utf8mb4&parseTime=True&loc=Local"

func SqlInstance() {
	var err1 error

	Db, err1 = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "wd_",
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	if err1 != nil {
		panic(err1)

	}

}
