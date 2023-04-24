package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LedgerAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the ledger account.
	Name string `json:"name,required"`
	// The description of the ledger account.
	Description string `json:"description,required,nullable"`
	// The normal balance of the ledger account.
	NormalBalance LedgerAccountNormalBalance `json:"normal_balance,required"`
	// The pending, posted, and available balances for this ledger account. The posted
	// balance is the sum of all posted entries on the account. The pending balance is
	// the sum of all pending and posted entries on the account. The available balance
	// is the posted incoming entries minus the sum of the pending and posted outgoing
	// amounts.
	Balances LedgerAccountBalances `json:"balances,required"`
	// Lock version of the ledger account.
	LockVersion int64 `json:"lock_version,required"`
	// The id of the ledger that this account belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType LedgerAccountLedgerableType `json:"ledgerable_type,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     LedgerAccountJSON
}

type LedgerAccountJSON struct {
	ID             pjson.Metadata
	Object         pjson.Metadata
	LiveMode       pjson.Metadata
	CreatedAt      pjson.Metadata
	UpdatedAt      pjson.Metadata
	DiscardedAt    pjson.Metadata
	Name           pjson.Metadata
	Description    pjson.Metadata
	NormalBalance  pjson.Metadata
	Balances       pjson.Metadata
	LockVersion    pjson.Metadata
	LedgerID       pjson.Metadata
	LedgerableID   pjson.Metadata
	LedgerableType pjson.Metadata
	Metadata       pjson.Metadata
	Raw            []byte
	Extras         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountNormalBalance string

const (
	LedgerAccountNormalBalanceCredit LedgerAccountNormalBalance = "credit"
	LedgerAccountNormalBalanceDebit  LedgerAccountNormalBalance = "debit"
)

type LedgerAccountBalances struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountBalancesPostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountBalancesAvailableBalance `json:"available_balance,required"`
	JSON             LedgerAccountBalancesJSON
}

type LedgerAccountBalancesJSON struct {
	PendingBalance   pjson.Metadata
	PostedBalance    pjson.Metadata
	AvailableBalance pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccountBalances using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountBalancesPendingBalanceJSON
}

type LedgerAccountBalancesPendingBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountBalancesPendingBalance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountBalancesPostedBalanceJSON
}

type LedgerAccountBalancesPostedBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountBalancesPostedBalance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountBalancesAvailableBalanceJSON
}

type LedgerAccountBalancesAvailableBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountBalancesAvailableBalance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountLedgerableType string

const (
	LedgerAccountLedgerableTypeExternalAccount LedgerAccountLedgerableType = "external_account"
	LedgerAccountLedgerableTypeInternalAccount LedgerAccountLedgerableType = "internal_account"
)
