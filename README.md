# go-dingtalk
DingTalk Go SDK

## Contents
- [Installation](#Installation)
- [Quick start](#quick-start)
- [ToDo](#todo)

## Installation
```shell
go get -u github.com/zdz1715/go-dingtalk@latest
```

## Quick start
```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zdz1715/go-dingtalk"
)

func main() {
	credential := &dingtalk.InternalAppCredential{
		AppKey:    "YourAppKey",
		AppSecret: "YourAppSecret",
	}

	client, err := dingtalk.NewClient(credential, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 获取部门列表
	departments, err := client.Contacts.ListDepartments(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("departments: %+v\n", departments)
}

```

## ToDo
> [!NOTE]
> 现在提供的方法不多，会逐渐完善，也欢迎来贡献代码，只需要编写参数结构体、响应结构体就可以很快添加一个方法，参考下方代码。
```go
// 响应
type Department struct {
	DeptId          int    `json:"dept_id"`
	ParentId        int    `json:"parent_id"`
	AutoAddUser     bool   `json:"auto_add_user"`
	CreateDeptGroup bool   `json:"create_dept_group"`
	Name            string `json:"name"`
}

// 请求参数
type ListDepartmentsOptions struct {
	DeptId   *int      `json:"dept_id,omitempty" query:"dept_id"`
	Language *Language `json:"language,omitempty" query:"language"`
}

// ListDepartments gets a list of departments
// API docs: https://open.dingtalk.com/document/orgapp/obtain-the-department-list-v2
func (s *ContactsService) ListDepartments(ctx context.Context, opts *ListDepartmentsOptions) ([]*Department, error) {
	const apiEndpoint = "https://oapi.dingtalk.com/topapi/v2/department/listsub"
	var respBody []*Department
	if err := s.client.InvokeByToken(ctx, http.MethodPost, apiEndpoint, opts, &respBody); err != nil {
		return nil, err
	}
	return respBody, nil
}
```