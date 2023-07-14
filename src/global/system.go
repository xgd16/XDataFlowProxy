package global

import (
	"XDataFlowProxy/src/types"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/xgd16/gf-x-tool/xTool"
)

// SystemConfig 系统配置信息
var SystemConfig *gjson.Json

// InitSystemConfig 初始化系统配置信息
func InitSystemConfig() {
	cfg, err := g.Cfg().Data(gctx.New())
	if err != nil {
		panic("初始化系统配置错误: " + err.Error())
	}
	SystemConfig = gjson.New(cfg, true)
}

// XDB 文件数据存储
var XDB = xTool.CreateXDB()

// ProxyMode 代理模式
var ProxyMode types.ProxyMode
