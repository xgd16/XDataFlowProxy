package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/xgd16/gf-x-tool/xTool"
)

// PrometheusGetServerInfoMiddleware 普罗米修斯 服务中间键
func PrometheusGetServerInfoMiddleware(r *ghttp.Request) {
	urlStr := gstr.TrimLeft(r.Request.URL.Path, "/")
	mode := gstr.Split(urlStr, "/")[0]
	xTool.MetricHttpRequestTotal.WithLabelValues(mode).Inc()
	xTool.MetricHttpRequestTotal.WithLabelValues(urlStr).Inc()

	r.Middleware.Next()
}
