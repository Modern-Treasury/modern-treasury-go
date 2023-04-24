package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type PaperItemService struct {
	Options []option.RequestOption
}

func NewPaperItemService(opts ...option.RequestOption) (r *PaperItemService) {
	r = &PaperItemService{}
	r.Options = opts
	return
}

// Get details on a single paper item.
func (r *PaperItemService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.PaperItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/paper_items/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of all paper items.
func (r *PaperItemService) List(ctx context.Context, query *requests.PaperItemListParams, opts ...option.RequestOption) (res *responses.Page[responses.PaperItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/paper_items"
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

// Get a list of all paper items.
func (r *PaperItemService) ListAutoPager(ctx context.Context, query *requests.PaperItemListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.PaperItem] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
