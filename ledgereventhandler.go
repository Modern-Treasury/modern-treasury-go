// File generated from our OpenAPI spec by Stainless.

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
func (r *LedgerEventHandlerService) New(ctx context.Context, body LedgerEventHandlerNewParams, opts ...option.RequestOption) (res *LedgerEventHandler, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_event_handlers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger event handler.
func (r *LedgerEventHandlerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerEventHandler, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_event_handlers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of ledger event handlers.
func (r *LedgerEventHandlerService) List(ctx context.Context, query LedgerEventHandlerListParams, opts ...option.RequestOption) (res *shared.Page[LedgerEventHandler], err error) {
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
func (r *LedgerEventHandlerService) ListAutoPaging(ctx context.Context, query LedgerEventHandlerListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerEventHandler] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Archive a ledger event handler.
func (r *LedgerEventHandlerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerEventHandler, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_event_handlers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LedgerEventHandler struct {
	ID         string                       `json:"id,required" format:"uuid"`
	Conditions LedgerEventHandlerConditions `json:"conditions,required,nullable"`
	CreatedAt  time.Time                    `json:"created_at,required" format:"date-time"`
	// An optional description.
	Description string    `json:"description,required,nullable"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The id of the ledger that this event handler belongs to.
	LedgerID                  string                                      `json:"ledger_id,required,nullable"`
	LedgerTransactionTemplate LedgerEventHandlerLedgerTransactionTemplate `json:"ledger_transaction_template,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required,nullable"`
	// Name of the ledger event handler.
	Name      string                                `json:"name,required"`
	Object    string                                `json:"object,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	Variables map[string]LedgerEventHandlerVariable `json:"variables,required,nullable"`
	JSON      ledgerEventHandlerJSON
}

// ledgerEventHandlerJSON contains the JSON metadata for the struct
// [LedgerEventHandler]
type ledgerEventHandlerJSON struct {
	ID                        apijson.Field
	Conditions                apijson.Field
	CreatedAt                 apijson.Field
	Description               apijson.Field
	DiscardedAt               apijson.Field
	LedgerID                  apijson.Field
	LedgerTransactionTemplate apijson.Field
	LiveMode                  apijson.Field
	Metadata                  apijson.Field
	Name                      apijson.Field
	Object                    apijson.Field
	UpdatedAt                 apijson.Field
	Variables                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LedgerEventHandler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerConditions struct {
	// The LHS of the conditional.
	Field string `json:"field,required"`
	// What the operator between the `field` and `value` is.
	Operator string `json:"operator,required"`
	// The RHS of the conditional.
	Value string `json:"value,required"`
	JSON  ledgerEventHandlerConditionsJSON
}

// ledgerEventHandlerConditionsJSON contains the JSON metadata for the struct
// [LedgerEventHandlerConditions]
type ledgerEventHandlerConditionsJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LedgerEventHandlerConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerLedgerTransactionTemplate struct {
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt string `json:"effective_at,required,nullable"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerEventHandlerLedgerTransactionTemplateLedgerEntry `json:"ledger_entries,required"`
	// To post a ledger transaction at creation, use `posted`.
	Status string `json:"status,required,nullable"`
	JSON   ledgerEventHandlerLedgerTransactionTemplateJSON
}

// ledgerEventHandlerLedgerTransactionTemplateJSON contains the JSON metadata for
// the struct [LedgerEventHandlerLedgerTransactionTemplate]
type ledgerEventHandlerLedgerTransactionTemplateJSON struct {
	Description   apijson.Field
	EffectiveAt   apijson.Field
	LedgerEntries apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LedgerEventHandlerLedgerTransactionTemplate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerLedgerTransactionTemplateLedgerEntry struct {
	// The LHS of the conditional.
	Amount string `json:"amount,required"`
	// What the operator between the `field` and `value` is.
	Direction string `json:"direction,required"`
	// The RHS of the conditional.
	LedgerAccountID string `json:"ledger_account_id,required"`
	JSON            ledgerEventHandlerLedgerTransactionTemplateLedgerEntryJSON
}

// ledgerEventHandlerLedgerTransactionTemplateLedgerEntryJSON contains the JSON
// metadata for the struct [LedgerEventHandlerLedgerTransactionTemplateLedgerEntry]
type ledgerEventHandlerLedgerTransactionTemplateLedgerEntryJSON struct {
	Amount          apijson.Field
	Direction       apijson.Field
	LedgerAccountID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LedgerEventHandlerLedgerTransactionTemplateLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerVariable struct {
	Query LedgerEventHandlerVariableQuery `json:"query,required"`
	// The type of object this variable is. Currently, only "ledger_account" is
	// supported.
	Type string `json:"type,required"`
	JSON ledgerEventHandlerVariableJSON
}

// ledgerEventHandlerVariableJSON contains the JSON metadata for the struct
// [LedgerEventHandlerVariable]
type ledgerEventHandlerVariableJSON struct {
	Query       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LedgerEventHandlerVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerVariableQuery struct {
	// The LHS of the conditional.
	Field string `json:"field,required"`
	// What the operator between the `field` and `value` is.
	Operator string `json:"operator,required"`
	// The RHS of the conditional.
	Value string `json:"value,required"`
	JSON  ledgerEventHandlerVariableQueryJSON
}

// ledgerEventHandlerVariableQueryJSON contains the JSON metadata for the struct
// [LedgerEventHandlerVariableQuery]
type ledgerEventHandlerVariableQueryJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LedgerEventHandlerVariableQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerEventHandlerVariableParam struct {
	Query param.Field[LedgerEventHandlerVariableQueryParam] `json:"query,required"`
	// The type of object this variable is. Currently, only "ledger_account" is
	// supported.
	Type param.Field[string] `json:"type,required"`
}

func (r LedgerEventHandlerVariableParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerVariableQueryParam struct {
	// The LHS of the conditional.
	Field param.Field[string] `json:"field,required"`
	// What the operator between the `field` and `value` is.
	Operator param.Field[string] `json:"operator,required"`
	// The RHS of the conditional.
	Value param.Field[string] `json:"value,required"`
}

func (r LedgerEventHandlerVariableQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParams struct {
	LedgerTransactionTemplate param.Field[LedgerEventHandlerNewParamsLedgerTransactionTemplate] `json:"ledger_transaction_template,required"`
	// Name of the ledger event handler.
	Name       param.Field[string]                                `json:"name,required"`
	Conditions param.Field[LedgerEventHandlerNewParamsConditions] `json:"conditions"`
	// An optional description.
	Description param.Field[string] `json:"description"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata  param.Field[map[string]string]                          `json:"metadata"`
	Variables param.Field[map[string]LedgerEventHandlerVariableParam] `json:"variables"`
}

func (r LedgerEventHandlerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParamsLedgerTransactionTemplate struct {
	// An optional description for internal use.
	Description param.Field[string] `json:"description,required"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[string] `json:"effective_at,required"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntry] `json:"ledger_entries,required"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[string] `json:"status,required"`
}

func (r LedgerEventHandlerNewParamsLedgerTransactionTemplate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntry struct {
	// The LHS of the conditional.
	Amount param.Field[string] `json:"amount,required"`
	// What the operator between the `field` and `value` is.
	Direction param.Field[string] `json:"direction,required"`
	// The RHS of the conditional.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required"`
}

func (r LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerEventHandlerNewParamsConditions struct {
	// The LHS of the conditional.
	Field param.Field[string] `json:"field,required"`
	// What the operator between the `field` and `value` is.
	Operator param.Field[string] `json:"operator,required"`
	// The RHS of the conditional.
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
