package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LedgerEntry struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount int64 `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction LedgerEntryDirection `json:"direction,required"`
	// Equal to the state of the ledger transaction when the ledger entry was created.
	// One of `pending`, `posted`, or `archived`.
	Status LedgerEntryStatus `json:"status,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID string `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required,nullable"`
	// The currency of the ledger account.
	LedgerAccountCurrency string `json:"ledger_account_currency,required"`
	// The currency exponent of the ledger account.
	LedgerAccountCurrencyExponent int64 `json:"ledger_account_currency_exponent,required"`
	// The ledger transaction that this ledger entry is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required"`
	// The pending, posted, and available balances for this ledger entry's ledger
	// account. The posted balance is the sum of all posted entries on the account. The
	// pending balance is the sum of all pending and posted entries on the account. The
	// available balance is the posted incoming entries minus the sum of the pending
	// and posted outgoing amounts. Please see
	// https://docs.moderntreasury.com/docs/transaction-status-and-balances for more
	// details.
	ResultingLedgerAccountBalances LedgerEntryResultingLedgerAccountBalances `json:"resulting_ledger_account_balances,required,nullable"`
	JSON                           LedgerEntryJSON
}

type LedgerEntryJSON struct {
	ID                             pjson.Metadata
	Object                         pjson.Metadata
	LiveMode                       pjson.Metadata
	CreatedAt                      pjson.Metadata
	UpdatedAt                      pjson.Metadata
	DiscardedAt                    pjson.Metadata
	Amount                         pjson.Metadata
	Direction                      pjson.Metadata
	Status                         pjson.Metadata
	LedgerAccountID                pjson.Metadata
	LedgerAccountLockVersion       pjson.Metadata
	LedgerAccountCurrency          pjson.Metadata
	LedgerAccountCurrencyExponent  pjson.Metadata
	LedgerTransactionID            pjson.Metadata
	ResultingLedgerAccountBalances pjson.Metadata
	Raw                            []byte
	Extras                         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerEntry using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerEntryDirection string

const (
	LedgerEntryDirectionCredit LedgerEntryDirection = "credit"
	LedgerEntryDirectionDebit  LedgerEntryDirection = "debit"
)

type LedgerEntryStatus string

const (
	LedgerEntryStatusArchived LedgerEntryStatus = "archived"
	LedgerEntryStatusPending  LedgerEntryStatus = "pending"
	LedgerEntryStatusPosted   LedgerEntryStatus = "posted"
)

type LedgerEntryResultingLedgerAccountBalances struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerEntryResultingLedgerAccountBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerEntryResultingLedgerAccountBalancesPostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerEntryResultingLedgerAccountBalancesAvailableBalance `json:"available_balance,required"`
	JSON             LedgerEntryResultingLedgerAccountBalancesJSON
}

type LedgerEntryResultingLedgerAccountBalancesJSON struct {
	PendingBalance   pjson.Metadata
	PostedBalance    pjson.Metadata
	AvailableBalance pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalances using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerEntryResultingLedgerAccountBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerEntryResultingLedgerAccountBalancesPendingBalanceJSON
}

type LedgerEntryResultingLedgerAccountBalancesPendingBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalancesPendingBalance using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerEntryResultingLedgerAccountBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerEntryResultingLedgerAccountBalancesPostedBalanceJSON
}

type LedgerEntryResultingLedgerAccountBalancesPostedBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalancesPostedBalance using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerEntryResultingLedgerAccountBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerEntryResultingLedgerAccountBalancesAvailableBalanceJSON
}

type LedgerEntryResultingLedgerAccountBalancesAvailableBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalancesAvailableBalance using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
