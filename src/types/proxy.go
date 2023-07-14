package types

// UrlLimitRule url 限制规则
type UrlLimitRule struct {
	LimitData LimitData `json:"limitData"`
}

type LimitData struct {
	ReqData []string `json:"reqData"`
	Header  []string `json:"header"`
}
