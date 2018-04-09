package myhttp

func (r *Router) WxRoutes() {
	r.Get("/", Index)
	r.Group("/wx/", func() {
		r.Post("/", WxHandle)
		r.Get("/")
	}, WxBaseFunc)

	go GetAccessToken()
}

