package main

import (
	"fmt"
	"gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("template/*")
	// 绑定服务器上的文件
	r.Static("/assets", "./static")

	stu1 := &student{
		Name: "xxx1",
		Age:  20,
	}

	stu2 := &student{
		Name: "xxx2",
		Age:  22,
	}

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "custom_fun.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2024, 3, 22, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":8989")
}
