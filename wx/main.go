package wx

import (
	"time"
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	AppID = "wxb4f36cf1b0708cba"
	AppSecret = "7e4e29a4414758b3bb753c3a63f3bff8"
	tokenErrorNumber = 5
	Api = "https://api.weixin.qq.com/cgi-bin/"
	Token = "VUe8o5z3H40FgEZNEs58sAs9k0XYjsfA"
)

type XmlHandle interface {
	Handle()
	response() []byte
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string
	Event string
	Body []byte `xml:"-"`
}

type BaseResponse struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string
	FromUserName string
	CreateTime time.Duration
	MsgType string
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

func (r *BaseResponse) Handle()  {

}

// 返回xml文本
func (r *BaseResponse) response() []byte {
	textMessage := TextMessage{}
	textMessage.MsgType = "text"
	textMessage.FromUserName, textMessage.ToUserName = r.ToUserName, r.FromUserName
	textMessage.CreateTime = time.Duration(time.Now().Unix())
	textMessage.Content = "处理错误"
	result, _ := xml.MarshalIndent(textMessage, "  ", "    ")
	return result
}

// 判断消息类型
func (wx *Response) TypeHandle() []byte {
	fmt.Println("MsgType:", wx.MsgType)
	var handle XmlHandle
	switch wx.MsgType {
	case "text":
		handle = &TextMessage{}
	case "image":
		handle = &ImageMessage{}
	case "voice":
		handle = &VoiceMessage{}
	case "video":
		handle = &VideoMessage{}
	case "shortvideo":
		handle = &VoiceMessage{}
	case "location":
		handle = &LocationMessage{}
	case "link":
		handle = &LinkMessage{}
	case "event":
		//wx.EventTypeHandle()
		handle = wx.EventTypeHandle()
	}
	xml.Unmarshal(wx.Body, handle)
	handle.Handle()
	return handle.response()
}

// 判断事件类型
func (wx *Response) EventTypeHandle() XmlHandle {
	fmt.Println("Event:", wx.Event)
	switch wx.Event {
	case "CLICK":
		return &MenuMessage{}
	}
	return nil
}