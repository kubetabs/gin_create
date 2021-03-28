package main

import (
	"gin-create/core"
	"gin-create/global"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
}
