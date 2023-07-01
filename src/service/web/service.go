package web

import (
	"XDataFlowProxy/src/service/web/controller/proxy"
	"XDataFlowProxy/src/service/web/route"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func Service() {
	server := g.Server()
	// 路由注册
	server.Group("/api", route.Api)
	server.BindHandler("/*", proxy.HttpProxy)
	// 基本配置
	server.BindMiddlewareDefault(ghttp.MiddlewareCORS)
	server.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Writefln(`
			<div style="text-align:center;"><div style="font-size: 5rem">404</div><div style="font-size: 3rem">%s</div></div>
		`, gtime.Now().Format("Y-m-d H:i:s"))
	})
	// 启动web服务
	server.Run()
}
