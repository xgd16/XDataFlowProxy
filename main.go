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
	"github.com/xgd16/gf-x-tool/xTool"
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
	// 创建 普罗米修斯数值
	xTool.InitPrometheusMetric(
		global.SystemConfig.Get("prometheus.namespace").String(),
		global.SystemConfig.Get("prometheus.subsystem").String(),
	)
	// 设置代理模式
	global.ProxyMode = func() types.ProxyMode {
		switch global.SystemConfig.Get("proxy.mode", 1).Int() {
		case 2:
			return new(proxyMode.RepeatedRejection)
		default:
			return new(proxyMode.SequentialAccess)
		}
	}()
}
