// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

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
	Accounting Accounting `json:"accounting,required"`
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
	// If the payment order's status is `held`, this will include the hold object's
	// data.
	CurrentHold Hold `json:"current_hold,required,nullable"`
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
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus PaymentOrderReconciliationStatus `json:"reconciliation_status,required"`
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
	CurrentHold                        apijson.Field
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
	ReconciliationStatus               apijson.Field
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

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type PaymentOrderReconciliationStatus string

const (
	PaymentOrderReconciliationStatusUnreconciled          PaymentOrderReconciliationStatus = "unreconciled"
	PaymentOrderReconciliationStatusTentativelyReconciled PaymentOrderReconciliationStatus = "tentatively_reconciled"
	PaymentOrderReconciliationStatusReconciled            PaymentOrderReconciliationStatus = "reconciled"
)

func (r PaymentOrderReconciliationStatus) IsKnown() bool {
	switch r {
	case PaymentOrderReconciliationStatusUnreconciled, PaymentOrderReconciliationStatusTentativelyReconciled, PaymentOrderReconciliationStatusReconciled:
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
	PaymentOrderReferenceNumbersReferenceNumberTypeBlockchainTransactionHash                      PaymentOrderReferenceNumbersReferenceNumberType = "blockchain_transaction_hash"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID                        PaymentOrderReferenceNumbersReferenceNumberType = "bnk_dev_prenotification_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevTransferID                               PaymentOrderReferenceNumbersReferenceNumberType = "bnk_dev_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBnyMellonTransactionReferenceNumber            PaymentOrderReferenceNumbersReferenceNumberType = "bny_mellon_transaction_reference_number"
	PaymentOrderReferenceNumbersReferenceNumberTypeBofaEndToEndID                                 PaymentOrderReferenceNumbersReferenceNumberType = "bofa_end_to_end_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBofaTransactionID                              PaymentOrderReferenceNumbersReferenceNumberType = "bofa_transaction_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeBraleTransferID                                PaymentOrderReferenceNumbersReferenceNumberType = "brale_transfer_id"
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
	PaymentOrderReferenceNumbersReferenceNumberTypeMtFlowACHNocID                                 PaymentOrderReferenceNumbersReferenceNumberType = "mt_flow_ach_noc_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeMtFlowTransferID                               PaymentOrderReferenceNumbersReferenceNumberType = "mt_flow_transfer_id"
	PaymentOrderReferenceNumbersReferenceNumberTypeOther                                          PaymentOrderReferenceNumbersReferenceNumberType = "other"
	PaymentOrderReferenceNumbersReferenceNumberTypePartialSwiftMir                                PaymentOrderReferenceNumbersReferenceNumberType = "partial_swift_mir"
	PaymentOrderReferenceNumbersReferenceNumberTypePaxosOrchestrationID                           PaymentOrderReferenceNumbersReferenceNumberType = "paxos_orchestration_id"
	PaymentOrderReferenceNumbersReferenceNumberTypePaxosTransferID                                PaymentOrderReferenceNumbersReferenceNumberType = "paxos_transfer_id"
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
	case PaymentOrderReferenceNumbersReferenceNumberTypeACHOriginalTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeACHTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentActivityDate, PaymentOrderReferenceNumbersReferenceNumberTypeBankprovPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeBlockchainTransactionHash, PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevPrenotificationID, PaymentOrderReferenceNumbersReferenceNumberTypeBnkDevTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeBnyMellonTransactionReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeBofaEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeBofaTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeBraleTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeCheckNumber, PaymentOrderReferenceNumbersReferenceNumberTypeCitibankReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeCitibankWorldlinkClearingSystemReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeColumnFxQuoteID, PaymentOrderReferenceNumbersReferenceNumberTypeColumnReversalPairTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeColumnTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverCoreTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverFedBatchID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverServiceMessage, PaymentOrderReferenceNumbersReferenceNumberTypeCrossRiverTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudConversionID, PaymentOrderReferenceNumbersReferenceNumberTypeCurrencycloudPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeDcBankTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeEftTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreBatch, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreFileKey, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveCoreSeq, PaymentOrderReferenceNumbersReferenceNumberTypeEvolveTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeFakeVendorPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeFedwireImad, PaymentOrderReferenceNumbersReferenceNumberTypeFedwireOmad, PaymentOrderReferenceNumbersReferenceNumberTypeFirstRepublicInternalID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsCollectionRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsPaymentRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeGoldmanSachsUniquePaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeHifiOfframpID, PaymentOrderReferenceNumbersReferenceNumberTypeHifiTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeInteracMessageID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCcn, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcClearingSystemReference, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcCustomerReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFirmRootID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcFxTrnID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcP3ID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentBatchID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentInformationID, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcPaymentReturnedDatetime, PaymentOrderReferenceNumbersReferenceNumberTypeJpmcTransactionReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeLobCheckID, PaymentOrderReferenceNumbersReferenceNumberTypeMtFlowACHNocID, PaymentOrderReferenceNumbersReferenceNumberTypeMtFlowTransferID, PaymentOrderReferenceNumbersReferenceNumberTypeOther, PaymentOrderReferenceNumbersReferenceNumberTypePartialSwiftMir, PaymentOrderReferenceNumbersReferenceNumberTypePaxosOrchestrationID, PaymentOrderReferenceNumbersReferenceNumberTypePaxosTransferID, PaymentOrderReferenceNumbersReferenceNumberTypePncClearingReference, PaymentOrderReferenceNumbersReferenceNumberTypePncInstructionID, PaymentOrderReferenceNumbersReferenceNumberTypePncMultipaymentID, PaymentOrderReferenceNumbersReferenceNumberTypePncPaymentTraceID, PaymentOrderReferenceNumbersReferenceNumberTypePncRequestForPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypePncTransactionReferenceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeRbcWireReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeRtpInstructionID, PaymentOrderReferenceNumbersReferenceNumberTypeSignetAPIReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeSignetConfirmationID, PaymentOrderReferenceNumbersReferenceNumberTypeSignetRequestID, PaymentOrderReferenceNumbersReferenceNumberTypeSilvergatePaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeSvbEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeSvbPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeSvbTransactionClearedForSanctionsReview, PaymentOrderReferenceNumbersReferenceNumberTypeSvbTransactionHeldForSanctionsReview, PaymentOrderReferenceNumbersReferenceNumberTypeSwiftMir, PaymentOrderReferenceNumbersReferenceNumberTypeSwiftUetr, PaymentOrderReferenceNumbersReferenceNumberTypeUmbProductPartnerAccountNumber, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentApplicationReferenceID, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPendingRtpPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeUsbankPostedRtpPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoEndToEndID, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoPaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoTraceNumber, PaymentOrderReferenceNumbersReferenceNumberTypeWellsFargoUetr, PaymentOrderReferenceNumbersReferenceNumberTypeWesternAlliancePaymentID, PaymentOrderReferenceNumbersReferenceNumberTypeWesternAllianceTransactionID, PaymentOrderReferenceNumbersReferenceNumberTypeWesternAllianceWireConfirmationNumber:
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
	// This field can have the runtime type of [[]AccountCapability].
	AccountCapabilities interface{} `json:"account_capabilities"`
	// Can be checking, savings or other.
	AccountType PaymentOrderUltimateOriginatingAccountAccountType `json:"account_type,nullable"`
	// Specifies which financial institution the accounts belong to.
	Connection Connection `json:"connection"`
	// If the internal account links to a contra ledger account in Modern Treasury, the
	// id of the contra ledger account will be populated here.
	ContraLedgerAccountID string `json:"contra_ledger_account_id,nullable" format:"uuid"`
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
	Status PaymentOrderUltimateOriginatingAccountStatus `json:"status,nullable"`
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
	ContraLedgerAccountID apijson.Field
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

// The internal account status.
type PaymentOrderUltimateOriginatingAccountStatus string

const (
	PaymentOrderUltimateOriginatingAccountStatusActive            PaymentOrderUltimateOriginatingAccountStatus = "active"
	PaymentOrderUltimateOriginatingAccountStatusClosed            PaymentOrderUltimateOriginatingAccountStatus = "closed"
	PaymentOrderUltimateOriginatingAccountStatusPendingActivation PaymentOrderUltimateOriginatingAccountStatus = "pending_activation"
	PaymentOrderUltimateOriginatingAccountStatusPendingClosure    PaymentOrderUltimateOriginatingAccountStatus = "pending_closure"
	PaymentOrderUltimateOriginatingAccountStatusSuspended         PaymentOrderUltimateOriginatingAccountStatus = "suspended"
)

func (r PaymentOrderUltimateOriginatingAccountStatus) IsKnown() bool {
	switch r {
	case PaymentOrderUltimateOriginatingAccountStatusActive, PaymentOrderUltimateOriginatingAccountStatusClosed, PaymentOrderUltimateOriginatingAccountStatusPendingActivation, PaymentOrderUltimateOriginatingAccountStatusPendingClosure, PaymentOrderUltimateOriginatingAccountStatusSuspended:
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

type PaymentOrderAsyncCreateParam struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[PaymentOrderAsyncCreateDirection] `json:"direction,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type param.Field[PaymentOrderType] `json:"type,required"`
	// Deprecated: deprecated
	Accounting param.Field[AccountingParam] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
	ChargeBearer param.Field[PaymentOrderAsyncCreateChargeBearer] `json:"charge_bearer"`
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
	FallbackType param.Field[PaymentOrderAsyncCreateFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[PaymentOrderAsyncCreateForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
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
	LineItems param.Field[[]LineItemParam] `json:"line_items"`
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
	Priority param.Field[PaymentOrderAsyncCreatePriority] `json:"priority"`
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
	ReceivingAccount param.Field[PaymentOrderAsyncCreateReceivingAccountParam] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus param.Field[PaymentOrderAsyncCreateReconciliationStatus] `json:"reconciliation_status"`
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

func (r PaymentOrderAsyncCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PaymentOrderAsyncCreateParam) ImplementsBulkRequestNewParamsResourceUnion() {}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type PaymentOrderAsyncCreateDirection string

const (
	PaymentOrderAsyncCreateDirectionCredit PaymentOrderAsyncCreateDirection = "credit"
	PaymentOrderAsyncCreateDirectionDebit  PaymentOrderAsyncCreateDirection = "debit"
)

func (r PaymentOrderAsyncCreateDirection) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateDirectionCredit, PaymentOrderAsyncCreateDirectionDebit:
		return true
	}
	return false
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type PaymentOrderAsyncCreateChargeBearer string

const (
	PaymentOrderAsyncCreateChargeBearerShared   PaymentOrderAsyncCreateChargeBearer = "shared"
	PaymentOrderAsyncCreateChargeBearerSender   PaymentOrderAsyncCreateChargeBearer = "sender"
	PaymentOrderAsyncCreateChargeBearerReceiver PaymentOrderAsyncCreateChargeBearer = "receiver"
)

func (r PaymentOrderAsyncCreateChargeBearer) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateChargeBearerShared, PaymentOrderAsyncCreateChargeBearerSender, PaymentOrderAsyncCreateChargeBearerReceiver:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type PaymentOrderAsyncCreateFallbackType string

const (
	PaymentOrderAsyncCreateFallbackTypeACH PaymentOrderAsyncCreateFallbackType = "ach"
)

func (r PaymentOrderAsyncCreateFallbackType) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type PaymentOrderAsyncCreateForeignExchangeIndicator string

const (
	PaymentOrderAsyncCreateForeignExchangeIndicatorFixedToVariable PaymentOrderAsyncCreateForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderAsyncCreateForeignExchangeIndicatorVariableToFixed PaymentOrderAsyncCreateForeignExchangeIndicator = "variable_to_fixed"
)

func (r PaymentOrderAsyncCreateForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateForeignExchangeIndicatorFixedToVariable, PaymentOrderAsyncCreateForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderAsyncCreatePriority string

const (
	PaymentOrderAsyncCreatePriorityHigh   PaymentOrderAsyncCreatePriority = "high"
	PaymentOrderAsyncCreatePriorityNormal PaymentOrderAsyncCreatePriority = "normal"
)

func (r PaymentOrderAsyncCreatePriority) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreatePriorityHigh, PaymentOrderAsyncCreatePriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderAsyncCreateReceivingAccountParam struct {
	AccountDetails param.Field[[]PaymentOrderAsyncCreateReceivingAccountAccountDetailParam] `json:"account_details"`
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
	PartyType param.Field[PaymentOrderAsyncCreateReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                      `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]PaymentOrderAsyncCreateReceivingAccountRoutingDetailParam] `json:"routing_details"`
}

func (r PaymentOrderAsyncCreateReceivingAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderAsyncCreateReceivingAccountAccountDetailParam struct {
	AccountNumber     param.Field[string]                                                                 `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderAsyncCreateReceivingAccountAccountDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeAuNumber        PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "au_number"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "base_address"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeClabe           PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeHkNumber        PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeIban            PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeIDNumber        PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "id_number"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeNzNumber        PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeOther           PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "other"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypePan             PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeSgNumber        PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeAuNumber, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeClabe, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeHkNumber, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeIban, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeIDNumber, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeNzNumber, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeOther, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypePan, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeSgNumber, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, PaymentOrderAsyncCreateReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type PaymentOrderAsyncCreateReceivingAccountPartyType string

const (
	PaymentOrderAsyncCreateReceivingAccountPartyTypeBusiness   PaymentOrderAsyncCreateReceivingAccountPartyType = "business"
	PaymentOrderAsyncCreateReceivingAccountPartyTypeIndividual PaymentOrderAsyncCreateReceivingAccountPartyType = "individual"
)

func (r PaymentOrderAsyncCreateReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateReceivingAccountPartyTypeBusiness, PaymentOrderAsyncCreateReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type PaymentOrderAsyncCreateReceivingAccountRoutingDetailParam struct {
	RoutingNumber     param.Field[string]                                                                 `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderAsyncCreateReceivingAccountRoutingDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeAba                     PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeChips                   PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeAba, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeChips, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeSwift, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeBase        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "base"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeChats       PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "chats"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeDkNets      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeEthereum    PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "ethereum"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeGBFps       PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeHuIcs       PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeMxCcen      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeNics        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "nics"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeNzBecs      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypePlElixir    PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypePolygon     PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "polygon"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeRoSent      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSeBankgirot PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSgGiro      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSic         PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "sic"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSknbi       PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "sknbi"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSolana      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "solana"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "wire"
	PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeZengin      PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeACH, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeAuBecs, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeBacs, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeBase, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeBook, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeCard, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeChats, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeCheck, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeCrossBorder, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeDkNets, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeEft, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeEthereum, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeGBFps, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeHuIcs, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeInterac, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeMasav, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeMxCcen, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeNeft, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeNics, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeNzBecs, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypePlElixir, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypePolygon, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeProvxchange, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeRoSent, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeRtp, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSen, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSepa, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSgGiro, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSic, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSignet, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSknbi, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeSolana, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeWire, PaymentOrderAsyncCreateReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type PaymentOrderAsyncCreateReconciliationStatus string

const (
	PaymentOrderAsyncCreateReconciliationStatusUnreconciled          PaymentOrderAsyncCreateReconciliationStatus = "unreconciled"
	PaymentOrderAsyncCreateReconciliationStatusTentativelyReconciled PaymentOrderAsyncCreateReconciliationStatus = "tentatively_reconciled"
	PaymentOrderAsyncCreateReconciliationStatusReconciled            PaymentOrderAsyncCreateReconciliationStatus = "reconciled"
)

func (r PaymentOrderAsyncCreateReconciliationStatus) IsKnown() bool {
	switch r {
	case PaymentOrderAsyncCreateReconciliationStatusUnreconciled, PaymentOrderAsyncCreateReconciliationStatusTentativelyReconciled, PaymentOrderAsyncCreateReconciliationStatusReconciled:
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

type PaymentOrderUpdateParam struct {
	// Deprecated: deprecated
	Accounting param.Field[AccountingParam] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount"`
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
	ChargeBearer param.Field[PaymentOrderUpdateChargeBearer] `json:"charge_bearer"`
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
	Direction param.Field[PaymentOrderUpdateDirection] `json:"direction"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[PaymentOrderUpdateFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[PaymentOrderUpdateForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]LineItemParam] `json:"line_items"`
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
	Priority param.Field[PaymentOrderUpdatePriority] `json:"priority"`
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
	ReceivingAccount param.Field[PaymentOrderUpdateReceivingAccountParam] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus param.Field[PaymentOrderUpdateReconciliationStatus] `json:"reconciliation_status"`
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
	Status param.Field[PaymentOrderUpdateStatus] `json:"status"`
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

func (r PaymentOrderUpdateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type PaymentOrderUpdateChargeBearer string

const (
	PaymentOrderUpdateChargeBearerShared   PaymentOrderUpdateChargeBearer = "shared"
	PaymentOrderUpdateChargeBearerSender   PaymentOrderUpdateChargeBearer = "sender"
	PaymentOrderUpdateChargeBearerReceiver PaymentOrderUpdateChargeBearer = "receiver"
)

func (r PaymentOrderUpdateChargeBearer) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateChargeBearerShared, PaymentOrderUpdateChargeBearerSender, PaymentOrderUpdateChargeBearerReceiver:
		return true
	}
	return false
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type PaymentOrderUpdateDirection string

const (
	PaymentOrderUpdateDirectionCredit PaymentOrderUpdateDirection = "credit"
	PaymentOrderUpdateDirectionDebit  PaymentOrderUpdateDirection = "debit"
)

func (r PaymentOrderUpdateDirection) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateDirectionCredit, PaymentOrderUpdateDirectionDebit:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type PaymentOrderUpdateFallbackType string

const (
	PaymentOrderUpdateFallbackTypeACH PaymentOrderUpdateFallbackType = "ach"
)

func (r PaymentOrderUpdateFallbackType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type PaymentOrderUpdateForeignExchangeIndicator string

const (
	PaymentOrderUpdateForeignExchangeIndicatorFixedToVariable PaymentOrderUpdateForeignExchangeIndicator = "fixed_to_variable"
	PaymentOrderUpdateForeignExchangeIndicatorVariableToFixed PaymentOrderUpdateForeignExchangeIndicator = "variable_to_fixed"
)

func (r PaymentOrderUpdateForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateForeignExchangeIndicatorFixedToVariable, PaymentOrderUpdateForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type PaymentOrderUpdatePriority string

const (
	PaymentOrderUpdatePriorityHigh   PaymentOrderUpdatePriority = "high"
	PaymentOrderUpdatePriorityNormal PaymentOrderUpdatePriority = "normal"
)

func (r PaymentOrderUpdatePriority) IsKnown() bool {
	switch r {
	case PaymentOrderUpdatePriorityHigh, PaymentOrderUpdatePriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type PaymentOrderUpdateReceivingAccountParam struct {
	AccountDetails param.Field[[]PaymentOrderUpdateReceivingAccountAccountDetailParam] `json:"account_details"`
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
	PartyType param.Field[PaymentOrderUpdateReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                 `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]PaymentOrderUpdateReceivingAccountRoutingDetailParam] `json:"routing_details"`
}

func (r PaymentOrderUpdateReceivingAccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateReceivingAccountAccountDetailParam struct {
	AccountNumber     param.Field[string]                                                            `json:"account_number,required"`
	AccountNumberType param.Field[PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r PaymentOrderUpdateReceivingAccountAccountDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType string

const (
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeAuNumber        PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "au_number"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "base_address"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeClabe           PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "clabe"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeHkNumber        PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeIban            PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "iban"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeIDNumber        PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "id_number"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeNzNumber        PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeOther           PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "other"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypePan             PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "pan"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeSgNumber        PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeAuNumber, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeClabe, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeHkNumber, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeIban, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeIDNumber, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeNzNumber, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeOther, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypePan, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeSgNumber, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, PaymentOrderUpdateReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type PaymentOrderUpdateReceivingAccountPartyType string

const (
	PaymentOrderUpdateReceivingAccountPartyTypeBusiness   PaymentOrderUpdateReceivingAccountPartyType = "business"
	PaymentOrderUpdateReceivingAccountPartyTypeIndividual PaymentOrderUpdateReceivingAccountPartyType = "individual"
)

func (r PaymentOrderUpdateReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateReceivingAccountPartyTypeBusiness, PaymentOrderUpdateReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type PaymentOrderUpdateReceivingAccountRoutingDetailParam struct {
	RoutingNumber     param.Field[string]                                                            `json:"routing_number,required"`
	RoutingNumberType param.Field[PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r PaymentOrderUpdateReceivingAccountRoutingDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType string

const (
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeAba                     PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeChips                   PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeAba, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeChips, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeSwift, PaymentOrderUpdateReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType string

const (
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeACH         PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "ach"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeAuBecs      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "au_becs"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeBacs        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "bacs"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeBase        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "base"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeBook        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "book"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeCard        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "card"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeChats       PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "chats"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeCheck       PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "check"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeCrossBorder PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "cross_border"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeDkNets      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeEft         PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "eft"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeEthereum    PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "ethereum"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeGBFps       PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeHuIcs       PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeInterac     PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "interac"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeMasav       PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "masav"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeMxCcen      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeNeft        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "neft"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeNics        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "nics"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeNzBecs      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypePlElixir    PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypePolygon     PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "polygon"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeProvxchange PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "provxchange"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeRoSent      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeRtp         PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "rtp"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSeBankgirot PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSen         PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "sen"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSepa        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "sepa"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSgGiro      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSic         PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "sic"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSignet      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "signet"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSknbi       PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "sknbi"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSolana      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "solana"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeWire        PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "wire"
	PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeZengin      PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeACH, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeAuBecs, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeBacs, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeBase, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeBook, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeCard, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeChats, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeCheck, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeCrossBorder, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeDkNets, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeEft, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeEthereum, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeGBFps, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeHuIcs, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeInterac, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeMasav, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeMxCcen, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeNeft, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeNics, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeNzBecs, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypePlElixir, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypePolygon, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeProvxchange, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeRoSent, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeRtp, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSen, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSepa, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSgGiro, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSic, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSignet, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSknbi, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeSolana, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeWire, PaymentOrderUpdateReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type PaymentOrderUpdateReconciliationStatus string

const (
	PaymentOrderUpdateReconciliationStatusUnreconciled          PaymentOrderUpdateReconciliationStatus = "unreconciled"
	PaymentOrderUpdateReconciliationStatusTentativelyReconciled PaymentOrderUpdateReconciliationStatus = "tentatively_reconciled"
	PaymentOrderUpdateReconciliationStatusReconciled            PaymentOrderUpdateReconciliationStatus = "reconciled"
)

func (r PaymentOrderUpdateReconciliationStatus) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateReconciliationStatusUnreconciled, PaymentOrderUpdateReconciliationStatusTentativelyReconciled, PaymentOrderUpdateReconciliationStatusReconciled:
		return true
	}
	return false
}

// To cancel a payment order, use `cancelled`. To redraft a returned payment order,
// use `approved`. To undo approval on a denied or approved payment order, use
// `needs_approval`.
type PaymentOrderUpdateStatus string

const (
	PaymentOrderUpdateStatusApproved      PaymentOrderUpdateStatus = "approved"
	PaymentOrderUpdateStatusCancelled     PaymentOrderUpdateStatus = "cancelled"
	PaymentOrderUpdateStatusCompleted     PaymentOrderUpdateStatus = "completed"
	PaymentOrderUpdateStatusDenied        PaymentOrderUpdateStatus = "denied"
	PaymentOrderUpdateStatusFailed        PaymentOrderUpdateStatus = "failed"
	PaymentOrderUpdateStatusHeld          PaymentOrderUpdateStatus = "held"
	PaymentOrderUpdateStatusNeedsApproval PaymentOrderUpdateStatus = "needs_approval"
	PaymentOrderUpdateStatusPending       PaymentOrderUpdateStatus = "pending"
	PaymentOrderUpdateStatusProcessing    PaymentOrderUpdateStatus = "processing"
	PaymentOrderUpdateStatusReturned      PaymentOrderUpdateStatus = "returned"
	PaymentOrderUpdateStatusReversed      PaymentOrderUpdateStatus = "reversed"
	PaymentOrderUpdateStatusSent          PaymentOrderUpdateStatus = "sent"
	PaymentOrderUpdateStatusStopped       PaymentOrderUpdateStatus = "stopped"
)

func (r PaymentOrderUpdateStatus) IsKnown() bool {
	switch r {
	case PaymentOrderUpdateStatusApproved, PaymentOrderUpdateStatusCancelled, PaymentOrderUpdateStatusCompleted, PaymentOrderUpdateStatusDenied, PaymentOrderUpdateStatusFailed, PaymentOrderUpdateStatusHeld, PaymentOrderUpdateStatusNeedsApproval, PaymentOrderUpdateStatusPending, PaymentOrderUpdateStatusProcessing, PaymentOrderUpdateStatusReturned, PaymentOrderUpdateStatusReversed, PaymentOrderUpdateStatusSent, PaymentOrderUpdateStatusStopped:
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
	Type       param.Field[PaymentOrderType] `json:"type,required"`
	Accounting param.Field[AccountingParam]  `json:"accounting"`
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
	Documents param.Field[[]DocumentCreateParam] `json:"documents"`
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
	LineItems param.Field[[]LineItemParam] `json:"line_items"`
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
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus param.Field[PaymentOrderNewParamsReconciliationStatus] `json:"reconciliation_status"`
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

func (r PaymentOrderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type PaymentOrderNewParamsReconciliationStatus string

const (
	PaymentOrderNewParamsReconciliationStatusUnreconciled          PaymentOrderNewParamsReconciliationStatus = "unreconciled"
	PaymentOrderNewParamsReconciliationStatusTentativelyReconciled PaymentOrderNewParamsReconciliationStatus = "tentatively_reconciled"
	PaymentOrderNewParamsReconciliationStatusReconciled            PaymentOrderNewParamsReconciliationStatus = "reconciled"
)

func (r PaymentOrderNewParamsReconciliationStatus) IsKnown() bool {
	switch r {
	case PaymentOrderNewParamsReconciliationStatusUnreconciled, PaymentOrderNewParamsReconciliationStatusTentativelyReconciled, PaymentOrderNewParamsReconciliationStatusReconciled:
		return true
	}
	return false
}

type PaymentOrderUpdateParams struct {
	PaymentOrderUpdate PaymentOrderUpdateParam `json:"payment_order_update"`
}

func (r PaymentOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PaymentOrderUpdate)
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
	PaymentOrderAsyncCreate PaymentOrderAsyncCreateParam `json:"payment_order_async_create,required"`
}

func (r PaymentOrderNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PaymentOrderAsyncCreate)
}
