package token

import (
	"time"
	"wx/config"
	"wx/common"
	"wx/log"
)

var AccessToken *Token

type Token struct {
	Token string
	Expires int64
}

// 从微信获取accessToken
func GetAccessToken() {
	AccessToken = &Token{}
	errorNumber := 0
	for errorNumber < config.TokenErrorNumber {
		nowTime := time.Now().Unix()
		if AccessToken.Token == "" || nowTime >= AccessToken.Expires {
			if s, err := common.Get(config.TokenUrl +"grant_type=client_credential&appid=" + config.AppID + "&secret=" + config.AppSecret); err == nil {
				errorNumber = 0
				AccessToken.Token = s.AccessToken
				AccessToken.Expires = s.ExpiresIn + nowTime
				time.Sleep(time.Second)
			} else {
				log.MessageLog(err.Error())
				errorNumber ++
			}
		}
	}
	log.MessageLog("token 错误次数超过" + string(config.TokenErrorNumber))
}

