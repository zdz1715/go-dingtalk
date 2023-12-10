package dingtalk

import (
	"context"
	"net/http"
)

type ContactsService service

type Department struct {
	DeptId          int    `json:"dept_id"`
	ParentId        int    `json:"parent_id"`
	AutoAddUser     bool   `json:"auto_add_user"`
	CreateDeptGroup bool   `json:"create_dept_group"`
	Name            string `json:"name"`
}

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

type DepartmentV1 struct {
	DeptId          int    `json:"id"`
	ParentId        int    `json:"parentid"`
	AutoAddUser     bool   `json:"autoAddUser"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	Name            string `json:"name"`
}

type DepartmentsResult struct {
	Result
	Department []*DepartmentV1 `json:"department"`
}

type ListDepartmentsV1Options struct {
	ID         string    `json:"id,omitempty" query:"id"`
	FetchChild bool      `json:"fetch_child,omitempty" query:"fetch_child"`
	Lang       *Language `json:"lang,omitempty" query:"lang"`
}

// ListDepartmentsV1 gets a list of all departments
// API docs: https://open.dingtalk.com/document/orgapp/obtain-the-department-list#h2-ys4-p08-78f
func (s *ContactsService) ListDepartmentsV1(ctx context.Context, opts *ListDepartmentsV1Options) ([]*DepartmentV1, error) {
	const apiEndpoint = "https://oapi.dingtalk.com/department/list"
	var respBody DepartmentsResult
	if err := s.client.InvokeByToken(ctx, http.MethodGet, apiEndpoint, opts, &respBody); err != nil {
		return nil, err
	}
	if err := respBody.Error(); err != nil {
		return nil, err
	}
	return respBody.Department, nil
}