package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LineItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the payment order or expected payment.
	ItemizableID string `json:"itemizable_id,required" format:"uuid"`
	// One of `payment_orders` or `expected_payments`.
	ItemizableType LineItemItemizableType `json:"itemizable_type,required"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// A free-form description of the line item.
	Description string `json:"description,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata   map[string]string  `json:"metadata,required"`
	Accounting LineItemAccounting `json:"accounting,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID string `json:"accounting_category_id,required,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	AccountingLedgerClassID string `json:"accounting_ledger_class_id,required,nullable" format:"uuid"`
	JSON                    LineItemJSON
}

type LineItemJSON struct {
	ID                      pjson.Metadata
	Object                  pjson.Metadata
	LiveMode                pjson.Metadata
	CreatedAt               pjson.Metadata
	UpdatedAt               pjson.Metadata
	ItemizableID            pjson.Metadata
	ItemizableType          pjson.Metadata
	Amount                  pjson.Metadata
	Description             pjson.Metadata
	Metadata                pjson.Metadata
	Accounting              pjson.Metadata
	AccountingCategoryID    pjson.Metadata
	AccountingLedgerClassID pjson.Metadata
	Raw                     []byte
	Extras                  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LineItem using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LineItem) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LineItemItemizableType string

const (
	LineItemItemizableTypeExpectedPayment LineItemItemizableType = "ExpectedPayment"
	LineItemItemizableTypePaymentOrder    LineItemItemizableType = "PaymentOrder"
)

type LineItemAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID string `json:"class_id,nullable" format:"uuid"`
	JSON    LineItemAccountingJSON
}

type LineItemAccountingJSON struct {
	AccountID pjson.Metadata
	ClassID   pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LineItemAccounting using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LineItemAccounting) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
