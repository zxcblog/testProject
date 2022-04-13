// Package middleware
// @Author:        asus
// @Description:   $
// @File:          token
// @Data:          2022/4/1116:16
//
package middleware

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"new-project/pkg/app"
	"new-project/services"
	"strings"
)

func Token(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		token = ctx.FormValue("Authorization")
	}

	// 没有上传token时， 让用户登录
	if strings.TrimSpace(token) == "" {
		ctx.StopWithJSON(401, app.UnauthorizedAuthNotExist)
	}

	// 获取当前正在登录用户
	user, err := services.UserTokenService.GetTokenUser(token)
	if err != nil {
		ctx.StopWithJSON(http.StatusUnauthorized, err)
	}

	ctx.Values().Set("user", user)
	ctx.Next()
}
