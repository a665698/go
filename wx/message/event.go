package message

import (
	"wx/log"
	"fmt"
	"strconv"
	"wx/user"
	"wx/config"
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

// 按钮点击
func (m *Message) Click() interface{}  {
	var content string
	if m.EventKey == config.MenuClickButtonKey {
		content = "您点击了按钮：" + config.MenuClickButtonName
	} else {
		content = "您点击了按钮,但我不知道您点击了哪个 -_-"
	}
	t := m.NewTextMessage(content)
	return t
}

// 跳转URL
func (m *Message) View() interface{} {
	if m.EventKey == config.MenuViewButtonUrl {
		log.MessageLog("点击了按钮：" + config.MenuViewButtonName)
	}
	return nil
}

// 扫码推
func (m *Message) ScanCodePush() interface{} {
	if m.EventKey == config.MenuScanCodePushButtonKey {
		content := "点击了按钮：" + config.MenuScanCodePushButtonName
		log.MessageLog(content)
	}
	return nil
}

// 扫码带提示
func (m *Message) ScanCodeWait() interface{} {
	var content string
	if m.EventKey == config.MenuScanCodeWaitButtonKey {
		content = "点击了按钮：" + config.MenuScanCodeWaitButtonName
	} else {
		content = "您点击了按钮,但我不知道您点击了哪个 -_-"
	}
	t := m.NewTextMessage(content)
	fmt.Println(t)
	return t
}

// 系统拍照发图
func (m *Message) PicSysPhoto() interface{} {
	if m.EventKey == config.MenuPicSysPhotoButtonKey {
		content := "点击了按钮：" + config.MenuPicSysPhotoButtonName
		log.MessageLog(content)
	}
	return nil
}

// 系统拍照 or 相册发图
func (m *Message) PicPhotoOrAlbum() interface{} {
	if m.EventKey == config.MenuPicPhotoOrAlbumButtonKey {
		content := "点击了按钮：" + config.MenuPicPhotoOrAlbumButtonName
		log.MessageLog(content)
	}
	return nil
}

// 相册发图
func (m *Message) PicAlbum() interface{} {
	if m.EventKey == config.MenuPicAlbumButtonKey {
		content := "点击了按钮：" + config.MenuPicAlbumButtonName
		log.MessageLog(content)
	}
	return nil
}

// 位置选择器
func (m *Message) LocationSelect() interface{} {
	if m.EventKey == config.MenuLocationButtonKey {
		content := "点击了按钮：" + config.MenuLocationButtonName
		log.MessageLog(content)
	}
	return nil
}