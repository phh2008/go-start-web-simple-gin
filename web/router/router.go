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
	Engine   *gin.Engine
	HelloApi *controller.HelloController
	TestApi  *controller.TestController
	Auth     *middleware.Auth
}

// Register 注册路由
func (a *Router) Register() {
	v1 := a.Engine.Group("/api/v1")
	{
		v1.GET("/hello", a.HelloApi.Hello)
		v1.GET("/test/token", a.TestApi.GetToken)
		// 需要登录
		v1.GET("/test/auth", a.Auth.Authenticate(), a.TestApi.Auth)
		// 需要权限校验
		v1.GET("/test/query", a.Auth.Authorization("view"), a.TestApi.Query)
	}
	// use ginSwagger middleware to serve the API docs
	profile := config.GetProfile()
	if profile.Server.Env == "dev" {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
