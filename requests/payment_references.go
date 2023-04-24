package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type PaymentReferenceListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// The id of the referenceable to search for. Must be accompanied by the
	// referenceable_type or will return an error.
	ReferenceableID field.Field[string] `query:"referenceable_id"`
	// One of the referenceable types. This must be accompanied by the id of the
	// referenceable or will return an error.
	ReferenceableType field.Field[PaymentReferenceListParamsReferenceableType] `query:"referenceable_type"`
	// The actual reference number assigned by the bank.
	ReferenceNumber field.Field[string] `query:"reference_number"`
}

// URLQuery serializes PaymentReferenceListParams into a url.Values of the query
// parameters associated with this value
func (r PaymentReferenceListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type PaymentReferenceListParamsReferenceableType string

const (
	PaymentReferenceListParamsReferenceableTypePaymentOrder PaymentReferenceListParamsReferenceableType = "payment_order"
	PaymentReferenceListParamsReferenceableTypeReturn       PaymentReferenceListParamsReferenceableType = "return"
	PaymentReferenceListParamsReferenceableTypeReversal     PaymentReferenceListParamsReferenceableType = "reversal"
)
