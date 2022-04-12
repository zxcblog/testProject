package services

import (
	"errors"
	"new-project/global"
	"new-project/models"
	"new-project/repositories"

	"go.uber.org/zap"
)

var ProductSkuService = NewProductSkuService()

type productSkuService struct{}

func NewProductSkuService() *productSkuService {
	return &productSkuService{}
}

func (this *productSkuService) BatchCreate(SkuSliceData []*models.ProductSku) error {
	err := repositories.ProductSkuRepositories.BatchCreate(global.DB, SkuSliceData)
	if err != nil {
		global.Logger.Error("商品Sku规格添加失败", zap.Error(err))
		return errors.New("商品Sku规格添加失败")
	}
	return err
}

func (this *productSkuService) GetProductIdAllData(productId uint) []*models.ProductSku {
	return repositories.ProductSkuRepositories.GetAllData(global.DB.Where("product_id", productId))
}
