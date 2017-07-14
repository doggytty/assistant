package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type MainController struct {
	BaseController
}

func (m *MainController) Get() {
	log.Println("main get")
	m.TplName = "index.tpl"
	m.Data["NavType"] = "main"
}

















type DemoController struct {
	beego.Controller
}

func (c *DemoController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "demo.tpl"
}
