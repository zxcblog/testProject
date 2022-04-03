package comm

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/v12"
	"new-project/cache"
	"new-project/controller/render"
	"new-project/global"
	"new-project/pkg/app"
	"new-project/pkg/config"
	"new-project/pkg/errcode"
	"new-project/services"
	"time"
)

type UploadController struct {
	Ctx iris.Context
}

// Post 单个文件上传
// @Summary      单个文件上传
// @Description  单个文件上传
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
// @param        root  body  InitiateMultipart  true  "文件上传元信息"
// @Tags         文件上传
// @Success      200  {object}  app.Response
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
// @param        uploadId  path      string  true  "文件uploadId"
// @param        num       path      int     true  "上传片"
// @param        file      formData  file    true  "文件"
// @Tags         文件上传
// @Success      200  {object}  app.Response{data=render.Upload}
// @Router       /comm/upload/part/{U} [post]
func (this *UploadController) PostPartBy(uploadId string, num int) *app.Response {
	startTime := time.Now().UnixMilli()
	file, fileHeader, err := this.Ctx.FormFile("file")
	if err != nil {
		return app.ToResponseErr(errcode.UploadFileError.SetMsg(err.Error()))
	}
	defer file.Close()
	fmt.Println("文件获取时间:", time.Now().UnixMilli()-startTime)
	fmt.Println(fileHeader.Size)

	startTime = time.Now().UnixMilli()
	imur := cache.UploadCache.GetImur(uploadId)
	fmt.Println("缓存获取时间:", time.Now().UnixMilli()-startTime)

	startTime = time.Now().UnixMilli()
	content := make([]byte, fileHeader.Size)
	_, err = file.Read(content)
	fmt.Println("文件读取错误：", err)

	res, err := global.Upload.Bucket.UploadPart(*imur, bytes.NewReader(content), fileHeader.Size, num)
	fmt.Println("分片上传时间:", time.Now().UnixMilli()-startTime)
	fmt.Println(res, err)
	if err != nil {
		return app.ToResponseErr(err)
	}

	cache.UploadCache.SetUploadParts(uploadId, res)
	return app.ResponseData(res)
}

// PostCompleteBy 通知上传完成合并操作
// @Summary      通知上传完成合并操作
// @Description  通知上传完成合并操作
// @Accept       mpfd
// @Produce      json
// @param        uploadId  path  string  true  "文件uploadId"
// @Tags         文件上传
// @Success      200  {object}  app.Response{data=render.Upload}
// @Router       /comm/upload/complete/{uploadId} [post]
func (this *UploadController) GetCompleteBy(uploadId string) *app.Response {
	parts := cache.UploadCache.GetUploadParts(uploadId)
	imur := cache.UploadCache.GetImur(uploadId)

	res, err := global.Upload.Bucket.CompleteMultipartUpload(*imur, parts)
	if err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseData(res)
}

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
