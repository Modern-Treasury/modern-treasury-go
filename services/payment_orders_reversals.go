package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type PaymentOrderReversalService struct {
	Options []option.RequestOption
}

func NewPaymentOrderReversalService(opts ...option.RequestOption) (r *PaymentOrderReversalService) {
	r = &PaymentOrderReversalService{}
	r.Options = opts
	return
}

// Create a reversal for a payment order.
func (r *PaymentOrderReversalService) New(ctx context.Context, payment_order_id string, body *requests.PaymentOrderReversalNewParams, opts ...option.RequestOption) (res *responses.Reversal, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s/reversals", payment_order_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get details on a single reversal of a payment order.
func (r *PaymentOrderReversalService) Get(ctx context.Context, payment_order_id string, reversal_id string, opts ...option.RequestOption) (res *responses.Reversal, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s/reversals/%s", payment_order_id, reversal_id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of all reversals of a payment order.
func (r *PaymentOrderReversalService) List(ctx context.Context, payment_order_id string, query *requests.PaymentOrderReversalListParams, opts ...option.RequestOption) (res *responses.Page[responses.Reversal], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/payment_orders/%s/reversals", payment_order_id)
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

// Get a list of all reversals of a payment order.
func (r *PaymentOrderReversalService) ListAutoPager(ctx context.Context, payment_order_id string, query *requests.PaymentOrderReversalListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Reversal] {
	return responses.NewPageAutoPager(r.List(ctx, payment_order_id, query, opts...))
}
