package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"] = append(beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"] = append(beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"] = append(beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"] = append(beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"] = append(beego.GlobalControllerRouter["mp-go/controllers:WechatUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["mp-go/controllers:WeixinMpController"] = append(beego.GlobalControllerRouter["mp-go/controllers:WeixinMpController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/index`,
			AllowHTTPMethods: []string{"get","post"},
			Params: nil})

}
