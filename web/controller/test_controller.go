package controller

import (
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/xgin"
	"com.gientech/selection/pkg/xjwt"
	"com.gientech/selection/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var TestSet = wire.NewSet(wire.Struct(new(TestController), "*"))

type TestController struct {
	TestService *service.TestService
	Jwt         *xjwt.JwtHelper
	Enforcer    *casbin.Enforcer
}

// GetToken 生成 token
func (a *TestController) GetToken(ctx *gin.Context) {
	userId := ctx.Query("userId")
	role := ctx.Query("role")
	user := xjwt.UserClaims{}
	user.ID = userId
	user.Role = role
	token, _ := a.Jwt.CreateToken(user)
	result.Ok[string](token.String()).Response(ctx)
}

func (a *TestController) Auth(ctx *gin.Context) {
	// 权限更新
	a.Enforcer.RemoveFilteredPolicy(0, ctx.Query("sub"))

	a.Enforcer.AddPolicy(ctx.Query("sub"), ctx.Query("obj"), ctx.Query("act"))
}

// Query godoc
//
//	@Summary		test data list
//	@Description	get list by ID
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Param			query	query		model.TestQuery	false	"查询条件"
//	@Success		200		{object}	result.Result[model.PageData[model.TestResult]]
//	@Router			/test/query [get]
func (a *TestController) Query(ctx *gin.Context) {
	var testQuery model.TestQuery
	if ok, err := xgin.ShouldBindQuery(ctx, &testQuery); !ok {
		result.Error[any](err).Response(ctx)
		return
	}
	res := a.TestService.Query(testQuery)
	res.Response(ctx)
}
