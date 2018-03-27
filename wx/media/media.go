package media

import (
	"wx/common"
	"wx/config"
	"wx/token"
	"fmt"
)

func AddTmpMedia() string {
	if err := common.UploadFile(
		fmt.Sprintf(config.MediaTmpAdd, token.AccessToken.Token, config.MediaTypeImage),
		"1.png");
		err != nil {
			return err.Error()
	}
	return "临时素材上传成功"
}

