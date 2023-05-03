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

// PaymentFlowService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewPaymentFlowService]
// method instead.
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
	path := fmt.Sprintf("api/payment_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update payment_flow
func (r *PaymentFlowService) Update(ctx context.Context, id string, body PaymentFlowUpdateParams, opts ...option.RequestOption) (res *PaymentFlow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_flows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list payment_flows
func (r *PaymentFlowService) List(ctx context.Context, query PaymentFlowListParams, opts ...option.RequestOption) (res *shared.Page[PaymentFlow], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *PaymentFlowService) ListAutoPaging(ctx context.Context, query PaymentFlowListParams, opts ...option.RequestOption) *shared.PageAutoPager[PaymentFlow] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type PaymentFlow struct {
	ID     string `json:"id" format:"uuid"`
	Object string `json:"object"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The client token of the payment flow. This token can be used to embed a payment
	// workflow in your client-side application.
	ClientToken string `json:"client_token"`
	// The current status of the payment flow. One of `pending`, `completed`,
	// `expired`, or `cancelled`.
	Status PaymentFlowStatus `json:"status"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount int64 `json:"amount"`
	// The currency of the payment.
	Currency string `json:"currency"`
	// Describes the direction money is flowing in the transaction. Can only be
	// `debit`. A `debit` pulls money from someone else's account to your own.
	Direction PaymentFlowDirection `json:"direction"`
	// The ID of a counterparty associated with the payment. As part of the payment
	// workflow an external account will be associated with this counterparty.
	CounterpartyID string `json:"counterparty_id,nullable" format:"uuid"`
	// If present, the ID of the external account created using this flow.
	ReceivingAccountID string `json:"receiving_account_id,nullable" format:"uuid"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID string `json:"originating_account_id,nullable" format:"uuid"`
	// If present, the ID of the payment order created using this flow.
	PaymentOrderID string `json:"payment_order_id,nullable" format:"uuid"`
	JSON           paymentFlowJSON
}

// paymentFlowJSON contains the JSON metadata for the struct [PaymentFlow]
type paymentFlowJSON struct {
	ID                   apijson.Field
	Object               apijson.Field
	LiveMode             apijson.Field
	CreatedAt            apijson.Field
	UpdatedAt            apijson.Field
	ClientToken          apijson.Field
	Status               apijson.Field
	Amount               apijson.Field
	Currency             apijson.Field
	Direction            apijson.Field
	CounterpartyID       apijson.Field
	ReceivingAccountID   apijson.Field
	OriginatingAccountID apijson.Field
	PaymentOrderID       apijson.Field
	raw                  string
	Extras               map[string]apijson.Field
}

func (r *PaymentFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentFlowStatus string

const (
	PaymentFlowStatusCancelled PaymentFlowStatus = "cancelled"
	PaymentFlowStatusCompleted PaymentFlowStatus = "completed"
	PaymentFlowStatusExpired   PaymentFlowStatus = "expired"
	PaymentFlowStatusPending   PaymentFlowStatus = "pending"
)

type PaymentFlowDirection string

const (
	PaymentFlowDirectionCredit PaymentFlowDirection = "credit"
	PaymentFlowDirectionDebit  PaymentFlowDirection = "debit"
)

type PaymentFlowNewParams struct {
	// Required. Value in specified currency's smallest unit. e.g. $10 would be
	// represented as 1000. Can be any integer up to 36 digits.
	Amount field.Field[int64] `json:"amount,required"`
	// Required. The currency of the payment.
	Currency field.Field[string] `json:"currency,required"`
	// Required. Describes the direction money is flowing in the transaction. Can only
	// be `debit`. A `debit` pulls money from someone else's account to your own.
	Direction field.Field[PaymentFlowNewParamsDirection] `json:"direction,required"`
	// Required. The ID of a counterparty associated with the payment. As part of the
	// payment workflow an external account will be associated with this model.
	CounterpartyID field.Field[string] `json:"counterparty_id,required" format:"uuid"`
	// Required. The ID of one of your organization's internal accounts.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required" format:"uuid"`
}

func (r PaymentFlowNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentFlowNewParamsDirection string

const (
	PaymentFlowNewParamsDirectionCredit PaymentFlowNewParamsDirection = "credit"
	PaymentFlowNewParamsDirectionDebit  PaymentFlowNewParamsDirection = "debit"
)

type PaymentFlowUpdateParams struct {
	// Required. The updated status of the payment flow. Can only be used to mark a
	// flow as `cancelled`.
	Status field.Field[PaymentFlowUpdateParamsStatus] `json:"status,required"`
}

func (r PaymentFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentFlowUpdateParamsStatus string

const (
	PaymentFlowUpdateParamsStatusCancelled PaymentFlowUpdateParamsStatus = "cancelled"
)

type PaymentFlowListParams struct {
	AfterCursor          field.Field[string] `query:"after_cursor,nullable"`
	PerPage              field.Field[int64]  `query:"per_page"`
	ClientToken          field.Field[string] `query:"client_token"`
	Status               field.Field[string] `query:"status"`
	CounterpartyID       field.Field[string] `query:"counterparty_id"`
	ReceivingAccountID   field.Field[string] `query:"receiving_account_id"`
	OriginatingAccountID field.Field[string] `query:"originating_account_id"`
	PaymentOrderID       field.Field[string] `query:"payment_order_id"`
}

// URLQuery serializes [PaymentFlowListParams]'s query parameters as `url.Values`.
func (r PaymentFlowListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
