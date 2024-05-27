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
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// LedgerAccountBalanceMonitorService contains methods and other services that help
// with interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerAccountBalanceMonitorService] method instead.
type LedgerAccountBalanceMonitorService struct {
	Options []option.RequestOption
}

// NewLedgerAccountBalanceMonitorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLedgerAccountBalanceMonitorService(opts ...option.RequestOption) (r *LedgerAccountBalanceMonitorService) {
	r = &LedgerAccountBalanceMonitorService{}
	r.Options = opts
	return
}

// Create a ledger account balance monitor.
func (r *LedgerAccountBalanceMonitorService) New(ctx context.Context, body LedgerAccountBalanceMonitorNewParams, opts ...option.RequestOption) (res *LedgerAccountBalanceMonitor, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_balance_monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger account balance monitor.
func (r *LedgerAccountBalanceMonitorService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountBalanceMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_balance_monitors/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a ledger account balance monitor.
func (r *LedgerAccountBalanceMonitorService) Update(ctx context.Context, id string, body LedgerAccountBalanceMonitorUpdateParams, opts ...option.RequestOption) (res *LedgerAccountBalanceMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_balance_monitors/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledger account balance monitors.
func (r *LedgerAccountBalanceMonitorService) List(ctx context.Context, query LedgerAccountBalanceMonitorListParams, opts ...option.RequestOption) (res *pagination.Page[LedgerAccountBalanceMonitor], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_account_balance_monitors"
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

// Get a list of ledger account balance monitors.
func (r *LedgerAccountBalanceMonitorService) ListAutoPaging(ctx context.Context, query LedgerAccountBalanceMonitorListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LedgerAccountBalanceMonitor] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger account balance monitor.
func (r *LedgerAccountBalanceMonitorService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountBalanceMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_balance_monitors/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LedgerAccountBalanceMonitor struct {
	ID string `json:"id,required" format:"uuid"`
	// Describes the condition that must be satisfied for the monitor to be triggered.
	AlertCondition LedgerAccountBalanceMonitorAlertCondition `json:"alert_condition,required"`
	CreatedAt      time.Time                                 `json:"created_at,required" format:"date-time"`
	// The ledger account's balances and the monitor state as of the current ledger
	// account lock version.
	CurrentLedgerAccountBalanceState LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceState `json:"current_ledger_account_balance_state,required"`
	// An optional, free-form description for internal use.
	Description string    `json:"description,required,nullable"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The ledger account associated with this balance monitor.
	LedgerAccountID string `json:"ledger_account_id,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata  map[string]string               `json:"metadata,required"`
	Object    string                          `json:"object,required"`
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	JSON      ledgerAccountBalanceMonitorJSON `json:"-"`
}

// ledgerAccountBalanceMonitorJSON contains the JSON metadata for the struct
// [LedgerAccountBalanceMonitor]
type ledgerAccountBalanceMonitorJSON struct {
	ID                               apijson.Field
	AlertCondition                   apijson.Field
	CreatedAt                        apijson.Field
	CurrentLedgerAccountBalanceState apijson.Field
	Description                      apijson.Field
	DiscardedAt                      apijson.Field
	LedgerAccountID                  apijson.Field
	LiveMode                         apijson.Field
	Metadata                         apijson.Field
	Object                           apijson.Field
	UpdatedAt                        apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorJSON) RawJSON() string {
	return r.raw
}

// Describes the condition that must be satisfied for the monitor to be triggered.
type LedgerAccountBalanceMonitorAlertCondition struct {
	// One of `available_balance_amount`, `pending_balance_amount`,
	// `posted_balance_amount`, `ledger_account_lock_version`.
	Field string `json:"field,required"`
	// A logical operator to compare the `field` against the `value`. One of
	// `less_than`, `less_than_or_equals`, `equals`, `greater_than_or_equals`,
	// `greater_than`.
	Operator string `json:"operator,required"`
	// The monitor's `current_ledger_account_balance_state.triggered` will be `true`
	// when comparing the `field` to this integer value using the `operator` is
	// logically true.
	Value int64                                         `json:"value,required"`
	JSON  ledgerAccountBalanceMonitorAlertConditionJSON `json:"-"`
}

// ledgerAccountBalanceMonitorAlertConditionJSON contains the JSON metadata for the
// struct [LedgerAccountBalanceMonitorAlertCondition]
type ledgerAccountBalanceMonitorAlertConditionJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitorAlertCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorAlertConditionJSON) RawJSON() string {
	return r.raw
}

// The ledger account's balances and the monitor state as of the current ledger
// account lock version.
type LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceState struct {
	Balances LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalances `json:"balances,required"`
	// The current lock version of the ledger account.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required"`
	// If `true`, the ledger account's balances satisfy the `alert_condition` at this
	// lock version.
	Triggered bool                                                            `json:"triggered,required"`
	JSON      ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateJSON `json:"-"`
}

// ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateJSON contains the
// JSON metadata for the struct
// [LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceState]
type ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateJSON struct {
	Balances                 apijson.Field
	LedgerAccountLockVersion apijson.Field
	Triggered                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateJSON) RawJSON() string {
	return r.raw
}

type LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalances struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesJSON          `json:"-"`
}

// ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesJSON contains
// the JSON metadata for the struct
// [LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalances]
type ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesJSON) RawJSON() string {
	return r.raw
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64                                                                                   `json:"currency_exponent,required"`
	Debits           int64                                                                                   `json:"debits,required"`
	JSON             ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalanceJSON `json:"-"`
}

// ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalanceJSON
// contains the JSON metadata for the struct
// [LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalance]
type ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesAvailableBalanceJSON) RawJSON() string {
	return r.raw
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64                                                                                 `json:"currency_exponent,required"`
	Debits           int64                                                                                 `json:"debits,required"`
	JSON             ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalanceJSON `json:"-"`
}

// ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalanceJSON
// contains the JSON metadata for the struct
// [LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalance]
type ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPendingBalanceJSON) RawJSON() string {
	return r.raw
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64                                                                                `json:"currency_exponent,required"`
	Debits           int64                                                                                `json:"debits,required"`
	JSON             ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalanceJSON `json:"-"`
}

// ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalanceJSON
// contains the JSON metadata for the struct
// [LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalance]
type ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountBalanceMonitorCurrentLedgerAccountBalanceStateBalancesPostedBalanceJSON) RawJSON() string {
	return r.raw
}

type LedgerAccountBalanceMonitorNewParams struct {
	// Describes the condition that must be satisfied for the monitor to be triggered.
	AlertCondition param.Field[LedgerAccountBalanceMonitorNewParamsAlertCondition] `json:"alert_condition,required"`
	// The ledger account associated with this balance monitor.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required"`
	// An optional, free-form description for internal use.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountBalanceMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Describes the condition that must be satisfied for the monitor to be triggered.
type LedgerAccountBalanceMonitorNewParamsAlertCondition struct {
	// One of `available_balance_amount`, `pending_balance_amount`,
	// `posted_balance_amount`, `ledger_account_lock_version`.
	Field param.Field[string] `json:"field,required"`
	// A logical operator to compare the `field` against the `value`. One of
	// `less_than`, `less_than_or_equals`, `equals`, `greater_than_or_equals`,
	// `greater_than`.
	Operator param.Field[string] `json:"operator,required"`
	// The monitor's `current_ledger_account_balance_state.triggered` will be `true`
	// when comparing the `field` to this integer value using the `operator` is
	// logically true.
	Value param.Field[int64] `json:"value,required"`
}

func (r LedgerAccountBalanceMonitorNewParamsAlertCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountBalanceMonitorUpdateParams struct {
	// An optional, free-form description for internal use.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountBalanceMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountBalanceMonitorListParams struct {
	// If you have specific IDs to retrieve in bulk, you can pass them as query
	// parameters delimited with `id[]=`, for example `?id[]=123&id[]=abc`.
	ID          param.Field[[]string] `query:"id"`
	AfterCursor param.Field[string]   `query:"after_cursor"`
	// Query the balance monitors for a single ledger account.
	LedgerAccountID param.Field[string] `query:"ledger_account_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
}

// URLQuery serializes [LedgerAccountBalanceMonitorListParams]'s query parameters
// as `url.Values`.
func (r LedgerAccountBalanceMonitorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
