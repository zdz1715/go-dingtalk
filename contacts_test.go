package dingtalk

import (
	"context"
	"testing"

	"github.com/zdz1715/ghttp"
	"github.com/zdz1715/go-utils/goutils"
)

func TestContactsService_ListDepartments(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Contacts.ListDepartments(context.Background(), &ListDepartmentsOptions{
		//DeptId: goutils.Ptr(1),
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
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	// 	查询全部部门
	reply, err := client.Contacts.ListDepartmentsV1(context.Background(), &ListDepartmentsV1Options{
		FetchChild: goutils.Ptr(true),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestContactsService_ListUsers(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Contacts.ListUsers(context.Background(), &ListUsersOptions{
		DeptId: goutils.Ptr(55436473),
		Cursor: goutils.Ptr(0),
		Size:   goutils.Ptr(2),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestContactsService_GetUserByMobile(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Contacts.GetUserByMobile(context.Background(), &GetUserByMobileOptions{
		Mobile:                        goutils.Ptr("110"),
		SupportExclusiveAccountSearch: goutils.Ptr(true),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}
