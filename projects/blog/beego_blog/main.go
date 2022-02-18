package main

import (
	_ "beego_blog_exercise/routers"
	"github.com/astaxie/beego"
)

func init()  {
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

