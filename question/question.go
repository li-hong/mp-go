package question

import (
	"mp-go/wx"
	"mp-go/utils"
)

func Handler(msg wx.WxMpXmlInMessage) interface{} {
	textMsg := wx.TextBuilder()
	content:=utils.GetCityWeather("北京")
	textMsg.Content = content

	textMsg.FromUserName = msg.ToUserName
	textMsg.ToUserName = msg.FromUserName

	if msg.MsgType == wx.MSG_EVENT {
		switch msg.Event {
		case wx.EVENT_SUBSCRIBE:wx.UserSubscribe(msg.FromUserName)
		case wx.EVENT_UNSUBSCRIBE:wx.UserUnSubscribe(msg.FromUserName)
		}
	}

	return textMsg
}

