package gee

import (
	"fmt"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) AddRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) Handle(ctx *Context) {
	pattern := ctx.Rep.Method + "-" + ctx.Rep.URL.Path
	if handler, ok := r.handlers[pattern]; ok {
		handler(ctx)
	} else {
		fmt.Fprintf(ctx.Writer, "404,NOT FOUND:%v", ctx.Rep.URL)
	}
}
