package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LineItemService struct {
	Options []option.RequestOption
}

func NewLineItemService(opts ...option.RequestOption) (r *LineItemService) {
	r = &LineItemService{}
	r.Options = opts
	return
}

// Get a single line item
func (r *LineItemService) Get(ctx context.Context, itemizable_type requests.LineItemGetParamsItemizableType, itemizable_id string, id string, opts ...option.RequestOption) (res *responses.LineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/line_items/%s", itemizable_type, itemizable_id, id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update line item
func (r *LineItemService) Update(ctx context.Context, itemizable_type requests.LineItemUpdateParamsItemizableType, itemizable_id string, id string, body *requests.LineItemUpdateParams, opts ...option.RequestOption) (res *responses.LineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/line_items/%s", itemizable_type, itemizable_id, id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of line items
func (r *LineItemService) List(ctx context.Context, itemizable_type requests.LineItemListParamsItemizableType, itemizable_id string, query *requests.LineItemListParams, opts ...option.RequestOption) (res *responses.Page[responses.LineItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/line_items", itemizable_type, itemizable_id)
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

// Get a list of line items
func (r *LineItemService) ListAutoPager(ctx context.Context, itemizable_type requests.LineItemListParamsItemizableType, itemizable_id string, query *requests.LineItemListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LineItem] {
	return responses.NewPageAutoPager(r.List(ctx, itemizable_type, itemizable_id, query, opts...))
}
