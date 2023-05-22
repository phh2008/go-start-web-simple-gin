package exception

import "fmt"

var SysError = NewBizError("500", "系统错误")
var NoLogin = NewBizError("501", "未登录")
var Unauthorized = NewBizError("502", "无权限")

// BizError 业务错误
type BizError struct {
	Code    string
	Message string
}

func NewBizError(code, message string) BizError {
	return BizError{Code: code, Message: message}
}

func (a BizError) Error() string {
	return fmt.Sprintf("code:%s, message:%s", a.Code, a.Message)
}
