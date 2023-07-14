package proxyMode

import (
	"XDataFlowProxy/src/types"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/xgd16/gf-x-tool/xTool"
)

var requestData = gmap.NewStrIntMap(true)

// RepeatedRejection 重复拒绝
type RepeatedRejection struct {
}

func (t *RepeatedRejection) ProxyBefore(back *types.ProxyCallBack) {
	if requestData.Contains(back.RuleKey) {
		xTool.FastResp(back.Request).ErrorStatus().Resp("上个请求未完成", 999999)
	} else {
		requestData.Set(back.RuleKey, 1)
	}
}

func (t *RepeatedRejection) ProxyAfter(back *types.ProxyCallBack) {
	if requestData.Contains(back.RuleKey) {
		requestData.Remove(back.RuleKey)
	}
}
