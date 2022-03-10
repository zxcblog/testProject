package comm

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"new-project/pkg/app"
	"new-project/pkg/errcode"
)

type UploadController struct {
	Ctx iris.Context
}

// Post 单个文件上传
// @Summary 单个文件上传
// @Description 单个文件上传
// @Accept mpfd
// @Produce json
// @param file formData file true "文件"
// @Tags 文件上传
// @Success 200 {object} app.Response{}
// @Router /comm/upload [post]
func (this *UploadController) Post() *app.Response {
	file, fileHeader, err := this.Ctx.FormFile("file")
	if err != nil {
		return app.ToResponseErr(errcode.UploadFileError.SetMsg(err.Error()))
	}

	fmt.Printf("打印file信息：%+v\n\n", file)
	fmt.Println(fileHeader.Filename, fileHeader.Size)
	return app.ResponseMsg("")
}
