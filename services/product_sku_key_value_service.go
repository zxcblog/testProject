package services

import (
	"errors"
	"new-project/global"
	"new-project/models"
	"new-project/repositories"

	"go.uber.org/zap"
)

var ProductSkuKeyValueService = NewProductSkuKeyValueService()

type productSkuKeyValueService struct{}

func NewProductSkuKeyValueService() *productSkuKeyValueService {
	return &productSkuKeyValueService{}
}

func (this *productSkuKeyValueService) Create(productSkuKeyValue *models.ProductSkuKeyValue) error {
	err := repositories.ProductSkuKeyValueRepositories.Create(global.DB, productSkuKeyValue)
	if err != nil {
		global.Logger.Error("商品Sku的key和value添加失败", zap.Error(err))
		return errors.New("商品Sku的key和value添加失败")
	}
	return nil
}

func (this *productSkuKeyValueService) BatchCreate(productSkuKeyValue []*models.ProductSkuKeyValue) error {
	err := repositories.ProductSkuKeyValueRepositories.BatchCreate(global.DB, productSkuKeyValue)
	if err != nil {
		global.Logger.Error("商品Sku的key和value添加失败", zap.Error(err))
		return errors.New("商品Sku的key和value添加失败")
	}
	return nil
}

func (this *productSkuKeyValueService) GetProductIdAllData(productId uint) []*models.ProductSkuKeyValue {
	return repositories.ProductSkuKeyValueRepositories.GetAllData(global.DB.Where("product_id  = ?", productId))
}
