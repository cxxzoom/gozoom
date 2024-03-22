package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(ctx *gee.Context) {
		// start timer
		t := time.Now()
		ctx.Json(http.StatusInternalServerError, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":8989")
}
