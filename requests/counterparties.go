package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type CounterpartyNewParams struct {
	// A human friendly name for this counterparty.
	Name field.Field[string] `json:"name,required,nullable"`
	// The accounts for this counterparty.
	Accounts field.Field[[]CounterpartyNewParamsAccounts] `json:"accounts"`
	// The counterparty's email.
	Email field.Field[string] `json:"email,nullable" format:"email"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// Send an email to the counterparty whenever an associated payment order is sent
	// to the bank.
	SendRemittanceAdvice field.Field[bool]                            `json:"send_remittance_advice"`
	Accounting           field.Field[CounterpartyNewParamsAccounting] `json:"accounting"`
	// An optional type to auto-sync the counterparty to your ledger. Either `customer`
	// or `vendor`.
	LedgerType field.Field[CounterpartyNewParamsLedgerType] `json:"ledger_type"`
	// Either a valid SSN or EIN.
	TaxpayerIdentifier field.Field[string] `json:"taxpayer_identifier"`
}

// MarshalJSON serializes CounterpartyNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CounterpartyNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CounterpartyNewParamsAccounts struct {
	// Can be `checking`, `savings` or `other`.
	AccountType field.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType field.Field[CounterpartyNewParamsAccountsPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress field.Field[CounterpartyNewParamsAccountsPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           field.Field[string]                                        `json:"name,nullable"`
	AccountDetails field.Field[[]CounterpartyNewParamsAccountsAccountDetails] `json:"account_details"`
	RoutingDetails field.Field[[]CounterpartyNewParamsAccountsRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       field.Field[string] `json:"party_name"`
	PartyIdentifier field.Field[string] `json:"party_identifier"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken field.Field[string]                                        `json:"plaid_processor_token"`
	ContactDetails      field.Field[[]CounterpartyNewParamsAccountsContactDetails] `json:"contact_details"`
}

type CounterpartyNewParamsAccountsPartyType string

const (
	CounterpartyNewParamsAccountsPartyTypeBusiness   CounterpartyNewParamsAccountsPartyType = "business"
	CounterpartyNewParamsAccountsPartyTypeIndividual CounterpartyNewParamsAccountsPartyType = "individual"
)

type CounterpartyNewParamsAccountsPartyAddress struct {
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

type CounterpartyNewParamsAccountsAccountDetails struct {
	AccountNumber     field.Field[string]                                                       `json:"account_number,required"`
	AccountNumberType field.Field[CounterpartyNewParamsAccountsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type CounterpartyNewParamsAccountsAccountDetailsAccountNumberType string

const (
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban          CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "iban"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeClabe         CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "clabe"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeWalletAddress CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "wallet_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePan           CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "pan"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeOther         CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "other"
)

type CounterpartyNewParamsAccountsRoutingDetails struct {
	RoutingNumber     field.Field[string]                                                       `json:"routing_number,required"`
	RoutingNumberType field.Field[CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       field.Field[CounterpartyNewParamsAccountsRoutingDetailsPaymentType]       `json:"payment_type"`
}

type CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba          CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "aba"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAuBsb        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "au_bsb"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeBrCodigo     CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "br_codigo"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCaCpa        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "ca_cpa"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeChips        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "chips"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCnaps        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "cnaps"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeGBSortCode   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "gb_sort_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeInIfsc       CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "in_ifsc"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMyBranchCode CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "my_branch_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSwift        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "swift"
)

type CounterpartyNewParamsAccountsRoutingDetailsPaymentType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ach"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeAuBecs      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "au_becs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBacs        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "bacs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBook        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "book"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCard        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "card"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCheck       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "check"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEft         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "eft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCrossBorder CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "cross_border"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeInterac     CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "interac"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMasav       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "masav"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNeft        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "neft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeProvxchange CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "provxchange"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRtp         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "rtp"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSen         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sen"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSepa        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sepa"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSignet      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "signet"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeWire        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "wire"
)

type CounterpartyNewParamsAccountsContactDetails struct {
	ContactIdentifier     field.Field[string]                                                           `json:"contact_identifier"`
	ContactIdentifierType field.Field[CounterpartyNewParamsAccountsContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type CounterpartyNewParamsAccountsContactDetailsContactIdentifierType string

const (
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail       CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "email"
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypePhoneNumber CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "phone_number"
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeWebsite     CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "website"
)

type CounterpartyNewParamsAccounting struct {
	// An optional type to auto-sync the counterparty to your ledger. Either `customer`
	// or `vendor`.
	Type field.Field[CounterpartyNewParamsAccountingType] `json:"type"`
}

type CounterpartyNewParamsAccountingType string

const (
	CounterpartyNewParamsAccountingTypeCustomer CounterpartyNewParamsAccountingType = "customer"
	CounterpartyNewParamsAccountingTypeVendor   CounterpartyNewParamsAccountingType = "vendor"
)

type CounterpartyNewParamsLedgerType string

const (
	CounterpartyNewParamsLedgerTypeCustomer CounterpartyNewParamsLedgerType = "customer"
	CounterpartyNewParamsLedgerTypeVendor   CounterpartyNewParamsLedgerType = "vendor"
)

type CounterpartyUpdateParams struct {
	// A new name for the counterparty. Will only update if passed.
	Name field.Field[string] `json:"name"`
	// A new email for the counterparty.
	Email field.Field[string] `json:"email" format:"email"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this is `true`, Modern Treasury will send an email to the counterparty
	// whenever an associated payment order is sent to the bank.
	SendRemittanceAdvice field.Field[bool] `json:"send_remittance_advice"`
	// Either a valid SSN or EIN.
	TaxpayerIdentifier field.Field[string] `json:"taxpayer_identifier"`
}

// MarshalJSON serializes CounterpartyUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r CounterpartyUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CounterpartyListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Performs a partial string match of the name field. This is also case
	// insensitive.
	Name field.Field[string] `query:"name"`
	// Performs a partial string match of the email field. This is also case
	// insensitive.
	Email field.Field[string] `query:"email" format:"email"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	// Used to return counterparties created after some datetime.
	CreatedAtLowerBound field.Field[time.Time] `query:"created_at_lower_bound" format:"date-time"`
	// Used to return counterparties created before some datetime.
	CreatedAtUpperBound field.Field[time.Time] `query:"created_at_upper_bound" format:"date-time"`
}

// URLQuery serializes CounterpartyListParams into a url.Values of the query
// parameters associated with this value
func (r CounterpartyListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type CounterpartyCollectAccountParams struct {
	// One of `credit` or `debit`. Use `credit` when you want to pay a counterparty.
	// Use `debit` when you need to charge a counterparty. This field helps us send a
	// more tailored email to your counterparties."
	Direction field.Field[CounterpartyCollectAccountParamsDirection] `json:"direction,required"`
	// By default, Modern Treasury will send an email to your counterparty that
	// includes a link to the form they must fill out. However, if you would like to
	// send the counterparty the link, you can set this parameter to `false`. The JSON
	// body will include the link to the secure Modern Treasury form.
	SendEmail field.Field[bool] `json:"send_email"`
	// The list of fields you want on the form. This field is optional and if it is not
	// set, will default to [\"nameOnAccount\", \"accountType\", \"accountNumber\",
	// \"routingNumber\", \"address\"]. The full list of options is [\"name\",
	// \"nameOnAccount\", \"taxpayerIdentifier\", \"accountType\", \"accountNumber\",
	// \"routingNumber\", \"address\", \"ibanNumber\", \"swiftCode\"].
	Fields field.Field[[]CounterpartyCollectAccountParamsFields] `json:"fields"`
	// The URL you want your customer to visit upon filling out the form. By default,
	// they will be sent to a Modern Treasury landing page. This must be a valid HTTPS
	// URL if set.
	CustomRedirect field.Field[string] `json:"custom_redirect" format:"uri"`
}

// MarshalJSON serializes CounterpartyCollectAccountParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r CounterpartyCollectAccountParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type CounterpartyCollectAccountParamsDirection string

const (
	CounterpartyCollectAccountParamsDirectionCredit CounterpartyCollectAccountParamsDirection = "credit"
	CounterpartyCollectAccountParamsDirectionDebit  CounterpartyCollectAccountParamsDirection = "debit"
)

type CounterpartyCollectAccountParamsFields string

const (
	CounterpartyCollectAccountParamsFieldsName                 CounterpartyCollectAccountParamsFields = "name"
	CounterpartyCollectAccountParamsFieldsNameOnAccount        CounterpartyCollectAccountParamsFields = "nameOnAccount"
	CounterpartyCollectAccountParamsFieldsTaxpayerIdentifier   CounterpartyCollectAccountParamsFields = "taxpayerIdentifier"
	CounterpartyCollectAccountParamsFieldsAccountType          CounterpartyCollectAccountParamsFields = "accountType"
	CounterpartyCollectAccountParamsFieldsAccountNumber        CounterpartyCollectAccountParamsFields = "accountNumber"
	CounterpartyCollectAccountParamsFieldsIbanNumber           CounterpartyCollectAccountParamsFields = "ibanNumber"
	CounterpartyCollectAccountParamsFieldsClabeNumber          CounterpartyCollectAccountParamsFields = "clabeNumber"
	CounterpartyCollectAccountParamsFieldsWalletAddress        CounterpartyCollectAccountParamsFields = "walletAddress"
	CounterpartyCollectAccountParamsFieldsPanNumber            CounterpartyCollectAccountParamsFields = "panNumber"
	CounterpartyCollectAccountParamsFieldsRoutingNumber        CounterpartyCollectAccountParamsFields = "routingNumber"
	CounterpartyCollectAccountParamsFieldsAbaWireRoutingNumber CounterpartyCollectAccountParamsFields = "abaWireRoutingNumber"
	CounterpartyCollectAccountParamsFieldsSwiftCode            CounterpartyCollectAccountParamsFields = "swiftCode"
	CounterpartyCollectAccountParamsFieldsAuBsb                CounterpartyCollectAccountParamsFields = "auBsb"
	CounterpartyCollectAccountParamsFieldsCaCpa                CounterpartyCollectAccountParamsFields = "caCpa"
	CounterpartyCollectAccountParamsFieldsCnaps                CounterpartyCollectAccountParamsFields = "cnaps"
	CounterpartyCollectAccountParamsFieldsGBSortCode           CounterpartyCollectAccountParamsFields = "gbSortCode"
	CounterpartyCollectAccountParamsFieldsInIfsc               CounterpartyCollectAccountParamsFields = "inIfsc"
	CounterpartyCollectAccountParamsFieldsMyBranchCode         CounterpartyCollectAccountParamsFields = "myBranchCode"
	CounterpartyCollectAccountParamsFieldsBrCodigo             CounterpartyCollectAccountParamsFields = "brCodigo"
	CounterpartyCollectAccountParamsFieldsRoutingNumberType    CounterpartyCollectAccountParamsFields = "routingNumberType"
	CounterpartyCollectAccountParamsFieldsAddress              CounterpartyCollectAccountParamsFields = "address"
)
