package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	// 如果没有实现 Engine的ServeHTTP
	// 就需要像原始方法一样，每个接口都 http.HandleFunc("/",indexHandle)
	// 然后再下面用 func indexHandle(w http.ResponseWriter, req *http.Request)
	// 相当于上面是把路由和方法绑定在一起，更加准确一点说应该叫映射
	// 这里相当于用engine来handle这些请求

	/**
	作者说：
	* 使用engine的方式，拦截了所有的http请求，并可以批量和自己处理
	* 比如加一些中间件啊统一的验证啊，路由分组啊之类的，日志等
	**/
	log.Fatal(http.ListenAndServe(":8989", engine))
}
