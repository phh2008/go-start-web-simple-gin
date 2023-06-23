// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
)

// Injectors from wire.go:

func BuildServer(configFolder config.ConfigFolder) *app.Server {
	configConfig := config.NewConfig(configFolder)
	db := orm.NewDB(configConfig)
	zapLogger := logger.NewLogger(configConfig)
	engine := gin.New()
	helloController := controller.NewHelloController()
	jwtHelper := xjwt.NewJwtHelper(configConfig)
	enforcer := xcasbin.NewCasbin(db, configConfig)
	auth := middleware.NewAuth(jwtHelper, enforcer)
	userDao := dao.NewUserDAO(db)
	userService := service.NewUserService(userDao, jwtHelper, enforcer)
	userController := controller.NewUserController(userService)
	permissionDao := dao.NewPermissionDAO(db)
	permissionService := service.NewPermissionService(permissionDao, enforcer)
	permissionController := controller.NewPermissionController(permissionService)
	roleDao := dao.NewRoleDAO(db)
	rolePermissionDao := dao.NewRolePermissionDAO(db)
	roleService := service.NewRoleService(roleDao, rolePermissionDao, permissionDao, enforcer, userDao)
	roleController := controller.NewRoleController(roleService, userService)
	routerRouter := &router.Router{
		Engine:        engine,
		HelloApi:      helloController,
		Auth:          auth,
		UserApi:       userController,
		PermissionApi: permissionController,
		RoleApi:       roleController,
	}
	server := &app.Server{
		Config: configConfig,
		DB:     db,
		Logger: zapLogger,
		Router: routerRouter,
		Engine: engine,
	}
	return server
}
