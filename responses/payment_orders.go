package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type PaymentOrder struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type PaymentOrderType `json:"type,required"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype PaymentOrderSubtype `json:"subtype,required,nullable"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount int64 `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction PaymentOrderDirection `json:"direction,required"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority PaymentOrderPriority `json:"priority,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID string `json:"originating_account_id,required" format:"uuid"`
	// The receiving account ID. Can be an `external_account` or `internal_account`.
	ReceivingAccountID string                 `json:"receiving_account_id,required" format:"uuid"`
	Accounting         PaymentOrderAccounting `json:"accounting,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID string `json:"accounting_category_id,required,nullable" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID string `json:"accounting_ledger_class_id,required,nullable" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency Currency `json:"currency,required,nullable"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor string `json:"statement_descriptor,required,nullable"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose string `json:"purpose,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer PaymentOrderChargeBearer `json:"charge_bearer,required,nullable"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator PaymentOrderForeignExchangeIndicator `json:"foreign_exchange_indicator,required,nullable"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract string `json:"foreign_exchange_contract,required,nullable"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected bool `json:"nsf_protected,required"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName string `json:"originating_party_name,required,nullable"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName string `json:"ultimate_originating_party_name,required,nullable"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier string `json:"ultimate_originating_party_identifier,required,nullable"`
	UltimateReceivingPartyName         string `json:"ultimate_receiving_party_name,required,nullable"`
	UltimateReceivingPartyIdentifier   string `json:"ultimate_receiving_party_identifier,required,nullable"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice bool `json:"send_remittance_advice,required,nullable"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt time.Time `json:"expires_at,required,nullable" format:"date-time"`
	// The current status of the payment order.
	Status               PaymentOrderStatus               `json:"status,required"`
	ReceivingAccountType PaymentOrderReceivingAccountType `json:"receiving_account_type,required"`
	// If the payment order is tied to a specific Counterparty, their id will appear,
	// otherwise `null`.
	CounterpartyID string `json:"counterparty_id,required,nullable" format:"uuid"`
	// The IDs of all the transactions associated to this payment order. Usually, you
	// will only have a single transaction ID. However, if a payment order initially
	// results in a Return, but gets redrafted and is later successfully completed, it
	// can have many transactions.
	TransactionIDs []string `json:"transaction_ids,required"`
	// The ID of the ledger transaction linked to the payment order.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// If the payment order's status is `returned`, this will include the return
	// object's data.
	CurrentReturn ReturnObject `json:"current_return,required,nullable"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	TransactionMonitoringEnabled bool `json:"transaction_monitoring_enabled,required"`
	// Custom key-value pair for usage in compliance rules. Please contact support
	// before making changes to this field.
	ComplianceRuleMetadata map[string]interface{}         `json:"compliance_rule_metadata,required,nullable"`
	ReferenceNumbers       []PaymentOrderReferenceNumbers `json:"reference_numbers,required"`
	// This field will be populated if a vendor (e.g. Currencycloud) failure occurs.
	// Logic shouldn't be built on its value as it is free-form.
	VendorFailureReason string `json:"vendor_failure_reason,required,nullable"`
	// The ID of the compliance decision for the payment order, if transaction
	// monitoring is enabled.
	DecisionID string `json:"decision_id,required,nullable" format:"uuid"`
	JSON       PaymentOrderJSON
}

type PaymentOrderJSON struct {
	ID                                 pjson.Metadata
	Object                             pjson.Metadata
	LiveMode                           pjson.Metadata
	CreatedAt                          pjson.Metadata
	UpdatedAt                          pjson.Metadata
	Type                               pjson.Metadata
	Subtype                            pjson.Metadata
	Amount                             pjson.Metadata
	Direction                          pjson.Metadata
	Priority                           pjson.Metadata
	OriginatingAccountID               pjson.Metadata
	ReceivingAccountID                 pjson.Metadata
	Accounting                         pjson.Metadata
	AccountingCategoryID               pjson.Metadata
	AccountingLedgerClassID            pjson.Metadata
	Currency                           pjson.Metadata
	EffectiveDate                      pjson.Metadata
	Description                        pjson.Metadata
	StatementDescriptor                pjson.Metadata
	RemittanceInformation              pjson.Metadata
	Purpose                            pjson.Metadata
	Metadata                           pjson.Metadata
	ChargeBearer                       pjson.Metadata
	ForeignExchangeIndicator           pjson.Metadata
	ForeignExchangeContract            pjson.Metadata
	NsfProtected                       pjson.Metadata
	OriginatingPartyName               pjson.Metadata
	UltimateOriginatingPartyName       pjson.Metadata
	UltimateOriginatingPartyIdentifier pjson.Metadata
	UltimateReceivingPartyName         pjson.Metadata
	UltimateReceivingPartyIdentifier   pjson.Metadata
	SendRemittanceAdvice               pjson.Metadata
	ExpiresAt                          pjson.Metadata
	Status                             pjson.Metadata
	ReceivingAccountType               pjson.Metadata
	CounterpartyID                     pjson.Metadata
	TransactionIDs                     pjson.Metadata
	LedgerTransactionID                pjson.Metadata
	CurrentReturn                      pjson.Metadata
	TransactionMonitoringEnabled       pjson.Metadata
	ComplianceRuleMetadata             pjson.Metadata
	ReferenceNumbers                   pjson.Metadata
	VendorFailureReason                pjson.Metadata
	DecisionID                         pjson.Metadata
	Raw                                []byte
	Extras                             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PaymentOrder using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PaymentOrder) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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

type PaymentOrderDirection string

const (
	PaymentOrderDirectionCredit PaymentOrderDirection = "credit"
	PaymentOrderDirectionDebit  PaymentOrderDirection = "debit"
)

type PaymentOrderPriority string

const (
	PaymentOrderPriorityHigh   PaymentOrderPriority = "high"
	PaymentOrderPriorityNormal PaymentOrderPriority = "normal"
)

type PaymentOrderAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID string `json:"class_id,nullable" format:"uuid"`
	JSON    PaymentOrderAccountingJSON
}

type PaymentOrderAccountingJSON struct {
	AccountID pjson.Metadata
	ClassID   pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PaymentOrderAccounting using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PaymentOrderAccounting) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PaymentOrderChargeBearer string

const (
	PaymentOrderChargeBearerShared   PaymentOrderChargeBearer = "shared"
	PaymentOrderChargeBearerSender   PaymentOrderChargeBearer = "sender"
	PaymentOrderChargeBearerReceiver PaymentOrderChargeBearer = "receiver"
)

type PaymentOrderForeignExchangeIndicator string

const (
	PaymentOrderForeignExchangeIndicatorFixedToVariable PaymentOrderForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderForeignExchangeIndicatorVariableToFixed PaymentOrderForeignExchangeIndicator = "variable_to_fixed"
)

type PaymentOrderStatus string

const (
	PaymentOrderStatusApproved      PaymentOrderStatus = "approved"
	PaymentOrderStatusCancelled     PaymentOrderStatus = "cancelled"
	PaymentOrderStatusCompleted     PaymentOrderStatus = "completed"
	PaymentOrderStatusDenied        PaymentOrderStatus = "denied"
	PaymentOrderStatusFailed        PaymentOrderStatus = "failed"
	PaymentOrderStatusNeedsApproval PaymentOrderStatus = "needs_approval"
	PaymentOrderStatusPending       PaymentOrderStatus = "pending"
	PaymentOrderStatusProcessing    PaymentOrderStatus = "processing"
	PaymentOrderStatusReturned      PaymentOrderStatus = "returned"
	PaymentOrderStatusReversed      PaymentOrderStatus = "reversed"
	PaymentOrderStatusSent          PaymentOrderStatus = "sent"
)

type PaymentOrderReceivingAccountType string

const (
	PaymentOrderReceivingAccountTypeInternalAccount PaymentOrderReceivingAccountType = "internal_account"
	PaymentOrderReceivingAccountTypeExternalAccount PaymentOrderReceivingAccountType = "external_account"
)

type PaymentOrderReferenceNumbers struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The vendor reference number.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of the reference number. Referring to the vendor payment id.
	ReferenceNumberType PaymentOrderReferenceNumbersReferenceNumberType `json:"reference_number_type,required"`
	JSON                PaymentOrderReferenceNumbersJSON
}

type PaymentOrderReferenceNumbersJSON struct {
	ID                  pjson.Metadata
	Object              pjson.Metadata
	LiveMode            pjson.Metadata
	CreatedAt           pjson.Metadata
	UpdatedAt           pjson.Metadata
	ReferenceNumber     pjson.Metadata
	ReferenceNumberType pjson.Metadata
	Raw                 []byte
	Extras              map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PaymentOrderReferenceNumbers
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PaymentOrderReferenceNumbers) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PaymentOrderReferenceNumbersReferenceNumberType string

const (
	PaymentOrderReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber          PaymentOrderReferenceNumbersReferenceNumberType = "ach_original_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeACHTraceNumber                  PaymentOrderReferenceNumbersReferenceNumberType = "ach_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate     PaymentOrderReferenceNumbersReferenceNumberType = "bankprov_payment_activity_date"
	PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentID               PaymentOrderReferenceNumbersReferenceNumberType = "bankprov_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID         PaymentOrderReferenceNumbersReferenceNumberType = "bnk_dev_prenotification_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevTransferID                PaymentOrderReferenceNumbersReferenceNumberType = "bnk_dev_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBofaEndToEndID                  PaymentOrderReferenceNumbersReferenceNumberType = "bofa_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBofaTransactionID               PaymentOrderReferenceNumbersReferenceNumberType = "bofa_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCheckNumber                     PaymentOrderReferenceNumbersReferenceNumberType = "check_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeColumnFxQuoteID                 PaymentOrderReferenceNumbersReferenceNumberType = "column_fx_quote_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeColumnTransferID                PaymentOrderReferenceNumbersReferenceNumberType = "column_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverPaymentID             PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverTransactionID         PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudConversionID       PaymentOrderReferenceNumbersReferenceNumberType = "currencycloud_conversion_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID          PaymentOrderReferenceNumbersReferenceNumberType = "currencycloud_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeDcBankTransactionID             PaymentOrderReferenceNumbersReferenceNumberType = "dc_bank_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeDwollaTransactionID             PaymentOrderReferenceNumbersReferenceNumberType = "dwolla_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeEftTraceNumber                  PaymentOrderReferenceNumbersReferenceNumberType = "eft_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeFedwireImad                     PaymentOrderReferenceNumbersReferenceNumberType = "fedwire_imad"
	PaymentOrderReferenceNumbersReferenceNumberTypeFedwireOmad                     PaymentOrderReferenceNumbersReferenceNumberType = "fedwire_omad"
	PaymentOrderReferenceNumbersReferenceNumberTypeFirstRepublicInternalID         PaymentOrderReferenceNumbersReferenceNumberType = "first_republic_internal_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_collection_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID          PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID    PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_payment_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID           PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID     PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_unique_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeInteracMessageID                PaymentOrderReferenceNumbersReferenceNumberType = "interac_message_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCcn                         PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_ccn"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID         PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_customer_reference_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcEndToEndID                  PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFirmRootID                  PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_firm_root_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcP3ID                        PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_p3_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID              PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_payment_batch_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID        PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_payment_information_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeLobCheckID                      PaymentOrderReferenceNumbersReferenceNumberType = "lob_check_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeOther                           PaymentOrderReferenceNumbersReferenceNumberType = "other"
	PaymentOrderReferenceNumbersReferenceNumberTypePartialSwiftMir                 PaymentOrderReferenceNumbersReferenceNumberType = "partial_swift_mir"
	PaymentOrderReferenceNumbersReferenceNumberTypePncClearingReference            PaymentOrderReferenceNumbersReferenceNumberType = "pnc_clearing_reference"
	PaymentOrderReferenceNumbersReferenceNumberTypePncInstructionID                PaymentOrderReferenceNumbersReferenceNumberType = "pnc_instruction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePncMultipaymentID               PaymentOrderReferenceNumbersReferenceNumberType = "pnc_multipayment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePncPaymentTraceID               PaymentOrderReferenceNumbersReferenceNumberType = "pnc_payment_trace_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeRtpInstructionID                PaymentOrderReferenceNumbersReferenceNumberType = "rtp_instruction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSignetAPIReferenceID            PaymentOrderReferenceNumbersReferenceNumberType = "signet_api_reference_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSignetConfirmationID            PaymentOrderReferenceNumbersReferenceNumberType = "signet_confirmation_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSignetRequestID                 PaymentOrderReferenceNumbersReferenceNumberType = "signet_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSilvergatePaymentID             PaymentOrderReferenceNumbersReferenceNumberType = "silvergate_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSwiftMir                        PaymentOrderReferenceNumbersReferenceNumberType = "swift_mir"
	PaymentOrderReferenceNumbersReferenceNumberTypeSwiftUetr                       PaymentOrderReferenceNumbersReferenceNumberType = "swift_uetr"
	PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentID                 PaymentOrderReferenceNumbersReferenceNumberType = "usbank_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoPaymentID             PaymentOrderReferenceNumbersReferenceNumberType = "wells_fargo_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber           PaymentOrderReferenceNumbersReferenceNumberType = "wells_fargo_trace_number"
)
