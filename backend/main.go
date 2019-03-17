package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"iris-start/spider"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	//输出html
	// 请求方式: GET
	// 访问地址: http://localhost:8080/welcome
	app.Handle("GET", "/bili/{mid:int}", func(ctx iris.Context) {
		mid, err := ctx.Params().GetInt("mid")
		if err != nil {
			panic(err)
		}
		list, err := spider.GetVideoList(mid)
		if err != nil {
			panic(err)
		}
		html := ""
		for _, video := range list {
			html += fmt.Sprint("<h1>",video.Title,"<\\h1>")
			html += fmt.Sprint("<h3>视频发布于 ",video.GetTime().Format("2006-01-02 15:04:05"),"<\\h3>")
			html += "<\\br><\\br>"
		}
		ctx.HTML(html)
	})
	//输出字符串
	// 类似于 app.Handle("GET", "/ping", [...])
	// 请求方式: GET
	// 请求地址: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	//输出json
	// 请求方式: GET
	// 请求地址: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	app.Run(iris.Addr(":9090")) //8080 监听端口
}
