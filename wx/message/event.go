package message

import (
	"wx/log"
	"fmt"
	"strconv"
	"wx/user"
)

// 关注事件
func (m *Message) Subscribe() interface{} {
	fmt.Println(m.EventKey)
	var t *TextMessage
	if u := user.GetUsetInfo(m.FromUserName);u == nil{
		t = m.NewTextMessage("欢迎关注公众号")
	} else {
		t = m.NewTextMessage("欢迎关注公众号：" + u.Nickname)
	}
	return t
}

// 取消关注
func (m *Message) UnSubscribe() interface{} {
	log.MessageLog("用户取消订阅,openId：" + m.FromUserName)
	return nil
}

// 上报地理位置
func (m *Message) Location() interface{} {
	log.MessageLog("纬度：" + strconv.FormatFloat(m.Latitude, 'g', -1, 64))
	log.MessageLog("经度：" + strconv.FormatFloat(m.Longitude, 'g', -1, 64))
	log.MessageLog("精度：" + strconv.FormatFloat(m.Precision, 'g', -1, 64))
	return nil
}
