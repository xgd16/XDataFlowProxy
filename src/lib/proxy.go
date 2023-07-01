package lib

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http/httputil"
	"net/url"
)

type ProxyCallBack struct {
	Proxy   *httputil.ReverseProxy
	Request *ghttp.Request
}

// SetProxy 设置代理
func SetProxy(r *ghttp.Request, toDomain string, callback func(back *ProxyCallBack) error) error {
	// 创建需要被代理的对象
	parse, err := url.Parse(toDomain)
	if err != nil {
		return err
	}
	// 创建反向代理对象
	proxy := httputil.NewSingleHostReverseProxy(parse)
	// 调用外部代码
	if callBackErr := callback(&ProxyCallBack{Proxy: proxy, Request: r}); err != nil {
		return callBackErr
	}
	// 转发代理后的请求
	proxy.ServeHTTP(r.Response.ResponseWriter, r.Response.Request.Request)
	return nil
}
