package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerTransactionService struct {
	Options  []option.RequestOption
	Versions *LedgerTransactionVersionService
}

func NewLedgerTransactionService(opts ...option.RequestOption) (r *LedgerTransactionService) {
	r = &LedgerTransactionService{}
	r.Options = opts
	r.Versions = NewLedgerTransactionVersionService(opts...)
	return
}

// Create a ledger transaction.
func (r *LedgerTransactionService) New(ctx context.Context, body *requests.LedgerTransactionNewParams, opts ...option.RequestOption) (res *responses.LedgerTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_transactions"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get details on a single ledger transaction.
func (r *LedgerTransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.LedgerTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_transactions/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update the details of a ledger transaction.
func (r *LedgerTransactionService) Update(ctx context.Context, id string, body *requests.LedgerTransactionUpdateParams, opts ...option.RequestOption) (res *responses.LedgerTransaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_transactions/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of ledger transactions.
func (r *LedgerTransactionService) List(ctx context.Context, query *requests.LedgerTransactionListParams, opts ...option.RequestOption) (res *responses.Page[responses.LedgerTransaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_transactions"
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

// Get a list of ledger transactions.
func (r *LedgerTransactionService) ListAutoPager(ctx context.Context, query *requests.LedgerTransactionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LedgerTransaction] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
