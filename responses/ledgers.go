package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type Ledger struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the ledger.
	Name string `json:"name,required"`
	// An optional free-form description for internal use.
	Description string `json:"description,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata    map[string]string      `json:"metadata,required"`
	ExtraFields map[string]interface{} `pjson:"-,extras"`
	JSON        LedgerJSON
}

type LedgerJSON struct {
	ID          pjson.Metadata
	Object      pjson.Metadata
	LiveMode    pjson.Metadata
	CreatedAt   pjson.Metadata
	UpdatedAt   pjson.Metadata
	DiscardedAt pjson.Metadata
	Name        pjson.Metadata
	Description pjson.Metadata
	Metadata    pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ledger using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Ledger) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
