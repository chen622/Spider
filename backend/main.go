package main

import (
	"./service/mail"
	"./service/spider"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
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
		author, list, err := spider.GetVideoList(mid)
		if err != nil {
			panic(err)
		}
		fmt.Println(list)
		if err := mail.Bilibili(author, list); err != nil {
			panic(err)
		}
		fmt.Println("Bilibili ‘", author, "’更新推送成功！")
		ctx.HTML("<h1>推送成功</h1>")
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
