// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// LedgerTransactionVersionService contains methods and other services that help
// with interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerTransactionVersionService] method instead.
type LedgerTransactionVersionService struct {
	Options []option.RequestOption
}

// NewLedgerTransactionVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLedgerTransactionVersionService(opts ...option.RequestOption) (r *LedgerTransactionVersionService) {
	r = &LedgerTransactionVersionService{}
	r.Options = opts
	return
}

// Get a list of ledger transaction versions.
func (r *LedgerTransactionVersionService) List(ctx context.Context, query LedgerTransactionVersionListParams, opts ...option.RequestOption) (res *pagination.Page[LedgerTransactionVersion], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_transaction_versions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of ledger transaction versions.
func (r *LedgerTransactionVersionService) ListAutoPaging(ctx context.Context, query LedgerTransactionVersionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LedgerTransactionVersion] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type LedgerTransactionVersion struct {
	ID string `json:"id,required" format:"uuid"`
	// System-set reason why the ledger transaction was archived; currently only
	// 'balance_lock_failure' for transactions that violated balance constraints. Only
	// populated when archive_on_balance_lock_failure is true and a balance lock
	// violation occurs, otherwise null.
	ArchivedReason string    `json:"archived_reason,required,nullable"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID string `json:"external_id,required,nullable"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerTransactionVersionLedgerEntry `json:"ledger_entries,required"`
	// The ID of the ledger this ledger transaction belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// The ID of the ledger transaction
	LedgerTransactionID string `json:"ledger_transaction_id,required" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType LedgerTransactionVersionLedgerableType `json:"ledgerable_type,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The ID of the ledger transaction that this ledger transaction partially posts.
	PartiallyPostsLedgerTransactionID string `json:"partially_posts_ledger_transaction_id,required,nullable"`
	// The time on which the ledger transaction posted. This is null if the ledger
	// transaction is pending.
	PostedAt time.Time `json:"posted_at,required,nullable" format:"date-time"`
	// The ID of the ledger transaction that reversed this ledger transaction.
	ReversedByLedgerTransactionID string `json:"reversed_by_ledger_transaction_id,required,nullable"`
	// The ID of the original ledger transaction. that this ledger transaction
	// reverses.
	ReversesLedgerTransactionID string `json:"reverses_ledger_transaction_id,required,nullable"`
	// One of `pending`, `posted`, or `archived`.
	Status LedgerTransactionVersionStatus `json:"status,required"`
	// Version number of the ledger transaction.
	Version int64                        `json:"version,required"`
	JSON    ledgerTransactionVersionJSON `json:"-"`
}

// ledgerTransactionVersionJSON contains the JSON metadata for the struct
// [LedgerTransactionVersion]
type ledgerTransactionVersionJSON struct {
	ID                                apijson.Field
	ArchivedReason                    apijson.Field
	CreatedAt                         apijson.Field
	Description                       apijson.Field
	EffectiveAt                       apijson.Field
	EffectiveDate                     apijson.Field
	ExternalID                        apijson.Field
	LedgerEntries                     apijson.Field
	LedgerID                          apijson.Field
	LedgerTransactionID               apijson.Field
	LedgerableID                      apijson.Field
	LedgerableType                    apijson.Field
	LiveMode                          apijson.Field
	Metadata                          apijson.Field
	Object                            apijson.Field
	PartiallyPostsLedgerTransactionID apijson.Field
	PostedAt                          apijson.Field
	ReversedByLedgerTransactionID     apijson.Field
	ReversesLedgerTransactionID       apijson.Field
	Status                            apijson.Field
	Version                           apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *LedgerTransactionVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerTransactionVersionJSON) RawJSON() string {
	return r.raw
}

type LedgerTransactionVersionLedgerEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount    int64     `json:"amount,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction shared.TransactionDirection `json:"direction,required"`
	// The currency of the ledger account.
	LedgerAccountCurrency string `json:"ledger_account_currency,required"`
	// The currency exponent of the ledger account.
	LedgerAccountCurrencyExponent int64 `json:"ledger_account_currency_exponent,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID string `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required,nullable"`
	// The ledger transaction that this ledger entry is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The pending, posted, and available balances for this ledger entry's ledger
	// account. The posted balance is the sum of all posted entries on the account. The
	// pending balance is the sum of all pending and posted entries on the account. The
	// available balance is the posted incoming entries minus the sum of the pending
	// and posted outgoing amounts. Please see
	// https://docs.moderntreasury.com/docs/transaction-status-and-balances for more
	// details.
	ResultingLedgerAccountBalances shared.LedgerBalances `json:"resulting_ledger_account_balances,required,nullable"`
	// Equal to the state of the ledger transaction when the ledger entry was created.
	// One of `pending`, `posted`, or `archived`.
	Status LedgerTransactionVersionLedgerEntriesStatus `json:"status,required"`
	JSON   ledgerTransactionVersionLedgerEntryJSON     `json:"-"`
}

// ledgerTransactionVersionLedgerEntryJSON contains the JSON metadata for the
// struct [LedgerTransactionVersionLedgerEntry]
type ledgerTransactionVersionLedgerEntryJSON struct {
	ID                             apijson.Field
	Amount                         apijson.Field
	CreatedAt                      apijson.Field
	Direction                      apijson.Field
	LedgerAccountCurrency          apijson.Field
	LedgerAccountCurrencyExponent  apijson.Field
	LedgerAccountID                apijson.Field
	LedgerAccountLockVersion       apijson.Field
	LedgerTransactionID            apijson.Field
	LiveMode                       apijson.Field
	Metadata                       apijson.Field
	Object                         apijson.Field
	ResultingLedgerAccountBalances apijson.Field
	Status                         apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *LedgerTransactionVersionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerTransactionVersionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

// Equal to the state of the ledger transaction when the ledger entry was created.
// One of `pending`, `posted`, or `archived`.
type LedgerTransactionVersionLedgerEntriesStatus string

const (
	LedgerTransactionVersionLedgerEntriesStatusArchived LedgerTransactionVersionLedgerEntriesStatus = "archived"
	LedgerTransactionVersionLedgerEntriesStatusPending  LedgerTransactionVersionLedgerEntriesStatus = "pending"
	LedgerTransactionVersionLedgerEntriesStatusPosted   LedgerTransactionVersionLedgerEntriesStatus = "posted"
)

func (r LedgerTransactionVersionLedgerEntriesStatus) IsKnown() bool {
	switch r {
	case LedgerTransactionVersionLedgerEntriesStatusArchived, LedgerTransactionVersionLedgerEntriesStatusPending, LedgerTransactionVersionLedgerEntriesStatusPosted:
		return true
	}
	return false
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type LedgerTransactionVersionLedgerableType string

const (
	LedgerTransactionVersionLedgerableTypeExpectedPayment       LedgerTransactionVersionLedgerableType = "expected_payment"
	LedgerTransactionVersionLedgerableTypeIncomingPaymentDetail LedgerTransactionVersionLedgerableType = "incoming_payment_detail"
	LedgerTransactionVersionLedgerableTypePaperItem             LedgerTransactionVersionLedgerableType = "paper_item"
	LedgerTransactionVersionLedgerableTypePaymentOrder          LedgerTransactionVersionLedgerableType = "payment_order"
	LedgerTransactionVersionLedgerableTypeReturn                LedgerTransactionVersionLedgerableType = "return"
	LedgerTransactionVersionLedgerableTypeReversal              LedgerTransactionVersionLedgerableType = "reversal"
)

func (r LedgerTransactionVersionLedgerableType) IsKnown() bool {
	switch r {
	case LedgerTransactionVersionLedgerableTypeExpectedPayment, LedgerTransactionVersionLedgerableTypeIncomingPaymentDetail, LedgerTransactionVersionLedgerableTypePaperItem, LedgerTransactionVersionLedgerableTypePaymentOrder, LedgerTransactionVersionLedgerableTypeReturn, LedgerTransactionVersionLedgerableTypeReversal:
		return true
	}
	return false
}

// One of `pending`, `posted`, or `archived`.
type LedgerTransactionVersionStatus string

const (
	LedgerTransactionVersionStatusArchived LedgerTransactionVersionStatus = "archived"
	LedgerTransactionVersionStatusPending  LedgerTransactionVersionStatus = "pending"
	LedgerTransactionVersionStatusPosted   LedgerTransactionVersionStatus = "posted"
)

func (r LedgerTransactionVersionStatus) IsKnown() bool {
	switch r {
	case LedgerTransactionVersionStatusArchived, LedgerTransactionVersionStatusPending, LedgerTransactionVersionStatusPosted:
		return true
	}
	return false
}

type LedgerTransactionVersionListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// created_at timestamp. For example, for all dates after Jan 1 2000 12:00 UTC, use
	// created_at%5Bgt%5D=2000-01-01T12:00:00Z.
	CreatedAt param.Field[map[string]time.Time] `query:"created_at" format:"date-time"`
	// Get all ledger transaction versions that are included in the ledger account
	// statement.
	LedgerAccountStatementID param.Field[string] `query:"ledger_account_statement_id"`
	// Get all the ledger transaction versions corresponding to the ID of a ledger
	// transaction.
	LedgerTransactionID param.Field[string] `query:"ledger_transaction_id"`
	PerPage             param.Field[int64]  `query:"per_page"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// version. For example, for all versions after 2, use version%5Bgt%5D=2.
	Version param.Field[map[string]int64] `query:"version"`
}

// URLQuery serializes [LedgerTransactionVersionListParams]'s query parameters as
// `url.Values`.
func (r LedgerTransactionVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
