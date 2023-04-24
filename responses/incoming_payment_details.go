package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type IncomingPaymentDetail struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the Internal Account for the incoming payment detail. This is always
	// present.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// If the incoming payment detail is in a virtual account, the ID of the Virtual
	// Account.
	VirtualAccountID string `json:"virtual_account_id,required,nullable" format:"uuid"`
	// If the incoming payment detail is in a virtual account, the serialized virtual
	// account object.
	VirtualAccount VirtualAccount `json:"virtual_account,required,nullable"`
	// The ID of the reconciled Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The ID of the reconciled Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
	// `wire`.
	Type IncomingPaymentDetailType `json:"type,required"`
	// The raw data from the payment pre-notification file that we get from the bank.
	Data interface{} `json:"data,required"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// The currency of the incoming payment detail.
	Currency Currency `json:"currency,required,nullable"`
	// One of `credit` or `debit`.
	Direction IncomingPaymentDetailDirection `json:"direction,required"`
	// The current status of the incoming payment order. One of `pending`, `completed`,
	// or `returned`.
	Status IncomingPaymentDetailStatus `json:"status,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The date on which the corresponding transaction will occur.
	AsOfDate time.Time `json:"as_of_date,required" format:"date"`
	// The identifier of the vendor bank.
	VendorID string `json:"vendor_id,required,nullable" format:"uuid"`
	JSON     IncomingPaymentDetailJSON
}

type IncomingPaymentDetailJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	InternalAccountID     pjson.Metadata
	VirtualAccountID      pjson.Metadata
	VirtualAccount        pjson.Metadata
	TransactionLineItemID pjson.Metadata
	TransactionID         pjson.Metadata
	Type                  pjson.Metadata
	Data                  pjson.Metadata
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	Direction             pjson.Metadata
	Status                pjson.Metadata
	Metadata              pjson.Metadata
	AsOfDate              pjson.Metadata
	VendorID              pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into IncomingPaymentDetail using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *IncomingPaymentDetail) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type IncomingPaymentDetailType string

const (
	IncomingPaymentDetailTypeACH     IncomingPaymentDetailType = "ach"
	IncomingPaymentDetailTypeBook    IncomingPaymentDetailType = "book"
	IncomingPaymentDetailTypeCheck   IncomingPaymentDetailType = "check"
	IncomingPaymentDetailTypeEft     IncomingPaymentDetailType = "eft"
	IncomingPaymentDetailTypeInterac IncomingPaymentDetailType = "interac"
	IncomingPaymentDetailTypeRtp     IncomingPaymentDetailType = "rtp"
	IncomingPaymentDetailTypeSepa    IncomingPaymentDetailType = "sepa"
	IncomingPaymentDetailTypeSignet  IncomingPaymentDetailType = "signet"
	IncomingPaymentDetailTypeWire    IncomingPaymentDetailType = "wire"
)

type IncomingPaymentDetailDirection string

const (
	IncomingPaymentDetailDirectionCredit IncomingPaymentDetailDirection = "credit"
	IncomingPaymentDetailDirectionDebit  IncomingPaymentDetailDirection = "debit"
)

type IncomingPaymentDetailStatus string

const (
	IncomingPaymentDetailStatusCompleted IncomingPaymentDetailStatus = "completed"
	IncomingPaymentDetailStatusPending   IncomingPaymentDetailStatus = "pending"
	IncomingPaymentDetailStatusReturned  IncomingPaymentDetailStatus = "returned"
)
