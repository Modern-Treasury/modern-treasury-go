// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
	"github.com/tidwall/gjson"
)

// EmbeddableFlowService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddableFlowService] method instead.
type EmbeddableFlowService struct {
	Options []option.RequestOption
}

// NewEmbeddableFlowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddableFlowService(opts ...option.RequestOption) (r *EmbeddableFlowService) {
	r = &EmbeddableFlowService{}
	r.Options = opts
	return
}

// create embeddable_flow
func (r *EmbeddableFlowService) New(ctx context.Context, body EmbeddableFlowNewParams, opts ...option.RequestOption) (res *EmbeddableFlow, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/embeddable_flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// get embeddable_flow
func (r *EmbeddableFlowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmbeddableFlow, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/embeddable_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// update embeddable_flow
func (r *EmbeddableFlowService) Update(ctx context.Context, id string, body EmbeddableFlowUpdateParams, opts ...option.RequestOption) (res *EmbeddableFlow, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/embeddable_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// list embeddable_flows
func (r *EmbeddableFlowService) List(ctx context.Context, query EmbeddableFlowListParams, opts ...option.RequestOption) (res *pagination.Page[EmbeddableFlow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/embeddable_flows"
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

// list embeddable_flows
func (r *EmbeddableFlowService) ListAutoPaging(ctx context.Context, query EmbeddableFlowListParams, opts ...option.RequestOption) *pagination.PageAutoPager[EmbeddableFlow] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type EmbeddableFlow struct {
	ID          string    `json:"id" api:"required" format:"uuid"`
	ClientToken string    `json:"client_token" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ExpiresAt   time.Time `json:"expires_at" api:"required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool                  `json:"live_mode" api:"required"`
	Object   string                `json:"object" api:"required"`
	Options  EmbeddableFlowOptions `json:"options" api:"required"`
	// The current status of the underlying flow.
	Status    EmbeddableFlowStatus   `json:"status" api:"required"`
	Template  EmbeddableFlowTemplate `json:"template" api:"required"`
	UpdatedAt time.Time              `json:"updated_at" api:"required" format:"date-time"`
	JSON      embeddableFlowJSON     `json:"-"`
}

// embeddableFlowJSON contains the JSON metadata for the struct [EmbeddableFlow]
type embeddableFlowJSON struct {
	ID          apijson.Field
	ClientToken apijson.Field
	CreatedAt   apijson.Field
	ExpiresAt   apijson.Field
	LiveMode    apijson.Field
	Object      apijson.Field
	Options     apijson.Field
	Status      apijson.Field
	Template    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmbeddableFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddableFlowJSON) RawJSON() string {
	return r.raw
}

type EmbeddableFlowOptions struct {
	Amount                         int64                                               `json:"amount"`
	CounterpartyID                 string                                              `json:"counterparty_id" api:"nullable" format:"uuid"`
	Currency                       string                                              `json:"currency"`
	Direction                      shared.TransactionDirection                         `json:"direction"`
	DueDate                        time.Time                                           `json:"due_date" api:"nullable" format:"date"`
	EffectiveDateSelectionEnabled  bool                                                `json:"effective_date_selection_enabled"`
	ExistingExternalAccountsFilter EmbeddableFlowOptionsExistingExternalAccountsFilter `json:"existing_external_accounts_filter" api:"nullable"`
	ExternalAccountCollection      EmbeddableFlowOptionsExternalAccountCollection      `json:"external_account_collection"`
	ExternalAccountID              string                                              `json:"external_account_id" api:"nullable" format:"uuid"`
	LegalEntityID                  string                                              `json:"legal_entity_id" format:"uuid"`
	OriginatingAccountID           string                                              `json:"originating_account_id" format:"uuid"`
	PaymentOrderID                 string                                              `json:"payment_order_id" api:"nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	PaymentTypes       interface{} `json:"payment_types"`
	ReceivingAccountID string      `json:"receiving_account_id" api:"nullable" format:"uuid"`
	// This field can have the runtime type of [[]string].
	ReceivingCountries    interface{}               `json:"receiving_countries"`
	SelectedEffectiveDate time.Time                 `json:"selected_effective_date" api:"nullable" format:"date"`
	Type                  EmbeddableFlowOptionsType `json:"type"`
	JSON                  embeddableFlowOptionsJSON `json:"-"`
	union                 EmbeddableFlowOptionsUnion
}

// embeddableFlowOptionsJSON contains the JSON metadata for the struct
// [EmbeddableFlowOptions]
type embeddableFlowOptionsJSON struct {
	Amount                         apijson.Field
	CounterpartyID                 apijson.Field
	Currency                       apijson.Field
	Direction                      apijson.Field
	DueDate                        apijson.Field
	EffectiveDateSelectionEnabled  apijson.Field
	ExistingExternalAccountsFilter apijson.Field
	ExternalAccountCollection      apijson.Field
	ExternalAccountID              apijson.Field
	LegalEntityID                  apijson.Field
	OriginatingAccountID           apijson.Field
	PaymentOrderID                 apijson.Field
	PaymentTypes                   apijson.Field
	ReceivingAccountID             apijson.Field
	ReceivingCountries             apijson.Field
	SelectedEffectiveDate          apijson.Field
	Type                           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r embeddableFlowOptionsJSON) RawJSON() string {
	return r.raw
}

func (r *EmbeddableFlowOptions) UnmarshalJSON(data []byte) (err error) {
	*r = EmbeddableFlowOptions{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EmbeddableFlowOptionsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [EmbeddableFlowOptionsObject],
// [EmbeddableFlowOptionsObject], [EmbeddableFlowOptionsObject].
func (r EmbeddableFlowOptions) AsUnion() EmbeddableFlowOptionsUnion {
	return r.union
}

// Union satisfied by [EmbeddableFlowOptionsObject], [EmbeddableFlowOptionsObject]
// or [EmbeddableFlowOptionsObject].
type EmbeddableFlowOptionsUnion interface {
	implementsEmbeddableFlowOptions()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EmbeddableFlowOptionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmbeddableFlowOptionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmbeddableFlowOptionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmbeddableFlowOptionsObject{}),
		},
	)
}

type EmbeddableFlowOptionsObject struct {
	LegalEntityID string                          `json:"legal_entity_id" api:"required" format:"uuid"`
	Type          EmbeddableFlowOptionsObjectType `json:"type" api:"required"`
	JSON          embeddableFlowOptionsObjectJSON `json:"-"`
}

// embeddableFlowOptionsObjectJSON contains the JSON metadata for the struct
// [EmbeddableFlowOptionsObject]
type embeddableFlowOptionsObjectJSON struct {
	LegalEntityID apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *EmbeddableFlowOptionsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r embeddableFlowOptionsObjectJSON) RawJSON() string {
	return r.raw
}

func (r EmbeddableFlowOptionsObject) implementsEmbeddableFlowOptions() {}

type EmbeddableFlowOptionsObjectType string

const (
	EmbeddableFlowOptionsObjectTypeIndividual EmbeddableFlowOptionsObjectType = "individual"
)

func (r EmbeddableFlowOptionsObjectType) IsKnown() bool {
	switch r {
	case EmbeddableFlowOptionsObjectTypeIndividual:
		return true
	}
	return false
}

type EmbeddableFlowOptionsExistingExternalAccountsFilter string

const (
	EmbeddableFlowOptionsExistingExternalAccountsFilterVerified EmbeddableFlowOptionsExistingExternalAccountsFilter = "verified"
)

func (r EmbeddableFlowOptionsExistingExternalAccountsFilter) IsKnown() bool {
	switch r {
	case EmbeddableFlowOptionsExistingExternalAccountsFilterVerified:
		return true
	}
	return false
}

type EmbeddableFlowOptionsExternalAccountCollection string

const (
	EmbeddableFlowOptionsExternalAccountCollectionEnabled  EmbeddableFlowOptionsExternalAccountCollection = "enabled"
	EmbeddableFlowOptionsExternalAccountCollectionDisabled EmbeddableFlowOptionsExternalAccountCollection = "disabled"
)

func (r EmbeddableFlowOptionsExternalAccountCollection) IsKnown() bool {
	switch r {
	case EmbeddableFlowOptionsExternalAccountCollectionEnabled, EmbeddableFlowOptionsExternalAccountCollectionDisabled:
		return true
	}
	return false
}

type EmbeddableFlowOptionsType string

const (
	EmbeddableFlowOptionsTypeIndividual EmbeddableFlowOptionsType = "individual"
)

func (r EmbeddableFlowOptionsType) IsKnown() bool {
	switch r {
	case EmbeddableFlowOptionsTypeIndividual:
		return true
	}
	return false
}

// The current status of the underlying flow.
type EmbeddableFlowStatus string

const (
	EmbeddableFlowStatusCancelled EmbeddableFlowStatus = "cancelled"
	EmbeddableFlowStatusCompleted EmbeddableFlowStatus = "completed"
	EmbeddableFlowStatusExpired   EmbeddableFlowStatus = "expired"
	EmbeddableFlowStatusPending   EmbeddableFlowStatus = "pending"
)

func (r EmbeddableFlowStatus) IsKnown() bool {
	switch r {
	case EmbeddableFlowStatusCancelled, EmbeddableFlowStatusCompleted, EmbeddableFlowStatusExpired, EmbeddableFlowStatusPending:
		return true
	}
	return false
}

type EmbeddableFlowTemplate string

const (
	EmbeddableFlowTemplateLegalEntityOnboarding EmbeddableFlowTemplate = "legal_entity_onboarding"
	EmbeddableFlowTemplateAccountCollection     EmbeddableFlowTemplate = "account_collection"
	EmbeddableFlowTemplatePayment               EmbeddableFlowTemplate = "payment"
)

func (r EmbeddableFlowTemplate) IsKnown() bool {
	switch r {
	case EmbeddableFlowTemplateLegalEntityOnboarding, EmbeddableFlowTemplateAccountCollection, EmbeddableFlowTemplatePayment:
		return true
	}
	return false
}

type EmbeddableFlowNewParams struct {
	// The parameters for the selected template's flow.
	Options  param.Field[EmbeddableFlowNewParamsOptionsUnion] `json:"options" api:"required"`
	Template param.Field[EmbeddableFlowNewParamsTemplate]     `json:"template" api:"required"`
}

func (r EmbeddableFlowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The parameters for the selected template's flow.
type EmbeddableFlowNewParamsOptions struct {
	// Required. Value in specified currency's smallest unit. e.g. $10 would be
	// represented as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount"`
	// Required.
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// Required. The currency of the payment.
	Currency param.Field[string] `json:"currency"`
	// Required. Describes the direction money is flowing in the transaction. Can only
	// be `debit`. A `debit` pulls money from someone else's account to your own.
	Direction param.Field[EmbeddableFlowNewParamsOptionsDirection] `json:"direction"`
	// Optional. Can only be passed in when `effective_date_selection_enabled` is
	// `true`. When set, the due date is shown to your end-user in the pre-built UI as
	// they are selecting a payment `effective_date`.
	DueDate param.Field[time.Time] `json:"due_date" format:"date"`
	// Required. The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string]                             `json:"originating_account_id" format:"uuid"`
	PaymentTypes         param.Field[interface{}]                        `json:"payment_types"`
	ReceivingCountries   param.Field[interface{}]                        `json:"receiving_countries"`
	Type                 param.Field[EmbeddableFlowNewParamsOptionsType] `json:"type"`
}

func (r EmbeddableFlowNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmbeddableFlowNewParamsOptions) implementsEmbeddableFlowNewParamsOptionsUnion() {}

// The parameters for the selected template's flow.
//
// Satisfied by [AccountCollectionFlowCreateRequestParam],
// [PaymentFlowCreateRequestParam], [EmbeddableFlowNewParamsOptionsType],
// [EmbeddableFlowNewParamsOptions].
type EmbeddableFlowNewParamsOptionsUnion interface {
	implementsEmbeddableFlowNewParamsOptionsUnion()
}

type EmbeddableFlowNewParamsOptionsType struct {
	Type param.Field[EmbeddableFlowNewParamsOptionsTypeType] `json:"type" api:"required"`
}

func (r EmbeddableFlowNewParamsOptionsType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmbeddableFlowNewParamsOptionsType) implementsEmbeddableFlowNewParamsOptionsUnion() {}

type EmbeddableFlowNewParamsOptionsTypeType string

const (
	EmbeddableFlowNewParamsOptionsTypeTypeIndividual EmbeddableFlowNewParamsOptionsTypeType = "individual"
)

func (r EmbeddableFlowNewParamsOptionsTypeType) IsKnown() bool {
	switch r {
	case EmbeddableFlowNewParamsOptionsTypeTypeIndividual:
		return true
	}
	return false
}

// Required. Describes the direction money is flowing in the transaction. Can only
// be `debit`. A `debit` pulls money from someone else's account to your own.
type EmbeddableFlowNewParamsOptionsDirection string

const (
	EmbeddableFlowNewParamsOptionsDirectionCredit EmbeddableFlowNewParamsOptionsDirection = "credit"
	EmbeddableFlowNewParamsOptionsDirectionDebit  EmbeddableFlowNewParamsOptionsDirection = "debit"
)

func (r EmbeddableFlowNewParamsOptionsDirection) IsKnown() bool {
	switch r {
	case EmbeddableFlowNewParamsOptionsDirectionCredit, EmbeddableFlowNewParamsOptionsDirectionDebit:
		return true
	}
	return false
}

type EmbeddableFlowNewParamsTemplate string

const (
	EmbeddableFlowNewParamsTemplateLegalEntityOnboarding EmbeddableFlowNewParamsTemplate = "legal_entity_onboarding"
	EmbeddableFlowNewParamsTemplateAccountCollection     EmbeddableFlowNewParamsTemplate = "account_collection"
	EmbeddableFlowNewParamsTemplatePayment               EmbeddableFlowNewParamsTemplate = "payment"
)

func (r EmbeddableFlowNewParamsTemplate) IsKnown() bool {
	switch r {
	case EmbeddableFlowNewParamsTemplateLegalEntityOnboarding, EmbeddableFlowNewParamsTemplateAccountCollection, EmbeddableFlowNewParamsTemplatePayment:
		return true
	}
	return false
}

type EmbeddableFlowUpdateParams struct {
	// The updated status of the flow. Can only be used to mark the flow as
	// `cancelled`.
	Status param.Field[EmbeddableFlowUpdateParamsStatus] `json:"status" api:"required"`
}

func (r EmbeddableFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The updated status of the flow. Can only be used to mark the flow as
// `cancelled`.
type EmbeddableFlowUpdateParamsStatus string

const (
	EmbeddableFlowUpdateParamsStatusCancelled EmbeddableFlowUpdateParamsStatus = "cancelled"
)

func (r EmbeddableFlowUpdateParamsStatus) IsKnown() bool {
	switch r {
	case EmbeddableFlowUpdateParamsStatusCancelled:
		return true
	}
	return false
}

type EmbeddableFlowListParams struct {
	AfterCursor param.Field[string]                           `query:"after_cursor"`
	PerPage     param.Field[int64]                            `query:"per_page"`
	Status      param.Field[string]                           `query:"status"`
	Template    param.Field[EmbeddableFlowListParamsTemplate] `query:"template"`
}

// URLQuery serializes [EmbeddableFlowListParams]'s query parameters as
// `url.Values`.
func (r EmbeddableFlowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EmbeddableFlowListParamsTemplate string

const (
	EmbeddableFlowListParamsTemplateLegalEntityOnboarding EmbeddableFlowListParamsTemplate = "legal_entity_onboarding"
	EmbeddableFlowListParamsTemplateAccountCollection     EmbeddableFlowListParamsTemplate = "account_collection"
	EmbeddableFlowListParamsTemplatePayment               EmbeddableFlowListParamsTemplate = "payment"
)

func (r EmbeddableFlowListParamsTemplate) IsKnown() bool {
	switch r {
	case EmbeddableFlowListParamsTemplateLegalEntityOnboarding, EmbeddableFlowListParamsTemplateAccountCollection, EmbeddableFlowListParamsTemplatePayment:
		return true
	}
	return false
}
