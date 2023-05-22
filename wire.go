//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"com.gientech/selection/app"
	"com.gientech/selection/dao"
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/logger"
	"com.gientech/selection/pkg/orm"
	"com.gientech/selection/pkg/xcasbin"
	"com.gientech/selection/pkg/xjwt"
	"com.gientech/selection/service"
	"com.gientech/selection/web/controller"
	"com.gientech/selection/web/middleware"
	"com.gientech/selection/web/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func BuildServer(configFolder config.ConfigFolder) *app.Server {
	wire.Build(
		config.ConfigSet,
		logger.LoggerSet,
		orm.InitDB,
		dao.DaoSet,
		service.ServiceSet,
		controller.ControllerSet,
		gin.New,
		xjwt.NewJwtHelper,
		xcasbin.NewCasbin,
		middleware.AuthSet,
		router.RouterSet,
		app.ServerSet,
	)
	return new(app.Server)
}
