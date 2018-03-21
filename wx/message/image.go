package message

type ImageMessage struct {
	BaseMessage
	Image struct{
		MediaId string
	}
}

func (m *ImageMessage) Handle (w *Message) {
	m.Image.MediaId = w.MediaId
}
