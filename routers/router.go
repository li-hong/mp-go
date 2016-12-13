package routers

import (
	"mp-go/controllers"
	"github.com/astaxie/beego"
)


func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/weixin",
			beego.NSInclude(
				&controllers.WeixinMpController{},
			),
		),
		beego.NSNamespace("/weixin/user",
			beego.NSInclude(
				&controllers.WechatUserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}



