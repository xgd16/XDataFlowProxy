package types

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http/httputil"
)

// UrlLimitRule url 限制规则
type UrlLimitRule struct {
	Mode      int       `json:"mode"`
	LimitData LimitData `json:"limitData"`
}

type LimitData struct {
	ReqData []string `json:"reqData"`
	Header  []string `json:"header"`
}

type ProxyCallBack struct {
	Proxy   *httputil.ReverseProxy
	Request *ghttp.Request
	Path    string
	RuleKey string
}

type ProxyMode interface {
	ProxyBefore(back *ProxyCallBack)
	ProxyAfter(back *ProxyCallBack)
	MemGC()
}
