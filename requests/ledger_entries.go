package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerEntryListParams struct {
	AfterCursor         field.Field[string]            `query:"after_cursor,nullable"`
	PerPage             field.Field[int64]             `query:"per_page"`
	ID                  field.Field[map[string]string] `query:"id"`
	LedgerAccountID     field.Field[string]            `query:"ledger_account_id"`
	LedgerTransactionID field.Field[string]            `query:"ledger_transaction_id"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// transaction's effective date. Format YYYY-MM-DD
	EffectiveDate field.Field[map[string]time.Time] `query:"effective_date"`
	// Use "gt" (>), "gte" (>=), "lt" (<), "lte" (<=), or "eq" (=) to filter by the
	// transaction's effective time. Format ISO8601
	EffectiveAt field.Field[map[string]string] `query:"effective_at"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt field.Field[map[string]time.Time] `query:"updated_at"`
	// Shows all ledger entries that were present on a ledger account at a particular
	// `lock_version`. You must also specify `ledger_account_id`.
	AsOfLockVersion field.Field[int64] `query:"as_of_lock_version"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// lock_version of a ledger account. For example, for all entries created at or
	// before before lock_version 1000 of a ledger account, use
	// `ledger_account_lock_version%5Blte%5D=1000`.
	LedgerAccountLockVersion field.Field[map[string]int64] `query:"ledger_account_lock_version"`
	// Get all ledger entries that match the direction specified. One of `credit`,
	// `debit`.
	LedgerAccountCategoryID field.Field[string] `query:"ledger_account_category_id"`
	// If true, response will include ledger entries that were deleted. When you update
	// a ledger transaction to specify a new set of entries, the previous entries are
	// deleted.
	ShowDeleted field.Field[bool] `query:"show_deleted"`
	// If true, response will include ledger entries that were deleted. When you update
	// a ledger transaction to specify a new set of entries, the previous entries are
	// deleted.
	Direction field.Field[LedgerEntryListParamsDirection] `query:"direction"`
	// Get all ledger entries that match the status specified. One of `pending`,
	// `posted`, or `archived`.
	Status field.Field[LedgerEntryListParamsStatus] `query:"status"`
	// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
	// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
	// by only one field at a time is supported.
	OrderBy field.Field[LedgerEntryListParamsOrderBy] `query:"order_by"`
}

// URLQuery serializes LedgerEntryListParams into a url.Values of the query
// parameters associated with this value
func (r LedgerEntryListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerEntryListParamsDirection string

const (
	LedgerEntryListParamsDirectionCredit LedgerEntryListParamsDirection = "credit"
	LedgerEntryListParamsDirectionDebit  LedgerEntryListParamsDirection = "debit"
)

type LedgerEntryListParamsStatus string

const (
	LedgerEntryListParamsStatusPending  LedgerEntryListParamsStatus = "pending"
	LedgerEntryListParamsStatusPosted   LedgerEntryListParamsStatus = "posted"
	LedgerEntryListParamsStatusArchived LedgerEntryListParamsStatus = "archived"
)

type LedgerEntryListParamsOrderBy struct {
	CreatedAt   field.Field[LedgerEntryListParamsOrderByCreatedAt]   `query:"created_at"`
	EffectiveAt field.Field[LedgerEntryListParamsOrderByEffectiveAt] `query:"effective_at"`
}

// URLQuery serializes LedgerEntryListParamsOrderBy into a url.Values of the query
// parameters associated with this value
func (r LedgerEntryListParamsOrderBy) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerEntryListParamsOrderByCreatedAt string

const (
	LedgerEntryListParamsOrderByCreatedAtAsc  LedgerEntryListParamsOrderByCreatedAt = "asc"
	LedgerEntryListParamsOrderByCreatedAtDesc LedgerEntryListParamsOrderByCreatedAt = "desc"
)

type LedgerEntryListParamsOrderByEffectiveAt string

const (
	LedgerEntryListParamsOrderByEffectiveAtAsc  LedgerEntryListParamsOrderByEffectiveAt = "asc"
	LedgerEntryListParamsOrderByEffectiveAtDesc LedgerEntryListParamsOrderByEffectiveAt = "desc"
)
