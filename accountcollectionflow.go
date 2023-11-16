// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
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
func (r *AccountCollectionFlowService) New(ctx context.Context, body AccountCollectionFlowNewParams, opts ...option.RequestOption) (res *AccountCollectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/account_collection_flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get account_collection_flow
func (r *AccountCollectionFlowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountCollectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update account_collection_flow
func (r *AccountCollectionFlowService) Update(ctx context.Context, id string, body AccountCollectionFlowUpdateParams, opts ...option.RequestOption) (res *AccountCollectionFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/account_collection_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list account_collection_flows
func (r *AccountCollectionFlowService) List(ctx context.Context, query AccountCollectionFlowListParams, opts ...option.RequestOption) (res *shared.Page[AccountCollectionFlow], err error) {
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
func (r *AccountCollectionFlowService) ListAutoPaging(ctx context.Context, query AccountCollectionFlowListParams, opts ...option.RequestOption) *shared.PageAutoPager[AccountCollectionFlow] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type AccountCollectionFlow struct {
	// The ID of a counterparty. An external account created with this flow will be
	// associated with this counterparty.
	CounterpartyID string                             `json:"counterparty_id,required" format:"uuid"`
	PaymentTypes   []AccountCollectionFlowPaymentType `json:"payment_types,required"`
	ID             string                             `json:"id" format:"uuid"`
	// The client token of the account collection flow. This token can be used to embed
	// account collection in your client-side application.
	ClientToken string    `json:"client_token"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	// If present, the ID of the external account created using this flow.
	ExternalAccountID string `json:"external_account_id,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode           bool                                    `json:"live_mode"`
	Object             string                                  `json:"object"`
	ReceivingCountries []AccountCollectionFlowReceivingCountry `json:"receiving_countries"`
	// The current status of the account collection flow. One of `pending`,
	// `completed`, `expired`, or `cancelled`.
	Status    AccountCollectionFlowStatus `json:"status"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      accountCollectionFlowJSON   `json:"-"`
}

// accountCollectionFlowJSON contains the JSON metadata for the struct
// [AccountCollectionFlow]
type accountCollectionFlowJSON struct {
	CounterpartyID     apijson.Field
	PaymentTypes       apijson.Field
	ID                 apijson.Field
	ClientToken        apijson.Field
	CreatedAt          apijson.Field
	ExternalAccountID  apijson.Field
	LiveMode           apijson.Field
	Object             apijson.Field
	ReceivingCountries apijson.Field
	Status             apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountCollectionFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An account created with this flow will support payments of one of these types.
type AccountCollectionFlowPaymentType string

const (
	AccountCollectionFlowPaymentTypeACH  AccountCollectionFlowPaymentType = "ach"
	AccountCollectionFlowPaymentTypeWire AccountCollectionFlowPaymentType = "wire"
)

// An account created with this flow will support wires from the US to these
// countries.
type AccountCollectionFlowReceivingCountry string

const (
	AccountCollectionFlowReceivingCountryUsa AccountCollectionFlowReceivingCountry = "USA"
	AccountCollectionFlowReceivingCountryAus AccountCollectionFlowReceivingCountry = "AUS"
	AccountCollectionFlowReceivingCountryBel AccountCollectionFlowReceivingCountry = "BEL"
	AccountCollectionFlowReceivingCountryCan AccountCollectionFlowReceivingCountry = "CAN"
	AccountCollectionFlowReceivingCountryChl AccountCollectionFlowReceivingCountry = "CHL"
	AccountCollectionFlowReceivingCountryChn AccountCollectionFlowReceivingCountry = "CHN"
	AccountCollectionFlowReceivingCountryCol AccountCollectionFlowReceivingCountry = "COL"
	AccountCollectionFlowReceivingCountryFra AccountCollectionFlowReceivingCountry = "FRA"
	AccountCollectionFlowReceivingCountryDeu AccountCollectionFlowReceivingCountry = "DEU"
	AccountCollectionFlowReceivingCountryHkg AccountCollectionFlowReceivingCountry = "HKG"
	AccountCollectionFlowReceivingCountryInd AccountCollectionFlowReceivingCountry = "IND"
	AccountCollectionFlowReceivingCountryIrl AccountCollectionFlowReceivingCountry = "IRL"
	AccountCollectionFlowReceivingCountryIta AccountCollectionFlowReceivingCountry = "ITA"
	AccountCollectionFlowReceivingCountryMex AccountCollectionFlowReceivingCountry = "MEX"
	AccountCollectionFlowReceivingCountryNld AccountCollectionFlowReceivingCountry = "NLD"
	AccountCollectionFlowReceivingCountryPer AccountCollectionFlowReceivingCountry = "PER"
	AccountCollectionFlowReceivingCountryEsp AccountCollectionFlowReceivingCountry = "ESP"
	AccountCollectionFlowReceivingCountryGbr AccountCollectionFlowReceivingCountry = "GBR"
)

// The current status of the account collection flow. One of `pending`,
// `completed`, `expired`, or `cancelled`.
type AccountCollectionFlowStatus string

const (
	AccountCollectionFlowStatusCancelled AccountCollectionFlowStatus = "cancelled"
	AccountCollectionFlowStatusCompleted AccountCollectionFlowStatus = "completed"
	AccountCollectionFlowStatusExpired   AccountCollectionFlowStatus = "expired"
	AccountCollectionFlowStatusPending   AccountCollectionFlowStatus = "pending"
)

type AccountCollectionFlowNewParams struct {
	// Required.
	CounterpartyID     param.Field[string]                                           `json:"counterparty_id,required" format:"uuid"`
	PaymentTypes       param.Field[[]string]                                         `json:"payment_types,required"`
	ReceivingCountries param.Field[[]AccountCollectionFlowNewParamsReceivingCountry] `json:"receiving_countries"`
}

func (r AccountCollectionFlowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional. Array of 3-digit ISO country codes.
type AccountCollectionFlowNewParamsReceivingCountry string

const (
	AccountCollectionFlowNewParamsReceivingCountryUsa AccountCollectionFlowNewParamsReceivingCountry = "USA"
	AccountCollectionFlowNewParamsReceivingCountryAus AccountCollectionFlowNewParamsReceivingCountry = "AUS"
	AccountCollectionFlowNewParamsReceivingCountryBel AccountCollectionFlowNewParamsReceivingCountry = "BEL"
	AccountCollectionFlowNewParamsReceivingCountryCan AccountCollectionFlowNewParamsReceivingCountry = "CAN"
	AccountCollectionFlowNewParamsReceivingCountryChl AccountCollectionFlowNewParamsReceivingCountry = "CHL"
	AccountCollectionFlowNewParamsReceivingCountryChn AccountCollectionFlowNewParamsReceivingCountry = "CHN"
	AccountCollectionFlowNewParamsReceivingCountryCol AccountCollectionFlowNewParamsReceivingCountry = "COL"
	AccountCollectionFlowNewParamsReceivingCountryFra AccountCollectionFlowNewParamsReceivingCountry = "FRA"
	AccountCollectionFlowNewParamsReceivingCountryDeu AccountCollectionFlowNewParamsReceivingCountry = "DEU"
	AccountCollectionFlowNewParamsReceivingCountryHkg AccountCollectionFlowNewParamsReceivingCountry = "HKG"
	AccountCollectionFlowNewParamsReceivingCountryInd AccountCollectionFlowNewParamsReceivingCountry = "IND"
	AccountCollectionFlowNewParamsReceivingCountryIrl AccountCollectionFlowNewParamsReceivingCountry = "IRL"
	AccountCollectionFlowNewParamsReceivingCountryIta AccountCollectionFlowNewParamsReceivingCountry = "ITA"
	AccountCollectionFlowNewParamsReceivingCountryMex AccountCollectionFlowNewParamsReceivingCountry = "MEX"
	AccountCollectionFlowNewParamsReceivingCountryNld AccountCollectionFlowNewParamsReceivingCountry = "NLD"
	AccountCollectionFlowNewParamsReceivingCountryPer AccountCollectionFlowNewParamsReceivingCountry = "PER"
	AccountCollectionFlowNewParamsReceivingCountryEsp AccountCollectionFlowNewParamsReceivingCountry = "ESP"
	AccountCollectionFlowNewParamsReceivingCountryGbr AccountCollectionFlowNewParamsReceivingCountry = "GBR"
)

type AccountCollectionFlowUpdateParams struct {
	// Required. The updated status of the account collection flow. Can only be used to
	// mark a flow as `cancelled`.
	Status param.Field[AccountCollectionFlowUpdateParamsStatus] `json:"status,required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
