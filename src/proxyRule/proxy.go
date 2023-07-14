package proxyRule

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/types"
	"sync"
)

type SystemProxyRuleData struct {
	sync.Mutex
	rule *map[string]types.UrlLimitRule
}

// SystemProxyRule 代理规则
var SystemProxyRule = new(SystemProxyRuleData)

// Get 获取规则
func (t *SystemProxyRuleData) Get() *map[string]types.UrlLimitRule {
	t.Lock()
	defer t.Unlock()
	return t.rule
}

// GetFormPath 根据路径获取规则
func (t *SystemProxyRuleData) GetFormPath(path string) (*types.UrlLimitRule, bool) {
	var ok bool
	if i, ok := (*t.rule)[path]; ok {
		return &i, ok
	}
	return nil, ok
}

// Refresh 刷新规则
func (t *SystemProxyRuleData) Refresh() {
	t.Lock()
	defer t.Unlock()
	m := new(map[string]types.UrlLimitRule)
	_ = global.XDB.GetGJson().Get("proxy").Scan(m)
	t.rule = m
}
