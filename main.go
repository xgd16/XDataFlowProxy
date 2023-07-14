package main

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/lib"
	"XDataFlowProxy/src/proxyMode"
	"XDataFlowProxy/src/proxyRule"
	"XDataFlowProxy/src/service"
	"XDataFlowProxy/src/types"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

func main() {
	// 初始化系统配置
	global.InitSystemConfig()
	// 基础初始化
	baseInit()
	// 初始化代理信息
	proxyRule.SystemProxyRule.Refresh()
	// 初始化系统服务
	service.InitService()
	// 维持
	lib.Maintain()
}

func baseInit() {
	global.ProxyMode = func() types.ProxyMode {
		switch global.SystemConfig.Get("proxy.mode", 1).Int() {
		case 2:
			return new(proxyMode.RepeatedRejection)
		default:
			return new(proxyMode.SequentialAccess)
		}
	}()
}
