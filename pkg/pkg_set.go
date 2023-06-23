package pkg

import (
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/logger"
	"com.gientech/selection/pkg/orm"
	"com.gientech/selection/pkg/xcasbin"
	"com.gientech/selection/pkg/xjwt"
	"github.com/google/wire"
)

// ProviderSet is pkg provider set
var ProviderSet = wire.NewSet(
	config.NewConfig,
	logger.NewLogger,
	orm.NewDB,
	xjwt.NewJwtHelper,
	xcasbin.NewCasbin,
)
