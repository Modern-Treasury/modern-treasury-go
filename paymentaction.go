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

// PaymentActionService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentActionService] method instead.
type PaymentActionService struct {
	Options []option.RequestOption
}

// NewPaymentActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentActionService(opts ...option.RequestOption) (r *PaymentActionService) {
	r = &PaymentActionService{}
	r.Options = opts
	return
}

// Create a payment action.
func (r *PaymentActionService) New(ctx context.Context, body PaymentActionNewParams, opts ...option.RequestOption) (res *PaymentActionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/payment_actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single payment action.
func (r *PaymentActionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentActionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_actions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a single payment action.
func (r *PaymentActionService) Update(ctx context.Context, id string, body PaymentActionUpdateParams, opts ...option.RequestOption) (res *PaymentActionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_actions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all payment actions.
func (r *PaymentActionService) List(ctx context.Context, query PaymentActionListParams, opts ...option.RequestOption) (res *pagination.Page[PaymentActionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_actions"
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

// Get a list of all payment actions.
func (r *PaymentActionService) ListAutoPaging(ctx context.Context, query PaymentActionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PaymentActionListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type PaymentActionNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the associated actionable object.
	ActionableID string `json:"actionable_id" api:"required,nullable" format:"uuid"`
	// The type of the associated actionable object. One of `payment_order`,
	// `expected_payment`.
	ActionableType string    `json:"actionable_type" api:"required,nullable"`
	CreatedAt      time.Time `json:"created_at" api:"required" format:"date-time"`
	// The specifc details of the payment action based on type.
	Details interface{} `json:"details" api:"required"`
	// The ID of the internal account associated with the payment action.
	InternalAccountID string `json:"internal_account_id" api:"required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The current status of the payment action. One of `pending`, `processing`,
	// `sent`, `acknowledged`, `cancelled`, or `failed`.
	Status string `json:"status" api:"required"`
	// The type of the payment action. Determines the action to be taken.
	Type      string                       `json:"type" api:"required"`
	UpdatedAt time.Time                    `json:"updated_at" api:"required" format:"date-time"`
	JSON      paymentActionNewResponseJSON `json:"-"`
}

// paymentActionNewResponseJSON contains the JSON metadata for the struct
// [PaymentActionNewResponse]
type paymentActionNewResponseJSON struct {
	ID                apijson.Field
	ActionableID      apijson.Field
	ActionableType    apijson.Field
	CreatedAt         apijson.Field
	Details           apijson.Field
	InternalAccountID apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentActionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentActionNewResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentActionGetResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the associated actionable object.
	ActionableID string `json:"actionable_id" api:"required,nullable" format:"uuid"`
	// The type of the associated actionable object. One of `payment_order`,
	// `expected_payment`.
	ActionableType string    `json:"actionable_type" api:"required,nullable"`
	CreatedAt      time.Time `json:"created_at" api:"required" format:"date-time"`
	// The specifc details of the payment action based on type.
	Details interface{} `json:"details" api:"required"`
	// The ID of the internal account associated with the payment action.
	InternalAccountID string `json:"internal_account_id" api:"required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The current status of the payment action. One of `pending`, `processing`,
	// `sent`, `acknowledged`, `cancelled`, or `failed`.
	Status string `json:"status" api:"required"`
	// The type of the payment action. Determines the action to be taken.
	Type      string                       `json:"type" api:"required"`
	UpdatedAt time.Time                    `json:"updated_at" api:"required" format:"date-time"`
	JSON      paymentActionGetResponseJSON `json:"-"`
}

// paymentActionGetResponseJSON contains the JSON metadata for the struct
// [PaymentActionGetResponse]
type paymentActionGetResponseJSON struct {
	ID                apijson.Field
	ActionableID      apijson.Field
	ActionableType    apijson.Field
	CreatedAt         apijson.Field
	Details           apijson.Field
	InternalAccountID apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentActionGetResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentActionUpdateResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the associated actionable object.
	ActionableID string `json:"actionable_id" api:"required,nullable" format:"uuid"`
	// The type of the associated actionable object. One of `payment_order`,
	// `expected_payment`.
	ActionableType string    `json:"actionable_type" api:"required,nullable"`
	CreatedAt      time.Time `json:"created_at" api:"required" format:"date-time"`
	// The specifc details of the payment action based on type.
	Details interface{} `json:"details" api:"required"`
	// The ID of the internal account associated with the payment action.
	InternalAccountID string `json:"internal_account_id" api:"required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The current status of the payment action. One of `pending`, `processing`,
	// `sent`, `acknowledged`, `cancelled`, or `failed`.
	Status string `json:"status" api:"required"`
	// The type of the payment action. Determines the action to be taken.
	Type      string                          `json:"type" api:"required"`
	UpdatedAt time.Time                       `json:"updated_at" api:"required" format:"date-time"`
	JSON      paymentActionUpdateResponseJSON `json:"-"`
}

// paymentActionUpdateResponseJSON contains the JSON metadata for the struct
// [PaymentActionUpdateResponse]
type paymentActionUpdateResponseJSON struct {
	ID                apijson.Field
	ActionableID      apijson.Field
	ActionableType    apijson.Field
	CreatedAt         apijson.Field
	Details           apijson.Field
	InternalAccountID apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentActionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentActionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentActionListResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the associated actionable object.
	ActionableID string `json:"actionable_id" api:"required,nullable" format:"uuid"`
	// The type of the associated actionable object. One of `payment_order`,
	// `expected_payment`.
	ActionableType string    `json:"actionable_type" api:"required,nullable"`
	CreatedAt      time.Time `json:"created_at" api:"required" format:"date-time"`
	// The specifc details of the payment action based on type.
	Details interface{} `json:"details" api:"required"`
	// The ID of the internal account associated with the payment action.
	InternalAccountID string `json:"internal_account_id" api:"required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The current status of the payment action. One of `pending`, `processing`,
	// `sent`, `acknowledged`, `cancelled`, or `failed`.
	Status string `json:"status" api:"required"`
	// The type of the payment action. Determines the action to be taken.
	Type      string                        `json:"type" api:"required"`
	UpdatedAt time.Time                     `json:"updated_at" api:"required" format:"date-time"`
	JSON      paymentActionListResponseJSON `json:"-"`
}

// paymentActionListResponseJSON contains the JSON metadata for the struct
// [PaymentActionListResponse]
type paymentActionListResponseJSON struct {
	ID                apijson.Field
	ActionableID      apijson.Field
	ActionableType    apijson.Field
	CreatedAt         apijson.Field
	Details           apijson.Field
	InternalAccountID apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	Status            apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PaymentActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentActionListResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentActionNewParams struct {
	// Required. The type of the payment action. Determines the action to be taken.
	Type param.Field[string] `json:"type" api:"required"`
	// Optional. The ID of the associated actionable object.
	ActionableID param.Field[string] `json:"actionable_id" format:"uuid"`
	// Optional. The type of the associated actionable object. One of `payment_order`,
	// `expected_payment`. Required if `actionable_id` is passed.
	ActionableType param.Field[string] `json:"actionable_type"`
	// Optional. The specifc details of the payment action based on type.
	Details param.Field[interface{}] `json:"details"`
	// Optional. The ID of one of your organization's internal accounts. Required if
	// `actionable_id` is not passed.
	InternalAccountID param.Field[string] `json:"internal_account_id" format:"uuid"`
}

func (r PaymentActionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentActionUpdateParams struct {
	// Optional. Set the status of the payment action to `cancelled` to cancel the
	// payment action. This will only work if the payment action is in a `pending`
	// state.
	Status param.Field[PaymentActionUpdateParamsStatus] `json:"status" api:"required"`
}

func (r PaymentActionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional. Set the status of the payment action to `cancelled` to cancel the
// payment action. This will only work if the payment action is in a `pending`
// state.
type PaymentActionUpdateParamsStatus string

const (
	PaymentActionUpdateParamsStatusPending      PaymentActionUpdateParamsStatus = "pending"
	PaymentActionUpdateParamsStatusProcessable  PaymentActionUpdateParamsStatus = "processable"
	PaymentActionUpdateParamsStatusProcessing   PaymentActionUpdateParamsStatus = "processing"
	PaymentActionUpdateParamsStatusSent         PaymentActionUpdateParamsStatus = "sent"
	PaymentActionUpdateParamsStatusAcknowledged PaymentActionUpdateParamsStatus = "acknowledged"
	PaymentActionUpdateParamsStatusFailed       PaymentActionUpdateParamsStatus = "failed"
	PaymentActionUpdateParamsStatusCancelled    PaymentActionUpdateParamsStatus = "cancelled"
)

func (r PaymentActionUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PaymentActionUpdateParamsStatusPending, PaymentActionUpdateParamsStatusProcessable, PaymentActionUpdateParamsStatusProcessing, PaymentActionUpdateParamsStatusSent, PaymentActionUpdateParamsStatusAcknowledged, PaymentActionUpdateParamsStatusFailed, PaymentActionUpdateParamsStatusCancelled:
		return true
	}
	return false
}

type PaymentActionListParams struct {
	// The ID of the associated actionable object.
	ActionableID param.Field[string] `query:"actionable_id"`
	// The type of the associated actionable object. One of `payment_order`,
	// `expected_payment`.
	ActionableType param.Field[string] `query:"actionable_type"`
	AfterCursor    param.Field[string] `query:"after_cursor"`
	// Filter by `created_at` using comparison operators like `gt` (>), `gte` (>=),
	// `lt` (<), `lte` (<=), or `eq` (=) to filter by the created at timestamp. For
	// example, `created_at[gte]=2024-01-01T00:00:00Z`
	CreatedAt param.Field[PaymentActionListParamsCreatedAt] `query:"created_at"`
	// The ID of one of your internal accounts.
	InternalAccountID param.Field[string] `query:"internal_account_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// Filter by payment actions of a specific status.
	Status param.Field[PaymentActionListParamsStatus] `query:"status"`
	// The type of payment action.
	Type param.Field[PaymentActionListParamsType] `query:"type"`
}

// URLQuery serializes [PaymentActionListParams]'s query parameters as
// `url.Values`.
func (r PaymentActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by `created_at` using comparison operators like `gt` (>), `gte` (>=),
// `lt` (<), `lte` (<=), or `eq` (=) to filter by the created at timestamp. For
// example, `created_at[gte]=2024-01-01T00:00:00Z`
type PaymentActionListParamsCreatedAt struct {
	Eq  param.Field[time.Time] `query:"eq" format:"date-time"`
	Gt  param.Field[time.Time] `query:"gt" format:"date-time"`
	Gte param.Field[time.Time] `query:"gte" format:"date-time"`
	Lt  param.Field[time.Time] `query:"lt" format:"date-time"`
	Lte param.Field[time.Time] `query:"lte" format:"date-time"`
}

// URLQuery serializes [PaymentActionListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r PaymentActionListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by payment actions of a specific status.
type PaymentActionListParamsStatus string

const (
	PaymentActionListParamsStatusPending      PaymentActionListParamsStatus = "pending"
	PaymentActionListParamsStatusProcessable  PaymentActionListParamsStatus = "processable"
	PaymentActionListParamsStatusProcessing   PaymentActionListParamsStatus = "processing"
	PaymentActionListParamsStatusSent         PaymentActionListParamsStatus = "sent"
	PaymentActionListParamsStatusAcknowledged PaymentActionListParamsStatus = "acknowledged"
	PaymentActionListParamsStatusFailed       PaymentActionListParamsStatus = "failed"
	PaymentActionListParamsStatusCancelled    PaymentActionListParamsStatus = "cancelled"
)

func (r PaymentActionListParamsStatus) IsKnown() bool {
	switch r {
	case PaymentActionListParamsStatusPending, PaymentActionListParamsStatusProcessable, PaymentActionListParamsStatusProcessing, PaymentActionListParamsStatusSent, PaymentActionListParamsStatusAcknowledged, PaymentActionListParamsStatusFailed, PaymentActionListParamsStatusCancelled:
		return true
	}
	return false
}

// The type of payment action.
type PaymentActionListParamsType string

const (
	PaymentActionListParamsTypeStop  PaymentActionListParamsType = "stop"
	PaymentActionListParamsTypeIssue PaymentActionListParamsType = "issue"
)

func (r PaymentActionListParamsType) IsKnown() bool {
	switch r {
	case PaymentActionListParamsTypeStop, PaymentActionListParamsTypeIssue:
		return true
	}
	return false
}
