package reptile

import "douban/proxy_pool"

func Main() {
	go proxy_pool.Main()
	tickTags()
}
