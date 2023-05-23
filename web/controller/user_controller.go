package controller

import (
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/xgin"
	"com.gientech/selection/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
)

var UserSet = wire.NewSet(wire.Struct(new(UserController), "*"))

type UserController struct {
	UserService *service.UserService
}

// List 用户管理列表
//
//	@Summary		用户管理列表
//	@Description	用户管理列表
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			user	query		model.UserListReq	false	"查询条件"
//	@Success		200		{object}	result.Result[model.PageData[model.UserModel]]
//	@Router			/user/list [get]
func (a *UserController) List(ctx *gin.Context) {
	var req model.UserListReq
	if ok, err := xgin.ShouldBindQuery(ctx, &req); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.UserService.ListPage(ctx, req).Response(ctx)
}

// CreateByEmail 创建用户
//
//	@Summary		邮箱注册用户
//	@Description	邮箱注册用户
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.UserEmailRegister	true	"邮箱与密码"
//	@Success		200		{object}	result.Result[model.UserModel]
//	@Router			/user/createByEmail [post]
func (a *UserController) CreateByEmail(ctx *gin.Context) {
	var email model.UserEmailRegister
	if ok, err := xgin.ShouldBindJSON(ctx, &email); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	res := a.UserService.CreateByEmail(ctx, email)
	res.Response(ctx)
}

// Login 邮箱登录
//
//	@Summary		使用邮箱登录
//	@Description	使用邮箱登录
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.UserLoginModel	true	"登录信息"
//	@Success		200		{object}	result.Result[any]
//	@Router			/user/login [post]
func (a *UserController) Login(ctx *gin.Context) {
	var logModel model.UserLoginModel
	if ok, err := xgin.ShouldBindJSON(ctx, &logModel); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.UserService.LoginByEmail(ctx, logModel).Response(ctx)
}

// AssignRole 给用户分配角色
//
//	@Summary		给用户分配角色
//	@Description	给用户分配角色
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			query	body		model.AssignRoleModel	true	"用户ID与角色编号"
//	@Success		200		{object}	result.Result[any]
//	@Router			/user/assignRole [post]
func (a *UserController) AssignRole(ctx *gin.Context) {
	var userRole model.AssignRoleModel
	if ok, err := xgin.ShouldBindJSON(ctx, &userRole); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.UserService.AssignRole(ctx, userRole).Response(ctx)
}

// DeleteById 刪除用戶
//
//	@Summary		刪除用戶
//	@Description	刪除用戶
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path		int	true	"用户ID"
//	@Success		200	{object}	result.Result[any]
//	@Router			/user/delete/:id [delete]
func (a *UserController) DeleteById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		result.Failure[any]("参数错误").Response(ctx)
		return
	}
	a.UserService.DeleteById(ctx, id).Response(ctx)
}
