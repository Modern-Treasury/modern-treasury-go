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

type LedgerAccountCategoryService struct {
	Options []option.RequestOption
}

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
func (r *LedgerAccountCategoryService) Delete(ctx context.Context, id string, query LedgerAccountCategoryDeleteParams, opts ...option.RequestOption) (res *LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return
}

// Add a ledger account category to an account.
func (r *LedgerAccountCategoryService) AddLedgerAccount(ctx context.Context, id string, ledger_account_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_accounts/%s", id, ledger_account_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Add a ledger account category to a ledger account category.
func (r *LedgerAccountCategoryService) AddNestedCategory(ctx context.Context, id string, sub_category_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_account_categories/%s", id, sub_category_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

// Delete a ledger account category from an account.
func (r *LedgerAccountCategoryService) RemoveLedgerAccount(ctx context.Context, id string, ledger_account_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_accounts/%s", id, ledger_account_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Delete a ledger account category from a ledger account category.
func (r *LedgerAccountCategoryService) RemoveNestedCategory(ctx context.Context, id string, sub_category_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_account_categories/%s", id, sub_category_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type LedgerAccountCategory struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the ledger account category.
	Name string `json:"name,required"`
	// The description of the ledger account category.
	Description string `json:"description,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The id of the ledger that this account category belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// The normal balance of the ledger account category.
	NormalBalance LedgerAccountCategoryNormalBalance `json:"normal_balance,required"`
	// The pending, posted, and available balances for this ledger account category.
	// The posted balance is the sum of all posted entries on the account. The pending
	// balance is the sum of all pending and posted entries on the account. The
	// available balance is the posted incoming entries minus the sum of the pending
	// and posted outgoing amounts.
	Balances LedgerAccountCategoryBalances `json:"balances,required"`
	JSON     LedgerAccountCategoryJSON
}

type LedgerAccountCategoryJSON struct {
	ID            apijson.Metadata
	Object        apijson.Metadata
	LiveMode      apijson.Metadata
	CreatedAt     apijson.Metadata
	UpdatedAt     apijson.Metadata
	DiscardedAt   apijson.Metadata
	Name          apijson.Metadata
	Description   apijson.Metadata
	Metadata      apijson.Metadata
	LedgerID      apijson.Metadata
	NormalBalance apijson.Metadata
	Balances      apijson.Metadata
	raw           string
	Extras        map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccountCategory using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerAccountCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryNormalBalance string

const (
	LedgerAccountCategoryNormalBalanceCredit LedgerAccountCategoryNormalBalance = "credit"
	LedgerAccountCategoryNormalBalanceDebit  LedgerAccountCategoryNormalBalance = "debit"
)

type LedgerAccountCategoryBalances struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountCategoryBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountCategoryBalancesPostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountCategoryBalancesAvailableBalance `json:"available_balance,required"`
	JSON             LedgerAccountCategoryBalancesJSON
}

type LedgerAccountCategoryBalancesJSON struct {
	PendingBalance   apijson.Metadata
	PostedBalance    apijson.Metadata
	AvailableBalance apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerAccountCategoryBalances
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *LedgerAccountCategoryBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountCategoryBalancesPendingBalanceJSON
}

type LedgerAccountCategoryBalancesPendingBalanceJSON struct {
	Credits          apijson.Metadata
	Debits           apijson.Metadata
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	CurrencyExponent apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountCategoryBalancesPendingBalance using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountCategoryBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountCategoryBalancesPostedBalanceJSON
}

type LedgerAccountCategoryBalancesPostedBalanceJSON struct {
	Credits          apijson.Metadata
	Debits           apijson.Metadata
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	CurrencyExponent apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountCategoryBalancesPostedBalance using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountCategoryBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             LedgerAccountCategoryBalancesAvailableBalanceJSON
}

type LedgerAccountCategoryBalancesAvailableBalanceJSON struct {
	Credits          apijson.Metadata
	Debits           apijson.Metadata
	Amount           apijson.Metadata
	Currency         apijson.Metadata
	CurrencyExponent apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// LedgerAccountCategoryBalancesAvailableBalance using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *LedgerAccountCategoryBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountCategoryNewParams struct {
	// The name of the ledger account category.
	Name field.Field[string] `json:"name,required"`
	// The description of the ledger account category.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The currency of the ledger account category.
	Currency field.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account category.
	CurrencyExponent field.Field[int64] `json:"currency_exponent,nullable"`
	// The id of the ledger that this account category belongs to.
	LedgerID field.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The normal balance of the ledger account category.
	NormalBalance field.Field[LedgerAccountCategoryNewParamsNormalBalance] `json:"normal_balance,required"`
}

// MarshalJSON serializes LedgerAccountCategoryNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r LedgerAccountCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountCategoryNewParamsNormalBalance string

const (
	LedgerAccountCategoryNewParamsNormalBalanceCredit LedgerAccountCategoryNewParamsNormalBalance = "credit"
	LedgerAccountCategoryNewParamsNormalBalanceDebit  LedgerAccountCategoryNewParamsNormalBalance = "debit"
)

type LedgerAccountCategoryGetParams struct {
	// For example, if you want the balances as of a particular effective date
	// (YYYY-MM-DD), the encoded query string would be
	// balances%5Bas_of_date%5D=2000-12-31. The balances as of a date are exclusive of
	// entries with that exact date.
	Balances field.Field[LedgerAccountCategoryGetParamsBalances] `query:"balances"`
}

// URLQuery serializes LedgerAccountCategoryGetParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryGetParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountCategoryGetParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountCategoryGetParamsBalances into a url.Values of
// the query parameters associated with this value
func (r LedgerAccountCategoryGetParamsBalances) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountCategoryUpdateParams struct {
	// The name of the ledger account category.
	Name field.Field[string] `json:"name"`
	// The description of the ledger account category.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerAccountCategoryUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r LedgerAccountCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes LedgerAccountCategoryUpdateParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryUpdateParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountCategoryUpdateParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountCategoryUpdateParamsBalances into a url.Values
// of the query parameters associated with this value
func (r LedgerAccountCategoryUpdateParamsBalances) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountCategoryListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	Name     field.Field[string]            `query:"name"`
	LedgerID field.Field[string]            `query:"ledger_id"`
	// Query categories that are nested underneath a parent category
	ParentLedgerAccountCategoryID field.Field[string] `query:"parent_ledger_account_category_id"`
}

// URLQuery serializes LedgerAccountCategoryListParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountCategoryDeleteParams struct {
	// For example, if you want the balances as of a particular effective date
	// (YYYY-MM-DD), the encoded query string would be
	// balances%5Bas_of_date%5D=2000-12-31. The balances as of a date are exclusive of
	// entries with that exact date.
	Balances field.Field[LedgerAccountCategoryDeleteParamsBalances] `query:"balances"`
}

// URLQuery serializes LedgerAccountCategoryDeleteParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryDeleteParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountCategoryDeleteParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountCategoryDeleteParamsBalances into a url.Values
// of the query parameters associated with this value
func (r LedgerAccountCategoryDeleteParamsBalances) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
