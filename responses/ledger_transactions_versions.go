package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LedgerTransactionVersion struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the ledger transaction
	LedgerTransactionID string `json:"ledger_transaction_id,required" format:"uuid"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// One of `pending`, `posted`, or `archived`
	Status LedgerTransactionVersionStatus `json:"status,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerTransactionVersionLedgerEntries `json:"ledger_entries,required"`
	// The time on which the ledger transaction posted. This is null if the ledger
	// transaction is pending.
	PostedAt string `json:"posted_at,required,nullable" format:"time"`
	// The ID of the ledger this ledger transaction belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType LedgerTransactionVersionLedgerableType `json:"ledgerable_type,required,nullable"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID string `json:"external_id,required,nullable"`
	// Version number of the ledger transaction.
	Version int64 `json:"version,required"`
	JSON    LedgerTransactionVersionJSON
}

type LedgerTransactionVersionJSON struct {
	ID                  pjson.Metadata
	Object              pjson.Metadata
	LiveMode            pjson.Metadata
	CreatedAt           pjson.Metadata
	LedgerTransactionID pjson.Metadata
	Description         pjson.Metadata
	Status              pjson.Metadata
	Metadata            pjson.Metadata
	EffectiveDate       pjson.Metadata
	LedgerEntries       pjson.Metadata
	PostedAt            pjson.Metadata
	LedgerID            pjson.Metadata
	LedgerableType      pjson.Metadata
	LedgerableID        pjson.Metadata
	ExternalID          pjson.Metadata
	Version             pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerTransactionVersion
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *LedgerTransactionVersion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionStatus string

const (
	LedgerTransactionVersionStatusArchived LedgerTransactionVersionStatus = "archived"
	LedgerTransactionVersionStatusPending  LedgerTransactionVersionStatus = "pending"
	LedgerTransactionVersionStatusPosted   LedgerTransactionVersionStatus = "posted"
)

type LedgerTransactionVersionLedgerEntries struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount int64 `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction LedgerTransactionVersionLedgerEntriesDirection `json:"direction,required"`
	// Equal to the state of the ledger transaction when the ledger entry was created.
	// One of `pending`, `posted`, or `archived`.
	Status LedgerTransactionVersionLedgerEntriesStatus `json:"status,required"`
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
	ResultingLedgerAccountBalances LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances `json:"resulting_ledger_account_balances,required,nullable"`
	JSON                           LedgerTransactionVersionLedgerEntriesJSON
}

type LedgerTransactionVersionLedgerEntriesJSON struct {
	ID                             pjson.Metadata
	Object                         pjson.Metadata
	LiveMode                       pjson.Metadata
	CreatedAt                      pjson.Metadata
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

// UnmarshalJSON deserializes the provided bytes into
// LedgerTransactionVersionLedgerEntries using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerTransactionVersionLedgerEntries) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerEntriesDirection string

const (
	LedgerTransactionVersionLedgerEntriesDirectionCredit LedgerTransactionVersionLedgerEntriesDirection = "credit"
	LedgerTransactionVersionLedgerEntriesDirectionDebit  LedgerTransactionVersionLedgerEntriesDirection = "debit"
)

type LedgerTransactionVersionLedgerEntriesStatus string

const (
	LedgerTransactionVersionLedgerEntriesStatusArchived LedgerTransactionVersionLedgerEntriesStatus = "archived"
	LedgerTransactionVersionLedgerEntriesStatusPending  LedgerTransactionVersionLedgerEntriesStatus = "pending"
	LedgerTransactionVersionLedgerEntriesStatusPosted   LedgerTransactionVersionLedgerEntriesStatus = "posted"
)

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance `json:"available_balance,required"`
	JSON             LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesJSON
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesJSON struct {
	PendingBalance   pjson.Metadata
	PostedBalance    pjson.Metadata
	AvailableBalance pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalanceJSON
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalanceJSON
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalanceJSON
}

type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalanceJSON struct {
	Credits          pjson.Metadata
	Debits           pjson.Metadata
	Amount           pjson.Metadata
	Currency         pjson.Metadata
	CurrencyExponent pjson.Metadata
	Raw              []byte
	Extras           map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerableType string

const (
	LedgerTransactionVersionLedgerableTypeCounterparty          LedgerTransactionVersionLedgerableType = "counterparty"
	LedgerTransactionVersionLedgerableTypeExpectedPayment       LedgerTransactionVersionLedgerableType = "expected_payment"
	LedgerTransactionVersionLedgerableTypeIncomingPaymentDetail LedgerTransactionVersionLedgerableType = "incoming_payment_detail"
	LedgerTransactionVersionLedgerableTypeInternalAccount       LedgerTransactionVersionLedgerableType = "internal_account"
	LedgerTransactionVersionLedgerableTypeLineItem              LedgerTransactionVersionLedgerableType = "line_item"
	LedgerTransactionVersionLedgerableTypePaperItem             LedgerTransactionVersionLedgerableType = "paper_item"
	LedgerTransactionVersionLedgerableTypePaymentOrder          LedgerTransactionVersionLedgerableType = "payment_order"
	LedgerTransactionVersionLedgerableTypePaymentOrderAttempt   LedgerTransactionVersionLedgerableType = "payment_order_attempt"
	LedgerTransactionVersionLedgerableTypeReturn                LedgerTransactionVersionLedgerableType = "return"
	LedgerTransactionVersionLedgerableTypeReversal              LedgerTransactionVersionLedgerableType = "reversal"
)
