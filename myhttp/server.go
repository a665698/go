package myhttp

import (
	"net/http"
	"fmt"
	"path"
)

const (
	ResourceFile = "resource"
)

type (
	Context struct {
		handlerFunc []HandleFunc
		request *http.Request
		response http.ResponseWriter
		index int8
	}
	HandleFunc func(content *Context)
	Router struct {
		path string
		HandleFunc []HandleFunc
		routers map[string]router
	}
	router struct {
		handlerFunc []HandleFunc
	}
)

func (r *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	basePath := fmt.Sprintf(path.Join(request.URL.Path, request.Method))
	r.path = request.URL.Path
	route, ok := r.routers[basePath]
	if !ok {
		Open(response, request)
	}
	c := Context{
		handlerFunc: route.handlerFunc,
		request: request,
		response: response,
		index: -1,
	}
	c.Next()
}

func Open(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response,request,ResourceFile + request.URL.Path)
}

func New() *Router {
	return &Router{
		path: "/",
		routers: make(map[string]router),
	}
}

func (r *Router) Use(handle ...HandleFunc) {
	r.HandleFunc = append(r.HandleFunc, handle...)
}

func (r *Router) Group(path string, f func(), handle ...HandleFunc) {
	tmpPath := r.path
	tmpHandel := r.HandleFunc
	r.path = path
	r.HandleFunc = r.combineFunc(handle)
	f()
	r.path = tmpPath
	r.HandleFunc = tmpHandel
}

func (r *Router) Get(url string, handle ...HandleFunc) {
	baseUrl := path.Join(r.path, url)
	handles := r.combineFunc(handle)
	r.Add(baseUrl, http.MethodGet, handles)
}

func (r *Router) Post(url string, handle ...HandleFunc) {
	baseUrl := path.Join(r.path, url)
	handles := r.combineFunc(handle)
	r.Add(baseUrl, http.MethodPost, handles)
}

func (r *Router) combineFunc(handle []HandleFunc) []HandleFunc {
	finallyLen := len(r.HandleFunc) + len(handle)
	finallyFunc := make([]HandleFunc, finallyLen)
	copy(finallyFunc, r.HandleFunc)
	copy(finallyFunc[len(r.HandleFunc):], handle)
	return finallyFunc
}

func (r *Router) Add(baseUrl , method string, handle []HandleFunc) {
	r.routers[path.Join(baseUrl, method)] = router{
		handlerFunc: handle,
	}
}

func (r * Router) Run()  {
	http.ListenAndServe(":80", r)
}

func (c *Context) Next()  {
	c.index ++
	if int8(len(c.handlerFunc)) > c.index  {
		c.handlerFunc[c.index](c)
	}
}







