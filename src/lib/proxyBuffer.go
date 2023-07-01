package lib

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/types"
	"sync"
)

// ProxyBuffer 代理缓冲区
var ProxyBuffer = CreateProxyBuffer()

// ProxyBufferType 代理缓冲
type ProxyBufferType struct {
	Domain   string
	RuleList map[string]*types.UrlLimitRule
	lock     sync.RWMutex
}

// CreateProxyBuffer 创建代理缓冲数据区
func CreateProxyBuffer() *ProxyBufferType {
	buffer := &ProxyBufferType{
		Domain:   global.SystemConfig.Get("proxy.domain").String(),
		RuleList: make(map[string]*types.UrlLimitRule),
	}
	// 首次创建时初始化
	buffer.Refresh()
	return buffer
}

// Refresh 刷新缓冲区
func (t *ProxyBufferType) Refresh() {
	t.lock.Lock()
	defer t.lock.Unlock()
	for route, item := range global.XDB.GetGJson().GetJsonMap("proxy") {
		t.RuleList[route] = &types.UrlLimitRule{
			LimitData: item.Get("limitData").Strings(),
		}
	}
}
