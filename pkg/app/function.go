package app

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

//md5密码加盐 到时候建一个全局function这些方法放在那里
func Md5Salt(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

//取得随机字符串:使用字符串拼接
func GetRandString(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

func CacheTimeGenerator(baseMinute, randMinute int64) time.Duration {
	// 先设置随机数 - 这里比较重要
	rand.Seed(time.Now().Unix())
	// 再设置缓存时间
	// baseMinute分钟 + [0~randMinute) 分钟的随机时间
	return time.Duration((baseMinute + rand.Int63n(randMinute)) * int64(time.Minute))
}
