package admin

import (
	"new-project/controller/render"
	"new-project/global"
	"new-project/pkg/app"
	"new-project/services"

	"github.com/kataras/iris/v12"
)

type ProductController struct {
	Ctx iris.Context
}

//添加商品校验
type PostProductAddCheckedParams struct {
	ProductName              string `validata:"required,max=100" label:"商品名称" json:"productName"`  //商品名称
	ProductTitle             string `validata:"required,max=100" label:"商品标题" json:"productTitle"` //商品标题
	CategoryID               uint   `validata:"required" label:"商品分类" json:"categoryID"`           //商品分类
	BrandId                  uint   `validata:"required" label:"品牌分类" json:"brandId"`              //品牌分类
	CategoryImgId            uint   `validata:"required" label:"商品主图ID" json:"categoryImgId"`      //商品主图ID
	PicImgIds                []uint `validata:"required" label:"商品缩略图" json:"picImgIds"`           //商品缩略图
	Introduction             string `label:"商品详情" json:"introduction"`                             //商品详情
	SkuAttributeKeyValueData []struct {
		AttributeKey   string   `validata:"required" json:"attributeKey"`
		AttributeValue []string `validata:"required" json:"attributeValue"`
	} `validata:"required" label:"商品规格" json:"skuAttributeKeyValueData"` //商品规格key和value
	SkuData []struct {
		SkuAttribute map[string]interface{} `validata:"required" json:"skuAttribute"`
		Stock        uint                   `validata:"required" json:"stock"`
		Price        uint                   `validata:"required" json:"price"`
	} `validata:"required" label:"商品规格" json:"skuData"` //商品规格
}

// Post 添加商品
// @Summary      添加商品
// @Description  后台管理人员添加商品
// @Accept       json
// @Produce      json
// @param        root  body  PostProductAddCheckedParams  true  "添加商品"
// @Tags         商品
// @Success      200  {object}  app.Response{}
// @Router       /admin/product [post]
func (this *ProductController) Post() *app.Response {
	params := &PostProductAddCheckedParams{}
	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return err
	}

	////TODO 增加事务
	//
	////商品model赋值
	//product := &models.Product{
	//	ProductName:   params.ProductName,
	//	ProductTitle:  params.ProductTitle,
	//	CategoryID:    params.CategoryID,
	//	BrandId:       params.BrandId,
	//	CategoryImgId: params.CategoryImgId,
	//	PicImgIds:     util.StructToString(params.PicImgIds),
	//	Introduction:  params.Introduction,
	//}
	//
	////添加商品主信息
	//err := services.ProductService.Create(product)
	//
	//if err != nil {
	//	return app.ToResponseErr(err)
	//}
	//
	////插入sku key与value
	//SkuKeyValueSliceData := make([]*models.ProductSkuKeyValue, 0)
	//for _, val := range params.SkuAttributeKeyValueData {
	//
	//	SkuKeyValueSliceData = append(SkuKeyValueSliceData, &models.ProductSkuKeyValue{
	//		ProductId:      product.ID,
	//		BrandId:        params.BrandId,
	//		AttributeKey:   val.AttributeKey,
	//		AttributeValue: util.StructToString(val.AttributeValue),
	//	})
	//}
	////批量添加规格key和value
	//err = services.ProductSkuKeyValueService.BatchCreate(SkuKeyValueSliceData)
	//if err != nil {
	//	return app.ToResponseErr(err)
	//}
	//
	//SkuSliceData := make([]*models.ProductSku, 0)
	////插入组装好的规格数据
	//for _, SkuDataVal := range params.SkuData {
	//
	//	SkuSliceData = append(SkuSliceData, &models.ProductSku{
	//		ProductId:    product.ID,
	//		SkuAttribute: util.StructToString(SkuDataVal.SkuAttribute),
	//		Stock:        SkuDataVal.Stock,
	//		Price:        SkuDataVal.Price,
	//	})
	//}
	//err = services.ProductSkuService.BatchCreate(SkuSliceData)
	//if err != nil {
	//	return app.ToResponseErr(err)
	//}

	return app.ResponseMsg("添加成功")
}

// Post 获取商品详情
// @Summary      获取商品详情
// @Description  后台管理人员获取商品详情
// @Produce      json
// @param        productId  path  uint  true  "商品id"
// @Tags         商品
// @Success      200  {object}  app.Response{data=app.Result}
// @Router       /admin/product/{productId} [get]
func (this *ProductController) GetBy(id uint) *app.Response {
	productRes := services.ProductService.Get(id)
	if productRes == nil {
		return app.NotFound
	}

	//sku数据
	productSkuRes := services.ProductSkuService.GetProductIdAllData(productRes.ID)

	//sku key和value值
	productSkuValueRes := services.ProductSkuKeyValueService.GetProductIdAllData(productRes.ID)

	return app.ToResponse("获取成功", app.Result{
		"product":                  render.BuildProduct(productRes),
		"skuAttributeKeyValueData": render.BuildProductSkuKeyValueList(productSkuValueRes),
		"skuData":                  render.BuildProductSkuList(productSkuRes),
	})

}

// Get 获取商品列表
// @Summary      获取商品列表
// @Description  获取商品列表
// @Produce      json
// @param        productName  query  string  false  "商品名称"  default("")
// @param        brandId      query  uint    false  "品牌id"  default(0)
// @param        categoryID   query  uint    false  "分类id"  default(0)
// @param        page         query  uint    false  "分页"    default(1)
// @param        pageSize     query  uint    false  "分页页数"  default(10)
// @Tags         商品
// @Success      200  {object}  app.Response{data=[]render.Product}
// @Router       /admin/product [get]
func (this *ProductController) Get() *app.Response {
	//组装where
	whereParams := make(map[string]string, 0)
	whereParams["productName"] = this.Ctx.FormValue("productName")
	whereParams["categoryID"] = this.Ctx.FormValue("categoryID")
	whereParams["brandId"] = this.Ctx.FormValue("brandId")

	page := app.GetPage(this.Ctx)
	pageSize := app.GetPageSize(this.Ctx)

	list, total := services.ProductService.GetListPage(whereParams, page, pageSize)
	return app.ResponseList(render.BuildProductList(list), total)
}
