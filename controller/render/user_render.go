// Package render
// @Author:        asus
// @Description:   $
// @File:          category_render
// @Data:          2022/2/2816:48
//
package render

import (
	"go.uber.org/zap"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"time"
)

type User struct {
	ID          uint   `json:"id"`
	Status      uint   `json:"status"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	AccountType uint   `json:"accountType"`
}

func BuildLoginSuccess(user *models.User) *app.Response {
	//生成Token
	expire := time.Now().Add(2*time.Hour + 30*time.Second)
	token, err := app.GenToken(user.Username, expire)
	if err != nil {
		global.Logger.Error("用户注册生成token失败", zap.Error(err))
		return app.UnauthorizedTokenGenerate
	}
	return app.ResponseData(app.Result{
		"token":    token,
		"expire":   expire.Unix(),
		"duration": 7200,
		"user":     BuildUser(user),
	})
}

func BuildUser(user *models.User) *User {
	if user == nil {
		return nil
	}

	return &User{
		ID:          user.ID,
		Status:      user.Status,
		Username:    user.Username,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		AccountType: user.AccountType,
	}
}

func BuildUsers(users []*models.User) []*User {
	list := make([]*User, 0, len(users))
	if len(users) < 1 {
		return list
	}

	for _, user := range users {
		list = append(list, BuildUser(user))
	}
	return list
}
