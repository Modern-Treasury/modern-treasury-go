package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type PaymentOrderService struct {
	Options   []option.RequestOption
	Reversals *PaymentOrderReversalService
}

func NewPaymentOrderService(opts ...option.RequestOption) (r *PaymentOrderService) {
	r = &PaymentOrderService{}
	r.Options = opts
	r.Reversals = NewPaymentOrderReversalService(opts...)
	return
}

// Create a new Payment Order
func (r *PaymentOrderService) New(ctx context.Context, body *requests.PaymentOrderNewParams, opts ...option.RequestOption) (res *responses.PaymentOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/payment_orders"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get details on a single payment order
func (r *PaymentOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.PaymentOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update a payment order
func (r *PaymentOrderService) Update(ctx context.Context, id string, body *requests.PaymentOrderUpdateParams, opts ...option.RequestOption) (res *responses.PaymentOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of all payment orders
func (r *PaymentOrderService) List(ctx context.Context, query *requests.PaymentOrderListParams, opts ...option.RequestOption) (res *responses.Page[responses.PaymentOrder], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_orders"
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

// Get a list of all payment orders
func (r *PaymentOrderService) ListAutoPager(ctx context.Context, query *requests.PaymentOrderListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.PaymentOrder] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create a new payment order asynchronously
func (r *PaymentOrderService) NewAsync(ctx context.Context, body *requests.PaymentOrderNewAsyncParams, opts ...option.RequestOption) (res *responses.AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/payment_orders/create_async"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
