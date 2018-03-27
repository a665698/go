package message

import (
	"encoding/xml"
	"time"
	"fmt"
	"wx/config"
)

type HandleFun interface {
	Handle(*Message)
	BaseHandle(*Message)
}

type BaseMessage struct {
	XMLName xml.Name `xml:"xml"`
	ToUserName string // 开发者微信号
	FromUserName string // 发送方帐号（一个OpenID）
	CreateTime time.Duration // 消息创建时间 （整型）
	MsgType string // 消息类型
}

type Message struct {
	BaseMessage

	// 基本消息
	Content string // 文本消息内容
	MsgId uint64 // 消息id，64位整型
	MediaId string // 消息媒体id，可以调用多媒体文件下载接口拉取数据。
	PicUrl string // 图片链接（由系统生成）
	Format string // 语音格式，如amr，speex等
	Recognition string // 语音识别结果，UTF8编码
	ThumbMediaId string // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	LocationX string `xml:"Location_X"` // 地理位置维度
	LocationY string `xml:"Location_Y"` // 地理位置经度
	Scale string // 地图缩放大小
	Label string // 地理位置信息
	Title string // 消息标题
	Description string // 消息描述
	Url string // 消息链接

	// 事件
	Event string // 事件类型
	EventKey string // 事件KEY值
	Ticket string // 二维码的ticket，可用来换取二维码图片
	Latitude float64 // 地理位置纬度
	Longitude float64 // 地理位置经度
	Precision float64 // 地理位置精度

	MenuID int // 指菜单ID，如果是个性化菜单，则可以通过这个字段，知道是哪个规则的菜单被点击了。

	ScanCodeInfo struct{
		ScanType string // 扫描类型，一般是qrcode
		ScanResult string // 扫描结果，即二维码对应的字符串信息
	}	// 扫描信息

	SendPicsInfo struct{
		Count int // 发送的图片数量
		PicList []struct{
			Item struct{
				PicMd5Sum string // 图片的MD5值，开发者若需要，可用于验证接收到图片
			} `xml:"item"`
		} // 图片列表
	}	// 	发送的图片信息

	SendLocationInfo struct{
		LocationX float64 `xml:"Location_X"` // X坐标信息
		LocationY float64 `xml:"Location_Y"` // Y坐标信息
		Scale int // 精度，可理解为精度或者比例尺、越精细的话 scale越高
		Label string // 地理位置的字符串信息
		PoiName string `xml:"Poiname"` // 朋友圈POI的名字，可能为空
	} // 发送的位置信息
}

func (m *Message) BaseHandle(b *BaseMessage) {
	b.FromUserName ,b.ToUserName = m.ToUserName, m.FromUserName
	b.CreateTime = time.Duration(time.Now().Unix())
}

// 判断消息类型
func (m *Message) TypeHandle() []byte {
	fmt.Println("MsgType:", m.MsgType)
	var i interface{}
	switch m.MsgType {
	case config.MsgTypeText:
		i = m.TextMessageHandle()
	case config.MediaTypeImage:
		i = m.NewImageMessage(m.MediaId)
	case config.MediaTypeVoice:
		i = m.NewTextMessage("收到语音消息")
	case config.MediaTypeVideo, config.MsgTypeShortVideo:
		i = m.NewTextMessage("收到视频消息")
	case config.MsgTypeLocation:
		i = m.NewTextMessage("收到地图消息")
	case config.MsgTypeLink:
		i = m.NewTextMessage("收到链接消息")
	case config.MsgTypeEvent:
		i = m.EventTypeHandle()
	default:
		i = nil
	}
	if i == nil {
		return []byte("success")
	}
	body, _ := xml.Marshal(i)
	return body
}

// 判断事件类型
func (m *Message) EventTypeHandle() interface{} {
	fmt.Println("Event:", m.Event)
	fmt.Println("EventKey", m.EventKey)
	switch m.Event {
	case config.EventSubscribe:
		return m.Subscribe()
	case config.EventUnSubscribe:
		return m.UnSubscribe()
	case config.EventLocation:
		return m.Location()
	case config.EventClick:
		return m.Click()
	case config.EventView:
		return m.View()
	case config.EventScanCodePush:
		return m.ScanCodePush()
	case config.EventScanCodeWait:
		return m.ScanCodeWait()
	case config.EventPicSysPhoto:
		return m.PicSysPhoto()
	case config.EventPicPhotoOrAlbum:
		return m.PicPhotoOrAlbum()
	case config.EventPicAlbum:
		return m.PicAlbum()
	case config.EventLocationSelect:
		return m.LocationSelect()
	}
	return nil
}

