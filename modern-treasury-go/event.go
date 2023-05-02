package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
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
	JSON     EventJSON
}

type EventJSON struct {
	ID        apijson.Metadata
	Object    apijson.Metadata
	LiveMode  apijson.Metadata
	CreatedAt apijson.Metadata
	UpdatedAt apijson.Metadata
	Resource  apijson.Metadata
	EventName apijson.Metadata
	EventTime apijson.Metadata
	Data      apijson.Metadata
	EntityID  apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Event using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// An inclusive lower bound for when the event occurred
	EventTimeStart field.Field[time.Time] `query:"event_time_start" format:"date-time"`
	// An inclusive upper bound for when the event occurred
	EventTimeEnd field.Field[time.Time] `query:"event_time_end" format:"date-time"`
	Resource     field.Field[string]    `query:"resource"`
	EntityID     field.Field[string]    `query:"entity_id"`
	EventName    field.Field[string]    `query:"event_name"`
}

// URLQuery serializes EventListParams into a url.Values of the query parameters
// associated with this value
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
