package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
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

// MarshalJSON serializes PaymentFlowNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r PaymentFlowNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
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

// MarshalJSON serializes PaymentFlowUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r PaymentFlowUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
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

// URLQuery serializes PaymentFlowListParams into a url.Values of the query
// parameters associated with this value
func (r PaymentFlowListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
