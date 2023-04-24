package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type CounterpartyService struct {
	Options []option.RequestOption
}

func NewCounterpartyService(opts ...option.RequestOption) (r *CounterpartyService) {
	r = &CounterpartyService{}
	r.Options = opts
	return
}

// Create a new counterparty.
func (r *CounterpartyService) New(ctx context.Context, body *requests.CounterpartyNewParams, opts ...option.RequestOption) (res *responses.Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/counterparties"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get details on a single counterparty.
func (r *CounterpartyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Updates a given counterparty with new information.
func (r *CounterpartyService) Update(ctx context.Context, id string, body *requests.CounterpartyUpdateParams, opts ...option.RequestOption) (res *responses.Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a paginated list of all counterparties.
func (r *CounterpartyService) List(ctx context.Context, query *requests.CounterpartyListParams, opts ...option.RequestOption) (res *responses.Page[responses.Counterparty], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/counterparties"
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

// Get a paginated list of all counterparties.
func (r *CounterpartyService) ListAutoPager(ctx context.Context, query *requests.CounterpartyListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Counterparty] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a given counterparty.
func (r *CounterpartyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}

// Send an email requesting account details.
func (r *CounterpartyService) CollectAccount(ctx context.Context, id string, body *requests.CounterpartyCollectAccountParams, opts ...option.RequestOption) (res *responses.CounterpartyCollectAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/counterparties/%s/collect_account", id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
