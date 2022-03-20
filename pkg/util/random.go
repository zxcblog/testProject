// Package service
// @Author:        asus
// @Description:   $
// @File:          random
// @Data:          2021/12/2318:44
//
package util

import (
	"math/rand"
	"time"
)

var (
	Chars  = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	Number = []byte("0123456789")
)

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

//随机生成字符串
func RandStr(length int, letter []byte) string {
	b := make([]byte, length)
	randomCharsLen := len(letter)

	for i := range b {
		b[i] = letter[rander.Intn(randomCharsLen)]
	}
	return string(b)
}

//RandomStr 获取指定长度的字符串
func RandomStr(length int) string {
	return RandStr(length, Chars)
}

//RandomNumber 获取指定长度的随机数值字符串
func RandomNumber(length int) string {
	return RandStr(length, Number)
}

// RandomInt 返回0-length之间的数值
func RandomInt(length int) int {
	return rand.Intn(length)
}

// RandomInt64 返回0-length区间的数值
func RandomInt64(length int64) int64 {
	return rand.Int63n(length)
}
