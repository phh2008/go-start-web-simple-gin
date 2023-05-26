package middleware

import (
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/exception"
	"com.gientech/selection/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const signKey = "encryptionToken"
const timestampKey = "timestampKey"

const expireTime int64 = 60

// SignVerify 签名验证
func SignVerify(ctx *gin.Context) {
	// 所有请求都要在请求头传递 encryptionToken,timestampKey(秒级)
	// 签名算法：sign=md5(token+timestamp)
	// 比较 encryptionToken与sign是否相等
	encryptionToken := ctx.GetHeader(signKey)
	timestamp := ctx.GetHeader(timestampKey)
	if encryptionToken == "" || timestamp == "" {
		result.Error[any](exception.SignVerifyError).Response(ctx)
		ctx.Abort()
		return
	}
	sign := util.CreateSign(config.GetProfile().Server.SignToken, timestamp)
	// 签名是否一致
	if sign != encryptionToken {
		result.Error[any](exception.SignVerifyError).Response(ctx)
		ctx.Abort()
		return
	}
	// 时间戳时否过期或超前（暂定60秒内有效）
	reqTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		result.Error[any](exception.SignVerifyError).Response(ctx)
		ctx.Abort()
		return
	}
	sub := time.Now().Unix() - reqTimestamp
	if sub > expireTime || -sub > expireTime {
		result.Error[any](exception.SignVerifyError).Response(ctx)
		ctx.Abort()
		return
	}
	ctx.Next()
}
