package proxy_pool

import "time"

func Main() {
	go CheckIp()
	for {
		Get66Ip()
		time.Sleep(time.Second * 10)
	}
}
