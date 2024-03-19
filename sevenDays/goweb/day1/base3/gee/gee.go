package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 这个是handleFunc的固定形式
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 定义路由和方法的映射
type Engine struct {
	router map[string]HandlerFunc
}

// New  实例化
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 将请求方法，接口地址与其方法对应起来
func (e *Engine) addRoute(method string, pattern string, fn HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = fn
}

// GET add get method to *Engine
func (e *Engine) GET(pattern string, fn HandlerFunc) {
	e.addRoute("GET", pattern, fn)
}

// POST add post method to *Engine
func (e *Engine) POST(pattern string, fn HandlerFunc) {
	e.addRoute("GET", pattern, fn)
}

// Run ListenAndServer
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

// ServeHTTP 接管请求并且找到对应的router对应的方法并执行
// 参考base2，就是接管了所有方法并根据具体的路径执行对应的方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if fn, ok := e.router[key]; ok {
		fn(w, req)
	} else {
		fmt.Fprintf(w, "404 not found: %s\n", req.URL)
	}
}
