package app

import "new-project/pkg/errcode"

type Result map[string]interface{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func ResponseMsg(msg string) *Response {
	return ToResponse(msg, nil)
}

func ResponseData(data interface{}) *Response {
	return ToResponse("", data)
}

func ToResponse(msg string, data interface{}) *Response {
	if data == nil {
		data = Result{}
	}

	if msg == "" {
		msg = errcode.Success.Msg()
	}

	return &Response{Code: errcode.Success.Code(), Data: data, Msg: msg}
}

func ToResponseList(list interface{}, totalRows int) *Response {
	return ToResponse("", Result{
		"list":       list,
		"total_rows": totalRows,
	})
}

func ResponseErrMsg(msg string) *Response {
	return ToResponseErr(errcode.RequestError.SetMsh(msg))
}

func ToResponseErr(err error) *Response {
	errResponse := &Response{Data: nil, Msg: err.Error()}
	if errCodeErr, ok := err.(*errcode.Error); ok {
		errResponse.Code = errCodeErr.Code()
	}

	return errResponse
}
