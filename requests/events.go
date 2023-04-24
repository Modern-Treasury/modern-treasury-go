package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

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
	return query.Marshal(r)
}
