// 封装了快速返回String/Json/Data/HTML的方法
// 提供了访问Query和PostForm表单的接口

package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context 封装头以及返回，获取路由里面的参数等等
type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
	handlers   []HandlerFunc
	index      int
}

// newContext 实例化
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// Next 执行下一个方法
// 这里有把所有
func (ctx *Context) Next() {
	ctx.index++
	s := len(ctx.handlers)
	for ; ctx.index < s; ctx.index++ {
		ctx.handlers[ctx.index](ctx)
	}
}

// Param 获取 context.params里面的某个值
func (ctx *Context) Param(key string) string {
	val := ctx.Params[key]
	return val
}

// PostForm 获取请From求里面的某一个参数
func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

// Query 从GET请求里获取某个参数
func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

// Status 写入状态码，并向返回头写入相同的返回码
func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

// SetHeader 往header里面写东西
func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

// String 设置返回参数的类型 这里是 text/plain
func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// Json 设置返回参数的类型 这里是 application/json
func (ctx *Context) Json(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// Data 直接写就行了？
func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

// HTML 返回HTML
func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
