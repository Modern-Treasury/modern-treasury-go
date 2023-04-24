package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LedgerAccountPayout struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the ledger account payout.
	Description string `json:"description,required,nullable"`
	// The status of the ledger account payout. One of `processing`, `pending`,
	// `posted`, `archiving` or `archived`.
	Status LedgerAccountPayoutStatus `json:"status,required"`
	// The id of the payout ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	PayoutLedgerAccountID string `json:"payout_ledger_account_id,required" format:"uuid"`
	// The id of the funding ledger account that sends to or receives funds from the
	// payout ledger account.
	FundingLedgerAccountID string `json:"funding_ledger_account_id,required" format:"uuid"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account payout. The default value is the created_at
	// timestamp of the ledger account payout.
	EffectiveAtUpperBound string `json:"effective_at_upper_bound,required" format:"time"`
	// The ledger transaction that this payout is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// The amount of the ledger account payout.
	Amount int64 `json:"amount,required,nullable"`
	// The currency of the ledger account payout.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account payout.
	CurrencyExponent int64 `json:"currency_exponent,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     LedgerAccountPayoutJSON
}

type LedgerAccountPayoutJSON struct {
	ID                     pjson.Metadata
	Object                 pjson.Metadata
	LiveMode               pjson.Metadata
	CreatedAt              pjson.Metadata
	UpdatedAt              pjson.Metadata
	Description            pjson.Metadata
	Status                 pjson.Metadata
	PayoutLedgerAccountID  pjson.Metadata
	FundingLedgerAccountID pjson.Metadata
	EffectiveAtUpperBound  pjson.Metadata
	LedgerTransactionID    pjson.Metadata
	Amount                 pjson.Metadata
	Currency               pjson.Metadata
	CurrencyExponent       pjson.Metadata
	Metadata               pjson.Metadata
	Raw                    []byte
	Extras                 map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccountPayout using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerAccountPayout) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountPayoutStatus string

const (
	LedgerAccountPayoutStatusArchived   LedgerAccountPayoutStatus = "archived"
	LedgerAccountPayoutStatusArchiving  LedgerAccountPayoutStatus = "archiving"
	LedgerAccountPayoutStatusPending    LedgerAccountPayoutStatus = "pending"
	LedgerAccountPayoutStatusPosted     LedgerAccountPayoutStatus = "posted"
	LedgerAccountPayoutStatusProcessing LedgerAccountPayoutStatus = "processing"
)
