// Package api
// @Author:        asus
// @Description:   $
// @File:          address_controller
// @Data:          2022/4/1116:17
//
package api

import (
	"github.com/kataras/iris/v12"
	"new-project/global"
	"new-project/models"
	"new-project/models/form"
	"new-project/pkg/app"
	"new-project/services"
)

type AddressController struct {
	Ctx iris.Context
}

// Post 新增地址
// @Summary      新增地址
// @Description  新增地址
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @param        root       body  form.Address  true  "地址管理"
// @Tags         用户地址
// @Success      200  {object}  app.Response{}
// @Router       /address [post]
func (this *AddressController) Post() *app.Response {
	param := &form.Address{}
	if err := app.FormValueJson(this.Ctx, global.Validate, param); err != nil {
		return err
	}

	user, err := app.GetUser(this.Ctx)
	if err != nil {
		return err
	}

	address := &models.Address{
		IsDefault:     param.IsDefault,
		UserId:        user.ID,
		Province:      param.Province,
		City:          param.City,
		Area:          param.Area,
		Street:        param.Street,
		Desc:          param.Desc,
		ContactName:   param.ContactName,
		ContactMobile: param.ContactMobile,
	}
	if err := services.AddressService.Create(address); err != nil {
		return app.CreateError.SetMsg(err.Error())
	}
	return app.Success
}

// Get 获取地址列表
// @Summary      地址列表
// @Description  地址列表
// @Produce      json
// @Security     ApiKeyAuth
// @param        page      query  int  false  "页数"
// @param        pageSize  query  int  false  "条数"
// @Tags         用户地址
// @Success      200  {object}  app.Response{}
// @Router       /address [get]
func (this *AddressController) Get() *app.Response {

	page := app.GetPage(this.Ctx)
	pageSize := app.GetPageSize(this.Ctx)
	user, err := app.GetUser(this.Ctx)
	if err != nil {
		return err
	}

	list, total := services.AddressService.GetList(user.ID, page, pageSize)
	return app.ResponseList(list, total)
}

// PutBy 修改地址
// @Summary      修改地址
// @Description  修改地址
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @param        addressId  path  uint  true  "地址id"
// @param        root  body  form.Address  true  "地址管理"
// @Tags         用户地址
// @Success      200  {object}  app.Response{}
// @Router       /address/{addressId} [put]
func (this *AddressController) PutBy(addressId uint) *app.Response {
	param := &form.Address{}
	if err := app.FormValueJson(this.Ctx, global.Validate, param); err != nil {
		return err
	}

	user, err := app.GetUser(this.Ctx)
	if err != nil {
		return err
	}

	address := services.AddressService.Get(addressId)
	if address == nil || address.UserId != user.ID {
		return app.NotFound
	}

	address.IsDefault = param.IsDefault
	address.Province = param.Province
	address.City = param.City
	address.Area = param.Area
	address.Street = param.Street
	address.Desc = param.Desc
	address.ContactName = param.ContactName
	address.ContactMobile = param.ContactMobile
	if err := services.AddressService.Update(address); err != nil {
		return app.UpdateError.SetMsg(err.Error())
	}
	return app.Success
}

// DeleteBy 删除地址
// @Summary      删除地址
// @Description  删除地址
// @Produce      json
// @Security     ApiKeyAuth
// @param        addressId  path  uint          true  "地址id"
// @Tags         用户地址
// @Success      200  {object}  app.Response{}
// @Router       /address/{addressId} [delete]
func (this *AddressController) DeleteBy(addressId uint) *app.Response {
	user, err := app.GetUser(this.Ctx)
	if err != nil {
		return err
	}

	address := services.AddressService.Get(addressId)
	if address == nil || address.UserId != user.ID {
		return app.Success
	}

	if err := services.AddressService.Delete(addressId); err != nil {
		return app.DelError.SetMsg(err.Error())
	}
	return app.Success
}
