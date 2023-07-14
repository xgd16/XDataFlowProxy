package proxyMode

import (
	"XDataFlowProxy/src/types"
	"github.com/gogf/gf/v2/os/gmlock"
)

var requestLock = gmlock.New()

// SequentialAccess 顺序访问
type SequentialAccess struct {
}

func (t *SequentialAccess) ProxyBefore(back *types.ProxyCallBack) {
	requestLock.Lock(back.RuleKey)
}

func (t *SequentialAccess) ProxyAfter(back *types.ProxyCallBack) {
	requestLock.Unlock(back.RuleKey)
}
