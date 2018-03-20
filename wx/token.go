package wx

import (
	"time"
	"fmt"
)

var accessToken *AccessToken

type AccessToken struct {
	Token string `json:"access_token"`
	Expires int64 `json:"expires_in"`
	ErrCode int `json:"errcode"`
}

// 验证微信端返回的Token信息是否正确
func (token *AccessToken) Validate() bool {
	if token.ErrCode != 0 {
		fmt.Println("token get error", GetError(token.ErrCode))
		return false
	} else if token.Token == "" || token.Expires == 0 {
		fmt.Println("token value error", token)
		return false
	}
	token.Expires = time.Now().Unix() + token.Expires
	return true
}

// 从微信获取accessToken
func GetAccessToken() {
	accessToken = &AccessToken{}
	errorNumber := 0
	for errorNumber < tokenErrorNumber {
		nowTime := time.Now().Unix()
		if accessToken.Token == "" || nowTime >= accessToken.Expires {
			if false == get(TokenUrl +"?grant_type=client_credential&appid=" + AppID + "&secret=" + AppSecret, accessToken) {
				errorNumber ++
				time.Sleep(time.Second)
			}
		}
	}
}
