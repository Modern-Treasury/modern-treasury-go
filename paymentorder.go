// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiform"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
	"github.com/tidwall/gjson"
)

// PaymentOrderService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentOrderService] method instead.
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
func (r *PaymentOrderService) New(ctx context.Context, body PaymentOrderNewParams, opts ...option.RequestOption) (res *PaymentOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/payment_orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single payment order
func (r *PaymentOrderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaymentOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a payment order
func (r *PaymentOrderService) Update(ctx context.Context, id string, body PaymentOrderUpdateParams, opts ...option.RequestOption) (res *PaymentOrder, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/payment_orders/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all payment orders
func (r *PaymentOrderService) List(ctx context.Context, query PaymentOrderListParams, opts ...option.RequestOption) (res *pagination.Page[PaymentOrder], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *PaymentOrderService) ListAutoPaging(ctx context.Context, query PaymentOrderListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PaymentOrder] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create a new payment order asynchronously
func (r *PaymentOrderService) NewAsync(ctx context.Context, body PaymentOrderNewAsyncParams, opts ...option.RequestOption) (res *shared.AsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/payment_orders/create_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ContactDetailCreateRequestParam struct {
	ContactIdentifier     param.Field[string]                                          `json:"contact_identifier"`
	ContactIdentifierType param.Field[ContactDetailCreateRequestContactIdentifierType] `json:"contact_identifier_type"`
}

func (r ContactDetailCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContactDetailCreateRequestContactIdentifierType string

const (
	ContactDetailCreateRequestContactIdentifierTypeEmail       ContactDetailCreateRequestContactIdentifierType = "email"
	ContactDetailCreateRequestContactIdentifierTypePhoneNumber ContactDetailCreateRequestContactIdentifierType = "phone_number"
	ContactDetailCreateRequestContactIdentifierTypeWebsite     ContactDetailCreateRequestContactIdentifierType = "website"
)

func (r ContactDetailCreateRequestContactIdentifierType) IsKnown() bool {
	switch r {
	case ContactDetailCreateRequestContactIdentifierTypeEmail, ContactDetailCreateRequestContactIdentifierTypePhoneNumber, ContactDetailCreateRequestContactIdentifierTypeWebsite:
		return true
	}
	return false
}

type PaymentOrder struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Accounting PaymentOrderAccounting `json:"accounting,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingCategoryID string `json:"accounting_category_id,required,nullable" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingLedgerClassID string `json:"accounting_ledger_class_id,required,nullable" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount int64 `json:"amount,required"`
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
	ChargeBearer PaymentOrderChargeBearer `json:"charge_bearer,required,nullable"`
	// If the payment order is tied to a specific Counterparty, their id will appear,
	// otherwise `null`.
	CounterpartyID string    `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// Defaults to the currency of the originating account.
	Currency shared.Currency `json:"currency,required"`
	// If the payment order's status is `returned`, this will include the return
	// object's data.
	CurrentReturn ReturnObject `json:"current_return,required,nullable"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction PaymentOrderDirection `json:"direction,required"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt time.Time `json:"expires_at,required,nullable" format:"date-time"`
	// An optional user-defined 180 character unique identifier.
	ExternalID string `json:"external_id,required,nullable"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract string `json:"foreign_exchange_contract,required,nullable"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator PaymentOrderForeignExchangeIndicator `json:"foreign_exchange_indicator,required,nullable"`
	// Associated serialized foreign exchange rate information.
	ForeignExchangeRate shared.ForeignExchangeRate `json:"foreign_exchange_rate,required,nullable"`
	// The ID of the ledger transaction linked to the payment order.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected bool   `json:"nsf_protected,required"`
	Object       string `json:"object,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID string `json:"originating_account_id,required" format:"uuid"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName string `json:"originating_party_name,required,nullable"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority PaymentOrderPriority `json:"priority,required"`
	// If present, Modern Treasury will not process the payment until after this time.
	// If `process_after` is past the cutoff for `effective_date`, `process_after` will
	// take precedence and `effective_date` will automatically update to reflect the
	// earliest possible sending date after `process_after`. Format is ISO8601
	// timestamp.
	ProcessAfter time.Time `json:"process_after,required,nullable" format:"date-time"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. For `eft`, this field is the 3
	// digit CPA Code that will be attached to the payment.
	Purpose string `json:"purpose,required,nullable"`
	// The receiving account ID. Can be an `external_account` or `internal_account`.
	ReceivingAccountID   string                           `json:"receiving_account_id,required" format:"uuid"`
	ReceivingAccountType PaymentOrderReceivingAccountType `json:"receiving_account_type,required"`
	ReferenceNumbers     []PaymentOrderReferenceNumber    `json:"reference_numbers,required"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice bool `json:"send_remittance_advice,required,nullable"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor string `json:"statement_descriptor,required,nullable"`
	// The current status of the payment order.
	Status PaymentOrderStatus `json:"status,required"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype PaymentOrderSubtype `json:"subtype,required,nullable"`
	// The IDs of all the transactions associated to this payment order. Usually, you
	// will only have a single transaction ID. However, if a payment order initially
	// results in a Return, but gets redrafted and is later successfully completed, it
	// can have many transactions.
	TransactionIDs []string `json:"transaction_ids,required" format:"uuid"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type PaymentOrderType `json:"type,required"`
	// The account to which the originating of this payment should be attributed to.
	// Can be a `virtual_account` or `internal_account`.
	UltimateOriginatingAccount PaymentOrderUltimateOriginatingAccount `json:"ultimate_originating_account,required,nullable"`
	// The ultimate originating account ID. Can be a `virtual_account` or
	// `internal_account`.
	UltimateOriginatingAccountID   string                                     `json:"ultimate_originating_account_id,required,nullable" format:"uuid"`
	UltimateOriginatingAccountType PaymentOrderUltimateOriginatingAccountType `json:"ultimate_originating_account_type,required,nullable"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier string `json:"ultimate_originating_party_identifier,required,nullable"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName     string    `json:"ultimate_originating_party_name,required,nullable"`
	UltimateReceivingPartyIdentifier string    `json:"ultimate_receiving_party_identifier,required,nullable"`
	UltimateReceivingPartyName       string    `json:"ultimate_receiving_party_name,required,nullable"`
	UpdatedAt                        time.Time `json:"updated_at,required" format:"date-time"`
	// Additional vendor specific fields for this payment. Data must be represented as
	// key-value pairs.
	VendorAttributes interface{} `json:"vendor_attributes,required,nullable"`
	// This field will be populated if a vendor failure occurs. Logic shouldn't be
	// built on its value as it is free-form.
	VendorFailureReason string           `json:"vendor_failure_reason,required,nullable"`
	JSON                paymentOrderJSON `json:"-"`
}

// paymentOrderJSON contains the JSON metadata for the struct [PaymentOrder]
type paymentOrderJSON struct {
	ID                                 apijson.Field
	Accounting                         apijson.Field
	AccountingCategoryID               apijson.Field
	AccountingLedgerClassID            apijson.Field
	Amount                             apijson.Field
	ChargeBearer                       apijson.Field
	CounterpartyID                     apijson.Field
	CreatedAt                          apijson.Field
	Currency                           apijson.Field
	CurrentReturn                      apijson.Field
	Description                        apijson.Field
	Direction                          apijson.Field
	EffectiveDate                      apijson.Field
	ExpiresAt                          apijson.Field
	ExternalID                         apijson.Field
	ForeignExchangeContract            apijson.Field
	ForeignExchangeIndicator           apijson.Field
	ForeignExchangeRate                apijson.Field
	LedgerTransactionID                apijson.Field
	LiveMode                           apijson.Field
	Metadata                           apijson.Field
	NsfProtected                       apijson.Field
	Object                             apijson.Field
	OriginatingAccountID               apijson.Field
	OriginatingPartyName               apijson.Field
	Priority                           apijson.Field
	ProcessAfter                       apijson.Field
	Purpose                            apijson.Field
	ReceivingAccountID                 apijson.Field
	ReceivingAccountType               apijson.Field
	ReferenceNumbers                   apijson.Field
	RemittanceInformation              apijson.Field
	SendRemittanceAdvice               apijson.Field
	StatementDescriptor                apijson.Field
	Status                             apijson.Field
	Subtype                            apijson.Field
	TransactionIDs                     apijson.Field
	Type                               apijson.Field
	UltimateOriginatingAccount         apijson.Field
	UltimateOriginatingAccountID       apijson.Field
	UltimateOriginatingAccountType     apijson.Field
	UltimateOriginatingPartyIdentifier apijson.Field
	UltimateOriginatingPartyName       apijson.Field
	UltimateReceivingPartyIdentifier   apijson.Field
	UltimateReceivingPartyName         apijson.Field
	UpdatedAt                          apijson.Field
	VendorAttributes                   apijson.Field
	VendorFailureReason                apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *PaymentOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentOrderJSON) RawJSON() string {
	return r.raw
}

func (r PaymentOrder) implementsBulkResultEntity() {}

// Deprecated: deprecated
type PaymentOrderAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	ClassID string                     `json:"class_id,nullable" format:"uuid"`
	JSON    paymentOrderAccountingJSON `json:"-"`
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

func (r paymentOrderAccountingJSON) RawJSON() string {
	return r.raw
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type PaymentOrderChargeBearer string

const (
	PaymentOrderChargeBearerShared   PaymentOrderChargeBearer = "shared"
	PaymentOrderChargeBearerSender   PaymentOrderChargeBearer = "sender"
	PaymentOrderChargeBearerReceiver PaymentOrderChargeBearer = "receiver"
)

func (r PaymentOrderChargeBearer) IsKnown() bool {
	switch r {
	case PaymentOrderChargeBearerShared, PaymentOrderChargeBearerSender, PaymentOrderChargeBearerReceiver:
		return true
	}
	return false
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type PaymentOrderDirection string

const (
	PaymentOrderDirectionCredit PaymentOrderDirection = "credit"
	PaymentOrderDirectionDebit  PaymentOrderDirection = "debit"
)

func (r PaymentOrderDirection) IsKnown() bool {
	switch r {
	case PaymentOrderDirectionCredit, PaymentOrderDirectionDebit:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type PaymentOrderForeignExchangeIndicator string

const (
	PaymentOrderForeignExchangeIndicatorFixedToVariable PaymentOrderForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderForeignExchangeIndicatorVariableToFixed PaymentOrderForeignExchangeIndicator = "variable_to_fixed"
)

func (r PaymentOrderForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case PaymentOrderForeignExchangeIndicatorFixedToVariable, PaymentOrderForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderPriority string

const (
	PaymentOrderPriorityHigh   PaymentOrderPriority = "high"
	PaymentOrderPriorityNormal PaymentOrderPriority = "normal"
)

func (r PaymentOrderPriority) IsKnown() bool {
	switch r {
	case PaymentOrderPriorityHigh, PaymentOrderPriorityNormal:
		return true
	}
	return false
}

type PaymentOrderReceivingAccountType string

const (
	PaymentOrderReceivingAccountTypeInternalAccount PaymentOrderReceivingAccountType = "internal_account"
	PaymentOrderReceivingAccountTypeExternalAccount PaymentOrderReceivingAccountType = "external_account"
)

func (r PaymentOrderReceivingAccountType) IsKnown() bool {
	switch r {
	case PaymentOrderReceivingAccountTypeInternalAccount, PaymentOrderReceivingAccountTypeExternalAccount:
		return true
	}
	return false
}

type PaymentOrderReferenceNumber struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The vendor reference number.
	ReferenceNumber string `json:"reference_number,required"`
	// The type of the reference number. Referring to the vendor payment id.
	ReferenceNumberType PaymentOrderReferenceNumbersReferenceNumberType `json:"reference_number_type,required"`
	UpdatedAt           time.Time                                       `json:"updated_at,required" format:"date-time"`
	JSON                paymentOrderReferenceNumberJSON                 `json:"-"`
}

// paymentOrderReferenceNumberJSON contains the JSON metadata for the struct
// [PaymentOrderReferenceNumber]
type paymentOrderReferenceNumberJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	LiveMode            apijson.Field
	Object              apijson.Field
	ReferenceNumber     apijson.Field
	ReferenceNumberType apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentOrderReferenceNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentOrderReferenceNumberJSON) RawJSON() string {
	return r.raw
}

// The type of the reference number. Referring to the vendor payment id.
type PaymentOrderReferenceNumbersReferenceNumberType string

const (
	PaymentOrderReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber                         PaymentOrderReferenceNumbersReferenceNumberType = "ach_original_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeACHTraceNumber                                 PaymentOrderReferenceNumbersReferenceNumberType = "ach_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate                    PaymentOrderReferenceNumbersReferenceNumberType = "bankprov_payment_activity_date"
	PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentID                              PaymentOrderReferenceNumbersReferenceNumberType = "bankprov_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID                        PaymentOrderReferenceNumbersReferenceNumberType = "bnk_dev_prenotification_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevTransferID                               PaymentOrderReferenceNumbersReferenceNumberType = "bnk_dev_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnyMellonTransactionReferenceNumber            PaymentOrderReferenceNumbersReferenceNumberType = "bny_mellon_transaction_reference_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeBofaEndToEndID                                 PaymentOrderReferenceNumbersReferenceNumberType = "bofa_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBofaTransactionID                              PaymentOrderReferenceNumbersReferenceNumberType = "bofa_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBraleTransferID                                PaymentOrderReferenceNumbersReferenceNumberType = "brale_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBridgeDestinationTransactionHash               PaymentOrderReferenceNumbersReferenceNumberType = "bridge_destination_transaction_hash"
	PaymentOrderReferenceNumbersReferenceNumberTypeBridgeSourceTransactionHash                    PaymentOrderReferenceNumbersReferenceNumberType = "bridge_source_transaction_hash"
	PaymentOrderReferenceNumbersReferenceNumberTypeBridgeTransferID                               PaymentOrderReferenceNumbersReferenceNumberType = "bridge_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCheckNumber                                    PaymentOrderReferenceNumbersReferenceNumberType = "check_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeCitibankReferenceNumber                        PaymentOrderReferenceNumbersReferenceNumberType = "citibank_reference_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber PaymentOrderReferenceNumbersReferenceNumberType = "citibank_worldlink_clearing_system_reference_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeColumnFxQuoteID                                PaymentOrderReferenceNumbersReferenceNumberType = "column_fx_quote_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeColumnReversalPairTransferID                   PaymentOrderReferenceNumbersReferenceNumberType = "column_reversal_pair_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeColumnTransferID                               PaymentOrderReferenceNumbersReferenceNumberType = "column_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverCoreTransactionID                    PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_core_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverFedBatchID                           PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_fed_batch_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverPaymentID                            PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverServiceMessage                       PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_service_message"
	PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverTransactionID                        PaymentOrderReferenceNumbersReferenceNumberType = "cross_river_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudConversionID                      PaymentOrderReferenceNumbersReferenceNumberType = "currencycloud_conversion_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID                         PaymentOrderReferenceNumbersReferenceNumberType = "currencycloud_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeDcBankTransactionID                            PaymentOrderReferenceNumbersReferenceNumberType = "dc_bank_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeEftTraceNumber                                 PaymentOrderReferenceNumbersReferenceNumberType = "eft_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreBatch                                PaymentOrderReferenceNumbersReferenceNumberType = "evolve_core_batch"
	PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreFileKey                              PaymentOrderReferenceNumbersReferenceNumberType = "evolve_core_file_key"
	PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreSeq                                  PaymentOrderReferenceNumbersReferenceNumberType = "evolve_core_seq"
	PaymentOrderReferenceNumbersReferenceNumberTypeEvolveTransactionID                            PaymentOrderReferenceNumbersReferenceNumberType = "evolve_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeFakeVendorPaymentID                            PaymentOrderReferenceNumbersReferenceNumberType = "fake_vendor_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeFedwireImad                                    PaymentOrderReferenceNumbersReferenceNumberType = "fedwire_imad"
	PaymentOrderReferenceNumbersReferenceNumberTypeFedwireOmad                                    PaymentOrderReferenceNumbersReferenceNumberType = "fedwire_omad"
	PaymentOrderReferenceNumbersReferenceNumberTypeFirstRepublicInternalID                        PaymentOrderReferenceNumbersReferenceNumberType = "first_republic_internal_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID                PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_collection_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID                         PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID                   PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_payment_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID                          PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID                    PaymentOrderReferenceNumbersReferenceNumberType = "goldman_sachs_unique_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeHifiOfframpID                                  PaymentOrderReferenceNumbersReferenceNumberType = "hifi_offramp_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeHifiTransferID                                 PaymentOrderReferenceNumbersReferenceNumberType = "hifi_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeInteracMessageID                               PaymentOrderReferenceNumbersReferenceNumberType = "interac_message_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCcn                                        PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_ccn"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcClearingSystemReference                    PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_clearing_system_reference"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID                        PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_customer_reference_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcEndToEndID                                 PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFirmRootID                                 PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_firm_root_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFxTrnID                                    PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_fx_trn_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcP3ID                                       PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_p3_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID                             PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_payment_batch_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID                       PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_payment_information_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentReturnedDatetime                    PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_payment_returned_datetime"
	PaymentOrderReferenceNumbersReferenceNumberTypeJpmcTransactionReferenceNumber                 PaymentOrderReferenceNumbersReferenceNumberType = "jpmc_transaction_reference_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeLobCheckID                                     PaymentOrderReferenceNumbersReferenceNumberType = "lob_check_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeMtFofTransferID                                PaymentOrderReferenceNumbersReferenceNumberType = "mt_fof_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeOther                                          PaymentOrderReferenceNumbersReferenceNumberType = "other"
	PaymentOrderReferenceNumbersReferenceNumberTypePartialSwiftMir                                PaymentOrderReferenceNumbersReferenceNumberType = "partial_swift_mir"
	PaymentOrderReferenceNumbersReferenceNumberTypePncClearingReference                           PaymentOrderReferenceNumbersReferenceNumberType = "pnc_clearing_reference"
	PaymentOrderReferenceNumbersReferenceNumberTypePncInstructionID                               PaymentOrderReferenceNumbersReferenceNumberType = "pnc_instruction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePncMultipaymentID                              PaymentOrderReferenceNumbersReferenceNumberType = "pnc_multipayment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePncPaymentTraceID                              PaymentOrderReferenceNumbersReferenceNumberType = "pnc_payment_trace_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePncRequestForPaymentID                         PaymentOrderReferenceNumbersReferenceNumberType = "pnc_request_for_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePncTransactionReferenceNumber                  PaymentOrderReferenceNumbersReferenceNumberType = "pnc_transaction_reference_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeRbcWireReferenceID                             PaymentOrderReferenceNumbersReferenceNumberType = "rbc_wire_reference_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeRtpInstructionID                               PaymentOrderReferenceNumbersReferenceNumberType = "rtp_instruction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSignetAPIReferenceID                           PaymentOrderReferenceNumbersReferenceNumberType = "signet_api_reference_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSignetConfirmationID                           PaymentOrderReferenceNumbersReferenceNumberType = "signet_confirmation_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSignetRequestID                                PaymentOrderReferenceNumbersReferenceNumberType = "signet_request_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSilvergatePaymentID                            PaymentOrderReferenceNumbersReferenceNumberType = "silvergate_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSvbEndToEndID                                  PaymentOrderReferenceNumbersReferenceNumberType = "svb_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSvbPaymentID                                   PaymentOrderReferenceNumbersReferenceNumberType = "svb_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeSvbTransactionClearedForSanctionsReview        PaymentOrderReferenceNumbersReferenceNumberType = "svb_transaction_cleared_for_sanctions_review"
	PaymentOrderReferenceNumbersReferenceNumberTypeSvbTransactionHeldForSanctionsReview           PaymentOrderReferenceNumbersReferenceNumberType = "svb_transaction_held_for_sanctions_review"
	PaymentOrderReferenceNumbersReferenceNumberTypeSwiftMir                                       PaymentOrderReferenceNumbersReferenceNumberType = "swift_mir"
	PaymentOrderReferenceNumbersReferenceNumberTypeSwiftUetr                                      PaymentOrderReferenceNumbersReferenceNumberType = "swift_uetr"
	PaymentOrderReferenceNumbersReferenceNumberTypeUmbProductPartnerAccountNumber                 PaymentOrderReferenceNumbersReferenceNumberType = "umb_product_partner_account_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentApplicationReferenceID            PaymentOrderReferenceNumbersReferenceNumberType = "usbank_payment_application_reference_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentID                                PaymentOrderReferenceNumbersReferenceNumberType = "usbank_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPendingRtpPaymentID                      PaymentOrderReferenceNumbersReferenceNumberType = "usbank_pending_rtp_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPostedRtpPaymentID                       PaymentOrderReferenceNumbersReferenceNumberType = "usbank_posted_rtp_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoEndToEndID                           PaymentOrderReferenceNumbersReferenceNumberType = "wells_fargo_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoPaymentID                            PaymentOrderReferenceNumbersReferenceNumberType = "wells_fargo_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber                          PaymentOrderReferenceNumbersReferenceNumberType = "wells_fargo_trace_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoUetr                                 PaymentOrderReferenceNumbersReferenceNumberType = "wells_fargo_uetr"
	PaymentOrderReferenceNumbersReferenceNumberTypeWesternAlliancePaymentID                       PaymentOrderReferenceNumbersReferenceNumberType = "western_alliance_payment_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWesternAllianceTransactionID                   PaymentOrderReferenceNumbersReferenceNumberType = "western_alliance_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeWesternAllianceWireConfirmationNumber          PaymentOrderReferenceNumbersReferenceNumberType = "western_alliance_wire_confirmation_number"
)

func (r PaymentOrderReferenceNumbersReferenceNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeACHTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate, PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID, PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeBnyMellonTransactionReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeBofaEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeBofaTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeBraleTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeBridgeDestinationTransactionHash, PaymentOrderReferenceNumbersReferenceNumberTypeBridgeSourceTransactionHash, PaymentOrderReferenceNumbersReferenceNumberTypeBridgeTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeCheckNumber, PaymentOrderReferenceNumbersReferenceNumberTypeCitibankReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeColumnFxQuoteID, PaymentOrderReferenceNumbersReferenceNumberTypeColumnReversalPairTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeColumnTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverCoreTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverFedBatchID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverServiceMessage, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudConversionID, PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeDcBankTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeEftTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreBatch, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreFileKey, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreSeq, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeFakeVendorPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeFedwireImad, PaymentOrderReferenceNumbersReferenceNumberTypeFedwireOmad, PaymentOrderReferenceNumbersReferenceNumberTypeFirstRepublicInternalID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeHifiOfframpID, PaymentOrderReferenceNumbersReferenceNumberTypeHifiTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeInteracMessageID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCcn, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcClearingSystemReference, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFirmRootID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFxTrnID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcP3ID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentReturnedDatetime, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcTransactionReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeLobCheckID, PaymentOrderReferenceNumbersReferenceNumberTypeMtFofTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeOther, PaymentOrderReferenceNumbersReferenceNumberTypePartialSwiftMir, PaymentOrderReferenceNumbersReferenceNumberTypePncClearingReference, PaymentOrderReferenceNumbersReferenceNumberTypePncInstructionID, PaymentOrderReferenceNumbersReferenceNumberTypePncMultipaymentID, PaymentOrderReferenceNumbersReferenceNumberTypePncPaymentTraceID, PaymentOrderReferenceNumbersReferenceNumberTypePncRequestForPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypePncTransactionReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeRbcWireReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeRtpInstructionID, PaymentOrderReferenceNumbersReferenceNumberTypeSignetAPIReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeSignetConfirmationID, PaymentOrderReferenceNumbersReferenceNumberTypeSignetRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeSilvergatePaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeSvbEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeSvbPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeSvbTransactionClearedForSanctionsReview, PaymentOrderReferenceNumbersReferenceNumberTypeSvbTransactionHeldForSanctionsReview, PaymentOrderReferenceNumbersReferenceNumberTypeSwiftMir, PaymentOrderReferenceNumbersReferenceNumberTypeSwiftUetr, PaymentOrderReferenceNumbersReferenceNumberTypeUmbProductPartnerAccountNumber, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentApplicationReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPendingRtpPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPostedRtpPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoUetr, PaymentOrderReferenceNumbersReferenceNumberTypeWesternAlliancePaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeWesternAllianceTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeWesternAllianceWireConfirmationNumber:
		return true
	}
	return false
}

// The current status of the payment order.
type PaymentOrderStatus string

const (
	PaymentOrderStatusApproved      PaymentOrderStatus = "approved"
	PaymentOrderStatusCancelled     PaymentOrderStatus = "cancelled"
	PaymentOrderStatusCompleted     PaymentOrderStatus = "completed"
	PaymentOrderStatusDenied        PaymentOrderStatus = "denied"
	PaymentOrderStatusFailed        PaymentOrderStatus = "failed"
	PaymentOrderStatusHeld          PaymentOrderStatus = "held"
	PaymentOrderStatusNeedsApproval PaymentOrderStatus = "needs_approval"
	PaymentOrderStatusPending       PaymentOrderStatus = "pending"
	PaymentOrderStatusProcessing    PaymentOrderStatus = "processing"
	PaymentOrderStatusReturned      PaymentOrderStatus = "returned"
	PaymentOrderStatusReversed      PaymentOrderStatus = "reversed"
	PaymentOrderStatusSent          PaymentOrderStatus = "sent"
	PaymentOrderStatusStopped       PaymentOrderStatus = "stopped"
)

func (r PaymentOrderStatus) IsKnown() bool {
	switch r {
	case PaymentOrderStatusApproved, PaymentOrderStatusCancelled, PaymentOrderStatusCompleted, PaymentOrderStatusDenied, PaymentOrderStatusFailed, PaymentOrderStatusHeld, PaymentOrderStatusNeedsApproval, PaymentOrderStatusPending, PaymentOrderStatusProcessing, PaymentOrderStatusReturned, PaymentOrderStatusReversed, PaymentOrderStatusSent, PaymentOrderStatusStopped:
		return true
	}
	return false
}

// The account to which the originating of this payment should be attributed to.
// Can be a `virtual_account` or `internal_account`.
type PaymentOrderUltimateOriginatingAccount struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of [[]AccountDetail].
	AccountDetails interface{} `json:"account_details,required"`
	// The ID of a counterparty that the virtual account belongs to. Optional.
	CounterpartyID string    `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// If the virtual account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// This field can have the runtime type of [map[string]string].
	Metadata interface{} `json:"metadata,required"`
	// The name of the virtual account.
	Name   string `json:"name,required,nullable"`
	Object string `json:"object,required"`
	// This field can have the runtime type of [[]RoutingDetail].
	RoutingDetails interface{} `json:"routing_details,required"`
	UpdatedAt      time.Time   `json:"updated_at,required" format:"date-time"`
	// This field can have the runtime type of [[]InternalAccountAccountCapability].
	AccountCapabilities interface{} `json:"account_capabilities"`
	// Can be checking, savings or other.
	AccountType PaymentOrderUltimateOriginatingAccountAccountType `json:"account_type,nullable"`
	// Specifies which financial institution the accounts belong to.
	Connection Connection `json:"connection"`
	// The ID of a credit normal ledger account. When money enters the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID string `json:"credit_ledger_account_id,nullable" format:"uuid"`
	// The currency of the account.
	Currency shared.Currency `json:"currency"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID string `json:"debit_ledger_account_id,nullable" format:"uuid"`
	// An optional free-form description for internal use.
	Description string    `json:"description,nullable"`
	DiscardedAt time.Time `json:"discarded_at,nullable" format:"date-time"`
	// The ID of the internal account that the virtual account is in.
	InternalAccountID string `json:"internal_account_id" format:"uuid"`
	// The Legal Entity associated to this account.
	LegalEntityID string `json:"legal_entity_id,nullable" format:"uuid"`
	// The parent InternalAccount of this account.
	ParentAccountID string `json:"parent_account_id,nullable" format:"uuid"`
	// The address associated with the owner or null.
	PartyAddress shared.Address `json:"party_address,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name"`
	// Either individual or business.
	PartyType PaymentOrderUltimateOriginatingAccountPartyType `json:"party_type,nullable"`
	// The internal account status.
	Status string `json:"status,nullable" format:"string"`
	// The vendor ID associated with this account.
	VendorID string                                     `json:"vendor_id,nullable" format:"string"`
	JSON     paymentOrderUltimateOriginatingAccountJSON `json:"-"`
	union    PaymentOrderUltimateOriginatingAccountUnion
}

// paymentOrderUltimateOriginatingAccountJSON contains the JSON metadata for the
// struct [PaymentOrderUltimateOriginatingAccount]
type paymentOrderUltimateOriginatingAccountJSON struct {
	ID                    apijson.Field
	AccountDetails        apijson.Field
	CounterpartyID        apijson.Field
	CreatedAt             apijson.Field
	LedgerAccountID       apijson.Field
	LiveMode              apijson.Field
	Metadata              apijson.Field
	Name                  apijson.Field
	Object                apijson.Field
	RoutingDetails        apijson.Field
	UpdatedAt             apijson.Field
	AccountCapabilities   apijson.Field
	AccountType           apijson.Field
	Connection            apijson.Field
	CreditLedgerAccountID apijson.Field
	Currency              apijson.Field
	DebitLedgerAccountID  apijson.Field
	Description           apijson.Field
	DiscardedAt           apijson.Field
	InternalAccountID     apijson.Field
	LegalEntityID         apijson.Field
	ParentAccountID       apijson.Field
	PartyAddress          apijson.Field
	PartyName             apijson.Field
	PartyType             apijson.Field
	Status                apijson.Field
	VendorID              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r paymentOrderUltimateOriginatingAccountJSON) RawJSON() string {
	return r.raw
}

func (r *PaymentOrderUltimateOriginatingAccount) UnmarshalJSON(data []byte) (err error) {
	*r = PaymentOrderUltimateOriginatingAccount{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PaymentOrderUltimateOriginatingAccountUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [VirtualAccount], [InternalAccount].
func (r PaymentOrderUltimateOriginatingAccount) AsUnion() PaymentOrderUltimateOriginatingAccountUnion {
	return r.union
}

// The account to which the originating of this payment should be attributed to.
// Can be a `virtual_account` or `internal_account`.
//
// Union satisfied by [VirtualAccount] or [InternalAccount].
type PaymentOrderUltimateOriginatingAccountUnion interface {
	implementsPaymentOrderUltimateOriginatingAccount()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PaymentOrderUltimateOriginatingAccountUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VirtualAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InternalAccount{}),
		},
	)
}

// Can be checking, savings or other.
type PaymentOrderUltimateOriginatingAccountAccountType string

const (
	PaymentOrderUltimateOriginatingAccountAccountTypeBaseWallet     PaymentOrderUltimateOriginatingAccountAccountType = "base_wallet"
	PaymentOrderUltimateOriginatingAccountAccountTypeCash           PaymentOrderUltimateOriginatingAccountAccountType = "cash"
	PaymentOrderUltimateOriginatingAccountAccountTypeChecking       PaymentOrderUltimateOriginatingAccountAccountType = "checking"
	PaymentOrderUltimateOriginatingAccountAccountTypeCryptoWallet   PaymentOrderUltimateOriginatingAccountAccountType = "crypto_wallet"
	PaymentOrderUltimateOriginatingAccountAccountTypeEthereumWallet PaymentOrderUltimateOriginatingAccountAccountType = "ethereum_wallet"
	PaymentOrderUltimateOriginatingAccountAccountTypeGeneralLedger  PaymentOrderUltimateOriginatingAccountAccountType = "general_ledger"
	PaymentOrderUltimateOriginatingAccountAccountTypeLoan           PaymentOrderUltimateOriginatingAccountAccountType = "loan"
	PaymentOrderUltimateOriginatingAccountAccountTypeNonResident    PaymentOrderUltimateOriginatingAccountAccountType = "non_resident"
	PaymentOrderUltimateOriginatingAccountAccountTypeOther          PaymentOrderUltimateOriginatingAccountAccountType = "other"
	PaymentOrderUltimateOriginatingAccountAccountTypeOverdraft      PaymentOrderUltimateOriginatingAccountAccountType = "overdraft"
	PaymentOrderUltimateOriginatingAccountAccountTypePolygonWallet  PaymentOrderUltimateOriginatingAccountAccountType = "polygon_wallet"
	PaymentOrderUltimateOriginatingAccountAccountTypeSavings        PaymentOrderUltimateOriginatingAccountAccountType = "savings"
	PaymentOrderUltimateOriginatingAccountAccountTypeSolanaWallet   PaymentOrderUltimateOriginatingAccountAccountType = "solana_wallet"
)

func (r PaymentOrderUltimateOriginatingAccountAccountType) IsKnown() bool {
	switch r {
	case PaymentOrderUltimateOriginatingAccountAccountTypeBaseWallet, PaymentOrderUltimateOriginatingAccountAccountTypeCash, PaymentOrderUltimateOriginatingAccountAccountTypeChecking, PaymentOrderUltimateOriginatingAccountAccountTypeCryptoWallet, PaymentOrderUltimateOriginatingAccountAccountTypeEthereumWallet, PaymentOrderUltimateOriginatingAccountAccountTypeGeneralLedger, PaymentOrderUltimateOriginatingAccountAccountTypeLoan, PaymentOrderUltimateOriginatingAccountAccountTypeNonResident, PaymentOrderUltimateOriginatingAccountAccountTypeOther, PaymentOrderUltimateOriginatingAccountAccountTypeOverdraft, PaymentOrderUltimateOriginatingAccountAccountTypePolygonWallet, PaymentOrderUltimateOriginatingAccountAccountTypeSavings, PaymentOrderUltimateOriginatingAccountAccountTypeSolanaWallet:
		return true
	}
	return false
}

// Either individual or business.
type PaymentOrderUltimateOriginatingAccountPartyType string

const (
	PaymentOrderUltimateOriginatingAccountPartyTypeBusiness   PaymentOrderUltimateOriginatingAccountPartyType = "business"
	PaymentOrderUltimateOriginatingAccountPartyTypeIndividual PaymentOrderUltimateOriginatingAccountPartyType = "individual"
)

func (r PaymentOrderUltimateOriginatingAccountPartyType) IsKnown() bool {
	switch r {
	case PaymentOrderUltimateOriginatingAccountPartyTypeBusiness, PaymentOrderUltimateOriginatingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type PaymentOrderUltimateOriginatingAccountType string

const (
	PaymentOrderUltimateOriginatingAccountTypeInternalAccount PaymentOrderUltimateOriginatingAccountType = "internal_account"
	PaymentOrderUltimateOriginatingAccountTypeVirtualAccount  PaymentOrderUltimateOriginatingAccountType = "virtual_account"
)

func (r PaymentOrderUltimateOriginatingAccountType) IsKnown() bool {
	switch r {
	case PaymentOrderUltimateOriginatingAccountTypeInternalAccount, PaymentOrderUltimateOriginatingAccountTypeVirtualAccount:
		return true
	}
	return false
}

// An additional layer of classification for the type of payment order you are
// doing. This field is only used for `ach` payment orders currently. For `ach`
// payment orders, the `subtype` represents the SEC code. We currently support
// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
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
	PaymentOrderSubtypeAuBecs                      PaymentOrderSubtype = "au_becs"
	PaymentOrderSubtypeBacs                        PaymentOrderSubtype = "bacs"
	PaymentOrderSubtypeChats                       PaymentOrderSubtype = "chats"
	PaymentOrderSubtypeDkNets                      PaymentOrderSubtype = "dk_nets"
	PaymentOrderSubtypeEft                         PaymentOrderSubtype = "eft"
	PaymentOrderSubtypeHuIcs                       PaymentOrderSubtype = "hu_ics"
	PaymentOrderSubtypeMasav                       PaymentOrderSubtype = "masav"
	PaymentOrderSubtypeMxCcen                      PaymentOrderSubtype = "mx_ccen"
	PaymentOrderSubtypeNeft                        PaymentOrderSubtype = "neft"
	PaymentOrderSubtypeNics                        PaymentOrderSubtype = "nics"
	PaymentOrderSubtypeNzBecs                      PaymentOrderSubtype = "nz_becs"
	PaymentOrderSubtypePlElixir                    PaymentOrderSubtype = "pl_elixir"
	PaymentOrderSubtypeRoSent                      PaymentOrderSubtype = "ro_sent"
	PaymentOrderSubtypeSeBankgirot                 PaymentOrderSubtype = "se_bankgirot"
	PaymentOrderSubtypeSepa                        PaymentOrderSubtype = "sepa"
	PaymentOrderSubtypeSgGiro                      PaymentOrderSubtype = "sg_giro"
	PaymentOrderSubtypeSic                         PaymentOrderSubtype = "sic"
	PaymentOrderSubtypeSknbi                       PaymentOrderSubtype = "sknbi"
	PaymentOrderSubtypeZengin                      PaymentOrderSubtype = "zengin"
)

func (r PaymentOrderSubtype) IsKnown() bool {
	switch r {
	case PaymentOrderSubtypeBacsNewInstruction, PaymentOrderSubtypeBacsCancellationInstruction, PaymentOrderSubtypeBacsConversionInstruction, PaymentOrderSubtypeCcd, PaymentOrderSubtypeCie, PaymentOrderSubtypeCtx, PaymentOrderSubtypeIat, PaymentOrderSubtypePpd, PaymentOrderSubtypeTel, PaymentOrderSubtypeWeb, PaymentOrderSubtypeAuBecs, PaymentOrderSubtypeBacs, PaymentOrderSubtypeChats, PaymentOrderSubtypeDkNets, PaymentOrderSubtypeEft, PaymentOrderSubtypeHuIcs, PaymentOrderSubtypeMasav, PaymentOrderSubtypeMxCcen, PaymentOrderSubtypeNeft, PaymentOrderSubtypeNics, PaymentOrderSubtypeNzBecs, PaymentOrderSubtypePlElixir, PaymentOrderSubtypeRoSent, PaymentOrderSubtypeSeBankgirot, PaymentOrderSubtypeSepa, PaymentOrderSubtypeSgGiro, PaymentOrderSubtypeSic, PaymentOrderSubtypeSknbi, PaymentOrderSubtypeZengin:
		return true
	}
	return false
}

// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
type PaymentOrderType string

const (
	PaymentOrderTypeACH         PaymentOrderType = "ach"
	PaymentOrderTypeAuBecs      PaymentOrderType = "au_becs"
	PaymentOrderTypeBacs        PaymentOrderType = "bacs"
	PaymentOrderTypeBase        PaymentOrderType = "base"
	PaymentOrderTypeBook        PaymentOrderType = "book"
	PaymentOrderTypeCard        PaymentOrderType = "card"
	PaymentOrderTypeChats       PaymentOrderType = "chats"
	PaymentOrderTypeCheck       PaymentOrderType = "check"
	PaymentOrderTypeCrossBorder PaymentOrderType = "cross_border"
	PaymentOrderTypeDkNets      PaymentOrderType = "dk_nets"
	PaymentOrderTypeEft         PaymentOrderType = "eft"
	PaymentOrderTypeEthereum    PaymentOrderType = "ethereum"
	PaymentOrderTypeGBFps       PaymentOrderType = "gb_fps"
	PaymentOrderTypeHuIcs       PaymentOrderType = "hu_ics"
	PaymentOrderTypeInterac     PaymentOrderType = "interac"
	PaymentOrderTypeMasav       PaymentOrderType = "masav"
	PaymentOrderTypeMxCcen      PaymentOrderType = "mx_ccen"
	PaymentOrderTypeNeft        PaymentOrderType = "neft"
	PaymentOrderTypeNics        PaymentOrderType = "nics"
	PaymentOrderTypeNzBecs      PaymentOrderType = "nz_becs"
	PaymentOrderTypePlElixir    PaymentOrderType = "pl_elixir"
	PaymentOrderTypePolygon     PaymentOrderType = "polygon"
	PaymentOrderTypeProvxchange PaymentOrderType = "provxchange"
	PaymentOrderTypeRoSent      PaymentOrderType = "ro_sent"
	PaymentOrderTypeRtp         PaymentOrderType = "rtp"
	PaymentOrderTypeSeBankgirot PaymentOrderType = "se_bankgirot"
	PaymentOrderTypeSen         PaymentOrderType = "sen"
	PaymentOrderTypeSepa        PaymentOrderType = "sepa"
	PaymentOrderTypeSgGiro      PaymentOrderType = "sg_giro"
	PaymentOrderTypeSic         PaymentOrderType = "sic"
	PaymentOrderTypeSignet      PaymentOrderType = "signet"
	PaymentOrderTypeSknbi       PaymentOrderType = "sknbi"
	PaymentOrderTypeSolana      PaymentOrderType = "solana"
	PaymentOrderTypeWire        PaymentOrderType = "wire"
	PaymentOrderTypeZengin      PaymentOrderType = "zengin"
)

func (r PaymentOrderType) IsKnown() bool {
	switch r {
	case PaymentOrderTypeACH, PaymentOrderTypeAuBecs, PaymentOrderTypeBacs, PaymentOrderTypeBase, PaymentOrderTypeBook, PaymentOrderTypeCard, PaymentOrderTypeChats, PaymentOrderTypeCheck, PaymentOrderTypeCrossBorder, PaymentOrderTypeDkNets, PaymentOrderTypeEft, PaymentOrderTypeEthereum, PaymentOrderTypeGBFps, PaymentOrderTypeHuIcs, PaymentOrderTypeInterac, PaymentOrderTypeMasav, PaymentOrderTypeMxCcen, PaymentOrderTypeNeft, PaymentOrderTypeNics, PaymentOrderTypeNzBecs, PaymentOrderTypePlElixir, PaymentOrderTypePolygon, PaymentOrderTypeProvxchange, PaymentOrderTypeRoSent, PaymentOrderTypeRtp, PaymentOrderTypeSeBankgirot, PaymentOrderTypeSen, PaymentOrderTypeSepa, PaymentOrderTypeSgGiro, PaymentOrderTypeSic, PaymentOrderTypeSignet, PaymentOrderTypeSknbi, PaymentOrderTypeSolana, PaymentOrderTypeWire, PaymentOrderTypeZengin:
		return true
	}
	return false
}

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
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type       param.Field[PaymentOrderType]                `json:"type,required"`
	Accounting param.Field[PaymentOrderNewParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
	ChargeBearer param.Field[PaymentOrderNewParamsChargeBearer] `json:"charge_bearer"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// An array of documents to be attached to the payment order. Note that if you
	// attach documents, the request's content type must be `multipart/form-data`.
	Documents param.Field[[]PaymentOrderNewParamsDocument] `json:"documents"`
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
	LedgerTransaction param.Field[shared.LedgerTransactionCreateRequestParam] `json:"ledger_transaction"`
	// Either ledger_transaction or ledger_transaction_id can be provided. Only a
	// pending ledger transaction can be attached upon payment order creation. Once the
	// payment order is created, the status of the ledger transaction tracks the
	// payment order automatically.
	LedgerTransactionID param.Field[string] `json:"ledger_transaction_id" format:"uuid"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]PaymentOrderNewParamsLineItem] `json:"line_items"`
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
	// If present, Modern Treasury will not process the payment until after this time.
	// If `process_after` is past the cutoff for `effective_date`, `process_after` will
	// take precedence and `effective_date` will automatically update to reflect the
	// earliest possible sending date after `process_after`. Format is ISO8601
	// timestamp.
	ProcessAfter param.Field[time.Time] `json:"process_after" format:"date-time"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. For `eft`, this field is the 3
	// digit CPA Code that will be attached to the payment.
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
	UltimateReceivingPartyName param.Field[string] `json:"ultimate_receiving_party_name"`
}

func (r PaymentOrderNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type PaymentOrderNewParamsDirection string

const (
	PaymentOrderNewParamsDirectionCredit PaymentOrderNewParamsDirection = "credit"
	PaymentOrderNewParamsDirectionDebit  PaymentOrderNewParamsDirection = "debit"
)

func (r PaymentOrderNewParamsDirection) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsDirectionCredit, PaymentOrderNewParamsDirectionDebit:
		return true
	}
	return false
}

// Deprecated: deprecated
type PaymentOrderNewParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r PaymentOrderNewParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type PaymentOrderNewParamsChargeBearer string

const (
	PaymentOrderNewParamsChargeBearerShared   PaymentOrderNewParamsChargeBearer = "shared"
	PaymentOrderNewParamsChargeBearerSender   PaymentOrderNewParamsChargeBearer = "sender"
	PaymentOrderNewParamsChargeBearerReceiver PaymentOrderNewParamsChargeBearer = "receiver"
)

func (r PaymentOrderNewParamsChargeBearer) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsChargeBearerShared, PaymentOrderNewParamsChargeBearerSender, PaymentOrderNewParamsChargeBearerReceiver:
		return true
	}
	return false
}

type PaymentOrderNewParamsDocument struct {
	// The unique identifier for the associated object.
	DocumentableID   param.Field[string]                                         `json:"documentable_id,required"`
	DocumentableType param.Field[PaymentOrderNewParamsDocumentsDocumentableType] `json:"documentable_type,required"`
	File             param.Field[io.Reader]                                      `json:"file,required" format:"binary"`
	// A category given to the document, can be `null`.
	DocumentType param.Field[string] `json:"document_type"`
}

func (r PaymentOrderNewParamsDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsDocumentsDocumentableType string

const (
	PaymentOrderNewParamsDocumentsDocumentableTypeCounterparties         PaymentOrderNewParamsDocumentsDocumentableType = "counterparties"
	PaymentOrderNewParamsDocumentsDocumentableTypeExpectedPayments       PaymentOrderNewParamsDocumentsDocumentableType = "expected_payments"
	PaymentOrderNewParamsDocumentsDocumentableTypeExternalAccounts       PaymentOrderNewParamsDocumentsDocumentableType = "external_accounts"
	PaymentOrderNewParamsDocumentsDocumentableTypeIdentifications        PaymentOrderNewParamsDocumentsDocumentableType = "identifications"
	PaymentOrderNewParamsDocumentsDocumentableTypeIncomingPaymentDetails PaymentOrderNewParamsDocumentsDocumentableType = "incoming_payment_details"
	PaymentOrderNewParamsDocumentsDocumentableTypeInternalAccounts       PaymentOrderNewParamsDocumentsDocumentableType = "internal_accounts"
	PaymentOrderNewParamsDocumentsDocumentableTypeOrganizations          PaymentOrderNewParamsDocumentsDocumentableType = "organizations"
	PaymentOrderNewParamsDocumentsDocumentableTypePaymentOrders          PaymentOrderNewParamsDocumentsDocumentableType = "payment_orders"
	PaymentOrderNewParamsDocumentsDocumentableTypeTransactions           PaymentOrderNewParamsDocumentsDocumentableType = "transactions"
	PaymentOrderNewParamsDocumentsDocumentableTypeConnections            PaymentOrderNewParamsDocumentsDocumentableType = "connections"
)

func (r PaymentOrderNewParamsDocumentsDocumentableType) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsDocumentsDocumentableTypeCounterparties, PaymentOrderNewParamsDocumentsDocumentableTypeExpectedPayments, PaymentOrderNewParamsDocumentsDocumentableTypeExternalAccounts, PaymentOrderNewParamsDocumentsDocumentableTypeIdentifications, PaymentOrderNewParamsDocumentsDocumentableTypeIncomingPaymentDetails, PaymentOrderNewParamsDocumentsDocumentableTypeInternalAccounts, PaymentOrderNewParamsDocumentsDocumentableTypeOrganizations, PaymentOrderNewParamsDocumentsDocumentableTypePaymentOrders, PaymentOrderNewParamsDocumentsDocumentableTypeTransactions, PaymentOrderNewParamsDocumentsDocumentableTypeConnections:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type PaymentOrderNewParamsFallbackType string

const (
	PaymentOrderNewParamsFallbackTypeACH PaymentOrderNewParamsFallbackType = "ach"
)

func (r PaymentOrderNewParamsFallbackType) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type PaymentOrderNewParamsForeignExchangeIndicator string

const (
	PaymentOrderNewParamsForeignExchangeIndicatorFixedToVariable PaymentOrderNewParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderNewParamsForeignExchangeIndicatorVariableToFixed PaymentOrderNewParamsForeignExchangeIndicator = "variable_to_fixed"
)

func (r PaymentOrderNewParamsForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsForeignExchangeIndicatorFixedToVariable, PaymentOrderNewParamsForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

type PaymentOrderNewParamsLineItem struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderNewParamsPriority string

const (
	PaymentOrderNewParamsPriorityHigh   PaymentOrderNewParamsPriority = "high"
	PaymentOrderNewParamsPriorityNormal PaymentOrderNewParamsPriority = "normal"
)

func (r PaymentOrderNewParamsPriority) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsPriorityHigh, PaymentOrderNewParamsPriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderNewParamsReceivingAccount struct {
	AccountDetails param.Field[[]PaymentOrderNewParamsReceivingAccountAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]               `json:"account_type"`
	ContactDetails param.Field[[]ContactDetailCreateRequestParam] `json:"contact_details"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[shared.LedgerAccountCreateRequestParam] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[shared.AddressRequestParam] `json:"party_address"`
	PartyIdentifier param.Field[string]                     `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[PaymentOrderNewParamsReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                               `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]PaymentOrderNewParamsReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r PaymentOrderNewParamsReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                               `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderNewParamsReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber        PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "au_number"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "base_address"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeClabe           PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeHkNumber        PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIban            PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIDNumber        PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "id_number"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeNzNumber        PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeOther           PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "other"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypePan             PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeSgNumber        PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeClabe, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeHkNumber, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIban, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIDNumber, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeNzNumber, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeOther, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypePan, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeSgNumber, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type PaymentOrderNewParamsReceivingAccountPartyType string

const (
	PaymentOrderNewParamsReceivingAccountPartyTypeBusiness   PaymentOrderNewParamsReceivingAccountPartyType = "business"
	PaymentOrderNewParamsReceivingAccountPartyTypeIndividual PaymentOrderNewParamsReceivingAccountPartyType = "individual"
)

func (r PaymentOrderNewParamsReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsReceivingAccountPartyTypeBusiness, PaymentOrderNewParamsReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type PaymentOrderNewParamsReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                               `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderNewParamsReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba                     PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips                   PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift, PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBase        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "base"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeChats       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "chats"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeDkNets      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeEthereum    PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "ethereum"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeGBFps       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeHuIcs       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeMxCcen      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNics        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "nics"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNzBecs      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypePlElixir    PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypePolygon     PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "polygon"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeRoSent      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSeBankgirot PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSgGiro      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSic         PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sic"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSknbi       PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "sknbi"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSolana      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "solana"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "wire"
	PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeZengin      PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBacs, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBase, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeBook, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCard, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeChats, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCheck, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeDkNets, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeEft, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeEthereum, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeGBFps, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeHuIcs, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeInterac, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeMasav, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeMxCcen, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNeft, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNics, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeNzBecs, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypePlElixir, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypePolygon, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeRoSent, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeRtp, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSen, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSepa, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSgGiro, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSic, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSignet, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSknbi, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeSolana, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeWire, PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

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
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
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
	LineItems param.Field[[]PaymentOrderUpdateParamsLineItem] `json:"line_items"`
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
	// If present, Modern Treasury will not process the payment until after this time.
	// If `process_after` is past the cutoff for `effective_date`, `process_after` will
	// take precedence and `effective_date` will automatically update to reflect the
	// earliest possible sending date after `process_after`. Format is ISO8601
	// timestamp.
	ProcessAfter param.Field[time.Time] `json:"process_after" format:"date-time"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. For `eft`, this field is the 3
	// digit CPA Code that will be attached to the payment.
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
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
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

// Deprecated: deprecated
type PaymentOrderUpdateParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r PaymentOrderUpdateParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type PaymentOrderUpdateParamsChargeBearer string

const (
	PaymentOrderUpdateParamsChargeBearerShared   PaymentOrderUpdateParamsChargeBearer = "shared"
	PaymentOrderUpdateParamsChargeBearerSender   PaymentOrderUpdateParamsChargeBearer = "sender"
	PaymentOrderUpdateParamsChargeBearerReceiver PaymentOrderUpdateParamsChargeBearer = "receiver"
)

func (r PaymentOrderUpdateParamsChargeBearer) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsChargeBearerShared, PaymentOrderUpdateParamsChargeBearerSender, PaymentOrderUpdateParamsChargeBearerReceiver:
		return true
	}
	return false
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type PaymentOrderUpdateParamsDirection string

const (
	PaymentOrderUpdateParamsDirectionCredit PaymentOrderUpdateParamsDirection = "credit"
	PaymentOrderUpdateParamsDirectionDebit  PaymentOrderUpdateParamsDirection = "debit"
)

func (r PaymentOrderUpdateParamsDirection) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsDirectionCredit, PaymentOrderUpdateParamsDirectionDebit:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type PaymentOrderUpdateParamsFallbackType string

const (
	PaymentOrderUpdateParamsFallbackTypeACH PaymentOrderUpdateParamsFallbackType = "ach"
)

func (r PaymentOrderUpdateParamsFallbackType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type PaymentOrderUpdateParamsForeignExchangeIndicator string

const (
	PaymentOrderUpdateParamsForeignExchangeIndicatorFixedToVariable PaymentOrderUpdateParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderUpdateParamsForeignExchangeIndicatorVariableToFixed PaymentOrderUpdateParamsForeignExchangeIndicator = "variable_to_fixed"
)

func (r PaymentOrderUpdateParamsForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsForeignExchangeIndicatorFixedToVariable, PaymentOrderUpdateParamsForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

type PaymentOrderUpdateParamsLineItem struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderUpdateParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderUpdateParamsPriority string

const (
	PaymentOrderUpdateParamsPriorityHigh   PaymentOrderUpdateParamsPriority = "high"
	PaymentOrderUpdateParamsPriorityNormal PaymentOrderUpdateParamsPriority = "normal"
)

func (r PaymentOrderUpdateParamsPriority) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsPriorityHigh, PaymentOrderUpdateParamsPriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderUpdateParamsReceivingAccount struct {
	AccountDetails param.Field[[]PaymentOrderUpdateParamsReceivingAccountAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]               `json:"account_type"`
	ContactDetails param.Field[[]ContactDetailCreateRequestParam] `json:"contact_details"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[shared.LedgerAccountCreateRequestParam] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[shared.AddressRequestParam] `json:"party_address"`
	PartyIdentifier param.Field[string]                     `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[PaymentOrderUpdateParamsReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                  `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]PaymentOrderUpdateParamsReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r PaymentOrderUpdateParamsReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                                  `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderUpdateParamsReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber        PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "au_number"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "base_address"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeClabe           PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeHkNumber        PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIban            PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIDNumber        PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "id_number"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeNzNumber        PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeOther           PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "other"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypePan             PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeSgNumber        PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeClabe, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeHkNumber, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIban, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIDNumber, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeNzNumber, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeOther, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypePan, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeSgNumber, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type PaymentOrderUpdateParamsReceivingAccountPartyType string

const (
	PaymentOrderUpdateParamsReceivingAccountPartyTypeBusiness   PaymentOrderUpdateParamsReceivingAccountPartyType = "business"
	PaymentOrderUpdateParamsReceivingAccountPartyTypeIndividual PaymentOrderUpdateParamsReceivingAccountPartyType = "individual"
)

func (r PaymentOrderUpdateParamsReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsReceivingAccountPartyTypeBusiness, PaymentOrderUpdateParamsReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type PaymentOrderUpdateParamsReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                                  `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderUpdateParamsReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba                     PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips                   PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBase        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "base"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeChats       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "chats"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeDkNets      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeEthereum    PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "ethereum"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeGBFps       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeHuIcs       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeMxCcen      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNics        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "nics"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNzBecs      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypePlElixir    PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypePolygon     PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "polygon"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeRoSent      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSeBankgirot PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSgGiro      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSic         PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sic"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSknbi       PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "sknbi"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSolana      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "solana"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "wire"
	PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeZengin      PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBacs, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBase, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeBook, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCard, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeChats, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCheck, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeDkNets, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeEft, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeEthereum, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeGBFps, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeHuIcs, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeInterac, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeMasav, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeMxCcen, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNeft, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNics, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeNzBecs, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypePlElixir, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypePolygon, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeRoSent, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeRtp, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSen, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSepa, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSgGiro, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSic, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSignet, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSknbi, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeSolana, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeWire, PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

// To cancel a payment order, use `cancelled`. To redraft a returned payment order,
// use `approved`. To undo approval on a denied or approved payment order, use
// `needs_approval`.
type PaymentOrderUpdateParamsStatus string

const (
	PaymentOrderUpdateParamsStatusApproved      PaymentOrderUpdateParamsStatus = "approved"
	PaymentOrderUpdateParamsStatusCancelled     PaymentOrderUpdateParamsStatus = "cancelled"
	PaymentOrderUpdateParamsStatusCompleted     PaymentOrderUpdateParamsStatus = "completed"
	PaymentOrderUpdateParamsStatusDenied        PaymentOrderUpdateParamsStatus = "denied"
	PaymentOrderUpdateParamsStatusFailed        PaymentOrderUpdateParamsStatus = "failed"
	PaymentOrderUpdateParamsStatusHeld          PaymentOrderUpdateParamsStatus = "held"
	PaymentOrderUpdateParamsStatusNeedsApproval PaymentOrderUpdateParamsStatus = "needs_approval"
	PaymentOrderUpdateParamsStatusPending       PaymentOrderUpdateParamsStatus = "pending"
	PaymentOrderUpdateParamsStatusProcessing    PaymentOrderUpdateParamsStatus = "processing"
	PaymentOrderUpdateParamsStatusReturned      PaymentOrderUpdateParamsStatus = "returned"
	PaymentOrderUpdateParamsStatusReversed      PaymentOrderUpdateParamsStatus = "reversed"
	PaymentOrderUpdateParamsStatusSent          PaymentOrderUpdateParamsStatus = "sent"
	PaymentOrderUpdateParamsStatusStopped       PaymentOrderUpdateParamsStatus = "stopped"
)

func (r PaymentOrderUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateParamsStatusApproved, PaymentOrderUpdateParamsStatusCancelled, PaymentOrderUpdateParamsStatusCompleted, PaymentOrderUpdateParamsStatusDenied, PaymentOrderUpdateParamsStatusFailed, PaymentOrderUpdateParamsStatusHeld, PaymentOrderUpdateParamsStatusNeedsApproval, PaymentOrderUpdateParamsStatusPending, PaymentOrderUpdateParamsStatusProcessing, PaymentOrderUpdateParamsStatusReturned, PaymentOrderUpdateParamsStatusReversed, PaymentOrderUpdateParamsStatusSent, PaymentOrderUpdateParamsStatusStopped:
		return true
	}
	return false
}

type PaymentOrderListParams struct {
	AfterCursor    param.Field[string] `query:"after_cursor"`
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// An inclusive upper bound for searching created_at
	CreatedAtEnd param.Field[time.Time] `query:"created_at_end" format:"date"`
	// An inclusive lower bound for searching created_at
	CreatedAtStart param.Field[time.Time]                   `query:"created_at_start" format:"date"`
	Direction      param.Field[shared.TransactionDirection] `query:"direction"`
	// An inclusive upper bound for searching effective_date
	EffectiveDateEnd param.Field[time.Time] `query:"effective_date_end" format:"date"`
	// An inclusive lower bound for searching effective_date
	EffectiveDateStart param.Field[time.Time] `query:"effective_date_start" format:"date"`
	ExternalID         param.Field[string]    `query:"external_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata             param.Field[map[string]string] `query:"metadata"`
	OriginatingAccountID param.Field[string]            `query:"originating_account_id"`
	PerPage              param.Field[int64]             `query:"per_page"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[PaymentOrderListParamsPriority] `query:"priority"`
	// An inclusive upper bound for searching process_after
	ProcessAfterEnd param.Field[time.Time] `query:"process_after_end" format:"date-time"`
	// An inclusive lower bound for searching process_after
	ProcessAfterStart param.Field[time.Time] `query:"process_after_start" format:"date-time"`
	// Query for records with the provided reference number
	ReferenceNumber param.Field[string]                       `query:"reference_number"`
	Status          param.Field[PaymentOrderListParamsStatus] `query:"status"`
	// The ID of a transaction that the payment order has been reconciled to.
	TransactionID param.Field[string]                     `query:"transaction_id"`
	Type          param.Field[PaymentOrderListParamsType] `query:"type"`
}

// URLQuery serializes [PaymentOrderListParams]'s query parameters as `url.Values`.
func (r PaymentOrderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderListParamsPriority string

const (
	PaymentOrderListParamsPriorityHigh   PaymentOrderListParamsPriority = "high"
	PaymentOrderListParamsPriorityNormal PaymentOrderListParamsPriority = "normal"
)

func (r PaymentOrderListParamsPriority) IsKnown() bool {
	switch r {
	case PaymentOrderListParamsPriorityHigh, PaymentOrderListParamsPriorityNormal:
		return true
	}
	return false
}

type PaymentOrderListParamsStatus string

const (
	PaymentOrderListParamsStatusApproved      PaymentOrderListParamsStatus = "approved"
	PaymentOrderListParamsStatusCancelled     PaymentOrderListParamsStatus = "cancelled"
	PaymentOrderListParamsStatusCompleted     PaymentOrderListParamsStatus = "completed"
	PaymentOrderListParamsStatusDenied        PaymentOrderListParamsStatus = "denied"
	PaymentOrderListParamsStatusFailed        PaymentOrderListParamsStatus = "failed"
	PaymentOrderListParamsStatusHeld          PaymentOrderListParamsStatus = "held"
	PaymentOrderListParamsStatusNeedsApproval PaymentOrderListParamsStatus = "needs_approval"
	PaymentOrderListParamsStatusPending       PaymentOrderListParamsStatus = "pending"
	PaymentOrderListParamsStatusProcessing    PaymentOrderListParamsStatus = "processing"
	PaymentOrderListParamsStatusReturned      PaymentOrderListParamsStatus = "returned"
	PaymentOrderListParamsStatusReversed      PaymentOrderListParamsStatus = "reversed"
	PaymentOrderListParamsStatusSent          PaymentOrderListParamsStatus = "sent"
	PaymentOrderListParamsStatusStopped       PaymentOrderListParamsStatus = "stopped"
)

func (r PaymentOrderListParamsStatus) IsKnown() bool {
	switch r {
	case PaymentOrderListParamsStatusApproved, PaymentOrderListParamsStatusCancelled, PaymentOrderListParamsStatusCompleted, PaymentOrderListParamsStatusDenied, PaymentOrderListParamsStatusFailed, PaymentOrderListParamsStatusHeld, PaymentOrderListParamsStatusNeedsApproval, PaymentOrderListParamsStatusPending, PaymentOrderListParamsStatusProcessing, PaymentOrderListParamsStatusReturned, PaymentOrderListParamsStatusReversed, PaymentOrderListParamsStatusSent, PaymentOrderListParamsStatusStopped:
		return true
	}
	return false
}

type PaymentOrderListParamsType string

const (
	PaymentOrderListParamsTypeACH         PaymentOrderListParamsType = "ach"
	PaymentOrderListParamsTypeAuBecs      PaymentOrderListParamsType = "au_becs"
	PaymentOrderListParamsTypeBacs        PaymentOrderListParamsType = "bacs"
	PaymentOrderListParamsTypeBase        PaymentOrderListParamsType = "base"
	PaymentOrderListParamsTypeBook        PaymentOrderListParamsType = "book"
	PaymentOrderListParamsTypeCard        PaymentOrderListParamsType = "card"
	PaymentOrderListParamsTypeChats       PaymentOrderListParamsType = "chats"
	PaymentOrderListParamsTypeCheck       PaymentOrderListParamsType = "check"
	PaymentOrderListParamsTypeCrossBorder PaymentOrderListParamsType = "cross_border"
	PaymentOrderListParamsTypeDkNets      PaymentOrderListParamsType = "dk_nets"
	PaymentOrderListParamsTypeEft         PaymentOrderListParamsType = "eft"
	PaymentOrderListParamsTypeEthereum    PaymentOrderListParamsType = "ethereum"
	PaymentOrderListParamsTypeGBFps       PaymentOrderListParamsType = "gb_fps"
	PaymentOrderListParamsTypeHuIcs       PaymentOrderListParamsType = "hu_ics"
	PaymentOrderListParamsTypeInterac     PaymentOrderListParamsType = "interac"
	PaymentOrderListParamsTypeMasav       PaymentOrderListParamsType = "masav"
	PaymentOrderListParamsTypeMxCcen      PaymentOrderListParamsType = "mx_ccen"
	PaymentOrderListParamsTypeNeft        PaymentOrderListParamsType = "neft"
	PaymentOrderListParamsTypeNics        PaymentOrderListParamsType = "nics"
	PaymentOrderListParamsTypeNzBecs      PaymentOrderListParamsType = "nz_becs"
	PaymentOrderListParamsTypePlElixir    PaymentOrderListParamsType = "pl_elixir"
	PaymentOrderListParamsTypePolygon     PaymentOrderListParamsType = "polygon"
	PaymentOrderListParamsTypeProvxchange PaymentOrderListParamsType = "provxchange"
	PaymentOrderListParamsTypeRoSent      PaymentOrderListParamsType = "ro_sent"
	PaymentOrderListParamsTypeRtp         PaymentOrderListParamsType = "rtp"
	PaymentOrderListParamsTypeSeBankgirot PaymentOrderListParamsType = "se_bankgirot"
	PaymentOrderListParamsTypeSen         PaymentOrderListParamsType = "sen"
	PaymentOrderListParamsTypeSepa        PaymentOrderListParamsType = "sepa"
	PaymentOrderListParamsTypeSgGiro      PaymentOrderListParamsType = "sg_giro"
	PaymentOrderListParamsTypeSic         PaymentOrderListParamsType = "sic"
	PaymentOrderListParamsTypeSignet      PaymentOrderListParamsType = "signet"
	PaymentOrderListParamsTypeSknbi       PaymentOrderListParamsType = "sknbi"
	PaymentOrderListParamsTypeSolana      PaymentOrderListParamsType = "solana"
	PaymentOrderListParamsTypeWire        PaymentOrderListParamsType = "wire"
	PaymentOrderListParamsTypeZengin      PaymentOrderListParamsType = "zengin"
)

func (r PaymentOrderListParamsType) IsKnown() bool {
	switch r {
	case PaymentOrderListParamsTypeACH, PaymentOrderListParamsTypeAuBecs, PaymentOrderListParamsTypeBacs, PaymentOrderListParamsTypeBase, PaymentOrderListParamsTypeBook, PaymentOrderListParamsTypeCard, PaymentOrderListParamsTypeChats, PaymentOrderListParamsTypeCheck, PaymentOrderListParamsTypeCrossBorder, PaymentOrderListParamsTypeDkNets, PaymentOrderListParamsTypeEft, PaymentOrderListParamsTypeEthereum, PaymentOrderListParamsTypeGBFps, PaymentOrderListParamsTypeHuIcs, PaymentOrderListParamsTypeInterac, PaymentOrderListParamsTypeMasav, PaymentOrderListParamsTypeMxCcen, PaymentOrderListParamsTypeNeft, PaymentOrderListParamsTypeNics, PaymentOrderListParamsTypeNzBecs, PaymentOrderListParamsTypePlElixir, PaymentOrderListParamsTypePolygon, PaymentOrderListParamsTypeProvxchange, PaymentOrderListParamsTypeRoSent, PaymentOrderListParamsTypeRtp, PaymentOrderListParamsTypeSeBankgirot, PaymentOrderListParamsTypeSen, PaymentOrderListParamsTypeSepa, PaymentOrderListParamsTypeSgGiro, PaymentOrderListParamsTypeSic, PaymentOrderListParamsTypeSignet, PaymentOrderListParamsTypeSknbi, PaymentOrderListParamsTypeSolana, PaymentOrderListParamsTypeWire, PaymentOrderListParamsTypeZengin:
		return true
	}
	return false
}

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
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type       param.Field[PaymentOrderType]                     `json:"type,required"`
	Accounting param.Field[PaymentOrderNewAsyncParamsAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
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
	LedgerTransaction param.Field[shared.LedgerTransactionCreateRequestParam] `json:"ledger_transaction"`
	// Either ledger_transaction or ledger_transaction_id can be provided. Only a
	// pending ledger transaction can be attached upon payment order creation. Once the
	// payment order is created, the status of the ledger transaction tracks the
	// payment order automatically.
	LedgerTransactionID param.Field[string] `json:"ledger_transaction_id" format:"uuid"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]PaymentOrderNewAsyncParamsLineItem] `json:"line_items"`
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
	// If present, Modern Treasury will not process the payment until after this time.
	// If `process_after` is past the cutoff for `effective_date`, `process_after` will
	// take precedence and `effective_date` will automatically update to reflect the
	// earliest possible sending date after `process_after`. Format is ISO8601
	// timestamp.
	ProcessAfter param.Field[time.Time] `json:"process_after" format:"date-time"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. For `eft`, this field is the 3
	// digit CPA Code that will be attached to the payment.
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
}

func (r PaymentOrderNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type PaymentOrderNewAsyncParamsDirection string

const (
	PaymentOrderNewAsyncParamsDirectionCredit PaymentOrderNewAsyncParamsDirection = "credit"
	PaymentOrderNewAsyncParamsDirectionDebit  PaymentOrderNewAsyncParamsDirection = "debit"
)

func (r PaymentOrderNewAsyncParamsDirection) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsDirectionCredit, PaymentOrderNewAsyncParamsDirectionDebit:
		return true
	}
	return false
}

// Deprecated: deprecated
type PaymentOrderNewAsyncParamsAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r PaymentOrderNewAsyncParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type PaymentOrderNewAsyncParamsChargeBearer string

const (
	PaymentOrderNewAsyncParamsChargeBearerShared   PaymentOrderNewAsyncParamsChargeBearer = "shared"
	PaymentOrderNewAsyncParamsChargeBearerSender   PaymentOrderNewAsyncParamsChargeBearer = "sender"
	PaymentOrderNewAsyncParamsChargeBearerReceiver PaymentOrderNewAsyncParamsChargeBearer = "receiver"
)

func (r PaymentOrderNewAsyncParamsChargeBearer) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsChargeBearerShared, PaymentOrderNewAsyncParamsChargeBearerSender, PaymentOrderNewAsyncParamsChargeBearerReceiver:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type PaymentOrderNewAsyncParamsFallbackType string

const (
	PaymentOrderNewAsyncParamsFallbackTypeACH PaymentOrderNewAsyncParamsFallbackType = "ach"
)

func (r PaymentOrderNewAsyncParamsFallbackType) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type PaymentOrderNewAsyncParamsForeignExchangeIndicator string

const (
	PaymentOrderNewAsyncParamsForeignExchangeIndicatorFixedToVariable PaymentOrderNewAsyncParamsForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderNewAsyncParamsForeignExchangeIndicatorVariableToFixed PaymentOrderNewAsyncParamsForeignExchangeIndicator = "variable_to_fixed"
)

func (r PaymentOrderNewAsyncParamsForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsForeignExchangeIndicatorFixedToVariable, PaymentOrderNewAsyncParamsForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

type PaymentOrderNewAsyncParamsLineItem struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PaymentOrderNewAsyncParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderNewAsyncParamsPriority string

const (
	PaymentOrderNewAsyncParamsPriorityHigh   PaymentOrderNewAsyncParamsPriority = "high"
	PaymentOrderNewAsyncParamsPriorityNormal PaymentOrderNewAsyncParamsPriority = "normal"
)

func (r PaymentOrderNewAsyncParamsPriority) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsPriorityHigh, PaymentOrderNewAsyncParamsPriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderNewAsyncParamsReceivingAccount struct {
	AccountDetails param.Field[[]PaymentOrderNewAsyncParamsReceivingAccountAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]               `json:"account_type"`
	ContactDetails param.Field[[]ContactDetailCreateRequestParam] `json:"contact_details"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[shared.LedgerAccountCreateRequestParam] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[shared.AddressRequestParam] `json:"party_address"`
	PartyIdentifier param.Field[string]                     `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[PaymentOrderNewAsyncParamsReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                    `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]PaymentOrderNewAsyncParamsReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                                    `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber        PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "au_number"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "base_address"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeClabe           PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeHkNumber        PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIban            PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIDNumber        PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "id_number"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeNzNumber        PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeOther           PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "other"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypePan             PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeSgNumber        PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeClabe, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeHkNumber, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIban, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIDNumber, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeNzNumber, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeOther, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypePan, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeSgNumber, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type PaymentOrderNewAsyncParamsReceivingAccountPartyType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountPartyTypeBusiness   PaymentOrderNewAsyncParamsReceivingAccountPartyType = "business"
	PaymentOrderNewAsyncParamsReceivingAccountPartyTypeIndividual PaymentOrderNewAsyncParamsReceivingAccountPartyType = "individual"
)

func (r PaymentOrderNewAsyncParamsReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsReceivingAccountPartyTypeBusiness, PaymentOrderNewAsyncParamsReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type PaymentOrderNewAsyncParamsReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                                    `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderNewAsyncParamsReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba                     PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips                   PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeChips, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeSwift, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBase        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "base"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeChats       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "chats"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeDkNets      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeEthereum    PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "ethereum"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeGBFps       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeHuIcs       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeMxCcen      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNics        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "nics"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNzBecs      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypePlElixir    PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypePolygon     PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "polygon"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeRoSent      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSeBankgirot PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSgGiro      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSic         PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sic"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSknbi       PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "sknbi"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSolana      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "solana"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "wire"
	PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeZengin      PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeAuBecs, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBacs, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBase, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeBook, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCard, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeChats, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCheck, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeCrossBorder, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeDkNets, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeEft, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeEthereum, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeGBFps, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeHuIcs, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeInterac, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeMasav, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeMxCcen, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNeft, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNics, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeNzBecs, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypePlElixir, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypePolygon, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeProvxchange, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeRoSent, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeRtp, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSen, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSepa, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSgGiro, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSic, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSignet, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSknbi, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeSolana, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeWire, PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}
