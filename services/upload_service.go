package services

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/config"
	"new-project/pkg/errcode"
	"new-project/pkg/util"
	"new-project/repositories"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var UploadService = NewUploadService()

type uploadService struct{}

func NewUploadService() *uploadService {
	return &uploadService{}
}

func (*uploadService) Get(id uint) *models.Upload {
	return repositories.UploadRepositories.Get(global.DB, id)
}

func (*uploadService) Upload(file multipart.File, fileHeader *multipart.FileHeader) (*models.Upload, error) {
	// 通过文件名获取文件后缀
	fileExt := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, ".")+1:]
	fileType := ""
	// 检查当前文件是否允许上传
	switch {
	case util.InArray(fileExt, config.GetService().UploadImageAllowExts):
		uploadImgSize := config.GetService().UploadImgMaxSize
		if util.BigToSmall(uploadImgSize, "m") < fileHeader.Size {
			return nil, errcode.UploadFileError.SetMsg(fmt.Sprintf("图片上传不能超过%fM", uploadImgSize))
		}
		fileType = "image"
	case util.InArray(fileExt, config.GetService().UploadVideoAllowExts):
		uploadVideoSize := config.GetService().UploadVideoMaxSize
		if util.BigToSmall(uploadVideoSize, "m") < fileHeader.Size {
			return nil, errcode.UploadFileError.SetMsg(fmt.Sprintf("视频上传不能超过%fM", uploadVideoSize))
		}
		fileType = "video"
	default:
		return nil, errcode.UploadFileError.SetMsg("文件上传类型不正确")
	}

	// 创建保存路径文件夹
	savePath := filepath.Join(config.GetService().UploadSavePath, fileType, time.Now().Format("20060102"))
	if err := util.MkdirOfNotExists(savePath); err != nil {
		global.Logger.Error("文件上传创建文件路径失败", zap.Error(err))
		return nil, errcode.UploadFileError.SetMsg("文件上传失败")
	}

	// 保存文件
	fileName := util.RandomStr(32) + "." + fileExt
	dest := filepath.Join(savePath, fileName)
	f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		global.Logger.Error("文件上传设置文件路径失败", zap.Error(err))
		return nil, errcode.UploadFileError.SetMsg("文件上传失败")
	}
	defer f.Close()

	if _, err := io.Copy(f, file); err != nil {
		global.Logger.Error("文件保存失败", zap.Error(err))
		return nil, errcode.UploadFileError.SetMsg("文件上传失败")
	}

	upload := &models.Upload{
		FileSize: uint(fileHeader.Size),
		SavePath: dest,
		OldName:  fileHeader.Filename,
		NewName:  fileName,
		FileType: fileType,
		FileExt:  fileExt,
	}
	if err := repositories.UploadRepositories.Create(global.DB, upload); err != nil {
		global.Logger.Error("文件上传存入数据库失败", zap.Error(err))
		util.DeleteFile(upload.SavePath) // 删除文件
		return nil, errcode.CreateError
	}

	return upload, nil
}
