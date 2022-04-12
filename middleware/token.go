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
	"new-project/pkg/app"
	"strings"
)

func Token(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		token = ctx.FormValue("Authorization")
	}

	// 没有上传token时， 让用户登录
	if strings.TrimSpace(token) == "" {
		ctx.JSON(app.UnauthorizedAuthNotExist)
		ctx.StopExecution()
	}

	// 验证token是否正确
	myClaims, err := app.ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myClaims)

	ctx.Next()
}
