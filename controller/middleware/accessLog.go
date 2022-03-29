package middleware

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"new-project/global"
	"regexp"
	"strings"
	"time"
)

func AccessLog(ctx iris.Context) {
	startTime := time.Now()

	ctx.Record() // 开启事务接收

	// 获取请求方法和请求路由
	method := ctx.Request().Method
	path := ctx.Request().URL.String()

	reg := regexp.MustCompile("\\s+")

	// 获取参数后将参数放回请求体中
	params := []byte{}

	contentType := ctx.GetHeader("Content-Type")
	if !strings.HasPrefix(contentType, "multipart/form-data") && (method == http.MethodPost || method == http.MethodPut) {
		fmt.Println(contentType, 123456789)
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = ioutil.NopCloser(buf)

			// 字符替换
			params = reg.ReplaceAll(body, []byte(""))
		}
	}

	ctx.Next()
	endTime := time.Now()

	// 记录访问日志
	global.AccessLog.Sugar().Infof(`
%s - - [%s %dμs] "%s %s %s" %d
Request Header Token: %s 
Request Header User-Agent: %s 
Request Body: %s 
Response Body: %s`,
		ctx.RemoteAddr(), startTime.Format("2006-01-02 15:04:05"), endTime.Sub(startTime).Microseconds(), method, path,
		ctx.Request().Proto, ctx.ResponseWriter().StatusCode(),
		ctx.Request().Header.Get("token"), ctx.Request().Header.Get("User-Agent"),
		params, reg.ReplaceAll(ctx.Recorder().Body(), []byte("")),
	)
}
