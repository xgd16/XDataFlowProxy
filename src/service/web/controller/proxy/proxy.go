package proxy

import (
	"XDataFlowProxy/src/lib"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xgd16/gf-x-tool/xTool"
)

// proxyHandler 代理处理层
func proxyHandler(back *lib.ProxyCallBack) error {

	return nil
}

// HttpProxy 代理执行
func HttpProxy(r *ghttp.Request) {
	// 发起代理
	proxyErr := lib.SetProxy(r, "http://127.0.0.1:8000", proxyHandler)
	// 判断出错进行错误返回
	xTool.FastResp(r, proxyErr, false).Resp("服务无响应...")
}
