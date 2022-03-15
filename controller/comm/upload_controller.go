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
