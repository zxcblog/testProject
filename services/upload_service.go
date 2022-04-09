package services

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"math"
	"mime/multipart"
	"new-project/cache"
	"new-project/global"
	"new-project/models"
	"new-project/models/form"
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

// Upload 上传文件
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

// InitialMultipart 初始化分片上传元信息
func (*uploadService) InitialMultipart(fileSize uint, fileName, fileType string) (string, int, error) {
	fileExt := fileName[strings.LastIndex(fileName, ".")+1:] // 通过文件名获取文件后缀
	newfileName := util.RandomStr(32) + "." + fileExt

	model := &models.Upload{
		FileSize: fileSize,
		OldName:  fileName,
		NewName:  newfileName,
		FileType: fileType,
		FileExt:  fileExt,
	}

	// 根据大小设置文件分块数量
	chunkNum := int(math.Ceil(float64(fileSize) / float64(config.GetService().GetChunkSize())))
	imur, err := global.Upload.InitiateMultipart(newfileName)
	if err != nil {
		global.Logger.Error("文件元信息分块初始化失败", zap.Error(err))
		return "", 0, errors.New("文件上传失败")
	}

	cache.UploadCache.InitiateMultipart(model, imur, chunkNum)

	return imur.UploadID, chunkNum, nil
}

// UploadPart 上传分块文件
func (*uploadService) UploadPart(part *form.UploadPart) error {
	imur := cache.UploadCache.GetImur(part.UploadId)
	content := make([]byte, part.FileSize)
	_, err := part.File.Read(content)
	if err != nil {
		global.Logger.Error("分片上传文件内容读取失败", zap.Error(err))
		return errors.New("上传失败")
	}

	res, err := global.Upload.Bucket.UploadPart(*imur, bytes.NewReader(content), part.FileSize, part.Num)
	if err != nil {
		global.Logger.Error("分片上传传送失败", zap.Error(err))
		return errors.New("上传失败")
	}

	cache.UploadCache.SetUploadParts(part.UploadId, res)
	return nil
}

// Complete 上传完成合并
func (*uploadService) Complete(uploadId string, userId uint) error {
	parts := cache.UploadCache.GetUploadParts(uploadId)
	imur := cache.UploadCache.GetImur(uploadId)

	// 通知上传合并
	res, err := global.Upload.Bucket.CompleteMultipartUpload(*imur, parts)
	if err != nil {
		global.Logger.Error("文件上传合并失败", zap.Error(err))
		return errors.New("文件保存失败")
	}

	// 将数据信息同步到mysql
	upload := cache.UploadCache.GetUpload(uploadId)
	if upload != nil {
		upload.UserId = userId
		upload.SavePath = res.Location
		if err := repositories.UploadRepositories.Create(global.DB, upload); err != nil {
			global.Logger.Error("文件上传合并保存到数据库失败", zap.Any("upload", upload), zap.Error(err))
			return errors.New("文件保存失败")
		}
	}
	return nil
}
