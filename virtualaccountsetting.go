// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
)

// VirtualAccountSettingService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVirtualAccountSettingService] method instead.
type VirtualAccountSettingService struct {
	Options []option.RequestOption
}

// NewVirtualAccountSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVirtualAccountSettingService(opts ...option.RequestOption) (r *VirtualAccountSettingService) {
	r = &VirtualAccountSettingService{}
	r.Options = opts
	return
}

// Create a virtual account setting.
func (r *VirtualAccountSettingService) New(ctx context.Context, body VirtualAccountSettingNewParams, opts ...option.RequestOption) (res *VirtualAccountSetting, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/virtual_account_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List virtual account settings.
func (r *VirtualAccountSettingService) List(ctx context.Context, query VirtualAccountSettingListParams, opts ...option.RequestOption) (res *pagination.Page[VirtualAccountSetting], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/virtual_account_settings"
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

// List virtual account settings.
func (r *VirtualAccountSettingService) ListAutoPaging(ctx context.Context, query VirtualAccountSettingListParams, opts ...option.RequestOption) *pagination.PageAutoPager[VirtualAccountSetting] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type VirtualAccountSetting struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A user-defined identifier for the virtual account setting.
	ExternalID string `json:"external_id" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                      `json:"live_mode" api:"required"`
	Object    string                    `json:"object" api:"required"`
	UpdatedAt time.Time                 `json:"updated_at" api:"required" format:"date-time"`
	JSON      virtualAccountSettingJSON `json:"-"`
}

// virtualAccountSettingJSON contains the JSON metadata for the struct
// [VirtualAccountSetting]
type virtualAccountSettingJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ExternalID  apijson.Field
	LiveMode    apijson.Field
	Object      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VirtualAccountSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r virtualAccountSettingJSON) RawJSON() string {
	return r.raw
}

type VirtualAccountSettingNewParams struct {
	// The method used to allocate virtual account numbers.
	AllocationType param.Field[string] `json:"allocation_type" api:"required"`
	// The ID of the internal account for the virtual account setting.
	InternalAccountID param.Field[string] `json:"internal_account_id" api:"required" format:"uuid"`
	// The prefix, suffix, or bank-assigned identifier for the virtual account numbers.
	AllocationIdentifier param.Field[string] `json:"allocation_identifier"`
	// The total length of generated virtual account numbers.
	AllocationLength param.Field[int64] `json:"allocation_length"`
	// The inclusive end of the virtual account number range.
	AllocationRangeEnd param.Field[string] `json:"allocation_range_end"`
	// The inclusive start of the virtual account number range.
	AllocationRangeStart param.Field[string] `json:"allocation_range_start"`
	// A user-defined identifier for the virtual account setting.
	ExternalID param.Field[string] `json:"external_id"`
	// The length of a generated virtual account setting prefix.
	GeneratedAllocationIdentifierLength param.Field[int64] `json:"generated_allocation_identifier_length"`
}

func (r VirtualAccountSettingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VirtualAccountSettingListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// A user-defined identifier for the virtual account setting.
	ExternalID param.Field[string] `query:"external_id"`
	PerPage    param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [VirtualAccountSettingListParams]'s query parameters as
// `url.Values`.
func (r VirtualAccountSettingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
