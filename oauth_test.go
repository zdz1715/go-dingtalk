package dingtalk

import (
	"context"
	"github.com/zdz1715/ghttp"
	"testing"
)

var testInternalAppCredential = &InternalAppCredential{
	AppKey:    "dingyu6tvxassvlwmrn3",
	AppSecret: "VctzwFg9LZ0PIHoKmYxx5tYwslCOhnQC0oH3tjVAvY2tsmcNwuQTGIAADjaQKiGP",
}

func TestOAuthService_GetAccessToken(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
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
