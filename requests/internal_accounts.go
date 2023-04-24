package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type InternalAccountNewParams struct {
	// The identifier of the financial institution the account belongs to.
	ConnectionID field.Field[string] `json:"connection_id,required"`
	// The nickname of the account.
	Name field.Field[string] `json:"name,required"`
	// The legal name of the entity which owns the account.
	PartyName field.Field[string] `json:"party_name,required"`
	// The address associated with the owner or null.
	PartyAddress field.Field[InternalAccountNewParamsPartyAddress] `json:"party_address"`
	// Either "USD" or "CAD". Internal accounts created at Increase only supports
	// "USD".
	Currency field.Field[InternalAccountNewParamsCurrency] `json:"currency,required"`
	// The identifier of the entity at Increase which owns the account.
	EntityID field.Field[string] `json:"entity_id"`
	// The parent internal account of this new account.
	ParentAccountID field.Field[string] `json:"parent_account_id"`
	// The Counterparty associated to this account.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
}

// MarshalJSON serializes InternalAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type InternalAccountNewParamsPartyAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InternalAccountNewParamsCurrency string

const (
	InternalAccountNewParamsCurrencyUsd InternalAccountNewParamsCurrency = "USD"
	InternalAccountNewParamsCurrencyCad InternalAccountNewParamsCurrency = "CAD"
)

type InternalAccountUpdateParams struct {
	// The nickname for the internal account.
	Name field.Field[string] `json:"name"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The parent internal account for this account.
	ParentAccountID field.Field[string] `json:"parent_account_id"`
	// The Counterparty associated to this account.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
}

// MarshalJSON serializes InternalAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type InternalAccountListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// The currency associated with the internal account.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// The counterparty associated with the internal account.
	CounterpartyID field.Field[string] `query:"counterparty_id"`
	// The type of payment that can be made by the internal account.
	PaymentType field.Field[InternalAccountListParamsPaymentType] `query:"payment_type"`
	// The direction of payments that can be made by internal account.
	PaymentDirection field.Field[InternalAccountListParamsPaymentDirection] `query:"payment_direction"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes InternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r InternalAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type InternalAccountListParamsPaymentType string

const (
	InternalAccountListParamsPaymentTypeACH         InternalAccountListParamsPaymentType = "ach"
	InternalAccountListParamsPaymentTypeAuBecs      InternalAccountListParamsPaymentType = "au_becs"
	InternalAccountListParamsPaymentTypeBacs        InternalAccountListParamsPaymentType = "bacs"
	InternalAccountListParamsPaymentTypeBook        InternalAccountListParamsPaymentType = "book"
	InternalAccountListParamsPaymentTypeCard        InternalAccountListParamsPaymentType = "card"
	InternalAccountListParamsPaymentTypeCheck       InternalAccountListParamsPaymentType = "check"
	InternalAccountListParamsPaymentTypeCrossBorder InternalAccountListParamsPaymentType = "cross_border"
	InternalAccountListParamsPaymentTypeEft         InternalAccountListParamsPaymentType = "eft"
	InternalAccountListParamsPaymentTypeInterac     InternalAccountListParamsPaymentType = "interac"
	InternalAccountListParamsPaymentTypeMasav       InternalAccountListParamsPaymentType = "masav"
	InternalAccountListParamsPaymentTypeNeft        InternalAccountListParamsPaymentType = "neft"
	InternalAccountListParamsPaymentTypeProvxchange InternalAccountListParamsPaymentType = "provxchange"
	InternalAccountListParamsPaymentTypeRtp         InternalAccountListParamsPaymentType = "rtp"
	InternalAccountListParamsPaymentTypeSen         InternalAccountListParamsPaymentType = "sen"
	InternalAccountListParamsPaymentTypeSepa        InternalAccountListParamsPaymentType = "sepa"
	InternalAccountListParamsPaymentTypeSignet      InternalAccountListParamsPaymentType = "signet"
	InternalAccountListParamsPaymentTypeWire        InternalAccountListParamsPaymentType = "wire"
)

type InternalAccountListParamsPaymentDirection string

const (
	InternalAccountListParamsPaymentDirectionCredit InternalAccountListParamsPaymentDirection = "credit"
	InternalAccountListParamsPaymentDirectionDebit  InternalAccountListParamsPaymentDirection = "debit"
)
