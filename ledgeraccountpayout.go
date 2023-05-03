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
func (r *LedgerAccountPayoutService) Retireve(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountPayout, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_payouts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LedgerAccountPayout struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The description of the ledger account payout.
	Description string `json:"description,required,nullable"`
	// The status of the ledger account payout. One of `processing`, `pending`,
	// `posted`, `archiving` or `archived`.
	Status LedgerAccountPayoutStatus `json:"status,required"`
	// The id of the payout ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	PayoutLedgerAccountID string `json:"payout_ledger_account_id,required" format:"uuid"`
	// The id of the funding ledger account that sends to or receives funds from the
	// payout ledger account.
	FundingLedgerAccountID string `json:"funding_ledger_account_id,required" format:"uuid"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account payout. The default value is the created_at
	// timestamp of the ledger account payout.
	EffectiveAtUpperBound string `json:"effective_at_upper_bound,required" format:"time"`
	// The ledger transaction that this payout is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// The amount of the ledger account payout.
	Amount int64 `json:"amount,required,nullable"`
	// The currency of the ledger account payout.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account payout.
	CurrencyExponent int64 `json:"currency_exponent,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     ledgerAccountPayoutJSON
}

// ledgerAccountPayoutJSON contains the JSON metadata for the struct
// [LedgerAccountPayout]
type ledgerAccountPayoutJSON struct {
	ID                     apijson.Field
	Object                 apijson.Field
	LiveMode               apijson.Field
	CreatedAt              apijson.Field
	UpdatedAt              apijson.Field
	Description            apijson.Field
	Status                 apijson.Field
	PayoutLedgerAccountID  apijson.Field
	FundingLedgerAccountID apijson.Field
	EffectiveAtUpperBound  apijson.Field
	LedgerTransactionID    apijson.Field
	Amount                 apijson.Field
	Currency               apijson.Field
	CurrencyExponent       apijson.Field
	Metadata               apijson.Field
	raw                    string
	Extras                 map[string]apijson.Field
}

func (r *LedgerAccountPayout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountPayoutStatus string

const (
	LedgerAccountPayoutStatusArchived   LedgerAccountPayoutStatus = "archived"
	LedgerAccountPayoutStatusArchiving  LedgerAccountPayoutStatus = "archiving"
	LedgerAccountPayoutStatusPending    LedgerAccountPayoutStatus = "pending"
	LedgerAccountPayoutStatusPosted     LedgerAccountPayoutStatus = "posted"
	LedgerAccountPayoutStatusProcessing LedgerAccountPayoutStatus = "processing"
)

type LedgerAccountPayoutNewParams struct {
	// The description of the ledger account payout.
	Description field.Field[string] `json:"description,nullable"`
	// The status of the ledger account payout. It is set to `pending` by default. To
	// post a ledger account payout at creation, use `posted`.
	Status field.Field[LedgerAccountPayoutNewParamsStatus] `json:"status,nullable"`
	// The id of the payout ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	PayoutLedgerAccountID field.Field[string] `json:"payout_ledger_account_id,required" format:"uuid"`
	// The id of the funding ledger account that sends to or receives funds from the
	// payout ledger account.
	FundingLedgerAccountID field.Field[string] `json:"funding_ledger_account_id,required" format:"uuid"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account payout. The default value is the created_at
	// timestamp of the ledger account payout.
	EffectiveAtUpperBound field.Field[string] `json:"effective_at_upper_bound,nullable" format:"time"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountPayoutNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountPayoutNewParamsStatus string

const (
	LedgerAccountPayoutNewParamsStatusPending LedgerAccountPayoutNewParamsStatus = "pending"
	LedgerAccountPayoutNewParamsStatusPosted  LedgerAccountPayoutNewParamsStatus = "posted"
)

type LedgerAccountPayoutUpdateParams struct {
	// The description of the ledger account payout.
	Description field.Field[string] `json:"description,nullable"`
	// To post a pending ledger account payout, use `posted`. To archive a pending
	// ledger transaction, use `archived`.
	Status field.Field[LedgerAccountPayoutUpdateParamsStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountPayoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerAccountPayoutUpdateParamsStatus string

const (
	LedgerAccountPayoutUpdateParamsStatusPosted   LedgerAccountPayoutUpdateParamsStatus = "posted"
	LedgerAccountPayoutUpdateParamsStatusArchived LedgerAccountPayoutUpdateParamsStatus = "archived"
)

type LedgerAccountPayoutListParams struct {
	AfterCursor           field.Field[string] `query:"after_cursor,nullable"`
	PerPage               field.Field[int64]  `query:"per_page"`
	PayoutLedgerAccountID field.Field[string] `query:"payout_ledger_account_id"`
}

// URLQuery serializes [LedgerAccountPayoutListParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountPayoutListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
