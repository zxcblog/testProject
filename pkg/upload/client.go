package upload

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"new-project/pkg/config"
	"os"
	"time"
)

type Client struct {
	oss    *oss.Client
	Bucket *oss.Bucket
}

func NewClient(option *config.Oss) (*Client, error) {
	ossClient, err := oss.New(option.Endpoint, option.AccessKeyID, option.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	if option.BucketName == "" {
		return nil, errors.New("BucketName不能为空，请手动创建Bucket")
	}

	bucket, err := ossClient.Bucket(option.BucketName)
	if err != nil {
		return nil, err
	}
	return &Client{oss: ossClient, Bucket: bucket}, nil
}

// FilePathUpload 通过文件路径上传文件
// fileName 上传后的文件名称
func (client *Client) FilePathUpload(fileName, filePath string) error {
	fd, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fd.Close()

	return client.FileUpload(fileName, fd)
}

// FileUpload 通过文件流上传文件
// fileName 上传后的文件名称
func (client *Client) FileUpload(fileName string, fd io.Reader) error {
	return client.Bucket.PutObject(fileName, fd)
}

// DeleteFile 通过文件名删除文件信息
func (client *Client) DeleteFile(fileName string) error {
	return client.Bucket.DeleteObject(fileName)
}

// InitiateMultipart 初始化分片上传事件
func (client *Client) InitiateMultipart(fileName string) (oss.InitiateMultipartUploadResult, error) {
	dayTime, _ := time.ParseDuration("24h")

	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		// 设置过期时间， 到期自动删除
		oss.Expires(time.Now().Add(dayTime)),
		// 指定该Object被下载时的网页缓存行为。
		oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		oss.ContentDisposition("attachment;filename=" + fileName),
		// 指定该Object的内容编码格式。
		oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		oss.EncodingType("url"),
	}

	return client.Bucket.InitiateMultipartUpload(fileName, options...)
}
