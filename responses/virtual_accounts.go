package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type VirtualAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the virtual account.
	Name string `json:"name,required"`
	// An optional free-form description for internal use.
	Description string `json:"description,required,nullable"`
	// The ID of a counterparty that the virtual account belongs to. Optional.
	CounterpartyID string `json:"counterparty_id,required,nullable" format:"uuid"`
	// The ID of the internal account that the virtual account is in.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// An array of account detail objects.
	AccountDetails []AccountDetail `json:"account_details,required"`
	// An array of routing detail objects. These will be the routing details of the
	// internal account.
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID string `json:"debit_ledger_account_id,required,nullable" format:"uuid"`
	// The ID of a credit normal ledger account. When money enters the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID string `json:"credit_ledger_account_id,required,nullable" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     VirtualAccountJSON
}

type VirtualAccountJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	DiscardedAt           pjson.Metadata
	Name                  pjson.Metadata
	Description           pjson.Metadata
	CounterpartyID        pjson.Metadata
	InternalAccountID     pjson.Metadata
	AccountDetails        pjson.Metadata
	RoutingDetails        pjson.Metadata
	DebitLedgerAccountID  pjson.Metadata
	CreditLedgerAccountID pjson.Metadata
	Metadata              pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into VirtualAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *VirtualAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
