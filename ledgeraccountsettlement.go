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
)

// LedgerAccountSettlementService contains methods and other services that help
// with interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerAccountSettlementService] method instead.
type LedgerAccountSettlementService struct {
	Options        []option.RequestOption
	AccountEntries *LedgerAccountSettlementAccountEntryService
}

// NewLedgerAccountSettlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLedgerAccountSettlementService(opts ...option.RequestOption) (r *LedgerAccountSettlementService) {
	r = &LedgerAccountSettlementService{}
	r.Options = opts
	r.AccountEntries = NewLedgerAccountSettlementAccountEntryService(opts...)
	return
}

// Create a ledger account settlement.
func (r *LedgerAccountSettlementService) New(ctx context.Context, body LedgerAccountSettlementNewParams, opts ...option.RequestOption) (res *LedgerAccountSettlement, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_settlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger account settlement.
func (r *LedgerAccountSettlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountSettlement, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_settlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the details of a ledger account settlement.
func (r *LedgerAccountSettlementService) Update(ctx context.Context, id string, body LedgerAccountSettlementUpdateParams, opts ...option.RequestOption) (res *LedgerAccountSettlement, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_account_settlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledger account settlements.
func (r *LedgerAccountSettlementService) List(ctx context.Context, query LedgerAccountSettlementListParams, opts ...option.RequestOption) (res *pagination.Page[LedgerAccountSettlement], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_account_settlements"
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

// Get a list of ledger account settlements.
func (r *LedgerAccountSettlementService) ListAutoPaging(ctx context.Context, query LedgerAccountSettlementListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LedgerAccountSettlement] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type LedgerAccountSettlement struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of the ledger account settlement.
	Amount int64 `json:"amount,required,nullable"`
	// The id of the contra ledger account that sends to or receives funds from the
	// settled ledger account.
	ContraLedgerAccountID string    `json:"contra_ledger_account_id,required" format:"uuid"`
	CreatedAt             time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the ledger account settlement.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account settlement.
	CurrencyExponent int64 `json:"currency_exponent,required,nullable"`
	// The description of the ledger account settlement.
	Description string `json:"description,required,nullable"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account settlement. The default value is the
	// created_at timestamp of the ledger account settlement.
	EffectiveAtUpperBound time.Time `json:"effective_at_upper_bound,required" format:"date-time"`
	// The id of the ledger that this ledger account settlement belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// The id of the ledger transaction that this settlement is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The id of the settled ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	SettledLedgerAccountID string `json:"settled_ledger_account_id,required" format:"uuid"`
	// The direction of the ledger entry with the settlement_ledger_account.
	SettlementEntryDirection string `json:"settlement_entry_direction,required,nullable"`
	// The status of the ledger account settlement. One of `processing`, `pending`,
	// `posted`, `archiving` or `archived`.
	Status    LedgerAccountSettlementStatus `json:"status,required"`
	UpdatedAt time.Time                     `json:"updated_at,required" format:"date-time"`
	JSON      ledgerAccountSettlementJSON   `json:"-"`
}

// ledgerAccountSettlementJSON contains the JSON metadata for the struct
// [LedgerAccountSettlement]
type ledgerAccountSettlementJSON struct {
	ID                       apijson.Field
	Amount                   apijson.Field
	ContraLedgerAccountID    apijson.Field
	CreatedAt                apijson.Field
	Currency                 apijson.Field
	CurrencyExponent         apijson.Field
	Description              apijson.Field
	EffectiveAtUpperBound    apijson.Field
	LedgerID                 apijson.Field
	LedgerTransactionID      apijson.Field
	LiveMode                 apijson.Field
	Metadata                 apijson.Field
	Object                   apijson.Field
	SettledLedgerAccountID   apijson.Field
	SettlementEntryDirection apijson.Field
	Status                   apijson.Field
	UpdatedAt                apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LedgerAccountSettlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerAccountSettlementJSON) RawJSON() string {
	return r.raw
}

// The status of the ledger account settlement. One of `processing`, `pending`,
// `posted`, `archiving` or `archived`.
type LedgerAccountSettlementStatus string

const (
	LedgerAccountSettlementStatusArchived   LedgerAccountSettlementStatus = "archived"
	LedgerAccountSettlementStatusArchiving  LedgerAccountSettlementStatus = "archiving"
	LedgerAccountSettlementStatusDrafting   LedgerAccountSettlementStatus = "drafting"
	LedgerAccountSettlementStatusPending    LedgerAccountSettlementStatus = "pending"
	LedgerAccountSettlementStatusPosted     LedgerAccountSettlementStatus = "posted"
	LedgerAccountSettlementStatusProcessing LedgerAccountSettlementStatus = "processing"
)

func (r LedgerAccountSettlementStatus) IsKnown() bool {
	switch r {
	case LedgerAccountSettlementStatusArchived, LedgerAccountSettlementStatusArchiving, LedgerAccountSettlementStatusDrafting, LedgerAccountSettlementStatusPending, LedgerAccountSettlementStatusPosted, LedgerAccountSettlementStatusProcessing:
		return true
	}
	return false
}

type LedgerAccountSettlementNewParams struct {
	// The id of the contra ledger account that sends to or receives funds from the
	// settled ledger account.
	ContraLedgerAccountID param.Field[string] `json:"contra_ledger_account_id,required" format:"uuid"`
	// The id of the settled ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	SettledLedgerAccountID param.Field[string] `json:"settled_ledger_account_id,required" format:"uuid"`
	// If true, the settlement amount and settlement_entry_direction will bring the
	// settlement ledger account's balance closer to zero, even if the balance is
	// negative.
	AllowEitherDirection param.Field[bool] `json:"allow_either_direction"`
	// The description of the ledger account settlement.
	Description param.Field[string] `json:"description"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account settlement. The default value is the
	// created_at timestamp of the ledger account settlement.
	EffectiveAtUpperBound param.Field[time.Time] `json:"effective_at_upper_bound" format:"date-time"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// It is set to `false` by default. It should be set to `true` when migrating
	// existing settlements.
	SkipSettlementLedgerTransaction param.Field[bool] `json:"skip_settlement_ledger_transaction"`
	// The status of the ledger account settlement. It is set to `pending` by default.
	// To post a ledger account settlement at creation, use `posted`.
	Status param.Field[LedgerAccountSettlementNewParamsStatus] `json:"status"`
}

func (r LedgerAccountSettlementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the ledger account settlement. It is set to `pending` by default.
// To post a ledger account settlement at creation, use `posted`.
type LedgerAccountSettlementNewParamsStatus string

const (
	LedgerAccountSettlementNewParamsStatusPending LedgerAccountSettlementNewParamsStatus = "pending"
	LedgerAccountSettlementNewParamsStatusPosted  LedgerAccountSettlementNewParamsStatus = "posted"
)

func (r LedgerAccountSettlementNewParamsStatus) IsKnown() bool {
	switch r {
	case LedgerAccountSettlementNewParamsStatusPending, LedgerAccountSettlementNewParamsStatusPosted:
		return true
	}
	return false
}

type LedgerAccountSettlementUpdateParams struct {
	// The description of the ledger account settlement.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a pending ledger account settlement, use `posted`. To archive a pending
	// ledger transaction, use `archived`.
	Status param.Field[LedgerAccountSettlementUpdateParamsStatus] `json:"status"`
}

func (r LedgerAccountSettlementUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// To post a pending ledger account settlement, use `posted`. To archive a pending
// ledger transaction, use `archived`.
type LedgerAccountSettlementUpdateParamsStatus string

const (
	LedgerAccountSettlementUpdateParamsStatusPosted   LedgerAccountSettlementUpdateParamsStatus = "posted"
	LedgerAccountSettlementUpdateParamsStatusArchived LedgerAccountSettlementUpdateParamsStatus = "archived"
)

func (r LedgerAccountSettlementUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LedgerAccountSettlementUpdateParamsStatusPosted, LedgerAccountSettlementUpdateParamsStatusArchived:
		return true
	}
	return false
}

type LedgerAccountSettlementListParams struct {
	// If you have specific IDs to retrieve in bulk, you can pass them as query
	// parameters delimited with `id[]=`, for example `?id[]=123&id[]=abc`.
	ID          param.Field[[]string] `query:"id"`
	AfterCursor param.Field[string]   `query:"after_cursor"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// created at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// created_at%5Bgt%5D=2000-01-01T12:00:00Z.
	CreatedAt           param.Field[map[string]time.Time] `query:"created_at" format:"date-time"`
	LedgerID            param.Field[string]               `query:"ledger_id"`
	LedgerTransactionID param.Field[string]               `query:"ledger_transaction_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata                 param.Field[map[string]string] `query:"metadata"`
	PerPage                  param.Field[int64]             `query:"per_page"`
	SettledLedgerAccountID   param.Field[string]            `query:"settled_ledger_account_id"`
	SettlementEntryDirection param.Field[string]            `query:"settlement_entry_direction"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// updated at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt param.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes [LedgerAccountSettlementListParams]'s query parameters as
// `url.Values`.
func (r LedgerAccountSettlementListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
