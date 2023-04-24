package responses

import (
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type RoutingNumberLookupRequest struct {
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number"`
	// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
	// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
	// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
	// respectively.
	RoutingNumberType RoutingNumberLookupRequestRoutingNumberType `json:"routing_number_type"`
	// An array of payment types that are supported for this routing number. This can
	// include `ach`, `wire`, `rtp`, `sepa`, `bacs`, `au_becs` currently.
	SupportedPaymentTypes []RoutingNumberLookupRequestSupportedPaymentTypes `json:"supported_payment_types"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The address of the bank.
	BankAddress RoutingNumberLookupRequestBankAddress `json:"bank_address"`
	// An object containing key-value pairs, each with a sanctions list as the key and
	// a boolean value representing whether the bank is on that particular sanctions
	// list. Currently, this includes eu_con, uk_hmt, us_ofac, and un sanctions lists.
	Sanctions interface{} `json:"sanctions"`
	JSON      RoutingNumberLookupRequestJSON
}

type RoutingNumberLookupRequestJSON struct {
	RoutingNumber         pjson.Metadata
	RoutingNumberType     pjson.Metadata
	SupportedPaymentTypes pjson.Metadata
	BankName              pjson.Metadata
	BankAddress           pjson.Metadata
	Sanctions             pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumberLookupRequest
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RoutingNumberLookupRequest) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RoutingNumberLookupRequestRoutingNumberType string

const (
	RoutingNumberLookupRequestRoutingNumberTypeAba        RoutingNumberLookupRequestRoutingNumberType = "aba"
	RoutingNumberLookupRequestRoutingNumberTypeAuBsb      RoutingNumberLookupRequestRoutingNumberType = "au_bsb"
	RoutingNumberLookupRequestRoutingNumberTypeCaCpa      RoutingNumberLookupRequestRoutingNumberType = "ca_cpa"
	RoutingNumberLookupRequestRoutingNumberTypeGBSortCode RoutingNumberLookupRequestRoutingNumberType = "gb_sort_code"
	RoutingNumberLookupRequestRoutingNumberTypeInIfsc     RoutingNumberLookupRequestRoutingNumberType = "in_ifsc"
	RoutingNumberLookupRequestRoutingNumberTypeSwift      RoutingNumberLookupRequestRoutingNumberType = "swift"
)

type RoutingNumberLookupRequestSupportedPaymentTypes string

const (
	RoutingNumberLookupRequestSupportedPaymentTypesACH         RoutingNumberLookupRequestSupportedPaymentTypes = "ach"
	RoutingNumberLookupRequestSupportedPaymentTypesAuBecs      RoutingNumberLookupRequestSupportedPaymentTypes = "au_becs"
	RoutingNumberLookupRequestSupportedPaymentTypesBacs        RoutingNumberLookupRequestSupportedPaymentTypes = "bacs"
	RoutingNumberLookupRequestSupportedPaymentTypesBook        RoutingNumberLookupRequestSupportedPaymentTypes = "book"
	RoutingNumberLookupRequestSupportedPaymentTypesCard        RoutingNumberLookupRequestSupportedPaymentTypes = "card"
	RoutingNumberLookupRequestSupportedPaymentTypesCheck       RoutingNumberLookupRequestSupportedPaymentTypes = "check"
	RoutingNumberLookupRequestSupportedPaymentTypesCrossBorder RoutingNumberLookupRequestSupportedPaymentTypes = "cross_border"
	RoutingNumberLookupRequestSupportedPaymentTypesEft         RoutingNumberLookupRequestSupportedPaymentTypes = "eft"
	RoutingNumberLookupRequestSupportedPaymentTypesInterac     RoutingNumberLookupRequestSupportedPaymentTypes = "interac"
	RoutingNumberLookupRequestSupportedPaymentTypesMasav       RoutingNumberLookupRequestSupportedPaymentTypes = "masav"
	RoutingNumberLookupRequestSupportedPaymentTypesNeft        RoutingNumberLookupRequestSupportedPaymentTypes = "neft"
	RoutingNumberLookupRequestSupportedPaymentTypesProvxchange RoutingNumberLookupRequestSupportedPaymentTypes = "provxchange"
	RoutingNumberLookupRequestSupportedPaymentTypesRtp         RoutingNumberLookupRequestSupportedPaymentTypes = "rtp"
	RoutingNumberLookupRequestSupportedPaymentTypesSen         RoutingNumberLookupRequestSupportedPaymentTypes = "sen"
	RoutingNumberLookupRequestSupportedPaymentTypesSepa        RoutingNumberLookupRequestSupportedPaymentTypes = "sepa"
	RoutingNumberLookupRequestSupportedPaymentTypesSignet      RoutingNumberLookupRequestSupportedPaymentTypes = "signet"
	RoutingNumberLookupRequestSupportedPaymentTypesWire        RoutingNumberLookupRequestSupportedPaymentTypes = "wire"
)

type RoutingNumberLookupRequestBankAddress struct {
	Line1 string `json:"line1,nullable"`
	Line2 string `json:"line2,nullable"`
	// Locality or City.
	Locality string `json:"locality,nullable"`
	// Region or State.
	Region string `json:"region,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,nullable"`
	JSON    RoutingNumberLookupRequestBankAddressJSON
}

type RoutingNumberLookupRequestBankAddressJSON struct {
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	Locality   pjson.Metadata
	Region     pjson.Metadata
	PostalCode pjson.Metadata
	Country    pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RoutingNumberLookupRequestBankAddress using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RoutingNumberLookupRequestBankAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
