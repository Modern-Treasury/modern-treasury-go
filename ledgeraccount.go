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

// LedgerAccountService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewLedgerAccountService]
// method instead.
type LedgerAccountService struct {
	Options []option.RequestOption
}

// NewLedgerAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLedgerAccountService(opts ...option.RequestOption) (r *LedgerAccountService) {
	r = &LedgerAccountService{}
	r.Options = opts
	return
}

// Create a ledger account.
func (r *LedgerAccountService) New(ctx context.Context, params LedgerAccountNewParams, opts ...option.RequestOption) (res *LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details on a single ledger account.
func (r *LedgerAccountService) Get(ctx context.Context, id string, query LedgerAccountGetParams, opts ...option.RequestOption) (res *LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update the details of a ledger account.
func (r *LedgerAccountService) Update(ctx context.Context, id string, body LedgerAccountUpdateParams, opts ...option.RequestOption) (res *LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledger accounts.
func (r *LedgerAccountService) List(ctx context.Context, query LedgerAccountListParams, opts ...option.RequestOption) (res *shared.Page[LedgerAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_accounts"
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

// Get a list of ledger accounts.
func (r *LedgerAccountService) ListAutoPaging(ctx context.Context, query LedgerAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger account.
func (r *LedgerAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LedgerAccount struct {
	ID string `json:"id,required" format:"uuid"`
	// The pending, posted, and available balances for this ledger account. The posted
	// balance is the sum of all posted entries on the account. The pending balance is
	// the sum of all pending and posted entries on the account. The available balance
	// is the posted incoming entries minus the sum of the pending and posted outgoing
	// amounts.
	Balances  LedgerAccountBalances `json:"balances,required"`
	CreatedAt time.Time             `json:"created_at,required" format:"date-time"`
	// The description of the ledger account.
	Description string    `json:"description,required,nullable"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The id of the ledger that this account belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType LedgerAccountLedgerableType `json:"ledgerable_type,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Lock version of the ledger account.
	LockVersion int64 `json:"lock_version,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The name of the ledger account.
	Name string `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance LedgerAccountNormalBalance `json:"normal_balance,required"`
	Object        string                     `json:"object,required"`
	UpdatedAt     time.Time                  `json:"updated_at,required" format:"date-time"`
	JSON          ledgerAccountJSON
}

// ledgerAccountJSON contains the JSON metadata for the struct [LedgerAccount]
type ledgerAccountJSON struct {
	ID             apijson.Field
	Balances       apijson.Field
	CreatedAt      apijson.Field
	Description    apijson.Field
	DiscardedAt    apijson.Field
	LedgerID       apijson.Field
	LedgerableID   apijson.Field
	LedgerableType apijson.Field
	LiveMode       apijson.Field
	LockVersion    apijson.Field
	Metadata       apijson.Field
	Name           apijson.Field
	NormalBalance  apijson.Field
	Object         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LedgerAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account. The posted
// balance is the sum of all posted entries on the account. The pending balance is
// the sum of all pending and posted entries on the account. The available balance
// is the posted incoming entries minus the sum of the pending and posted outgoing
// amounts.
type LedgerAccountBalances struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountBalancesAvailableBalance `json:"available_balance,required"`
	// The inclusive lower bound of the effective_at timestamp for the returned
	// balances.
	EffectiveAtLowerBound string `json:"effective_at_lower_bound,required,nullable" format:"time"`
	// The exclusive upper bound of the effective_at timestamp for the returned
	// balances.
	EffectiveAtUpperBound string `json:"effective_at_upper_bound,required,nullable" format:"time"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountBalancesPostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountBalancesJSON
}

// ledgerAccountBalancesJSON contains the JSON metadata for the struct
// [LedgerAccountBalances]
type ledgerAccountBalancesJSON struct {
	AvailableBalance      apijson.Field
	EffectiveAtLowerBound apijson.Field
	EffectiveAtUpperBound apijson.Field
	PendingBalance        apijson.Field
	PostedBalance         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *LedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountBalancesAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountBalancesAvailableBalanceJSON
}

// ledgerAccountBalancesAvailableBalanceJSON contains the JSON metadata for the
// struct [LedgerAccountBalancesAvailableBalance]
type ledgerAccountBalancesAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountBalancesPendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountBalancesPendingBalanceJSON
}

// ledgerAccountBalancesPendingBalanceJSON contains the JSON metadata for the
// struct [LedgerAccountBalancesPendingBalance]
type ledgerAccountBalancesPendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountBalancesPostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountBalancesPostedBalanceJSON
}

// ledgerAccountBalancesPostedBalanceJSON contains the JSON metadata for the struct
// [LedgerAccountBalancesPostedBalance]
type ledgerAccountBalancesPostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type LedgerAccountLedgerableType string

const (
	LedgerAccountLedgerableTypeExternalAccount LedgerAccountLedgerableType = "external_account"
	LedgerAccountLedgerableTypeInternalAccount LedgerAccountLedgerableType = "internal_account"
)

// The normal balance of the ledger account.
type LedgerAccountNormalBalance string

const (
	LedgerAccountNormalBalanceCredit LedgerAccountNormalBalance = "credit"
	LedgerAccountNormalBalanceDebit  LedgerAccountNormalBalance = "debit"
)

type LedgerAccountNewParams struct {
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[LedgerAccountNewParamsNormalBalance] `json:"normal_balance,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[LedgerAccountNewParamsLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata       param.Field[map[string]string] `json:"metadata"`
	IdempotencyKey param.Field[string]            `header:"Idempotency-Key"`
}

func (r LedgerAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The normal balance of the ledger account.
type LedgerAccountNewParamsNormalBalance string

const (
	LedgerAccountNewParamsNormalBalanceCredit LedgerAccountNewParamsNormalBalance = "credit"
	LedgerAccountNewParamsNormalBalanceDebit  LedgerAccountNewParamsNormalBalance = "debit"
)

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type LedgerAccountNewParamsLedgerableType string

const (
	LedgerAccountNewParamsLedgerableTypeExternalAccount LedgerAccountNewParamsLedgerableType = "external_account"
	LedgerAccountNewParamsLedgerableTypeInternalAccount LedgerAccountNewParamsLedgerableType = "internal_account"
)

type LedgerAccountGetParams struct {
	// Use `balances[effective_at_lower_bound]` and
	// `balances[effective_at_upper_bound]` to get the balances change between the two
	// timestamps. The lower bound is inclusive while the upper bound is exclusive of
	// the provided timestamps. If no value is supplied the balances will be retrieved
	// not including that bound. Use `balances[as_of_lock_version]` to retrieve a
	// balance as of a specific Ledger Account `lock_version`.
	Balances param.Field[LedgerAccountGetParamsBalances] `query:"balances"`
}

// URLQuery serializes [LedgerAccountGetParams]'s query parameters as `url.Values`.
func (r LedgerAccountGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use `balances[effective_at_lower_bound]` and
// `balances[effective_at_upper_bound]` to get the balances change between the two
// timestamps. The lower bound is inclusive while the upper bound is exclusive of
// the provided timestamps. If no value is supplied the balances will be retrieved
// not including that bound. Use `balances[as_of_lock_version]` to retrieve a
// balance as of a specific Ledger Account `lock_version`.
type LedgerAccountGetParamsBalances struct {
	AsOfDate              param.Field[time.Time] `query:"as_of_date" format:"date"`
	AsOfLockVersion       param.Field[int64]     `query:"as_of_lock_version"`
	EffectiveAt           param.Field[time.Time] `query:"effective_at" format:"date-time"`
	EffectiveAtLowerBound param.Field[time.Time] `query:"effective_at_lower_bound" format:"date-time"`
	EffectiveAtUpperBound param.Field[time.Time] `query:"effective_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [LedgerAccountGetParamsBalances]'s query parameters as
// `url.Values`.
func (r LedgerAccountGetParamsBalances) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerAccountUpdateParams struct {
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name"`
}

func (r LedgerAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountListParams struct {
	ID          param.Field[string] `query:"id"`
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Use `balances[effective_at_lower_bound]` and
	// `balances[effective_at_upper_bound]` to get the balances change between the two
	// timestamps. The lower bound is inclusive while the upper bound is exclusive of
	// the provided timestamps. If no value is supplied the balances will be retrieved
	// not including that bound.
	Balances param.Field[LedgerAccountListParamsBalances] `query:"balances"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// created at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// created_at%5Bgt%5D=2000-01-01T12:00:00Z.
	CreatedAt               param.Field[map[string]time.Time] `query:"created_at" format:"date-time"`
	LedgerAccountCategoryID param.Field[string]               `query:"ledger_account_category_id"`
	LedgerID                param.Field[string]               `query:"ledger_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	Name     param.Field[string]            `query:"name"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// updated at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt param.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes [LedgerAccountListParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use `balances[effective_at_lower_bound]` and
// `balances[effective_at_upper_bound]` to get the balances change between the two
// timestamps. The lower bound is inclusive while the upper bound is exclusive of
// the provided timestamps. If no value is supplied the balances will be retrieved
// not including that bound.
type LedgerAccountListParamsBalances struct {
	AsOfDate              param.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt           param.Field[time.Time] `query:"effective_at" format:"date-time"`
	EffectiveAtLowerBound param.Field[time.Time] `query:"effective_at_lower_bound" format:"date-time"`
	EffectiveAtUpperBound param.Field[time.Time] `query:"effective_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [LedgerAccountListParamsBalances]'s query parameters as
// `url.Values`.
func (r LedgerAccountListParamsBalances) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
