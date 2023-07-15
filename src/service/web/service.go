package web

import (
	"XDataFlowProxy/src/middleware"
	"XDataFlowProxy/src/service/web/controller/proxy"
	"XDataFlowProxy/src/service/web/route"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xgd16/gf-x-tool/xTool"
)

func Service() {
	server := g.Server()
	// 路由注册
	group := server.Group("/")
	{
		// 注册中间键
		group.Middleware(ghttp.MiddlewareCORS, middleware.PrometheusGetServerInfoMiddleware)
		// 创建
		group.Group("/xdf-api", route.Api)
		group.ALL("/*", proxy.HttpProxy)
	}
	// 注册普罗米修斯
	server.BindHandler("/metrics", xTool.PrometheusHttp)
	// 处理 404 页面
	server.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Writefln(`
			<div style="text-align:center;"><div style="font-size: 5rem">404</div><div style="font-size: 3rem">%s</div></div>
		`, gtime.Now().Format("Y-m-d H:i:s"))
	})
	// 启动web服务
	server.Run()
}
