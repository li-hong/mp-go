package wx

import (
	"mp-go/models"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"mp-go/utils"
	"github.com/google/uuid"
	"time"
)

const (
	MSG_EVENT = "event"
	MSG_LOCATION = "location"
	EVENT_SUBSCRIBE = "subscribe"
	EVENT_UNSUBSCRIBE = "unsubscribe"
)

type WxMpUserResponse struct {
	City          string    `json:"city"`
	Country       string    `json:"country"`
	Groupid       int       `json:"groupid"`
	Headimgurl    string    `json:"headimgurl"`
	Language      string    `json:"language"`
	Nickname      string    `json:"nickname"`
	Openid        string    `json:"openid"`
	Province      string    `json:"province"`
	Sex           int    `json:"sex"`
	Subscribe     int       `json:"subscribe"`
	SubscribeTime int       `json:"subscribe_time"`
	Unionid       string    `json:"unionid"`
}

func MpUserInfo(openid string) (user *models.WechatUser, err error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?openid=%s&access_token=%s", openid, GetAccessToken())

	data, err := utils.Get(url)
	if err != nil {
		beego.Error(err)
		return
	}
	mpUser := WxMpUserResponse{}

	if err := json.Unmarshal(data, &mpUser); nil == err {
		user = &models.WechatUser{}
		user.Country = mpUser.Country
		user.Province = mpUser.Province
		user.City = mpUser.City
		user.Nickname = mpUser.Nickname
		user.Headimgurl = mpUser.Headimgurl
		user.Unionid = mpUser.Openid
		user.Openid = mpUser.Openid
		user.Groupid = mpUser.Groupid
		user.Language = mpUser.Language
		user.Subscribe = mpUser.Subscribe
		user.SubscribeTime = mpUser.SubscribeTime
		switch mpUser.Sex {
		case 0:
			user.Sex = "未知"
		case 1:
			user.Sex = "男"
		case 2:
			user.Sex = "女"
		}
		return user, err
	}
	return
}

func UserSubscribe(openId string) {
	user, err := models.GetWechatUserByOpenid(openId)
	if nil != err {
		beego.Error(err)
	}
	if nil == user {
		user, err = MpUserInfo(openId)
		beego.Error(err)
		if nil == err && nil != user {
			user.Id = uuid.New().String()
			user.CreateTime = time.Now()
			id, errIns := models.AddWechatUser(user)
			if nil != errIns {
				beego.Error(err)
			}
			beego.Info(id)
		}
	}
}

func UserUnSubscribe(openId string) {
	user, err := models.GetWechatUserByOpenid(openId)
	if nil != err {
		beego.Error(err)
	}
	//设置用户关注状态不关注
	if nil != user {
		user.Subscribe = 0
		user.LastModifyTime = time.Now()
		models.UpdateWechatUserById(user)

	}
}
