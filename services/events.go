package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type EventService struct {
	Options []option.RequestOption
}

func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// get event
func (r *EventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/events/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// list events
func (r *EventService) List(ctx context.Context, query *requests.EventListParams, opts ...option.RequestOption) (res *responses.Page[responses.Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/events"
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

// list events
func (r *EventService) ListAutoPager(ctx context.Context, query *requests.EventListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.Event] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}
