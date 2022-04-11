// Package errcode
// @Author:        asus
// @Description:   $
// @File:          error
// @Data:          2022/4/1117:39
//
package app

import (
	"fmt"
)

var (
	Success                   = NewResponseError(0, "成功")
	ServerError               = NewResponseError(1000, "服务内部错误")
	InvalidParams             = NewResponseError(1001, "入参错误")
	NotFound                  = NewResponseError(1002, "找不到")
	UnauthorizedAuthNotExist  = NewResponseError(1003, "请登录后重试")
	UnauthorizedTokenError    = NewResponseError(1004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewResponseError(1005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewResponseError(1006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewResponseError(1007, "请求过多")
	UploadFileError           = NewResponseError(1008, "文件上传失败")
	RequestError              = NewResponseError(1009, "普通请求失败")

	// 数据库错误使用2开头
	SelectError      = NewResponseError(2001, "查找失败")
	CreateError      = NewResponseError(2002, "创建失败")
	UpdateError      = NewResponseError(2003, "修改失败")
	DelError         = NewResponseError(2004, "删除失败")
	TransactionError = NewResponseError(2005, "事务操作失败")
)

var codes = map[int]string{}

func NewResponseError(code int, msg string) *Response {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已存在，请更换新的错误码", code))
	}

	return &Response{Code: code, Msg: msg}
}

func ResponseErrMsg(msg string) *Response {
	return &Response{Code: RequestError.Code, Msg: msg}
}
