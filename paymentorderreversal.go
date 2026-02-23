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
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// PaymentOrderReversalService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentOrderReversalService] method instead.
type PaymentOrderReversalService struct {
	Options []option.RequestOption
}

// NewPaymentOrderReversalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentOrderReversalService(opts ...option.RequestOption) (r *PaymentOrderReversalService) {
	r = &PaymentOrderReversalService{}
	r.Options = opts
	return
}

// Create a reversal for a payment order.
func (r *PaymentOrderReversalService) New(ctx context.Context, paymentOrderID string, body PaymentOrderReversalNewParams, opts ...option.RequestOption) (res *Reversal, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentOrderID == "" {
		err = errors.New("missing required payment_order_id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_orders/%s/reversals", paymentOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single reversal of a payment order.
func (r *PaymentOrderReversalService) Get(ctx context.Context, paymentOrderID string, reversalID string, opts ...option.RequestOption) (res *Reversal, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentOrderID == "" {
		err = errors.New("missing required payment_order_id parameter")
		return
	}
	if reversalID == "" {
		err = errors.New("missing required reversal_id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_orders/%s/reversals/%s", paymentOrderID, reversalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all reversals of a payment order.
func (r *PaymentOrderReversalService) List(ctx context.Context, paymentOrderID string, query PaymentOrderReversalListParams, opts ...option.RequestOption) (res *pagination.Page[Reversal], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if paymentOrderID == "" {
		err = errors.New("missing required payment_order_id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_orders/%s/reversals", paymentOrderID)
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

// Get a list of all reversals of a payment order.
func (r *PaymentOrderReversalService) ListAutoPaging(ctx context.Context, paymentOrderID string, query PaymentOrderReversalListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Reversal] {
	return pagination.NewPageAutoPager(r.List(ctx, paymentOrderID, query, opts...))
}

type Reversal struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required,nullable" format:"uuid"`
	// The ID of the ledger transaction linked to the reversal.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The ID of the relevant Payment Order.
	PaymentOrderID string `json:"payment_order_id,required,nullable" format:"uuid"`
	// The reason for the reversal.
	Reason ReversalReason `json:"reason,required"`
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus ReversalReconciliationStatus `json:"reconciliation_status,required"`
	// The current status of the reversal.
	Status         ReversalStatus `json:"status,required"`
	TransactionIDs []string       `json:"transaction_ids,required" format:"uuid"`
	UpdatedAt      time.Time      `json:"updated_at,required" format:"date-time"`
	JSON           reversalJSON   `json:"-"`
}

// reversalJSON contains the JSON metadata for the struct [Reversal]
type reversalJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	InternalAccountID    apijson.Field
	LedgerTransactionID  apijson.Field
	LiveMode             apijson.Field
	Metadata             apijson.Field
	Object               apijson.Field
	PaymentOrderID       apijson.Field
	Reason               apijson.Field
	ReconciliationStatus apijson.Field
	Status               apijson.Field
	TransactionIDs       apijson.Field
	UpdatedAt            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Reversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reversalJSON) RawJSON() string {
	return r.raw
}

// The reason for the reversal.
type ReversalReason string

const (
	ReversalReasonDuplicate                 ReversalReason = "duplicate"
	ReversalReasonIncorrectAmount           ReversalReason = "incorrect_amount"
	ReversalReasonIncorrectReceivingAccount ReversalReason = "incorrect_receiving_account"
	ReversalReasonDateEarlierThanIntended   ReversalReason = "date_earlier_than_intended"
	ReversalReasonDateLaterThanIntended     ReversalReason = "date_later_than_intended"
)

func (r ReversalReason) IsKnown() bool {
	switch r {
	case ReversalReasonDuplicate, ReversalReasonIncorrectAmount, ReversalReasonIncorrectReceivingAccount, ReversalReasonDateEarlierThanIntended, ReversalReasonDateLaterThanIntended:
		return true
	}
	return false
}

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type ReversalReconciliationStatus string

const (
	ReversalReconciliationStatusUnreconciled          ReversalReconciliationStatus = "unreconciled"
	ReversalReconciliationStatusTentativelyReconciled ReversalReconciliationStatus = "tentatively_reconciled"
	ReversalReconciliationStatusReconciled            ReversalReconciliationStatus = "reconciled"
)

func (r ReversalReconciliationStatus) IsKnown() bool {
	switch r {
	case ReversalReconciliationStatusUnreconciled, ReversalReconciliationStatusTentativelyReconciled, ReversalReconciliationStatusReconciled:
		return true
	}
	return false
}

// The current status of the reversal.
type ReversalStatus string

const (
	ReversalStatusCompleted  ReversalStatus = "completed"
	ReversalStatusFailed     ReversalStatus = "failed"
	ReversalStatusPending    ReversalStatus = "pending"
	ReversalStatusProcessing ReversalStatus = "processing"
	ReversalStatusReturned   ReversalStatus = "returned"
	ReversalStatusSent       ReversalStatus = "sent"
)

func (r ReversalStatus) IsKnown() bool {
	switch r {
	case ReversalStatusCompleted, ReversalStatusFailed, ReversalStatusPending, ReversalStatusProcessing, ReversalStatusReturned, ReversalStatusSent:
		return true
	}
	return false
}

type PaymentOrderReversalNewParams struct {
	// The reason for the reversal. Must be one of `duplicate`, `incorrect_amount`,
	// `incorrect_receiving_account`, `date_earlier_than_intended`,
	// `date_later_than_intended`.
	Reason param.Field[PaymentOrderReversalNewParamsReason] `json:"reason,required"`
	// Specifies a ledger transaction object that will be created with the reversal. If
	// the ledger transaction cannot be created, then the reversal creation will fail.
	// The resulting ledger transaction will mirror the status of the reversal.
	LedgerTransaction param.Field[shared.LedgerTransactionCreateRequestParam] `json:"ledger_transaction"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderReversalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The reason for the reversal. Must be one of `duplicate`, `incorrect_amount`,
// `incorrect_receiving_account`, `date_earlier_than_intended`,
// `date_later_than_intended`.
type PaymentOrderReversalNewParamsReason string

const (
	PaymentOrderReversalNewParamsReasonDuplicate                 PaymentOrderReversalNewParamsReason = "duplicate"
	PaymentOrderReversalNewParamsReasonIncorrectAmount           PaymentOrderReversalNewParamsReason = "incorrect_amount"
	PaymentOrderReversalNewParamsReasonIncorrectReceivingAccount PaymentOrderReversalNewParamsReason = "incorrect_receiving_account"
	PaymentOrderReversalNewParamsReasonDateEarlierThanIntended   PaymentOrderReversalNewParamsReason = "date_earlier_than_intended"
	PaymentOrderReversalNewParamsReasonDateLaterThanIntended     PaymentOrderReversalNewParamsReason = "date_later_than_intended"
)

func (r PaymentOrderReversalNewParamsReason) IsKnown() bool {
	switch r {
	case PaymentOrderReversalNewParamsReasonDuplicate, PaymentOrderReversalNewParamsReasonIncorrectAmount, PaymentOrderReversalNewParamsReasonIncorrectReceivingAccount, PaymentOrderReversalNewParamsReasonDateEarlierThanIntended, PaymentOrderReversalNewParamsReasonDateLaterThanIntended:
		return true
	}
	return false
}

type PaymentOrderReversalListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [PaymentOrderReversalListParams]'s query parameters as
// `url.Values`.
func (r PaymentOrderReversalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
