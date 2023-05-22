package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ConfigSet = wire.NewSet(NewConfig)

// Env 当前环境，比如：dev，测会加载 config-dev.yml的配置
var Env string

var profile Profile

// ConfigFolder 配置文件目录
type ConfigFolder string

type Config struct {
	ConfigDir ConfigFolder
	Viper     *viper.Viper
	Profile   Profile
}

func NewConfig(configFolder ConfigFolder) *Config {
	var env string
	if Env != "" {
		env = "-" + Env
	}
	vp := viper.New()
	vp.SetConfigName("config" + env)
	vp.SetConfigType("yml")
	vp.AddConfigPath(string(configFolder))
	err := vp.ReadInConfig()
	if err != nil {
		zap.S().Errorf("加载配置错误,error:%s", err.Error())
		panic(err)
	}
	if err = vp.Unmarshal(&profile); err != nil {
		zap.S().Errorf("绑定配置出错,error:%s", err.Error())
		panic(err)
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("config file changed:%s", e.Name)
		if err = vp.Unmarshal(&profile); err != nil {
			zap.S().Errorf("更新配置出错,error:%s", err.Error())
		}
	})
	return &Config{
		ConfigDir: configFolder,
		Viper:     vp,
		Profile:   profile,
	}
}

func (a *Config) GetString(key string) string {
	return a.Viper.GetString(key)
}

func (a *Config) Get(key string) interface{} {
	return a.Viper.Get(key)
}

func GetProfile() Profile {
	p := profile
	return p
}
