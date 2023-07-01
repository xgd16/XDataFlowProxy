package api

import (
	"XDataFlowProxy/src/global"
	"XDataFlowProxy/src/types"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xgd16/gf-x-tool/xTool"
	"net/url"
)

func SetRule(r *ghttp.Request) {
	routeT := r.Get("route")     // 路径
	limitT := r.Get("limitData") // 限制参数
	// 判断输入参数是否为空
	xTool.FastResp(r, routeT.IsEmpty() || limitT.IsEmpty(), false).Resp("参数错误")
	// 解析路径地址
	parsedURL, err := url.Parse(routeT.String())
	xTool.FastResp(r, err).Resp("操作失败")
	// 存入数据
	xTool.FastResp(r, global.XDB.Set("proxy", parsedURL.Path, types.UrlLimitRule{
		LimitData: limitT.Strings(),
	})).Resp()
	xTool.FastResp(r).Resp()
}

func DelRule(r *ghttp.Request) {
	routeT := r.Get("route")
	// 判断输入参数是否为空
	xTool.FastResp(r, routeT.IsEmpty(), false).Resp("参数错误")
	// 解析路径地址
	parsedURL, err := url.Parse(routeT.String())
	xTool.FastResp(r, err).Resp("操作失败")
	// 删除数据
	xTool.FastResp(r, global.XDB.Del("proxy", parsedURL.Path)).Resp()
	xTool.FastResp(r).Resp()
}
