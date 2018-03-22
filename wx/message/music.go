package message

import (
	"wx/config"
)

type MusicMessage struct {
	BaseMessage
	Music struct{
		Title string
		Description string
		MusicUrl string
		HQMusicUrl string
		ThumbMediaId string
	}
}

func (m *Message) NewMusicMessage () *MusicMessage {
	mu := &MusicMessage{}
	mu.MsgType = config.MsgTypeMusic
	return mu
}
