// Package util
// @Author:        asus
// @Description:   $
// @File:          json
// @Data:          2022/3/219:51
//
package util

import "encoding/json"

func StructToString(val interface{}) string {
	src, err := json.Marshal(val)
	if err != nil {
		return ""
	}
	return string(src)
}

func StringToStruct(str string, val interface{}) {
	json.Unmarshal([]byte(str), val)
}
