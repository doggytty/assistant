package routers

import (
	"github.com/doggytty/assistant/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/demo", &controllers.DemoController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/index", &controllers.MainController{})
}
