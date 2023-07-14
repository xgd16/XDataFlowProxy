package proxy

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/lib"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/xgd16/gf-x-tool/xTool"
)

var requestLock = gmlock.New()

// 代理处理层
func proxyBefore(back *lib.ProxyCallBack) {
	requestLock.Lock(back.RuleKey)
}

func proxyAfter(back *lib.ProxyCallBack) {
	requestLock.Unlock(back.RuleKey)
}

// HttpProxy 代理执行
func HttpProxy(r *ghttp.Request) {
	// 发起代理
	proxyErr := lib.SetProxy(
		r,
		global.SystemConfig.Get("proxy.domain").String(),
		proxyBefore,
		proxyAfter,
	)
	// 判断出错进行错误返回
	xTool.FastResp(r, proxyErr, false).Resp("服务无响应...")
}
