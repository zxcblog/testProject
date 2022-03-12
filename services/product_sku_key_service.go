package services

import (
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
	"new-project/repositories"

	"go.uber.org/zap"
)

var ProductSkuKeyService = NewProductSkuKeyService()

type productSkuKeyService struct{}

func NewProductSkuKeyService() *productSkuKeyService {
	return &productSkuKeyService{}
}

func (this *productSkuKeyService) Create(productSkuKey *models.ProductSkuKey) error {
	err := repositories.ProductSkuKeyRepositories.Create(global.DB, productSkuKey)
	if err != nil {
		global.Logger.Error("商品Sku规格添加失败", zap.Error(err))
		return errcode.CreateError.SetMsg("商品Sku规格添加失败")
	}
	return nil
}
