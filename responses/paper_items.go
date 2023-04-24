package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type PaperItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the reconciled Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The ID of the reconciled Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The current status of the paper item. One of `pending`, `completed`, or
	// `returned`.
	Status PaperItemStatus `json:"status,required"`
	// The identifier for the lockbox assigned by the bank.
	LockboxNumber string `json:"lockbox_number,required"`
	// The date the paper item was deposited into your organization's bank account.
	DepositDate time.Time `json:"deposit_date,required" format:"date"`
	// The amount of the paper item.
	Amount int64 `json:"amount,required"`
	// The currency of the paper item.
	Currency Currency `json:"currency,required,nullable"`
	// The account number on the paper item.
	AccountNumber string `json:"account_number,required,nullable"`
	// The last 4 digits of the account_number.
	AccountNumberSafe string `json:"account_number_safe,required,nullable"`
	// The routing number on the paper item.
	RoutingNumber string `json:"routing_number,required,nullable"`
	// The check number on the paper item.
	CheckNumber string `json:"check_number,required,nullable"`
	// The name of the remitter on the paper item.
	RemitterName string `json:"remitter_name,required,nullable"`
	// The memo field on the paper item.
	MemoField string `json:"memo_field,required,nullable"`
	JSON      PaperItemJSON
}

type PaperItemJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	TransactionLineItemID pjson.Metadata
	TransactionID         pjson.Metadata
	Status                pjson.Metadata
	LockboxNumber         pjson.Metadata
	DepositDate           pjson.Metadata
	Amount                pjson.Metadata
	Currency              pjson.Metadata
	AccountNumber         pjson.Metadata
	AccountNumberSafe     pjson.Metadata
	RoutingNumber         pjson.Metadata
	CheckNumber           pjson.Metadata
	RemitterName          pjson.Metadata
	MemoField             pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PaperItem using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PaperItem) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PaperItemStatus string

const (
	PaperItemStatusCompleted PaperItemStatus = "completed"
	PaperItemStatusPending   PaperItemStatus = "pending"
	PaperItemStatusReturned  PaperItemStatus = "returned"
)
