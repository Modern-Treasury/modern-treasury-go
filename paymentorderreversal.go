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

// PaymentOrderReversalService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewPaymentOrderReversalService] method instead.
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
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s/reversals", paymentOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single reversal of a payment order.
func (r *PaymentOrderReversalService) Get(ctx context.Context, paymentOrderID string, reversalID string, opts ...option.RequestOption) (res *Reversal, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s/reversals/%s", paymentOrderID, reversalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all reversals of a payment order.
func (r *PaymentOrderReversalService) List(ctx context.Context, paymentOrderID string, query PaymentOrderReversalListParams, opts ...option.RequestOption) (res *shared.Page[Reversal], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *PaymentOrderReversalService) ListAutoPaging(ctx context.Context, paymentOrderID string, query PaymentOrderReversalListParams, opts ...option.RequestOption) *shared.PageAutoPager[Reversal] {
	return shared.NewPageAutoPager(r.List(ctx, paymentOrderID, query, opts...))
}

type Reversal struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
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
	// The current status of the reversal.
	Status    ReversalStatus `json:"status,required"`
	UpdatedAt time.Time      `json:"updated_at,required" format:"date-time"`
	JSON      reversalJSON
}

// reversalJSON contains the JSON metadata for the struct [Reversal]
type reversalJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	LiveMode       apijson.Field
	Metadata       apijson.Field
	Object         apijson.Field
	PaymentOrderID apijson.Field
	Reason         apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Reversal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type PaymentOrderReversalNewParams struct {
	// The reason for the reversal. Must be one of `duplicate`, `incorrect_amount`,
	// `incorrect_receiving_account`, `date_earlier_than_intended`,
	// `date_later_than_intended`.
	Reason param.Field[PaymentOrderReversalNewParamsReason] `json:"reason,required"`
	// Specifies a ledger transaction object that will be created with the reversal. If
	// the ledger transaction cannot be created, then the reversal creation will fail.
	// The resulting ledger transaction will mirror the status of the reversal.
	LedgerTransaction param.Field[PaymentOrderReversalNewParamsLedgerTransaction] `json:"ledger_transaction"`
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

// Specifies a ledger transaction object that will be created with the reversal. If
// the ledger transaction cannot be created, then the reversal creation will fail.
// The resulting ledger transaction will mirror the status of the reversal.
type PaymentOrderReversalNewParamsLedgerTransaction struct {
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]PaymentOrderReversalNewParamsLedgerTransactionLedgerEntry] `json:"ledger_entries,required"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[PaymentOrderReversalNewParamsLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[PaymentOrderReversalNewParamsLedgerTransactionStatus] `json:"status"`
}

func (r PaymentOrderReversalNewParamsLedgerTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderReversalNewParamsLedgerTransactionLedgerEntry struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r PaymentOrderReversalNewParamsLedgerTransactionLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type PaymentOrderReversalNewParamsLedgerTransactionLedgerableType string

const (
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeCounterparty          PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "counterparty"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeExpectedPayment       PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "expected_payment"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeIncomingPaymentDetail PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "incoming_payment_detail"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeInternalAccount       PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "internal_account"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeLineItem              PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "line_item"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypePaperItem             PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "paper_item"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypePaymentOrder          PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "payment_order"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypePaymentOrderAttempt   PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "payment_order_attempt"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeReturn                PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "return"
	PaymentOrderReversalNewParamsLedgerTransactionLedgerableTypeReversal              PaymentOrderReversalNewParamsLedgerTransactionLedgerableType = "reversal"
)

// To post a ledger transaction at creation, use `posted`.
type PaymentOrderReversalNewParamsLedgerTransactionStatus string

const (
	PaymentOrderReversalNewParamsLedgerTransactionStatusArchived PaymentOrderReversalNewParamsLedgerTransactionStatus = "archived"
	PaymentOrderReversalNewParamsLedgerTransactionStatusPending  PaymentOrderReversalNewParamsLedgerTransactionStatus = "pending"
	PaymentOrderReversalNewParamsLedgerTransactionStatusPosted   PaymentOrderReversalNewParamsLedgerTransactionStatus = "posted"
)

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
