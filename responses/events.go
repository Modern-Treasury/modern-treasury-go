package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

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
	Data interface{} `json:"data,required"`
	// The ID of the entity for the event.
	EntityID string `json:"entity_id,required"`
	JSON     EventJSON
}

type EventJSON struct {
	ID        pjson.Metadata
	Object    pjson.Metadata
	LiveMode  pjson.Metadata
	CreatedAt pjson.Metadata
	UpdatedAt pjson.Metadata
	Resource  pjson.Metadata
	EventName pjson.Metadata
	EventTime pjson.Metadata
	Data      pjson.Metadata
	EntityID  pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Event using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Event) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
