package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type RoutingDetailService struct {
	Options []option.RequestOption
}

func NewRoutingDetailService(opts ...option.RequestOption) (r *RoutingDetailService) {
	r = &RoutingDetailService{}
	r.Options = opts
	return
}

// Create a routing detail for a single external account.
func (r *RoutingDetailService) New(ctx context.Context, accounts_type requests.RoutingDetailNewParamsAccountsType, account_id string, body *requests.RoutingDetailNewParams, opts ...option.RequestOption) (res *responses.RoutingDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details", accounts_type, account_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get a single routing detail for a single internal or external account.
func (r *RoutingDetailService) Get(ctx context.Context, accounts_type requests.AccountsType, account_id string, id string, opts ...option.RequestOption) (res *responses.RoutingDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details/%s", accounts_type, account_id, id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) List(ctx context.Context, accounts_type requests.AccountsType, account_id string, query *requests.RoutingDetailListParams, opts ...option.RequestOption) (res *responses.Page[responses.RoutingDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details", accounts_type, account_id)
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

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) ListAutoPager(ctx context.Context, accounts_type requests.AccountsType, account_id string, query *requests.RoutingDetailListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.RoutingDetail] {
	return responses.NewPageAutoPager(r.List(ctx, accounts_type, account_id, query, opts...))
}

// Delete a routing detail for a single external account.
func (r *RoutingDetailService) Delete(ctx context.Context, accounts_type requests.RoutingDetailDeleteParamsAccountsType, account_id string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details/%s", accounts_type, account_id, id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}
