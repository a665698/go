package message

import (
	"wx/config"
	"wx/menu"
)

type TextMessage struct {
	BaseMessage
	Content string
}

func (m *Message) TextMessageHandle() *TextMessage {
	content := ""
	switch m.Content {
	case "设置菜单":
		content = menu.Create()
	case "删除菜单":
		content = menu.Delete()
	default:
		content = "你输入的是：" + m.Content
	}
	i := m.NewTextMessage(content)
	return i
}

func (m *Message) NewTextMessage (content string) *TextMessage {
	t := &TextMessage{}
	t.Content = content
	t.MsgType = config.MsgTypeText
	m.BaseHandle(&t.BaseMessage)
	return t
}

