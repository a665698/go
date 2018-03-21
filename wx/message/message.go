package message

import (
	"encoding/xml"
	"time"
	"fmt"
)

type HandleFun interface {
	Handle(*Message)
	BaseHandle(*Message)
}

type BaseMessage struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string
	FromUserName string
	CreateTime time.Duration
	MsgType string
}

type Message struct {
	BaseMessage

	// 基本消息
	Content string
	MsgId uint64
	MediaId string
	PicUrl string
	Format string
	ThumbMediaId string
	LocationX string `xml:"location_x"`
	LocationY string `xml:"location_y"`
	Scale string
	Label string
	Title string
	Description string
	Url string
}

func (m *BaseMessage) BaseHandle(w *Message) {
	m.FromUserName ,m.ToUserName = w.ToUserName, w.FromUserName
	m.MsgType = w.MsgType
	m.CreateTime = time.Duration(time.Now().Unix())
}

func (m *BaseMessage) Handle(*Message) {
	m.MsgType = "text"
}

// 判断消息类型
func (m *Message) TypeHandle() []byte {
	fmt.Println("MsgType:", m.MsgType)
	var handle HandleFun
	switch m.MsgType {
	case "text":
		handle = &TextMessage{}
	case "image":
		handle = &ImageMessage{}
	//case "voice":
	//	handle = &VoiceMessage{}
	//case "video":
	//	handle = &VideoMessage{}
	//case "shortvideo":
	//	handle = &VoiceMessage{}
	//case "location":
	//	handle = &LocationMessage{}
	//case "link":
	//	handle = &LinkMessage{}
	//case "event":
	//	//wx.EventTypeHandle()
	//	handle = wx.EventTypeHandle()
	default:
		return []byte("")
	}
	handle.BaseHandle(m)
	handle.Handle(m)
	body, _ := xml.Marshal(handle)
	return body
	//return []byte("fdsa")
	//xml.Unmarshal(wx.Body, handle)
}

//// 判断事件类型
//func (m *Message) EventTypeHandle() XmlHandle {
//	fmt.Println("Event:", wx.Event)
//	switch wx.Event {
//	case "CLICK":
//		return &MenuMessage{}
//	}
//	return &BaseResponse{}
//}

