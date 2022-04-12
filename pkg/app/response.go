package app

type Result map[string]interface{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (r *Response) SetMsg(msg string) *Response {
	err := *r
	err.Msg = msg
	return &err
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
		msg = Success.Msg
	}

	return &Response{Code: Success.Code, Data: data, Msg: msg}
}

func ToResponseList(list interface{}, totalRows int64) *Response {
	return ToResponse("", Result{
		"list":      list,
		"totalRows": totalRows,
	})
}
