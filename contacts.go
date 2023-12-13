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

type ListUsersOptions struct {
	DeptId             *int      `json:"dept_id,omitempty" query:"dept_id"`
	Cursor             *int      `json:"cursor,omitempty" query:"cursor"`
	Size               *int      `json:"size,omitempty" query:"size"`
	OrderField         *string   `json:"order_field,omitempty" query:"order_field"`
	ContainAccessLimit *bool     `json:"contain_access_limit,omitempty" query:"contain_access_limit"`
	Language           *Language `json:"language,omitempty" query:"language"`
}

type User struct {
	Userid           string   `json:"userid"`
	Unionid          string   `json:"unionid"`
	Name             string   `json:"name"`
	Avatar           string   `json:"avatar"`
	StateCode        string   `json:"state_code"`
	Mobile           string   `json:"mobile"`
	HideMobile       string   `json:"hide_mobile"`
	Telephone        string   `json:"telephone"`
	JobNumber        string   `json:"job_number"`
	Title            string   `json:"title"`
	Email            string   `json:"email"`
	OrgEmail         string   `json:"org_email"`
	Remark           string   `json:"remark"`
	DeptIDList       []string `json:"dept_id_list"`
	DeptOrder        int      `json:"dept_order"`
	Extension        string   `json:"extension"`
	HiredDate        int      `json:"hired_date"`
	Active           bool     `json:"active"`
	Admin            bool     `json:"admin"`
	Boss             bool     `json:"boss"`
	Leader           bool     `json:"leader"`
	ExclusiveAccount bool     `json:"exclusive_account"`
}

type UsersResult struct {
	HasMore    bool    `json:"has_more"`
	NextCursor int     `json:"next_cursor"`
	List       []*User `json:"list"`
}

// ListUsers gets a list of user detail
// API docs: https://open.dingtalk.com/document/orgapp/queries-the-complete-information-of-a-department-user
func (s *ContactsService) ListUsers(ctx context.Context, opts *ListUsersOptions) (*UsersResult, error) {
	const apiEndpoint = "https://oapi.dingtalk.com/topapi/v2/user/list"
	var respBody UsersResult
	if err := s.client.InvokeByToken(ctx, http.MethodPost, apiEndpoint, opts, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}

type DepartmentV1 struct {
	DeptId          int    `json:"id"`
	ParentId        int    `json:"parentid"`
	AutoAddUser     bool   `json:"autoAddUser"`
	CreateDeptGroup bool   `json:"createDeptGroup"`
	Name            string `json:"name"`
}

type DepartmentsResult struct {
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
	return respBody.Department, nil
}
