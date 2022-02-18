package routers

import (
	"github.com/astaxie/beego"
	"go-demos/study/beego_study/lesson1/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
