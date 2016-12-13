package utils

import "github.com/astaxie/beego"

const sms_url  = "http://yl.mobsms.net/send/gsend.aspx"

func SendSms(phone string, msg string) (res string, err error) {
	param :=make(map[string]string)
	param["name"]="zawh"
	param["pwd"]="za123"
	param["dst"]=phone
	param["msg"]=msg
	data, err := PostWithProxy(sms_url, param)
	res = string(data)
	beego.Info(res)
	return
}
