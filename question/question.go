package question

import (
	"mp-go/wx"
	"mp-go/utils"
)

func Handler(msg wx.WxMpXmlInMessage) interface{} {
	textMsg := wx.TextBuilder()
	city := "北京" //默认城市设为北京
	openid := msg.FromUserName //查询用户地址

	user, err := wx.MpUserInfo(openid)
	if err == nil {
		if (user.City != "") {
			city = user.City
		}
	}

	content := utils.GetCityWeather(city)
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

