package moderntreasury

import (
	"context"
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

type ConnectionService struct {
	Options []option.RequestOption
}

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
	JSON       ConnectionJSON
}

type ConnectionJSON struct {
	ID               apijson.Metadata
	Object           apijson.Metadata
	LiveMode         apijson.Metadata
	CreatedAt        apijson.Metadata
	UpdatedAt        apijson.Metadata
	DiscardedAt      apijson.Metadata
	VendorID         apijson.Metadata
	VendorCustomerID apijson.Metadata
	VendorName       apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Connection using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// An identifier assigned by the vendor to your organization.
	VendorCustomerID field.Field[string] `query:"vendor_customer_id" format:"uuid"`
	// A string code representing the vendor (i.e. bank).
	Entity field.Field[string] `query:"entity"`
}

// URLQuery serializes ConnectionListParams into a url.Values of the query
// parameters associated with this value
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
