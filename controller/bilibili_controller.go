package controller

import (
	"Spider/database"
	"Spider/model"
	"Spider/spider"
	"Spider/utils"
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

func GetUpInfo(ctx iris.Context) {
	if upId, err := strconv.ParseInt(ctx.URLParam("mid"), 10, 64); err != nil {
		ctx.JSON(utils.NotOk(utils.E_401()))
	} else {
		bilibiliUp, err := model.FindBilibiliUpById(upId)
		if err != nil {
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

func Subscribe(ctx iris.Context) {
	up := new(model.BilibiliUp)
	if err := ctx.ReadJSON(&up); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(utils.NotOk(utils.E_500()))
	} else {
		userId := ctx.Values().Get("userId").(uint)
		user := &model.User{ID: userId}
		fmt.Println(database.DB.Model(&user).Association("BilibiliUps").Append(up).Error)
		ctx.JSON(utils.Ok(database.DB.First(&user, userId).Value, "订阅成功"))
	}
}
