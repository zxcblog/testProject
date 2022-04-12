package services

import (
	"errors"
	"new-project/global"
	"new-project/models"
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
		return errors.New("商品添加失败")
	}
	return nil
}

func (this *productService) Get(id uint) *models.Product {
	return repositories.ProductRepositories.GetProductByID(global.DB, id)
}

func (this *productService) GetListPage(params map[string]string, page, pageSize int) ([]*models.Product, int64) {
	return repositories.ProductRepositories.GetListPage(global.DB, params, page, pageSize)
}
