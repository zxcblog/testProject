package cache

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"new-project/models"
)

var UploadCache = NewUploadCache()

type uploadCache struct{}

func NewUploadCache() *uploadCache {
	return &uploadCache{}
}

const (
	uploadPrefix = "upload"
)

// InitiateMultipart 文件分片上传信息存储
func (*uploadCache) InitiateMultipart(upload *models.Upload, imur oss.InitiateMultipartUploadResult, chunkNum int) {
	// 将文件信息， 分片信息， 分片数量进行保存

	//imur.UploadID

	//global.Redis.S
}
