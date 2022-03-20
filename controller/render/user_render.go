// Package render
// @Author:        asus
// @Description:   $
// @File:          category_render
// @Data:          2022/2/2816:48
//
package render

import (
	"new-project/models"
)

type User struct {
	ID          uint   `json:"id" example:"1"`
	Status      uint   `json:"status" example:"1"`
	Username    string `json:"username" example:"zxc7310"`
	Nickname    string `json:"nickname" example:"周晓钏"`
	Avatar      string `json:"avatar" example:""`
	AccountType uint   `json:"accountType" example:"1"`
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
