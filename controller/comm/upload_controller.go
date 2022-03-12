package comm

import (
	"github.com/kataras/iris/v12"
	"new-project/controller/render"
	"new-project/pkg/app"
	"new-project/pkg/errcode"
	"new-project/services"
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
// @Success 200 {object} app.Response{data=render.Upload}
// @Router /comm/upload [post]
func (this *UploadController) Post() *app.Response {
	file, fileHeader, err := this.Ctx.FormFile("file")
	if err != nil {
		return app.ToResponseErr(errcode.UploadFileError.SetMsg(err.Error()))
	}
	defer file.Close()

	upload, err := services.UploadService.Upload(file, fileHeader)
	if err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseData(render.BuildUpload(upload))
}
