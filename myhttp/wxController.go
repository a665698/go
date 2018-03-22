package myhttp

import (
	"strings"
	"net/http"
	"io/ioutil"
	"wx"
)

func GetAccessToken() {
	wx.GetAccessToken()
}

// 验证是否是微信端发送的请求
func WxBaseFunc(c *Context) {
	c.request.ParseForm()
	timestamp := strings.Join(c.request.Form["timestamp"], "")
	nonce := strings.Join(c.request.Form["nonce"], "")
	signature := strings.Join(c.request.Form["signature"], "")
	if signature == wx.MakeSignature(timestamp, nonce) {
		if c.request.Method == http.MethodGet {
			schostr := strings.Join(c.request.Form["echostr"], "")
			c.response.Write([]byte(schostr))
		} else {
			c.response.Header().Set("Content-Type", "text/xml")
			c.Next()
		}
	} else {
		http.Error(c.response, "invalid URL path", http.StatusBadRequest)
	}
}

func WxHandle(c *Context) {
	body, err := ioutil.ReadAll(c.request.Body)
	if err != nil {
		http.Error(c.response, "error", http.StatusBadRequest)
		return
	}
	c.response.Write(wx.Handle(body))
}

func Index(c *Context) {
	c.response.Write([]byte("index"))
}



