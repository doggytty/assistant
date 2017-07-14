package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	log.Println(l.GetSession("uid"))
	log.Println(l.Ctx.Input.CruSession.Get("uid"))
	l.TplName = "login.tpl"
}

func (l *LoginController) Post() {
	l.Data["Website"] = "beego.me"
	l.Data["Email"] = "astaxie@gmail.com"
	l.SetSession("uid", "110")
	l.Redirect("/index", 302)
}
