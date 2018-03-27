package message

import (
	"wx/config"
)

type VoiceMessage struct {
	BaseMessage

	Voice struct{
		MediaId string
	}
}

func (m *Message) NewVoiceMessage () *VoiceMessage {
	v := &VoiceMessage{}
	v.Voice.MediaId = m.MediaId
	v.MsgType = config.MediaTypeVoice
	return v
}