package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerService struct {
	Options []option.RequestOption
}

func NewLedgerService(opts ...option.RequestOption) (r *LedgerService) {
	r = &LedgerService{}
	r.Options = opts
	return
}

// Create a ledger.
func (r *LedgerService) New(ctx context.Context, body *requests.LedgerNewParams, opts ...option.RequestOption) (res *responses.Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledgers"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get details on a single ledger.
func (r *LedgerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgers/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update the details of a ledger.
func (r *LedgerService) Update(ctx context.Context, id string, body *requests.LedgerUpdateParams, opts ...option.RequestOption) (res *responses.Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgers/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of ledgers.
func (r *LedgerService) List(ctx context.Context, query *requests.LedgerListParams, opts ...option.RequestOption) (res *responses.Page[responses.Ledger], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledgers"
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

// Get a list of ledgers.
func (r *LedgerService) ListAutoPager(ctx context.Context, query *requests.LedgerListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Ledger] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger.
func (r *LedgerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgers/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
