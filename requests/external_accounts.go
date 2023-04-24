package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type ExternalAccountType string

const (
	ExternalAccountTypeCash        ExternalAccountType = "cash"
	ExternalAccountTypeChecking    ExternalAccountType = "checking"
	ExternalAccountTypeLoan        ExternalAccountType = "loan"
	ExternalAccountTypeNonResident ExternalAccountType = "non_resident"
	ExternalAccountTypeOther       ExternalAccountType = "other"
	ExternalAccountTypeOverdraft   ExternalAccountType = "overdraft"
	ExternalAccountTypeSavings     ExternalAccountType = "savings"
)

type ExternalAccountNewParams struct {
	// Can be `checking`, `savings` or `other`.
	AccountType field.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType field.Field[ExternalAccountNewParamsPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress field.Field[ExternalAccountNewParamsPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           field.Field[string]                                   `json:"name,nullable"`
	CounterpartyID field.Field[string]                                   `json:"counterparty_id,required,nullable" format:"uuid"`
	AccountDetails field.Field[[]ExternalAccountNewParamsAccountDetails] `json:"account_details"`
	RoutingDetails field.Field[[]ExternalAccountNewParamsRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       field.Field[string] `json:"party_name"`
	PartyIdentifier field.Field[string] `json:"party_identifier"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken field.Field[string]                                   `json:"plaid_processor_token"`
	ContactDetails      field.Field[[]ExternalAccountNewParamsContactDetails] `json:"contact_details"`
}

// MarshalJSON serializes ExternalAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ExternalAccountNewParamsPartyType string

const (
	ExternalAccountNewParamsPartyTypeBusiness   ExternalAccountNewParamsPartyType = "business"
	ExternalAccountNewParamsPartyTypeIndividual ExternalAccountNewParamsPartyType = "individual"
)

type ExternalAccountNewParamsPartyAddress struct {
	Line1 field.Field[string] `json:"line1,nullable"`
	Line2 field.Field[string] `json:"line2,nullable"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,nullable"`
	// Region or State.
	Region field.Field[string] `json:"region,nullable"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,nullable"`
}

type ExternalAccountNewParamsAccountDetails struct {
	AccountNumber     field.Field[string]                                                  `json:"account_number,required"`
	AccountNumberType field.Field[ExternalAccountNewParamsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type ExternalAccountNewParamsAccountDetailsAccountNumberType string

const (
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban          ExternalAccountNewParamsAccountDetailsAccountNumberType = "iban"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeClabe         ExternalAccountNewParamsAccountDetailsAccountNumberType = "clabe"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeWalletAddress ExternalAccountNewParamsAccountDetailsAccountNumberType = "wallet_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypePan           ExternalAccountNewParamsAccountDetailsAccountNumberType = "pan"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeOther         ExternalAccountNewParamsAccountDetailsAccountNumberType = "other"
)

type ExternalAccountNewParamsRoutingDetails struct {
	RoutingNumber     field.Field[string]                                                  `json:"routing_number,required"`
	RoutingNumberType field.Field[ExternalAccountNewParamsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       field.Field[ExternalAccountNewParamsRoutingDetailsPaymentType]       `json:"payment_type"`
}

type ExternalAccountNewParamsRoutingDetailsRoutingNumberType string

const (
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba          ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "aba"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAuBsb        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "au_bsb"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeBrCodigo     ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "br_codigo"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCaCpa        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "ca_cpa"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeChips        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "chips"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCnaps        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "cnaps"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeGBSortCode   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "gb_sort_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeInIfsc       ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "in_ifsc"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMyBranchCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "my_branch_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSwift        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "swift"
)

type ExternalAccountNewParamsRoutingDetailsPaymentType string

const (
	ExternalAccountNewParamsRoutingDetailsPaymentTypeACH         ExternalAccountNewParamsRoutingDetailsPaymentType = "ach"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeAuBecs      ExternalAccountNewParamsRoutingDetailsPaymentType = "au_becs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBacs        ExternalAccountNewParamsRoutingDetailsPaymentType = "bacs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBook        ExternalAccountNewParamsRoutingDetailsPaymentType = "book"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCard        ExternalAccountNewParamsRoutingDetailsPaymentType = "card"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCheck       ExternalAccountNewParamsRoutingDetailsPaymentType = "check"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeEft         ExternalAccountNewParamsRoutingDetailsPaymentType = "eft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCrossBorder ExternalAccountNewParamsRoutingDetailsPaymentType = "cross_border"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeInterac     ExternalAccountNewParamsRoutingDetailsPaymentType = "interac"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeMasav       ExternalAccountNewParamsRoutingDetailsPaymentType = "masav"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeNeft        ExternalAccountNewParamsRoutingDetailsPaymentType = "neft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeProvxchange ExternalAccountNewParamsRoutingDetailsPaymentType = "provxchange"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeRtp         ExternalAccountNewParamsRoutingDetailsPaymentType = "rtp"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSen         ExternalAccountNewParamsRoutingDetailsPaymentType = "sen"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSepa        ExternalAccountNewParamsRoutingDetailsPaymentType = "sepa"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSignet      ExternalAccountNewParamsRoutingDetailsPaymentType = "signet"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeWire        ExternalAccountNewParamsRoutingDetailsPaymentType = "wire"
)

type ExternalAccountNewParamsContactDetails struct {
	ContactIdentifier     field.Field[string]                                                      `json:"contact_identifier"`
	ContactIdentifierType field.Field[ExternalAccountNewParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type ExternalAccountNewParamsContactDetailsContactIdentifierType string

const (
	ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail       ExternalAccountNewParamsContactDetailsContactIdentifierType = "email"
	ExternalAccountNewParamsContactDetailsContactIdentifierTypePhoneNumber ExternalAccountNewParamsContactDetailsContactIdentifierType = "phone_number"
	ExternalAccountNewParamsContactDetailsContactIdentifierTypeWebsite     ExternalAccountNewParamsContactDetailsContactIdentifierType = "website"
)

type ExternalAccountUpdateParams struct {
	// Either `individual` or `business`.
	PartyType field.Field[ExternalAccountUpdateParamsPartyType] `json:"party_type,nullable"`
	// Can be `checking`, `savings` or `other`.
	AccountType    field.Field[ExternalAccountType] `json:"account_type"`
	CounterpartyID field.Field[string]              `json:"counterparty_id,nullable" format:"uuid"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name field.Field[string] `json:"name,nullable"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName    field.Field[string]                                  `json:"party_name"`
	PartyAddress field.Field[ExternalAccountUpdateParamsPartyAddress] `json:"party_address"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes ExternalAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ExternalAccountUpdateParamsPartyType string

const (
	ExternalAccountUpdateParamsPartyTypeBusiness   ExternalAccountUpdateParamsPartyType = "business"
	ExternalAccountUpdateParamsPartyTypeIndividual ExternalAccountUpdateParamsPartyType = "individual"
)

type ExternalAccountUpdateParamsPartyAddress struct {
	Line1 field.Field[string] `json:"line1,nullable"`
	Line2 field.Field[string] `json:"line2,nullable"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,nullable"`
	// Region or State.
	Region field.Field[string] `json:"region,nullable"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,nullable"`
}

type ExternalAccountListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Searches the ExternalAccount's party_name AND the Counterparty's party_name
	PartyName      field.Field[string] `query:"party_name"`
	CounterpartyID field.Field[string] `query:"counterparty_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes ExternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r ExternalAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type ExternalAccountCompleteVerificationParams struct {
	Amounts field.Field[[]int64] `json:"amounts"`
}

// MarshalJSON serializes ExternalAccountCompleteVerificationParams into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r ExternalAccountCompleteVerificationParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ExternalAccountVerifyParams struct {
	// The ID of the internal account where the micro-deposits originate from. Both
	// credit and debit capabilities must be enabled.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Both ach and eft are supported payment types.
	PaymentType field.Field[ExternalAccountVerifyParamsPaymentType] `json:"payment_type,required"`
	// Defaults to the currency of the originating account.
	Currency field.Field[Currency] `json:"currency,nullable"`
}

// MarshalJSON serializes ExternalAccountVerifyParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountVerifyParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ExternalAccountVerifyParamsPaymentType string

const (
	ExternalAccountVerifyParamsPaymentTypeACH         ExternalAccountVerifyParamsPaymentType = "ach"
	ExternalAccountVerifyParamsPaymentTypeAuBecs      ExternalAccountVerifyParamsPaymentType = "au_becs"
	ExternalAccountVerifyParamsPaymentTypeBacs        ExternalAccountVerifyParamsPaymentType = "bacs"
	ExternalAccountVerifyParamsPaymentTypeBook        ExternalAccountVerifyParamsPaymentType = "book"
	ExternalAccountVerifyParamsPaymentTypeCard        ExternalAccountVerifyParamsPaymentType = "card"
	ExternalAccountVerifyParamsPaymentTypeCheck       ExternalAccountVerifyParamsPaymentType = "check"
	ExternalAccountVerifyParamsPaymentTypeCrossBorder ExternalAccountVerifyParamsPaymentType = "cross_border"
	ExternalAccountVerifyParamsPaymentTypeEft         ExternalAccountVerifyParamsPaymentType = "eft"
	ExternalAccountVerifyParamsPaymentTypeInterac     ExternalAccountVerifyParamsPaymentType = "interac"
	ExternalAccountVerifyParamsPaymentTypeMasav       ExternalAccountVerifyParamsPaymentType = "masav"
	ExternalAccountVerifyParamsPaymentTypeNeft        ExternalAccountVerifyParamsPaymentType = "neft"
	ExternalAccountVerifyParamsPaymentTypeProvxchange ExternalAccountVerifyParamsPaymentType = "provxchange"
	ExternalAccountVerifyParamsPaymentTypeRtp         ExternalAccountVerifyParamsPaymentType = "rtp"
	ExternalAccountVerifyParamsPaymentTypeSen         ExternalAccountVerifyParamsPaymentType = "sen"
	ExternalAccountVerifyParamsPaymentTypeSepa        ExternalAccountVerifyParamsPaymentType = "sepa"
	ExternalAccountVerifyParamsPaymentTypeSignet      ExternalAccountVerifyParamsPaymentType = "signet"
	ExternalAccountVerifyParamsPaymentTypeWire        ExternalAccountVerifyParamsPaymentType = "wire"
)
