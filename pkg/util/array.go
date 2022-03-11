// Package util
// @Author:        asus
// @Description:   $
// @File:          array
// @Data:          2022/3/1116:23
//
package util

import (
	"reflect"
)

// 查看是否在
func InArray(needle interface{}, haystack interface{}) bool {
	if reflect.TypeOf(haystack).Kind() != reflect.Slice {
		return false
	}

	val := reflect.ValueOf(haystack)
	length := val.Len()
	for i := 0; i < length; i++ {
		if reflect.DeepEqual(needle, val.Index(i).Interface()) {
			return true
		}
	}

	return false
}
