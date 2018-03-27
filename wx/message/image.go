package message

import (
	"wx/config"
)

type ImageMessage struct {
	BaseMessage
	Image struct{
		MediaId string
	}
}

func (m *Message) NewImageMessage (mediaId string) *ImageMessage {
	i := &ImageMessage{}
	i.Image.MediaId = mediaId
	i.MsgType = config.MediaTypeImage
	m.BaseHandle(&i.BaseMessage)
	return i
}
