package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type AccountConnectionFlow struct {
	ID     string `json:"id" format:"uuid"`
	Object string `json:"object"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The client token of the account collection flow. This token can be used to embed
	// account collection in your client-side application.
	ClientToken string `json:"client_token"`
	// The current status of the account collection flow. One of `pending`,
	// `completed`, `expired`, or `cancelled`.
	Status AccountConnectionFlowStatus `json:"status"`
	// The ID of a counterparty. An external account created with this flow will be
	// associated with this counterparty.
	CounterpartyID string `json:"counterparty_id,required" format:"uuid"`
	// If present, the ID of the external account created using this flow.
	ExternalAccountID string                              `json:"external_account_id,nullable" format:"uuid"`
	PaymentTypes      []AccountConnectionFlowPaymentTypes `json:"payment_types,required"`
	JSON              AccountConnectionFlowJSON
}

type AccountConnectionFlowJSON struct {
	ID                pjson.Metadata
	Object            pjson.Metadata
	LiveMode          pjson.Metadata
	CreatedAt         pjson.Metadata
	UpdatedAt         pjson.Metadata
	ClientToken       pjson.Metadata
	Status            pjson.Metadata
	CounterpartyID    pjson.Metadata
	ExternalAccountID pjson.Metadata
	PaymentTypes      pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountConnectionFlow using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountConnectionFlow) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountConnectionFlowStatus string

const (
	AccountConnectionFlowStatusCancelled AccountConnectionFlowStatus = "cancelled"
	AccountConnectionFlowStatusCompleted AccountConnectionFlowStatus = "completed"
	AccountConnectionFlowStatusExpired   AccountConnectionFlowStatus = "expired"
	AccountConnectionFlowStatusPending   AccountConnectionFlowStatus = "pending"
)

type AccountConnectionFlowPaymentTypes string

const (
	AccountConnectionFlowPaymentTypesACH  AccountConnectionFlowPaymentTypes = "ach"
	AccountConnectionFlowPaymentTypesWire AccountConnectionFlowPaymentTypes = "wire"
)
