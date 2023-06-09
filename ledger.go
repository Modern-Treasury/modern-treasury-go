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

// LedgerService contains methods and other services that help with interacting
// with the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewLedgerService] method instead.
type LedgerService struct {
	Options []option.RequestOption
}

// NewLedgerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLedgerService(opts ...option.RequestOption) (r *LedgerService) {
	r = &LedgerService{}
	r.Options = opts
	return
}

// Create a ledger.
func (r *LedgerService) New(ctx context.Context, params LedgerNewParams, opts ...option.RequestOption) (res *Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledgers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details on a single ledger.
func (r *LedgerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the details of a ledger.
func (r *LedgerService) Update(ctx context.Context, id string, body LedgerUpdateParams, opts ...option.RequestOption) (res *Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledgers.
func (r *LedgerService) List(ctx context.Context, query LedgerListParams, opts ...option.RequestOption) (res *shared.Page[Ledger], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledgers"
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

// Get a list of ledgers.
func (r *LedgerService) ListAutoPaging(ctx context.Context, query LedgerListParams, opts ...option.RequestOption) *shared.PageAutoPager[Ledger] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger.
func (r *LedgerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Ledger struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the ledger.
	Name string `json:"name,required"`
	// An optional free-form description for internal use.
	Description string `json:"description,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata    map[string]string      `json:"metadata,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        ledgerJSON
}

// ledgerJSON contains the JSON metadata for the struct [Ledger]
type ledgerJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	LiveMode    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	DiscardedAt apijson.Field
	Name        apijson.Field
	Description apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerNewParams struct {
	// The name of the ledger.
	Name param.Field[string] `json:"name,required"`
	// An optional free-form description for internal use.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata       param.Field[map[string]string] `json:"metadata"`
	IdempotencyKey param.Field[string]            `header:"Idempotency-Key"`
}

func (r LedgerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerUpdateParams struct {
	// An optional free-form description for internal use.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The name of the ledger.
	Name param.Field[string] `json:"name"`
}

func (r LedgerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt param.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes [LedgerListParams]'s query parameters as `url.Values`.
func (r LedgerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
