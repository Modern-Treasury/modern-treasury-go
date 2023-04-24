package requests

import (
	"io"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type PaymentOrderType string

const (
	PaymentOrderTypeACH         PaymentOrderType = "ach"
	PaymentOrderTypeAuBecs      PaymentOrderType = "au_becs"
	PaymentOrderTypeBacs        PaymentOrderType = "bacs"
	PaymentOrderTypeBook        PaymentOrderType = "book"
	PaymentOrderTypeCard        PaymentOrderType = "card"
	PaymentOrderTypeCheck       PaymentOrderType = "check"
	PaymentOrderTypeCrossBorder PaymentOrderType = "cross_border"
	PaymentOrderTypeEft         PaymentOrderType = "eft"
	PaymentOrderTypeInterac     PaymentOrderType = "interac"
	PaymentOrderTypeMasav       PaymentOrderType = "masav"
	PaymentOrderTypeNeft        PaymentOrderType = "neft"
	PaymentOrderTypeProvxchange PaymentOrderType = "provxchange"
	PaymentOrderTypeRtp         PaymentOrderType = "rtp"
	PaymentOrderTypeSen         PaymentOrderType = "sen"
	PaymentOrderTypeSepa        PaymentOrderType = "sepa"
	PaymentOrderTypeSignet      PaymentOrderType = "signet"
	PaymentOrderTypeWire        PaymentOrderType = "wire"
)

type PaymentOrderSubtype string

const (
	PaymentOrderSubtypeCcd PaymentOrderSubtype = "CCD"
	PaymentOrderSubtypeCie PaymentOrderSubtype = "CIE"
	PaymentOrderSubtypeCtx PaymentOrderSubtype = "CTX"
	PaymentOrderSubtypeIat PaymentOrderSubtype = "IAT"
	PaymentOrderSubtypePpd PaymentOrderSubtype = "PPD"
	PaymentOrderSubtypeTel PaymentOrderSubtype = "TEL"
	PaymentOrderSubtypeWeb PaymentOrderSubtype = "WEB"
)

type PaymentOrderNewParams struct {
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type field.Field[PaymentOrderType] `json:"type,required"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype field.Field[PaymentOrderSubtype] `json:"subtype,nullable"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[PaymentOrderNewParamsDirection] `json:"direction,required"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority field.Field[PaymentOrderNewParamsPriority] `json:"priority"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID field.Field[string]                          `json:"receiving_account_id" format:"uuid"`
	Accounting         field.Field[PaymentOrderNewParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID field.Field[string] `json:"accounting_category_id,nullable" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID field.Field[string] `json:"accounting_ledger_class_id,nullable" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate field.Field[time.Time] `json:"effective_date" format:"date"`
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor field.Field[string] `json:"statement_descriptor,nullable"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation field.Field[string] `json:"remittance_information,nullable"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose field.Field[string] `json:"purpose,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer field.Field[PaymentOrderNewParamsChargeBearer] `json:"charge_bearer,nullable"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator field.Field[PaymentOrderNewParamsForeignExchangeIndicator] `json:"foreign_exchange_indicator,nullable"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract field.Field[string] `json:"foreign_exchange_contract,nullable"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected field.Field[bool] `json:"nsf_protected"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName field.Field[string] `json:"originating_party_name,nullable"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName field.Field[string] `json:"ultimate_originating_party_name,nullable"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier field.Field[string] `json:"ultimate_originating_party_identifier,nullable"`
	// Name of the ultimate funds recipient.
	UltimateReceivingPartyName field.Field[string] `json:"ultimate_receiving_party_name,nullable"`
	// Identifier of the ultimate funds recipient.
	UltimateReceivingPartyIdentifier field.Field[string] `json:"ultimate_receiving_party_identifier,nullable"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice field.Field[bool] `json:"send_remittance_advice,nullable"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt field.Field[time.Time] `json:"expires_at,nullable" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType field.Field[PaymentOrderNewParamsFallbackType] `json:"fallback_type"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount field.Field[PaymentOrderNewParamsReceivingAccount] `json:"receiving_account"`
	// Specifies a ledger transaction object that will be created with the payment
	// order. If the ledger transaction cannot be created, then the payment order
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the payment order.
	LedgerTransaction field.Field[PaymentOrderNewParamsLedgerTransaction] `json:"ledger_transaction"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems field.Field[[]PaymentOrderNewParamsLineItems] `json:"line_items"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	TransactionMonitoringEnabled field.Field[bool] `json:"transaction_monitoring_enabled"`
	// An array of documents to be attached to the payment order. Note that if you
	// attach documents, the request's content type must be `multipart/form-data`.
	Documents field.Field[[]PaymentOrderNewParamsDocuments] `json:"documents"`
}

// MarshalJSON serializes PaymentOrderNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r PaymentOrderNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type PaymentOrderNewParamsDirection string

const (
	PaymentOrderNewParamsDirectionCredit PaymentOrderNewParamsDirection = "credit"
	PaymentOrderNewParamsDirectionDebit  PaymentOrderNewParamsDirection = "debit"
)

type PaymentOrderNewParamsPriority string

const (
	PaymentOrderNewParamsPriorityHigh   PaymentOrderNewParamsPriority = "high"
	PaymentOrderNewParamsPriorityNormal PaymentOrderNewParamsPriority = "normal"
)

type PaymentOrderNewParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID field.Field[string] `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID field.Field[string] `json:"class_id,nullable" format:"uuid"`
}

type PaymentOrderNewParamsChargeBearer string

const (
	PaymentOrderNewParamsChargeBearerShared   PaymentOrderNewParamsChargeBearer = "shared"
	PaymentOrderNewParamsChargeBearerSender   PaymentOrderNewParamsChargeBearer = "sender"
	PaymentOrderNewParamsChargeBearerReceiver PaymentOrderNewParamsChargeBearer = "receiver"
)

type PaymentOrderNewParamsForeignExchangeIndicator string

const (
	PaymentOrderNewParamsForeignExchangeIndicatorFixedToVariable PaymentOrderNewParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderNewParamsForeignExchangeIndicatorVariableToFixed PaymentOrderNewParamsForeignExchangeIndicator = "variable_to_fixed"
)

type PaymentOrderNewParamsFallbackType string

const (
	PaymentOrderNewParamsFallbackTypeACH PaymentOrderNewParamsFallbackType = "ach"
)

type PaymentOrderNewParamsReceivingAccount struct {
	// Can be `checking`, `savings` or `other`.
	AccountType field.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType field.Field[PaymentOrderNewParamsReceivingAccountPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress field.Field[PaymentOrderNewParamsReceivingAccountPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           field.Field[string]                                                `json:"name,nullable"`
	AccountDetails field.Field[[]PaymentOrderNewParamsReceivingAccountAccountDetails] `json:"account_details"`
	RoutingDetails field.Field[[]PaymentOrderNewParamsReceivingAccountRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       field.Field[string] `json:"party_name"`
	PartyIdentifier field.Field[string] `json:"party_identifier"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken field.Field[string]                                                `json:"plaid_processor_token"`
	ContactDetails      field.Field[[]PaymentOrderNewParamsReceivingAccountContactDetails] `json:"contact_details"`
}

type PaymentOrderNewParamsReceivingAccountPartyType string

const (
	PaymentOrderNewParamsReceivingAccountPartyTypeBusiness   PaymentOrderNewParamsReceivingAccountPartyType = "business"
	PaymentOrderNewParamsReceivingAccountPartyTypeIndividual PaymentOrderNewParamsReceivingAccountPartyType = "individual"
)

type PaymentOrderNewParamsReceivingAccountPartyAddress struct {
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

type PaymentOrderNewParamsReceivingAccountAccountDetails struct {
	AccountNumber     field.Field[string]                                                               `json:"account_number,required"`
	AccountNumberType field.Field[PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIban          PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeClabe         PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypePan           PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeOther         PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "other"
)

type PaymentOrderNewParamsReceivingAccountRoutingDetails struct {
	RoutingNumber     field.Field[string]                                                               `json:"routing_number,required"`
	RoutingNumberType field.Field[PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       field.Field[PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

type PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba          PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb        PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo     PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa        PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips        PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps        PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode   PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc       PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift        PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "swift"
)

type PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "wire"
)

type PaymentOrderNewParamsReceivingAccountContactDetails struct {
	ContactIdentifier     field.Field[string]                                                                   `json:"contact_identifier"`
	ContactIdentifierType field.Field[PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType string

const (
	PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeEmail       PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType = "email"
	PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypePhoneNumber PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeWebsite     PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType = "website"
)

type PaymentOrderNewParamsLedgerTransaction struct {
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status field.Field[PaymentOrderNewParamsLedgerTransactionStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate field.Field[time.Time] `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries field.Field[[]PaymentOrderNewParamsLedgerTransactionLedgerEntries] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID field.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType field.Field[PaymentOrderNewParamsLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID field.Field[string] `json:"ledgerable_id" format:"uuid"`
}

type PaymentOrderNewParamsLedgerTransactionStatus string

const (
	PaymentOrderNewParamsLedgerTransactionStatusArchived PaymentOrderNewParamsLedgerTransactionStatus = "archived"
	PaymentOrderNewParamsLedgerTransactionStatusPending  PaymentOrderNewParamsLedgerTransactionStatus = "pending"
	PaymentOrderNewParamsLedgerTransactionStatusPosted   PaymentOrderNewParamsLedgerTransactionStatus = "posted"
)

type PaymentOrderNewParamsLedgerTransactionLedgerEntries struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID field.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion field.Field[int64] `json:"lock_version,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount field.Field[map[string]int64] `json:"pending_balance_amount,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount field.Field[map[string]int64] `json:"posted_balance_amount,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount field.Field[map[string]int64] `json:"available_balance_amount,nullable"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances field.Field[bool] `json:"show_resulting_ledger_account_balances,nullable"`
}

type PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirection string

const (
	PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirectionCredit PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirection = "credit"
	PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirectionDebit  PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirection = "debit"
)

type PaymentOrderNewParamsLedgerTransactionLedgerableType string

const (
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeCounterparty          PaymentOrderNewParamsLedgerTransactionLedgerableType = "counterparty"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeExpectedPayment       PaymentOrderNewParamsLedgerTransactionLedgerableType = "expected_payment"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeIncomingPaymentDetail PaymentOrderNewParamsLedgerTransactionLedgerableType = "incoming_payment_detail"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeInternalAccount       PaymentOrderNewParamsLedgerTransactionLedgerableType = "internal_account"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeLineItem              PaymentOrderNewParamsLedgerTransactionLedgerableType = "line_item"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypePaperItem             PaymentOrderNewParamsLedgerTransactionLedgerableType = "paper_item"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypePaymentOrder          PaymentOrderNewParamsLedgerTransactionLedgerableType = "payment_order"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypePaymentOrderAttempt   PaymentOrderNewParamsLedgerTransactionLedgerableType = "payment_order_attempt"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeReturn                PaymentOrderNewParamsLedgerTransactionLedgerableType = "return"
	PaymentOrderNewParamsLedgerTransactionLedgerableTypeReversal              PaymentOrderNewParamsLedgerTransactionLedgerableType = "reversal"
)

type PaymentOrderNewParamsLineItems struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount field.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description field.Field[string] `json:"description,nullable"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID field.Field[string] `json:"accounting_category_id,nullable"`
}

type PaymentOrderNewParamsDocuments struct {
	// A category given to the document, can be `null`.
	DocumentType field.Field[string]    `json:"document_type"`
	File         field.Field[io.Reader] `json:"file,required" format:"binary"`
}

type PaymentOrderUpdateParams struct {
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type field.Field[PaymentOrderType] `json:"type"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype field.Field[PaymentOrderSubtype] `json:"subtype,nullable"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount field.Field[int64] `json:"amount"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[PaymentOrderUpdateParamsDirection] `json:"direction"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority field.Field[PaymentOrderUpdateParamsPriority] `json:"priority"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID field.Field[string] `json:"originating_account_id" format:"uuid"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID field.Field[string]                             `json:"receiving_account_id" format:"uuid"`
	Accounting         field.Field[PaymentOrderUpdateParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID field.Field[string] `json:"accounting_category_id,nullable" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID field.Field[string] `json:"accounting_ledger_class_id,nullable" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate field.Field[time.Time] `json:"effective_date" format:"date"`
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor field.Field[string] `json:"statement_descriptor,nullable"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation field.Field[string] `json:"remittance_information,nullable"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose field.Field[string] `json:"purpose,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer field.Field[PaymentOrderUpdateParamsChargeBearer] `json:"charge_bearer,nullable"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator field.Field[PaymentOrderUpdateParamsForeignExchangeIndicator] `json:"foreign_exchange_indicator,nullable"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract field.Field[string] `json:"foreign_exchange_contract,nullable"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected field.Field[bool] `json:"nsf_protected"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName field.Field[string] `json:"originating_party_name,nullable"`
	// This represents the name of the person that the payment is on behalf of when
	// using the CIE subtype for ACH payments. Only the first 15 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateOriginatingPartyName field.Field[string] `json:"ultimate_originating_party_name,nullable"`
	// This represents the identifier by which the person is known to the receiver when
	// using the CIE subtype for ACH payments. Only the first 22 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateOriginatingPartyIdentifier field.Field[string] `json:"ultimate_originating_party_identifier,nullable"`
	// This represents the identifier by which the merchant is known to the person
	// initiating an ACH payment with CIE subtype. Only the first 15 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateReceivingPartyName field.Field[string] `json:"ultimate_receiving_party_name,nullable"`
	// This represents the name of the merchant that the payment is being sent to when
	// using the CIE subtype for ACH payments. Only the first 22 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateReceivingPartyIdentifier field.Field[string] `json:"ultimate_receiving_party_identifier,nullable"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice field.Field[bool] `json:"send_remittance_advice,nullable"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt field.Field[time.Time] `json:"expires_at,nullable" format:"date-time"`
	// To cancel a payment order, use `cancelled`. To redraft a returned payment order,
	// use `approved`. To undo approval on a denied or approved payment order, use
	// `needs_approval`.
	Status field.Field[PaymentOrderUpdateParamsStatus] `json:"status"`
	// Required when receiving_account_id is passed the ID of an external account.
	CounterpartyID field.Field[string] `json:"counterparty_id,nullable" format:"uuid"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType field.Field[PaymentOrderUpdateParamsFallbackType] `json:"fallback_type"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount field.Field[PaymentOrderUpdateParamsReceivingAccount] `json:"receiving_account"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems field.Field[[]PaymentOrderUpdateParamsLineItems] `json:"line_items"`
}

// MarshalJSON serializes PaymentOrderUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r PaymentOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsDirection string

const (
	PaymentOrderUpdateParamsDirectionCredit PaymentOrderUpdateParamsDirection = "credit"
	PaymentOrderUpdateParamsDirectionDebit  PaymentOrderUpdateParamsDirection = "debit"
)

type PaymentOrderUpdateParamsPriority string

const (
	PaymentOrderUpdateParamsPriorityHigh   PaymentOrderUpdateParamsPriority = "high"
	PaymentOrderUpdateParamsPriorityNormal PaymentOrderUpdateParamsPriority = "normal"
)

type PaymentOrderUpdateParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID field.Field[string] `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID field.Field[string] `json:"class_id,nullable" format:"uuid"`
}

type PaymentOrderUpdateParamsChargeBearer string

const (
	PaymentOrderUpdateParamsChargeBearerShared   PaymentOrderUpdateParamsChargeBearer = "shared"
	PaymentOrderUpdateParamsChargeBearerSender   PaymentOrderUpdateParamsChargeBearer = "sender"
	PaymentOrderUpdateParamsChargeBearerReceiver PaymentOrderUpdateParamsChargeBearer = "receiver"
)

type PaymentOrderUpdateParamsForeignExchangeIndicator string

const (
	PaymentOrderUpdateParamsForeignExchangeIndicatorFixedToVariable PaymentOrderUpdateParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderUpdateParamsForeignExchangeIndicatorVariableToFixed PaymentOrderUpdateParamsForeignExchangeIndicator = "variable_to_fixed"
)

type PaymentOrderUpdateParamsStatus string

const (
	PaymentOrderUpdateParamsStatusApproved      PaymentOrderUpdateParamsStatus = "approved"
	PaymentOrderUpdateParamsStatusCancelled     PaymentOrderUpdateParamsStatus = "cancelled"
	PaymentOrderUpdateParamsStatusCompleted     PaymentOrderUpdateParamsStatus = "completed"
	PaymentOrderUpdateParamsStatusDenied        PaymentOrderUpdateParamsStatus = "denied"
	PaymentOrderUpdateParamsStatusFailed        PaymentOrderUpdateParamsStatus = "failed"
	PaymentOrderUpdateParamsStatusNeedsApproval PaymentOrderUpdateParamsStatus = "needs_approval"
	PaymentOrderUpdateParamsStatusPending       PaymentOrderUpdateParamsStatus = "pending"
	PaymentOrderUpdateParamsStatusProcessing    PaymentOrderUpdateParamsStatus = "processing"
	PaymentOrderUpdateParamsStatusReturned      PaymentOrderUpdateParamsStatus = "returned"
	PaymentOrderUpdateParamsStatusReversed      PaymentOrderUpdateParamsStatus = "reversed"
	PaymentOrderUpdateParamsStatusSent          PaymentOrderUpdateParamsStatus = "sent"
)

type PaymentOrderUpdateParamsFallbackType string

const (
	PaymentOrderUpdateParamsFallbackTypeACH PaymentOrderUpdateParamsFallbackType = "ach"
)

type PaymentOrderUpdateParamsReceivingAccount struct {
	// Can be `checking`, `savings` or `other`.
	AccountType field.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType field.Field[PaymentOrderUpdateParamsReceivingAccountPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress field.Field[PaymentOrderUpdateParamsReceivingAccountPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           field.Field[string]                                                   `json:"name,nullable"`
	AccountDetails field.Field[[]PaymentOrderUpdateParamsReceivingAccountAccountDetails] `json:"account_details"`
	RoutingDetails field.Field[[]PaymentOrderUpdateParamsReceivingAccountRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       field.Field[string] `json:"party_name"`
	PartyIdentifier field.Field[string] `json:"party_identifier"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken field.Field[string]                                                   `json:"plaid_processor_token"`
	ContactDetails      field.Field[[]PaymentOrderUpdateParamsReceivingAccountContactDetails] `json:"contact_details"`
}

type PaymentOrderUpdateParamsReceivingAccountPartyType string

const (
	PaymentOrderUpdateParamsReceivingAccountPartyTypeBusiness   PaymentOrderUpdateParamsReceivingAccountPartyType = "business"
	PaymentOrderUpdateParamsReceivingAccountPartyTypeIndividual PaymentOrderUpdateParamsReceivingAccountPartyType = "individual"
)

type PaymentOrderUpdateParamsReceivingAccountPartyAddress struct {
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

type PaymentOrderUpdateParamsReceivingAccountAccountDetails struct {
	AccountNumber     field.Field[string]                                                                  `json:"account_number,required"`
	AccountNumberType field.Field[PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIban          PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeClabe         PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypePan           PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeOther         PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "other"
)

type PaymentOrderUpdateParamsReceivingAccountRoutingDetails struct {
	RoutingNumber     field.Field[string]                                                                  `json:"routing_number,required"`
	RoutingNumberType field.Field[PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       field.Field[PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

type PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba          PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo     PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode   PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "swift"
)

type PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "wire"
)

type PaymentOrderUpdateParamsReceivingAccountContactDetails struct {
	ContactIdentifier     field.Field[string]                                                                      `json:"contact_identifier"`
	ContactIdentifierType field.Field[PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType string

const (
	PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeEmail       PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType = "email"
	PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypePhoneNumber PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeWebsite     PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType = "website"
)

type PaymentOrderUpdateParamsLineItems struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount field.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description field.Field[string] `json:"description,nullable"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID field.Field[string] `json:"accounting_category_id,nullable"`
}

type PaymentOrderListParams struct {
	AfterCursor field.Field[string]                     `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]                      `query:"per_page"`
	Type        field.Field[PaymentOrderListParamsType] `query:"type"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority             field.Field[PaymentOrderListParamsPriority] `query:"priority"`
	CounterpartyID       field.Field[string]                         `query:"counterparty_id" format:"uuid"`
	OriginatingAccountID field.Field[string]                         `query:"originating_account_id" format:"uuid"`
	// The ID of a transaction that the payment order has been reconciled to.
	TransactionID field.Field[string]                          `query:"transaction_id" format:"uuid"`
	Status        field.Field[PaymentOrderListParamsStatus]    `query:"status"`
	Direction     field.Field[PaymentOrderListParamsDirection] `query:"direction"`
	// Query for records with the provided reference number
	ReferenceNumber field.Field[string] `query:"reference_number"`
	// An inclusive lower bound for searching effective_date
	EffectiveDateStart field.Field[time.Time] `query:"effective_date_start" format:"date"`
	// An inclusive upper bound for searching effective_date
	EffectiveDateEnd field.Field[time.Time] `query:"effective_date_end" format:"date"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes PaymentOrderListParams into a url.Values of the query
// parameters associated with this value
func (r PaymentOrderListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type PaymentOrderListParamsType string

const (
	PaymentOrderListParamsTypeACH         PaymentOrderListParamsType = "ach"
	PaymentOrderListParamsTypeAuBecs      PaymentOrderListParamsType = "au_becs"
	PaymentOrderListParamsTypeBacs        PaymentOrderListParamsType = "bacs"
	PaymentOrderListParamsTypeBook        PaymentOrderListParamsType = "book"
	PaymentOrderListParamsTypeCard        PaymentOrderListParamsType = "card"
	PaymentOrderListParamsTypeCheck       PaymentOrderListParamsType = "check"
	PaymentOrderListParamsTypeCrossBorder PaymentOrderListParamsType = "cross_border"
	PaymentOrderListParamsTypeEft         PaymentOrderListParamsType = "eft"
	PaymentOrderListParamsTypeInterac     PaymentOrderListParamsType = "interac"
	PaymentOrderListParamsTypeMasav       PaymentOrderListParamsType = "masav"
	PaymentOrderListParamsTypeNeft        PaymentOrderListParamsType = "neft"
	PaymentOrderListParamsTypeProvxchange PaymentOrderListParamsType = "provxchange"
	PaymentOrderListParamsTypeRtp         PaymentOrderListParamsType = "rtp"
	PaymentOrderListParamsTypeSen         PaymentOrderListParamsType = "sen"
	PaymentOrderListParamsTypeSepa        PaymentOrderListParamsType = "sepa"
	PaymentOrderListParamsTypeSignet      PaymentOrderListParamsType = "signet"
	PaymentOrderListParamsTypeWire        PaymentOrderListParamsType = "wire"
)

type PaymentOrderListParamsPriority string

const (
	PaymentOrderListParamsPriorityHigh   PaymentOrderListParamsPriority = "high"
	PaymentOrderListParamsPriorityNormal PaymentOrderListParamsPriority = "normal"
)

type PaymentOrderListParamsStatus string

const (
	PaymentOrderListParamsStatusApproved      PaymentOrderListParamsStatus = "approved"
	PaymentOrderListParamsStatusCancelled     PaymentOrderListParamsStatus = "cancelled"
	PaymentOrderListParamsStatusCompleted     PaymentOrderListParamsStatus = "completed"
	PaymentOrderListParamsStatusDenied        PaymentOrderListParamsStatus = "denied"
	PaymentOrderListParamsStatusFailed        PaymentOrderListParamsStatus = "failed"
	PaymentOrderListParamsStatusNeedsApproval PaymentOrderListParamsStatus = "needs_approval"
	PaymentOrderListParamsStatusPending       PaymentOrderListParamsStatus = "pending"
	PaymentOrderListParamsStatusProcessing    PaymentOrderListParamsStatus = "processing"
	PaymentOrderListParamsStatusReturned      PaymentOrderListParamsStatus = "returned"
	PaymentOrderListParamsStatusReversed      PaymentOrderListParamsStatus = "reversed"
	PaymentOrderListParamsStatusSent          PaymentOrderListParamsStatus = "sent"
)

type PaymentOrderListParamsDirection string

const (
	PaymentOrderListParamsDirectionCredit PaymentOrderListParamsDirection = "credit"
	PaymentOrderListParamsDirectionDebit  PaymentOrderListParamsDirection = "debit"
)

type PaymentOrderNewAsyncParams struct {
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type field.Field[PaymentOrderType] `json:"type,required"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype field.Field[PaymentOrderSubtype] `json:"subtype,nullable"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[PaymentOrderNewAsyncParamsDirection] `json:"direction,required"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority field.Field[PaymentOrderNewAsyncParamsPriority] `json:"priority"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID field.Field[string]                               `json:"receiving_account_id" format:"uuid"`
	Accounting         field.Field[PaymentOrderNewAsyncParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID field.Field[string] `json:"accounting_category_id,nullable" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID field.Field[string] `json:"accounting_ledger_class_id,nullable" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate field.Field[time.Time] `json:"effective_date" format:"date"`
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor field.Field[string] `json:"statement_descriptor,nullable"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation field.Field[string] `json:"remittance_information,nullable"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose field.Field[string] `json:"purpose,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer field.Field[PaymentOrderNewAsyncParamsChargeBearer] `json:"charge_bearer,nullable"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator field.Field[PaymentOrderNewAsyncParamsForeignExchangeIndicator] `json:"foreign_exchange_indicator,nullable"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract field.Field[string] `json:"foreign_exchange_contract,nullable"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected field.Field[bool] `json:"nsf_protected"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName field.Field[string] `json:"originating_party_name,nullable"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName field.Field[string] `json:"ultimate_originating_party_name,nullable"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier field.Field[string] `json:"ultimate_originating_party_identifier,nullable"`
	// Name of the ultimate funds recipient.
	UltimateReceivingPartyName field.Field[string] `json:"ultimate_receiving_party_name,nullable"`
	// Identifier of the ultimate funds recipient.
	UltimateReceivingPartyIdentifier field.Field[string] `json:"ultimate_receiving_party_identifier,nullable"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice field.Field[bool] `json:"send_remittance_advice,nullable"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt field.Field[time.Time] `json:"expires_at,nullable" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType field.Field[PaymentOrderNewAsyncParamsFallbackType] `json:"fallback_type"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount field.Field[PaymentOrderNewAsyncParamsReceivingAccount] `json:"receiving_account"`
	// Specifies a ledger transaction object that will be created with the payment
	// order. If the ledger transaction cannot be created, then the payment order
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the payment order.
	LedgerTransaction field.Field[PaymentOrderNewAsyncParamsLedgerTransaction] `json:"ledger_transaction"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems field.Field[[]PaymentOrderNewAsyncParamsLineItems] `json:"line_items"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	TransactionMonitoringEnabled field.Field[bool] `json:"transaction_monitoring_enabled"`
}

// MarshalJSON serializes PaymentOrderNewAsyncParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r PaymentOrderNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsDirection string

const (
	PaymentOrderNewAsyncParamsDirectionCredit PaymentOrderNewAsyncParamsDirection = "credit"
	PaymentOrderNewAsyncParamsDirectionDebit  PaymentOrderNewAsyncParamsDirection = "debit"
)

type PaymentOrderNewAsyncParamsPriority string

const (
	PaymentOrderNewAsyncParamsPriorityHigh   PaymentOrderNewAsyncParamsPriority = "high"
	PaymentOrderNewAsyncParamsPriorityNormal PaymentOrderNewAsyncParamsPriority = "normal"
)

type PaymentOrderNewAsyncParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID field.Field[string] `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID field.Field[string] `json:"class_id,nullable" format:"uuid"`
}

type PaymentOrderNewAsyncParamsChargeBearer string

const (
	PaymentOrderNewAsyncParamsChargeBearerShared   PaymentOrderNewAsyncParamsChargeBearer = "shared"
	PaymentOrderNewAsyncParamsChargeBearerSender   PaymentOrderNewAsyncParamsChargeBearer = "sender"
	PaymentOrderNewAsyncParamsChargeBearerReceiver PaymentOrderNewAsyncParamsChargeBearer = "receiver"
)

type PaymentOrderNewAsyncParamsForeignExchangeIndicator string

const (
	PaymentOrderNewAsyncParamsForeignExchangeIndicatorFixedToVariable PaymentOrderNewAsyncParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderNewAsyncParamsForeignExchangeIndicatorVariableToFixed PaymentOrderNewAsyncParamsForeignExchangeIndicator = "variable_to_fixed"
)

type PaymentOrderNewAsyncParamsFallbackType string

const (
	PaymentOrderNewAsyncParamsFallbackTypeACH PaymentOrderNewAsyncParamsFallbackType = "ach"
)

type PaymentOrderNewAsyncParamsReceivingAccount struct {
	// Can be `checking`, `savings` or `other`.
	AccountType field.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType field.Field[PaymentOrderNewAsyncParamsReceivingAccountPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress field.Field[PaymentOrderNewAsyncParamsReceivingAccountPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           field.Field[string]                                                     `json:"name,nullable"`
	AccountDetails field.Field[[]PaymentOrderNewAsyncParamsReceivingAccountAccountDetails] `json:"account_details"`
	RoutingDetails field.Field[[]PaymentOrderNewAsyncParamsReceivingAccountRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       field.Field[string] `json:"party_name"`
	PartyIdentifier field.Field[string] `json:"party_identifier"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken field.Field[string]                                                     `json:"plaid_processor_token"`
	ContactDetails      field.Field[[]PaymentOrderNewAsyncParamsReceivingAccountContactDetails] `json:"contact_details"`
}

type PaymentOrderNewAsyncParamsReceivingAccountPartyType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountPartyTypeBusiness   PaymentOrderNewAsyncParamsReceivingAccountPartyType = "business"
	PaymentOrderNewAsyncParamsReceivingAccountPartyTypeIndividual PaymentOrderNewAsyncParamsReceivingAccountPartyType = "individual"
)

type PaymentOrderNewAsyncParamsReceivingAccountPartyAddress struct {
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

type PaymentOrderNewAsyncParamsReceivingAccountAccountDetails struct {
	AccountNumber     field.Field[string]                                                                    `json:"account_number,required"`
	AccountNumberType field.Field[PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIban          PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeClabe         PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypePan           PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeOther         PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "other"
)

type PaymentOrderNewAsyncParamsReceivingAccountRoutingDetails struct {
	RoutingNumber     field.Field[string]                                                                    `json:"routing_number,required"`
	RoutingNumberType field.Field[PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       field.Field[PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

type PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba          PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo     PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode   PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "swift"
)

type PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "wire"
)

type PaymentOrderNewAsyncParamsReceivingAccountContactDetails struct {
	ContactIdentifier     field.Field[string]                                                                        `json:"contact_identifier"`
	ContactIdentifierType field.Field[PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeEmail       PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType = "email"
	PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypePhoneNumber PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeWebsite     PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType = "website"
)

type PaymentOrderNewAsyncParamsLedgerTransaction struct {
	// An optional description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status field.Field[PaymentOrderNewAsyncParamsLedgerTransactionStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate field.Field[time.Time] `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries field.Field[[]PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntries] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID field.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType field.Field[PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID field.Field[string] `json:"ledgerable_id" format:"uuid"`
}

type PaymentOrderNewAsyncParamsLedgerTransactionStatus string

const (
	PaymentOrderNewAsyncParamsLedgerTransactionStatusArchived PaymentOrderNewAsyncParamsLedgerTransactionStatus = "archived"
	PaymentOrderNewAsyncParamsLedgerTransactionStatusPending  PaymentOrderNewAsyncParamsLedgerTransactionStatus = "pending"
	PaymentOrderNewAsyncParamsLedgerTransactionStatusPosted   PaymentOrderNewAsyncParamsLedgerTransactionStatus = "posted"
)

type PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntries struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount field.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction field.Field[PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID field.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion field.Field[int64] `json:"lock_version,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount field.Field[map[string]int64] `json:"pending_balance_amount,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount field.Field[map[string]int64] `json:"posted_balance_amount,nullable"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount field.Field[map[string]int64] `json:"available_balance_amount,nullable"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances field.Field[bool] `json:"show_resulting_ledger_account_balances,nullable"`
}

type PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirection string

const (
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirectionCredit PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirection = "credit"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirectionDebit  PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirection = "debit"
)

type PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType string

const (
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeCounterparty          PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "counterparty"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeExpectedPayment       PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "expected_payment"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeIncomingPaymentDetail PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "incoming_payment_detail"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeInternalAccount       PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "internal_account"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeLineItem              PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "line_item"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypePaperItem             PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "paper_item"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypePaymentOrder          PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "payment_order"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypePaymentOrderAttempt   PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "payment_order_attempt"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeReturn                PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "return"
	PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeReversal              PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType = "reversal"
)

type PaymentOrderNewAsyncParamsLineItems struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount field.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description field.Field[string] `json:"description,nullable"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID field.Field[string] `json:"accounting_category_id,nullable"`
}
