// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// LedgerEntryService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewLedgerEntryService]
// method instead.
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
func (r *LedgerEntryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerEntry, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_entries/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all ledger entries.
func (r *LedgerEntryService) List(ctx context.Context, query LedgerEntryListParams, opts ...option.RequestOption) (res *shared.Page[LedgerEntry], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *LedgerEntryService) ListAutoPaging(ctx context.Context, query LedgerEntryListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerEntry] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

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
	JSON                           ledgerEntryJSON
}

// ledgerEntryJSON contains the JSON metadata for the struct [LedgerEntry]
type ledgerEntryJSON struct {
	ID                             apijson.Field
	Object                         apijson.Field
	LiveMode                       apijson.Field
	CreatedAt                      apijson.Field
	UpdatedAt                      apijson.Field
	DiscardedAt                    apijson.Field
	Amount                         apijson.Field
	Direction                      apijson.Field
	Status                         apijson.Field
	LedgerAccountID                apijson.Field
	LedgerAccountLockVersion       apijson.Field
	LedgerAccountCurrency          apijson.Field
	LedgerAccountCurrencyExponent  apijson.Field
	LedgerTransactionID            apijson.Field
	ResultingLedgerAccountBalances apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *LedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// The pending, posted, and available balances for this ledger entry's ledger
// account. The posted balance is the sum of all posted entries on the account. The
// pending balance is the sum of all pending and posted entries on the account. The
// available balance is the posted incoming entries minus the sum of the pending
// and posted outgoing amounts. Please see
// https://docs.moderntreasury.com/docs/transaction-status-and-balances for more
// details.
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
	JSON             ledgerEntryResultingLedgerAccountBalancesJSON
}

// ledgerEntryResultingLedgerAccountBalancesJSON contains the JSON metadata for the
// struct [LedgerEntryResultingLedgerAccountBalances]
type ledgerEntryResultingLedgerAccountBalancesJSON struct {
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	AvailableBalance apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerEntryResultingLedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerEntryResultingLedgerAccountBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerEntryResultingLedgerAccountBalancesPendingBalanceJSON
}

// ledgerEntryResultingLedgerAccountBalancesPendingBalanceJSON contains the JSON
// metadata for the struct
// [LedgerEntryResultingLedgerAccountBalancesPendingBalance]
type ledgerEntryResultingLedgerAccountBalancesPendingBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerEntryResultingLedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerEntryResultingLedgerAccountBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerEntryResultingLedgerAccountBalancesPostedBalanceJSON
}

// ledgerEntryResultingLedgerAccountBalancesPostedBalanceJSON contains the JSON
// metadata for the struct [LedgerEntryResultingLedgerAccountBalancesPostedBalance]
type ledgerEntryResultingLedgerAccountBalancesPostedBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerEntryResultingLedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerEntryResultingLedgerAccountBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerEntryResultingLedgerAccountBalancesAvailableBalanceJSON
}

// ledgerEntryResultingLedgerAccountBalancesAvailableBalanceJSON contains the JSON
// metadata for the struct
// [LedgerEntryResultingLedgerAccountBalancesAvailableBalance]
type ledgerEntryResultingLedgerAccountBalancesAvailableBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerEntryResultingLedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEntryListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Shows all ledger entries that were present on a ledger account at a particular
	// `lock_version`. You must also specify `ledger_account_id`.
	AsOfLockVersion param.Field[int64] `query:"as_of_lock_version"`
	// If true, response will include ledger entries that were deleted. When you update
	// a ledger transaction to specify a new set of entries, the previous entries are
	// deleted.
	Direction param.Field[LedgerEntryListParamsDirection] `query:"direction"`
	// Use "gt" (>), "gte" (>=), "lt" (<), "lte" (<=), or "eq" (=) to filter by the
	// transaction's effective time. Format ISO8601
	EffectiveAt param.Field[map[string]string] `query:"effective_at" format:"time"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// transaction's effective date. Format YYYY-MM-DD
	EffectiveDate param.Field[map[string]time.Time] `query:"effective_date" format:"date"`
	ID            param.Field[map[string]string]    `query:"id"`
	// Get all ledger entries that match the direction specified. One of `credit`,
	// `debit`.
	LedgerAccountCategoryID param.Field[string] `query:"ledger_account_category_id"`
	LedgerAccountID         param.Field[string] `query:"ledger_account_id"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// lock_version of a ledger account. For example, for all entries created at or
	// before before lock_version 1000 of a ledger account, use
	// `ledger_account_lock_version%5Blte%5D=1000`.
	LedgerAccountLockVersion param.Field[map[string]int64] `query:"ledger_account_lock_version"`
	LedgerTransactionID      param.Field[string]           `query:"ledger_transaction_id"`
	// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
	// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
	// by only one field at a time is supported.
	OrderBy param.Field[LedgerEntryListParamsOrderBy] `query:"order_by"`
	PerPage param.Field[int64]                        `query:"per_page"`
	// If true, response will include ledger entries that were deleted. When you update
	// a ledger transaction to specify a new set of entries, the previous entries are
	// deleted.
	ShowDeleted param.Field[bool] `query:"show_deleted"`
	// Get all ledger entries that match the status specified. One of `pending`,
	// `posted`, or `archived`.
	Status param.Field[LedgerEntryListParamsStatus] `query:"status"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt param.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes [LedgerEntryListParams]'s query parameters as `url.Values`.
func (r LedgerEntryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerEntryListParamsDirection string

const (
	LedgerEntryListParamsDirectionCredit LedgerEntryListParamsDirection = "credit"
	LedgerEntryListParamsDirectionDebit  LedgerEntryListParamsDirection = "debit"
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type LedgerEntryListParamsStatus string

const (
	LedgerEntryListParamsStatusPending  LedgerEntryListParamsStatus = "pending"
	LedgerEntryListParamsStatusPosted   LedgerEntryListParamsStatus = "posted"
	LedgerEntryListParamsStatusArchived LedgerEntryListParamsStatus = "archived"
)
