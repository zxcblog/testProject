package services

import (
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
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
		return errcode.CreateError.SetMsg("商品Sku规格添加失败")
	}
	return err
}

func (this *productSkuService) GetProductIdAllData(productId uint) []*models.ProductSku {
	return repositories.ProductSkuRepositories.GetAllData(global.DB.Where("product_id", productId))
}
