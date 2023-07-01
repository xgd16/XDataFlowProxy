package proxy

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/lib"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gmutex"
	"github.com/xgd16/gf-x-tool/xTool"
)

var ReqLockMap = gmap.NewStrAnyMap(true)

// proxyHandler 代理处理层
func proxyBefore(back *lib.ProxyCallBack) {
	if item, ok := lib.ProxyBuffer.RuleList[back.Request.URL.Path]; ok {
		var str string
		for _, datum := range item.LimitData {
			s := back.Request.Get(datum).String()
			if s == "" {
				return
			}
			str += fmt.Sprintf("%s-%s", datum, s)
		}
		// 生成一个用来限制的key
		key := fmt.Sprintf("%s_%s", back.Request.URL.Path, str)
		back.Key = key
		back.ReqLock = ReqLockMap.GetOrSetFunc(key, func() any {
			mutex := gmutex.New()
			mutex.Lock()
			return mutex
		}).(*gmutex.Mutex)
		fmt.Println("进入")
	}
	return
}

func proxyAfter(back *lib.ProxyCallBack) {
	back.ReqLock.Unlock()
}

// HttpProxy 代理执行
func HttpProxy(r *ghttp.Request) {
	// 发起代理
	proxyErr := lib.SetProxy(r, global.SystemConfig.Get("proxy.domain").String(), proxyBefore, proxyAfter)
	// 判断出错进行错误返回
	xTool.FastResp(r, proxyErr, false).Resp("服务无响应...")
}
