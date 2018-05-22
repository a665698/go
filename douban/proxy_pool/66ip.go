package proxy_pool

import (
	"douban/common"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func Get66Ip() {
	b, err := common.GetHttpRes("http://www.66ip.cn/nmtq.php?getnum=1000", "")
	if err != nil {
		common.NoticeLog(err)
		return
	}
	document, err := goquery.NewDocumentFromReader(b.Body)
	if err != nil {
		common.NoticeLog(err)
		return
	}
	html := document.Text()
	rows := strings.Split(html, "\n")
	reg, err := regexp.Compile("\\d+.\\d+.\\d+.\\d:\\d+")
	if err != nil {
		common.NoticeLog(err)
		return
	}
	for _, ip := range rows {
		if reg.MatchString(ip) {
			WritePendingIP(strings.TrimSpace(ip))
		}
	}
}
