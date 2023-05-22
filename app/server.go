package app

import (
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/logger"
	"com.gientech/selection/web/middleware"
	"com.gientech/selection/web/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ServerSet = wire.NewSet(wire.Struct(new(Server), "*"))

type Server struct {
	Config *config.Config
	DB     *gorm.DB
	Logger *zap.Logger
	Router *router.Router
	Engine *gin.Engine
}

func (a *Server) Start() {
	//gin.SetMode(gin.ReleaseMode)
	profile := config.GetProfile()
	logger.Infof(">>>>>>>>>>>服务正在起启动，运行环境为：%s <<<<<<<<<<<<", profile.Server.Env)
	a.Engine.Use(middleware.GinLogger)
	a.Engine.Use(middleware.GinRecovery(true))
	a.Engine.Use(middleware.Translations())
	a.Engine.Use(middleware.Cors(profile.Cors))
	a.Router.Register()
	err := a.Engine.Run(fmt.Sprintf("0.0.0.0:%s", profile.Server.Port))
	if err != nil {
		logger.Errorf("server start error: %s", err.Error())
	}
}
