package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type ReturnService struct {
	Options []option.RequestOption
}

func NewReturnService(opts ...option.RequestOption) (r *ReturnService) {
	r = &ReturnService{}
	r.Options = opts
	return
}

// Create a return.
func (r *ReturnService) New(ctx context.Context, body *requests.ReturnNewParams, opts ...option.RequestOption) (res *responses.ReturnObject, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/returns"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get a single return.
func (r *ReturnService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.ReturnObject, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/returns/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Get a list of returns.
func (r *ReturnService) List(ctx context.Context, query *requests.ReturnListParams, opts ...option.RequestOption) (res *responses.Page[responses.ReturnObject], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/returns"
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

// Get a list of returns.
func (r *ReturnService) ListAutoPager(ctx context.Context, query *requests.ReturnListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.ReturnObject] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
