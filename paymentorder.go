// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// PaymentOrderService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewPaymentOrderService]
// method instead.
type PaymentOrderService struct {
	Options   []option.RequestOption
	Reversals *PaymentOrderReversalService
}

// NewPaymentOrderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentOrderService(opts ...option.RequestOption) (r *PaymentOrderService) {
	r = &PaymentOrderService{}
	r.Options = opts
	r.Reversals = NewPaymentOrderReversalService(opts...)
	return
}

// Create a new Payment Order
func (r *PaymentOrderService) New(ctx context.Context, params PaymentOrderNewParams, opts ...option.RequestOption) (res *PaymentOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/payment_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details on a single payment order
func (r *PaymentOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a payment order
func (r *PaymentOrderService) Update(ctx context.Context, id string, body PaymentOrderUpdateParams, opts ...option.RequestOption) (res *PaymentOrder, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/payment_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all payment orders
func (r *PaymentOrderService) List(ctx context.Context, query PaymentOrderListParams, opts ...option.RequestOption) (res *shared.Page[PaymentOrder], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/payment_orders"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of all payment orders
func (r *PaymentOrderService) ListAutoPaging(ctx context.Context, query PaymentOrderListParams, opts ...option.RequestOption) *shared.PageAutoPager[PaymentOrder] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create a new payment order asynchronously
func (r *PaymentOrderService) NewAsync(ctx context.Context, params PaymentOrderNewAsyncParams, opts ...option.RequestOption) (res *shared.AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/payment_orders/create_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

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
	Currency shared.Currency `json:"currency,required,nullable"`
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
	TransactionIDs []string `json:"transaction_ids,required" format:"uuid"`
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
	JSON       paymentOrderJSON
}

// paymentOrderJSON contains the JSON metadata for the struct [PaymentOrder]
type paymentOrderJSON struct {
	ID                                 apijson.Field
	Object                             apijson.Field
	LiveMode                           apijson.Field
	CreatedAt                          apijson.Field
	UpdatedAt                          apijson.Field
	Type                               apijson.Field
	Subtype                            apijson.Field
	Amount                             apijson.Field
	Direction                          apijson.Field
	Priority                           apijson.Field
	OriginatingAccountID               apijson.Field
	ReceivingAccountID                 apijson.Field
	Accounting                         apijson.Field
	AccountingCategoryID               apijson.Field
	AccountingLedgerClassID            apijson.Field
	Currency                           apijson.Field
	EffectiveDate                      apijson.Field
	Description                        apijson.Field
	StatementDescriptor                apijson.Field
	RemittanceInformation              apijson.Field
	Purpose                            apijson.Field
	Metadata                           apijson.Field
	ChargeBearer                       apijson.Field
	ForeignExchangeIndicator           apijson.Field
	ForeignExchangeContract            apijson.Field
	NsfProtected                       apijson.Field
	OriginatingPartyName               apijson.Field
	UltimateOriginatingPartyName       apijson.Field
	UltimateOriginatingPartyIdentifier apijson.Field
	UltimateReceivingPartyName         apijson.Field
	UltimateReceivingPartyIdentifier   apijson.Field
	SendRemittanceAdvice               apijson.Field
	ExpiresAt                          apijson.Field
	Status                             apijson.Field
	ReceivingAccountType               apijson.Field
	CounterpartyID                     apijson.Field
	TransactionIDs                     apijson.Field
	LedgerTransactionID                apijson.Field
	CurrentReturn                      apijson.Field
	TransactionMonitoringEnabled       apijson.Field
	ComplianceRuleMetadata             apijson.Field
	ReferenceNumbers                   apijson.Field
	VendorFailureReason                apijson.Field
	DecisionID                         apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *PaymentOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	PaymentOrderSubtypeBacsNewInstruction          PaymentOrderSubtype = "0C"
	PaymentOrderSubtypeBacsCancellationInstruction PaymentOrderSubtype = "0N"
	PaymentOrderSubtypeBacsConversionInstruction   PaymentOrderSubtype = "0S"
	PaymentOrderSubtypeCcd                         PaymentOrderSubtype = "CCD"
	PaymentOrderSubtypeCie                         PaymentOrderSubtype = "CIE"
	PaymentOrderSubtypeCtx                         PaymentOrderSubtype = "CTX"
	PaymentOrderSubtypeIat                         PaymentOrderSubtype = "IAT"
	PaymentOrderSubtypePpd                         PaymentOrderSubtype = "PPD"
	PaymentOrderSubtypeTel                         PaymentOrderSubtype = "TEL"
	PaymentOrderSubtypeWeb                         PaymentOrderSubtype = "WEB"
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
	JSON    paymentOrderAccountingJSON
}

// paymentOrderAccountingJSON contains the JSON metadata for the struct
// [PaymentOrderAccounting]
type paymentOrderAccountingJSON struct {
	AccountID   apijson.Field
	ClassID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentOrderAccounting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON                paymentOrderReferenceNumbersJSON
}

// paymentOrderReferenceNumbersJSON contains the JSON metadata for the struct
// [PaymentOrderReferenceNumbers]
type paymentOrderReferenceNumbersJSON struct {
	ID                  apijson.Field
	Object              apijson.Field
	LiveMode            apijson.Field
	CreatedAt           apijson.Field
	UpdatedAt           apijson.Field
	ReferenceNumber     apijson.Field
	ReferenceNumberType apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentOrderReferenceNumbers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	PaymentOrderReferenceNumbersReferenceNumberTypeEvolveTransactionID             PaymentOrderReferenceNumbersReferenceNumberType = "evolve_transaction_id"
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
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentReturnedDatetime     PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_payment_returned_datetime"
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

type PaymentOrderNewParams struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[PaymentOrderNewParamsDirection] `json:"direction,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type       param.Field[PaymentOrderType]                `json:"type,required"`
	Accounting param.Field[PaymentOrderNewParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer param.Field[PaymentOrderNewParamsChargeBearer] `json:"charge_bearer"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// An array of documents to be attached to the payment order. Note that if you
	// attach documents, the request's content type must be `multipart/form-data`.
	Documents param.Field[[]PaymentOrderNewParamsDocuments] `json:"documents"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[PaymentOrderNewParamsFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[PaymentOrderNewParamsForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// Specifies a ledger transaction object that will be created with the payment
	// order. If the ledger transaction cannot be created, then the payment order
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the payment order.
	LedgerTransaction param.Field[PaymentOrderNewParamsLedgerTransaction] `json:"ledger_transaction"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]PaymentOrderNewParamsLineItems] `json:"line_items"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected param.Field[bool] `json:"nsf_protected"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName param.Field[string] `json:"originating_party_name"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[PaymentOrderNewParamsPriority] `json:"priority"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose param.Field[string] `json:"purpose"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount param.Field[PaymentOrderNewParamsReceivingAccount] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice param.Field[bool] `json:"send_remittance_advice"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype param.Field[PaymentOrderSubtype] `json:"subtype"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	TransactionMonitoringEnabled param.Field[bool] `json:"transaction_monitoring_enabled"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier param.Field[string] `json:"ultimate_originating_party_identifier"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName param.Field[string] `json:"ultimate_originating_party_name"`
	// Identifier of the ultimate funds recipient.
	UltimateReceivingPartyIdentifier param.Field[string] `json:"ultimate_receiving_party_identifier"`
	// Name of the ultimate funds recipient.
	UltimateReceivingPartyName param.Field[string]                           `json:"ultimate_receiving_party_name"`
	ContentType                param.Field[PaymentOrderNewParamsContentType] `header:"Content-Type"`
	IdempotencyKey             param.Field[string]                           `header:"Idempotency-Key"`
}

func (r PaymentOrderNewParams) MarshalMultipart() (data []byte, err error) {
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	{
		bdy, err := apijson.Marshal(r.Amount)
		if err != nil {
			return nil, err
		}
		writer.WriteField("amount", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Direction)
		if err != nil {
			return nil, err
		}
		writer.WriteField("direction", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.OriginatingAccountID)
		if err != nil {
			return nil, err
		}
		writer.WriteField("originating_account_id", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Type)
		if err != nil {
			return nil, err
		}
		writer.WriteField("type", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Accounting)
		if err != nil {
			return nil, err
		}
		writer.WriteField("accounting", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.AccountingCategoryID)
		if err != nil {
			return nil, err
		}
		writer.WriteField("accounting_category_id", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.AccountingLedgerClassID)
		if err != nil {
			return nil, err
		}
		writer.WriteField("accounting_ledger_class_id", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ChargeBearer)
		if err != nil {
			return nil, err
		}
		writer.WriteField("charge_bearer", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Currency)
		if err != nil {
			return nil, err
		}
		writer.WriteField("currency", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Description)
		if err != nil {
			return nil, err
		}
		writer.WriteField("description", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Documents)
		if err != nil {
			return nil, err
		}
		writer.WriteField("documents", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.EffectiveDate)
		if err != nil {
			return nil, err
		}
		writer.WriteField("effective_date", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ExpiresAt)
		if err != nil {
			return nil, err
		}
		writer.WriteField("expires_at", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.FallbackType)
		if err != nil {
			return nil, err
		}
		writer.WriteField("fallback_type", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ForeignExchangeContract)
		if err != nil {
			return nil, err
		}
		writer.WriteField("foreign_exchange_contract", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ForeignExchangeIndicator)
		if err != nil {
			return nil, err
		}
		writer.WriteField("foreign_exchange_indicator", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.LedgerTransaction)
		if err != nil {
			return nil, err
		}
		writer.WriteField("ledger_transaction", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.LineItems)
		if err != nil {
			return nil, err
		}
		writer.WriteField("line_items", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Metadata)
		if err != nil {
			return nil, err
		}
		writer.WriteField("metadata", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.NsfProtected)
		if err != nil {
			return nil, err
		}
		writer.WriteField("nsf_protected", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.OriginatingPartyName)
		if err != nil {
			return nil, err
		}
		writer.WriteField("originating_party_name", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Priority)
		if err != nil {
			return nil, err
		}
		writer.WriteField("priority", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Purpose)
		if err != nil {
			return nil, err
		}
		writer.WriteField("purpose", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ReceivingAccount)
		if err != nil {
			return nil, err
		}
		writer.WriteField("receiving_account", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ReceivingAccountID)
		if err != nil {
			return nil, err
		}
		writer.WriteField("receiving_account_id", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.RemittanceInformation)
		if err != nil {
			return nil, err
		}
		writer.WriteField("remittance_information", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.SendRemittanceAdvice)
		if err != nil {
			return nil, err
		}
		writer.WriteField("send_remittance_advice", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.StatementDescriptor)
		if err != nil {
			return nil, err
		}
		writer.WriteField("statement_descriptor", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.Subtype)
		if err != nil {
			return nil, err
		}
		writer.WriteField("subtype", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.TransactionMonitoringEnabled)
		if err != nil {
			return nil, err
		}
		writer.WriteField("transaction_monitoring_enabled", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.UltimateOriginatingPartyIdentifier)
		if err != nil {
			return nil, err
		}
		writer.WriteField("ultimate_originating_party_identifier", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.UltimateOriginatingPartyName)
		if err != nil {
			return nil, err
		}
		writer.WriteField("ultimate_originating_party_name", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.UltimateReceivingPartyIdentifier)
		if err != nil {
			return nil, err
		}
		writer.WriteField("ultimate_receiving_party_identifier", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.UltimateReceivingPartyName)
		if err != nil {
			return nil, err
		}
		writer.WriteField("ultimate_receiving_party_name", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.ContentType)
		if err != nil {
			return nil, err
		}
		writer.WriteField("Content-Type", string(bdy))
	}
	{
		bdy, err := apijson.Marshal(r.IdempotencyKey)
		if err != nil {
			return nil, err
		}
		writer.WriteField("Idempotency-Key", string(bdy))
	}
	return body.Bytes(), nil
}

type PaymentOrderNewParamsDirection string

const (
	PaymentOrderNewParamsDirectionCredit PaymentOrderNewParamsDirection = "credit"
	PaymentOrderNewParamsDirectionDebit  PaymentOrderNewParamsDirection = "debit"
)

type PaymentOrderNewParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r PaymentOrderNewParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsChargeBearer string

const (
	PaymentOrderNewParamsChargeBearerShared   PaymentOrderNewParamsChargeBearer = "shared"
	PaymentOrderNewParamsChargeBearerSender   PaymentOrderNewParamsChargeBearer = "sender"
	PaymentOrderNewParamsChargeBearerReceiver PaymentOrderNewParamsChargeBearer = "receiver"
)

type PaymentOrderNewParamsDocuments struct {
	// A category given to the document, can be `null`.
	DocumentType param.Field[string]    `json:"document_type"`
	File         param.Field[io.Reader] `json:"file,required" format:"binary"`
}

func (r PaymentOrderNewParamsDocuments) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsFallbackType string

const (
	PaymentOrderNewParamsFallbackTypeACH PaymentOrderNewParamsFallbackType = "ach"
)

type PaymentOrderNewParamsForeignExchangeIndicator string

const (
	PaymentOrderNewParamsForeignExchangeIndicatorFixedToVariable PaymentOrderNewParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderNewParamsForeignExchangeIndicatorVariableToFixed PaymentOrderNewParamsForeignExchangeIndicator = "variable_to_fixed"
)

// Specifies a ledger transaction object that will be created with the payment
// order. If the ledger transaction cannot be created, then the payment order
// creation will fail. The resulting ledger transaction will mirror the status of
// the payment order.
type PaymentOrderNewParamsLedgerTransaction struct {
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[PaymentOrderNewParamsLedgerTransactionStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]PaymentOrderNewParamsLedgerTransactionLedgerEntries] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[PaymentOrderNewParamsLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
}

func (r PaymentOrderNewParamsLedgerTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r PaymentOrderNewParamsLedgerTransactionLedgerEntries) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Amount param.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
}

func (r PaymentOrderNewParamsLineItems) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsPriority string

const (
	PaymentOrderNewParamsPriorityHigh   PaymentOrderNewParamsPriority = "high"
	PaymentOrderNewParamsPriorityNormal PaymentOrderNewParamsPriority = "normal"
)

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderNewParamsReceivingAccount struct {
	// Can be `checking`, `savings` or `other`.
	AccountType param.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType param.Field[PaymentOrderNewParamsReceivingAccountPartyType] `json:"party_type"`
	// Required if receiving wire payments.
	PartyAddress param.Field[PaymentOrderNewParamsReceivingAccountPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           param.Field[string]                                                `json:"name"`
	AccountDetails param.Field[[]PaymentOrderNewParamsReceivingAccountAccountDetails] `json:"account_details"`
	RoutingDetails param.Field[[]PaymentOrderNewParamsReceivingAccountRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       param.Field[string] `json:"party_name"`
	PartyIdentifier param.Field[string] `json:"party_identifier"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[PaymentOrderNewParamsReceivingAccountLedgerAccount] `json:"ledger_account"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                `json:"plaid_processor_token"`
	ContactDetails      param.Field[[]PaymentOrderNewParamsReceivingAccountContactDetails] `json:"contact_details"`
}

func (r PaymentOrderNewParamsReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountPartyType string

const (
	PaymentOrderNewParamsReceivingAccountPartyTypeBusiness   PaymentOrderNewParamsReceivingAccountPartyType = "business"
	PaymentOrderNewParamsReceivingAccountPartyTypeIndividual PaymentOrderNewParamsReceivingAccountPartyType = "individual"
)

// Required if receiving wire payments.
type PaymentOrderNewParamsReceivingAccountPartyAddress struct {
	Line1 param.Field[string] `json:"line1"`
	Line2 param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// Region or State.
	Region param.Field[string] `json:"region"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
}

func (r PaymentOrderNewParamsReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountAccountDetails struct {
	AccountNumber     param.Field[string]                                                               `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderNewParamsReceivingAccountAccountDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	RoutingNumber     param.Field[string]                                                               `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderNewParamsReceivingAccountRoutingDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type PaymentOrderNewParamsReceivingAccountLedgerAccount struct {
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalance] `json:"normal_balance,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderNewParamsReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalance string

const (
	PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalanceCredit PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalance = "credit"
	PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalanceDebit  PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalance = "debit"
)

type PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableType string

const (
	PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableTypeExternalAccount PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableType = "external_account"
	PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableTypeInternalAccount PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableType = "internal_account"
)

type PaymentOrderNewParamsReceivingAccountContactDetails struct {
	ContactIdentifier     param.Field[string]                                                                   `json:"contact_identifier"`
	ContactIdentifierType param.Field[PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r PaymentOrderNewParamsReceivingAccountContactDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType string

const (
	PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeEmail       PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType = "email"
	PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypePhoneNumber PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeWebsite     PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierType = "website"
)

type PaymentOrderNewParamsContentType string

const (
	PaymentOrderNewParamsContentTypeApplicationJson   PaymentOrderNewParamsContentType = "application/json"
	PaymentOrderNewParamsContentTypeMultipartFormData PaymentOrderNewParamsContentType = "multipart/form-data"
)

type PaymentOrderUpdateParams struct {
	Accounting param.Field[PaymentOrderUpdateParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer param.Field[PaymentOrderUpdateParamsChargeBearer] `json:"charge_bearer"`
	// Required when receiving_account_id is passed the ID of an external account.
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[PaymentOrderUpdateParamsDirection] `json:"direction"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[PaymentOrderUpdateParamsFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[PaymentOrderUpdateParamsForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]PaymentOrderUpdateParamsLineItems] `json:"line_items"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected param.Field[bool] `json:"nsf_protected"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id" format:"uuid"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName param.Field[string] `json:"originating_party_name"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[PaymentOrderUpdateParamsPriority] `json:"priority"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose param.Field[string] `json:"purpose"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount param.Field[PaymentOrderUpdateParamsReceivingAccount] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice param.Field[bool] `json:"send_remittance_advice"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	// To cancel a payment order, use `cancelled`. To redraft a returned payment order,
	// use `approved`. To undo approval on a denied or approved payment order, use
	// `needs_approval`.
	Status param.Field[PaymentOrderUpdateParamsStatus] `json:"status"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype param.Field[PaymentOrderSubtype] `json:"subtype"`
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type param.Field[PaymentOrderType] `json:"type"`
	// This represents the identifier by which the person is known to the receiver when
	// using the CIE subtype for ACH payments. Only the first 22 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateOriginatingPartyIdentifier param.Field[string] `json:"ultimate_originating_party_identifier"`
	// This represents the name of the person that the payment is on behalf of when
	// using the CIE subtype for ACH payments. Only the first 15 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateOriginatingPartyName param.Field[string] `json:"ultimate_originating_party_name"`
	// This represents the name of the merchant that the payment is being sent to when
	// using the CIE subtype for ACH payments. Only the first 22 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateReceivingPartyIdentifier param.Field[string] `json:"ultimate_receiving_party_identifier"`
	// This represents the identifier by which the merchant is known to the person
	// initiating an ACH payment with CIE subtype. Only the first 15 characters of this
	// string will be used. Any additional characters will be truncated.
	UltimateReceivingPartyName param.Field[string] `json:"ultimate_receiving_party_name"`
}

func (r PaymentOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r PaymentOrderUpdateParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsChargeBearer string

const (
	PaymentOrderUpdateParamsChargeBearerShared   PaymentOrderUpdateParamsChargeBearer = "shared"
	PaymentOrderUpdateParamsChargeBearerSender   PaymentOrderUpdateParamsChargeBearer = "sender"
	PaymentOrderUpdateParamsChargeBearerReceiver PaymentOrderUpdateParamsChargeBearer = "receiver"
)

type PaymentOrderUpdateParamsDirection string

const (
	PaymentOrderUpdateParamsDirectionCredit PaymentOrderUpdateParamsDirection = "credit"
	PaymentOrderUpdateParamsDirectionDebit  PaymentOrderUpdateParamsDirection = "debit"
)

type PaymentOrderUpdateParamsFallbackType string

const (
	PaymentOrderUpdateParamsFallbackTypeACH PaymentOrderUpdateParamsFallbackType = "ach"
)

type PaymentOrderUpdateParamsForeignExchangeIndicator string

const (
	PaymentOrderUpdateParamsForeignExchangeIndicatorFixedToVariable PaymentOrderUpdateParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderUpdateParamsForeignExchangeIndicatorVariableToFixed PaymentOrderUpdateParamsForeignExchangeIndicator = "variable_to_fixed"
)

type PaymentOrderUpdateParamsLineItems struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
}

func (r PaymentOrderUpdateParamsLineItems) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsPriority string

const (
	PaymentOrderUpdateParamsPriorityHigh   PaymentOrderUpdateParamsPriority = "high"
	PaymentOrderUpdateParamsPriorityNormal PaymentOrderUpdateParamsPriority = "normal"
)

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderUpdateParamsReceivingAccount struct {
	// Can be `checking`, `savings` or `other`.
	AccountType param.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType param.Field[PaymentOrderUpdateParamsReceivingAccountPartyType] `json:"party_type"`
	// Required if receiving wire payments.
	PartyAddress param.Field[PaymentOrderUpdateParamsReceivingAccountPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           param.Field[string]                                                   `json:"name"`
	AccountDetails param.Field[[]PaymentOrderUpdateParamsReceivingAccountAccountDetails] `json:"account_details"`
	RoutingDetails param.Field[[]PaymentOrderUpdateParamsReceivingAccountRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       param.Field[string] `json:"party_name"`
	PartyIdentifier param.Field[string] `json:"party_identifier"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[PaymentOrderUpdateParamsReceivingAccountLedgerAccount] `json:"ledger_account"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                   `json:"plaid_processor_token"`
	ContactDetails      param.Field[[]PaymentOrderUpdateParamsReceivingAccountContactDetails] `json:"contact_details"`
}

func (r PaymentOrderUpdateParamsReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountPartyType string

const (
	PaymentOrderUpdateParamsReceivingAccountPartyTypeBusiness   PaymentOrderUpdateParamsReceivingAccountPartyType = "business"
	PaymentOrderUpdateParamsReceivingAccountPartyTypeIndividual PaymentOrderUpdateParamsReceivingAccountPartyType = "individual"
)

// Required if receiving wire payments.
type PaymentOrderUpdateParamsReceivingAccountPartyAddress struct {
	Line1 param.Field[string] `json:"line1"`
	Line2 param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// Region or State.
	Region param.Field[string] `json:"region"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
}

func (r PaymentOrderUpdateParamsReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountAccountDetails struct {
	AccountNumber     param.Field[string]                                                                  `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderUpdateParamsReceivingAccountAccountDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	RoutingNumber     param.Field[string]                                                                  `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderUpdateParamsReceivingAccountRoutingDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type PaymentOrderUpdateParamsReceivingAccountLedgerAccount struct {
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalance] `json:"normal_balance,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderUpdateParamsReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalance string

const (
	PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalanceCredit PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalance = "credit"
	PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalanceDebit  PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalance = "debit"
)

type PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableType string

const (
	PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableTypeExternalAccount PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableType = "external_account"
	PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableTypeInternalAccount PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableType = "internal_account"
)

type PaymentOrderUpdateParamsReceivingAccountContactDetails struct {
	ContactIdentifier     param.Field[string]                                                                      `json:"contact_identifier"`
	ContactIdentifierType param.Field[PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r PaymentOrderUpdateParamsReceivingAccountContactDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType string

const (
	PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeEmail       PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType = "email"
	PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypePhoneNumber PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeWebsite     PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierType = "website"
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

type PaymentOrderListParams struct {
	AfterCursor    param.Field[string]                          `query:"after_cursor"`
	CounterpartyID param.Field[string]                          `query:"counterparty_id" format:"uuid"`
	Direction      param.Field[PaymentOrderListParamsDirection] `query:"direction"`
	// An inclusive upper bound for searching effective_date
	EffectiveDateEnd param.Field[time.Time] `query:"effective_date_end" format:"date"`
	// An inclusive lower bound for searching effective_date
	EffectiveDateStart param.Field[time.Time] `query:"effective_date_start" format:"date"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata             param.Field[map[string]string] `query:"metadata"`
	OriginatingAccountID param.Field[string]            `query:"originating_account_id" format:"uuid"`
	PerPage              param.Field[int64]             `query:"per_page"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[PaymentOrderListParamsPriority] `query:"priority"`
	// Query for records with the provided reference number
	ReferenceNumber param.Field[string]                       `query:"reference_number"`
	Status          param.Field[PaymentOrderListParamsStatus] `query:"status"`
	// The ID of a transaction that the payment order has been reconciled to.
	TransactionID param.Field[string]                     `query:"transaction_id" format:"uuid"`
	Type          param.Field[PaymentOrderListParamsType] `query:"type"`
}

// URLQuery serializes [PaymentOrderListParams]'s query parameters as `url.Values`.
func (r PaymentOrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentOrderListParamsDirection string

const (
	PaymentOrderListParamsDirectionCredit PaymentOrderListParamsDirection = "credit"
	PaymentOrderListParamsDirectionDebit  PaymentOrderListParamsDirection = "debit"
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

type PaymentOrderNewAsyncParams struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[PaymentOrderNewAsyncParamsDirection] `json:"direction,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// One of `ach`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`, `bacs`,
	// `au_becs`, `interac`, `signet`, `provexchange`.
	Type       param.Field[PaymentOrderType]                     `json:"type,required"`
	Accounting param.Field[PaymentOrderNewAsyncParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer param.Field[PaymentOrderNewAsyncParamsChargeBearer] `json:"charge_bearer"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[PaymentOrderNewAsyncParamsFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[PaymentOrderNewAsyncParamsForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// Specifies a ledger transaction object that will be created with the payment
	// order. If the ledger transaction cannot be created, then the payment order
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the payment order.
	LedgerTransaction param.Field[PaymentOrderNewAsyncParamsLedgerTransaction] `json:"ledger_transaction"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]PaymentOrderNewAsyncParamsLineItems] `json:"line_items"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected param.Field[bool] `json:"nsf_protected"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName param.Field[string] `json:"originating_party_name"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[PaymentOrderNewAsyncParamsPriority] `json:"priority"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose param.Field[string] `json:"purpose"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount param.Field[PaymentOrderNewAsyncParamsReceivingAccount] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice param.Field[bool] `json:"send_remittance_advice"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype param.Field[PaymentOrderSubtype] `json:"subtype"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	TransactionMonitoringEnabled param.Field[bool] `json:"transaction_monitoring_enabled"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier param.Field[string] `json:"ultimate_originating_party_identifier"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName param.Field[string] `json:"ultimate_originating_party_name"`
	// Identifier of the ultimate funds recipient.
	UltimateReceivingPartyIdentifier param.Field[string] `json:"ultimate_receiving_party_identifier"`
	// Name of the ultimate funds recipient.
	UltimateReceivingPartyName param.Field[string] `json:"ultimate_receiving_party_name"`
	IdempotencyKey             param.Field[string] `header:"Idempotency-Key"`
}

func (r PaymentOrderNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsDirection string

const (
	PaymentOrderNewAsyncParamsDirectionCredit PaymentOrderNewAsyncParamsDirection = "credit"
	PaymentOrderNewAsyncParamsDirectionDebit  PaymentOrderNewAsyncParamsDirection = "debit"
)

type PaymentOrderNewAsyncParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r PaymentOrderNewAsyncParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsChargeBearer string

const (
	PaymentOrderNewAsyncParamsChargeBearerShared   PaymentOrderNewAsyncParamsChargeBearer = "shared"
	PaymentOrderNewAsyncParamsChargeBearerSender   PaymentOrderNewAsyncParamsChargeBearer = "sender"
	PaymentOrderNewAsyncParamsChargeBearerReceiver PaymentOrderNewAsyncParamsChargeBearer = "receiver"
)

type PaymentOrderNewAsyncParamsFallbackType string

const (
	PaymentOrderNewAsyncParamsFallbackTypeACH PaymentOrderNewAsyncParamsFallbackType = "ach"
)

type PaymentOrderNewAsyncParamsForeignExchangeIndicator string

const (
	PaymentOrderNewAsyncParamsForeignExchangeIndicatorFixedToVariable PaymentOrderNewAsyncParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderNewAsyncParamsForeignExchangeIndicatorVariableToFixed PaymentOrderNewAsyncParamsForeignExchangeIndicator = "variable_to_fixed"
)

// Specifies a ledger transaction object that will be created with the payment
// order. If the ledger transaction cannot be created, then the payment order
// creation will fail. The resulting ledger transaction will mirror the status of
// the payment order.
type PaymentOrderNewAsyncParamsLedgerTransaction struct {
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[PaymentOrderNewAsyncParamsLedgerTransactionStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntries] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[PaymentOrderNewAsyncParamsLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
}

func (r PaymentOrderNewAsyncParamsLedgerTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntries) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Amount param.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
}

func (r PaymentOrderNewAsyncParamsLineItems) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsPriority string

const (
	PaymentOrderNewAsyncParamsPriorityHigh   PaymentOrderNewAsyncParamsPriority = "high"
	PaymentOrderNewAsyncParamsPriorityNormal PaymentOrderNewAsyncParamsPriority = "normal"
)

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderNewAsyncParamsReceivingAccount struct {
	// Can be `checking`, `savings` or `other`.
	AccountType param.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType param.Field[PaymentOrderNewAsyncParamsReceivingAccountPartyType] `json:"party_type"`
	// Required if receiving wire payments.
	PartyAddress param.Field[PaymentOrderNewAsyncParamsReceivingAccountPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           param.Field[string]                                                     `json:"name"`
	AccountDetails param.Field[[]PaymentOrderNewAsyncParamsReceivingAccountAccountDetails] `json:"account_details"`
	RoutingDetails param.Field[[]PaymentOrderNewAsyncParamsReceivingAccountRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       param.Field[string] `json:"party_name"`
	PartyIdentifier param.Field[string] `json:"party_identifier"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[PaymentOrderNewAsyncParamsReceivingAccountLedgerAccount] `json:"ledger_account"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                     `json:"plaid_processor_token"`
	ContactDetails      param.Field[[]PaymentOrderNewAsyncParamsReceivingAccountContactDetails] `json:"contact_details"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountPartyType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountPartyTypeBusiness   PaymentOrderNewAsyncParamsReceivingAccountPartyType = "business"
	PaymentOrderNewAsyncParamsReceivingAccountPartyTypeIndividual PaymentOrderNewAsyncParamsReceivingAccountPartyType = "individual"
)

// Required if receiving wire payments.
type PaymentOrderNewAsyncParamsReceivingAccountPartyAddress struct {
	Line1 param.Field[string] `json:"line1"`
	Line2 param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// Region or State.
	Region param.Field[string] `json:"region"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountAccountDetails struct {
	AccountNumber     param.Field[string]                                                                    `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountAccountDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	RoutingNumber     param.Field[string]                                                                    `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountRoutingDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type PaymentOrderNewAsyncParamsReceivingAccountLedgerAccount struct {
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalance] `json:"normal_balance,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalance string

const (
	PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalanceCredit PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalance = "credit"
	PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalanceDebit  PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalance = "debit"
)

type PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableTypeExternalAccount PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableType = "external_account"
	PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableTypeInternalAccount PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableType = "internal_account"
)

type PaymentOrderNewAsyncParamsReceivingAccountContactDetails struct {
	ContactIdentifier     param.Field[string]                                                                        `json:"contact_identifier"`
	ContactIdentifierType param.Field[PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountContactDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeEmail       PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType = "email"
	PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypePhoneNumber PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeWebsite     PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierType = "website"
)
