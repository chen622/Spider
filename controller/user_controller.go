package controller

import (
	"Spider/middleware"
	"Spider/model"
	"Spider/utils"
	"fmt"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

func Login(ctx iris.Context) {
	user := new(model.User)
	if err := ctx.ReadJSON(&user); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(utils.NotOk(utils.E_500()))
	} else {
		var o = map[string]interface{}{
			"token": middleware.CreateToken(123),
		}
		ctx.JSON(o)
	}
}

func Register(ctx iris.Context) {
	user := new(model.User)

	if err := ctx.ReadJSON(&user); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(utils.NotOk(utils.E_500()))
	} else {
		err := utils.Validate.Struct(user)
		if err != nil {
			errMap := map[string]interface{}{}
			for _, err := range err.(validator.ValidationErrors) {
				errMap[err.Field()] = err.ActualTag()
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.Type())
				fmt.Println(err.Param())
				fmt.Println(err.ActualTag())
			}
			ctx.JSON(utils.NotOk(utils.E_402(errMap)))
		} else {
			if model.FindUserByUsername(user.Username) != nil {

			}
			u, err := model.CreateUser(user)
			if err != nil {
				ctx.JSON(utils.NotOk(utils.E_500()))
			} else {
				ctx.StatusCode(iris.StatusOK)
				ctx.JSON(utils.Ok(u, ""))
			}
		}
	}
}

func UserInfo(ctx iris.Context) {
	userId := ctx.Values().Get("userId").(uint)
	ctx.JSON(utils.Ok(userId, "success"))
}
