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

// LedgerAccountCategoryService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLedgerAccountCategoryService] method instead.
type LedgerAccountCategoryService struct {
	Options []option.RequestOption
}

// NewLedgerAccountCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLedgerAccountCategoryService(opts ...option.RequestOption) (r *LedgerAccountCategoryService) {
	r = &LedgerAccountCategoryService{}
	r.Options = opts
	return
}

// Create a ledger account category.
func (r *LedgerAccountCategoryService) New(ctx context.Context, body LedgerAccountCategoryNewParams, opts ...option.RequestOption) (res *LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the details on a single ledger account category.
func (r *LedgerAccountCategoryService) Get(ctx context.Context, id string, query LedgerAccountCategoryGetParams, opts ...option.RequestOption) (res *LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update the details of a ledger account category.
func (r *LedgerAccountCategoryService) Update(ctx context.Context, id string, body LedgerAccountCategoryUpdateParams, opts ...option.RequestOption) (res *LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledger account categories.
func (r *LedgerAccountCategoryService) List(ctx context.Context, query LedgerAccountCategoryListParams, opts ...option.RequestOption) (res *shared.Page[LedgerAccountCategory], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_account_categories"
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

// Get a list of ledger account categories.
func (r *LedgerAccountCategoryService) ListAutoPaging(ctx context.Context, query LedgerAccountCategoryListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerAccountCategory] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger account category.
func (r *LedgerAccountCategoryService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a ledger account to a ledger account category.
func (r *LedgerAccountCategoryService) AddLedgerAccount(ctx context.Context, id string, ledgerAccountID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_accounts/%s", id, ledgerAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Add a ledger account category to a ledger account category.
func (r *LedgerAccountCategoryService) AddNestedCategory(ctx context.Context, id string, subCategoryID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_account_categories/%s", id, subCategoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Remove a ledger account from a ledger account category.
func (r *LedgerAccountCategoryService) RemoveLedgerAccount(ctx context.Context, id string, ledgerAccountID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_accounts/%s", id, ledgerAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete a ledger account category from a ledger account category.
func (r *LedgerAccountCategoryService) RemoveNestedCategory(ctx context.Context, id string, subCategoryID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_account_categories/%s", id, subCategoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type LedgerAccountCategory struct {
	ID string `json:"id,required" format:"uuid"`
	// The pending, posted, and available balances for this ledger account category.
	// The posted balance is the sum of all posted entries on the account. The pending
	// balance is the sum of all pending and posted entries on the account. The
	// available balance is the posted incoming entries minus the sum of the pending
	// and posted outgoing amounts.
	Balances  LedgerAccountCategoryBalances `json:"balances,required"`
	CreatedAt time.Time                     `json:"created_at,required" format:"date-time"`
	// The description of the ledger account category.
	Description string    `json:"description,required,nullable"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The id of the ledger that this account category belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The name of the ledger account category.
	Name string `json:"name,required"`
	// The normal balance of the ledger account category.
	NormalBalance LedgerAccountCategoryNormalBalance `json:"normal_balance,required"`
	Object        string                             `json:"object,required"`
	UpdatedAt     time.Time                          `json:"updated_at,required" format:"date-time"`
	JSON          ledgerAccountCategoryJSON
}

// ledgerAccountCategoryJSON contains the JSON metadata for the struct
// [LedgerAccountCategory]
type ledgerAccountCategoryJSON struct {
	ID            apijson.Field
	Balances      apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	DiscardedAt   apijson.Field
	LedgerID      apijson.Field
	LiveMode      apijson.Field
	Metadata      apijson.Field
	Name          apijson.Field
	NormalBalance apijson.Field
	Object        apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LedgerAccountCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account category.
// The posted balance is the sum of all posted entries on the account. The pending
// balance is the sum of all pending and posted entries on the account. The
// available balance is the posted incoming entries minus the sum of the pending
// and posted outgoing amounts.
type LedgerAccountCategoryBalances struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountCategoryBalancesAvailableBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountCategoryBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountCategoryBalancesPostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountCategoryBalancesJSON
}

// ledgerAccountCategoryBalancesJSON contains the JSON metadata for the struct
// [LedgerAccountCategoryBalances]
type ledgerAccountCategoryBalancesJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountCategoryBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountCategoryBalancesAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountCategoryBalancesAvailableBalanceJSON
}

// ledgerAccountCategoryBalancesAvailableBalanceJSON contains the JSON metadata for
// the struct [LedgerAccountCategoryBalancesAvailableBalance]
type ledgerAccountCategoryBalancesAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountCategoryBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountCategoryBalancesPendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountCategoryBalancesPendingBalanceJSON
}

// ledgerAccountCategoryBalancesPendingBalanceJSON contains the JSON metadata for
// the struct [LedgerAccountCategoryBalancesPendingBalance]
type ledgerAccountCategoryBalancesPendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountCategoryBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountCategoryBalancesPostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountCategoryBalancesPostedBalanceJSON
}

// ledgerAccountCategoryBalancesPostedBalanceJSON contains the JSON metadata for
// the struct [LedgerAccountCategoryBalancesPostedBalance]
type ledgerAccountCategoryBalancesPostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountCategoryBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The normal balance of the ledger account category.
type LedgerAccountCategoryNormalBalance string

const (
	LedgerAccountCategoryNormalBalanceCredit LedgerAccountCategoryNormalBalance = "credit"
	LedgerAccountCategoryNormalBalanceDebit  LedgerAccountCategoryNormalBalance = "debit"
)

type LedgerAccountCategoryNewParams struct {
	// The currency of the ledger account category.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account category belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account category.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account category.
	NormalBalance param.Field[LedgerAccountCategoryNewParamsNormalBalance] `json:"normal_balance,required"`
	// The currency exponent of the ledger account category.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account category.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The normal balance of the ledger account category.
type LedgerAccountCategoryNewParamsNormalBalance string

const (
	LedgerAccountCategoryNewParamsNormalBalanceCredit LedgerAccountCategoryNewParamsNormalBalance = "credit"
	LedgerAccountCategoryNewParamsNormalBalanceDebit  LedgerAccountCategoryNewParamsNormalBalance = "debit"
)

type LedgerAccountCategoryGetParams struct {
	// For example, if you want the balances as of a particular time (ISO8601), the
	// encoded query string would be `balances%5Beffective_at%5D=2000-12-31T12:00:00Z`.
	// The balances as of a time are inclusive of entries with that exact time.
	Balances param.Field[LedgerAccountCategoryGetParamsBalances] `query:"balances"`
}

// URLQuery serializes [LedgerAccountCategoryGetParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountCategoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For example, if you want the balances as of a particular time (ISO8601), the
// encoded query string would be `balances%5Beffective_at%5D=2000-12-31T12:00:00Z`.
// The balances as of a time are inclusive of entries with that exact time.
type LedgerAccountCategoryGetParamsBalances struct {
	AsOfDate    param.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt param.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes [LedgerAccountCategoryGetParamsBalances]'s query parameters
// as `url.Values`.
func (r LedgerAccountCategoryGetParamsBalances) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerAccountCategoryUpdateParams struct {
	// The description of the ledger account category.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The name of the ledger account category.
	Name param.Field[string] `json:"name"`
}

func (r LedgerAccountCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountCategoryListParams struct {
	// If you have specific IDs to retrieve in bulk, you can pass them as query
	// parameters delimited with `id[]=`, for example `?id[]=123&id[]=abc`.
	ID          param.Field[[]string] `query:"id"`
	AfterCursor param.Field[string]   `query:"after_cursor"`
	// For example, if you want the balances as of a particular time (ISO8601), the
	// encoded query string would be `balances%5Beffective_at%5D=2000-12-31T12:00:00Z`.
	// The balances as of a time are inclusive of entries with that exact time.
	Balances param.Field[LedgerAccountCategoryListParamsBalances] `query:"balances"`
	// Query categories which contain a ledger account directly or through child
	// categories.
	LedgerAccountID param.Field[string] `query:"ledger_account_id"`
	LedgerID        param.Field[string] `query:"ledger_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	Name     param.Field[string]            `query:"name"`
	// Query categories that are nested underneath a parent category
	ParentLedgerAccountCategoryID param.Field[string] `query:"parent_ledger_account_category_id"`
	PerPage                       param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [LedgerAccountCategoryListParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For example, if you want the balances as of a particular time (ISO8601), the
// encoded query string would be `balances%5Beffective_at%5D=2000-12-31T12:00:00Z`.
// The balances as of a time are inclusive of entries with that exact time.
type LedgerAccountCategoryListParamsBalances struct {
	EffectiveAt param.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes [LedgerAccountCategoryListParamsBalances]'s query parameters
// as `url.Values`.
func (r LedgerAccountCategoryListParamsBalances) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
