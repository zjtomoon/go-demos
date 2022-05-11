package dao

import (
	"context"
	"fmt"

	"gitee.com/chunanyong/zorm"
	_ "github.com/go-sql-driver/mysql"
	//_ "gitee.com/chunanyong/dm" // 达梦数据库驱动
)

//var user models.User
var db *zorm.DBDao
var ctx = context.Background()

func init() {
	var err error
	dbConfig := zorm.DataSourceConfig{
		// 连接数据库DSN
		DSN: "root:123456@tcp(127.0.0.1:3306)/document_server?charset=utf8&parseTime=true",
		// 数据库类型
		DriverName:            "mysql",
		DBType:                "mysql",
		PrintSQL:              true,
		MaxOpenConns:          50,
		MaxIdleConns:          50,
		ConnMaxLifetimeSecond: 0,
		DefaultTxOptions:      nil,
	}

	db, err = zorm.NewDBDao(&dbConfig)
	if err != nil {
		fmt.Errorf("数据库连接异常 %v", err)
		panic(err)
	}

	fmt.Println("数据库连接成功")
}
