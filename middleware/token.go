// Package middleware
// @Author:        asus
// @Description:   $
// @File:          token
// @Data:          2022/4/1116:16
//
package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"new-project/cache"
	"new-project/pkg/app"
	"strings"
	"time"
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

	// 验证token是否正确
	myClaims, err := app.ParseToken(token)
	if err != nil {
		fmt.Println(err)
		ctx.StopWithJSON(401, app.UnauthorizedTokenError)
	}

	if time.Now().Unix() > myClaims.ExpiresAt {
		ctx.StopWithJSON(401, app.UnauthorizedTokenTimeout)
	}

	// 获取正在登录中的用户
	user := cache.UserCache.Get(myClaims.Username)
	if user == nil {
		ctx.StopWithJSON(401, app.UnauthorizedAuthNotExist)
	}

	ctx.Values().Set("user", user)
	ctx.Next()
}
