package gee

import "net/http"

// 原本以为是这样的，但是这节改了一下，变成了传参是*Context，在gee.go里面定义的
// type HandlerFunc func(w http.ResponseWriter, r *http.Request)
type router struct {
	route map[string]HandlerFunc
}

// newRoute 返回路由，并返回引用
func newRoute() *router {
	return &router{route: make(map[string]HandlerFunc)}
}

// addRoute 把路由和方法添加到router的映射里
func (r *router) addRoute(method string, pattern string, fn HandlerFunc) {
	key := method + "-" + pattern
	r.route[key] = fn
}

// handleReq 接收请求并根据根据路由找到对应的方法并处理
func (r *router) handleReq(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if fn, ok := r.route[key]; ok {
		fn(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
