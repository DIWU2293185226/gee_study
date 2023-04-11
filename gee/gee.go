package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(ctx *Context)

type RouterGroup struct {
	//分组前缀
	prefix string
	//注册进来的中间件
	middlewares []HandlerFunc
	//指向引擎的指针，部分业务会需要用到引擎的功能
	engine *Engine
}

type Engine struct {
	//engine可以被抽象成最顶层的group，他应当继承group的所有能力
	*RouterGroup
	//存储的路由组
	groups []*RouterGroup
	//路由树
	router *Router
}

func NewEngine() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	//?
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(pre string) *RouterGroup {
	engine := group.engine
	new_group := &RouterGroup{
		prefix:      pre,
		middlewares: make([]HandlerFunc, 0),
		engine:      engine,
	}
	engine.groups = append(engine.groups, new_group)
	return new_group
}

func (group *RouterGroup) Addroute(mathod string, path string, handler HandlerFunc) {
	pattern := group.prefix + path
	group.engine.AddRouter(mathod, pattern, handler)
}

func (group *RouterGroup) Get(path string, handler HandlerFunc) {
	group.Addroute("GET", path, handler)
}

func (group *RouterGroup) Post(path string, handler HandlerFunc) {
	group.Addroute("POST", path, handler)
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
	var middlewares []HandlerFunc
	for _, group := range g.groups {
		if strings.HasPrefix(rep.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, rep)
	c.Handler = middlewares
	g.router.Handle(c)
}
