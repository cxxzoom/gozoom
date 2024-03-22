package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc 定义方法
type HandlerFunc func(ctx *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// Engine 定义路由
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// 实例化
func New() *Engine {
	engine := &Engine{router: newRoute()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: engine,
	}
	g.engine.groups = append(g.engine.groups, newGroup)

	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, fn HandlerFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.router.addRoute(method, pattern, fn)
}

// GET 添加GET 方法
func (g *RouterGroup) GET(pattern string, fn HandlerFunc) {
	g.addRoute("GET", pattern, fn)
}

// POST 添加post方法
func (g *RouterGroup) POST(pattern string, fn HandlerFunc) {
	g.addRoute("POST", pattern, fn)
}

// Use  把中间件加到group里
func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// Run 监听端口并执行
func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}

// ServeHTTP 接管请求
// 当请求进来的时候，获取所有的中间件，并把中间件放在ctx.handlers里面
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc

	for _, group := range e.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	ctx := newContext(w, req)
	ctx.handlers = middlewares
	e.router.handle(ctx)
}
