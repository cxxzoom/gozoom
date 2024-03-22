package gee

import (
	"log"
	"net/http"
	"strings"
)

// 暂时先这么写，然后后面再替换

// router 定义router的结构
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// newRouter 实例化
func newRoute() *router {
	return &router{roots: make(map[string]*node), handlers: make(map[string]HandlerFunc)}
}

// parsePattern 解析pattern变成[]string
// 如果是*(通配符)则跳出并返回
func parsePattern(pattern string) []string {
	p := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, part := range p {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}

	}

	return parts
}

// addRoute 解析pattern并添加到路由里
// handlers里面存的方法
// roots里面存的树形结构的node
func (r *router) addRoute(method string, pattern string, fn HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	parts := parsePattern(pattern)

	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = fn
}

// getRoute
func (r *router) getRoute(method string, pattern string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	parts := parsePattern(pattern)
	log.Printf("parts=%+v,pattern=%+v", parts, pattern)
	search := root.search(parts, 0)
	params := make(map[string]string)
	if search != nil {
		parts2 := parsePattern(search.pattern)
		for index, part := range parts2 {
			// 如果是 : 就取冒号后面那一截，并继续匹配
			if part[0] == ':' {
				params[part[1:]] = parts[index]
			}

			// 如果是通配符，就把通配符后面的部分全部加上
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return search, params
	}

	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
