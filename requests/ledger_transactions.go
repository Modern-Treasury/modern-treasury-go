package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerTransactionNewParams struct {
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status field.Field[LedgerTransactionNewParamsStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate field.Field[time.Time] `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries field.Field[[]LedgerTransactionNewParamsLedgerEntries] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID field.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType field.Field[LedgerTransactionNewParamsLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID field.Field[string] `json:"ledgerable_id" format:"uuid"`
}

// MarshalJSON serializes LedgerTransactionNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r LedgerTransactionNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerTransactionNewParamsStatus string

const (
	LedgerTransactionNewParamsStatusArchived LedgerTransactionNewParamsStatus = "archived"
	LedgerTransactionNewParamsStatusPending  LedgerTransactionNewParamsStatus = "pending"
	LedgerTransactionNewParamsStatusPosted   LedgerTransactionNewParamsStatus = "posted"
)

type LedgerTransactionNewParamsLedgerEntries struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[LedgerTransactionNewParamsLedgerEntriesDirection] `json:"direction,required"`
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

type LedgerTransactionNewParamsLedgerEntriesDirection string

const (
	LedgerTransactionNewParamsLedgerEntriesDirectionCredit LedgerTransactionNewParamsLedgerEntriesDirection = "credit"
	LedgerTransactionNewParamsLedgerEntriesDirectionDebit  LedgerTransactionNewParamsLedgerEntriesDirection = "debit"
)

type LedgerTransactionNewParamsLedgerableType string

const (
	LedgerTransactionNewParamsLedgerableTypeCounterparty          LedgerTransactionNewParamsLedgerableType = "counterparty"
	LedgerTransactionNewParamsLedgerableTypeExpectedPayment       LedgerTransactionNewParamsLedgerableType = "expected_payment"
	LedgerTransactionNewParamsLedgerableTypeIncomingPaymentDetail LedgerTransactionNewParamsLedgerableType = "incoming_payment_detail"
	LedgerTransactionNewParamsLedgerableTypeInternalAccount       LedgerTransactionNewParamsLedgerableType = "internal_account"
	LedgerTransactionNewParamsLedgerableTypeLineItem              LedgerTransactionNewParamsLedgerableType = "line_item"
	LedgerTransactionNewParamsLedgerableTypePaperItem             LedgerTransactionNewParamsLedgerableType = "paper_item"
	LedgerTransactionNewParamsLedgerableTypePaymentOrder          LedgerTransactionNewParamsLedgerableType = "payment_order"
	LedgerTransactionNewParamsLedgerableTypePaymentOrderAttempt   LedgerTransactionNewParamsLedgerableType = "payment_order_attempt"
	LedgerTransactionNewParamsLedgerableTypeReturn                LedgerTransactionNewParamsLedgerableType = "return"
	LedgerTransactionNewParamsLedgerableTypeReversal              LedgerTransactionNewParamsLedgerableType = "reversal"
)

type LedgerTransactionUpdateParams struct {
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status field.Field[LedgerTransactionUpdateParamsStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// An array of ledger entry objects.
	LedgerEntries field.Field[[]LedgerTransactionUpdateParamsLedgerEntries] `json:"ledger_entries"`
}

// MarshalJSON serializes LedgerTransactionUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r LedgerTransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerTransactionUpdateParamsStatus string

const (
	LedgerTransactionUpdateParamsStatusArchived LedgerTransactionUpdateParamsStatus = "archived"
	LedgerTransactionUpdateParamsStatusPending  LedgerTransactionUpdateParamsStatus = "pending"
	LedgerTransactionUpdateParamsStatusPosted   LedgerTransactionUpdateParamsStatus = "posted"
)

type LedgerTransactionUpdateParamsLedgerEntries struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[LedgerTransactionUpdateParamsLedgerEntriesDirection] `json:"direction,required"`
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

type LedgerTransactionUpdateParamsLedgerEntriesDirection string

const (
	LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit LedgerTransactionUpdateParamsLedgerEntriesDirection = "credit"
	LedgerTransactionUpdateParamsLedgerEntriesDirectionDebit  LedgerTransactionUpdateParamsLedgerEntriesDirection = "debit"
)

type LedgerTransactionListParams struct {
	AfterCursor field.Field[string]            `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]             `query:"per_page"`
	ID          field.Field[map[string]string] `query:"id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata        field.Field[map[string]string] `query:"metadata"`
	LedgerID        field.Field[string]            `query:"ledger_id"`
	LedgerAccountID field.Field[string]            `query:"ledger_account_id"`
	// Use "gt" (>), "gte" (>=), "lt" (<), "lte" (<=), or "eq" (=) to filter by
	// effective at. For example, for all transactions after Jan 1 2000, use
	// effective_at%5Bgt%5D=2000-01-01T00:00:00:00.000Z.
	EffectiveAt field.Field[map[string]string] `query:"effective_at"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by
	// effective date. For example, for all dates after Jan 1 2000, use
	// effective_date%5Bgt%5D=2000-01-01.
	EffectiveDate field.Field[map[string]time.Time] `query:"effective_date"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// posted_at%5Bgt%5D=2000-01-01T12:00:00Z.
	PostedAt field.Field[map[string]time.Time] `query:"posted_at"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt field.Field[map[string]time.Time] `query:"updated_at"`
	// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
	// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
	// by only one field at a time is supported.
	OrderBy                 field.Field[LedgerTransactionListParamsOrderBy] `query:"order_by"`
	Status                  field.Field[LedgerTransactionListParamsStatus]  `query:"status"`
	ExternalID              field.Field[string]                             `query:"external_id"`
	LedgerAccountCategoryID field.Field[string]                             `query:"ledger_account_category_id"`
}

// URLQuery serializes LedgerTransactionListParams into a url.Values of the query
// parameters associated with this value
func (r LedgerTransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerTransactionListParamsOrderBy struct {
	CreatedAt   field.Field[LedgerTransactionListParamsOrderByCreatedAt]   `query:"created_at"`
	EffectiveAt field.Field[LedgerTransactionListParamsOrderByEffectiveAt] `query:"effective_at"`
}

// URLQuery serializes LedgerTransactionListParamsOrderBy into a url.Values of the
// query parameters associated with this value
func (r LedgerTransactionListParamsOrderBy) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerTransactionListParamsOrderByCreatedAt string

const (
	LedgerTransactionListParamsOrderByCreatedAtAsc  LedgerTransactionListParamsOrderByCreatedAt = "asc"
	LedgerTransactionListParamsOrderByCreatedAtDesc LedgerTransactionListParamsOrderByCreatedAt = "desc"
)

type LedgerTransactionListParamsOrderByEffectiveAt string

const (
	LedgerTransactionListParamsOrderByEffectiveAtAsc  LedgerTransactionListParamsOrderByEffectiveAt = "asc"
	LedgerTransactionListParamsOrderByEffectiveAtDesc LedgerTransactionListParamsOrderByEffectiveAt = "desc"
)

type LedgerTransactionListParamsStatus string

const (
	LedgerTransactionListParamsStatusPending  LedgerTransactionListParamsStatus = "pending"
	LedgerTransactionListParamsStatusPosted   LedgerTransactionListParamsStatus = "posted"
	LedgerTransactionListParamsStatusArchived LedgerTransactionListParamsStatus = "archived"
)
