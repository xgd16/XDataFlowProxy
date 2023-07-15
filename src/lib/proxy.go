package lib

import (
	"XDataFlowProxy/src/proxyRule"
	"XDataFlowProxy/src/types"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http/httputil"
	"net/url"
)

// SetProxy 设置代理
func SetProxy(r *ghttp.Request, proxyMode types.ProxyMode, toDomain string, cbBefore, cbAfter func(back *types.ProxyCallBack)) error {
	// 创建需要被代理的对象
	parse, err := url.Parse(toDomain)
	if err != nil {
		return err
	}
	// 创建反向代理对象
	proxy := httputil.NewSingleHostReverseProxy(parse)
	// 调用外部代码
	proxyData := &types.ProxyCallBack{
		Proxy:   proxy,
		Request: r,
		Path:    r.Request.URL.Path,
	}
	// 处理前置数据
	ok := handlerBeforeData(proxyData)
	// 前置处理
	if ok {
		proxyMode.ProxyBefore(proxyData)
		cbBefore(proxyData)
	}
	// 转发代理后的请求
	proxy.ServeHTTP(r.Response.Writer.RawWriter(), r.Request)
	// 后置处理
	if ok {
		cbAfter(proxyData)
		proxyMode.ProxyAfter(proxyData)
	}
	return nil
}

func handlerBeforeData(data *types.ProxyCallBack) (ok bool) {
	// 获取基础配置规则
	rule, ok := proxyRule.SystemProxyRule.GetFormPath(data.Path)
	if ok {
		data.RuleKey = getRequestRuleKey(
			data.Request,
			garray.NewStrArrayFrom(rule.LimitData.ReqData).Sort(false).Slice(),
			garray.NewStrArrayFrom(rule.LimitData.Header).Sort(false).Slice(),
		)
	}
	return
}

func getRequestRuleKey(r *ghttp.Request, reqData, header []string) string {
	var k []string
	for _, i1 := range reqData {
		k = append(k, fmt.Sprintf("%s=%s", i1, r.Get(i1).String()))
	}
	for _, i1 := range header {
		k = append(k, fmt.Sprintf("%s=%s", i1, r.GetHeader(i1)))
	}
	return gmd5.MustEncrypt(gstr.Join(k, "_"))
}
