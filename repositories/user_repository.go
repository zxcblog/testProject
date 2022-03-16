package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var UserRepositories = NewUserRepositories()

type userRepositories struct{}

func NewUserRepositories() *userRepositories {
	return &userRepositories{}
}

// Create 创建
func (u *userRepositories) Create(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

// 根据账号查询用户
func (u *userRepositories) GetUsernameData(db *gorm.DB, userName string) *models.User {
	ret := &models.User{}
	if err := db.First(ret, "username = ?", userName).Error; err != nil {
		return nil
	}

	return ret
}
