package main

import (
	"fmt"
	"gin-blog/controller"
	"gin-blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:passwd@tcp(localhost:3306)/ginblog?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		fmt.Println(err)
	}
	//加载静态文件
	router.Static("/static/", "./static")

	//加载模板
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)           //首页
	router.GET("/category/", controller.CategoryList) //分类列表页面
	_ = router.Run(":8001")
}
