package common

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"
)

// 发送GET返回body
func GetHttp(tUrl, proxyIp string) ([]byte, error) {
	res, err := GetHttpRes(tUrl, proxyIp)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 获取http返回包
func GetHttpRes(tUrl, proxyIp string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}
	req, err := http.NewRequest("GET", tUrl, nil)
	if proxyIp != "" {
		parseUrl, _ := url.Parse("http://" + proxyIp)
		transport := http.Transport{
			Proxy: http.ProxyURL(parseUrl),
		}
		client.Transport = &transport
	}

	if err != nil {
		return nil, err
	}
	u, _ := url.Parse(tUrl)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Host", u.Host)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 删除字符串中的空白字符
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
