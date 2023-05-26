package util

import (
	"crypto/md5"
	"encoding/hex"
)

// CreateSign 生成签名
func CreateSign(token, timestamp string) string {
	sign := md5.Sum([]byte(token + timestamp))
	return hex.EncodeToString(sign[:])
}
