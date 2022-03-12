package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var ProductSkuKeyRepositories = NewProductSkuKeyRepositories()

type productSkuKeyRepositories struct{}

func NewProductSkuKeyRepositories() *productSkuKeyRepositories {
	return &productSkuKeyRepositories{}
}

// Create 创建
func (p *productSkuKeyRepositories) Create(db *gorm.DB, productSkuKey *models.ProductSkuKey) error {
	return db.Create(productSkuKey).Error
}
