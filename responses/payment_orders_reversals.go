package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type Reversal struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The current status of the reversal.
	Status ReversalStatus `json:"status,required"`
	// The ID of the relevant Payment Order.
	PaymentOrderID string `json:"payment_order_id,required,nullable" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The reason for the reversal.
	Reason ReversalReason `json:"reason,required"`
	JSON   ReversalJSON
}

type ReversalJSON struct {
	ID             pjson.Metadata
	Object         pjson.Metadata
	LiveMode       pjson.Metadata
	CreatedAt      pjson.Metadata
	UpdatedAt      pjson.Metadata
	Status         pjson.Metadata
	PaymentOrderID pjson.Metadata
	Metadata       pjson.Metadata
	Reason         pjson.Metadata
	Raw            []byte
	Extras         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Reversal using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Reversal) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ReversalStatus string

const (
	ReversalStatusCompleted  ReversalStatus = "completed"
	ReversalStatusFailed     ReversalStatus = "failed"
	ReversalStatusPending    ReversalStatus = "pending"
	ReversalStatusProcessing ReversalStatus = "processing"
	ReversalStatusReturned   ReversalStatus = "returned"
	ReversalStatusSent       ReversalStatus = "sent"
)

type ReversalReason string

const (
	ReversalReasonDuplicate                 ReversalReason = "duplicate"
	ReversalReasonIncorrectAmount           ReversalReason = "incorrect_amount"
	ReversalReasonIncorrectReceivingAccount ReversalReason = "incorrect_receiving_account"
	ReversalReasonDateEarlierThanIntended   ReversalReason = "date_earlier_than_intended"
	ReversalReasonDateLaterThanIntended     ReversalReason = "date_later_than_intended"
)
