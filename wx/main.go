package wx

import (
	"wx/message"
	"encoding/xml"
	"wx/config"
	"crypto/sha1"
	"strings"
	"fmt"
	"sort"
	"wx/token"
)

func Handle(body []byte) []byte {
	w := &message.Message{}
	xml.Unmarshal(body, w)
	return w.TypeHandle()
}

// 获取access_token
func GetAccessToken() {
	token.GetAccessToken()
}

// 生成签名用来判断签名是否正确
func MakeSignature(timestamp, nonce string) string {
	str := []string{timestamp, nonce, config.Token}
	sort.Strings(str)
	s := sha1.New()
	s.Write([]byte(strings.Join(str, "")))
	return fmt.Sprintf("%x", s.Sum(nil))
}