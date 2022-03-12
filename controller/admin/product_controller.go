package admin

import (
	"encoding/json"
	"fmt"
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
	PicImgIds                string `validata:"required" label:"商品缩略图" json:"picImgIds"`           //商品缩略图
	Introduction             string `label:"商品详情" json:"introduction"`                             //商品详情
	SkuAttributeKeyValueData []struct {
		AttributeKey   string   `validata:"required" json:"attributeKey"`
		AttributeValue []string `validata:"required" json:"attributeValue"`
	} `validata:"required" label:"商品规格" json:"skuAttributeKeyValueData"` //商品规格key和value
	SkuData []struct {
		SkuAttribute map[string]interface{} `validata:"required" json:"skuAttribute"`
		Stock        uint                   `validata:"required" json:"stock"`
	} `validata:"required" label:"商品规格" json:"skuData"` //商品规格
}

//商品添加
func (p *ProductController) PostProductadd() *app.Response {
	params := &PostProductAddCheckedParams{}

	if err := app.FormValueJson(p.Ctx, global.Validate, params); err != nil {
		return app.ToResponseErr(err)
	}

	//商品model赋值
	product := &models.Product{
		ProductName:   params.ProductName,
		ProductTitle:  params.ProductTitle,
		CategoryID:    params.CategoryID,
		BrandId:       params.BrandId,
		CategoryImgId: params.CategoryImgId,
		PicImgIds:     params.PicImgIds,
		Introduction:  params.Introduction,
	}

	//添加商品主信息
	err := services.ProductService.Create(product)

	if err != nil {
		return app.ToResponseErr(err)
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
		productSkuValue := make([]models.ProductSkuValue, 0)
		for _, val := range val.AttributeValue {
			//组装数据
			productSkuValue = append(productSkuValue, models.ProductSkuValue{
				ProductId:       product.ID,
				ProductSkuKeyId: productSkuKey.ID,
				AttributeValue:  val,
			})
		}
		//批量添加规格value
		skuValueRrr := services.ProductSkuValueService.BatchCreate(&productSkuValue)
		if skuValueRrr != nil {
			return app.ToResponseErr(skuValueRrr)
		}
	}

	//插入组装好的规格数据
	for _, SkuDataVal := range params.SkuData {
		jsonSkuAttribute, err := json.Marshal(SkuDataVal.SkuAttribute)
		//Marshal失败时err!=nil
		if err != nil {
			fmt.Println("生成json字符串错误")
		}
		//jsonSkuAttribute是[]byte类型，转化成string类型便于查看
		fmt.Println(string(jsonSkuAttribute))
	}

	fmt.Println(product.ID)

	return app.ResponseMsg("添加成功")
}
