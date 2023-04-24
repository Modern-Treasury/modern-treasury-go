package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerTransactionVersionService struct {
	Options []option.RequestOption
}

func NewLedgerTransactionVersionService(opts ...option.RequestOption) (r *LedgerTransactionVersionService) {
	r = &LedgerTransactionVersionService{}
	r.Options = opts
	return
}

// Get a list of ledger transaction versions.
func (r *LedgerTransactionVersionService) List(ctx context.Context, id string, query *requests.LedgerTransactionVersionListParams, opts ...option.RequestOption) (res *responses.Page[responses.LedgerTransactionVersion], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/ledger_transactions/%s/versions", id)
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

// Get a list of ledger transaction versions.
func (r *LedgerTransactionVersionService) ListAutoPager(ctx context.Context, id string, query *requests.LedgerTransactionVersionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LedgerTransactionVersion] {
	return responses.NewPageAutoPager(r.List(ctx, id, query, opts...))
}
