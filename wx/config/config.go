package config

const (
	AppID = "wxb4f36cf1b0708cba"
	AppSecret = "7e4e29a4414758b3bb753c3a63f3bff8"
	TokenErrorNumber = 5
	Api = "https://api.weixin.qq.com/cgi-bin/"
	Token = "VUe8o5z3H40FgEZNEs58sAs9k0XYjsfA"
)

const (
	TokenUrl = Api + "token?"
	UserInfoUrl = Api + "user/info?"
	CreateMenu = Api + "menu/create"
)

const (
	// 消息
	MsgTypeText = "text"
	MsgTypeImage = "image"
	MsgTypeVoice = "voice"
	MsgTypeVideo = "video"
	MsgTypeMusic = "music"
	MsgTypeNews = "news"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation = "location"
	MsgTypeLink = "link"
	MsgTypeEvent = "event"

	// 事件
	EventSubscribe = "subscribe" // 订阅
	EventUnSubscribe = "unsubscribe" // 取消订阅
	EventScan = "SCAN"
	EventLocation = "LOCATION"
	EventClick = "CLICK"
	EventView = "VIEW"
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
)

