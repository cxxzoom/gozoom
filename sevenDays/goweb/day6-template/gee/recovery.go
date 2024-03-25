package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recovery() HandlerFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				ctx.String(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		ctx.Next()
	}
}

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	// 为什么这里是3层：
	// 第0层， runtime.Caller
	// 第1层， trace，即这个函数体
	// 第2层， defer func(){}() ，因为在这个里面调用的trace
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\n TraceBack:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}

	return str.String()
}
