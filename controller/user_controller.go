package controller

import (
	"Spider/middleware"
	"Spider/model"
	"Spider/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx iris.Context) {
	user := new(model.User)
	if err := ctx.ReadJSON(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(utils.NotOk(utils.E_500()))
	} else {
		rightUser, err := model.UserAdminCheckLogin(user.Username)
		if ok := bcrypt.CompareHashAndPassword([]byte(rightUser.Password), []byte(user.Password)); ok != nil || err != nil {
			fmt.Println(err)
			ctx.JSON(utils.NotOk(utils.E_403()))
			ctx.Next()
		} else {
			var o = map[string]interface{}{
				"token": middleware.CreateToken(rightUser.ID),
				"user":  rightUser,
			}
			ctx.JSON(utils.Ok(o, "登录成功!"))
		}
	}
}

func Register(ctx iris.Context) {
	ctx.JSON(utils.NotOk(utils.E_201()))
	//user := new(model.User)
	//if err := ctx.ReadJSON(&user); err != nil {
	//	fmt.Println(err.Error())
	//	ctx.JSON(utils.NotOk(utils.E_500()))
	//} else {
	//	err := utils.Validate.Struct(user)
	//	if err != nil { //表单验证失败
	//		errMap := map[string]interface{}{}
	//		for _, err := range err.(validator.ValidationErrors) {
	//			errMap[err.Field()] = err.ActualTag()
	//			fmt.Println(err.Namespace())
	//			fmt.Println(err.Field())
	//			fmt.Println(err.Type())
	//			fmt.Println(err.Param())
	//			fmt.Println(err.ActualTag())
	//		}
	//		ctx.JSON(utils.NotOk(utils.E_402(errMap)))
	//	} else {
	//		oldUser := model.FindUserByUsername(user.Username)
	//		if oldUser.ID != 0 { //用户名存在
	//			ctx.JSON(utils.NotOk(utils.E_405()))
	//		} else { //注册成功
	//			u, err := model.CreateUser(user)
	//			if err != nil {
	//				ctx.JSON(utils.NotOk(utils.E_500()))
	//			} else {
	//				ctx.JSON(utils.Ok(u, "注册成功"))
	//			}
	//		}
	//	}
	//}
}

func CheckToken(ctx iris.Context) {
	userToken := middleware.MyJwtMiddleware.Get(ctx)
	if _, ok := userToken.Claims.(jwt.MapClaims); ok {
		if userToken.Valid {
			ctx.JSON(utils.Ok(nil, ""))
		} else {
			ctx.JSON(utils.TokenInvalid("token 已经过期"))
		}
	} else {
		ctx.JSON(utils.TokenInvalid("token 无效"))
		return
	}
}

func UserInfo(ctx iris.Context) {
	userId := ctx.Values().Get("userId").(uint)
	user := model.FindUserById(userId)
	if user == nil {
		ctx.JSON(utils.NotOk(utils.E_500()))
	}
	ctx.JSON(utils.Ok(model.FindUserById(userId), "success"))
}
