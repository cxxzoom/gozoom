package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义方法
type HandlerFunc func(ctx *Context)

// Engine 定义路由
type Engine struct {
	router *router
}

// 实例化
func New() *Engine {
	return &Engine{router: newRoute()}
}

// GET 添加GET 方法
func (e *Engine) GET(pattern string, fn HandlerFunc) {
	e.router.addRoute("GET", pattern, fn)
}

// POST 添加post方法
func (e *Engine) POST(pattern string, fn HandlerFunc) {
	e.router.addRoute("POST", pattern, fn)
}

// Run 监听端口并执行
func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}

// ServeHTTP 接管请求
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	e.router.handleReq(ctx)
}
