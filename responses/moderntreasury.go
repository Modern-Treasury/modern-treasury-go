package responses

import (
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type PingResponse struct {
	Ping string `json:"ping,required"`
	JSON PingResponseJSON
}

type PingResponseJSON struct {
	Ping   pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PingResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PingResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
