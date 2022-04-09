package form

import "mime/multipart"

//Upload 单个文件上传
type Upload struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadPart struct {
	Num      int            // 第几次分片
	FileSize int64          // 文件大小
	UploadId string         // 分片上传唯一标识
	File     multipart.File // 上传文件
}
