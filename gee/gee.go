package gee

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *Router
}

func NewEngine() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (g *Engine) AddRouter(mathod string, path string, handler HandlerFunc) {
	// key := mathod + "-" + path
	g.router.AddRouter(mathod, path, handler)
}

func (g *Engine) GET(path string, handler HandlerFunc) {
	g.AddRouter("GET", path, handler)
}

func (g *Engine) POST(path string, handler HandlerFunc) {
	g.AddRouter("POST", path, handler)
}

func (g *Engine) DELETE(path string, handler HandlerFunc) {
	g.AddRouter("DELETE", path, handler)
}

func (g *Engine) PUT(path string, handler HandlerFunc) {
	g.AddRouter("PUT", path, handler)
}

func (g *Engine) Run(path string) error {
	return http.ListenAndServe(path, g)
}

func (g *Engine) ServeHTTP(w http.ResponseWriter, rep *http.Request) {
	c := NewContext(w, rep)
	g.router.Handle(c)
}
