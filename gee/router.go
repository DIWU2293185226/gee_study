package gee

import (
	"fmt"
	"strings"
)

type Router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*Node
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*Node),
	}
}

// 拆分pattern串，直到发现*或者拆分完成
func ParsePattern(pattern string) []string {
	parts := make([]string, 0)
	vs := strings.Split(pattern, "/")
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 注册路由，注册函数&&将注册路由插入tire树
func (r *Router) AddRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	parts := ParsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Node{}
	}
	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// 根据请求方式和路径，拿到路由节点和url中的参数
func (r *Router) GetRouter(method string, path string) (*Node, map[string]string) {
	path_parse := ParsePattern(path)
	params := make(map[string]string)
	n, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	//搜索路由是否已经注册
	result := n.Search(path_parse, 0)
	//拿参数
	if result != nil {
		//将匹配到的路由拆分方便处理
		parts := ParsePattern(result.Pattern)
		//拿params的传参
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = path_parse[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(path_parse[index:], "/")
				break
			}
		}
		return result, params

	}
	return nil, nil

}

func (r *Router) Handle(ctx *Context) {
	//用GetRouter进行路由匹配，n不为空则匹配到，这样实现了字典树的动态匹配
	n, param := r.GetRouter(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = param
		pattern := ctx.Rep.Method + "-" + n.Pattern
		r.handlers[pattern](ctx)
	} else {
		fmt.Fprintf(ctx.Writer, "404,NOT FOUND:%v", ctx.Rep.URL)
	}
	// if handler, ok := r.handlers[pattern]; ok {
	// 	handler(ctx)
	// } else {
	// 	fmt.Fprintf(ctx.Writer, "404,NOT FOUND:%v", ctx.Rep.URL)
	// }
}

/*
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
*/
