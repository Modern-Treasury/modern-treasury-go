package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type VirtualAccountService struct {
	Options []option.RequestOption
}

func NewVirtualAccountService(opts ...option.RequestOption) (r *VirtualAccountService) {
	r = &VirtualAccountService{}
	r.Options = opts
	return
}

// create virtual_account
func (r *VirtualAccountService) New(ctx context.Context, body *requests.VirtualAccountNewParams, opts ...option.RequestOption) (res *responses.VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/virtual_accounts"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get virtual_account
func (r *VirtualAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, nil, opts...)
	return
}

// update virtual_account
func (r *VirtualAccountService) Update(ctx context.Context, id string, body *requests.VirtualAccountUpdateParams, opts ...option.RequestOption) (res *responses.VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of virtual accounts.
func (r *VirtualAccountService) List(ctx context.Context, query *requests.VirtualAccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.VirtualAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/virtual_accounts"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of virtual accounts.
func (r *VirtualAccountService) ListAutoPager(ctx context.Context, query *requests.VirtualAccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.VirtualAccount] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete virtual_account
func (r *VirtualAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
