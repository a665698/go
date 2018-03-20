package wx

import (
	"time"
	"encoding/xml"
)

type (
	// 普通消息
	Message struct {
		BaseResponse
		MsgId uint64
	}
	// 消息媒介ID，可以调用多媒体文件下载接口拉取数据。
	Media struct {
		MediaId string
	}
	// 文本消息
	TextMessage struct {
		Message
		BaseResponse
		Content string
	}
	// 图片消息
	ImageMessage struct {
		Message
		Media
		PicUrl string
	}
	// 语音消息
	VoiceMessage struct {
		Message
		Media
		Format string
	}
	// 视频消息、小视频消息
	VideoMessage struct {
		Message
		Media
		ThumbMediaId string
	}
	// 地理位置消息
	LocationMessage struct {
		Message
		LocationX string `xml:"location_x"`
		LocationY string `xml:"location_y"`
		Scale string
		Label string
	}
	// 链接消息
	LinkMessage struct {
		Message
		Title string
		Description string
		Url string
	}
)

// 文本消息处理
func (m *TextMessage) Handle()  {
	m.Content = "你输入的消息为:" + m.Content
}
// 文本消息返回
func (m *TextMessage) response() []byte {
	m.FromUserName, m.ToUserName = m.ToUserName, m.FromUserName
	m.CreateTime = time.Duration(time.Now().Unix())
	result, _ := xml.MarshalIndent(m, "  ", "    ")
	return result
}
