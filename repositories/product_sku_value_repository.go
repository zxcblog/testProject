package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var ProductSkuValueRepositories = NewProductSkuValueRepositories()

type productSkuValueRepositories struct{}

func NewProductSkuValueRepositories() *productSkuValueRepositories {
	return &productSkuValueRepositories{}
}

// Create 创建
func (p *productSkuValueRepositories) BatchCreate(db *gorm.DB, productSkuValue *[]models.ProductSkuValue) error {
	return db.Create(productSkuValue).Error
}
