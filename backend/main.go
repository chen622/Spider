package main

import (
	"./model"
	"./sercet"
	"./service/mail"
	"./service/spider"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	db, err := gorm.Open(sercet.DB["type"], sercet.DB["url"])

	db.SingularTable(true)
	if err != nil {
		panic(err)
	}
	app.Handle("GET", "/bili/{mid:int}", func(ctx iris.Context) {
		mid, err := ctx.Params().GetInt64("mid")
		if err != nil {
			panic(err)
		}
		up := model.BilibiliUp{Mid: mid}
		db.First(&up)
		if up.Name == "" {
			ctx.WriteString("No record")
			return
		}

		list, err := spider.GetVideoList(mid)
		if err != nil {
			panic(err)
		}
		fmt.Println(list)
		for _, video := range list {
			db.NewRecord(video)
			db.Create(&video)
			fmt.Println(video)
		}
		if err := mail.Bilibili(up.Name, list); err != nil {
			panic(err)
		}
		fmt.Println("Bilibili ‘", up.Name, "’更新推送成功！")
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
