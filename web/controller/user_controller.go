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

// CreateByEmail 创建用户
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
func (a *UserController) Login(ctx *gin.Context) {
	var logModel model.UserLoginModel
	if ok, err := xgin.ShouldBindJSON(ctx, &logModel); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.UserService.LoginByEmail(ctx, logModel).Response(ctx)
}

// AssignRole 给用户分配角色
func (a *UserController) AssignRole(ctx *gin.Context) {
	var userRole model.AssignRoleModel
	if ok, err := xgin.ShouldBindJSON(ctx, &userRole); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	a.UserService.AssignRole(ctx, userRole).Response(ctx)
}

// DeleteById 刪除用戶
func (a *UserController) DeleteById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		result.Failure[any]("参数错误").Response(ctx)
		return
	}
	a.UserService.DeleteById(ctx, id).Response(ctx)
}
