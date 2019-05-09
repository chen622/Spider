package controller

import (
	"Spider/middleware"
	"Spider/utils"
	"github.com/kataras/iris"
)

func Login(ctx iris.Context) {
	var o = map[string]interface{}{
		"token": middleware.CreateToken(123),
	}
	ctx.JSON(o)
}

func Register(ctx iris.Context) {

}

func UserInfo(ctx iris.Context) {
	userId := ctx.Values().Get("userId").(uint)
	ctx.JSON(utils.Ok(userId, "success"))
}
