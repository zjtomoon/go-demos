package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-demos/study/beego_study/lesson2/models"
	_ "go-demos/study/beego_study/lesson2/routers"
)

func init() {
	//注册数据库
	models.RegisterDB()
}

func main() {
	//开启ORM调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	beego.Run()
}
