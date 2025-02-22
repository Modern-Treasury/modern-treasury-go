// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
)

// PaymentFlowService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentFlowService] method instead.
type PaymentFlowService struct {
	Options []option.RequestOption
}

// NewPaymentFlowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentFlowService(opts ...option.RequestOption) (r *PaymentFlowService) {
	r = &PaymentFlowService{}
	r.Options = opts
	return
}

// create payment_flow
func (r *PaymentFlowService) New(ctx context.Context, body PaymentFlowNewParams, opts ...option.RequestOption) (res *PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/payment_flows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get payment_flow
func (r *PaymentFlowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update payment_flow
func (r *PaymentFlowService) Update(ctx context.Context, id string, body PaymentFlowUpdateParams, opts ...option.RequestOption) (res *PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list payment_flows
func (r *PaymentFlowService) List(ctx context.Context, query PaymentFlowListParams, opts ...option.RequestOption) (res *pagination.Page[PaymentFlow], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_flows"
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

// list payment_flows
func (r *PaymentFlowService) ListAutoPaging(ctx context.Context, query PaymentFlowListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PaymentFlow] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type PaymentFlow struct {
	ID string `json:"id" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount int64 `json:"amount"`
	// The client token of the payment flow. This token can be used to embed a payment
	// workflow in your client-side application.
	ClientToken string `json:"client_token"`
	// The ID of a counterparty associated with the payment. As part of the payment
	// workflow an external account will be associated with this counterparty.
	CounterpartyID string    `json:"counterparty_id,nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at" format:"date-time"`
	// The currency of the payment.
	Currency string `json:"currency"`
	// Describes the direction money is flowing in the transaction. Can only be
	// `debit`. A `debit` pulls money from someone else's account to your own.
	Direction PaymentFlowDirection `json:"direction"`
	// The due date for the flow. Can only be passed in when
	// `effective_date_selection_enabled` is `true`.
	DueDate time.Time `json:"due_date,nullable" format:"date"`
	// When `true`, your end-user can schedule the payment `effective_date` while
	// completing the pre-built UI.
	EffectiveDateSelectionEnabled bool `json:"effective_date_selection_enabled"`
	// When `verified` and `external_account_collection` is `enabled`, filters the list
	// of external accounts your end-user can select to those with a
	// `verification_status` of `verified`.
	ExistingExternalAccountsFilter PaymentFlowExistingExternalAccountsFilter `json:"existing_external_accounts_filter,nullable"`
	// When `enabled`, your end-user can select from an existing external account when
	// completing the flow. When `disabled`, your end-user must add new payment details
	// when completing the flow.
	ExternalAccountCollection PaymentFlowExternalAccountCollection `json:"external_account_collection"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode"`
	Object   string `json:"object"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID string `json:"originating_account_id,nullable" format:"uuid"`
	// If present, the ID of the payment order created using this flow.
	PaymentOrderID string `json:"payment_order_id,nullable" format:"uuid"`
	// If present, the ID of the external account created using this flow.
	ReceivingAccountID string `json:"receiving_account_id,nullable" format:"uuid"`
	// This field is set after your end-user selects a payment date while completing
	// the pre-built UI. This field is always `null` unless
	// `effective_date_selection_enabled` is `true`.
	SelectedEffectiveDate time.Time `json:"selected_effective_date,nullable" format:"date"`
	// The current status of the payment flow. One of `pending`, `completed`,
	// `expired`, or `cancelled`.
	Status    PaymentFlowStatus `json:"status"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      paymentFlowJSON   `json:"-"`
}

// paymentFlowJSON contains the JSON metadata for the struct [PaymentFlow]
type paymentFlowJSON struct {
	ID                             apijson.Field
	Amount                         apijson.Field
	ClientToken                    apijson.Field
	CounterpartyID                 apijson.Field
	CreatedAt                      apijson.Field
	Currency                       apijson.Field
	Direction                      apijson.Field
	DueDate                        apijson.Field
	EffectiveDateSelectionEnabled  apijson.Field
	ExistingExternalAccountsFilter apijson.Field
	ExternalAccountCollection      apijson.Field
	LiveMode                       apijson.Field
	Object                         apijson.Field
	OriginatingAccountID           apijson.Field
	PaymentOrderID                 apijson.Field
	ReceivingAccountID             apijson.Field
	SelectedEffectiveDate          apijson.Field
	Status                         apijson.Field
	UpdatedAt                      apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PaymentFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFlowJSON) RawJSON() string {
	return r.raw
}

// Describes the direction money is flowing in the transaction. Can only be
// `debit`. A `debit` pulls money from someone else's account to your own.
type PaymentFlowDirection string

const (
	PaymentFlowDirectionCredit PaymentFlowDirection = "credit"
	PaymentFlowDirectionDebit  PaymentFlowDirection = "debit"
)

func (r PaymentFlowDirection) IsKnown() bool {
	switch r {
	case PaymentFlowDirectionCredit, PaymentFlowDirectionDebit:
		return true
	}
	return false
}

// When `verified` and `external_account_collection` is `enabled`, filters the list
// of external accounts your end-user can select to those with a
// `verification_status` of `verified`.
type PaymentFlowExistingExternalAccountsFilter string

const (
	PaymentFlowExistingExternalAccountsFilterVerified PaymentFlowExistingExternalAccountsFilter = "verified"
)

func (r PaymentFlowExistingExternalAccountsFilter) IsKnown() bool {
	switch r {
	case PaymentFlowExistingExternalAccountsFilterVerified:
		return true
	}
	return false
}

// When `enabled`, your end-user can select from an existing external account when
// completing the flow. When `disabled`, your end-user must add new payment details
// when completing the flow.
type PaymentFlowExternalAccountCollection string

const (
	PaymentFlowExternalAccountCollectionDisabled PaymentFlowExternalAccountCollection = "disabled"
	PaymentFlowExternalAccountCollectionEnabled  PaymentFlowExternalAccountCollection = "enabled"
)

func (r PaymentFlowExternalAccountCollection) IsKnown() bool {
	switch r {
	case PaymentFlowExternalAccountCollectionDisabled, PaymentFlowExternalAccountCollectionEnabled:
		return true
	}
	return false
}

// The current status of the payment flow. One of `pending`, `completed`,
// `expired`, or `cancelled`.
type PaymentFlowStatus string

const (
	PaymentFlowStatusCancelled PaymentFlowStatus = "cancelled"
	PaymentFlowStatusCompleted PaymentFlowStatus = "completed"
	PaymentFlowStatusExpired   PaymentFlowStatus = "expired"
	PaymentFlowStatusPending   PaymentFlowStatus = "pending"
)

func (r PaymentFlowStatus) IsKnown() bool {
	switch r {
	case PaymentFlowStatusCancelled, PaymentFlowStatusCompleted, PaymentFlowStatusExpired, PaymentFlowStatusPending:
		return true
	}
	return false
}

type PaymentFlowNewParams struct {
	// Required. Value in specified currency's smallest unit. e.g. $10 would be
	// represented as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// Required. The ID of a counterparty associated with the payment. As part of the
	// payment workflow an external account will be associated with this model.
	CounterpartyID param.Field[string] `json:"counterparty_id,required" format:"uuid"`
	// Required. The currency of the payment.
	Currency param.Field[string] `json:"currency,required"`
	// Required. Describes the direction money is flowing in the transaction. Can only
	// be `debit`. A `debit` pulls money from someone else's account to your own.
	Direction param.Field[PaymentFlowNewParamsDirection] `json:"direction,required"`
	// Required. The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Optional. Can only be passed in when `effective_date_selection_enabled` is
	// `true`. When set, the due date is shown to your end-user in the pre-built UI as
	// they are selecting a payment `effective_date`.
	DueDate param.Field[time.Time] `json:"due_date" format:"date"`
}

func (r PaymentFlowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required. Describes the direction money is flowing in the transaction. Can only
// be `debit`. A `debit` pulls money from someone else's account to your own.
type PaymentFlowNewParamsDirection string

const (
	PaymentFlowNewParamsDirectionCredit PaymentFlowNewParamsDirection = "credit"
	PaymentFlowNewParamsDirectionDebit  PaymentFlowNewParamsDirection = "debit"
)

func (r PaymentFlowNewParamsDirection) IsKnown() bool {
	switch r {
	case PaymentFlowNewParamsDirectionCredit, PaymentFlowNewParamsDirectionDebit:
		return true
	}
	return false
}

type PaymentFlowUpdateParams struct {
	// Required. The updated status of the payment flow. Can only be used to mark a
	// flow as `cancelled`.
	Status param.Field[PaymentFlowUpdateParamsStatus] `json:"status,required"`
}

func (r PaymentFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required. The updated status of the payment flow. Can only be used to mark a
// flow as `cancelled`.
type PaymentFlowUpdateParamsStatus string

const (
	PaymentFlowUpdateParamsStatusCancelled PaymentFlowUpdateParamsStatus = "cancelled"
)

func (r PaymentFlowUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PaymentFlowUpdateParamsStatusCancelled:
		return true
	}
	return false
}

type PaymentFlowListParams struct {
	AfterCursor          param.Field[string] `query:"after_cursor"`
	ClientToken          param.Field[string] `query:"client_token"`
	CounterpartyID       param.Field[string] `query:"counterparty_id"`
	OriginatingAccountID param.Field[string] `query:"originating_account_id"`
	PaymentOrderID       param.Field[string] `query:"payment_order_id"`
	PerPage              param.Field[int64]  `query:"per_page"`
	ReceivingAccountID   param.Field[string] `query:"receiving_account_id"`
	Status               param.Field[string] `query:"status"`
}

// URLQuery serializes [PaymentFlowListParams]'s query parameters as `url.Values`.
func (r PaymentFlowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
