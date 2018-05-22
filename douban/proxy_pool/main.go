package proxy_pool

func Main() {
	go CheckIp()
	for {
		Get66Ip()
	}
}
