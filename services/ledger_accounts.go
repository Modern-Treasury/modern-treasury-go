package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerAccountService struct {
	Options []option.RequestOption
}

func NewLedgerAccountService(opts ...option.RequestOption) (r *LedgerAccountService) {
	r = &LedgerAccountService{}
	r.Options = opts
	return
}

// Create a ledger account.
func (r *LedgerAccountService) New(ctx context.Context, body *requests.LedgerAccountNewParams, opts ...option.RequestOption) (res *responses.LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_accounts"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get details on a single ledger account.
func (r *LedgerAccountService) Get(ctx context.Context, id string, query *requests.LedgerAccountGetParams, opts ...option.RequestOption) (res *responses.LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}

// Update the details of a ledger account.
func (r *LedgerAccountService) Update(ctx context.Context, id string, body *requests.LedgerAccountUpdateParams, opts ...option.RequestOption) (res *responses.LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of ledger accounts.
func (r *LedgerAccountService) List(ctx context.Context, query *requests.LedgerAccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.LedgerAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_accounts"
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

// Get a list of ledger accounts.
func (r *LedgerAccountService) ListAutoPager(ctx context.Context, query *requests.LedgerAccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LedgerAccount] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger account.
func (r *LedgerAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
