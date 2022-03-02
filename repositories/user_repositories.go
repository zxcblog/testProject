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
