package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type ExpectedPaymentService struct {
	Options []option.RequestOption
}

func NewExpectedPaymentService(opts ...option.RequestOption) (r *ExpectedPaymentService) {
	r = &ExpectedPaymentService{}
	r.Options = opts
	return
}

// create expected payment
func (r *ExpectedPaymentService) New(ctx context.Context, body *requests.ExpectedPaymentNewParams, opts ...option.RequestOption) (res *responses.ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/expected_payments"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get expected payment
func (r *ExpectedPaymentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update expected payment
func (r *ExpectedPaymentService) Update(ctx context.Context, id string, body *requests.ExpectedPaymentUpdateParams, opts ...option.RequestOption) (res *responses.ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list expected_payments
func (r *ExpectedPaymentService) List(ctx context.Context, query *requests.ExpectedPaymentListParams, opts ...option.RequestOption) (res *responses.Page[responses.ExpectedPayment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/expected_payments"
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

// list expected_payments
func (r *ExpectedPaymentService) ListAutoPager(ctx context.Context, query *requests.ExpectedPaymentListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.ExpectedPayment] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete expected payment
func (r *ExpectedPaymentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
