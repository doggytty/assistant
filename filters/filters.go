package filters

import (
	"github.com/astaxie/beego/context"
	"log"
	"github.com/doggytty/goutils/collections"
)

var ExcludeUrl = []string{"/login", "/logout"}

func FilterLogin(ctx *context.Context) {
	uid := ctx.Input.CruSession.Get("uid")
	currURI := ctx.Request.RequestURI
	log.Println(collections.Contains(ExcludeUrl, currURI))
	// 检查是否必须过滤的url
	if collections.Contains(ExcludeUrl, currURI) {
		log.Printf("url: %s\n", currURI)
	} else {
		log.Printf("not url: %s\n", currURI)
		if uid == nil {
			ctx.Redirect(302, "/login")
		} else {
			uidValue, ok := uid.(int)
			if !ok {
				ctx.Redirect(302, "/login")
			} else {
				log.Printf("now uid %d\n", uidValue)
			}
		}
	}
}
