package route

import (
	"XDataFlowProxy/src/service/web/controller/api"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Api 路由注册
func Api(r *ghttp.RouterGroup) {
	xProxy := r.Group("/x-proxy")

	xProxy.POST("/setRule", api.SetRule)
	xProxy.POST("/delRule", api.DelRule)
	xProxy.GET("/ruleList", api.RuleList)
}
