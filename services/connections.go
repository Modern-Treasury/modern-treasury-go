package services

import (
	"context"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type ConnectionService struct {
	Options []option.RequestOption
}

func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

// Get a list of all connections.
func (r *ConnectionService) List(ctx context.Context, query *requests.ConnectionListParams, opts ...option.RequestOption) (res *responses.Page[responses.Connection], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/connections"
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

// Get a list of all connections.
func (r *ConnectionService) ListAutoPager(ctx context.Context, query *requests.ConnectionListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Connection] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
