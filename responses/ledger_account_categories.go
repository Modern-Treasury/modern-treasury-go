package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LedgerAccountCategory struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the ledger account category.
	Name string `json:"name,required"`
	// The description of the ledger account category.
	Description string `json:"description,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The id of the ledger that this account category belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// The normal balance of the ledger account category.
	NormalBalance LedgerAccountCategoryNormalBalance `json:"normal_balance,required"`
	// The pending, posted, and available balances for this ledger account category.
	// The posted balance is the sum of all posted entries on the account. The pending
	// balance is the sum of all pending and posted entries on the account. The
	// available balance is the posted incoming entries minus the sum of the pending
	// and posted outgoing amounts.
	Balances LedgerAccountCategoryBalances `json:"balances,required"`
	JSON     LedgerAccountCategoryJSON
}

type LedgerAccountCategoryJSON struct {
	ID            pjson.Metadata
	Object        pjson.Metadata
	LiveMode      pjson.Metadata
	CreatedAt     pjson.Metadata
	UpdatedAt     pjson.Metadata
	DiscardedAt   pjson.Metadata
	Name          pjson.Metadata
	Description   pjson.Metadata
	Metadata      pjson.Metadata
	LedgerID      pjson.Metadata
	NormalBalance pjson.Metadata
	Balances      pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccountCategory using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerAccountCategory) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryNormalBalance string

const (
	LedgerAccountCategoryNormalBalanceCredit LedgerAccountCategoryNormalBalance = "credit"
	LedgerAccountCategoryNormalBalanceDebit  LedgerAccountCategoryNormalBalance = "debit"
)

type LedgerAccountCategoryBalances struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountCategoryBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountCategoryBalancesPostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountCategoryBalancesAvailableBalance `json:"available_balance,required"`
	JSON             LedgerAccountCategoryBalancesJSON
}

type LedgerAccountCategoryBalancesJSON struct {
	PendingBalance   pjson.Metadata
	PostedBalance    pjson.Metadata
	AvailableBalance pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccountCategoryBalances
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *LedgerAccountCategoryBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountCategoryBalancesPendingBalanceJSON
}

type LedgerAccountCategoryBalancesPendingBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountCategoryBalancesPendingBalance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountCategoryBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountCategoryBalancesPostedBalanceJSON
}

type LedgerAccountCategoryBalancesPostedBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountCategoryBalancesPostedBalance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountCategoryBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountCategoryBalancesAvailableBalanceJSON
}

type LedgerAccountCategoryBalancesAvailableBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountCategoryBalancesAvailableBalance using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountCategoryBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
