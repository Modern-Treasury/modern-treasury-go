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

// AccountCollectionFlowService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountCollectionFlowService] method instead.
type AccountCollectionFlowService struct {
	Options []option.RequestOption
}

// NewAccountCollectionFlowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCollectionFlowService(opts ...option.RequestOption) (r *AccountCollectionFlowService) {
	r = &AccountCollectionFlowService{}
	r.Options = opts
	return
}

// create account_collection_flow
func (r *AccountCollectionFlowService) New(ctx context.Context, params AccountCollectionFlowNewParams, opts ...option.RequestOption) (res *AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/account_collection_flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// get account_collection_flow
func (r *AccountCollectionFlowService) Get(ctx context.Context, id string, query AccountCollectionFlowGetParams, opts ...option.RequestOption) (res *AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// update account_collection_flow
func (r *AccountCollectionFlowService) Update(ctx context.Context, id string, params AccountCollectionFlowUpdateParams, opts ...option.RequestOption) (res *AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// list account_collection_flows
func (r *AccountCollectionFlowService) List(ctx context.Context, query AccountCollectionFlowListParams, opts ...option.RequestOption) (res *shared.Page[AccountConnectionFlow], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/account_collection_flows"
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

// list account_collection_flows
func (r *AccountCollectionFlowService) ListAutoPaging(ctx context.Context, query AccountCollectionFlowListParams, opts ...option.RequestOption) *shared.PageAutoPager[AccountConnectionFlow] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type AccountConnectionFlow struct {
	ID     string `json:"id" format:"uuid"`
	Object string `json:"object"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The client token of the account collection flow. This token can be used to embed
	// account collection in your client-side application.
	ClientToken string `json:"client_token"`
	// The current status of the account collection flow. One of `pending`,
	// `completed`, `expired`, or `cancelled`.
	Status AccountConnectionFlowStatus `json:"status"`
	// The ID of a counterparty. An external account created with this flow will be
	// associated with this counterparty.
	CounterpartyID string `json:"counterparty_id,required" format:"uuid"`
	// If present, the ID of the external account created using this flow.
	ExternalAccountID string                              `json:"external_account_id,nullable" format:"uuid"`
	PaymentTypes      []AccountConnectionFlowPaymentTypes `json:"payment_types,required"`
	JSON              accountConnectionFlowJSON
}

// accountConnectionFlowJSON contains the JSON metadata for the struct
// [AccountConnectionFlow]
type accountConnectionFlowJSON struct {
	ID                apijson.Field
	Object            apijson.Field
	LiveMode          apijson.Field
	CreatedAt         apijson.Field
	UpdatedAt         apijson.Field
	ClientToken       apijson.Field
	Status            apijson.Field
	CounterpartyID    apijson.Field
	ExternalAccountID apijson.Field
	PaymentTypes      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountConnectionFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the account collection flow. One of `pending`,
// `completed`, `expired`, or `cancelled`.
type AccountConnectionFlowStatus string

const (
	AccountConnectionFlowStatusCancelled AccountConnectionFlowStatus = "cancelled"
	AccountConnectionFlowStatusCompleted AccountConnectionFlowStatus = "completed"
	AccountConnectionFlowStatusExpired   AccountConnectionFlowStatus = "expired"
	AccountConnectionFlowStatusPending   AccountConnectionFlowStatus = "pending"
)

// An account created with this flow will support payments of one of these types.
type AccountConnectionFlowPaymentTypes string

const (
	AccountConnectionFlowPaymentTypesACH  AccountConnectionFlowPaymentTypes = "ach"
	AccountConnectionFlowPaymentTypesWire AccountConnectionFlowPaymentTypes = "wire"
)

type AccountCollectionFlowNewParams struct {
	// Required.
	CounterpartyID param.Field[string]   `json:"counterparty_id,required" format:"uuid"`
	PaymentTypes   param.Field[[]string] `json:"payment_types,required"`
	IdempotencyKey param.Field[string]   `header:"Idempotency-Key"`
}

func (r AccountCollectionFlowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCollectionFlowGetParams struct {
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

type AccountCollectionFlowUpdateParams struct {
	// Required. The updated status of the account collection flow. Can only be used to
	// mark a flow as `cancelled`.
	Status         param.Field[AccountCollectionFlowUpdateParamsStatus] `json:"status,required"`
	IdempotencyKey param.Field[string]                                  `header:"Idempotency-Key"`
}

func (r AccountCollectionFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required. The updated status of the account collection flow. Can only be used to
// mark a flow as `cancelled`.
type AccountCollectionFlowUpdateParamsStatus string

const (
	AccountCollectionFlowUpdateParamsStatusCancelled AccountCollectionFlowUpdateParamsStatus = "cancelled"
)

type AccountCollectionFlowListParams struct {
	AfterCursor       param.Field[string] `query:"after_cursor"`
	ClientToken       param.Field[string] `query:"client_token"`
	CounterpartyID    param.Field[string] `query:"counterparty_id"`
	ExternalAccountID param.Field[string] `query:"external_account_id"`
	PerPage           param.Field[int64]  `query:"per_page"`
	Status            param.Field[string] `query:"status"`
}

// URLQuery serializes [AccountCollectionFlowListParams]'s query parameters as
// `url.Values`.
func (r AccountCollectionFlowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
