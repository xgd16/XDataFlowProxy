package main

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/lib"
	"XDataFlowProxy/src/service"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	lib.ProxyBuffer.Refresh()
	// 初始化系统配置
	global.InitSystemConfig()
	// 初始化系统服务
	service.InitService()
	// 维持
	lib.Maintain()
}
