package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/zdz1715/ghttp"
)

type service struct {
	client *Client
}

type Options struct {
	ClientOpts         []ghttp.ClientOption
	InitSkipCredential bool
}

type Client struct {
	cc *ghttp.Client

	opts *Options

	common service

	// Services used for talking to different parts of the Jira API.
	OAuth *OAuthService
	// 通讯录
	Contacts *ContactsService
}

func NewClient(credential Credential, opts *Options) (*Client, error) {
	if opts == nil {
		opts = &Options{}
	}

	clientOptions := make([]ghttp.ClientOption, 0)

	if len(opts.ClientOpts) > 0 {
		clientOptions = append(clientOptions, opts.ClientOpts...)
	}

	// 覆盖Endpoint
	clientOptions = append(clientOptions, ghttp.WithEndpoint("https://api.dingtalk.com"))
	// 新版API错误处理
	clientOptions = append(clientOptions, ghttp.WithNot2xxError(&Error{}))

	cc, err := ghttp.NewClient(context.Background(), clientOptions...)
	if err != nil {
		return nil, err
	}

	c := &Client{
		cc:   cc,
		opts: opts,
	}

	c.common.client = c

	c.OAuth = &OAuthService{client: c.common.client}
	c.Contacts = (*ContactsService)(&c.common)

	if opts.InitSkipCredential {
		return c, nil
	}

	if err = c.SetCredential(credential); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) SetCredential(credential Credential) error {
	if credential == nil {
		return ErrCredential
	}

	if err := credential.Valid(); err != nil {
		return err
	}

	if c.OAuth != nil {
		c.OAuth.credential = credential
	}

	return nil
}

func (c *Client) IsOldAPI(req *http.Request, checkV1 bool) bool {
	if req == nil || req.URL == nil {
		return false
	}
	if !checkV1 {
		return req.URL.Host == "oapi.dingtalk.com"
	}
	return req.URL.Host == "oapi.dingtalk.com" && strings.HasPrefix(req.URL.Path, "/topapi/")
}

func (c *Client) InvokeByToken(ctx context.Context, method, path string, args interface{}, reply interface{}) error {
	accessToken, err := c.OAuth.GetAccessToken(ctx)
	if err != nil {
		return err
	}
	return c.Invoke(ctx, method, path, args, reply, accessToken.AccessToken)
}

func (c *Client) Invoke(ctx context.Context, method, path string, args interface{}, reply interface{}, token string) error {
	callOpts := &ghttp.CallOptions{
		BeforeHook: func(request *http.Request) error {
			if token != "" {
				if c.IsOldAPI(request, false) {
					query := request.URL.Query()
					query.Set("access_token", token)
					request.URL.RawQuery = query.Encode()
				} else {
					// 新版API鉴权方式
					request.Header.Set("x-acs-dingtalk-access-token", token)
				}
			}
			return nil
		},
		// 处理旧版API返回错误
		AfterHook: func(response *http.Response) error {
			if !c.IsOldAPI(response.Request, false) {
				return nil
			}
			// 处理 https://oapi.dingtalk.com/
			all, err := io.ReadAll(response.Body)
			if err != nil {
				return err
			}
			_ = response.Body.Close()

			result := new(Result)
			if err = json.Unmarshal(all, result); err != nil {
				return err
			}
			if err = result.Error(); err != nil {
				return err
			}
			if c.IsOldAPI(response.Request, true) {
				// 处理 https://oapi.dingtalk.com/topapi/
				byt, err := json.Marshal(result.Result)
				if err != nil {
					return err
				}
				response.Body = io.NopCloser(bytes.NewReader(byt))
			} else {
				response.Body = io.NopCloser(bytes.NewReader(all))
			}
			return nil
		},
	}

	if method == http.MethodGet && args != nil {
		callOpts.Query = args
		args = nil
	}

	_, err := c.cc.Invoke(ctx, method, path, args, reply, callOpts)
	return err
}
