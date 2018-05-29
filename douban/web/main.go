package web

func Main() {
	engine := Routes()
	engine.Run(":3000")
}
