package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
)

var header map[string]string

const (
	api_url = "http://apis.baidu.com/heweather/pro/weather?city=%s"
)

type rootObject struct {
	Result []weather `json:"HeWeather data service 3.0"`
}

type weather struct {
	Aqi           aqi `json:"aqi"`
	Now           now `json:"now"`
	Suggestion    sug `json:"suggestion"`
	DailyForecast []dailyForecast `json:"daily_forecast"`
	Status        string `json:"status"`
}

type aqi struct {
	City city `json:"city"`
}

type suggestion struct {
	Brf string `json:"brf"`
	Txt string `json:"txt"`
}

type now struct {
	Cond cond `json:"cond"`
	Fl   string `json:"fl"`
	Tmp  string `json:"tmp"`
}

type sug struct {
	Drsg  suggestion `json:"drsg"`
	Sport suggestion  `json:"sport"`
}

type cond struct {
	Code string `json:"code"`
	Txt  string `json:"txt"`
}

type city struct {
	Qlty string `json:"qlty"`
	Pm25 string `json:"pm25"`
}

type dailyForecast struct {
	Tmp tmp `json:"tmp"`
}

type tmp struct {
	Max string `json:"max"`
	Min string `json:"min"`
}

func GetCityWeather(city string) (resp string) {
	url := fmt.Sprintf(api_url, city)
	data, err := GetWithHeader(url, header)
	if err != nil {
		beego.Error(err)
		return
	}

	result := &rootObject{}

	err = json.Unmarshal(data, result)
	if err != nil {
		beego.Error(err)
	}

	beego.Error(result.Result[0])
	if "unknown city" == result.Result[0].Status {
		resp = ""
		return
	}

	now := result.Result[0].Now
	today := result.Result[0].DailyForecast[0]
	drsg := result.Result[0].Suggestion.Drsg
	sport := result.Result[0].Suggestion.Sport
	aqi := result.Result[0].Aqi
	resp = fmt.Sprintf("%s今天%s,%s℃~%s℃,pm2.5 %s %s,当前温度%s℃,体感温度%s°,穿衣指数:%s,%s运动指数:%s,%s", city, now.Cond.Txt, today.Tmp.Min, today.Tmp.Max, aqi.City.Pm25, aqi.City.Qlty, now.Tmp, now.Fl, drsg.Brf, drsg.Txt, sport.Brf, sport.Txt)
	return
}

func init() {
	header = map[string]string{"apikey":"9272fa42651c51e1a10c90a61b6807aa"}
}

