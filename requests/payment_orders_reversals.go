package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type PaymentOrderReversalNewParams struct {
	// The reason for the reversal. Must be one of `duplicate`, `incorrect_amount`,
	// `incorrect_receiving_account`, `date_earlier_than_intended`,
	// `date_later_than_intended`.
	Reason field.Field[PaymentOrderReversalNewParamsReason] `json:"reason,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// Specifies a ledger transaction object that will be created with the reversal. If
	// the ledger transaction cannot be created, then the reversal creation will fail.
	// The resulting ledger transaction will mirror the status of the reversal.
	LedgerTransaction field.Field[PaymentOrderReversalNewParamsLedgerTransaction] `json:"ledger_transaction"`
}

// MarshalJSON serializes PaymentOrderReversalNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r PaymentOrderReversalNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type PaymentOrderReversalNewParamsReason string

const (
	PaymentOrderReversalNewParamsReasonDuplicate                 PaymentOrderReversalNewParamsReason = "duplicate"
	PaymentOrderReversalNewParamsReasonIncorrectAmount           PaymentOrderReversalNewParamsReason = "incorrect_amount"
	PaymentOrderReversalNewParamsReasonIncorrectReceivingAccount PaymentOrderReversalNewParamsReason = "incorrect_receiving_account"
	PaymentOrderReversalNewParamsReasonDateEarlierThanIntended   PaymentOrderReversalNewParamsReason = "date_earlier_than_intended"
	PaymentOrderReversalNewParamsReasonDateLaterThanIntended     PaymentOrderReversalNewParamsReason = "date_later_than_intended"
)

type PaymentOrderReversalNewParamsLedgerTransaction struct {
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status field.Field[PaymentOrderReversalNewParamsLedgerTransactionStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate field.Field[time.Time] `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries field.Field[[]PaymentOrderReversalNewParamsLedgerTransactionLedgerEntries] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID field.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType field.Field[PaymentOrderReversalNewParamsLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID field.Field[string] `json:"ledgerable_id" format:"uuid"`
}

type PaymentOrderReversalNewParamsLedgerTransactionStatus string

const (
	PaymentOrderReversalNewParamsLedgerTransactionStatusArchived PaymentOrderReversalNewParamsLedgerTransactionStatus = "archived"
	PaymentOrderReversalNewParamsLedgerTransactionStatusPending  PaymentOrderReversalNewParamsLedgerTransactionStatus = "pending"
	PaymentOrderReversalNewParamsLedgerTransactionStatusPosted   PaymentOrderReversalNewParamsLedgerTransactionStatus = "posted"
)

type PaymentOrderReversalNewParamsLedgerTransactionLedgerEntries struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID field.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion field.Field[int64] `json:"lock_version,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount field.Field[map[string]int64] `json:"pending_balance_amount,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount field.Field[map[string]int64] `json:"posted_balance_amount,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount field.Field[map[string]int64] `json:"available_balance_amount,nullable"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances field.Field[bool] `json:"show_resulting_ledger_account_balances,nullable"`
}

type PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirection string

const (
	PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirectionCredit PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirection = "credit"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirectionDebit  PaymentOrderReversalNewParamsLedgerTransactionLedgerEntriesDirection = "debit"
)

type PaymentOrderReversalNewParamsLedgerTransactionLedgerableType string

const (
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeCounterparty          PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "counterparty"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeExpectedPayment       PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "expected_payment"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeIncomingPaymentDetail PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "incoming_payment_detail"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeInternalAccount       PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "internal_account"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeLineItem              PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "line_item"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypePaperItem             PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "paper_item"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypePaymentOrder          PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "payment_order"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypePaymentOrderAttempt   PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "payment_order_attempt"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeReturn                PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "return"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeReversal              PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "reversal"
)

type PaymentOrderReversalListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes PaymentOrderReversalListParams into a url.Values of the
// query parameters associated with this value
func (r PaymentOrderReversalListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
