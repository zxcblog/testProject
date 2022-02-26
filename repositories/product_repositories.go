// Package repositories
// @Author:        asus
// @Description:   $
// @File:          product_repositories
// @Data:          2022/2/2614:24
//
package repositories

import (
	"gorm.io/gorm"
	"new-project/models"
)

type IPorduct interface {
	Insert(*gorm.DB, *models.Product) (int64, error)
	Delete(*gorm.DB, uint) bool
	Update(*gorm.DB, *models.Product) error
	SelectByKey(*gorm.DB, uint) *models.Product
	SelectAll(*gorm.DB) ([]*models.Product, error)
}

type ProductRepositories struct{}

func NewProductRepositories() IPorduct {
	return &ProductRepositories{}
}

func (p *ProductRepositories) Insert(db *gorm.DB, product *models.Product) (int64, error) {
	panic("implement me")
}

func (p *ProductRepositories) Delete(db *gorm.DB, id uint) bool {
	panic("implement me")
}

func (p *ProductRepositories) Update(db *gorm.DB, product *models.Product) error {
	panic("implement me")
}

func (p *ProductRepositories) SelectByKey(db *gorm.DB, id uint) *models.Product {
	panic("implement me")
}

func (p *ProductRepositories) SelectAll(db *gorm.DB) ([]*models.Product, error) {
	panic("implement me")
}
