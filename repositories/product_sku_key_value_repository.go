package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var ProductSkuKeyValueRepositories = NewProductSkuKeyRepositories()

type productSkuKeyValueRepositories struct{}

func NewProductSkuKeyRepositories() *productSkuKeyValueRepositories {
	return &productSkuKeyValueRepositories{}
}

// Create 创建
func (this *productSkuKeyValueRepositories) Create(db *gorm.DB, productSkuKeyValue *models.ProductSkuKeyValue) error {
	return db.Create(productSkuKeyValue).Error
}

//批量添加
func (this *productSkuKeyValueRepositories) BatchCreate(db *gorm.DB, productSkuKeyValue []*models.ProductSkuKeyValue) error {
	return db.Create(productSkuKeyValue).Error
}

func (this *productSkuKeyValueRepositories) GetAllData(db *gorm.DB) []*models.ProductSkuKeyValue {
	listData := make([]*models.ProductSkuKeyValue, 0)
	db.Model(models.ProductSkuKeyValue{}).Find(&listData)
	return listData
}
