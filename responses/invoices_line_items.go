package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type InvoiceLineItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The name of the line item, typically a product or SKU name.
	Name string `json:"name,required"`
	// An optional free-form description of the line item.
	Description string `json:"description,required"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity int64 `json:"quantity,required"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit.
	UnitAmount int64 `json:"unit_amount,required"`
	// Either `debit` or `credit`. `debit` indicates that a client owes the business
	// money and increases the invoice's `total_amount` due. `credit` has the opposite
	// intention and effect.
	Direction string `json:"direction,required"`
	// The total amount for this line item specified in the invoice currency's smallest
	// unit.
	Amount int64 `json:"amount,required"`
	JSON   InvoiceLineItemJSON
}

type InvoiceLineItemJSON struct {
	ID          pjson.Metadata
	Object      pjson.Metadata
	LiveMode    pjson.Metadata
	CreatedAt   pjson.Metadata
	UpdatedAt   pjson.Metadata
	Name        pjson.Metadata
	Description pjson.Metadata
	Quantity    pjson.Metadata
	UnitAmount  pjson.Metadata
	Direction   pjson.Metadata
	Amount      pjson.Metadata
	Raw         []byte
	Extras      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InvoiceLineItem using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
