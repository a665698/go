package message

import (
	"encoding/xml"
	"time"
	"fmt"
	"wx/config"
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

	// 事件
	Event string
	EventKey string // 事件KEY值
	Ticket string // 二维码的ticket，可用来换取二维码图片
	Latitude float64 // 地理位置纬度
	Longitude float64 // 地理位置经度
	Precision float64 // 地理位置精度
}

func (m *Message) BaseHandle(b *BaseMessage) {
	b.FromUserName ,b.ToUserName = m.ToUserName, m.FromUserName
	b.CreateTime = time.Duration(time.Now().Unix())
}

// 判断消息类型
func (m *Message) TypeHandle() []byte {
	fmt.Println("MsgType:", m.MsgType)
	var i interface{}
	switch m.MsgType {
	case config.MsgTypeText:
		i = m.NewTextMessage("你输入的是：" + m.Content)
	case config.MsgTypeImage:
		i = m.NewImageMessage(m.MediaId)
	case config.MsgTypeVoice:
		i = m.NewTextMessage("收到语音消息")
	case config.MsgTypeVideo, config.MsgTypeShortVideo:
		i = m.NewTextMessage("收到视频消息")
	case config.MsgTypeLocation:
		i = m.NewTextMessage("收到地图消息")
	case config.MsgTypeLink:
		i = m.NewTextMessage("收到链接消息")
	case config.MsgTypeEvent:
		i = m.EventTypeHandle()
	default:
		i = nil
	}
	if i == nil {
		return []byte("")
	}
	body, _ := xml.Marshal(i)
	return body
}

// 判断事件类型
func (m *Message) EventTypeHandle() interface{} {
	fmt.Println("Event:", m.Event)
	switch m.Event {
	case config.EventSubscribe:
		return m.Subscribe()
	case config.EventUnSubscribe:
		return m.UnSubscribe()
	case config.EventLocation:
		return m.Location()
	}
	return nil
}

