package main

import (
	"gin-create/core"
	"gin-create/global"
	"gin-create/initialize"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	if global.GVA_DB != nil {
		// 程序结束前关闭数据库连接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
}
