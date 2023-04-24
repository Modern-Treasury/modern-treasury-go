package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type VirtualAccountNewParams struct {
	// The name of the virtual account.
	Name field.Field[string] `json:"name,required"`
	// An optional description for internal use.
	Description field.Field[string] `json:"description"`
	// The ID of the counterparty that the virtual account belongs to.
	CounterpartyID field.Field[string] `json:"counterparty_id" format:"uuid"`
	// The ID of the internal account that this virtual account is associated with.
	InternalAccountID field.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// An array of account detail objects.
	AccountDetails field.Field[[]AccountDetail] `json:"account_details"`
	// An array of routing detail objects.
	RoutingDetails field.Field[[]RoutingDetail] `json:"routing_details"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID field.Field[string] `json:"debit_ledger_account_id" format:"uuid"`
	// The ID of a credit normal ledger account. When money leaves the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID field.Field[string] `json:"credit_ledger_account_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes VirtualAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r VirtualAccountNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type VirtualAccountUpdateParams struct {
	Name           field.Field[string]            `json:"name,nullable"`
	CounterpartyID field.Field[string]            `json:"counterparty_id" format:"uuid"`
	Metadata       field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes VirtualAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r VirtualAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type VirtualAccountListParams struct {
	AfterCursor       field.Field[string] `query:"after_cursor,nullable"`
	PerPage           field.Field[int64]  `query:"per_page"`
	InternalAccountID field.Field[string] `query:"internal_account_id" format:"uuid"`
	CounterpartyID    field.Field[string] `query:"counterparty_id" format:"uuid"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes VirtualAccountListParams into a url.Values of the query
// parameters associated with this value
func (r VirtualAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
