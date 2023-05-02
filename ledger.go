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

type LedgerService struct {
	Options []option.RequestOption
}

func NewLedgerService(opts ...option.RequestOption) (r *LedgerService) {
	r = &LedgerService{}
	r.Options = opts
	return
}

// Create a ledger.
func (r *LedgerService) New(ctx context.Context, body LedgerNewParams, opts ...option.RequestOption) (res *Ledger, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledgers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	JSON        LedgerJSON
}

type LedgerJSON struct {
	ID          apijson.Metadata
	Object      apijson.Metadata
	LiveMode    apijson.Metadata
	CreatedAt   apijson.Metadata
	UpdatedAt   apijson.Metadata
	DiscardedAt apijson.Metadata
	Name        apijson.Metadata
	Description apijson.Metadata
	Metadata    apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ledger using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Ledger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerNewParams struct {
	// The name of the ledger.
	Name field.Field[string] `json:"name,required"`
	// An optional free-form description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r LedgerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerUpdateParams struct {
	// The name of the ledger.
	Name field.Field[string] `json:"name"`
	// An optional free-form description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r LedgerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt field.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes LedgerListParams into a url.Values of the query parameters
// associated with this value
func (r LedgerListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
