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

func (u *productService) Create(product *models.Product) (uint, error) {
	productId, err := repositories.ProductRepositories.Create(global.DB, product)
	if err != nil {
		global.Logger.Error("添加失败", zap.Error(err))
		return 0, errcode.CreateError.SetMsg("添加失败")
	}
	return productId, nil
}
