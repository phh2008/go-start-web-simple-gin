package router

import (
	_ "com.gientech/selection/docs"
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/web/controller"
	"com.gientech/selection/web/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

type Router struct {
	Engine        *gin.Engine
	HelloApi      *controller.HelloController
	TestApi       *controller.TestController
	Auth          *middleware.Auth
	UserApi       *controller.UserController
	PermissionApi *controller.PermissionController
	RoleApi       *controller.RoleController
}

// Register 注册路由
func (a *Router) Register() {
	v1 := a.Engine.Group("/api/v1")
	{
		//测试用的API，请删除
		v1.GET("/hello", a.HelloApi.Hello)
		v1.GET("/test/token", a.TestApi.GetToken)
		v1.GET("/test/auth", a.Auth.Authenticate(), a.TestApi.Auth)          // 要登录
		v1.GET("/test/query", a.Auth.Authorization("view"), a.TestApi.Query) //要权限
	}
	{
		//用户API
		v1.POST("/user/login", a.UserApi.Login)
		v1.POST("/user/createByEmail", a.Auth.Authenticate(), a.UserApi.CreateByEmail)
		v1.POST("/user/assignRole", a.UserApi.AssignRole)
		v1.DELETE("/user/delete/:id", a.Auth.Authenticate(), a.UserApi.DeleteById)
		//角色API
		v1.POST("/role/add", a.RoleApi.Add)
		v1.POST("/role/assignPermission", a.RoleApi.AssignPermission)
		//权限API
		v1.POST("/permission/add", a.PermissionApi.AddPermission)
	}

	// use ginSwagger middleware to serve the API docs
	profile := config.GetProfile()
	if profile.Server.Env == "dev" {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
