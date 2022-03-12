// Package repositories
// @Author:        asus
// @Description:   $
// @File:          product_repositories
// @Data:          2022/2/2614:24
//
package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var ProductRepositories = NewProductRepositories()

type productRepositories struct{}

func NewProductRepositories() *productRepositories {
	return &productRepositories{}
}

// Create 创建
func (p *productRepositories) Create(db *gorm.DB, product *models.Product) error {
	return db.Create(product).Error
}
