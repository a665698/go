package wx

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
	"wx/message"
	"encoding/xml"
)

const (
	AppID = "wxb4f36cf1b0708cba"
	AppSecret = "7e4e29a4414758b3bb753c3a63f3bff8"
	tokenErrorNumber = 5
	Api = "https://api.weixin.qq.com/cgi-bin/"
	Token = "VUe8o5z3H40FgEZNEs58sAs9k0XYjsfA"
)

func Handle(body []byte) []byte {
	w := &message.Message{}
	xml.Unmarshal(body, w)
	return w.TypeHandle()
}

// 发送GET请求获取数据
func get(url string, v interface{}) bool {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("url get error", url , err)
		return false
	}
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("url body analysis error", url, err)
		return false
	}
	if err = json.Unmarshal(r, v); err != nil {
		fmt.Println("data to json error", url, err)
		return false
	}
	switch v.(type) {
	case *AccessToken:
		return v.(*AccessToken).Validate()
	}
	return true
}

func post(url, postInfo string) bool {
	result, err := http.Post(url, "application/json; encoding=utf-8", strings.NewReader(postInfo))
	if err != nil {
		fmt.Println(err)
		return false
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	var r map[string]interface{}
	if err = json.Unmarshal(body, &r); err != nil {
		fmt.Println(err)
		return false
	}
	if errCode, ok := r["errcode"]; !ok {
		fmt.Println("errcode not fount")
		return false
	} else {
		code := int(errCode.(float64))
		if code == 0 {
			return true
		} else {
			fmt.Println(GetError(code))
			return false
		}
	}
}