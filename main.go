package main

import (
	"Spider/config"
	"Spider/controller"
	"Spider/database"
	"Spider/middleware"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func newApp() (app *iris.Application) {
	app = iris.New()
	app.Use(logger.New())

	//database.DB.AutoMigrate(
	//	&model.User{},
	//	&model.BilibiliUp{},
	//	&model.BilibiliVideo{},
	//)

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
	ccm.Get("/user/check-token", middleware.MyJwtMiddleware.Serve, controller.CheckToken)
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
