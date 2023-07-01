package lib

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gmutex"
	"net/http/httputil"
	"net/url"
)

type ProxyCallBack struct {
	Proxy   *httputil.ReverseProxy
	Request *ghttp.Request
	Key     string
	ReqLock *gmutex.Mutex
}

// SetProxy 设置代理
func SetProxy(r *ghttp.Request, toDomain string, cbBefore, cbAfter func(back *ProxyCallBack)) error {
	// 创建需要被代理的对象
	parse, err := url.Parse(toDomain)
	if err != nil {
		return err
	}
	// 创建反向代理对象
	proxy := httputil.NewSingleHostReverseProxy(parse)
	// 调用外部代码
	proxyData := &ProxyCallBack{Proxy: proxy, Request: r}
	cbBefore(proxyData)
	// 转发代理后的请求
	proxy.ServeHTTP(r.Response.ResponseWriter, r.Response.Request.Request)
	cbAfter(proxyData)
	return nil
}
