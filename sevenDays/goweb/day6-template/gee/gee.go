package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
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
	router        *router
	groups        []*RouterGroup
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

// New 实例化
func New() *Engine {
	engine := &Engine{router: newRoute()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 添加分组
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

// addRoute
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

// Use 把中间件加到group里
func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// createStaticHandler 找到文件的真实路径并交给htt.FileSystem
// http.FileSystem已经实现了怎么返回一个文件
func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(ctx *Context) {
		file := ctx.Param("filepath")

		if _, err := fs.Open(file); err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(ctx.Writer, ctx.Req)
	}
}

// Static 提供静态文件服务
func (g *RouterGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handler
	g.GET(urlPattern, handler)
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
	ctx.engine = e
	e.router.handle(ctx)
}

// SetFuncMap 设置啥子模板方法？
func (e *Engine) SetFuncMap(funcMap template.FuncMap) {
	e.funcMap = funcMap
}

// LoadHTMLGlob 晓不得这个是搞什么的
func (e *Engine) LoadHTMLGlob(pattern string) {
	e.htmlTemplates = template.Must(template.New("").Funcs(e.funcMap).ParseGlob(pattern))
}
