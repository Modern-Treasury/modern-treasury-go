package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type AccountCollectionFlowService struct {
	Options []option.RequestOption
}

func NewAccountCollectionFlowService(opts ...option.RequestOption) (r *AccountCollectionFlowService) {
	r = &AccountCollectionFlowService{}
	r.Options = opts
	return
}

// create account_collection_flow
func (r *AccountCollectionFlowService) New(ctx context.Context, body *requests.AccountCollectionFlowNewParams, opts ...option.RequestOption) (res *responses.AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/account_collection_flows"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get account_collection_flow
func (r *AccountCollectionFlowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update account_collection_flow
func (r *AccountCollectionFlowService) Update(ctx context.Context, id string, body *requests.AccountCollectionFlowUpdateParams, opts ...option.RequestOption) (res *responses.AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list account_collection_flows
func (r *AccountCollectionFlowService) List(ctx context.Context, query *requests.AccountCollectionFlowListParams, opts ...option.RequestOption) (res *responses.Page[responses.AccountConnectionFlow], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/account_collection_flows"
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

// list account_collection_flows
func (r *AccountCollectionFlowService) ListAutoPager(ctx context.Context, query *requests.AccountCollectionFlowListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.AccountConnectionFlow] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
