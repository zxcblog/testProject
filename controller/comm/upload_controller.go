package comm

import (
	"github.com/kataras/iris/v12"
	"new-project/controller/render"
	"new-project/global"
	"new-project/models"
	"new-project/models/form"
	"new-project/pkg/app"
	"new-project/pkg/config"
	"new-project/pkg/errcode"
	"new-project/services"
)

type UploadController struct {
	Ctx iris.Context
}

// Post 单个文件上传
// @Summary      单个文件上传
// @Description  单个文件上传
// @Security ApiKeyAuth
// @Accept       mpfd
// @Produce      json
// @param        file  formData  file  true  "文件"
// @Tags         文件上传
// @Success      200  {object}  app.Response{data=render.Upload}
// @Router       /comm/upload [post]
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

type InitiateMultipart struct {
	FileSize uint   `validate:"required,min=102400,max=536870912000" label:"文件大小" json:"fileSize"` // 文件大小 最小不能小于100KB 最大不能超过500GB
	FileName string `validate:"required" label:"文件名" json:"fileName"`                              // 文件名
	FileType string `validate:"required" label:"文件类型" json:"FileType"`                             // 文件类型
	//MD5Check string `validate:"required" label:"md5校验码" json:"MD5Check"`
}

// PostInitiateMultipart 大文件分块上传元信息
// @Summary      大文件分块上传元信息
// @Description  大文件分块上传元信息
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @param        root  body  InitiateMultipart  true  "文件上传元信息"
// @Tags         文件上传
// @Success      200  {object}  app.Response{data=app.Result{uploadId=string,chunkNum=int,chunkSize=int}} "uploadId上传分块文件标识 chunkNum文件上传数量 chunkSize每块文件切割大小"
// @Router       /comm/upload/initiate/multipart [post]
func (this *UploadController) PostInitiateMultipart() *app.Response {
	param := &InitiateMultipart{}
	if err := app.FormValueJson(this.Ctx, global.Validate, param); err != nil {
		return app.ToResponseErr(err)
	}

	// 对文件进行重命名
	uploadId, chunkNum, err := services.UploadService.InitialMultipart(param.FileSize, param.FileName, param.FileType)
	if err != nil {
		return app.ToResponseErr(err)
	}

	return app.ResponseData(app.Result{
		"uploadId":  uploadId,
		"chunkNum":  chunkNum,
		"chunkSize": config.GetService().GetChunkSize(),
	})
}

// PostPartBy 上传文件分块信息
// @Summary      上传文件分块信息
// @Description  上传文件分块信息
// @Accept       mpfd
// @Produce      json
// @Security ApiKeyAuth
// @param        uploadId  path      string  true  "文件uploadId"
// @param        num       path      int     true  "上传片"
// @param        file      formData  file    true  "文件"
// @Tags         文件上传
// @Success      200  {object}  app.Response{data=render.Upload}
// @Router       /comm/upload/part/{U} [post]
func (this *UploadController) PostPartBy(uploadId string, num int) *app.Response {
	file, fileHeader, err := this.Ctx.FormFile("file")
	if err != nil {
		return app.ToResponseErr(errcode.UploadFileError.SetMsg(err.Error()))
	}
	defer file.Close()

	req := &form.UploadPart{
		FileSize: fileHeader.Size,
		Num:      num,
		UploadId: uploadId,
		File:     file,
	}
	if err := services.UploadService.UploadPart(req); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseMsg("上传成功")
}

// PostCompleteBy 通知上传完成合并操作
// @Summary      通知上传完成合并操作
// @Description  通知上传完成合并操作
// @Accept       mpfd
// @Produce      json
// @Security ApiKeyAuth
// @param        uploadId  path  string  true  "文件uploadId"
// @Tags         文件上传
// @Success      200  {object}  app.Response{data=render.Upload}
// @Router       /comm/upload/complete/{uploadId} [post]
func (this *UploadController) PostCompleteBy(uploadId string) *app.Response {
	// TODO 登录用户id
	user := &models.User{}
	user.ID = 6
	if err := services.UploadService.Complete(uploadId, user.ID); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseMsg("上传成功")
}
