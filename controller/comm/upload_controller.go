package comm

import (
	"github.com/kataras/iris/v12"
	"new-project/controller/render"
	"new-project/global"
	"new-project/pkg/app"
	"new-project/pkg/config"
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

type InitiateMultipart struct {
	FileSize uint   `validate:"required,min=102400,max=536870912000" label:"文件大小" json:"fileSize"` // 文件大小 最小不能小于100KB 最大不能超过500GB
	FileName string `validate:"required" label:"文件名" json:"fileName"`                              // 文件名
	FileType string `validate:"required" label:"文件类型" json:"FileType"`                             // 文件类型
	MD5Check string `validate:"required" label:"md5校验码" json:"MD5Check"`
}

// PostInitiateMultipart 大文件分块上传元信息
// @Summary 大文件分块上传元信息
// @Description 大文件分块上传元信息
// @Accept json
// @Produce json
// @param root body InitiateMultipart true "文件上传元信息"
// @Tags 文件上传
// @Success 200 {object} app.Response
// @Router /comm/upload/initiate/multipart [post]
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

//// PostPartUpload 文件分块信息
//// @Summary 文件分块信息
//// @Description 文件分块信息
//// @Accept mpfd
//// @Produce json
//// @param file formData file true "文件"
//// @Tags 文件上传
//// @Success 200 {object} app.Response{data=render.Upload}
//// @Router /comm/upload [post]
//func (this *UploadController) PostPartUpload() *app.Response {
//	// 本地hash值和服务器hash值进行校验
//
//	// 获得文件句柄， 用于存储分块内容
//
//	// 更新redis缓存状态
//
//	return nil
//}
//
//// PostComplete 通知上传完成合并操作
//// @Summary 通知上传完成合并操作
//// @Description 通知上传完成合并操作
//// @Accept mpfd
//// @Produce json
//// @param file formData file true "文件"
//// @Tags 文件上传
//// @Success 200 {object} app.Response{data=render.Upload}
//// @Router /comm/upload [post]
//func (this *UploadController) PostComplete() *app.Response {
//	// 通过uploadId查询redis并判断是否所有分块上传完成
//
//	// 合并分块
//
//	// 更新唯一文件表
//
//	// 响应处理结果
//	return nil
//}

//func (this *UploadController) GetBy(id uint) {
//upload := services.UploadService.Get(id)
//if upload == nil {
//	fmt.Println("找不到")
//	return
//}
//
//file, err := ioutil.ReadFile(upload.SavePath)
//if err != nil {
//	fmt.Println("文件打开失败")
//	return
//}
//
//if upload.FileType == "image" {
//	this.Ctx.ContentType("image/*")
//} else {
//	this.Ctx.ContentType("video/*")
//}
//this.Ctx.Header("Transfer-Encoding", "chunked")
//
//err = this.Ctx.StreamWriter(func(w io.Writer) error {
//	var i uint = 0
//	for i < upload.FileSize {
//		endi := i + 10000
//		if endi >= upload.FileSize {
//			endi = upload.FileSize
//		}
//		time.Sleep(1 * time.Second)
//		_, err := w.Write(file[i:endi])
//		if err != nil {
//			return err
//		}
//		fmt.Println(i, endi, upload.FileSize)
//		i = endi
//		if i >= upload.FileSize {
//			return io.EOF
//		}
//	}
//	return nil
//})
//if err != nil {
//	fmt.Println("流文件传输失败", err)
//	return
//}
//}
