package proxyMode

import (
	"XDataFlowProxy/src/types"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/os/gmlock"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

var requestLock = gmlock.New()
var memClean = gmap.NewStrIntMap(true)

// SequentialAccess 顺序访问
type SequentialAccess struct {
}

func (t *SequentialAccess) MemGC() {
	for s, i := range memClean.Map() {
		if int64(i) < gtime.Now().Unix() {
			memClean.Remove(s)
		}
	}
}

func (t *SequentialAccess) ProxyBefore(back *types.ProxyCallBack) {
	requestLock.Lock(back.RuleKey)
}

func (t *SequentialAccess) ProxyAfter(back *types.ProxyCallBack) {
	memClean.Set(back.RuleKey, int(gtime.Now().Add(60*time.Second).Unix()))
	requestLock.Unlock(back.RuleKey)
}
