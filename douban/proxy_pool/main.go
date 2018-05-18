package proxy_pool

import (
	"net/http"
	"net/url"
	"time"
)

func Main() {
	go CheckIp()
	Get66Ip()
}

func GetHttp(tUrl, proxyIp string) (*http.Response, error) {
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
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Mobile Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
