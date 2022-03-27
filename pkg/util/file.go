package util

import (
	"fmt"
	"os"
)

// FileExists 判断文件或文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// MkdirOfNotExists 创建文件夹， 不存在就创建
func MkdirOfNotExists(path string) error {
	if !FileExists(path) {
		return os.MkdirAll(path, 0644)
	}
	return nil
}

// DeleteFile 删除文件
func DeleteFile(path string) bool {
	err := os.Remove(path)
	fmt.Println(err)
	return err == nil
}
