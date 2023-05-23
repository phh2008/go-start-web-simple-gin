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
//
//	@Summary		角色管理列表
//	@Description	角色管理列表
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			role	query		model.RoleListReq	false	"查询条件"
//	@Success		200		{object}	result.Result[model.PageData[model.RoleModel]]
//	@Router			/role/list [get]
func (a *RoleController) List(ctx *gin.Context) {
	var req model.RoleListReq
	if ok, err := xgin.ShouldBindQuery(ctx, &req); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.RoleService.ListPage(ctx, req).Response(ctx)
}

// Add 添加角色
//
//	@Summary		添加角色
//	@Description	添加角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			role	body		model.RoleModel	false	"角色信息"
//	@Success		200		{object}	result.Result[entity.RoleEntity]
//	@Router			/role/add [post]
func (a *RoleController) Add(ctx *gin.Context) {
	var roleModel model.RoleModel
	if ok, err := xgin.ShouldBindJSON(ctx, &roleModel); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.RoleService.Add(ctx, roleModel).Response(ctx)
}

// Delete 删除角色
//
//	@Summary		删除角色
//	@Description	删除角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path		int	true	"角色ID"
//	@Success		200	{object}	result.Result[any]
//	@Router			/role/delete/:id [delete]
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
//
//	@Summary		分配权限
//	@Description	分配权限
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			role	body		model.RoleAssignPermModel	true	"角色与权限信息"
//	@Success		200		{object}	result.Result[any]
//	@Router			/role/assignPermission [post]
func (a *RoleController) AssignPermission(ctx *gin.Context) {
	var assignModel model.RoleAssignPermModel
	if ok, err := xgin.ShouldBindJSON(ctx, &assignModel); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.RoleService.AssignPermission(ctx, assignModel).Response(ctx)
}
