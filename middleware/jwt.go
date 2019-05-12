package middleware

import (
	"Spider/config"
	"Spider/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"time"
)

var (
	MyJwtMiddleware = New()
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func New() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Conf.Get("secret.key").(string)), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}

func AuthToken(ctx iris.Context) {
	userToken := MyJwtMiddleware.Get(ctx)
	if claims, ok := userToken.Claims.(jwt.MapClaims); ok {
		if userToken.Valid {
			fmt.Println(claims["userId"])
			userId := uint(claims["userId"].(float64))
			ctx.Values().Set("userId", userId)
			ctx.Next()
		} else {
			ctx.StatusCode(401)
			ctx.JSON(utils.TokenInvalid("token 已经过期"))
		}
	} else {
		ctx.StatusCode(401)
		ctx.JSON(utils.TokenInvalid("token 无效"))
		return
	}
}

func CreateToken(userId uint) (tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":    time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Conf.Get("secret.key").(string)))
	if err != nil {
		panic(fmt.Sprintln(err.Error()))
	}
	return
}
