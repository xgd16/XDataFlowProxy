package proxy

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/lib"
	"XDataFlowProxy/src/types"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xgd16/gf-x-tool/xTool"
)

// 代理处理层
func proxyBefore(back *types.ProxyCallBack) {

}

func proxyAfter(back *types.ProxyCallBack) {

}

// HttpProxy 代理执行
func HttpProxy(r *ghttp.Request) {
	// 发起代理
	proxyErr := lib.SetProxy(
		r,
		global.ProxyMode,
		global.SystemConfig.Get("proxy.domain").String(),
		proxyBefore,
		proxyAfter,
	)
	fmt.Println("代理出错", proxyErr)
	// 判断出错进行错误返回
	xTool.FastResp(r, proxyErr).Resp("服务无响应...")
}
