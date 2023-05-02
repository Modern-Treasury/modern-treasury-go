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

type AccountCollectionFlowService struct {
	Options []option.RequestOption
}

func NewAccountCollectionFlowService(opts ...option.RequestOption) (r *AccountCollectionFlowService) {
	r = &AccountCollectionFlowService{}
	r.Options = opts
	return
}

// create account_collection_flow
func (r *AccountCollectionFlowService) New(ctx context.Context, body AccountCollectionFlowNewParams, opts ...option.RequestOption) (res *AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/account_collection_flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get account_collection_flow
func (r *AccountCollectionFlowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update account_collection_flow
func (r *AccountCollectionFlowService) Update(ctx context.Context, id string, body AccountCollectionFlowUpdateParams, opts ...option.RequestOption) (res *AccountConnectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
	JSON              AccountConnectionFlowJSON
}

type AccountConnectionFlowJSON struct {
	ID                apijson.Metadata
	Object            apijson.Metadata
	LiveMode          apijson.Metadata
	CreatedAt         apijson.Metadata
	UpdatedAt         apijson.Metadata
	ClientToken       apijson.Metadata
	Status            apijson.Metadata
	CounterpartyID    apijson.Metadata
	ExternalAccountID apijson.Metadata
	PaymentTypes      apijson.Metadata
	raw               string
	Extras            map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountConnectionFlow using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountConnectionFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountConnectionFlowStatus string

const (
	AccountConnectionFlowStatusCancelled AccountConnectionFlowStatus = "cancelled"
	AccountConnectionFlowStatusCompleted AccountConnectionFlowStatus = "completed"
	AccountConnectionFlowStatusExpired   AccountConnectionFlowStatus = "expired"
	AccountConnectionFlowStatusPending   AccountConnectionFlowStatus = "pending"
)

type AccountConnectionFlowPaymentTypes string

const (
	AccountConnectionFlowPaymentTypesACH  AccountConnectionFlowPaymentTypes = "ach"
	AccountConnectionFlowPaymentTypesWire AccountConnectionFlowPaymentTypes = "wire"
)

type AccountCollectionFlowNewParams struct {
	// Required.
	CounterpartyID field.Field[string]   `json:"counterparty_id,required" format:"uuid"`
	PaymentTypes   field.Field[[]string] `json:"payment_types,required"`
}

// MarshalJSON serializes AccountCollectionFlowNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r AccountCollectionFlowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCollectionFlowUpdateParams struct {
	// Required. The updated status of the account collection flow. Can only be used to
	// mark a flow as `cancelled`.
	Status field.Field[AccountCollectionFlowUpdateParamsStatus] `json:"status,required"`
}

// MarshalJSON serializes AccountCollectionFlowUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r AccountCollectionFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCollectionFlowUpdateParamsStatus string

const (
	AccountCollectionFlowUpdateParamsStatusCancelled AccountCollectionFlowUpdateParamsStatus = "cancelled"
)

type AccountCollectionFlowListParams struct {
	AfterCursor       field.Field[string] `query:"after_cursor,nullable"`
	PerPage           field.Field[int64]  `query:"per_page"`
	ClientToken       field.Field[string] `query:"client_token"`
	Status            field.Field[string] `query:"status"`
	CounterpartyID    field.Field[string] `query:"counterparty_id"`
	ExternalAccountID field.Field[string] `query:"external_account_id"`
}

// URLQuery serializes AccountCollectionFlowListParams into a url.Values of the
// query parameters associated with this value
func (r AccountCollectionFlowListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
