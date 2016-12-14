package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
)

const (
	geocoder_api_url = "http://api.map.baidu.com/geocoder/v2/?location=%s,%s&output=json&pois=1&ak=D87fed887b258ae6cd1eba15a5a2e1ed"
)

type geocoderResp struct {
	GeocoderResult geocoderResult `json:"result"`
}

type geocoderResult struct {
	AddressComponent addressComponent `json:"addressComponent"`
}

type addressComponent struct {
	District string `json:"district"`
}

func GetLocationDistrict(lat, lng string) (district string) {
	url := fmt.Sprintf(geocoder_api_url, lat, lng)
	data, err := Get(url)
	if err != nil {
		beego.Error(err)
		return
	}

	resp := &geocoderResp{}

	err = json.Unmarshal(data, resp)
	if err != nil {
		beego.Error(err)
	}
	district = resp.GeocoderResult.AddressComponent.District
	return
}
