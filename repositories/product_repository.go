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
func (p *productRepositories) Create(db *gorm.DB, product *models.Product) (uint, error) {
	if db.Create(product).Error != nil {
		return 0, db.Create(product).Error
	}
	return product.ID, nil
}
