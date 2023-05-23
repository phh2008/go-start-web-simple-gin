package controller

import (
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/exception"
	"com.gientech/selection/pkg/xgin"
	"com.gientech/selection/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleController), "*"))

type RoleController struct {
	RoleService *service.RoleService
	UserService *service.UserService
}

// List 角色管理列表
func (a *RoleController) List(ctx *gin.Context) {
	var req model.RoleListReq
	if ok, err := xgin.ShouldBindQuery(ctx, &req); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.RoleService.ListPage(ctx, req).Response(ctx)
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

// Delete 删除角色
func (a *RoleController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		result.Error[any](exception.ParamError)
		return
	}
	a.RoleService.DeleteById(ctx, id).Response(ctx)
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
