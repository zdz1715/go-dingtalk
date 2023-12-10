package dingtalk

import (
	"context"
	"time"
)

// OAuthService
// API Docs: https://open.dingtalk.com/document/orgapp/authorization-overview
type OAuthService struct {
	client     *Client
	credential Credential
	store
}

type store struct {
	val    *AccessToken
	expire time.Time
}

func (s *store) value() *AccessToken {
	return s.val
}

func (s *store) memory(at *AccessToken) {
	s.val = at
	if at != nil {
		// 提前5分钟过期, 避免网络带来的延时
		s.expire = time.Now().Add(time.Duration(at.ExpireIn)*time.Second - 5*time.Minute)
	}
}

func (s *store) IsExpired() bool {
	if s.val == nil {
		return true
	}
	if s.val.AccessToken == "" {
		return true
	}
	return time.Now().After(s.expire)
}

type AccessToken struct {
	AccessToken string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	ExpireIn    int64  `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

// GetAccessToken get access token.
// API Docs:  https://open.dingtalk.com/document/orgapp/obtain-the-access_token-of-an-internal-app
func (s *OAuthService) GetAccessToken(ctx context.Context) (*AccessToken, error) {
	if !s.store.IsExpired() {
		return s.store.value(), nil
	}
	if s.credential == nil {
		return nil, ErrCredential
	}
	if err := s.credential.Valid(); err != nil {
		return nil, err
	}
	switch s.credential.AuthType() {
	case AuthTypeApp:
	default:
		return nil, ErrAuthType
	}
	var respBody AccessToken
	err := s.client.Invoke(
		ctx,
		s.credential.Method(),
		s.credential.URL(),
		s.credential.Body(),
		&respBody,
		"",
	)
	if err != nil {
		return nil, err
	}

	s.store.memory(&respBody)

	return &respBody, nil
}
