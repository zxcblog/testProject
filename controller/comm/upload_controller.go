package comm

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"new-project/global"
	"new-project/pkg/app"
	"new-project/pkg/config"
	"new-project/pkg/errcode"
	"new-project/pkg/util"
	"path/filepath"
	"strings"
	"time"
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
	_, fileHeader, err := this.Ctx.FormFile("file")
	if err != nil {
		return app.ToResponseErr(errcode.UploadFileError.SetMsg(err.Error()))
	}

	// 通过文件名获取文件后缀
	fileExt := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, ".")+1:]
	switch {
	case util.InArray(fileExt, config.GetService().UploadImageAllowExts):
		uploadImgSize := config.GetService().UploadImgMaxSize
		if util.BigToSmall(uploadImgSize, "m") < fileHeader.Size {
			return app.ToResponseErr(errcode.UploadFileError.SetMsg(fmt.Sprintf("图片上传不能超过%fM", uploadImgSize)))
		}

	case util.InArray(fileExt, config.GetService().UploadVideoAllowExts):
		uploadVideoSize := config.GetService().UploadVideoMaxSize
		if util.BigToSmall(uploadVideoSize, "m") < fileHeader.Size {
			return app.ToResponseErr(errcode.UploadFileError.SetMsg(fmt.Sprintf("视频上传不能超过%fM", uploadVideoSize)))
		}
	default:
		return app.ToResponseErr(errcode.UploadFileError.SetMsg("文件上传类型不正确"))
	}

	// 创建文件夹

	dest := filepath.Join(config.GetService().UploadSavePath, time.Now().Format("20060102"), util.RandomStr(32)+"."+fileExt)
	_, err = this.Ctx.SaveFormFile(fileHeader, dest)
	if err != nil {
		global.Logger.Error("文件上传失败", zap.Error(err))
		return app.ToResponseErr(errcode.UploadFileError.SetMsg("文件上传失败"))
	}

	//
	//upload := &models.Upload{
	//	FileSize: uint(fileHeader.Size),
	//	SavePath: "",
	//	OldName:  "",
	//	NewName:  "",
	//	FileType: "",
	//}
	//
	return app.ResponseMsg("上传成功")
}
