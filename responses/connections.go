package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type Connection struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Unique identifier for the bank or vendor.
	VendorID string `json:"vendor_id,required" format:"uuid"`
	// An identifier given to this connection by the bank.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable" format:"uuid"`
	// A human-friendly name for the bank or vendor.
	VendorName string `json:"vendor_name,required"`
	JSON       ConnectionJSON
}

type ConnectionJSON struct {
	ID               pjson.Metadata
	Object           pjson.Metadata
	LiveMode         pjson.Metadata
	CreatedAt        pjson.Metadata
	UpdatedAt        pjson.Metadata
	DiscardedAt      pjson.Metadata
	VendorID         pjson.Metadata
	VendorCustomerID pjson.Metadata
	VendorName       pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Connection using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
