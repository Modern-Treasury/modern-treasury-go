package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type PaymentReferenceService struct {
	Options []option.RequestOption
}

func NewPaymentReferenceService(opts ...option.RequestOption) (r *PaymentReferenceService) {
	r = &PaymentReferenceService{}
	r.Options = opts
	return
}

// list payment_references
func (r *PaymentReferenceService) List(ctx context.Context, query *requests.PaymentReferenceListParams, opts ...option.RequestOption) (res *responses.Page[responses.PaymentReference], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_references"
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

// list payment_references
func (r *PaymentReferenceService) ListAutoPager(ctx context.Context, query *requests.PaymentReferenceListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.PaymentReference] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// get payment_reference
func (r *PaymentReferenceService) Retireve(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.PaymentReference, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_references/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
