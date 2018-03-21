package myhttp

import (
	"strings"
	"sort"
	"crypto/sha1"
	"fmt"
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
	if signature == makeSignature(timestamp, nonce) {
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

// 生成签名用来判断签名是否正确
func makeSignature(timestamp, nonce string) string {
	str := []string{timestamp, nonce, wx.Token}
	sort.Strings(str)
	s := sha1.New()
	s.Write([]byte(strings.Join(str, "")))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func WxHandle(c *Context) {
	body, err := ioutil.ReadAll(c.request.Body)
	if err != nil {
		http.Error(c.response, "error", http.StatusBadRequest)
		return
	}
	c.response.Write(wx.Handle(body))

	//wxResult := &wx.Response{}
	//wxResult.Body = body
	//xml.Unmarshal(body, wxResult)
	//result := wxResult.TypeHandle()
	//c.response.Write(result)

	//response,err := xml.MarshalIndent(wxResult, "", "")
	//if err != nil {
	//	http.Error(c.response, "error", http.StatusBadRequest)
	//	return
	//}
	// 获取用户信息
	//wxResult.GetUserInfo()
	//wx.MenuCreate()
	//c.response.Write(wxResult.BaseResponse())
}

func Index(c *Context) {
	c.response.Write([]byte("index"))
}



