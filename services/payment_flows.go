package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type PaymentFlowService struct {
	Options []option.RequestOption
}

func NewPaymentFlowService(opts ...option.RequestOption) (r *PaymentFlowService) {
	r = &PaymentFlowService{}
	r.Options = opts
	return
}

// create payment_flow
func (r *PaymentFlowService) New(ctx context.Context, body *requests.PaymentFlowNewParams, opts ...option.RequestOption) (res *responses.PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/payment_flows"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get payment_flow
func (r *PaymentFlowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_flows/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update payment_flow
func (r *PaymentFlowService) Update(ctx context.Context, id string, body *requests.PaymentFlowUpdateParams, opts ...option.RequestOption) (res *responses.PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_flows/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list payment_flows
func (r *PaymentFlowService) List(ctx context.Context, query *requests.PaymentFlowListParams, opts ...option.RequestOption) (res *responses.Page[responses.PaymentFlow], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_flows"
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

// list payment_flows
func (r *PaymentFlowService) ListAutoPager(ctx context.Context, query *requests.PaymentFlowListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.PaymentFlow] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
