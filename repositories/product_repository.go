// Package repositories
// @Author:        asus
// @Description:   $
// @File:          product_repositories
// @Data:          2022/2/2614:24
//
package repositories

import (
	"new-project/models"

	"gorm.io/gorm"
)

var ProductRepositories = NewProductRepositories()

type productRepositories struct{}

func NewProductRepositories() *productRepositories {
	return &productRepositories{}
}

// Create 创建
func (p *productRepositories) Create(db *gorm.DB, product *models.Product) error {
	return db.Create(product).Error
}

func (p *productRepositories) GetProductByID(db *gorm.DB, id uint) *models.Product {
	return p.GetWhereFirst(db.Where("id = ?", id))
}

// GetWhereFirst 通过条件查找某一个
func (p *productRepositories) GetWhereFirst(db *gorm.DB) *models.Product {
	ret := &models.Product{}
	if err := db.First(ret).Error; err != nil {
		return nil
	}

	return ret
}

func (p *productRepositories) GetListPage(db *gorm.DB, params map[string]string, page, pageSize int) ([]*models.Product, int64) {
	list := make([]*models.Product, 0, pageSize)
	var total int64
	dbModel := db.Model(models.Product{})
	dbModel = p.BuildParams(dbModel, params)

	dbModel.Count(&total).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&list)

	return list, total
}

func (p *productRepositories) BuildParams(db *gorm.DB, params map[string]string) *gorm.DB {
	if params["categoryID"] != "" {
		db.Where("category_id = ?", params["categoryID"])
	}

	if params["brandId"] != "" {
		db.Where("brand_id = ?", params["brandId"])
	}

	if params["productName"] != "" {
		db.Where("product_name LIKE ?", "%"+params["productName"]+"%")
	}

	return db
}
