package controller

import (
	"Spider/utils"
	"github.com/kataras/iris"
)

func E_404(ctx iris.Context) {
	ctx.JSON(utils.NotOk(utils.E_404()))
}

func E_500(ctx iris.Context) {
	ctx.JSON(utils.NotOk(utils.E_500()))
}

func E_All(ctx iris.Context) {
	ctx.JSON(utils.NotOk(utils.E_All()))
}
