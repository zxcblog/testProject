package cache

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/util"
	"time"
)

var UploadCache = NewUploadCache()

type uploadCache struct{}

func NewUploadCache() *uploadCache {
	return &uploadCache{}
}

const (
	uploadPrefix = "upload:"
)

// InitiateMultipart 文件分片上传信息存储
func (*uploadCache) InitiateMultipart(upload *models.Upload, imur oss.InitiateMultipartUploadResult, chunkNum int) {
	imurStr := util.StructToString(imur)
	uploadStr := util.StructToString(upload)

	global.Redis.Set(context.Background(), uploadPrefix+imur.UploadID+"imur", imurStr, 24*time.Hour)
	global.Redis.Set(context.Background(), uploadPrefix+imur.UploadID+"upload", uploadStr, 24*time.Hour)
}

func (*uploadCache) GetImur(uploadId string) *oss.InitiateMultipartUploadResult {
	imur := &oss.InitiateMultipartUploadResult{}
	str, err := global.Redis.Get(context.Background(), uploadPrefix+uploadId+"imur").Result()
	if err != nil || str == "" {
		return nil
	}
	util.StringToStruct(str, imur)
	return imur
}

// SetUploadParts 设置分片上传切片
func (this *uploadCache) SetUploadParts(uploadId string, part oss.UploadPart) {
	parts := this.GetUploadParts(uploadId)
	fmt.Println(parts)
	parts = append(parts, part)

	global.Redis.Set(context.Background(), uploadPrefix+uploadId+"parts", util.StructToString(parts), 48*time.Hour)
}

// GetUploadParts 获取分片上传切片
func (*uploadCache) GetUploadParts(uploadId string) (parts []oss.UploadPart) {
	if flag, _ := global.Redis.Exists(context.Background(), uploadPrefix+uploadId+"parts").Result(); flag != 1 {
		return
	}

	str, err := global.Redis.Get(context.Background(), uploadPrefix+uploadId+"parts").Result()
	if err != nil {
		global.Logger.Error("读取上传文件切片列表失败", zap.String("uploadId", uploadId), zap.Error(err))
		return
	}
	util.StringToStruct(str, &parts)
	return
}

// GetUpload 获取上传文件信息
func (*uploadCache) GetUpload(uploadId string) (upload *models.Upload) {
	res, err := global.Redis.Get(context.Background(), uploadPrefix+uploadId+"upload").Result()
	if err != nil {
		global.Logger.Error("读取上传文件元信息失败", zap.String("uploadId", uploadId), zap.Error(err))
		return
	}

	util.StringToStruct(res, &upload)
	return
}
