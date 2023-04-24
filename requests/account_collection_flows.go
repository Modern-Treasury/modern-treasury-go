package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
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
	return pjson.MarshalRoot(r)
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
	return pjson.MarshalRoot(r)
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
	return query.Marshal(r)
}
