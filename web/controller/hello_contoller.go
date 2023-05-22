package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

var HelloSet = wire.NewSet(wire.Struct(new(HelloController), "*"))

type HelloController struct {
}

// Hello 测试 api
func (a *HelloController) Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "请求成功：success")
}
