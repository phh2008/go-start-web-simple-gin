package controller

import (
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/xgin"
	"com.gientech/selection/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var PermissionSet = wire.NewSet(wire.Struct(new(PermissionController), "*"))

type PermissionController struct {
	PermissionService *service.PermissionService
}

// AddPermission 添加权限
func (a *PermissionController) AddPermission(ctx *gin.Context) {
	var perm model.PermissionModel
	if ok, err := xgin.ShouldBindJSON(ctx, &perm); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.PermissionService.Add(ctx, perm).Response(ctx)
}
