package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type PaymentFlow struct {
	ID     string `json:"id" format:"uuid"`
	Object string `json:"object"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The client token of the payment flow. This token can be used to embed a payment
	// workflow in your client-side application.
	ClientToken string `json:"client_token"`
	// The current status of the payment flow. One of `pending`, `completed`,
	// `expired`, or `cancelled`.
	Status PaymentFlowStatus `json:"status"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount int64 `json:"amount"`
	// The currency of the payment.
	Currency string `json:"currency"`
	// Describes the direction money is flowing in the transaction. Can only be
	// `debit`. A `debit` pulls money from someone else's account to your own.
	Direction PaymentFlowDirection `json:"direction"`
	// The ID of a counterparty associated with the payment. As part of the payment
	// workflow an external account will be associated with this counterparty.
	CounterpartyID string `json:"counterparty_id,nullable" format:"uuid"`
	// If present, the ID of the external account created using this flow.
	ReceivingAccountID string `json:"receiving_account_id,nullable" format:"uuid"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID string `json:"originating_account_id,nullable" format:"uuid"`
	// If present, the ID of the payment order created using this flow.
	PaymentOrderID string `json:"payment_order_id,nullable" format:"uuid"`
	JSON           PaymentFlowJSON
}

type PaymentFlowJSON struct {
	ID                   pjson.Metadata
	Object               pjson.Metadata
	LiveMode             pjson.Metadata
	CreatedAt            pjson.Metadata
	UpdatedAt            pjson.Metadata
	ClientToken          pjson.Metadata
	Status               pjson.Metadata
	Amount               pjson.Metadata
	Currency             pjson.Metadata
	Direction            pjson.Metadata
	CounterpartyID       pjson.Metadata
	ReceivingAccountID   pjson.Metadata
	OriginatingAccountID pjson.Metadata
	PaymentOrderID       pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PaymentFlow using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PaymentFlow) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PaymentFlowStatus string

const (
	PaymentFlowStatusCancelled PaymentFlowStatus = "cancelled"
	PaymentFlowStatusCompleted PaymentFlowStatus = "completed"
	PaymentFlowStatusExpired   PaymentFlowStatus = "expired"
	PaymentFlowStatusPending   PaymentFlowStatus = "pending"
)

type PaymentFlowDirection string

const (
	PaymentFlowDirectionCredit PaymentFlowDirection = "credit"
	PaymentFlowDirectionDebit  PaymentFlowDirection = "debit"
)
