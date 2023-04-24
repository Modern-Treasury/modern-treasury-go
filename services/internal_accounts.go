package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type InternalAccountService struct {
	Options        []option.RequestOption
	BalanceReports *InternalAccountBalanceReportService
}

func NewInternalAccountService(opts ...option.RequestOption) (r *InternalAccountService) {
	r = &InternalAccountService{}
	r.Options = opts
	r.BalanceReports = NewInternalAccountBalanceReportService(opts...)
	return
}

// create internal account
func (r *InternalAccountService) New(ctx context.Context, body *requests.InternalAccountNewParams, opts ...option.RequestOption) (res *responses.InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/internal_accounts"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get internal account
func (r *InternalAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update internal account
func (r *InternalAccountService) Update(ctx context.Context, id string, body *requests.InternalAccountUpdateParams, opts ...option.RequestOption) (res *responses.InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list internal accounts
func (r *InternalAccountService) List(ctx context.Context, query *requests.InternalAccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.InternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/internal_accounts"
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

// list internal accounts
func (r *InternalAccountService) ListAutoPager(ctx context.Context, query *requests.InternalAccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.InternalAccount] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
