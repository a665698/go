package message

import (
	"wx/config"
)

type NewsMessage struct {
	BaseMessage

	ArticleCount uint8
	Articles struct{
		Item []NewsMessageItem `xml:"item"`
	}
}

type NewsMessageItem struct {
	Title string
	Description string
	PicUrl string
	Url string
}

func (m *Message) NewNewsMessage () *NewsMessage {
	n := &NewsMessage{}
	n.MsgType = config.MsgTypeNews
	return n
}