// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// BulkRequestService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkRequestService] method instead.
type BulkRequestService struct {
	Options []option.RequestOption
}

// NewBulkRequestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBulkRequestService(opts ...option.RequestOption) (r *BulkRequestService) {
	r = &BulkRequestService{}
	r.Options = opts
	return
}

// create bulk_request
func (r *BulkRequestService) New(ctx context.Context, body BulkRequestNewParams, opts ...option.RequestOption) (res *BulkRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/bulk_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get bulk_request
func (r *BulkRequestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BulkRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/bulk_requests/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// list bulk_requests
func (r *BulkRequestService) List(ctx context.Context, query BulkRequestListParams, opts ...option.RequestOption) (res *pagination.Page[BulkRequest], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/bulk_requests"
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

// list bulk_requests
func (r *BulkRequestService) ListAutoPaging(ctx context.Context, query BulkRequestListParams, opts ...option.RequestOption) *pagination.PageAutoPager[BulkRequest] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type BulkRequest struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// One of create, or update.
	ActionType BulkRequestActionType `json:"action_type" api:"required"`
	CreatedAt  time.Time             `json:"created_at" api:"required" format:"date-time"`
	// Total number of failed bulk results so far for this request
	FailedResultCount int64 `json:"failed_result_count" api:"required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode" api:"required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata" api:"required"`
	Object   string            `json:"object" api:"required"`
	// One of payment_order, expected_payment, or ledger_transaction.
	ResourceType BulkRequestResourceType `json:"resource_type" api:"required"`
	// One of pending, processing, or completed.
	Status BulkRequestStatus `json:"status" api:"required"`
	// Total number of successful bulk results so far for this request
	SuccessResultCount int64 `json:"success_result_count" api:"required"`
	// Total number of items in the `resources` array. Once a bulk request is
	// completed, `success_result_count` + `failed_result_count` will be equal to
	// `total_result_count`.
	TotalResourceCount int64           `json:"total_resource_count" api:"required"`
	UpdatedAt          time.Time       `json:"updated_at" api:"required" format:"date-time"`
	JSON               bulkRequestJSON `json:"-"`
}

// bulkRequestJSON contains the JSON metadata for the struct [BulkRequest]
type bulkRequestJSON struct {
	ID                 apijson.Field
	ActionType         apijson.Field
	CreatedAt          apijson.Field
	FailedResultCount  apijson.Field
	LiveMode           apijson.Field
	Metadata           apijson.Field
	Object             apijson.Field
	ResourceType       apijson.Field
	Status             apijson.Field
	SuccessResultCount apijson.Field
	TotalResourceCount apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BulkRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkRequestJSON) RawJSON() string {
	return r.raw
}

// One of create, or update.
type BulkRequestActionType string

const (
	BulkRequestActionTypeCreate BulkRequestActionType = "create"
	BulkRequestActionTypeUpdate BulkRequestActionType = "update"
	BulkRequestActionTypeDelete BulkRequestActionType = "delete"
)

func (r BulkRequestActionType) IsKnown() bool {
	switch r {
	case BulkRequestActionTypeCreate, BulkRequestActionTypeUpdate, BulkRequestActionTypeDelete:
		return true
	}
	return false
}

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestResourceType string

const (
	BulkRequestResourceTypePaymentOrder        BulkRequestResourceType = "payment_order"
	BulkRequestResourceTypeLedgerAccount       BulkRequestResourceType = "ledger_account"
	BulkRequestResourceTypeLedgerTransaction   BulkRequestResourceType = "ledger_transaction"
	BulkRequestResourceTypeExpectedPayment     BulkRequestResourceType = "expected_payment"
	BulkRequestResourceTypeTransaction         BulkRequestResourceType = "transaction"
	BulkRequestResourceTypeTransactionLineItem BulkRequestResourceType = "transaction_line_item"
	BulkRequestResourceTypeEntityLink          BulkRequestResourceType = "entity_link"
)

func (r BulkRequestResourceType) IsKnown() bool {
	switch r {
	case BulkRequestResourceTypePaymentOrder, BulkRequestResourceTypeLedgerAccount, BulkRequestResourceTypeLedgerTransaction, BulkRequestResourceTypeExpectedPayment, BulkRequestResourceTypeTransaction, BulkRequestResourceTypeTransactionLineItem, BulkRequestResourceTypeEntityLink:
		return true
	}
	return false
}

// One of pending, processing, or completed.
type BulkRequestStatus string

const (
	BulkRequestStatusPending    BulkRequestStatus = "pending"
	BulkRequestStatusProcessing BulkRequestStatus = "processing"
	BulkRequestStatusCompleted  BulkRequestStatus = "completed"
)

func (r BulkRequestStatus) IsKnown() bool {
	switch r {
	case BulkRequestStatusPending, BulkRequestStatusProcessing, BulkRequestStatusCompleted:
		return true
	}
	return false
}

type BulkRequestNewParams struct {
	// One of create, or update.
	ActionType param.Field[BulkRequestNewParamsActionType] `json:"action_type" api:"required"`
	// One of payment_order, expected_payment, or ledger_transaction.
	ResourceType param.Field[BulkRequestNewParamsResourceType] `json:"resource_type" api:"required"`
	// An array of objects where each object contains the input params for a single
	// `action_type` request on a `resource_type` resource
	Resources param.Field[[]BulkRequestNewParamsResourceUnion] `json:"resources" api:"required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of create, or update.
type BulkRequestNewParamsActionType string

const (
	BulkRequestNewParamsActionTypeCreate BulkRequestNewParamsActionType = "create"
	BulkRequestNewParamsActionTypeUpdate BulkRequestNewParamsActionType = "update"
	BulkRequestNewParamsActionTypeDelete BulkRequestNewParamsActionType = "delete"
)

func (r BulkRequestNewParamsActionType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsActionTypeCreate, BulkRequestNewParamsActionTypeUpdate, BulkRequestNewParamsActionTypeDelete:
		return true
	}
	return false
}

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestNewParamsResourceType string

const (
	BulkRequestNewParamsResourceTypePaymentOrder        BulkRequestNewParamsResourceType = "payment_order"
	BulkRequestNewParamsResourceTypeLedgerAccount       BulkRequestNewParamsResourceType = "ledger_account"
	BulkRequestNewParamsResourceTypeLedgerTransaction   BulkRequestNewParamsResourceType = "ledger_transaction"
	BulkRequestNewParamsResourceTypeExpectedPayment     BulkRequestNewParamsResourceType = "expected_payment"
	BulkRequestNewParamsResourceTypeTransaction         BulkRequestNewParamsResourceType = "transaction"
	BulkRequestNewParamsResourceTypeTransactionLineItem BulkRequestNewParamsResourceType = "transaction_line_item"
	BulkRequestNewParamsResourceTypeEntityLink          BulkRequestNewParamsResourceType = "entity_link"
)

func (r BulkRequestNewParamsResourceType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourceTypePaymentOrder, BulkRequestNewParamsResourceTypeLedgerAccount, BulkRequestNewParamsResourceTypeLedgerTransaction, BulkRequestNewParamsResourceTypeExpectedPayment, BulkRequestNewParamsResourceTypeTransaction, BulkRequestNewParamsResourceTypeTransactionLineItem, BulkRequestNewParamsResourceTypeEntityLink:
		return true
	}
	return false
}

type BulkRequestNewParamsResource struct {
	ID         param.Field[string]      `json:"id" format:"uuid"`
	Accounting param.Field[interface{}] `json:"accounting"`
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
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound"`
	// The amount reconciled for this expected payment. Value in specified currency's
	// smallest unit. e.g. $10 would be represented as 1000.
	AmountReconciled param.Field[int64] `json:"amount_reconciled"`
	// One of credit or debit. Indicates whether amount_reconciled is a credit or debit
	// amount.
	AmountReconciledDirection param.Field[BulkRequestNewParamsResourcesAmountReconciledDirection] `json:"amount_reconciled_direction"`
	// The amount that remains unreconciled for this expected payment. Value in
	// specified currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUnreconciled param.Field[int64] `json:"amount_unreconciled"`
	// One of credit or debit. Indicates whether amount_unreconciled is a credit or
	// debit amount.
	AmountUnreconciledDirection param.Field[BulkRequestNewParamsResourcesAmountUnreconciledDirection] `json:"amount_unreconciled_direction"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound"`
	// The date on which the transaction occurred.
	AsOfDate param.Field[time.Time] `json:"as_of_date" format:"date"`
	// The party that will pay the fees for the payment order. See
	// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
	// differences between the options.
	ChargeBearer param.Field[BulkRequestNewParamsResourcesChargeBearer] `json:"charge_bearer"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound" format:"date"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound" format:"date"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[string] `json:"direction"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[BulkRequestNewParamsResourcesFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[BulkRequestNewParamsResourcesForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID        param.Field[string]      `json:"internal_account_id" format:"uuid"`
	LedgerAccountCategoryIDs param.Field[interface{}] `json:"ledger_account_category_ids"`
	LedgerEntries            param.Field[interface{}] `json:"ledger_entries"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id" format:"uuid"`
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
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesLedgerableType] `json:"ledgerable_type"`
	LineItems      param.Field[interface{}]                                 `json:"line_items"`
	Metadata       param.Field[interface{}]                                 `json:"metadata"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[shared.TransactionDirection] `json:"normal_balance"`
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
	// This field will be `true` if the transaction has posted to the account.
	Posted param.Field[bool] `json:"posted"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[BulkRequestNewParamsResourcesPriority] `json:"priority"`
	// If present, Modern Treasury will not process the payment until after this time.
	// If `process_after` is past the cutoff for `effective_date`, `process_after` will
	// take precedence and `effective_date` will automatically update to reflect the
	// earliest possible sending date after `process_after`. Format is ISO8601
	// timestamp.
	ProcessAfter param.Field[time.Time] `json:"process_after" format:"date-time"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. For `eft`, this field is the 3
	// digit CPA Code that will be attached to the payment.
	Purpose          param.Field[string]      `json:"purpose"`
	ReceivingAccount param.Field[interface{}] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID          param.Field[string]      `json:"receiving_account_id" format:"uuid"`
	ReconciliationFilters       param.Field[interface{}] `json:"reconciliation_filters"`
	ReconciliationGroups        param.Field[interface{}] `json:"reconciliation_groups"`
	ReconciliationRuleVariables param.Field[interface{}] `json:"reconciliation_rule_variables"`
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus param.Field[BulkRequestNewParamsResourcesReconciliationStatus] `json:"reconciliation_status"`
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
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesStatus] `json:"status"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype param.Field[PaymentOrderSubtype] `json:"subtype"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	//
	// Deprecated: deprecated
	TransactionMonitoringEnabled param.Field[bool] `json:"transaction_monitoring_enabled"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type param.Field[PaymentOrderType] `json:"type"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier param.Field[string] `json:"ultimate_originating_party_identifier"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName param.Field[string] `json:"ultimate_originating_party_name"`
	// Identifier of the ultimate funds recipient.
	UltimateReceivingPartyIdentifier param.Field[string] `json:"ultimate_receiving_party_identifier"`
	// Name of the ultimate funds recipient.
	UltimateReceivingPartyName param.Field[string] `json:"ultimate_receiving_party_name"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode param.Field[string] `json:"vendor_code"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, `us_bank`, or others.
	VendorCodeType param.Field[string] `json:"vendor_code_type"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID param.Field[string] `json:"vendor_customer_id"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription param.Field[string] `json:"vendor_description"`
}

func (r BulkRequestNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResource) ImplementsBulkRequestNewParamsResourceUnion() {}

// Satisfied by [BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest],
// [BulkRequestNewParamsResourcesExpectedPaymentCreateRequest],
// [shared.LedgerTransactionCreateRequestParam],
// [shared.LedgerAccountCreateRequestParam],
// [BulkRequestNewParamsResourcesTransactionCreateRequest],
// [BulkRequestNewParamsResourcesID],
// [BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID],
// [BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID],
// [BulkRequestNewParamsResourcesTransactionUpdateRequestWithID],
// [BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID],
// [BulkRequestNewParamsResource].
type BulkRequestNewParamsResourceUnion interface {
	ImplementsBulkRequestNewParamsResourceUnion()
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirection] `json:"direction" api:"required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id" api:"required" format:"uuid"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type param.Field[PaymentOrderType] `json:"type" api:"required"`
	// Deprecated: deprecated
	Accounting param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting] `json:"accounting"`
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
	ChargeBearer param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearer] `json:"charge_bearer"`
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
	FallbackType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
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
	LineItems param.Field[[]BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem] `json:"line_items"`
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
	Priority param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriority] `json:"priority"`
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
	ReceivingAccount param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatus] `json:"reconciliation_status"`
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
	//
	// Deprecated: deprecated
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

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest) ImplementsBulkRequestNewParamsResourceUnion() {
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirection string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionCredit BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirection = "credit"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionDebit  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirection = "debit"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionCredit, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionDebit:
		return true
	}
	return false
}

// Deprecated: deprecated
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting struct {
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

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearer string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerShared   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearer = "shared"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerSender   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearer = "sender"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerReceiver BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearer = "receiver"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearer) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerShared, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerSender, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerReceiver:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackTypeACH BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackType = "ach"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicator string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorFixedToVariable BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicator = "fixed_to_variable"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorVariableToFixed BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicator = "variable_to_fixed"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorFixedToVariable, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriority string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityHigh   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriority = "high"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityNormal BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriority = "normal"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriority) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityHigh, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount struct {
	AccountDetails param.Field[[]BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail] `json:"account_details"`
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
	PartyType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                                                     `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                                                                     `json:"account_number" api:"required"`
	AccountNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "au_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "base_address"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeCardToken       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "card_token"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeClabe           BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "clabe"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeHkNumber        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban            BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "iban"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIDNumber        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "id_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeNzNumber        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeOther           BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "other"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypePan             BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "pan"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeSgNumber        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeCardToken, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeClabe, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeHkNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIDNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeNzNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeOther, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypePan, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeSgNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeBusiness   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyType = "business"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeIndividual BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyType = "individual"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeBusiness, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                                                                     `json:"routing_number" api:"required"`
	RoutingNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type" api:"required"`
	PaymentType       param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba                     BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeChips                   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeChips, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSwift, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "ach"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeAuBecs      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "au_becs"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBacs        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "bacs"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBase        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "base"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBook        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "book"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCard        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "card"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeChats       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "chats"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCheck       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "check"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCrossBorder BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "cross_border"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeDkNets      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeEft         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "eft"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeEthereum    BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "ethereum"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeGBFps       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeHuIcs       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeInterac     BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "interac"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMasav       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "masav"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMxCcen      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNeft        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "neft"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNics        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "nics"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNzBecs      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypePlElixir    BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypePolygon     BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "polygon"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeProvxchange BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "provxchange"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeRoSent      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeRtp         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "rtp"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSeBankgirot BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSen         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "sen"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSepa        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "sepa"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSgGiro      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSic         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "sic"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSignet      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "signet"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSknbi       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "sknbi"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSolana      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "solana"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeWire        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "wire"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeZengin      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeAuBecs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBacs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBase, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBook, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCard, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeChats, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCheck, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCrossBorder, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeDkNets, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeEft, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeEthereum, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeGBFps, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeHuIcs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeInterac, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMasav, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMxCcen, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNeft, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNics, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNzBecs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypePlElixir, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypePolygon, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeProvxchange, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeRoSent, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeRtp, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSen, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSepa, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSgGiro, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSic, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSignet, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSknbi, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSolana, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeWire, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatus string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusUnreconciled          BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatus = "unreconciled"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusTentativelyReconciled BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatus = "tentatively_reconciled"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusReconciled            BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatus = "reconciled"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusUnreconciled, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusTentativelyReconciled, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusReconciled:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesExpectedPaymentCreateRequest struct {
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound"`
	// The amount reconciled for this expected payment. Value in specified currency's
	// smallest unit. e.g. $10 would be represented as 1000.
	AmountReconciled param.Field[int64] `json:"amount_reconciled"`
	// One of credit or debit. Indicates whether amount_reconciled is a credit or debit
	// amount.
	AmountReconciledDirection param.Field[BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirection] `json:"amount_reconciled_direction"`
	// The amount that remains unreconciled for this expected payment. Value in
	// specified currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUnreconciled param.Field[int64] `json:"amount_unreconciled"`
	// One of credit or debit. Indicates whether amount_unreconciled is a credit or
	// debit amount.
	AmountUnreconciledDirection param.Field[BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirection] `json:"amount_unreconciled_direction"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound" format:"date"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound" format:"date"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction param.Field[BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirection] `json:"direction"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID param.Field[string] `json:"internal_account_id" format:"uuid"`
	// Specifies a ledger transaction object that will be created with the expected
	// payment. If the ledger transaction cannot be created, then the expected payment
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the expected payment.
	LedgerTransaction param.Field[shared.LedgerTransactionCreateRequestParam] `json:"ledger_transaction"`
	// Either ledger_transaction or ledger_transaction_id can be provided. Only a
	// pending ledger transaction can be attached upon expected payment creation. Once
	// the expected payment is created, the status of the ledger transaction tracks the
	// expected payment automatically.
	LedgerTransactionID param.Field[string]                                                              `json:"ledger_transaction_id" format:"uuid"`
	LineItems           param.Field[[]BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLineItem] `json:"line_items"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The reconciliation filters you have for this payment.
	ReconciliationFilters param.Field[interface{}] `json:"reconciliation_filters"`
	// The reconciliation groups you have for this payment.
	ReconciliationGroups param.Field[interface{}] `json:"reconciliation_groups"`
	// An array of reconciliation rule variables for this payment.
	ReconciliationRuleVariables param.Field[[]ReconciliationRuleParam] `json:"reconciliation_rule_variables"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// The statement description you expect to see on the transaction. For ACH
	// payments, this will be the full line item passed from the bank. For wire
	// payments, this will be the OBI field on the wire. For check payments, this will
	// be the memo field.
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type param.Field[ExpectedPaymentType] `json:"type"`
}

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequest) ImplementsBulkRequestNewParamsResourceUnion() {
}

// One of credit or debit. Indicates whether amount_reconciled is a credit or debit
// amount.
type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirection string

const (
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirectionCredit BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirection = "credit"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirectionDebit  BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirection = "debit"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirectionCredit, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountReconciledDirectionDebit:
		return true
	}
	return false
}

// One of credit or debit. Indicates whether amount_unreconciled is a credit or
// debit amount.
type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirection string

const (
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirectionCredit BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirection = "credit"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirectionDebit  BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirection = "debit"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirectionCredit, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestAmountUnreconciledDirectionDebit:
		return true
	}
	return false
}

// One of credit or debit. When you are receiving money, use credit. When you are
// being charged, use debit.
type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirection string

const (
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirectionCredit BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirection = "credit"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirectionDebit  BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirection = "debit"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirectionCredit, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestDirectionDebit:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLineItem struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesTransactionCreateRequest struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The date on which the transaction occurred.
	AsOfDate param.Field[time.Time] `json:"as_of_date" api:"required" format:"date"`
	// Either `credit` or `debit`.
	Direction param.Field[string] `json:"direction" api:"required"`
	// The ID of the relevant Internal Account.
	InternalAccountID param.Field[string] `json:"internal_account_id" api:"required" format:"uuid"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode param.Field[string] `json:"vendor_code" api:"required"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, `us_bank`, or others.
	VendorCodeType param.Field[string] `json:"vendor_code_type" api:"required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// This field will be `true` if the transaction has posted to the account.
	Posted param.Field[bool] `json:"posted"`
	// The type of the transaction. Examples could be
	// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
	Type param.Field[BulkRequestNewParamsResourcesTransactionCreateRequestType] `json:"type"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID param.Field[string] `json:"vendor_customer_id"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription param.Field[string] `json:"vendor_description"`
}

func (r BulkRequestNewParamsResourcesTransactionCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesTransactionCreateRequest) ImplementsBulkRequestNewParamsResourceUnion() {
}

// The type of the transaction. Examples could be
// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
type BulkRequestNewParamsResourcesTransactionCreateRequestType string

const (
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeACH         BulkRequestNewParamsResourcesTransactionCreateRequestType = "ach"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeAuBecs      BulkRequestNewParamsResourcesTransactionCreateRequestType = "au_becs"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeBacs        BulkRequestNewParamsResourcesTransactionCreateRequestType = "bacs"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeBase        BulkRequestNewParamsResourcesTransactionCreateRequestType = "base"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeBook        BulkRequestNewParamsResourcesTransactionCreateRequestType = "book"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeCard        BulkRequestNewParamsResourcesTransactionCreateRequestType = "card"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeChats       BulkRequestNewParamsResourcesTransactionCreateRequestType = "chats"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeCheck       BulkRequestNewParamsResourcesTransactionCreateRequestType = "check"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeCrossBorder BulkRequestNewParamsResourcesTransactionCreateRequestType = "cross_border"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeDkNets      BulkRequestNewParamsResourcesTransactionCreateRequestType = "dk_nets"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeEft         BulkRequestNewParamsResourcesTransactionCreateRequestType = "eft"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeEthereum    BulkRequestNewParamsResourcesTransactionCreateRequestType = "ethereum"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeGBFps       BulkRequestNewParamsResourcesTransactionCreateRequestType = "gb_fps"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeHuIcs       BulkRequestNewParamsResourcesTransactionCreateRequestType = "hu_ics"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeInterac     BulkRequestNewParamsResourcesTransactionCreateRequestType = "interac"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeMasav       BulkRequestNewParamsResourcesTransactionCreateRequestType = "masav"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeMxCcen      BulkRequestNewParamsResourcesTransactionCreateRequestType = "mx_ccen"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeNeft        BulkRequestNewParamsResourcesTransactionCreateRequestType = "neft"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeNics        BulkRequestNewParamsResourcesTransactionCreateRequestType = "nics"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeNzBecs      BulkRequestNewParamsResourcesTransactionCreateRequestType = "nz_becs"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypePlElixir    BulkRequestNewParamsResourcesTransactionCreateRequestType = "pl_elixir"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypePolygon     BulkRequestNewParamsResourcesTransactionCreateRequestType = "polygon"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeProvxchange BulkRequestNewParamsResourcesTransactionCreateRequestType = "provxchange"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeRoSent      BulkRequestNewParamsResourcesTransactionCreateRequestType = "ro_sent"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeRtp         BulkRequestNewParamsResourcesTransactionCreateRequestType = "rtp"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSeBankgirot BulkRequestNewParamsResourcesTransactionCreateRequestType = "se_bankgirot"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSen         BulkRequestNewParamsResourcesTransactionCreateRequestType = "sen"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSepa        BulkRequestNewParamsResourcesTransactionCreateRequestType = "sepa"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSgGiro      BulkRequestNewParamsResourcesTransactionCreateRequestType = "sg_giro"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSic         BulkRequestNewParamsResourcesTransactionCreateRequestType = "sic"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSignet      BulkRequestNewParamsResourcesTransactionCreateRequestType = "signet"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSknbi       BulkRequestNewParamsResourcesTransactionCreateRequestType = "sknbi"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeSolana      BulkRequestNewParamsResourcesTransactionCreateRequestType = "solana"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeWire        BulkRequestNewParamsResourcesTransactionCreateRequestType = "wire"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeZengin      BulkRequestNewParamsResourcesTransactionCreateRequestType = "zengin"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeOther       BulkRequestNewParamsResourcesTransactionCreateRequestType = "other"
)

func (r BulkRequestNewParamsResourcesTransactionCreateRequestType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesTransactionCreateRequestTypeACH, BulkRequestNewParamsResourcesTransactionCreateRequestTypeAuBecs, BulkRequestNewParamsResourcesTransactionCreateRequestTypeBacs, BulkRequestNewParamsResourcesTransactionCreateRequestTypeBase, BulkRequestNewParamsResourcesTransactionCreateRequestTypeBook, BulkRequestNewParamsResourcesTransactionCreateRequestTypeCard, BulkRequestNewParamsResourcesTransactionCreateRequestTypeChats, BulkRequestNewParamsResourcesTransactionCreateRequestTypeCheck, BulkRequestNewParamsResourcesTransactionCreateRequestTypeCrossBorder, BulkRequestNewParamsResourcesTransactionCreateRequestTypeDkNets, BulkRequestNewParamsResourcesTransactionCreateRequestTypeEft, BulkRequestNewParamsResourcesTransactionCreateRequestTypeEthereum, BulkRequestNewParamsResourcesTransactionCreateRequestTypeGBFps, BulkRequestNewParamsResourcesTransactionCreateRequestTypeHuIcs, BulkRequestNewParamsResourcesTransactionCreateRequestTypeInterac, BulkRequestNewParamsResourcesTransactionCreateRequestTypeMasav, BulkRequestNewParamsResourcesTransactionCreateRequestTypeMxCcen, BulkRequestNewParamsResourcesTransactionCreateRequestTypeNeft, BulkRequestNewParamsResourcesTransactionCreateRequestTypeNics, BulkRequestNewParamsResourcesTransactionCreateRequestTypeNzBecs, BulkRequestNewParamsResourcesTransactionCreateRequestTypePlElixir, BulkRequestNewParamsResourcesTransactionCreateRequestTypePolygon, BulkRequestNewParamsResourcesTransactionCreateRequestTypeProvxchange, BulkRequestNewParamsResourcesTransactionCreateRequestTypeRoSent, BulkRequestNewParamsResourcesTransactionCreateRequestTypeRtp, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSeBankgirot, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSen, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSepa, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSgGiro, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSic, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSignet, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSknbi, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSolana, BulkRequestNewParamsResourcesTransactionCreateRequestTypeWire, BulkRequestNewParamsResourcesTransactionCreateRequestTypeZengin, BulkRequestNewParamsResourcesTransactionCreateRequestTypeOther:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
}

func (r BulkRequestNewParamsResourcesID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesID) ImplementsBulkRequestNewParamsResourceUnion() {}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// Deprecated: deprecated
	Accounting param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDAccounting] `json:"accounting"`
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
	ChargeBearer param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearer] `json:"charge_bearer"`
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
	Direction param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirection] `json:"direction"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDLineItem] `json:"line_items"`
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
	Priority param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriority] `json:"priority"`
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
	ReceivingAccount param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccount] `json:"receiving_account"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
	ReconciliationStatus param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatus] `json:"reconciliation_status"`
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
	Status param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus] `json:"status"`
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

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID) ImplementsBulkRequestNewParamsResourceUnion() {
}

// Deprecated: deprecated
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDAccounting struct {
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

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearer string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearerShared   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearer = "shared"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearerSender   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearer = "sender"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearerReceiver BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearer = "receiver"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearer) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearerShared, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearerSender, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDChargeBearerReceiver:
		return true
	}
	return false
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirection string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirectionCredit BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirection = "credit"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirectionDebit  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirection = "debit"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirectionCredit, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDDirectionDebit:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDFallbackType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDFallbackTypeACH BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDFallbackType = "ach"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDFallbackType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicator string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicatorFixedToVariable BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicator = "fixed_to_variable"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicatorVariableToFixed BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicator = "variable_to_fixed"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicatorFixedToVariable, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDLineItem struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount" api:"required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriority string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriorityHigh   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriority = "high"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriorityNormal BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriority = "normal"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriority) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriorityHigh, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDPriorityNormal:
		return true
	}
	return false
}

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccount struct {
	AccountDetails param.Field[[]BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetail] `json:"account_details"`
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
	PartyType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                                                      `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                                                                      `json:"account_number" api:"required"`
	AccountNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeAuNumber        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "au_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeBaseAddress     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "base_address"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeCardToken       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "card_token"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeClabe           BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "clabe"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "ethereum_address"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeHkNumber        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIban            BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "iban"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIDNumber        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "id_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeNzNumber        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeOther           BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "other"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypePan             BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "pan"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypePolygonAddress  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "polygon_address"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeSgNumber        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "solana_address"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeWalletAddress   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeAuNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeBaseAddress, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeCardToken, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeClabe, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeEthereumAddress, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeHkNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIban, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIDNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeNzNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeOther, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypePan, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypePolygonAddress, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeSgNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeSolanaAddress, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyTypeBusiness   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyType = "business"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyTypeIndividual BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyType = "individual"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyTypeBusiness, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyTypeIndividual:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                                                                      `json:"routing_number" api:"required"`
	RoutingNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type" api:"required"`
	PaymentType       param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeAba                     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeChips                   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode             BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "id_sknbi_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode              BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "il_bank_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeAba, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeChips, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeIlBankCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSwift, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeACH         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "ach"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeAuBecs      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "au_becs"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBacs        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "bacs"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBase        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "base"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBook        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "book"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCard        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "card"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeChats       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "chats"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCheck       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "check"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCrossBorder BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "cross_border"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeDkNets      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeEft         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "eft"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeEthereum    BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "ethereum"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeGBFps       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "gb_fps"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeHuIcs       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeInterac     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "interac"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMasav       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "masav"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMxCcen      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNeft        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "neft"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNics        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "nics"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNzBecs      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypePlElixir    BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypePolygon     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "polygon"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeProvxchange BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "provxchange"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeRoSent      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "ro_sent"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeRtp         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "rtp"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSeBankgirot BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSen         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "sen"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSepa        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "sepa"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSgGiro      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSic         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "sic"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSignet      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "signet"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSknbi       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "sknbi"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSolana      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "solana"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeWire        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "wire"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeZengin      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeACH, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeAuBecs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBacs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBase, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBook, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCard, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeChats, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCheck, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCrossBorder, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeDkNets, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeEft, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeEthereum, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeGBFps, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeHuIcs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeInterac, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMasav, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMxCcen, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNeft, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNics, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNzBecs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypePlElixir, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypePolygon, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeProvxchange, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeRoSent, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeRtp, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSen, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSepa, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSgGiro, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSic, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSignet, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSknbi, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSolana, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeWire, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatus string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatusUnreconciled          BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatus = "unreconciled"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatusTentativelyReconciled BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatus = "tentatively_reconciled"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatusReconciled            BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatus = "reconciled"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatusUnreconciled, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatusTentativelyReconciled, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReconciliationStatusReconciled:
		return true
	}
	return false
}

// To cancel a payment order, use `cancelled`. To redraft a returned payment order,
// use `approved`. To undo approval on a denied or approved payment order, use
// `needs_approval`.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusApproved      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "approved"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusCancelled     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "cancelled"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusCompleted     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "completed"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusDenied        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "denied"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusFailed        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "failed"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusHeld          BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "held"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusNeedsApproval BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "needs_approval"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusPending       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "pending"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusProcessing    BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "processing"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReturned      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "returned"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReversed      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "reversed"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusSent          BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "sent"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusStopped       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "stopped"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusApproved, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusCancelled, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusCompleted, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusDenied, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusFailed, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusHeld, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusNeedsApproval, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusPending, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusProcessing, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReturned, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReversed, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusSent, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusStopped:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound"`
	// The amount reconciled for this expected payment. Value in specified currency's
	// smallest unit. e.g. $10 would be represented as 1000.
	AmountReconciled param.Field[int64] `json:"amount_reconciled"`
	// One of credit or debit. Indicates whether amount_reconciled is a credit or debit
	// amount.
	AmountReconciledDirection param.Field[BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirection] `json:"amount_reconciled_direction"`
	// The amount that remains unreconciled for this expected payment. Value in
	// specified currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUnreconciled param.Field[int64] `json:"amount_unreconciled"`
	// One of credit or debit. Indicates whether amount_unreconciled is a credit or
	// debit amount.
	AmountUnreconciledDirection param.Field[BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirection] `json:"amount_unreconciled_direction"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound" format:"date"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound" format:"date"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction param.Field[BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirection] `json:"direction"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID param.Field[string] `json:"internal_account_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The reconciliation filters you have for this payment.
	ReconciliationFilters param.Field[interface{}] `json:"reconciliation_filters"`
	// The reconciliation groups you have for this payment.
	ReconciliationGroups param.Field[interface{}] `json:"reconciliation_groups"`
	// An array of reconciliation rule variables for this payment.
	ReconciliationRuleVariables param.Field[[]ReconciliationRuleParam] `json:"reconciliation_rule_variables"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// The statement description you expect to see on the transaction. For ACH
	// payments, this will be the full line item passed from the bank. For wire
	// payments, this will be the OBI field on the wire. For check payments, this will
	// be the memo field.
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	// The Expected Payment's status can be updated from partially_reconciled to
	// reconciled.
	Status param.Field[BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDStatus] `json:"status"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type param.Field[ExpectedPaymentType] `json:"type"`
}

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID) ImplementsBulkRequestNewParamsResourceUnion() {
}

// One of credit or debit. Indicates whether amount_reconciled is a credit or debit
// amount.
type BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirection string

const (
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirectionCredit BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirection = "credit"
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirectionDebit  BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirection = "debit"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirectionCredit, BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountReconciledDirectionDebit:
		return true
	}
	return false
}

// One of credit or debit. Indicates whether amount_unreconciled is a credit or
// debit amount.
type BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirection string

const (
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirectionCredit BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirection = "credit"
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirectionDebit  BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirection = "debit"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirectionCredit, BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDAmountUnreconciledDirectionDebit:
		return true
	}
	return false
}

// One of credit or debit. When you are receiving money, use credit. When you are
// being charged, use debit.
type BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirection string

const (
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirectionCredit BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirection = "credit"
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirectionDebit  BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirection = "debit"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirectionCredit, BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDDirectionDebit:
		return true
	}
	return false
}

// The Expected Payment's status can be updated from partially_reconciled to
// reconciled.
type BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDStatus string

const (
	BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDStatusReconciled BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDStatus = "reconciled"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithIDStatusReconciled:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesTransactionUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesTransactionUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesTransactionUpdateRequestWithID) ImplementsBulkRequestNewParamsResourceUnion() {
}

type BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]shared.LedgerEntryCreateRequestParam] `json:"ledger_entries"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatus] `json:"status"`
}

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID) ImplementsBulkRequestNewParamsResourceUnion() {
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType string

const (
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReturn                BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "return"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReversal              BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "reversal"
)

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReturn, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatus string

const (
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatusArchived BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatus = "archived"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatusPending  BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatus = "pending"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatusPosted   BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatus = "posted"
)

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatusArchived, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatusPending, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDStatusPosted:
		return true
	}
	return false
}

// One of credit or debit. Indicates whether amount_reconciled is a credit or debit
// amount.
type BulkRequestNewParamsResourcesAmountReconciledDirection string

const (
	BulkRequestNewParamsResourcesAmountReconciledDirectionCredit BulkRequestNewParamsResourcesAmountReconciledDirection = "credit"
	BulkRequestNewParamsResourcesAmountReconciledDirectionDebit  BulkRequestNewParamsResourcesAmountReconciledDirection = "debit"
)

func (r BulkRequestNewParamsResourcesAmountReconciledDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesAmountReconciledDirectionCredit, BulkRequestNewParamsResourcesAmountReconciledDirectionDebit:
		return true
	}
	return false
}

// One of credit or debit. Indicates whether amount_unreconciled is a credit or
// debit amount.
type BulkRequestNewParamsResourcesAmountUnreconciledDirection string

const (
	BulkRequestNewParamsResourcesAmountUnreconciledDirectionCredit BulkRequestNewParamsResourcesAmountUnreconciledDirection = "credit"
	BulkRequestNewParamsResourcesAmountUnreconciledDirectionDebit  BulkRequestNewParamsResourcesAmountUnreconciledDirection = "debit"
)

func (r BulkRequestNewParamsResourcesAmountUnreconciledDirection) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesAmountUnreconciledDirectionCredit, BulkRequestNewParamsResourcesAmountUnreconciledDirectionDebit:
		return true
	}
	return false
}

// The party that will pay the fees for the payment order. See
// https://docs.moderntreasury.com/payments/docs/charge-bearer to understand the
// differences between the options.
type BulkRequestNewParamsResourcesChargeBearer string

const (
	BulkRequestNewParamsResourcesChargeBearerShared   BulkRequestNewParamsResourcesChargeBearer = "shared"
	BulkRequestNewParamsResourcesChargeBearerSender   BulkRequestNewParamsResourcesChargeBearer = "sender"
	BulkRequestNewParamsResourcesChargeBearerReceiver BulkRequestNewParamsResourcesChargeBearer = "receiver"
)

func (r BulkRequestNewParamsResourcesChargeBearer) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesChargeBearerShared, BulkRequestNewParamsResourcesChargeBearerSender, BulkRequestNewParamsResourcesChargeBearerReceiver:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type BulkRequestNewParamsResourcesFallbackType string

const (
	BulkRequestNewParamsResourcesFallbackTypeACH BulkRequestNewParamsResourcesFallbackType = "ach"
)

func (r BulkRequestNewParamsResourcesFallbackType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesFallbackTypeACH:
		return true
	}
	return false
}

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type BulkRequestNewParamsResourcesForeignExchangeIndicator string

const (
	BulkRequestNewParamsResourcesForeignExchangeIndicatorFixedToVariable BulkRequestNewParamsResourcesForeignExchangeIndicator = "fixed_to_variable"
	BulkRequestNewParamsResourcesForeignExchangeIndicatorVariableToFixed BulkRequestNewParamsResourcesForeignExchangeIndicator = "variable_to_fixed"
)

func (r BulkRequestNewParamsResourcesForeignExchangeIndicator) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesForeignExchangeIndicatorFixedToVariable, BulkRequestNewParamsResourcesForeignExchangeIndicatorVariableToFixed:
		return true
	}
	return false
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type BulkRequestNewParamsResourcesLedgerableType string

const (
	BulkRequestNewParamsResourcesLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesLedgerableTypeReturn                BulkRequestNewParamsResourcesLedgerableType = "return"
	BulkRequestNewParamsResourcesLedgerableTypeReversal              BulkRequestNewParamsResourcesLedgerableType = "reversal"
	BulkRequestNewParamsResourcesLedgerableTypeCounterparty          BulkRequestNewParamsResourcesLedgerableType = "counterparty"
	BulkRequestNewParamsResourcesLedgerableTypeExternalAccount       BulkRequestNewParamsResourcesLedgerableType = "external_account"
	BulkRequestNewParamsResourcesLedgerableTypeInternalAccount       BulkRequestNewParamsResourcesLedgerableType = "internal_account"
	BulkRequestNewParamsResourcesLedgerableTypeVirtualAccount        BulkRequestNewParamsResourcesLedgerableType = "virtual_account"
)

func (r BulkRequestNewParamsResourcesLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesLedgerableTypeReturn, BulkRequestNewParamsResourcesLedgerableTypeReversal, BulkRequestNewParamsResourcesLedgerableTypeCounterparty, BulkRequestNewParamsResourcesLedgerableTypeExternalAccount, BulkRequestNewParamsResourcesLedgerableTypeInternalAccount, BulkRequestNewParamsResourcesLedgerableTypeVirtualAccount:
		return true
	}
	return false
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type BulkRequestNewParamsResourcesPriority string

const (
	BulkRequestNewParamsResourcesPriorityHigh   BulkRequestNewParamsResourcesPriority = "high"
	BulkRequestNewParamsResourcesPriorityNormal BulkRequestNewParamsResourcesPriority = "normal"
)

func (r BulkRequestNewParamsResourcesPriority) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPriorityHigh, BulkRequestNewParamsResourcesPriorityNormal:
		return true
	}
	return false
}

// One of `unreconciled`, `tentatively_reconciled` or `reconciled`.
type BulkRequestNewParamsResourcesReconciliationStatus string

const (
	BulkRequestNewParamsResourcesReconciliationStatusUnreconciled          BulkRequestNewParamsResourcesReconciliationStatus = "unreconciled"
	BulkRequestNewParamsResourcesReconciliationStatusTentativelyReconciled BulkRequestNewParamsResourcesReconciliationStatus = "tentatively_reconciled"
	BulkRequestNewParamsResourcesReconciliationStatusReconciled            BulkRequestNewParamsResourcesReconciliationStatus = "reconciled"
)

func (r BulkRequestNewParamsResourcesReconciliationStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesReconciliationStatusUnreconciled, BulkRequestNewParamsResourcesReconciliationStatusTentativelyReconciled, BulkRequestNewParamsResourcesReconciliationStatusReconciled:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesStatus string

const (
	BulkRequestNewParamsResourcesStatusArchived      BulkRequestNewParamsResourcesStatus = "archived"
	BulkRequestNewParamsResourcesStatusPending       BulkRequestNewParamsResourcesStatus = "pending"
	BulkRequestNewParamsResourcesStatusPosted        BulkRequestNewParamsResourcesStatus = "posted"
	BulkRequestNewParamsResourcesStatusApproved      BulkRequestNewParamsResourcesStatus = "approved"
	BulkRequestNewParamsResourcesStatusCancelled     BulkRequestNewParamsResourcesStatus = "cancelled"
	BulkRequestNewParamsResourcesStatusCompleted     BulkRequestNewParamsResourcesStatus = "completed"
	BulkRequestNewParamsResourcesStatusDenied        BulkRequestNewParamsResourcesStatus = "denied"
	BulkRequestNewParamsResourcesStatusFailed        BulkRequestNewParamsResourcesStatus = "failed"
	BulkRequestNewParamsResourcesStatusHeld          BulkRequestNewParamsResourcesStatus = "held"
	BulkRequestNewParamsResourcesStatusNeedsApproval BulkRequestNewParamsResourcesStatus = "needs_approval"
	BulkRequestNewParamsResourcesStatusProcessing    BulkRequestNewParamsResourcesStatus = "processing"
	BulkRequestNewParamsResourcesStatusReturned      BulkRequestNewParamsResourcesStatus = "returned"
	BulkRequestNewParamsResourcesStatusReversed      BulkRequestNewParamsResourcesStatus = "reversed"
	BulkRequestNewParamsResourcesStatusSent          BulkRequestNewParamsResourcesStatus = "sent"
	BulkRequestNewParamsResourcesStatusStopped       BulkRequestNewParamsResourcesStatus = "stopped"
	BulkRequestNewParamsResourcesStatusReconciled    BulkRequestNewParamsResourcesStatus = "reconciled"
)

func (r BulkRequestNewParamsResourcesStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesStatusArchived, BulkRequestNewParamsResourcesStatusPending, BulkRequestNewParamsResourcesStatusPosted, BulkRequestNewParamsResourcesStatusApproved, BulkRequestNewParamsResourcesStatusCancelled, BulkRequestNewParamsResourcesStatusCompleted, BulkRequestNewParamsResourcesStatusDenied, BulkRequestNewParamsResourcesStatusFailed, BulkRequestNewParamsResourcesStatusHeld, BulkRequestNewParamsResourcesStatusNeedsApproval, BulkRequestNewParamsResourcesStatusProcessing, BulkRequestNewParamsResourcesStatusReturned, BulkRequestNewParamsResourcesStatusReversed, BulkRequestNewParamsResourcesStatusSent, BulkRequestNewParamsResourcesStatusStopped, BulkRequestNewParamsResourcesStatusReconciled:
		return true
	}
	return false
}

type BulkRequestListParams struct {
	// One of create, or update.
	ActionType  param.Field[BulkRequestListParamsActionType] `query:"action_type"`
	AfterCursor param.Field[string]                          `query:"after_cursor"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// One of payment_order, expected_payment, or ledger_transaction.
	ResourceType param.Field[BulkRequestListParamsResourceType] `query:"resource_type"`
	// One of pending, processing, or completed.
	Status param.Field[BulkRequestListParamsStatus] `query:"status"`
}

// URLQuery serializes [BulkRequestListParams]'s query parameters as `url.Values`.
func (r BulkRequestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One of create, or update.
type BulkRequestListParamsActionType string

const (
	BulkRequestListParamsActionTypeCreate BulkRequestListParamsActionType = "create"
	BulkRequestListParamsActionTypeUpdate BulkRequestListParamsActionType = "update"
	BulkRequestListParamsActionTypeDelete BulkRequestListParamsActionType = "delete"
)

func (r BulkRequestListParamsActionType) IsKnown() bool {
	switch r {
	case BulkRequestListParamsActionTypeCreate, BulkRequestListParamsActionTypeUpdate, BulkRequestListParamsActionTypeDelete:
		return true
	}
	return false
}

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestListParamsResourceType string

const (
	BulkRequestListParamsResourceTypePaymentOrder        BulkRequestListParamsResourceType = "payment_order"
	BulkRequestListParamsResourceTypeLedgerAccount       BulkRequestListParamsResourceType = "ledger_account"
	BulkRequestListParamsResourceTypeLedgerTransaction   BulkRequestListParamsResourceType = "ledger_transaction"
	BulkRequestListParamsResourceTypeExpectedPayment     BulkRequestListParamsResourceType = "expected_payment"
	BulkRequestListParamsResourceTypeTransaction         BulkRequestListParamsResourceType = "transaction"
	BulkRequestListParamsResourceTypeTransactionLineItem BulkRequestListParamsResourceType = "transaction_line_item"
	BulkRequestListParamsResourceTypeEntityLink          BulkRequestListParamsResourceType = "entity_link"
)

func (r BulkRequestListParamsResourceType) IsKnown() bool {
	switch r {
	case BulkRequestListParamsResourceTypePaymentOrder, BulkRequestListParamsResourceTypeLedgerAccount, BulkRequestListParamsResourceTypeLedgerTransaction, BulkRequestListParamsResourceTypeExpectedPayment, BulkRequestListParamsResourceTypeTransaction, BulkRequestListParamsResourceTypeTransactionLineItem, BulkRequestListParamsResourceTypeEntityLink:
		return true
	}
	return false
}

// One of pending, processing, or completed.
type BulkRequestListParamsStatus string

const (
	BulkRequestListParamsStatusPending    BulkRequestListParamsStatus = "pending"
	BulkRequestListParamsStatusProcessing BulkRequestListParamsStatus = "processing"
	BulkRequestListParamsStatusCompleted  BulkRequestListParamsStatus = "completed"
)

func (r BulkRequestListParamsStatus) IsKnown() bool {
	switch r {
	case BulkRequestListParamsStatusPending, BulkRequestListParamsStatusProcessing, BulkRequestListParamsStatusCompleted:
		return true
	}
	return false
}
