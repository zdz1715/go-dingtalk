package dingtalk

import (
	"errors"
	"net/http"
)

var (
	ErrCredential = errors.New("invalid credential")
	ErrNilBody    = errors.New("nil body")
)

type Credential interface {
	URL() string
	Method() string
	Body(opts *GetAccessTokenOptions) any
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

func (c *InternalAppCredential) Method() string {
	return http.MethodPost
}

func (c *InternalAppCredential) Body(opts *GetAccessTokenOptions) any {
	return c
}

func (c *InternalAppCredential) Valid() error {
	if c.AppKey == "" || c.AppSecret == "" {
		return ErrCredential
	}
	return nil
}
