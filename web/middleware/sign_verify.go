package middleware

import (
	"bytes"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/exception"
	"com.gientech/selection/pkg/logger"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const signKey = "sign"
const timestampKey = "timestamp"
const anonKey = "anon"

const expireTime int64 = 60

// SignVerify 签名验证
func SignVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 签名验证
		var data map[string]interface{}
		body, err := io.ReadAll(ctx.Request.Body)
		ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
		err = json.Unmarshal(body, &data)
		if err != nil {
			logger.Errorf("参数解析错误：%s", err)
			result.Error[any](exception.ParamError).Response(ctx)
			ctx.Abort()
			return
		}
		if v, ok := data[anonKey]; !ok || cast.ToString(v) == "" {
			result.Error[any](errors.New("缺少参数：" + anonKey)).Response(ctx)
			ctx.Abort()
			return
		}
		if v, ok := data[timestampKey]; !ok || cast.ToString(v) == "" {
			result.Error[any](errors.New("缺少参数：" + timestampKey)).Response(ctx)
			ctx.Abort()
			return
		}
		if _, ok := data[signKey]; !ok {
			result.Error[any](errors.New("缺少参数：" + signKey)).Response(ctx)
			ctx.Abort()
			return
		}
		// 生成签名
		sign := createSign(data)
		// 签名是否一致
		if sign != cast.ToString(data[signKey]) {
			result.Error[any](exception.SignVerifyError).Response(ctx)
			ctx.Abort()
			return
		}
		// 时间戳时否过期或超前（暂定60秒内有效）
		timestamp := cast.ToString(data[timestampKey])
		reqTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
		if err != nil {
			result.Error[any](exception.SignVerifyError).Response(ctx)
			ctx.Abort()
			return
		}
		sub := time.Now().Unix() - reqTimestamp
		expired := getExpired()
		if sub > expired || -sub > expired {
			result.Error[any](exception.SignVerifyError).Response(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// getExpired 获取过期时长
func getExpired() int64 {
	expire := config.GetProfile().Server.ExpireTime
	if expire <= 0 {
		return expireTime
	}
	return expire
}

// createSign 生成签名串
func createSign(data map[string]interface{}) string {
	var keys []string
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		if k == signKey {
			continue
		}
		val := data[k]
		if val != nil {
			valType := reflect.ValueOf(val)
			kind := valType.Type().Kind()
			if kind == reflect.Struct || kind == reflect.Slice || kind == reflect.Map {
				continue
			}
		}
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(cast.ToString(val))
		sb.WriteString("&")
	}
	values := strings.TrimSuffix(sb.String(), "&") + config.GetProfile().Server.SignToken
	sign := md5.Sum([]byte(values))
	return hex.EncodeToString(sign[:])
}
