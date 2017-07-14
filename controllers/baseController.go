package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
)

type BaseController struct {
	beego.Controller
}

func (base *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{})  {
	// 调用默认初始化设置
	base.Controller.Init(ctx, controllerName, actionName, app)
	// 自定义初始化设置
	log.Println("base controller Init")
}


func (base *BaseController) Prepare() {
	log.Println("base controller Prepare")
	base.Layout = "layout/base.tpl"
	base.LayoutSections = make(map[string]string)
	base.LayoutSections["Header"] = "layout/header.tpl"
	base.LayoutSections["Scripts"] = "layout/script.tpl"
	base.LayoutSections["Styles"] = "layout/styles.tpl"
	base.LayoutSections["Footer"] = "layout/footer.tpl"
	base.LayoutSections["Navigation"] = "layout/navigation.tpl"


	//if ctx.BeegoInput.Query("_method")!="" && ctx.BeegoInput.IsPost(){
	//this.Ctx.Input.IsPost()
	//log.Println(this.Ctx.Input.IsPost())
	//处理uid,username

	//tmpId := base.Ctx.Input.CruSession.Get("uid")
	//if tmpId == nil {
	//	return
	//}
	//uid := tmpId.(int64)
	//if uid > 0 {
	//	if this.Ctx.Input.IsAjax() {
	//		log.Println("ajax request!")
	//	} else {
	//		user := m.GetUserById(uid)
	//		this.Data["uid"] = uid
	//		this.Data["username"] = user.Username
	//	}
	//}
	//log.Println(this.Ctx.Request)
}


func (base *BaseController) Rsp(result bool, err string) {
	base.Data["json"] = map[string]interface{}{"result": result, "msg": err}
	base.ServeJSON()
}

func (base *BaseController) Noauth()  {
	if base.IsAjax() {
		base.Rsp(false, "no auth to access!")
	} else {
		base.TplName = "noauth.tpl"
	}
}