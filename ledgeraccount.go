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
func (r *LedgerAccountService) New(ctx context.Context, body LedgerAccountNewParams, opts ...option.RequestOption) (res *LedgerAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the ledger account.
	Name string `json:"name,required"`
	// The description of the ledger account.
	Description string `json:"description,required,nullable"`
	// The normal balance of the ledger account.
	NormalBalance LedgerAccountNormalBalance `json:"normal_balance,required"`
	// The pending, posted, and available balances for this ledger account. The posted
	// balance is the sum of all posted entries on the account. The pending balance is
	// the sum of all pending and posted entries on the account. The available balance
	// is the posted incoming entries minus the sum of the pending and posted outgoing
	// amounts.
	Balances LedgerAccountBalances `json:"balances,required"`
	// Lock version of the ledger account.
	LockVersion int64 `json:"lock_version,required"`
	// The id of the ledger that this account belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType LedgerAccountLedgerableType `json:"ledgerable_type,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     ledgerAccountJSON
}

// ledgerAccountJSON contains the JSON metadata for the struct [LedgerAccount]
type ledgerAccountJSON struct {
	ID             apijson.Field
	Object         apijson.Field
	LiveMode       apijson.Field
	CreatedAt      apijson.Field
	UpdatedAt      apijson.Field
	DiscardedAt    apijson.Field
	Name           apijson.Field
	Description    apijson.Field
	NormalBalance  apijson.Field
	Balances       apijson.Field
	LockVersion    apijson.Field
	LedgerID       apijson.Field
	LedgerableID   apijson.Field
	LedgerableType apijson.Field
	Metadata       apijson.Field
	raw            string
	Extras         map[string]apijson.Field
}

func (r *LedgerAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountNormalBalance string

const (
	LedgerAccountNormalBalanceCredit LedgerAccountNormalBalance = "credit"
	LedgerAccountNormalBalanceDebit  LedgerAccountNormalBalance = "debit"
)

// The pending, posted, and available balances for this ledger account. The posted
// balance is the sum of all posted entries on the account. The pending balance is
// the sum of all pending and posted entries on the account. The available balance
// is the posted incoming entries minus the sum of the pending and posted outgoing
// amounts.
type LedgerAccountBalances struct {
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
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountBalancesAvailableBalance `json:"available_balance,required"`
	JSON             ledgerAccountBalancesJSON
}

// ledgerAccountBalancesJSON contains the JSON metadata for the struct
// [LedgerAccountBalances]
type ledgerAccountBalancesJSON struct {
	EffectiveAtLowerBound apijson.Field
	EffectiveAtUpperBound apijson.Field
	PendingBalance        apijson.Field
	PostedBalance         apijson.Field
	AvailableBalance      apijson.Field
	raw                   string
	Extras                map[string]apijson.Field
}

func (r *LedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountBalancesPendingBalanceJSON
}

// ledgerAccountBalancesPendingBalanceJSON contains the JSON metadata for the
// struct [LedgerAccountBalancesPendingBalance]
type ledgerAccountBalancesPendingBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountBalancesPostedBalanceJSON
}

// ledgerAccountBalancesPostedBalanceJSON contains the JSON metadata for the struct
// [LedgerAccountBalancesPostedBalance]
type ledgerAccountBalancesPostedBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountBalancesAvailableBalanceJSON
}

// ledgerAccountBalancesAvailableBalanceJSON contains the JSON metadata for the
// struct [LedgerAccountBalancesAvailableBalance]
type ledgerAccountBalancesAvailableBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountLedgerableType string

const (
	LedgerAccountLedgerableTypeExternalAccount LedgerAccountLedgerableType = "external_account"
	LedgerAccountLedgerableTypeInternalAccount LedgerAccountLedgerableType = "internal_account"
)

type LedgerAccountNewParams struct {
	// The name of the ledger account.
	Name field.Field[string] `json:"name,required"`
	// The description of the ledger account.
	Description field.Field[string] `json:"description,nullable"`
	// The normal balance of the ledger account.
	NormalBalance field.Field[LedgerAccountNewParamsNormalBalance] `json:"normal_balance,required"`
	// The id of the ledger that this account belongs to.
	LedgerID field.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The currency of the ledger account.
	Currency field.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent field.Field[int64] `json:"currency_exponent,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountNewParamsNormalBalance string

const (
	LedgerAccountNewParamsNormalBalanceCredit LedgerAccountNewParamsNormalBalance = "credit"
	LedgerAccountNewParamsNormalBalanceDebit  LedgerAccountNewParamsNormalBalance = "debit"
)

type LedgerAccountGetParams struct {
	// Use balances[effective_at_lower_bound] and balances[effective_at_upper_bound] to
	// get the balances change between the two timestamps. The lower bound is inclusive
	// while the upper bound is exclusive of the provided timestamps. If no value is
	// supplied the balances will be retrieved not including that bound.
	Balances field.Field[LedgerAccountGetParamsBalances] `query:"balances"`
}

// URLQuery serializes [LedgerAccountGetParams]'s query parameters as `url.Values`.
func (r LedgerAccountGetParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// Use balances[effective_at_lower_bound] and balances[effective_at_upper_bound] to
// get the balances change between the two timestamps. The lower bound is inclusive
// while the upper bound is exclusive of the provided timestamps. If no value is
// supplied the balances will be retrieved not including that bound.
type LedgerAccountGetParamsBalances struct {
	AsOfDate              field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt           field.Field[time.Time] `query:"effective_at" format:"date-time"`
	EffectiveAtLowerBound field.Field[time.Time] `query:"effective_at_lower_bound" format:"date-time"`
	EffectiveAtUpperBound field.Field[time.Time] `query:"effective_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [LedgerAccountGetParamsBalances]'s query parameters as
// `url.Values`.
func (r LedgerAccountGetParamsBalances) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LedgerAccountUpdateParams struct {
	// The name of the ledger account.
	Name field.Field[string] `json:"name"`
	// The description of the ledger account.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	ID       field.Field[string]            `query:"id"`
	Name     field.Field[string]            `query:"name"`
	LedgerID field.Field[string]            `query:"ledger_id"`
	// Use balances[effective_at_lower_bound] and balances[effective_at_upper_bound] to
	// get the balances change between the two timestamps. The lower bound is inclusive
	// while the upper bound is exclusive of the provided timestamps. If no value is
	// supplied the balances will be retrieved not including that bound.
	Balances field.Field[LedgerAccountListParamsBalances] `query:"balances"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt               field.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
	LedgerAccountCategoryID field.Field[string]               `query:"ledger_account_category_id"`
}

// URLQuery serializes [LedgerAccountListParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

// Use balances[effective_at_lower_bound] and balances[effective_at_upper_bound] to
// get the balances change between the two timestamps. The lower bound is inclusive
// while the upper bound is exclusive of the provided timestamps. If no value is
// supplied the balances will be retrieved not including that bound.
type LedgerAccountListParamsBalances struct {
	AsOfDate              field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt           field.Field[time.Time] `query:"effective_at" format:"date-time"`
	EffectiveAtLowerBound field.Field[time.Time] `query:"effective_at_lower_bound" format:"date-time"`
	EffectiveAtUpperBound field.Field[time.Time] `query:"effective_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [LedgerAccountListParamsBalances]'s query parameters as
// `url.Values`.
func (r LedgerAccountListParamsBalances) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
