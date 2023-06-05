package middleware

import (
	"com.gientech/selection/pkg/config"
	"fmt"
	"testing"
	"time"
)

// TestCreateSign 测试生成签名
func TestCreateSign(t *testing.T) {
	config.NewConfig("../../config")
	var data = map[string]interface{}{
		signKey:      "",
		anonKey:      "abc",
		timestampKey: time.Now().Unix(),
		"openId":     "oylxy5HqJ630VVOw1CR3diqUzfuQ",
	}
	fmt.Println(data)
	sign := createSign(data)
	fmt.Println(sign)
}
