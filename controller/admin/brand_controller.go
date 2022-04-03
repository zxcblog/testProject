// Package admin
// @Author:        asus
// @Description:   $
// @File:          category_controller
// @Data:          2022/2/2710:44
//
package admin

import (
	"github.com/kataras/iris/v12"
	"new-project/controller/render"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/errcode"
	"new-project/services"
)

type BrandController struct {
	Ctx iris.Context
}

// BrandRequest 创建品牌
type BrandRequest struct {
	BrandName   string `validate:"required,min=1,max=20" label:"品牌名称" json:"brandName"` // 品牌名称
	BrandImg    string `validate:"-" label:"品牌图片" json:"brandImg" default:""`           // 品牌图片地址链接
	Description string `validate:"-" label:"品牌描述" json:"description" default:""`        // 品牌描述
	CategoryID  uint   `validate:"required" label:"所属分类" json:"categoryID"`             // 所属分类
}

// Post 添加商品品牌
// @Summary      添加商品品牌
// @Description  后台管理人员添加商品品牌
// @Accept       json
// @Produce      json
// @param        root  body  BrandRequest  true  "添加商品品牌"
// @Tags         品牌
// @Success      200  {object}  app.Response{data=render.Brand}
// @Router       /admin/brand [post]
func (b *BrandController) Post() *app.Response {
	param := &BrandRequest{}
	if err := app.FormValueJson(b.Ctx, global.Validate, param); err != nil {
		return app.ToResponseErr(err)
	}

	brand := &models.Brand{
		BrandName:   param.BrandName,
		BrandImg:    param.BrandImg,
		Description: param.Description,
		CategoryID:  param.CategoryID,
	}

	if err := services.BrandService.Create(brand); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseData(render.BuildBrand(brand))
}

// Get 获取品牌列表
// @Summary      获取品牌列表
// @Description  获取品牌列表
// @Produce      json
// @param        categoryID  query  uint  false  "分类id"  default(0)
// @param        page        query  uint  false  "分页"    default(1)
// @param        pageSize    query  uint  false  "分页页数"  default(10)
// @Tags         品牌
// @Success      200  {object}  app.Response{data=[]render.Brand}
// @Router       /admin/brand [get]
func (b *BrandController) Get() *app.Response {
	categoryID := app.FormValueUintDefault(b.Ctx, "categoryID", 0)
	page := app.GetPage(b.Ctx)
	pageSize := app.GetPageSize(b.Ctx)
	list, total := services.BrandService.GetListPage(categoryID, page, pageSize)

	return app.ToResponseList(render.BuildBrands(list), total)
}

// GetBy 通过id获取品牌信息
// @Summary      获取品牌详情
// @Description  通过品牌id获取品牌详情
// @Produce      json
// @param        brandID  path  uint  true  "品牌id"
// @tags         品牌
// @Success      200  {object}  app.Response{data=render.Brand}
// @Router       /admin/brand/{brandID} [get]
func (b *BrandController) GetBy(id uint) *app.Response {
	res := services.BrandService.Get(id)

	if res == nil {
		return app.ToResponseErr(errcode.NotFound)
	}
	return app.ResponseData(render.BuildBrand(res))
}

// PutBy 修改品牌信息
// @Summary      修改品牌信息
// @Description  修改品牌信息
// @Accept       json
// @Produce      json
// @param        brandID  path  uint          true  "品牌ID"
// @param        root     body  BrandRequest  true  "修改品牌信息"
// @Tags         品牌
// @Success      200  {object}  app.Response{data=render.Brand}
// @Router       /admin/brand/{brandID} [put]
func (b *BrandController) PutBy(id uint) *app.Response {
	brand := services.BrandService.Get(id)
	if brand == nil {
		return app.ToResponseErr(errcode.NotFound)
	}

	param := &BrandRequest{}
	if err := app.FormValueJson(b.Ctx, global.Validate, param); err != nil {
		return app.ToResponseErr(err)
	}

	brand.BrandName = param.BrandName
	brand.BrandImg = param.BrandImg
	brand.Description = param.Description
	brand.CategoryID = param.CategoryID

	if err := services.BrandService.Update(brand); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseMsg("param")
}

// DeleteBy 删除品牌信息
// @Summary      删除品牌信息
// @Description  删除品牌信息
// @Produce      json
// @param        brandID  path  uint  true  "品牌ID"
// @Tags         品牌
// @Success      200  {object}  app.Response
// @Router       /admin/brand/{brandID} [delete]
func (b *BrandController) DeleteBy(id uint) *app.Response {
	if err := services.BrandService.Delete(id); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseMsg("删除成功")
}
