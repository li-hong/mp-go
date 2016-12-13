package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/xml"
	"mp-go/wx"
	"mp-go/question"
	"github.com/huichen/sego"
	"mp-go/utils"
)

var segmenter sego.Segmenter

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


	utils.SendSms("18910881431","hello your verfi code is 1234")



	// 分词
	text := []byte(msg.Content)
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	beego.Info(sego.SegmentsToString(segments, false))

	outMsg := question.Handler(msg)
	if nil != outMsg {
		obj, err := xml.Marshal(outMsg)
		if nil == err {
			c.Ctx.ResponseWriter.Write(obj)
		}
		return
	}

	beego.Info(msg)
	beego.Info(outMsg)

	c.Data["json"] = msg

	c.ServeXML()
}

func init() {
	segmenter.LoadDictionary("../github.com/huichen/sego/data/dictionary.txt")
}
