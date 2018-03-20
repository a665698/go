package main

import (
	"myhttp"
)

func main() {
	app := myhttp.New()
	app.WxRoutes()
	app.Run()
}
