package pkg

import (
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/logger"
	"com.gientech/selection/pkg/orm"
	"com.gientech/selection/pkg/xcasbin"
	"com.gientech/selection/pkg/xjwt"
	"github.com/google/wire"
)

var ToolSet = wire.NewSet(
	config.ConfigSet,
	logger.LoggerSet,
	orm.InitDB,
	xjwt.NewJwtHelper,
	xcasbin.NewCasbin,
)
