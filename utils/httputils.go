package utils

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"strings"
	"github.com/astaxie/beego"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetWithHeader(url string, header map[string]string) (data []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)

	for k, v := range header {
		req.Header.Add(k, v)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func PostWithProxy(uri string, param map[string]string) (data []byte, err error) {

	values := url.Values{}
	for k, v := range param {
		values.Add(k, v)
	}
	beego.Error(values)
	body := ioutil.NopCloser(strings.NewReader(values.Encode()))

	req, err := http.NewRequest("POST", uri, body)

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://123.206.18.16:8888")
	}
	transport := &http.Transport{Proxy: proxy}
	client := http.Client{Transport:transport}

	beego.Info(req)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
