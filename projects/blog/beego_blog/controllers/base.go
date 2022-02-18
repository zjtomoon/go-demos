package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

func (p *baseController) Prepare()  {
	controllerName,actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0:len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	if strings.ToLower(p.controllerName) == "admin" && strings.ToLower(p.actionName) != "login" {
		if p.GetSession("user") == nil {

		}
	}
}

func (p *baseController) History(msg string,url string)  {
	if url == "" {
		p.Ctx.WriteString("")
	}
}