package moderntreasury

import (
	"context"
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

// ConnectionService contains methods and other services that help with interacting
// with the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

// Get a list of all connections.
func (r *ConnectionService) List(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) (res *shared.Page[Connection], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/connections"
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

// Get a list of all connections.
func (r *ConnectionService) ListAutoPaging(ctx context.Context, query ConnectionListParams, opts ...option.RequestOption) *shared.PageAutoPager[Connection] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Connection struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Unique identifier for the bank or vendor.
	VendorID string `json:"vendor_id,required" format:"uuid"`
	// An identifier given to this connection by the bank.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable" format:"uuid"`
	// A human-friendly name for the bank or vendor.
	VendorName string `json:"vendor_name,required"`
	JSON       connectionJSON
}

// connectionJSON contains the JSON metadata for the struct [Connection]
type connectionJSON struct {
	ID               apijson.Field
	Object           apijson.Field
	LiveMode         apijson.Field
	CreatedAt        apijson.Field
	UpdatedAt        apijson.Field
	DiscardedAt      apijson.Field
	VendorID         apijson.Field
	VendorCustomerID apijson.Field
	VendorName       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
	// An identifier assigned by the vendor to your organization.
	VendorCustomerID param.Field[string] `query:"vendor_customer_id" format:"uuid"`
	// A string code representing the vendor (i.e. bank).
	Entity param.Field[string] `query:"entity"`
}

// URLQuery serializes [ConnectionListParams]'s query parameters as `url.Values`.
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
