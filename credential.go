package dingtalk

import (
	"errors"
	"net/http"
)

var (
	ErrCredential = errors.New("invalid credential")
	ErrAuthType   = errors.New("invalid authType")
)

type Credential interface {
	AuthType() AuthType
	Method() string
	URL() string
	Body() interface{}
	Valid() error
}

// InternalAppCredential
// API docs: https://open.dingtalk.com/document/orgapp/obtain-the-access_token-of-an-internal-app
type InternalAppCredential struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

func (c *InternalAppCredential) URL() string {
	return "/v1.0/oauth2/accessToken"
}

func (c *InternalAppCredential) Body() interface{} {
	return c
}

func (c *InternalAppCredential) Valid() error {
	if c.AppKey == "" || c.AppSecret == "" {
		return ErrCredential
	}
	return nil
}

func (c *InternalAppCredential) AuthType() AuthType {
	return AuthTypeApp
}

func (c *InternalAppCredential) Method() string {
	return http.MethodPost
}
