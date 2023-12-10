package dingtalk

import (
	"context"
	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-utils/goutils"
	"testing"
)

func TestContactsService_ListDepartments(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Contacts.ListDepartments(context.Background(), &ListDepartmentsOptions{
		DeptId: goutils.Ptr(1),
		//Language: utils.Ptr(EN_US),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestContactsService_ListDepartmentsV1(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(true),
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	// 	查询全部部门
	reply, err := client.Contacts.ListDepartmentsV1(context.Background(), &ListDepartmentsV1Options{
		FetchChild: true,
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
