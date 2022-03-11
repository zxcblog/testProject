// Package util
// @Author:        asus
// @Description:   $
// @File:          byte_conv
// @Data:          2022/3/1116:58
//
package util

import (
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

// BigToSmall 将字节转换为对应大小的数值
func BigToSmall(size float64, unit string) int64 {

	switch strings.ToUpper(unit) {
	case "B":
		size = size * B
	case "KB", "K":
		size = size * KB
	case "MB", "M":
		size = size * MB
	case "GB", "G":
		size = size * GB
	case "TB", "T":
		size = size * TB
	case "PB", "P":
		size = size * PB
	}

	return int64(size)
}
