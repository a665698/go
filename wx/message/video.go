package message

import (
	"wx/config"
)

type VideoMessage struct {
	BaseMessage

	Video struct{
		MediaId string
		Title string
		Description string
	}
}

func (m *Message) NewVideoMessage () *VideoMessage {
	v := &VideoMessage{}
	v.Video.MediaId = m.MediaId
	v.Video.Title = "标题"
	v.Video.Description = "描述"
	v.MsgType = config.MediaTypeVideo
	return v
}
