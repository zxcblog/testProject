package services

import (
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
	"new-project/repositories"

	"go.uber.org/zap"
)

var ProductService = NewProductService()

type productService struct{}

func NewProductService() *productService {
	return &productService{}
}

//添加商品主信息
func (this *productService) Create(product *models.Product) error {
	err := repositories.ProductRepositories.Create(global.DB, product)
	if err != nil {
		global.Logger.Error("商品添加失败", zap.Error(err))
		return errcode.CreateError.SetMsg("商品添加失败")
	}
	return nil
}
