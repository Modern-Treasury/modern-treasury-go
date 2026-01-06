// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
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

// HoldService contains methods and other services that help with interacting with
// the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHoldService] method instead.
type HoldService struct {
	Options []option.RequestOption
}

// NewHoldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHoldService(opts ...option.RequestOption) (r *HoldService) {
	r = &HoldService{}
	r.Options = opts
	return
}

// Create a new hold
func (r *HoldService) New(ctx context.Context, body HoldNewParams, opts ...option.RequestOption) (res *Hold, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/holds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific hold
func (r *HoldService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Hold, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/holds/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a hold
func (r *HoldService) Update(ctx context.Context, id string, body HoldUpdateParams, opts ...option.RequestOption) (res *Hold, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/holds/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of holds.
func (r *HoldService) List(ctx context.Context, query HoldListParams, opts ...option.RequestOption) (res *pagination.Page[Hold], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/holds"
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

// Get a list of holds.
func (r *HoldService) ListAutoPaging(ctx context.Context, query HoldListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Hold] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Hold struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The type of object
	Object HoldObject `json:"object,required"`
	// The status of the hold
	Status HoldStatus `json:"status,required"`
	// The ID of the target being held
	TargetID string `json:"target_id,required" format:"uuid"`
	// The type of target being held
	TargetType HoldTargetType `json:"target_type,required"`
	UpdatedAt  time.Time      `json:"updated_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode"`
	// Additional metadata for the hold
	Metadata map[string]string `json:"metadata,nullable"`
	// The reason for the hold
	Reason string `json:"reason,nullable"`
	// The resolution of the hold
	Resolution string `json:"resolution,nullable"`
	// When the hold was resolved
	ResolvedAt time.Time `json:"resolved_at,nullable" format:"date-time"`
	JSON       holdJSON  `json:"-"`
}

// holdJSON contains the JSON metadata for the struct [Hold]
type holdJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Object      apijson.Field
	Status      apijson.Field
	TargetID    apijson.Field
	TargetType  apijson.Field
	UpdatedAt   apijson.Field
	LiveMode    apijson.Field
	Metadata    apijson.Field
	Reason      apijson.Field
	Resolution  apijson.Field
	ResolvedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Hold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdJSON) RawJSON() string {
	return r.raw
}

// The type of object
type HoldObject string

const (
	HoldObjectHold HoldObject = "hold"
)

func (r HoldObject) IsKnown() bool {
	switch r {
	case HoldObjectHold:
		return true
	}
	return false
}

// The status of the hold
type HoldStatus string

const (
	HoldStatusActive   HoldStatus = "active"
	HoldStatusResolved HoldStatus = "resolved"
)

func (r HoldStatus) IsKnown() bool {
	switch r {
	case HoldStatusActive, HoldStatusResolved:
		return true
	}
	return false
}

// The type of target being held
type HoldTargetType string

const (
	HoldTargetTypePaymentOrder HoldTargetType = "payment_order"
)

func (r HoldTargetType) IsKnown() bool {
	switch r {
	case HoldTargetTypePaymentOrder:
		return true
	}
	return false
}

type HoldNewParams struct {
	// The status of the hold
	Status param.Field[HoldNewParamsStatus] `json:"status,required"`
	// The ID of the target to hold
	TargetID param.Field[string] `json:"target_id,required" format:"uuid"`
	// The type of target to hold
	TargetType param.Field[HoldNewParamsTargetType] `json:"target_type,required"`
	// Additional metadata for the hold
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The reason for the hold
	Reason param.Field[string] `json:"reason"`
}

func (r HoldNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the hold
type HoldNewParamsStatus string

const (
	HoldNewParamsStatusActive HoldNewParamsStatus = "active"
)

func (r HoldNewParamsStatus) IsKnown() bool {
	switch r {
	case HoldNewParamsStatusActive:
		return true
	}
	return false
}

// The type of target to hold
type HoldNewParamsTargetType string

const (
	HoldNewParamsTargetTypePaymentOrder HoldNewParamsTargetType = "payment_order"
)

func (r HoldNewParamsTargetType) IsKnown() bool {
	switch r {
	case HoldNewParamsTargetTypePaymentOrder:
		return true
	}
	return false
}

type HoldUpdateParams struct {
	// The new status of the hold
	Status param.Field[HoldUpdateParamsStatus] `json:"status,required"`
	// The resolution of the hold
	Resolution param.Field[string] `json:"resolution"`
}

func (r HoldUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The new status of the hold
type HoldUpdateParamsStatus string

const (
	HoldUpdateParamsStatusResolved HoldUpdateParamsStatus = "resolved"
)

func (r HoldUpdateParamsStatus) IsKnown() bool {
	switch r {
	case HoldUpdateParamsStatusResolved:
		return true
	}
	return false
}

type HoldListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// Translation missing: en.openapi.descriptions.payment_order.query_params.status
	Status param.Field[HoldListParamsStatus] `query:"status"`
	// Translation missing:
	// en.openapi.descriptions.payment_order.query_params.target_id
	TargetID param.Field[string] `query:"target_id"`
	// Translation missing:
	// en.openapi.descriptions.payment_order.query_params.target_type
	TargetType param.Field[HoldListParamsTargetType] `query:"target_type"`
}

// URLQuery serializes [HoldListParams]'s query parameters as `url.Values`.
func (r HoldListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Translation missing: en.openapi.descriptions.payment_order.query_params.status
type HoldListParamsStatus string

const (
	HoldListParamsStatusActive   HoldListParamsStatus = "active"
	HoldListParamsStatusResolved HoldListParamsStatus = "resolved"
)

func (r HoldListParamsStatus) IsKnown() bool {
	switch r {
	case HoldListParamsStatusActive, HoldListParamsStatusResolved:
		return true
	}
	return false
}

// Translation missing:
// en.openapi.descriptions.payment_order.query_params.target_type
type HoldListParamsTargetType string

const (
	HoldListParamsTargetTypePaymentOrder HoldListParamsTargetType = "payment_order"
)

func (r HoldListParamsTargetType) IsKnown() bool {
	switch r {
	case HoldListParamsTargetTypePaymentOrder:
		return true
	}
	return false
}
