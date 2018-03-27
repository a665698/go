package config

const (
	AppID = "wxb4f36cf1b0708cba"
	AppSecret = "7e4e29a4414758b3bb753c3a63f3bff8"
	TokenErrorNumber = 5
	Domain = "https://api.weixin.qq.com/"
	Api = Domain + "cgi-bin/"
	Token = "VUe8o5z3H40FgEZNEs58sAs9k0XYjsfA"
)

const (
	TokenUrl = Api + "token?" // 获取accessToken
	UserInfoUrl = Api + "user/info?" // 获取用户信息
	MenuCreateUrl = Api + "menu/create?" // 新建自定义菜单
	MenuDeleteUrl = Api + "menu/delete?" // 删除自定义菜单
)

const (
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
)

const (
	// 消息
	MsgTypeText = "text"
	MsgTypeMusic = "music"
	MsgTypeNews = "news"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation = "location"
	MsgTypeLink = "link"
	MsgTypeEvent = "event"

	// 事件
	EventSubscribe = "subscribe" // 订阅
	EventUnSubscribe = "unsubscribe" // 取消订阅
	EventLocation = "LOCATION" // 发送地图
	EventClick = "CLICK" // 点击
	EventView = "VIEW" // 跳转
	EventScanCodePush = "scancode_push" // 扫码推
	EventScanCodeWait = "scancode_waitmsg"
	EventPicSysPhoto = "pic_sysphoto"
	EventPicPhotoOrAlbum = "pic_photo_or_album"
	EventPicAlbum = "pic_weixin"
	EventLocationSelect = "location_select"
)

// 菜单
const (
	MenuClick = "click"
	MenuView = "view"
	MenuScanCodePush = "scancode_push"
	MenuScanCodeWait = "scancode_waitmsg"
	MenuPicSysPhoto = "pic_sysphoto"
	MenuPicPhotoOrAlbum = "pic_photo_or_album"
	MenuPicAlbum = "pic_weixin"
	MenuLocation = "location_select"
	MenuMedia = "media_id"
	MenuMediaView = "view_limited"

	MenuClickButtonName = "点击按钮"
	MenuClickButtonKey = "clickButton"
	MenuViewButtonName = "跳转按钮"
	MenuViewButtonUrl = "http://www.zsuch.com"
	MenuLocationButtonName = "发送位置"
	MenuLocationButtonKey = "locationButton"
	MenuScanCodePushButtonName = "扫码推"
	MenuScanCodePushButtonKey = "scanCodePushButton"
	MenuScanCodeWaitButtonName = "扫码带提示"
	MenuScanCodeWaitButtonKey = "scanCodeWaitButton"
	MenuPicSysPhotoButtonName = "系统拍照发图"
	MenuPicSysPhotoButtonKey = "picSysPhotoButton"
	MenuPicPhotoOrAlbumButtonName = "拍照或者相册发图"
	MenuPicPhotoOrAlbumButtonKey = "PicPhotoOrAlbumButton"
	MenuPicAlbumButtonName = "相册发图"
	MenuPicAlbumButtonKey = "albumButton"
)

const (
	CustomApi = Domain + "customservice/kfaccount/"
	CustomAdd = CustomApi + "add?"
	CustomUpdate = CustomApi + "update?"
	CustomDel = CustomApi + "del?"
)

const (
	MediaTmpAdd = Api + "media/upload?access_token=%s&type=%s"
)

