package controllers

import (
	"beego_study/lesson3/models"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	//检查是否有操作
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 302)
		return
	}
	this.Data["isActive"] = "category"
	this.TplName = "category.html"
	this.Data["isLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Category"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
		//logs.Error(err)
	}
}
