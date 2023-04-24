package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type TransactionService struct {
	Options []option.RequestOption
}

func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	return
}

// Get details on a single transaction.
func (r *TransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/transactions/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update a single transaction.
func (r *TransactionService) Update(ctx context.Context, id string, body *requests.TransactionUpdateParams, opts ...option.RequestOption) (res *responses.Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/transactions/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of all transactions.
func (r *TransactionService) List(ctx context.Context, query *requests.TransactionListParams, opts ...option.RequestOption) (res *responses.Page[responses.Transaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/transactions"
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

// Get a list of all transactions.
func (r *TransactionService) ListAutoPager(ctx context.Context, query *requests.TransactionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Transaction] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
