package menu

import (
	"wx/config"
	"encoding/json"
	"wx/common"
	"wx/token"
	"wx/log"
)

type Button struct {
	Button []*Button `json:"button,omitempty"`
	SubButton []*Button `json:"sub_button,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Key string `json:"key,omitempty"`
	Url string `json:"url,omitempty"`
	MediaId string `json:"media_id,omitempty"`
}

// 设置点击按钮
func (m *Button) SetClick(name, key string) {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuClick
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置子菜单按钮
func (m *Button) SetSub(name string) {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = ""
	m.Key = ""
	m.Url = ""
	m.MediaId = ""
}

// 设置跳转按钮
func (m *Button) SetView(name, url string) {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuView
	m.Key = ""
	m.Url = url
	m.MediaId = ""
}

// 设置扫码推按钮
func (m *Button) SetScanCodePush(name ,key string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuScanCodePush
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置扫码带提示按钮
func (m *Button) SetScanCodeWait(name ,key string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuScanCodeWait
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置系统拍照发图按钮
func (m *Button) SetPicSysPhoto(name ,key string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuPicSysPhoto
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置拍照或相册发图按钮
func (m *Button) SetPicPhotoOrAlbum(name ,key string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuPicPhotoOrAlbum
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置相册发图按钮
func (m *Button) SetPicAlbum(name ,key string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuPicAlbum
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置发送位置按钮
func (m *Button) SetLocationSelect(name ,key string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuLocation
	m.Key = key
	m.Url = ""
	m.MediaId = ""
}

// 设置下发消息按钮
func (m *Button) SetMedia(name ,mediaId string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuMedia
	m.Key = ""
	m.Url = ""
	m.MediaId = mediaId
}

// 设置图文消息按钮
func (m *Button) SetMediaView(name ,mediaId string)  {
	m.Button = nil
	m.SubButton = nil
	m.Name = name
	m.Type = config.MenuMediaView
	m.Key = ""
	m.Url = ""
	m.MediaId = mediaId
}

// 合并子菜单按钮
func (b *Button) combineSubButton (cb *Button)  {
	l := len(b.SubButton)
	s := make([]*Button, l +1)
	for i := 0; i< l ; i++ {
		s[i] = b.SubButton[i]
	}
	s[l] = cb
	b.SubButton = s
}

// 合并主菜单按钮
func (b *Button) combineButton (cb *Button)  {
	l := len(b.Button)
	s := make([]*Button, l +1)
	for i := 0; i< l ; i++ {
		s[i] = b.Button[i]
	}
	s[l] = cb
	b.Button = s
}

// 合并菜单按钮
func (b *Button) Combine (cb *Button, master bool) {
	if master {
		b.combineButton(cb)
	} else {
		b.combineSubButton(cb)
	}
}

// 新建点击按钮
func CreateClick(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetClick(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建主菜单
func CreateMaster() *Button {
	b := &Button{}
	return b
}

// 新建跳转按钮
func CreateView(name, url string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetView(name, url)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建扫码推按钮
func CreateScanCodePush(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetScanCodePush(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建扫码带提示按钮
func CreateScanCodeWait(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetScanCodeWait(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建系统拍照发图按钮
func CreatePicSysPhoto(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetPicSysPhoto(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建相册发图按钮
func CreatePicAlbum(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetPicAlbum(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建发送位置按钮
func CreateLocation(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetLocationSelect(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建下发消息按钮
func CreateMedia(name, mediaId string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetMedia(name, mediaId)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建图文消息按钮
func CreateMediaView(name, mediaId string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetMediaView(name, mediaId)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建拍照或相册发图按钮
func CreatePicPhotoOrAlbum(name, key string, b *Button, master bool) *Button {
	c := &Button{}
	c.SetPicPhotoOrAlbum(name, key)
	if b != nil {
		b.Combine(c, master)
		return nil
	}
	return c
}

// 新建子菜单按钮
func CreateSub(name string) *Button {
	c := &Button{}
	c.SetSub(name)
	return c
}

// 创建自定义菜单
func Create() string {
	b := CreateMaster()

	s1 := CreateSub("菜单1")
	CreateClick(config.MenuClickButtonName, config.MenuClickButtonKey, s1, false)
	CreateView(config.MenuViewButtonName, config.MenuViewButtonUrl, s1, false)
	CreateLocation(config.MenuLocationButtonName, config.MenuLocationButtonKey, s1, false)
	b.Combine(s1, true)

	s2 := CreateSub("菜单2")
	CreateScanCodePush(config.MenuScanCodePushButtonName, config.MenuScanCodePushButtonKey, s2, false)
	CreateScanCodeWait(config.MenuScanCodeWaitButtonName, config.MenuScanCodeWaitButtonKey, s2, false)
	b.Combine(s2, true)

	s3 := CreateSub("菜单3")
	CreatePicSysPhoto(config.MenuPicSysPhotoButtonName, config.MenuPicSysPhotoButtonKey, s3, false)
	CreatePicPhotoOrAlbum(config.MenuPicPhotoOrAlbumButtonName, config.MenuPicPhotoOrAlbumButtonKey, s3, false)
	CreatePicAlbum(config.MenuPicAlbumButtonName, config.MenuPicAlbumButtonKey, s3, false)
	b.Combine(s3, true)

	j, _ := json.Marshal(b)
	err := common.Post(config.MenuCreateUrl + "access_token=" + token.AccessToken.Token, j)
	content := ""
	if err != nil {
		content = err.Error()
	} else {
		content = "菜单设置成功"
	}
	log.MessageLog(content)
	return content
}

// 删除自定义菜单
func Delete() string {
	content := ""
	if _, err := common.Get(config.MenuDeleteUrl + "access_token=" + token.AccessToken.Token); err != nil {
		content = err.Error()
	} else {
		content = "菜单删除成功"
	}
	log.MessageLog(content)
	return content
}
