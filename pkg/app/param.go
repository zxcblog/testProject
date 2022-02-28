// Package app
// @Author:        asus
// @Description:   $
// @File:          param
// @Data:          2022/2/2817:04
//
package app

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"strconv"
)

func GetPage(ctx iris.Context) int {
	page := FormValueIntDefault(ctx, "page", 1)
	if page < 1 {
		return 1
	}
	return page
}

func GetPageSize(ctx iris.Context) int {
	pageSize := FormValueIntDefault(ctx, "pageSize", 10)
	switch {
	case pageSize <= 1:
		return 10
	case pageSize >= 100:
		return 100
	}
	return pageSize
}

func FormValueUint(ctx iris.Context, name string) (uint, error) {
	res, err := FormValueInt(ctx, name)
	if err != nil {
		return 0, err
	}

	return uint(res), nil
}

func FormValueUintDefault(ctx iris.Context, name string, def uint) uint {
	if v, err := FormValueUint(ctx, name); err == nil {
		return v
	}
	return def
}

func FormValueIntDefault(ctx iris.Context, name string, def int) int {
	if v, err := FormValueInt(ctx, name); err == nil {
		return v
	}
	return def
}

func FormValueInt(ctx iris.Context, name string) (int, error) {
	str := ctx.FormValue(name)
	if str == "" {
		return 0, paramError(name)
	}

	return strconv.Atoi(str)
}

func FormValueUInt64(ctx iris.Context, name string) (uint64, error) {
	str := ctx.FormValue(name)
	if str == "" {
		return 0, paramError(name)
	}
	return strconv.ParseUint(str, 10, 0)
}

// param error
func paramError(name string) error {
	return errors.New(fmt.Sprintf("找不到参数值 '%s'", name))
}
