package controller

import (
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/xgin"
	"com.gientech/selection/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleController), "*"))

type RoleController struct {
	RoleService *service.RoleService
	UserService *service.UserService
}

// Add 添加角色
func (a *RoleController) Add(ctx *gin.Context) {
	var roleModel model.RoleModel
	if ok, err := xgin.ShouldBindJSON(ctx, &roleModel); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.RoleService.Add(ctx, roleModel).Response(ctx)
}

// AssignPermission 分配权限
func (a *RoleController) AssignPermission(ctx *gin.Context) {
	var assignModel model.RoleAssignPermModel
	if ok, err := xgin.ShouldBindJSON(ctx, &assignModel); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.RoleService.AssignPermission(ctx, assignModel).Response(ctx)
}
