package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type InvoiceService struct {
	Options   []option.RequestOption
	LineItems *InvoiceLineItemService
}

func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	r.LineItems = NewInvoiceLineItemService(opts...)
	return
}

// create invoice
func (r *InvoiceService) New(ctx context.Context, body *requests.InvoiceNewParams, opts ...option.RequestOption) (res *responses.Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/invoices"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get invoice
func (r *InvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update invoice
func (r *InvoiceService) Update(ctx context.Context, id string, body *requests.InvoiceUpdateParams, opts ...option.RequestOption) (res *responses.Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list invoices
func (r *InvoiceService) List(ctx context.Context, query *requests.InvoiceListParams, opts ...option.RequestOption) (res *responses.Page[responses.Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/invoices"
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

// list invoices
func (r *InvoiceService) ListAutoPager(ctx context.Context, query *requests.InvoiceListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Invoice] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
