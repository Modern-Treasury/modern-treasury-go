package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type InvoiceLineItemService struct {
	Options []option.RequestOption
}

func NewInvoiceLineItemService(opts ...option.RequestOption) (r *InvoiceLineItemService) {
	r = &InvoiceLineItemService{}
	r.Options = opts
	return
}

// create invoice_line_item
func (r *InvoiceLineItemService) New(ctx context.Context, invoice_id string, body *requests.InvoiceLineItemNewParams, opts ...option.RequestOption) (res *responses.InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items", invoice_id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// get invoice_line_item
func (r *InvoiceLineItemService) Get(ctx context.Context, invoice_id string, id string, opts ...option.RequestOption) (res *responses.InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoice_id, id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update invoice_line_item
func (r *InvoiceLineItemService) Update(ctx context.Context, invoice_id string, id string, body *requests.InvoiceLineItemUpdateParams, opts ...option.RequestOption) (res *responses.InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoice_id, id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list invoice_line_items
func (r *InvoiceLineItemService) List(ctx context.Context, invoice_id string, query *requests.InvoiceLineItemListParams, opts ...option.RequestOption) (res *responses.Page[responses.InvoiceLineItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items", invoice_id)
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

// list invoice_line_items
func (r *InvoiceLineItemService) ListAutoPager(ctx context.Context, invoice_id string, query *requests.InvoiceLineItemListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.InvoiceLineItem] {
	return responses.NewPageAutoPager(r.List(ctx, invoice_id, query, opts...))
}

// delete invoice_line_item
func (r *InvoiceLineItemService) Delete(ctx context.Context, invoice_id string, id string, opts ...option.RequestOption) (res *responses.InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoice_id, id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
