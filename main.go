package main

import (
	"com.gientech/selection/pkg/config"
	"flag"
)

//	@title						XXX API
//	@version					1.0
//	@description				This is a sample server celler server.
//	@host						localhost:8089
//	@BasePath					/api/v1
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	// 参数
	var configFolder string
	flag.StringVar(&configFolder, "config", "./config", "指定配置文件目录，示例: -config ./config")
	flag.StringVar(&config.Env, "env", "dev", "指定当前运行环境,示例: -env test")
	flag.Parse()
	// wire
	server := BuildServer(config.ConfigFolder(configFolder))
	// 启动服务
	server.Start()
}
