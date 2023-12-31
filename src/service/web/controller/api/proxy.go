package api

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/proxyRule"
	"XDataFlowProxy/src/types"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xgd16/gf-x-tool/xTool"
	"net/url"
)

func verifyKey(r *ghttp.Request) {
	xTool.FastResp(r, r.Get("key").String() != global.SystemConfig.Get("server.key").String(), false).Resp("无效请求")
}

// RuleList 规则列表
func RuleList(r *ghttp.Request) {
	verifyKey(r)
	xTool.FastResp(r).SetData(proxyRule.SystemProxyRule.Get()).Resp()
}

func SetRule(r *ghttp.Request) {
	verifyKey(r)
	routeT := r.Get("route")     // 路径
	limitT := r.Get("limitData") // 限制参数
	// 判断输入参数是否为空
	xTool.FastResp(r, routeT.IsEmpty() || limitT.IsEmpty(), false).Resp("参数错误")
	// 解析路径地址
	parsedURL, err := url.Parse(routeT.String())
	xTool.FastResp(r, err).Resp("操作失败")
	// 存入数据
	limitData := new(types.UrlLimitRule)
	data, err := r.GetJson()
	xTool.FastResp(r, err).Resp()
	xTool.FastResp(r, data.Scan(limitData)).Resp()
	xTool.FastResp(r, !xTool.InArr(limitData.Mode, []int{0, 1, 2}), false).Response("参数错误")
	xTool.FastResp(r, global.XDB.Set("proxy", parsedURL.Path, limitData)).Resp()
	proxyRule.SystemProxyRule.Refresh()
	xTool.FastResp(r).Resp()
}

func DelRule(r *ghttp.Request) {
	verifyKey(r)
	routeT := r.Get("route")
	// 判断输入参数是否为空
	xTool.FastResp(r, routeT.IsEmpty(), false).Resp("参数错误")
	// 解析路径地址
	parsedURL, err := url.Parse(routeT.String())
	xTool.FastResp(r, err).Resp("操作失败")
	// 删除数据
	xTool.FastResp(r, global.XDB.Del("proxy", parsedURL.Path)).Resp()
	proxyRule.SystemProxyRule.Refresh()
	xTool.FastResp(r).Resp()
}
