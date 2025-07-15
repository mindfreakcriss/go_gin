package util

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	// 这里可以使用实际的MD5实现
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str // 仅为示例，实际应返回MD5哈希值
}
