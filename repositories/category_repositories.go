// Package repositories
// @Author:        asus
// @Description:   $
// @File:          category_repositories
// @Data:          2022/2/2815:58
//
package repositories

import (
	"gorm.io/gorm"
	"new-project/models"
)

var CategoryRepositories = NewCategoryRepositories()

type categoryRepositories struct{}

func NewCategoryRepositories() *categoryRepositories {
	return &categoryRepositories{}
}

func (c *categoryRepositories) Get(db *gorm.DB, id uint) *models.Category {
	ret := &models.Category{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

func (c *categoryRepositories) GetList(db *gorm.DB, page, pageSize int) ([]*models.Category, int64) {
	list := make([]*models.Category, pageSize)
	var total int64
	db.Model(models.Category{}).Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&list)
	return list, total
}

func (c *categoryRepositories) Create(db *gorm.DB, category *models.Category) error {
	return db.Create(category).Error
}

func (c *categoryRepositories) Update(db *gorm.DB, category *models.Category) error {
	return db.Save(category).Error
}

func (c *categoryRepositories) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&models.Category{}, "id = ?", id).Error
}

func (c *categoryRepositories) DeleteWhere(db *gorm.DB) error {
	return db.Delete(&models.Category{}).Error
}
