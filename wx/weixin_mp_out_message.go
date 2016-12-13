package wx

import (
	"time"
	"encoding/xml"
)

type WxMpXmlOutMessage struct {
	XMLName      xml.Name `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
}

type WxMpXmlOutTextMessage struct {
	WxMpXmlOutMessage
	Content string `xml:"Content"`
}

func TextBuilder() WxMpXmlOutTextMessage {
	textMsg := WxMpXmlOutTextMessage{}
	textMsg.MsgType = "text"
	textMsg.CreateTime = time.Now().Unix()
	return textMsg
}



