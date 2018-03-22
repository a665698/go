package message

import (
	"wx/config"
)

type TextMessage struct {
	BaseMessage
	Content string
}

func (m *Message) NewTextMessage (content string) *TextMessage {
	t := &TextMessage{}
	t.Content = content
	t.MsgType = config.MsgTypeText
	m.BaseHandle(&t.BaseMessage)
	return t
}

