package controller

import (
	"Spider/model"
	"Spider/spider"
	"Spider/utils"
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

func GetUpInfo(ctx iris.Context) {
	upId, err := strconv.ParseInt(ctx.URLParam("mid"), 10, 64)
	if err != nil {
		ctx.JSON(utils.NotOk(utils.E_501()))
	} else {
		bilibiliUp := model.FindBilibiliUpById(upId)
		if bilibiliUp.ID == 0 {
			bilibiliUp, err := spider.GetUpInfo(upId)
			if err != nil {
				fmt.Println(err.Error())
				ctx.JSON(utils.NotOk(utils.E_501()))
			} else {
				ctx.JSON(utils.Ok(bilibiliUp, "获取并添加成功"))
			}
		} else {
			ctx.JSON(utils.Ok(bilibiliUp, "获取成功"))
		}
	}
}
