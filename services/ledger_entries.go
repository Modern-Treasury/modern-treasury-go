package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerEntryService struct {
	Options []option.RequestOption
}

func NewLedgerEntryService(opts ...option.RequestOption) (r *LedgerEntryService) {
	r = &LedgerEntryService{}
	r.Options = opts
	return
}

// Get details on a single ledger entry.
func (r *LedgerEntryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.LedgerEntry, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_entries/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of all ledger entries.
func (r *LedgerEntryService) List(ctx context.Context, query *requests.LedgerEntryListParams, opts ...option.RequestOption) (res *responses.Page[responses.LedgerEntry], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_entries"
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

// Get a list of all ledger entries.
func (r *LedgerEntryService) ListAutoPager(ctx context.Context, query *requests.LedgerEntryListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LedgerEntry] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
