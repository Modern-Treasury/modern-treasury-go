package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type AccountDetailService struct {
	Options []option.RequestOption
}

func NewAccountDetailService(opts ...option.RequestOption) (r *AccountDetailService) {
	r = &AccountDetailService{}
	r.Options = opts
	return
}

// Create an account detail for an external account.
func (r *AccountDetailService) New(ctx context.Context, accounts_type requests.AccountDetailNewParamsAccountsType, account_id string, body *requests.AccountDetailNewParams, opts ...option.RequestOption) (res *responses.AccountDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/account_details", accounts_type, account_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get a single account detail for a single internal or external account.
func (r *AccountDetailService) Get(ctx context.Context, accounts_type requests.AccountsType, account_id string, id string, opts ...option.RequestOption) (res *responses.AccountDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/account_details/%s", accounts_type, account_id, id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of account details for a single internal or external account.
func (r *AccountDetailService) List(ctx context.Context, accounts_type requests.AccountsType, account_id string, query *requests.AccountDetailListParams, opts ...option.RequestOption) (res *responses.Page[responses.AccountDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/account_details", accounts_type, account_id)
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

// Get a list of account details for a single internal or external account.
func (r *AccountDetailService) ListAutoPager(ctx context.Context, accounts_type requests.AccountsType, account_id string, query *requests.AccountDetailListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.AccountDetail] {
	return responses.NewPageAutoPager(r.List(ctx, accounts_type, account_id, query, opts...))
}

// Delete a single account detail for an external account.
func (r *AccountDetailService) Delete(ctx context.Context, accounts_type requests.AccountDetailDeleteParamsAccountsType, account_id string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/%s/%s/account_details/%s", accounts_type, account_id, id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}
