package main

import (
	"Spider/config"
	"Spider/controller"
	"Spider/database"
	"Spider/middleware"
	"Spider/model"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func newApp() (app *iris.Application) {
	app = iris.New()
	app.Use(logger.New())

	database.DB.AutoMigrate(
		&model.User{},
		&model.BilibiliUp{},
		&model.BilibiliVideo{},
	)

	//"github.com/iris-contrib/middleware/cors"
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	iris.RegisterOnInterrupt(func() {
		database.DB.Close()
	})

	// 错误处理
	app.OnAnyErrorCode(controller.E_All)
	app.OnErrorCode(iris.StatusNotFound, controller.E_404)
	app.OnErrorCode(iris.StatusInternalServerError, controller.E_500)

	ccm := app.Party("/ccm", crs).AllowMethods(iris.MethodOptions)
	ccm.Post("/user/login", controller.Login)
	ccm.Post("/user/register", controller.Register)
	ccm.PartyFunc("/user", func(users iris.Party) {
		users.Use(middleware.MyJwtMiddleware.Serve, middleware.AuthToken)
		users.Get("/info", controller.UserInfo)
	})

	return
}

func main() {
	app := newApp()
	err := app.Run(iris.Addr(config.Conf.Get("app.port").(string)))
	if err != nil {
		panic(fmt.Sprintf("App Start Error", err.Error()))
	}
}

//package main
//
//import (
//	"github.com/kataras/iris"
//
//	"github.com/iris-contrib/middleware/cors"
//)
//
//func main() {
//	app := iris.New()
//
//	crs := cors.New(cors.Options{
//		AllowedOrigins: []string{"*"}, // allows everything, use that to change the hosts.
//		AllowedHeaders: []string{"*"},
//		AllowCredentials: true,
//	})
//
//	v1 := app.Party("/api/v1", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
//	{
//		v1.Get("/home", func(ctx iris.Context) {
//			ctx.WriteString("Hello from /home")
//		})
//		v1.Get("/about", func(ctx iris.Context) {
//			ctx.WriteString("Hello from /about")
//		})
//		v1.Post("/send", func(ctx iris.Context) {
//			ctx.WriteString("sent")
//		})
//		v1.Put("/send", func(ctx iris.Context) {
//			ctx.WriteString("updated")
//		})
//		v1.Delete("/send", func(ctx iris.Context) {
//			ctx.WriteString("deleted")
//		})
//	}
//
//	app.Run(iris.Addr("localhost:9090"))
//}
