// Package middleware
// @Author:        asus
// @Description:   $
// @File:          limiter
// @Data:          2022/3/1415:52
//
package middleware

import (
	"github.com/kataras/iris/v12"
	"new-project/pkg/app"
	"new-project/pkg/limiter"
	"time"
)

var limit = make(map[string]*limiter.Bucket)

func Limiter(ctx iris.Context) {
	key := ctx.Path()
	bucket, ok := limit[key]
	if !ok {
		bucket = limiter.NewBucketWithQuantum(time.Second, 10, 10)
		limit[key] = bucket
	}

	count := bucket.TakeAvailable(1)
	if count == 0 {
		ctx.JSON(app.TooManyRequests)
		ctx.StopExecution()
		return
	}

	ctx.Next()
}
