package dingtalk

import (
	"context"
	"os"
	"testing"

	"github.com/zdz1715/ghttp"
)

var testInternalAppCredential = &InternalAppCredential{
	AppKey:    os.Getenv("TEST_DINGTALK_APP_KEY"),
	AppSecret: os.Getenv("TEST_DINGTALK_APP_SECRET"),
}

func TestOAuthService_GetAccessToken(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
		//InitSkipCredential: true,
	})

	if err != nil {
		t.Fatal(err)
	}

	token, err := client.OAuth.GetAccessToken(context.Background())
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", token)
	}

	t.Logf("is expired: %t", client.OAuth.IsExpired())

	// 在有效期内不会再请求接口
	token, err = client.OAuth.GetAccessToken(context.Background())
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%+v", token)
	}
}
