package wx

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"mp-go/utils"
	"mp-go/cache"
	"encoding/json"
	"time"
	"github.com/astaxie/beego"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int `json:"expires_in"`
}

var ACCESS_TOKEN_KEY = "ACCESS_TOKEN_KEY"

func GetAccessToken() (token string) {

	var err error
	if v, _ := redis.String(cache.Get(ACCESS_TOKEN_KEY), err); v != "" {
		token = v
		return token
	} else {
		url := fmt.Sprintf(
			"%stoken?grant_type=client_credential&appid=%s&secret=%s",
			"https://api.weixin.qq.com/cgi-bin/",
			"wx478a0894ea001e37",
			"a44026286b9bc89b544ab8cbbab47ac2",
		)

		data, err := utils.Get(url)
		if err != nil {
			beego.Error(err)
			return ""
		}
		var ac AccessToken
		if err := json.Unmarshal(data, &ac); err != nil {
			return ""
		}
		//放入redis中
		timeoutDuration := time.Second * 7000//(ac.ExpiresIn - 200)
		err = cache.Put(ACCESS_TOKEN_KEY, ac.AccessToken, timeoutDuration)
		if nil != err {
			beego.Error(err)
		}
		return ac.AccessToken
	}
}
