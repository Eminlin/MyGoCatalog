package main

import (
	"newframework/internal/middleware"
	"newframework/internal/router"
	"newframework/pkg/config"
	"newframework/pkg/db"
	"newframework/pkg/log"
	"newframework/pkg/version"
	"newframework/server"
)

func main() {
	// 解析服务器启动参数
	appOpt := &server.AppOptions{}
	server.ResolveAppOptions(appOpt)
	if appOpt.PrintVersion {
		version.PrintVersion()
	}
	// 加载配置文件
	config.Load(appOpt.ConfigFilePath)
	log.InitLogger() // 日志
	db.InitDB()      // 创建数据库链接，使用默认的实现方式
	// 创建HTTPServer
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		if db.Database != nil {
			sqlDB, err := db.Database.DB()
			if err != nil {
				log.Log().Error(err)
			}
			if err = sqlDB.Close(); err != nil {
				log.Log().Error(err)
			}
		}
	})
	router := router.InitRouter()
	srv.Run(middleware.NewMiddleware(), router)
}
