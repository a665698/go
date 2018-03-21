package message

type TextMessage struct {
	BaseMessage
	Content string
}

func (m *TextMessage) Handle (w *Message) {
	m.Content = "您输入的文字：" + w.Content
}

