package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type TransactionUpdateParams struct {
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes TransactionUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type TransactionListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Specify `internal_account_id` if you wish to see transactions to/from a specific
	// account.
	InternalAccountID field.Field[string] `query:"internal_account_id" format:"uuid"`
	VirtualAccountID  field.Field[string] `query:"virtual_account_id" format:"uuid"`
	// Either `true` or `false`.
	Posted field.Field[bool] `query:"posted"`
	// Filters transactions with an `as_of_date` starting on or after the specified
	// date (YYYY-MM-DD).
	AsOfDateStart field.Field[time.Time] `query:"as_of_date_start" format:"date"`
	// Filters transactions with an `as_of_date` starting on or before the specified
	// date (YYYY-MM-DD).
	AsOfDateEnd      field.Field[time.Time] `query:"as_of_date_end" format:"date"`
	Direction        field.Field[string]    `query:"direction"`
	CounterpartyID   field.Field[string]    `query:"counterparty_id" format:"uuid"`
	PaymentType      field.Field[string]    `query:"payment_type"`
	TransactableType field.Field[string]    `query:"transactable_type"`
	// Filters for transactions including the queried string in the description.
	Description field.Field[string] `query:"description"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes TransactionListParams into a url.Values of the query
// parameters associated with this value
func (r TransactionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
