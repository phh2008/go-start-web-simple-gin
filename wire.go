//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"com.gientech/selection/app"
	"com.gientech/selection/dao"
	"com.gientech/selection/pkg"
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/service"
	"com.gientech/selection/web/controller"
	"com.gientech/selection/web/middleware"
	"com.gientech/selection/web/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func BuildServer(configFolder config.ConfigFolder) *app.Server {
	wire.Build(
		pkg.ToolSet,
		dao.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		gin.New,
		middleware.ProviderSet,
		router.ProviderSet,
		app.ServerSet,
	)
	return new(app.Server)
}
