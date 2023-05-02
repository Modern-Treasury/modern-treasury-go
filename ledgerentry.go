package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type LedgerEntryService struct {
	Options []option.RequestOption
}

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
	JSON                           LedgerEntryJSON
}

type LedgerEntryJSON struct {
	ID                             apijson.Metadata
	Object                         apijson.Metadata
	LiveMode                       apijson.Metadata
	CreatedAt                      apijson.Metadata
	UpdatedAt                      apijson.Metadata
	DiscardedAt                    apijson.Metadata
	Amount                         apijson.Metadata
	Direction                      apijson.Metadata
	Status                         apijson.Metadata
	LedgerAccountID                apijson.Metadata
	LedgerAccountLockVersion       apijson.Metadata
	LedgerAccountCurrency          apijson.Metadata
	LedgerAccountCurrencyExponent  apijson.Metadata
	LedgerTransactionID            apijson.Metadata
	ResultingLedgerAccountBalances apijson.Metadata
	raw                            string
	Extras                         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerEntry using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
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
	PendingBalance   apijson.Metadata
	PostedBalance    apijson.Metadata
	AvailableBalance apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalances using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Credits          apijson.Metadata
	Debits           apijson.Metadata
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	CurrencyExponent apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalancesPendingBalance using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Credits          apijson.Metadata
	Debits           apijson.Metadata
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	CurrencyExponent apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalancesPostedBalance using the internal json
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Credits          apijson.Metadata
	Debits           apijson.Metadata
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	CurrencyExponent apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerEntryResultingLedgerAccountBalancesAvailableBalance using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerEntryResultingLedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEntryListParams struct {
	AfterCursor         field.Field[string]            `query:"after_cursor,nullable"`
	PerPage             field.Field[int64]             `query:"per_page"`
	ID                  field.Field[map[string]string] `query:"id"`
	LedgerAccountID     field.Field[string]            `query:"ledger_account_id"`
	LedgerTransactionID field.Field[string]            `query:"ledger_transaction_id"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// transaction's effective date. Format YYYY-MM-DD
	EffectiveDate field.Field[map[string]time.Time] `query:"effective_date" format:"date"`
	// Use "gt" (>), "gte" (>=), "lt" (<), "lte" (<=), or "eq" (=) to filter by the
	// transaction's effective time. Format ISO8601
	EffectiveAt field.Field[map[string]string] `query:"effective_at" format:"time"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt field.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
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
	return apiquery.Marshal(r)
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
	return apiquery.Marshal(r)
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
