package xcasbin

import (
	"com.gientech/selection/pkg/config"
	"com.gientech/selection/pkg/logger"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"path/filepath"
)

func NewCasbin(db *gorm.DB, conf *config.Config) *casbin.Enforcer {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		logger.Errorf("casbin gorm 适配器创建失败,error:%s", err.Error())
		panic(err)
	}
	configFile := filepath.Join(string(conf.ConfigDir), "rbac_model.conf")
	rbacEnforcer, err := casbin.NewEnforcer(configFile, adapter)
	if err != nil {
		logger.Errorf("casbin.NewEnforcer 错误,error:%s", err.Error())
		panic(err)
	}
	rbacEnforcer.EnableAutoSave(true)
	// Load the policy from DB.
	_ = rbacEnforcer.LoadPolicy()
	return rbacEnforcer
}
