// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// EventService contains methods and other services that help with interacting with
// the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// get event
func (r *EventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Event, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/events/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// list events
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *shared.Page[Event], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *shared.PageAutoPager[Event] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Event struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The type of resource for the event.
	Resource string `json:"resource,required"`
	// The name of the event.
	EventName string `json:"event_name,required"`
	// The time of the event.
	EventTime time.Time `json:"event_time,required" format:"date-time"`
	// The body of the event.
	Data map[string]interface{} `json:"data,required"`
	// The ID of the entity for the event.
	EntityID string `json:"entity_id,required"`
	JSON     eventJSON
}

// eventJSON contains the JSON metadata for the struct [Event]
type eventJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	LiveMode    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	Resource    apijson.Field
	EventName   apijson.Field
	EventTime   apijson.Field
	Data        apijson.Field
	EntityID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	EntityID    param.Field[string] `query:"entity_id"`
	EventName   param.Field[string] `query:"event_name"`
	// An inclusive upper bound for when the event occurred
	EventTimeEnd param.Field[time.Time] `query:"event_time_end" format:"date-time"`
	// An inclusive lower bound for when the event occurred
	EventTimeStart param.Field[time.Time] `query:"event_time_start" format:"date-time"`
	PerPage        param.Field[int64]     `query:"per_page"`
	Resource       param.Field[string]    `query:"resource"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
