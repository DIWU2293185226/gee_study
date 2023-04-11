package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// 封装一个随请求而诞生的Context，对外简化接口，将复杂内容在内部完成
type Context struct {
	//基本请求与响应
	Writer http.ResponseWriter
	Rep    *http.Request
	//请求paylod
	Path   string
	Method string
	Params map[string]string
	//响应
	StatusCodes int
	Handler     []HandlerFunc
	index       int
}

func NewContext(w http.ResponseWriter, rep *http.Request) *Context {
	return &Context{
		Writer: w,
		Rep:    rep,
		Path:   rep.URL.Path,
		Params: make(map[string]string),
		Method: rep.Method,
		index:  -1,
	}
}

// 切到下一个中间件
func (ctx *Context) Next() {
	ctx.index++
	l := len(ctx.Handler)
	for ; ctx.index < l; ctx.index++ {
		ctx.Handler[ctx.index](ctx)
	}
}

// 注册中间件到路由组中
func (group *RouterGroup) Use(handler ...HandlerFunc) {
	group.middlewares = append(group.middlewares, handler...)
}

// 从Context结构体中拿Params参数
func (ctx *Context) Param(key string) string {
	value, _ := ctx.Params[key]
	return value
}

// 从Context结构体中拿取数据
func (ctx *Context) PostForm(key string) string {
	return ctx.Rep.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Rep.URL.Query().Get(key)
}

func (ctx *Context) Status(codes int) {
	ctx.StatusCodes = codes
	ctx.Writer.WriteHeader(codes)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

// 封装不同响应类型
func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
