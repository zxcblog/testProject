package admin

import (
	"encoding/json"
	"new-project/global"
	"new-project/models"
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
// @Summary 添加商品
// @Description 后台管理人员添加商品
// @Accept json
// @Produce json
// @param root body PostProductAddCheckedParams true "添加商品"
// @Tags 商品
// @Success 200 {object} app.Response{}
// @Router /admin/product/productadd [post]
func (this *ProductController) PostProductadd() *app.Response {
	params := &PostProductAddCheckedParams{}

	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return app.ToResponseErr(err)
	}

	jsonPicImgIds, err := json.Marshal(params.PicImgIds)
	//Marshal失败时err!=nil
	if err != nil {
		return app.ToResponseErr(err)
	}

	//商品model赋值
	product := &models.Product{
		ProductName:   params.ProductName,
		ProductTitle:  params.ProductTitle,
		CategoryID:    params.CategoryID,
		BrandId:       params.BrandId,
		CategoryImgId: params.CategoryImgId,
		PicImgIds:     string(jsonPicImgIds),
		Introduction:  params.Introduction,
	}

	//添加商品主信息
	productErr := services.ProductService.Create(product)

	if productErr != nil {
		return app.ToResponseErr(productErr)
	}

	//插入sku key与value
	for _, val := range params.SkuAttributeKeyValueData {
		productSkuKey := &models.ProductSkuKey{
			ProductId:    product.ID,
			BrandId:      params.BrandId,
			AttributeKey: val.AttributeKey,
		}
		//添加规格key
		skuKeyErr := services.ProductSkuKeyService.Create(productSkuKey)
		if skuKeyErr != nil {
			return app.ToResponseErr(skuKeyErr)
		}
		//使用切片批量添加value
		productSliceSkuValue := make([]models.ProductSkuValue, 0)
		for _, val := range val.AttributeValue {
			//组装数据
			productSliceSkuValue = append(productSliceSkuValue, models.ProductSkuValue{
				ProductId:       product.ID,
				ProductSkuKeyId: productSkuKey.ID,
				AttributeValue:  val,
			})
		}
		//批量添加规格value
		skuValueRrr := services.ProductSkuValueService.BatchCreate(&productSliceSkuValue)
		if skuValueRrr != nil {
			return app.ToResponseErr(skuValueRrr)
		}
	}

	SkuSliceData := make([]models.ProductSku, 0)
	//插入组装好的规格数据
	for _, SkuDataVal := range params.SkuData {
		jsonSkuAttribute, err := json.Marshal(SkuDataVal.SkuAttribute)
		//Marshal失败时err!=nil
		if err != nil {
			return app.ToResponseErr(err)
		}
		SkuSliceData = append(SkuSliceData, models.ProductSku{
			ProductId:    product.ID,
			SkuAttribute: string(jsonSkuAttribute),
			Stock:        SkuDataVal.Stock,
			Price:        SkuDataVal.Price,
		})
	}
	skuErr := services.ProductSkuService.BatchCreate(&SkuSliceData)
	if skuErr != nil {
		return app.ToResponseErr(skuErr)
	}

	return app.ResponseMsg("添加成功")
}
