package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zdz1715/ghttp"
)

type service struct {
	client *Client
}

type Options struct {
	ClientOpts []ghttp.ClientOption
}

type Client struct {
	cc *ghttp.Client

	opts *Options

	common service

	// Services used for talking to different parts of the Jira API.
	OAuth *OAuthService
	// 通讯录
	Contacts *ContactsService
	// 消息通知
	Message *MessagesService
	// 机器人
	Robot *RobotsService
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
	clientOptions = append(clientOptions, ghttp.WithNot2xxError(func() ghttp.Not2xxError {
		return new(Error)
	}))

	c := &Client{
		cc:   ghttp.NewClient(clientOptions...),
		opts: opts,
	}

	c.common.client = c

	c.OAuth = &OAuthService{client: c.common.client}
	c.Contacts = (*ContactsService)(&c.common)
	c.Message = (*MessagesService)(&c.common)
	c.Robot = (*RobotsService)(&c.common)

	if credential != nil {
		if err := c.SetCredential(credential); err != nil {
			return nil, err
		}
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

func (c *Client) IsOldAPI(req *http.Request, topapiV2 bool) bool {
	if req == nil || req.URL == nil {
		return false
	}
	if !topapiV2 {
		return req.URL.Host == "oapi.dingtalk.com"
	}
	return req.URL.Host == "oapi.dingtalk.com" && strings.HasPrefix(req.URL.Path, "/topapi/v2/")
}

func (c *Client) InvokeByToken(ctx context.Context, method, path string, args any, reply any) error {
	accessToken, err := c.OAuth.GetAccessToken(ctx)
	if err != nil {
		return err
	}
	return c.Invoke(ctx, method, path, args, reply, accessToken.AccessToken)
}

func (c *Client) Invoke(ctx context.Context, method, path string, args any, reply any, token string) error {
	callOpts := &ghttp.CallOptions{
		BeforeHook: func(request *http.Request) error {
			if token != "" {
				if request.URL.Host == "oapi.dingtalk.com" {
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
			// 新版api根据http code处理
			if response.Request.URL.Host != "oapi.dingtalk.com" {
				return nil
			}
			// 处理旧版api
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
			if strings.HasPrefix(response.Request.URL.Path, "/topapi/v2/") {
				// 处理 https://oapi.dingtalk.com/topapi/v2/
				byt, err := json.Marshal(result.Result)
				if err != nil {
					return err
				}
				response.Body = io.NopCloser(bytes.NewReader(byt))
			} else {
				// 处理 https://oapi.dingtalk.com/topapi/
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
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  any    `json:"result"`
}

func (r *Result) Error() error {
	if r.Errcode == 0 {
		return nil
	}
	return fmt.Errorf("errcode=%d errmsg=%s", r.Errcode, r.Errmsg)
}
