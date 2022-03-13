package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var ProductSkuRepositories = NewProductSkuRepositories()

type productSkuRepositories struct{}

func NewProductSkuRepositories() *productSkuRepositories {
	return &productSkuRepositories{}
}

// Create 创建
func (p *productSkuRepositories) Create(db *gorm.DB, productSku *models.ProductSku) error {
	return db.Create(productSku).Error
}

//批量添加
func (p *productSkuRepositories) BatchCreate(db *gorm.DB, productSku *[]models.ProductSku) error {
	return db.Create(productSku).Error
}
