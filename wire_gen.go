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
	db := orm.InitDB(configConfig)
	zapLogger := logger.InitLogger(configConfig)
	engine := gin.New()
	helloController := &controller.HelloController{}
	testDAO := &dao.TestDAO{
		Db: db,
	}
	testService := &service.TestService{
		TestDao: testDAO,
	}
	jwtHelper := xjwt.NewJwtHelper(configConfig)
	enforcer := xcasbin.NewCasbin(db, configConfig)
	testController := &controller.TestController{
		TestService: testService,
		Jwt:         jwtHelper,
		Enforcer:    enforcer,
	}
	auth := &middleware.Auth{
		Jwt:      jwtHelper,
		Enforcer: enforcer,
	}
	routerRouter := &router.Router{
		Engine:   engine,
		HelloApi: helloController,
		TestApi:  testController,
		Auth:     auth,
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
