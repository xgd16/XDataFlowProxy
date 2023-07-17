package timer

import (
	"XDataFlowProxy/src/proxyMode"
	"XDataFlowProxy/src/types"
	"time"
)

// 此处只注册需要定时内存回收的代理模式
var gcArr = []types.ProxyMode{
	new(proxyMode.SequentialAccess),
}

// Service 定时器 (目前用于内存回收)
func Service() {
	for {
		for _, mode := range gcArr {
			mode.MemGC()
		}
		time.Sleep(3000 * time.Millisecond)
	}
}
