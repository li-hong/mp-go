package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/xml"
	"mp-go/wx"
	"mp-go/question"
)

type WeixinMpController struct {
	beego.Controller
}

func (c *WeixinMpController) URLMapping() {
	c.Mapping("index", c.Index)
}


// @Title get,post
// @Description 处理微信请求
// @Param	echostr	query	string false
// @Success 200
// @router /index [get,post]
func (c *WeixinMpController) Index() {
	// signature := c.Ctx.Input.Param("signature");
	//nonce :=  c.Ctx.Input.Param("nonce");
	//timestamp := c.Ctx.Input.Param("timestamp");
	echostr := c.GetString("echostr");
	fmt.Println(echostr)
	if ("" != echostr) {
		c.Ctx.ResponseWriter.Write([]byte(echostr))
		return;
	}

	msg := wx.WxMpXmlInMessage{}
	xmlerr := xml.Unmarshal(c.Ctx.Input.RequestBody, &msg)
	if xmlerr != nil {
		beego.Error(xmlerr)
	}

	outMsg := question.Handler(msg)
	beego.Info(outMsg)

	if nil != outMsg {
		obj, err := xml.Marshal(outMsg)
		if nil == err {
			c.Ctx.ResponseWriter.Write(obj)
		}
		return
	}

	beego.Info(msg)

	c.Data["json"] = msg

	c.ServeXML()
}
