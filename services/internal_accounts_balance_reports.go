package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type InternalAccountBalanceReportService struct {
	Options []option.RequestOption
}

func NewInternalAccountBalanceReportService(opts ...option.RequestOption) (r *InternalAccountBalanceReportService) {
	r = &InternalAccountBalanceReportService{}
	r.Options = opts
	return
}

// Get a single balance report for a given internal account.
func (r *InternalAccountBalanceReportService) Get(ctx context.Context, internal_account_id string, id string, opts ...option.RequestOption) (res *responses.BalanceReport, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports/%s", internal_account_id, id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get all balance reports for a given internal account.
func (r *InternalAccountBalanceReportService) List(ctx context.Context, internal_account_id string, query *requests.BalanceReportListParams, opts ...option.RequestOption) (res *responses.Page[responses.BalanceReport], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports", internal_account_id)
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

// Get all balance reports for a given internal account.
func (r *InternalAccountBalanceReportService) ListAutoPager(ctx context.Context, internal_account_id string, query *requests.BalanceReportListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.BalanceReport] {
	return responses.NewPageAutoPager(r.List(ctx, internal_account_id, query, opts...))
}
