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

type CategoryController struct {
	Ctx iris.Context
}

// CategoryRequest 创建分类
type CategoryRequest struct {
	CategoryName string `validate:"required,min=1,max=20" label:"分类名称" json:"categoryName"` // 分类名称
	CategoryImg  string `validate:"-" label:"分类图片" json:"categoryImg" default:""`           // 分类图片地址链接
	CategoryID   uint   `validate:"-" label:"所属分类" json:"categoryID" default:"0"`           // 所属分类
	Sort         uint   `validate:"min=0,max=100" label:"排序" json:"sort" default:"0"`       //排序
}

// Post 添加商品分类
// @Summary      添加商品分类
// @Description  后台管理人员添加商品分类
// @Accept       json
// @Produce      json
// @param        root  body  CategoryRequest  true  "添加商品分类"
// @Tags         商品分类
// @Success      200  {object}  app.Response{data=render.Category}
// @Router       /admin/category [post]
func (this *CategoryController) Post() *app.Response {
	param := &CategoryRequest{}
	if err := app.FormValueJson(this.Ctx, global.Validate, param); err != nil {
		return app.ToResponseErr(err)
	}

	category := &models.Category{
		CategoryName: param.CategoryName,
		CategoryImg:  param.CategoryImg,
		Sort:         param.Sort,
		CategoryID:   param.CategoryID,
	}

	if err := services.CategoryService.Create(category); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseData(render.BuildCreategory(category))
}

// Get 根据分类id获取子集分类
// @Summary      获取子分类列表
// @Description  根据分类id获取子分类列表
// @Produce      json
// @param        categoryID  query  uint  false  "分类id"  default(0)
// @param        page        query  uint  false  "分页"    default(1)
// @param        pageSize    query  uint  false  "分页页数"  default(10)
// @Tags         商品分类
// @Success      200  {object}  app.Response{data=[]render.Category}
// @Router       /admin/category [get]
func (this *CategoryController) Get() *app.Response {
	categoryID := app.FormValueUintDefault(this.Ctx, "categoryID", 0)
	page := app.GetPage(this.Ctx)
	pageSize := app.GetPageSize(this.Ctx)
	list, total := services.CategoryService.GetListPage(global.DB.Where("category_id", categoryID), page, pageSize)

	return app.ToResponseList(render.BuildCreategories(list), total)
}

// GetBy 通过id获取分类信息
// @Summary      获取分类详情
// @Description  通过分类id获取分类详情
// @Produce      json
// @param        categoryID  path  uint  true  "分类id"
// @tags         商品分类
// @Success      200  {object}  app.Response{data=render.Category}
// @Router       /admin/category/{categoryID} [get]
func (this *CategoryController) GetBy(id uint) *app.Response {
	res := services.CategoryService.Get(id)

	if res == nil {
		return app.ToResponseErr(errcode.NotFound)
	}
	return app.ResponseData(render.BuildCreategory(res))
}

// PutBy 修改分类信息
// @Summary      修改分类信息
// @Description  修改分类信息
// @Accept       json
// @Produce      json
// @param        categoryID  path  uint             true  "分类ID"
// @param        root        body  CategoryRequest  true  "修改商品分类"
// @Tags         商品分类
// @Success      200  {object}  app.Response{data=render.Category}
// @Router       /admin/category/{categoryID} [put]
func (this *CategoryController) PutBy(id uint) *app.Response {
	category := services.CategoryService.Get(id)
	if category == nil {
		return app.ToResponseErr(errcode.NotFound)
	}

	param := &CategoryRequest{}
	if err := app.FormValueJson(this.Ctx, global.Validate, param); err != nil {
		return app.ToResponseErr(err)
	}

	category.CategoryName = param.CategoryName
	category.CategoryImg = param.CategoryImg
	category.CategoryID = param.CategoryID
	category.Sort = param.Sort

	if category.ID == category.CategoryID {
		return app.ResponseErrMsg("自己不能为自己父级")
	}

	if err := services.CategoryService.Update(category); err != nil {
		return app.ToResponseErr(err)
	}

	return app.ResponseMsg("param")
}

// DeleteBy 删除分类信息
// @Summary      删除分类信息
// @Description  删除分类信息
// @Produce      json
// @param        categoryID  path  uint  true  "分类ID"
// @Tags         商品分类
// @Success      200  {object}  app.Response{data=render.Category}
// @Router       /admin/category/{categoryID} [delete]
func (this *CategoryController) DeleteBy(id uint) *app.Response {
	if err := services.CategoryService.Delete(id); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseMsg("")
}

type CategoryQueryName struct {
	CategoryName string `validate:"required,min=1,max=20" label:"分类名称" json:"categoryName"` // 分类名称
}

// PostQueryName 通过分类名称查询
// @Summary      分类名称搜索
// @Description  通过分类名称搜索可绑定的分类信息
// @Produce      json
// @Param        root  body  CategoryQueryName  true  "分类名称"
// @Tags         商品分类
// @Success      200  {object}  app.Response{data=[]render.Category}
// @Router       /admin/category/query/name [post]
func (this *CategoryController) PostQueryName() *app.Response {
	param := &CategoryQueryName{}
	if err := app.FormValueJson(this.Ctx, global.Validate, param); err != nil {
		return app.ToResponseErr(err)
	}

	categories, err := services.CategoryService.QueryName(param.CategoryName)
	if err != nil {
		return app.ToResponseErr(err)
	}

	return app.ResponseData(render.BuildCreategoryPathName(categories))
}
