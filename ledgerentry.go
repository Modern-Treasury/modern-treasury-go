// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
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

// LedgerEntryService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerEntryService] method instead.
type LedgerEntryService struct {
	Options []option.RequestOption
}

// NewLedgerEntryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLedgerEntryService(opts ...option.RequestOption) (r *LedgerEntryService) {
	r = &LedgerEntryService{}
	r.Options = opts
	return
}

// Get details on a single ledger entry.
func (r *LedgerEntryService) Get(ctx context.Context, id string, query LedgerEntryGetParams, opts ...option.RequestOption) (res *LedgerEntry, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_entries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update the details of a ledger entry.
func (r *LedgerEntryService) Update(ctx context.Context, id string, body LedgerEntryUpdateParams, opts ...option.RequestOption) (res *LedgerEntry, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_entries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all ledger entries.
func (r *LedgerEntryService) List(ctx context.Context, query LedgerEntryListParams, opts ...option.RequestOption) (res *pagination.Page[LedgerEntry], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_entries"
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

// Get a list of all ledger entries.
func (r *LedgerEntryService) ListAutoPaging(ctx context.Context, query LedgerEntryListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LedgerEntry] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type LedgerEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount    int64     `json:"amount,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction   shared.TransactionDirection `json:"direction,required"`
	DiscardedAt time.Time                   `json:"discarded_at,required,nullable" format:"date-time"`
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
	Status    LedgerEntryStatus `json:"status,required"`
	UpdatedAt time.Time         `json:"updated_at,required" format:"date-time"`
	JSON      ledgerEntryJSON   `json:"-"`
}

// ledgerEntryJSON contains the JSON metadata for the struct [LedgerEntry]
type ledgerEntryJSON struct {
	ID                             apijson.Field
	Amount                         apijson.Field
	CreatedAt                      apijson.Field
	Direction                      apijson.Field
	DiscardedAt                    apijson.Field
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
	UpdatedAt                      apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *LedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerEntryJSON) RawJSON() string {
	return r.raw
}

// Equal to the state of the ledger transaction when the ledger entry was created.
// One of `pending`, `posted`, or `archived`.
type LedgerEntryStatus string

const (
	LedgerEntryStatusArchived LedgerEntryStatus = "archived"
	LedgerEntryStatusPending  LedgerEntryStatus = "pending"
	LedgerEntryStatusPosted   LedgerEntryStatus = "posted"
)

func (r LedgerEntryStatus) IsKnown() bool {
	switch r {
	case LedgerEntryStatusArchived, LedgerEntryStatusPending, LedgerEntryStatusPosted:
		return true
	}
	return false
}

type LedgerEntryGetParams struct {
	// If true, response will include the balances attached to the ledger entry. If
	// there is no balance available, null will be returned instead.
	ShowBalances param.Field[bool] `query:"show_balances"`
}

// URLQuery serializes [LedgerEntryGetParams]'s query parameters as `url.Values`.
func (r LedgerEntryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerEntryUpdateParams struct {
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEntryListParams struct {
	// If you have specific IDs to retrieve in bulk, you can pass them as query
	// parameters delimited with `id[]=`, for example `?id[]=123&id[]=abc`.
	ID          param.Field[[]string] `query:"id"`
	AfterCursor param.Field[string]   `query:"after_cursor"`
	// Shows all ledger entries that were present on a ledger account at a particular
	// `lock_version`. You must also specify `ledger_account_id`.
	AsOfLockVersion param.Field[int64] `query:"as_of_lock_version"`
	// If true, response will include ledger entries that were deleted. When you update
	// a ledger transaction to specify a new set of entries, the previous entries are
	// deleted.
	Direction param.Field[shared.TransactionDirection] `query:"direction"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// transaction's effective time. Format ISO8601
	EffectiveAt param.Field[map[string]time.Time] `query:"effective_at" format:"date-time"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// transaction's effective date. Format YYYY-MM-DD
	EffectiveDate param.Field[map[string]time.Time] `query:"effective_date" format:"date"`
	// Get all ledger entries that match the direction specified. One of `credit`,
	// `debit`.
	LedgerAccountCategoryID param.Field[string] `query:"ledger_account_category_id"`
	LedgerAccountID         param.Field[string] `query:"ledger_account_id"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// lock_version of a ledger account. For example, for all entries created at or
	// before before lock_version 1000 of a ledger account, use
	// `ledger_account_lock_version%5Blte%5D=1000`.
	LedgerAccountLockVersion  param.Field[map[string]int64] `query:"ledger_account_lock_version"`
	LedgerAccountPayoutID     param.Field[string]           `query:"ledger_account_payout_id"`
	LedgerAccountSettlementID param.Field[string]           `query:"ledger_account_settlement_id"`
	// Get all ledger entries that are included in the ledger account statement.
	LedgerAccountStatementID param.Field[string] `query:"ledger_account_statement_id"`
	LedgerTransactionID      param.Field[string] `query:"ledger_transaction_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
	// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
	// by only one field at a time is supported.
	OrderBy param.Field[LedgerEntryListParamsOrderBy] `query:"order_by"`
	PerPage param.Field[int64]                        `query:"per_page"`
	// If true, response will include the balances attached to the ledger entry. If
	// there is no balance available, null will be returned instead.
	ShowBalances param.Field[bool] `query:"show_balances"`
	// If true, response will include ledger entries that were deleted. When you update
	// a ledger transaction to specify a new set of entries, the previous entries are
	// deleted.
	ShowDeleted param.Field[bool] `query:"show_deleted"`
	// Get all ledger entries that match the status specified. One of `pending`,
	// `posted`, or `archived`. For multiple statuses, use
	// `status[]=pending&status[]=posted`.
	Status param.Field[LedgerEntryListParamsStatusUnion] `query:"status"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt param.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes [LedgerEntryListParams]'s query parameters as `url.Values`.
func (r LedgerEntryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
// by only one field at a time is supported.
type LedgerEntryListParamsOrderBy struct {
	CreatedAt   param.Field[LedgerEntryListParamsOrderByCreatedAt]   `query:"created_at"`
	EffectiveAt param.Field[LedgerEntryListParamsOrderByEffectiveAt] `query:"effective_at"`
}

// URLQuery serializes [LedgerEntryListParamsOrderBy]'s query parameters as
// `url.Values`.
func (r LedgerEntryListParamsOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerEntryListParamsOrderByCreatedAt string

const (
	LedgerEntryListParamsOrderByCreatedAtAsc  LedgerEntryListParamsOrderByCreatedAt = "asc"
	LedgerEntryListParamsOrderByCreatedAtDesc LedgerEntryListParamsOrderByCreatedAt = "desc"
)

func (r LedgerEntryListParamsOrderByCreatedAt) IsKnown() bool {
	switch r {
	case LedgerEntryListParamsOrderByCreatedAtAsc, LedgerEntryListParamsOrderByCreatedAtDesc:
		return true
	}
	return false
}

type LedgerEntryListParamsOrderByEffectiveAt string

const (
	LedgerEntryListParamsOrderByEffectiveAtAsc  LedgerEntryListParamsOrderByEffectiveAt = "asc"
	LedgerEntryListParamsOrderByEffectiveAtDesc LedgerEntryListParamsOrderByEffectiveAt = "desc"
)

func (r LedgerEntryListParamsOrderByEffectiveAt) IsKnown() bool {
	switch r {
	case LedgerEntryListParamsOrderByEffectiveAtAsc, LedgerEntryListParamsOrderByEffectiveAtDesc:
		return true
	}
	return false
}

// Get all ledger entries that match the status specified. One of `pending`,
// `posted`, or `archived`. For multiple statuses, use
// `status[]=pending&status[]=posted`.
//
// Satisfied by [LedgerEntryListParamsStatusString],
// [LedgerEntryListParamsStatusArray].
type LedgerEntryListParamsStatusUnion interface {
	implementsLedgerEntryListParamsStatusUnion()
}

type LedgerEntryListParamsStatusString string

const (
	LedgerEntryListParamsStatusStringPending  LedgerEntryListParamsStatusString = "pending"
	LedgerEntryListParamsStatusStringPosted   LedgerEntryListParamsStatusString = "posted"
	LedgerEntryListParamsStatusStringArchived LedgerEntryListParamsStatusString = "archived"
)

func (r LedgerEntryListParamsStatusString) IsKnown() bool {
	switch r {
	case LedgerEntryListParamsStatusStringPending, LedgerEntryListParamsStatusStringPosted, LedgerEntryListParamsStatusStringArchived:
		return true
	}
	return false
}

func (r LedgerEntryListParamsStatusString) implementsLedgerEntryListParamsStatusUnion() {}

type LedgerEntryListParamsStatusArray []LedgerEntryListParamsStatusArrayItem

func (r LedgerEntryListParamsStatusArray) implementsLedgerEntryListParamsStatusUnion() {}

type LedgerEntryListParamsStatusArrayItem string

const (
	LedgerEntryListParamsStatusArrayItemPending  LedgerEntryListParamsStatusArrayItem = "pending"
	LedgerEntryListParamsStatusArrayItemPosted   LedgerEntryListParamsStatusArrayItem = "posted"
	LedgerEntryListParamsStatusArrayItemArchived LedgerEntryListParamsStatusArrayItem = "archived"
)

func (r LedgerEntryListParamsStatusArrayItem) IsKnown() bool {
	switch r {
	case LedgerEntryListParamsStatusArrayItemPending, LedgerEntryListParamsStatusArrayItemPosted, LedgerEntryListParamsStatusArrayItemArchived:
		return true
	}
	return false
}
