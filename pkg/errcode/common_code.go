package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000, "服务内部错误")
	InvalidParams             = NewError(1001, "入参错误")
	NotFound                  = NewError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1003, "请登录后重试")
	UnauthorizedTokenError    = NewError(1004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(1005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(1006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(1007, "请求过多")
	UploadFileError           = NewError(1008, "文件上传失败")
	RequestError              = NewError(1009, "普通请求失败")

	// 数据库错误使用2开头
	SelectError      = NewError(2001, "查找失败")
	CreateError      = NewError(2002, "创建失败")
	UpdateError      = NewError(2003, "修改失败")
	DelError         = NewError(2004, "删除失败")
	TransactionError = NewError(2005, "事务操作失败")
)
