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

// LedgerEventHandlerService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLedgerEventHandlerService] method instead.
type LedgerEventHandlerService struct {
	Options []option.RequestOption
}

// NewLedgerEventHandlerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLedgerEventHandlerService(opts ...option.RequestOption) (r *LedgerEventHandlerService) {
	r = &LedgerEventHandlerService{}
	r.Options = opts
	return
}

// create ledger_event_handler
func (r *LedgerEventHandlerService) New(ctx context.Context, params LedgerEventHandlerNewParams, opts ...option.RequestOption) (res *LedgerEventHandlerCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_event_handlers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details on a single ledger event handler.
func (r *LedgerEventHandlerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerEventHandlerCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_event_handlers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of ledger event handlers.
func (r *LedgerEventHandlerService) List(ctx context.Context, query LedgerEventHandlerListParams, opts ...option.RequestOption) (res *shared.Page[LedgerEventHandlerCreateResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_event_handlers"
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

// Get a list of ledger event handlers.
func (r *LedgerEventHandlerService) ListAutoPaging(ctx context.Context, query LedgerEventHandlerListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerEventHandlerCreateResponse] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a ledger event handler.
func (r *LedgerEventHandlerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerEventHandlerCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_event_handlers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LedgerEventHandlerCreateResponse struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Name of the ledger event handler.
	Name string `json:"name,required"`
	// An optional description.
	Description               string                                                    `json:"description,required,nullable"`
	LedgerTransactionTemplate LedgerEventHandlerCreateResponseLedgerTransactionTemplate `json:"ledger_transaction_template,required"`
	Conditions                LedgerEventHandlerCreateResponseConditions                `json:"conditions,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required,nullable"`
	JSON     ledgerEventHandlerCreateResponseJSON
}

// ledgerEventHandlerCreateResponseJSON contains the JSON metadata for the struct
// [LedgerEventHandlerCreateResponse]
type ledgerEventHandlerCreateResponseJSON struct {
	ID                        apijson.Field
	Object                    apijson.Field
	LiveMode                  apijson.Field
	CreatedAt                 apijson.Field
	UpdatedAt                 apijson.Field
	DiscardedAt               apijson.Field
	Name                      apijson.Field
	Description               apijson.Field
	LedgerTransactionTemplate apijson.Field
	Conditions                apijson.Field
	Metadata                  apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LedgerEventHandlerCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerCreateResponseLedgerTransactionTemplate struct {
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt string `json:"effective_at,required,nullable" format:"datetime"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntries `json:"ledger_entries,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required,nullable"`
	JSON     ledgerEventHandlerCreateResponseLedgerTransactionTemplateJSON
}

// ledgerEventHandlerCreateResponseLedgerTransactionTemplateJSON contains the JSON
// metadata for the struct
// [LedgerEventHandlerCreateResponseLedgerTransactionTemplate]
type ledgerEventHandlerCreateResponseLedgerTransactionTemplateJSON struct {
	Description   apijson.Field
	EffectiveAt   apijson.Field
	LedgerEntries apijson.Field
	Metadata      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LedgerEventHandlerCreateResponseLedgerTransactionTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntries struct {
	// The field you're fetching from the `ledgerable_event`.
	Amount string `json:"amount,required"`
	// What the operator between the `field` and `value` is. Currently only supports
	// `equals`.
	Direction string `json:"direction,required"`
	// What raw string you are comparing the `field` against.
	LedgerAccountID string `json:"ledger_account_id,required"`
	JSON            ledgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntriesJSON
}

// ledgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntriesJSON
// contains the JSON metadata for the struct
// [LedgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntries]
type ledgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntriesJSON struct {
	Amount          apijson.Field
	Direction       apijson.Field
	LedgerAccountID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LedgerEventHandlerCreateResponseLedgerTransactionTemplateLedgerEntries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerCreateResponseConditions struct {
	// The field you're fetching from the `ledgerable_event`.
	Field string `json:"field,required"`
	// What the operator between the `field` and `value` is. Currently only supports
	// `equals`.
	Operator string `json:"operator,required"`
	// What raw string you are comparing the `field` against.
	Value string `json:"value,required"`
	JSON  ledgerEventHandlerCreateResponseConditionsJSON
}

// ledgerEventHandlerCreateResponseConditionsJSON contains the JSON metadata for
// the struct [LedgerEventHandlerCreateResponseConditions]
type ledgerEventHandlerCreateResponseConditionsJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LedgerEventHandlerCreateResponseConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerNewParams struct {
	LedgerTransactionTemplate param.Field[LedgerEventHandlerNewParamsLedgerTransactionTemplate] `json:"ledger_transaction_template,required"`
	// Name of the ledger event handler.
	Name       param.Field[string]                                `json:"name,required"`
	Conditions param.Field[LedgerEventHandlerNewParamsConditions] `json:"conditions"`
	// An optional description.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata       param.Field[map[string]string] `json:"metadata"`
	IdempotencyKey param.Field[string]            `header:"Idempotency-Key"`
}

func (r LedgerEventHandlerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParamsLedgerTransactionTemplate struct {
	// An optional description for internal use.
	Description param.Field[string] `json:"description,required"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[string] `json:"effective_at,required" format:"datetime"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntries] `json:"ledger_entries,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata,required"`
}

func (r LedgerEventHandlerNewParamsLedgerTransactionTemplate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntries struct {
	// The field you're fetching from the `ledgerable_event`.
	Amount param.Field[string] `json:"amount,required"`
	// What the operator between the `field` and `value` is. Currently only supports
	// `equals`.
	Direction param.Field[string] `json:"direction,required"`
	// What raw string you are comparing the `field` against.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required"`
}

func (r LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntries) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParamsConditions struct {
	// The field you're fetching from the `ledgerable_event`.
	Field param.Field[string] `json:"field,required"`
	// What the operator between the `field` and `value` is. Currently only supports
	// `equals`.
	Operator param.Field[string] `json:"operator,required"`
	// What raw string you are comparing the `field` against.
	Value param.Field[string] `json:"value,required"`
}

func (r LedgerEventHandlerNewParamsConditions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// created_at%5Bgt%5D=2000-01-01T12:00:00Z.
	CreatedAt param.Field[map[string]time.Time] `query:"created_at" format:"date-time"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	Name     param.Field[string]            `query:"name"`
	PerPage  param.Field[int64]             `query:"per_page"`
}

// URLQuery serializes [LedgerEventHandlerListParams]'s query parameters as
// `url.Values`.
func (r LedgerEventHandlerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
