package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(ctx *Context) {
		// start time
		t := time.Now()
		// process request
		ctx.Next()
		//
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
