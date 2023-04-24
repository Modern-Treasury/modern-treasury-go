package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type IncomingPaymentDetailService struct {
	Options []option.RequestOption
}

func NewIncomingPaymentDetailService(opts ...option.RequestOption) (r *IncomingPaymentDetailService) {
	r = &IncomingPaymentDetailService{}
	r.Options = opts
	return
}

// Get an existing Incoming Payment Detail.
func (r *IncomingPaymentDetailService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.IncomingPaymentDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/incoming_payment_details/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Update an existing Incoming Payment Detail.
func (r *IncomingPaymentDetailService) Update(ctx context.Context, id string, body *requests.IncomingPaymentDetailUpdateParams, opts ...option.RequestOption) (res *responses.IncomingPaymentDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/incoming_payment_details/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of Incoming Payment Details.
func (r *IncomingPaymentDetailService) List(ctx context.Context, query *requests.IncomingPaymentDetailListParams, opts ...option.RequestOption) (res *responses.Page[responses.IncomingPaymentDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/incoming_payment_details"
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

// Get a list of Incoming Payment Details.
func (r *IncomingPaymentDetailService) ListAutoPager(ctx context.Context, query *requests.IncomingPaymentDetailListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.IncomingPaymentDetail] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Simulate Incoming Payment Detail
func (r *IncomingPaymentDetailService) NewAsync(ctx context.Context, body *requests.IncomingPaymentDetailNewAsyncParams, opts ...option.RequestOption) (res *responses.AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/simulations/incoming_payment_details/create_async"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
