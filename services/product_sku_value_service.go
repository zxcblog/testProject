package services

import (
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
	"new-project/repositories"

	"go.uber.org/zap"
)

var ProductSkuValueService = NewProductSkuValueService()

type productSkuValueService struct{}

func NewProductSkuValueService() *productSkuValueService {
	return &productSkuValueService{}
}

func (this *productSkuValueService) BatchCreate(productSkuValue *[]models.ProductSkuValue) error {
	err := repositories.ProductSkuValueRepositories.BatchCreate(global.DB, productSkuValue)
	if err != nil {
		global.Logger.Error("商品Sku规格添加失败", zap.Error(err))
		return errcode.CreateError.SetMsg("商品Sku规格添加失败")
	}
	return err
}
