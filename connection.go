// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
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
	ID          string    `json:"id,required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An identifier given to this connection by the bank.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable" format:"uuid"`
	// Unique identifier for the bank or vendor.
	VendorID string `json:"vendor_id,required" format:"uuid"`
	// A human-friendly name for the bank or vendor.
	VendorName string         `json:"vendor_name,required"`
	JSON       connectionJSON `json:"-"`
}

// connectionJSON contains the JSON metadata for the struct [Connection]
type connectionJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	DiscardedAt      apijson.Field
	LiveMode         apijson.Field
	Object           apijson.Field
	UpdatedAt        apijson.Field
	VendorCustomerID apijson.Field
	VendorID         apijson.Field
	VendorName       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionJSON) RawJSON() string {
	return r.raw
}

type ConnectionListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// A string code representing the vendor (i.e. bank).
	Entity  param.Field[string] `query:"entity"`
	PerPage param.Field[int64]  `query:"per_page"`
	// An identifier assigned by the vendor to your organization.
	VendorCustomerID param.Field[string] `query:"vendor_customer_id"`
}

// URLQuery serializes [ConnectionListParams]'s query parameters as `url.Values`.
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
