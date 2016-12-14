package wx

type WxMpXmlInMessage struct {
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
	Event        string   `xml:"Event"`
	EventKey     string   `xml:"EventKey"`
	Lat          string   `xml:"Location_X"`
	Lng          string   `xml:"Location_Y"`
}

