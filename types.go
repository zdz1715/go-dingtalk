package dingtalk

import "fmt"

// Error
// 新版 API docs: https://open.dingtalk.com/document/orgapp/error-code
type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestid"`
}

func (e *Error) String() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Code != "" {
		return e.Code
	}
	return ""
}

func (e *Error) Reset() {
	e.Code = ""
	e.Message = ""
}

// Result
// 旧版 API Docs: https://open.dingtalk.com/document/orgapp/server-api-error-codes-1
type Result struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Result  interface{} `json:"result"`
}

func (r *Result) Error() error {
	if r.Errcode == 0 {
		return nil
	}
	return fmt.Errorf("errcode=%d errmsg=%s", r.Errcode, r.Errmsg)
}

type AuthType uint8

const (
	AuthTypeApp = iota
	// todo: 获取登录用户的访问凭证
	// https://open.dingtalk.com/document/orgapp/obtain-identity-credentials
	//AuthTypeCode
)

type Language string

const (
	ZH_CN Language = "zh_CN"
	EN_US Language = "en_US"
)
