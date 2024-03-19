// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// LedgerAccountPayoutService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLedgerAccountPayoutService] method instead.
type LedgerAccountPayoutService struct {
	Options []option.RequestOption
}

// NewLedgerAccountPayoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLedgerAccountPayoutService(opts ...option.RequestOption) (r *LedgerAccountPayoutService) {
	r = &LedgerAccountPayoutService{}
	r.Options = opts
	return
}

// Create a ledger account payout.
func (r *LedgerAccountPayoutService) New(ctx context.Context, body LedgerAccountPayoutNewParams, opts ...option.RequestOption) (res *LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_payouts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger account payout.
func (r *LedgerAccountPayoutService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the details of a ledger account payout.
func (r *LedgerAccountPayoutService) Update(ctx context.Context, id string, body LedgerAccountPayoutUpdateParams, opts ...option.RequestOption) (res *LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledger account payouts.
func (r *LedgerAccountPayoutService) List(ctx context.Context, query LedgerAccountPayoutListParams, opts ...option.RequestOption) (res *shared.Page[LedgerAccountPayout], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_account_payouts"
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

// Get a list of ledger account payouts.
func (r *LedgerAccountPayoutService) ListAutoPaging(ctx context.Context, query LedgerAccountPayoutListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerAccountPayout] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Get details on a single ledger account payout.
//
// Deprecated: use `Get` instead
func (r *LedgerAccountPayoutService) Retireve(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LedgerAccountPayout struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of the ledger account payout.
	Amount    int64     `json:"amount,required,nullable"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the ledger account payout.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account payout.
	CurrencyExponent int64 `json:"currency_exponent,required,nullable"`
	// The description of the ledger account payout.
	Description string `json:"description,required,nullable"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account payout. The default value is the created_at
	// timestamp of the ledger account payout.
	EffectiveAtUpperBound time.Time `json:"effective_at_upper_bound,required" format:"date-time"`
	// The id of the funding ledger account that sends to or receives funds from the
	// payout ledger account.
	FundingLedgerAccountID string `json:"funding_ledger_account_id,required" format:"uuid"`
	// The id of the ledger that this ledger account payout belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// The id of the ledger transaction that this payout is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The direction of the ledger entry with the payout_ledger_account.
	PayoutEntryDirection string `json:"payout_entry_direction,required,nullable"`
	// The id of the payout ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	PayoutLedgerAccountID string `json:"payout_ledger_account_id,required" format:"uuid"`
	// The status of the ledger account payout. One of `processing`, `pending`,
	// `posted`, `archiving`, `archived`.
	Status    LedgerAccountPayoutStatus `json:"status,required"`
	UpdatedAt time.Time                 `json:"updated_at,required" format:"date-time"`
	JSON      ledgerAccountPayoutJSON   `json:"-"`
}

// ledgerAccountPayoutJSON contains the JSON metadata for the struct
// [LedgerAccountPayout]
type ledgerAccountPayoutJSON struct {
	ID                     apijson.Field
	Amount                 apijson.Field
	CreatedAt              apijson.Field
	Currency               apijson.Field
	CurrencyExponent       apijson.Field
	Description            apijson.Field
	EffectiveAtUpperBound  apijson.Field
	FundingLedgerAccountID apijson.Field
	LedgerID               apijson.Field
	LedgerTransactionID    apijson.Field
	LiveMode               apijson.Field
	Metadata               apijson.Field
	Object                 apijson.Field
	PayoutEntryDirection   apijson.Field
	PayoutLedgerAccountID  apijson.Field
	Status                 apijson.Field
	UpdatedAt              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *LedgerAccountPayout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountPayoutJSON) RawJSON() string {
	return r.raw
}

// The status of the ledger account payout. One of `processing`, `pending`,
// `posted`, `archiving`, `archived`.
type LedgerAccountPayoutStatus string

const (
	LedgerAccountPayoutStatusArchived   LedgerAccountPayoutStatus = "archived"
	LedgerAccountPayoutStatusArchiving  LedgerAccountPayoutStatus = "archiving"
	LedgerAccountPayoutStatusPending    LedgerAccountPayoutStatus = "pending"
	LedgerAccountPayoutStatusPosted     LedgerAccountPayoutStatus = "posted"
	LedgerAccountPayoutStatusProcessing LedgerAccountPayoutStatus = "processing"
)

func (r LedgerAccountPayoutStatus) IsKnown() bool {
	switch r {
	case LedgerAccountPayoutStatusArchived, LedgerAccountPayoutStatusArchiving, LedgerAccountPayoutStatusPending, LedgerAccountPayoutStatusPosted, LedgerAccountPayoutStatusProcessing:
		return true
	}
	return false
}

type LedgerAccountPayoutNewParams struct {
	// The id of the funding ledger account that sends to or receives funds from the
	// payout ledger account.
	FundingLedgerAccountID param.Field[string] `json:"funding_ledger_account_id,required" format:"uuid"`
	// The id of the payout ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	PayoutLedgerAccountID param.Field[string] `json:"payout_ledger_account_id,required" format:"uuid"`
	// If true, the payout amount and payout_entry_direction will bring the payout
	// ledger account's balance closer to zero, even if the balance is negative.
	AllowEitherDirection param.Field[bool] `json:"allow_either_direction"`
	// The description of the ledger account payout.
	Description param.Field[string] `json:"description"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account payout. The default value is the created_at
	// timestamp of the ledger account payout.
	EffectiveAtUpperBound param.Field[time.Time] `json:"effective_at_upper_bound" format:"date-time"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// It is set to `false` by default. It should be set to `true` when migrating
	// existing payouts.
	SkipPayoutLedgerTransaction param.Field[bool] `json:"skip_payout_ledger_transaction"`
	// The status of the ledger account payout. It is set to `pending` by default. To
	// post a ledger account payout at creation, use `posted`.
	Status param.Field[LedgerAccountPayoutNewParamsStatus] `json:"status"`
}

func (r LedgerAccountPayoutNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the ledger account payout. It is set to `pending` by default. To
// post a ledger account payout at creation, use `posted`.
type LedgerAccountPayoutNewParamsStatus string

const (
	LedgerAccountPayoutNewParamsStatusPending LedgerAccountPayoutNewParamsStatus = "pending"
	LedgerAccountPayoutNewParamsStatusPosted  LedgerAccountPayoutNewParamsStatus = "posted"
)

func (r LedgerAccountPayoutNewParamsStatus) IsKnown() bool {
	switch r {
	case LedgerAccountPayoutNewParamsStatusPending, LedgerAccountPayoutNewParamsStatusPosted:
		return true
	}
	return false
}

type LedgerAccountPayoutUpdateParams struct {
	// The description of the ledger account payout.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a pending ledger account payout, use `posted`. To archive a pending
	// ledger transaction, use `archived`.
	Status param.Field[LedgerAccountPayoutUpdateParamsStatus] `json:"status"`
}

func (r LedgerAccountPayoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// To post a pending ledger account payout, use `posted`. To archive a pending
// ledger transaction, use `archived`.
type LedgerAccountPayoutUpdateParamsStatus string

const (
	LedgerAccountPayoutUpdateParamsStatusPosted   LedgerAccountPayoutUpdateParamsStatus = "posted"
	LedgerAccountPayoutUpdateParamsStatusArchived LedgerAccountPayoutUpdateParamsStatus = "archived"
)

func (r LedgerAccountPayoutUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LedgerAccountPayoutUpdateParamsStatusPosted, LedgerAccountPayoutUpdateParamsStatusArchived:
		return true
	}
	return false
}

type LedgerAccountPayoutListParams struct {
	// If you have specific IDs to retrieve in bulk, you can pass them as query
	// parameters delimited with `id[]=`, for example `?id[]=123&id[]=abc`.
	ID          param.Field[[]string] `query:"id"`
	AfterCursor param.Field[string]   `query:"after_cursor"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata              param.Field[map[string]string] `query:"metadata"`
	PayoutEntryDirection  param.Field[string]            `query:"payout_entry_direction"`
	PayoutLedgerAccountID param.Field[string]            `query:"payout_ledger_account_id"`
	PerPage               param.Field[int64]             `query:"per_page"`
}

// URLQuery serializes [LedgerAccountPayoutListParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountPayoutListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
