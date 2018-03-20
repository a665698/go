package wx

import (
	"time"
)

type User struct {
	Subscribe int8 `json:"subscribe"`
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int8 `json:"sex"`
	Language string `json:"language"`
	City string `json:"city"`
	Province string `json:"province"`
	Country string `json:"country"`
	HeadImgUrl string `json:"headimgurl"`
	SubscribeTime time.Duration `json:"subscribe_time"`
	UnionId string `json:"unionid"`
	Remark string `json:"remark"`
	GroupId int8 `json:"groupid"`
	TagIdList []int `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene int `json:"qr_scene"`
	QrSceneStr string `json:"qr_scene_str"`
}

// 获取用户信息
func (wx *Response) GetUserInfo() {
	//user := &User{}
	//if false == get(GetUserInfo + "?access_token=" + accessToken.Token + "&openid=" + wx.FromUserName + "&lang=zh_CN", user) {
	//	wx.Content = "获取用户信息失败"
	//} else {
	//	wx.Content = "hello " + user.Nickname
	//}
	//wx.BaseResponse()
}

