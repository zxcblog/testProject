// Package repositories
// @Author:        asus
// @Description:   $
// @File:          address_repository
// @Data:          2022/4/1415:47
//
package repositories

import (
	"gorm.io/gorm"
	"new-project/models"
)

var AddressRepository = NewAddressRepository()

type addressRepository struct{}

func NewAddressRepository() *addressRepository {
	return &addressRepository{}
}

func (*addressRepository) Create(db *gorm.DB, model *models.Address) error {
	return db.Create(model).Error
}

func (*addressRepository) GetList(db *gorm.DB, page, pageSize int) ([]*models.Address, int64, error) {
	var total int64
	list := make([]*models.Address, 0, pageSize)
	err := db.Model(models.Address{}).Count(&total).Limit(pageSize).Offset((page - 1) * pageSize).Find(&list).Error

	return list, total, err
}

func (*addressRepository) Get(db *gorm.DB, id uint) *models.Address {
	model := &models.Address{}
	if db.First(model, "id", id).Error != nil {
		return nil
	}
	return model
}

func (*addressRepository) Update(db *gorm.DB, model *models.Address) error {
	return db.Save(model).Error
}

func (*addressRepository) UpdateColumn(db *gorm.DB, field string, val interface{}) error {
	return db.Model(models.Address{}).UpdateColumn(field, val).Error
}

func (*addressRepository) Delete(db *gorm.DB, id uint) error {
	return db.Unscoped().Delete(models.Address{}, "id", id).Error
}
