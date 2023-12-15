package dingtalk

import (
	"context"
	"time"
)

type GetAccessTokenOptions struct {
	Code         string
	RefreshToken string
}

type AccessToken struct {
	AccessToken  string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	RefreshToken string `xml:"refreshToken" xml:"refreshToken"`
	CorpId       string `json:"corpId" xml:"corpId"`
	ExpireIn     int64  `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

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

// GetAccessToken get access token.
// API Docs:  https://open.dingtalk.com/document/orgapp/obtain-the-access_token-of-an-internal-app
func (s *OAuthService) GetAccessToken(ctx context.Context, opts ...*GetAccessTokenOptions) (*AccessToken, error) {
	opt := new(GetAccessTokenOptions)
	if len(opts) > 0 && opts[0] != nil {
		opt = opts[0]
	}

	// 如果是刷新token，则强制刷新
	if opt.RefreshToken == "" && !s.store.IsExpired() {
		return s.store.value(), nil
	}

	if s.credential == nil {
		return nil, ErrCredential
	}
	if err := s.credential.Valid(); err != nil {
		return nil, err
	}

	body := s.credential.Body(opt)

	if body == nil {
		return nil, ErrNilBody
	}

	var respBody AccessToken
	err := s.client.Invoke(
		ctx,
		s.credential.Method(),
		s.credential.URL(),
		body,
		&respBody,
		"",
	)
	if err != nil {
		return nil, err
	}

	s.store.memory(&respBody)

	return &respBody, nil
}
