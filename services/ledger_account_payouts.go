package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerAccountPayoutService struct {
	Options []option.RequestOption
}

func NewLedgerAccountPayoutService(opts ...option.RequestOption) (r *LedgerAccountPayoutService) {
	r = &LedgerAccountPayoutService{}
	r.Options = opts
	return
}

// Create a ledger account payout.
func (r *LedgerAccountPayoutService) New(ctx context.Context, body *requests.LedgerAccountPayoutNewParams, opts ...option.RequestOption) (res *responses.LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_payouts"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Update the details of a ledger account payout.
func (r *LedgerAccountPayoutService) Update(ctx context.Context, id string, body *requests.LedgerAccountPayoutUpdateParams, opts ...option.RequestOption) (res *responses.LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_payouts/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of ledger account payouts.
func (r *LedgerAccountPayoutService) List(ctx context.Context, query *requests.LedgerAccountPayoutListParams, opts ...option.RequestOption) (res *responses.Page[responses.LedgerAccountPayout], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_account_payouts"
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

// Get a list of ledger account payouts.
func (r *LedgerAccountPayoutService) ListAutoPager(ctx context.Context, query *requests.LedgerAccountPayoutListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LedgerAccountPayout] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Get details on a single ledger account payout.
func (r *LedgerAccountPayoutService) Retireve(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_payouts/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
