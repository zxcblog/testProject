package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

//Md5Encrypt md5加密
func Md5Encrypt(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
