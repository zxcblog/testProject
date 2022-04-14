// Package services
// @Author:        asus
// @Description:   $
// @File:          address_service.go
// @Data:          2022/4/1415:42
//
package services

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-project/global"
	"new-project/models"
	"new-project/repositories"
)

var AddressService = NewAddressService()

type addressService struct{}

func NewAddressService() *addressService {
	return &addressService{}
}

func (*addressService) Create(address *models.Address) error {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 修改当前登录用户其他地址的默认地址
		if address.IsDefault {
			err := repositories.AddressRepository.UpdateColumn(tx.Where("user_id", address.UserId), "is_default", false)
			if err != nil {
				return err
			}
		}

		return repositories.AddressRepository.Create(tx, address)
	})

	if err != nil {
		global.Logger.Error("地址创建失败", zap.Any("address", address), zap.Error(err))
		return errors.New("地址创建失败")
	}
	return nil
}

func (*addressService) GetList(userId uint, page, pageSize int) ([]*models.Address, int64) {
	list, total, err := repositories.AddressRepository.GetList(global.DB.Where("user_id", userId), page, pageSize)
	if err != nil {
		global.Logger.Error("获取用户地址列表失败", zap.Uint("userId", userId), zap.Error(err))
	}
	return list, total
}

func (*addressService) Get(addressId uint) *models.Address {
	return repositories.AddressRepository.Get(global.DB, addressId)
}

func (*addressService) Update(address *models.Address) error {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 修改当前登录用户其他地址的默认地址
		if address.IsDefault {
			err := repositories.AddressRepository.UpdateColumn(tx.Where("user_id", address.UserId), "is_default", false)
			if err != nil {
				return err
			}
		}

		return repositories.AddressRepository.Update(tx, address)
	})

	if err != nil {
		global.Logger.Error("用户地址修改失败", zap.Any("address", address), zap.Error(err))
		return errors.New("地址修改失败")
	}
	return nil
}

func (*addressService) Delete(addressId uint) error {
	if err := repositories.AddressRepository.Delete(global.DB, addressId); err != nil {
		global.Logger.Error("地址删除失败", zap.Uint("addressId", addressId), zap.Error(err))
		return errors.New("地址删除失败")
	}
	return nil
}
