package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zdz1715/go-dingtalk"
)

func main() {
	credential := dingtalk.InternalAppCredential{
		AppKey:    "",
		AppSecret: "",
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
