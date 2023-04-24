package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LineItemGetParamsItemizableType string

const (
	LineItemGetParamsItemizableTypeExpectedPayments LineItemGetParamsItemizableType = "expected_payments"
	LineItemGetParamsItemizableTypePaymentOrders    LineItemGetParamsItemizableType = "payment_orders"
)

type LineItemUpdateParams struct {
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LineItemUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r LineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LineItemUpdateParamsItemizableType string

const (
	LineItemUpdateParamsItemizableTypeExpectedPayments LineItemUpdateParamsItemizableType = "expected_payments"
	LineItemUpdateParamsItemizableTypePaymentOrders    LineItemUpdateParamsItemizableType = "payment_orders"
)

type LineItemListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes LineItemListParams into a url.Values of the query parameters
// associated with this value
func (r LineItemListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LineItemListParamsItemizableType string

const (
	LineItemListParamsItemizableTypeExpectedPayments LineItemListParamsItemizableType = "expected_payments"
	LineItemListParamsItemizableTypePaymentOrders    LineItemListParamsItemizableType = "payment_orders"
)
