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

// Get 通过id进行搜索
func (c *categoryRepositories) Get(db *gorm.DB, id uint) *models.Category {
	return c.GetWhereFirst(db.Where("id = ?", id))
}

// GetWhereFirst 通过条件查找某一个
func (c *categoryRepositories) GetWhereFirst(db *gorm.DB) *models.Category {
	ret := &models.Category{}
	if err := db.First(ret).Error; err != nil {
		return nil
	}

	return ret
}

// GetWhereList 通过条件查找集合
func (c *categoryRepositories) GetWhereList(db *gorm.DB) ([]*models.Category, error) {
	ret := []*models.Category{}
	if err := db.Find(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

// GetList 获取分类列表
func (c *categoryRepositories) GetList(db *gorm.DB, page, pageSize int) ([]*models.Category, int64) {
	list := make([]*models.Category, 0, pageSize)
	var total int64
	db.Model(models.Category{}).Count(&total).
		Order("sort desc").
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&list)
	return list, total
}

// Create 创建
func (c *categoryRepositories) Create(db *gorm.DB, category *models.Category) error {
	return db.Create(category).Error
}

// Update 修改
func (c *categoryRepositories) Update(db *gorm.DB, category *models.Category) error {
	return db.Save(category).Error
}

// Delete 通过id删除
func (c *categoryRepositories) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&models.Category{}, "id = ?", id).Error
}

// DeleteWhere 通过where条件删除
func (c *categoryRepositories) DeleteWhere(db *gorm.DB) error {
	return db.Delete(&models.Category{}).Error
}
