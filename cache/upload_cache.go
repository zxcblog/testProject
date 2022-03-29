package cache

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

func (this *uploadCache) SetUploadParts(uploadId string, part oss.UploadPart) {
	parts := this.GetUploadParts(uploadId)
	parts = append(parts, part)

	global.Redis.Set(context.Background(), uploadPrefix+uploadId+"parts", util.StructToString(parts), 48*time.Hour)
}

func (*uploadCache) GetUploadParts(uploadId string) []oss.UploadPart {
	parts := make([]oss.UploadPart, 0)
	if flag, _ := global.Redis.Exists(context.Background(), uploadPrefix+uploadId+"parts").Result(); flag != 1 {
		return parts
	}

	str, _ := global.Redis.Get(context.Background(), uploadPrefix+uploadId+"parts").Result()
	util.StringToStruct(str, parts)
	return parts
}
