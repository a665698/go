package user

import (
	"wx/common"
	"wx/config"
	"wx/token"
	"wx/log"
)

func GetUsetInfo(openId string) *common.GetResponse {
	s, err := common.Get(config.UserInfoUrl + "access_token=" + token.AccessToken.Token + "&openid=" + openId + "&lang=zh_CN")
	if err != nil {
		log.MessageLog(err.Error())
		return  nil
	}
	return s
}
