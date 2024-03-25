package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.Use(gee.Recovery())

	r.GET("/", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello geek\n")
	})

	r.GET("/panic", func(ctx *gee.Context) {
		names := []string{"geektutu"}
		ctx.String(http.StatusOK, names[100])
	})

	r.Run(":8989")
}
