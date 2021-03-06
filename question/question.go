package question

import (
	"mp-go/wx"
	"mp-go/utils"
	"github.com/astaxie/beego"
)

func Handler(msg wx.WxMpXmlInMessage) interface{} {
	textMsg := wx.TextBuilder()
	city := "北京" //默认城市设为北京
	//openid := msg.FromUserName //查询用户地址

	////接收的是位置消息
	if (msg.MsgType == wx.MSG_LOCATION) {
		beego.Info(msg)
		dis := utils.GetLocationDistrict(msg.Lat, msg.Lng)
		if dis != "" {
			city = dis
		}
	}

	var content string
	content = utils.GetCityWeather(msg.Content)
	if content == "" {
		content = utils.GetCityWeather(city)
	}
	textMsg.Content = content

	textMsg.FromUserName = msg.ToUserName
	textMsg.ToUserName = msg.FromUserName

	//if msg.MsgType == wx.MSG_EVENT {
	//	switch msg.Event {
	//	case wx.EVENT_SUBSCRIBE:wx.UserSubscribe(msg.FromUserName)
	//	case wx.EVENT_UNSUBSCRIBE:wx.UserUnSubscribe(msg.FromUserName)
	//	}
	//}

	return textMsg
}

