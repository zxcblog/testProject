package app

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

	return &Response{Code: 0, Data: data, Msg: msg}
}

func ToResponseList(list interface{}, totalRows int) *Response {
	return ToResponse("", Result{
		"list":       list,
		"total_rows": totalRows,
	})
}

func ToResponseErr(err error) *Response {
	return &Response{
		//	Code: err,
		//	Data: nil,
		//	Msg:  err.Error(),
	}
}
