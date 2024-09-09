// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
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
	opts = append(r.Options[:], opts...)
	path := "api/bulk_requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get bulk_request
func (r *BulkRequestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BulkRequest, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	ID string `json:"id,required" format:"uuid"`
	// One of create, or update.
	ActionType BulkRequestActionType `json:"action_type,required"`
	CreatedAt  time.Time             `json:"created_at,required" format:"date-time"`
	// Total number of failed bulk results so far for this request
	FailedResultCount int64 `json:"failed_result_count,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// One of payment_order, expected_payment, or ledger_transaction.
	ResourceType BulkRequestResourceType `json:"resource_type,required"`
	// One of pending, processing, or completed.
	Status BulkRequestStatus `json:"status,required"`
	// Total number of successful bulk results so far for this request
	SuccessResultCount int64 `json:"success_result_count,required"`
	// Total number of items in the `resources` array. Once a bulk request is
	// completed, `success_result_count` + `failed_result_count` will be equal to
	// `total_result_count`.
	TotalResourceCount int64           `json:"total_resource_count,required"`
	UpdatedAt          time.Time       `json:"updated_at,required" format:"date-time"`
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
)

func (r BulkRequestActionType) IsKnown() bool {
	switch r {
	case BulkRequestActionTypeCreate, BulkRequestActionTypeUpdate:
		return true
	}
	return false
}

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestResourceType string

const (
	BulkRequestResourceTypePaymentOrder      BulkRequestResourceType = "payment_order"
	BulkRequestResourceTypeLedgerTransaction BulkRequestResourceType = "ledger_transaction"
	BulkRequestResourceTypeTransaction       BulkRequestResourceType = "transaction"
	BulkRequestResourceTypeExpectedPayment   BulkRequestResourceType = "expected_payment"
)

func (r BulkRequestResourceType) IsKnown() bool {
	switch r {
	case BulkRequestResourceTypePaymentOrder, BulkRequestResourceTypeLedgerTransaction, BulkRequestResourceTypeTransaction, BulkRequestResourceTypeExpectedPayment:
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
	ActionType param.Field[BulkRequestNewParamsActionType] `json:"action_type,required"`
	// One of payment_order, expected_payment, or ledger_transaction.
	ResourceType param.Field[BulkRequestNewParamsResourceType] `json:"resource_type,required"`
	// An array of objects where each object contains the input params for a single
	// `action_type` request on a `resource_type` resource
	Resources param.Field[[]BulkRequestNewParamsResourceUnion] `json:"resources,required"`
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
)

func (r BulkRequestNewParamsActionType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsActionTypeCreate, BulkRequestNewParamsActionTypeUpdate:
		return true
	}
	return false
}

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestNewParamsResourceType string

const (
	BulkRequestNewParamsResourceTypePaymentOrder      BulkRequestNewParamsResourceType = "payment_order"
	BulkRequestNewParamsResourceTypeLedgerTransaction BulkRequestNewParamsResourceType = "ledger_transaction"
	BulkRequestNewParamsResourceTypeTransaction       BulkRequestNewParamsResourceType = "transaction"
	BulkRequestNewParamsResourceTypeExpectedPayment   BulkRequestNewParamsResourceType = "expected_payment"
)

func (r BulkRequestNewParamsResourceType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourceTypePaymentOrder, BulkRequestNewParamsResourceTypeLedgerTransaction, BulkRequestNewParamsResourceTypeTransaction, BulkRequestNewParamsResourceTypeExpectedPayment:
		return true
	}
	return false
}

type BulkRequestNewParamsResource struct {
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type param.Field[PaymentOrderType] `json:"type"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype param.Field[PaymentOrderSubtype] `json:"subtype"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[string] `json:"direction"`
	// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
	// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
	// an overnight check rather than standard mail.
	Priority param.Field[BulkRequestNewParamsResourcesPriority] `json:"priority"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id" format:"uuid"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccountID param.Field[string]      `json:"receiving_account_id" format:"uuid"`
	Accounting         param.Field[interface{}] `json:"accounting,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// An optional descriptor which will appear in the receiver's statement. For
	// `check` payments this field will be used as the memo line. For `ach` the maximum
	// length is 10 characters. Note that for ACH payments, the name on your bank
	// account will be included automatically by the bank, so you can use the
	// characters for other useful information. For `eft` the maximum length is 15
	// characters.
	StatementDescriptor param.Field[string] `json:"statement_descriptor"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information"`
	// If present, Modern Treasury will not process the payment until after this time.
	// If `process_after` is past the cutoff for `effective_date`, `process_after` will
	// take precedence and `effective_date` will automatically update to reflect the
	// earliest possible sending date after `process_after`. Format is ISO8601
	// timestamp.
	ProcessAfter param.Field[time.Time] `json:"process_after" format:"date-time"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. For `eft`, this field is the 3
	// digit CPA Code that will be attached to the payment.
	Purpose  param.Field[string]      `json:"purpose"`
	Metadata param.Field[interface{}] `json:"metadata,required"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer param.Field[BulkRequestNewParamsResourcesChargeBearer] `json:"charge_bearer"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[BulkRequestNewParamsResourcesForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// A boolean to determine if NSF Protection is enabled for this payment order. Note
	// that this setting must also be turned on in your organization settings page.
	NsfProtected param.Field[bool] `json:"nsf_protected"`
	// If present, this will replace your default company name on receiver's bank
	// statement. This field can only be used for ACH payments currently. For ACH, only
	// the first 16 characters of this string will be used. Any additional characters
	// will be truncated.
	OriginatingPartyName param.Field[string] `json:"originating_party_name"`
	// Name of the ultimate originator of the payment order.
	UltimateOriginatingPartyName param.Field[string] `json:"ultimate_originating_party_name"`
	// Identifier of the ultimate originator of the payment order.
	UltimateOriginatingPartyIdentifier param.Field[string] `json:"ultimate_originating_party_identifier"`
	// Name of the ultimate funds recipient.
	UltimateReceivingPartyName param.Field[string] `json:"ultimate_receiving_party_name"`
	// Identifier of the ultimate funds recipient.
	UltimateReceivingPartyIdentifier param.Field[string] `json:"ultimate_receiving_party_identifier"`
	// Send an email to the counterparty when the payment order is sent to the bank. If
	// `null`, `send_remittance_advice` on the Counterparty is used.
	SendRemittanceAdvice param.Field[bool] `json:"send_remittance_advice"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType      param.Field[BulkRequestNewParamsResourcesFallbackType] `json:"fallback_type"`
	ReceivingAccount  param.Field[interface{}]                               `json:"receiving_account,required"`
	LedgerTransaction param.Field[interface{}]                               `json:"ledger_transaction,required"`
	// Either ledger_transaction or ledger_transaction_id can be provided. Only a
	// pending ledger transaction can be attached upon payment order creation. Once the
	// payment order is created, the status of the ledger transaction tracks the
	// payment order automatically.
	LedgerTransactionID param.Field[string]      `json:"ledger_transaction_id" format:"uuid"`
	LineItems           param.Field[interface{}] `json:"line_items,required"`
	// A flag that determines whether a payment order should go through transaction
	// monitoring.
	TransactionMonitoringEnabled param.Field[bool] `json:"transaction_monitoring_enabled"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID param.Field[string] `json:"internal_account_id" format:"uuid"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound" format:"date"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound" format:"date"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID              param.Field[string]      `json:"counterparty_id" format:"uuid"`
	ReconciliationGroups        param.Field[interface{}] `json:"reconciliation_groups,required"`
	ReconciliationFilters       param.Field[interface{}] `json:"reconciliation_filters,required"`
	ReconciliationRuleVariables param.Field[interface{}] `json:"reconciliation_rule_variables,required"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesStatus] `json:"status"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt   param.Field[time.Time]   `json:"effective_at" format:"date-time"`
	LedgerEntries param.Field[interface{}] `json:"ledger_entries,required"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
	// reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesLedgerableType] `json:"ledgerable_type"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription param.Field[string] `json:"vendor_description"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode param.Field[string] `json:"vendor_code"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, `us_bank`, or others.
	VendorCodeType param.Field[string] `json:"vendor_code_type"`
	// The date on which the transaction occurred.
	AsOfDate param.Field[time.Time] `json:"as_of_date" format:"date"`
	// This field will be `true` if the transaction has posted to the account.
	Posted param.Field[bool]   `json:"posted"`
	ID     param.Field[string] `json:"id" format:"uuid"`
}

func (r BulkRequestNewParamsResource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResource) implementsBulkRequestNewParamsResourceUnion() {}

// Satisfied by [BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest],
// [BulkRequestNewParamsResourcesExpectedPaymentCreateRequest],
// [BulkRequestNewParamsResourcesLedgerTransactionCreateRequest],
// [BulkRequestNewParamsResourcesTransactionCreateRequest],
// [BulkRequestNewParamsResourcesID],
// [BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID],
// [BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID],
// [BulkRequestNewParamsResourcesTransactionUpdateRequestWithID],
// [BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID],
// [BulkRequestNewParamsResource].
type BulkRequestNewParamsResourceUnion interface {
	implementsBulkRequestNewParamsResourceUnion()
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirection] `json:"direction,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	Type       param.Field[PaymentOrderType]                                                      `json:"type,required"`
	Accounting param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
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
	LedgerTransaction param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransaction] `json:"ledger_transaction"`
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

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest) implementsBulkRequestNewParamsResourceUnion() {
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

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. Only applies to wire
// payment orders. Can be one of shared, sender, or receiver, which correspond
// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
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

// Specifies a ledger transaction object that will be created with the payment
// order. If the ledger transaction cannot be created, then the payment order
// creation will fail. The resulting ledger transaction will mirror the status of
// the payment order.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransaction struct {
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerEntry] `json:"ledger_entries,required"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
	// reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatus] `json:"status"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerEntry struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
// reversal.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypePaperItem             BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeReturn                BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType = "return"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeReversal              BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType = "reversal"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypePaperItem, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeReturn, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatus string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusArchived BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatus = "archived"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusPending  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatus = "pending"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusPosted   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatus = "posted"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusArchived, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusPending, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusPosted:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem struct {
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
	AccountType    param.Field[ExternalAccountType]                                                                        `json:"account_type"`
	ContactDetails param.Field[[]BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetail] `json:"contact_details"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccount] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyAddress] `json:"party_address"`
	PartyIdentifier param.Field[string]                                                                                  `json:"party_identifier"`
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
	AccountNumber     param.Field[string]                                                                                                     `json:"account_number,required"`
	AccountNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "au_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeClabe         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "clabe"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeHkNumber      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban          BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "iban"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIDNumber      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "id_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeNzNumber      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeOther         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "other"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypePan           BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "pan"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeSgNumber      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeWalletAddress BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeClabe, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeHkNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIDNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeNzNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeOther, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypePan, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeSgNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetail struct {
	ContactIdentifier     param.Field[string]                                                                                                         `json:"contact_identifier"`
	ContactIdentifierType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierType = "email"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypePhoneNumber BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeWebsite     BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierType = "website"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypePhoneNumber, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeWebsite:
		return true
	}
	return false
}

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccount struct {
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[shared.TransactionDirection] `json:"normal_balance,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The array of ledger account category ids that this ledger account should be a
	// child of.
	LedgerAccountCategoryIDs param.Field[[]string] `json:"ledger_account_category_ids" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeCounterparty    BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType = "counterparty"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeExternalAccount BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType = "external_account"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeInternalAccount BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType = "internal_account"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeVirtualAccount  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType = "virtual_account"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeCounterparty, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeExternalAccount, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeInternalAccount, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeVirtualAccount:
		return true
	}
	return false
}

// Required if receiving wire payments.
type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
	Line1   param.Field[string] `json:"line1"`
	Line2   param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Region or State.
	Region param.Field[string] `json:"region"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	RoutingNumber     param.Field[string]                                                                                                     `json:"routing_number,required"`
	RoutingNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
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
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeChips, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSwift, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType string

const (
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "ach"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeAuBecs      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "au_becs"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBacs        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "bacs"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBook        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "book"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCard        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "card"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeChats       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "chats"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCheck       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "check"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCrossBorder BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "cross_border"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeDkNets      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeEft         BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "eft"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeHuIcs       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeInterac     BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "interac"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMasav       BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "masav"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMxCcen      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNeft        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "neft"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNics        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "nics"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNzBecs      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypePlElixir    BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
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
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeWire        BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "wire"
	BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeZengin      BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeAuBecs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBacs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeBook, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCard, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeChats, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCheck, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeCrossBorder, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeDkNets, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeEft, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeHuIcs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeInterac, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMasav, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeMxCcen, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNeft, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNics, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeNzBecs, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypePlElixir, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeProvxchange, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeRoSent, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeRtp, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSen, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSepa, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSgGiro, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSic, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSignet, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeSknbi, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeWire, BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesExpectedPaymentCreateRequest struct {
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound,required"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound,required"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
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
	// Specifies a ledger transaction object that will be created with the expected
	// payment. If the ledger transaction cannot be created, then the expected payment
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the expected payment.
	LedgerTransaction param.Field[BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransaction] `json:"ledger_transaction"`
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
	ReconciliationRuleVariables param.Field[[]map[string]string] `json:"reconciliation_rule_variables"`
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

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequest) implementsBulkRequestNewParamsResourceUnion() {
}

// Specifies a ledger transaction object that will be created with the expected
// payment. If the ledger transaction cannot be created, then the expected payment
// creation will fail. The resulting ledger transaction will mirror the status of
// the expected payment.
type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransaction struct {
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerEntry] `json:"ledger_entries,required"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
	// reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatus] `json:"status"`
}

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerEntry struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
// reversal.
type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType string

const (
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypePaperItem             BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeReturn                BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType = "return"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeReversal              BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType = "reversal"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypePaperItem, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeReturn, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatus string

const (
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatusArchived BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatus = "archived"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatusPending  BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatus = "pending"
	BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatusPosted   BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatus = "posted"
)

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatusArchived, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatusPending, BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLedgerTransactionStatusPosted:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLineItem struct {
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

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesLedgerTransactionCreateRequest struct {
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerEntry] `json:"ledger_entries,required"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
	// reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus] `json:"status"`
}

func (r BulkRequestNewParamsResourcesLedgerTransactionCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesLedgerTransactionCreateRequest) implementsBulkRequestNewParamsResourceUnion() {
}

type BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerEntry struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
// reversal.
type BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType string

const (
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaperItem             BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeReturn                BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "return"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeReversal              BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "reversal"
)

func (r BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaperItem, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeReturn, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus string

const (
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusArchived BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus = "archived"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusPending  BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus = "pending"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusPosted   BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus = "posted"
)

func (r BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusArchived, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusPending, BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusPosted:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesTransactionCreateRequest struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// The date on which the transaction occurred.
	AsOfDate param.Field[time.Time] `json:"as_of_date,required" format:"date"`
	// Either `credit` or `debit`.
	Direction param.Field[string] `json:"direction,required"`
	// The ID of the relevant Internal Account.
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode param.Field[string] `json:"vendor_code,required"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, `us_bank`, or others.
	VendorCodeType param.Field[string] `json:"vendor_code_type,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// This field will be `true` if the transaction has posted to the account.
	Posted param.Field[bool] `json:"posted"`
	// The type of the transaction. Examples could be
	// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
	Type param.Field[BulkRequestNewParamsResourcesTransactionCreateRequestType] `json:"type"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription param.Field[string] `json:"vendor_description"`
}

func (r BulkRequestNewParamsResourcesTransactionCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesTransactionCreateRequest) implementsBulkRequestNewParamsResourceUnion() {
}

// The type of the transaction. Examples could be
// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
type BulkRequestNewParamsResourcesTransactionCreateRequestType string

const (
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeACH         BulkRequestNewParamsResourcesTransactionCreateRequestType = "ach"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeAuBecs      BulkRequestNewParamsResourcesTransactionCreateRequestType = "au_becs"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeBacs        BulkRequestNewParamsResourcesTransactionCreateRequestType = "bacs"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeBook        BulkRequestNewParamsResourcesTransactionCreateRequestType = "book"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeCard        BulkRequestNewParamsResourcesTransactionCreateRequestType = "card"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeChats       BulkRequestNewParamsResourcesTransactionCreateRequestType = "chats"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeCheck       BulkRequestNewParamsResourcesTransactionCreateRequestType = "check"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeCrossBorder BulkRequestNewParamsResourcesTransactionCreateRequestType = "cross_border"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeDkNets      BulkRequestNewParamsResourcesTransactionCreateRequestType = "dk_nets"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeEft         BulkRequestNewParamsResourcesTransactionCreateRequestType = "eft"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeHuIcs       BulkRequestNewParamsResourcesTransactionCreateRequestType = "hu_ics"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeInterac     BulkRequestNewParamsResourcesTransactionCreateRequestType = "interac"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeMasav       BulkRequestNewParamsResourcesTransactionCreateRequestType = "masav"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeMxCcen      BulkRequestNewParamsResourcesTransactionCreateRequestType = "mx_ccen"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeNeft        BulkRequestNewParamsResourcesTransactionCreateRequestType = "neft"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeNics        BulkRequestNewParamsResourcesTransactionCreateRequestType = "nics"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeNzBecs      BulkRequestNewParamsResourcesTransactionCreateRequestType = "nz_becs"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypePlElixir    BulkRequestNewParamsResourcesTransactionCreateRequestType = "pl_elixir"
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
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeWire        BulkRequestNewParamsResourcesTransactionCreateRequestType = "wire"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeZengin      BulkRequestNewParamsResourcesTransactionCreateRequestType = "zengin"
	BulkRequestNewParamsResourcesTransactionCreateRequestTypeOther       BulkRequestNewParamsResourcesTransactionCreateRequestType = "other"
)

func (r BulkRequestNewParamsResourcesTransactionCreateRequestType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesTransactionCreateRequestTypeACH, BulkRequestNewParamsResourcesTransactionCreateRequestTypeAuBecs, BulkRequestNewParamsResourcesTransactionCreateRequestTypeBacs, BulkRequestNewParamsResourcesTransactionCreateRequestTypeBook, BulkRequestNewParamsResourcesTransactionCreateRequestTypeCard, BulkRequestNewParamsResourcesTransactionCreateRequestTypeChats, BulkRequestNewParamsResourcesTransactionCreateRequestTypeCheck, BulkRequestNewParamsResourcesTransactionCreateRequestTypeCrossBorder, BulkRequestNewParamsResourcesTransactionCreateRequestTypeDkNets, BulkRequestNewParamsResourcesTransactionCreateRequestTypeEft, BulkRequestNewParamsResourcesTransactionCreateRequestTypeHuIcs, BulkRequestNewParamsResourcesTransactionCreateRequestTypeInterac, BulkRequestNewParamsResourcesTransactionCreateRequestTypeMasav, BulkRequestNewParamsResourcesTransactionCreateRequestTypeMxCcen, BulkRequestNewParamsResourcesTransactionCreateRequestTypeNeft, BulkRequestNewParamsResourcesTransactionCreateRequestTypeNics, BulkRequestNewParamsResourcesTransactionCreateRequestTypeNzBecs, BulkRequestNewParamsResourcesTransactionCreateRequestTypePlElixir, BulkRequestNewParamsResourcesTransactionCreateRequestTypeProvxchange, BulkRequestNewParamsResourcesTransactionCreateRequestTypeRoSent, BulkRequestNewParamsResourcesTransactionCreateRequestTypeRtp, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSeBankgirot, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSen, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSepa, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSgGiro, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSic, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSignet, BulkRequestNewParamsResourcesTransactionCreateRequestTypeSknbi, BulkRequestNewParamsResourcesTransactionCreateRequestTypeWire, BulkRequestNewParamsResourcesTransactionCreateRequestTypeZengin, BulkRequestNewParamsResourcesTransactionCreateRequestTypeOther:
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

func (r BulkRequestNewParamsResourcesID) implementsBulkRequestNewParamsResourceUnion() {}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID struct {
	ID         param.Field[string]                                                                 `json:"id" format:"uuid"`
	Accounting param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDAccounting] `json:"accounting"`
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

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithID) implementsBulkRequestNewParamsResourceUnion() {
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. Only applies to wire
// payment orders. Can be one of shared, sender, or receiver, which correspond
// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
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
	AccountType    param.Field[ExternalAccountType]                                                                         `json:"account_type"`
	ContactDetails param.Field[[]BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetail] `json:"contact_details"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccount] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyAddress] `json:"party_address"`
	PartyIdentifier param.Field[string]                                                                                   `json:"party_identifier"`
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
	AccountNumber     param.Field[string]                                                                                                      `json:"account_number,required"`
	AccountNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeAuNumber      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "au_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeClabe         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "clabe"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeHkNumber      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "hk_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIban          BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "iban"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIDNumber      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "id_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeNzNumber      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "nz_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeOther         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "other"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypePan           BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "pan"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeSgNumber      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "sg_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeWalletAddress BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeAuNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeClabe, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeHkNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIban, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeIDNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeNzNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeOther, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypePan, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeSgNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetail struct {
	ContactIdentifier     param.Field[string]                                                                                                          `json:"contact_identifier"`
	ContactIdentifierType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierTypeEmail       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierType = "email"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierTypePhoneNumber BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierTypeWebsite     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierType = "website"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierTypeEmail, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierTypePhoneNumber, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountContactDetailsContactIdentifierTypeWebsite:
		return true
	}
	return false
}

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccount struct {
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[shared.TransactionDirection] `json:"normal_balance,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The array of ledger account category ids that this ledger account should be a
	// child of.
	LedgerAccountCategoryIDs param.Field[[]string] `json:"ledger_account_category_ids" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeCounterparty    BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType = "counterparty"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeExternalAccount BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType = "external_account"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeInternalAccount BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType = "internal_account"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeVirtualAccount  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType = "virtual_account"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeCounterparty, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeExternalAccount, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeInternalAccount, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountLedgerAccountLedgerableTypeVirtualAccount:
		return true
	}
	return false
}

// Required if receiving wire payments.
type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
	Line1   param.Field[string] `json:"line1"`
	Line2   param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Region or State.
	Region param.Field[string] `json:"region"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	RoutingNumber     param.Field[string]                                                                                                      `json:"routing_number,required"`
	RoutingNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
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
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeAba, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeChips, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeCnaps, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeIDSknbiCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeMxBankIdentifier, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypePlNationalClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeSwift, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType string

const (
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeACH         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "ach"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeAuBecs      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "au_becs"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBacs        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "bacs"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBook        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "book"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCard        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "card"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeChats       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "chats"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCheck       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "check"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCrossBorder BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "cross_border"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeDkNets      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "dk_nets"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeEft         BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "eft"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeHuIcs       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "hu_ics"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeInterac     BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "interac"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMasav       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "masav"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMxCcen      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "mx_ccen"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNeft        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "neft"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNics        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "nics"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNzBecs      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypePlElixir    BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "pl_elixir"
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
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeWire        BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "wire"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeZengin      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType = "zengin"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeACH, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeAuBecs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBacs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeBook, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCard, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeChats, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCheck, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeCrossBorder, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeDkNets, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeEft, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeHuIcs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeInterac, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMasav, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeMxCcen, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNeft, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNics, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeNzBecs, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypePlElixir, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeProvxchange, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeRoSent, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeRtp, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSeBankgirot, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSen, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSepa, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSgGiro, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSic, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSignet, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeSknbi, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeWire, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDReceivingAccountRoutingDetailsPaymentTypeZengin:
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
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusNeedsApproval BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "needs_approval"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusPending       BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "pending"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusProcessing    BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "processing"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReturned      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "returned"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReversed      BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "reversed"
	BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusSent          BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus = "sent"
)

func (r BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusApproved, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusCancelled, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusCompleted, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusDenied, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusFailed, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusNeedsApproval, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusPending, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusProcessing, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReturned, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusReversed, BulkRequestNewParamsResourcesPaymentOrderUpdateRequestWithIDStatusSent:
		return true
	}
	return false
}

type BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound"`
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
	Direction param.Field[shared.TransactionDirection] `json:"direction"`
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
	ReconciliationRuleVariables param.Field[[]map[string]string] `json:"reconciliation_rule_variables"`
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

func (r BulkRequestNewParamsResourcesExpectedPaymentUpdateRequestWithID) implementsBulkRequestNewParamsResourceUnion() {
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

func (r BulkRequestNewParamsResourcesTransactionUpdateRequestWithID) implementsBulkRequestNewParamsResourceUnion() {
}

type BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerEntry] `json:"ledger_entries"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
	// reversal.
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

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithID) implementsBulkRequestNewParamsResourceUnion() {
}

type BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerEntry struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
// reversal.
type BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType string

const (
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypePaperItem             BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReturn                BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "return"
	BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReversal              BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType = "reversal"
)

func (r BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypePaperItem, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReturn, BulkRequestNewParamsResourcesLedgerTransactionUpdateRequestWithIDLedgerableTypeReversal:
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

// The party that will pay the fees for the payment order. Only applies to wire
// payment orders. Can be one of shared, sender, or receiver, which correspond
// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
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
	BulkRequestNewParamsResourcesStatusNeedsApproval BulkRequestNewParamsResourcesStatus = "needs_approval"
	BulkRequestNewParamsResourcesStatusProcessing    BulkRequestNewParamsResourcesStatus = "processing"
	BulkRequestNewParamsResourcesStatusReturned      BulkRequestNewParamsResourcesStatus = "returned"
	BulkRequestNewParamsResourcesStatusReversed      BulkRequestNewParamsResourcesStatus = "reversed"
	BulkRequestNewParamsResourcesStatusSent          BulkRequestNewParamsResourcesStatus = "sent"
	BulkRequestNewParamsResourcesStatusReconciled    BulkRequestNewParamsResourcesStatus = "reconciled"
)

func (r BulkRequestNewParamsResourcesStatus) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesStatusArchived, BulkRequestNewParamsResourcesStatusPending, BulkRequestNewParamsResourcesStatusPosted, BulkRequestNewParamsResourcesStatusApproved, BulkRequestNewParamsResourcesStatusCancelled, BulkRequestNewParamsResourcesStatusCompleted, BulkRequestNewParamsResourcesStatusDenied, BulkRequestNewParamsResourcesStatusFailed, BulkRequestNewParamsResourcesStatusNeedsApproval, BulkRequestNewParamsResourcesStatusProcessing, BulkRequestNewParamsResourcesStatusReturned, BulkRequestNewParamsResourcesStatusReversed, BulkRequestNewParamsResourcesStatusSent, BulkRequestNewParamsResourcesStatusReconciled:
		return true
	}
	return false
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
// reversal.
type BulkRequestNewParamsResourcesLedgerableType string

const (
	BulkRequestNewParamsResourcesLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesLedgerableTypePaperItem             BulkRequestNewParamsResourcesLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesLedgerableTypeReturn                BulkRequestNewParamsResourcesLedgerableType = "return"
	BulkRequestNewParamsResourcesLedgerableTypeReversal              BulkRequestNewParamsResourcesLedgerableType = "reversal"
)

func (r BulkRequestNewParamsResourcesLedgerableType) IsKnown() bool {
	switch r {
	case BulkRequestNewParamsResourcesLedgerableTypeExpectedPayment, BulkRequestNewParamsResourcesLedgerableTypeIncomingPaymentDetail, BulkRequestNewParamsResourcesLedgerableTypePaperItem, BulkRequestNewParamsResourcesLedgerableTypePaymentOrder, BulkRequestNewParamsResourcesLedgerableTypeReturn, BulkRequestNewParamsResourcesLedgerableTypeReversal:
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
)

func (r BulkRequestListParamsActionType) IsKnown() bool {
	switch r {
	case BulkRequestListParamsActionTypeCreate, BulkRequestListParamsActionTypeUpdate:
		return true
	}
	return false
}

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestListParamsResourceType string

const (
	BulkRequestListParamsResourceTypePaymentOrder      BulkRequestListParamsResourceType = "payment_order"
	BulkRequestListParamsResourceTypeLedgerTransaction BulkRequestListParamsResourceType = "ledger_transaction"
	BulkRequestListParamsResourceTypeTransaction       BulkRequestListParamsResourceType = "transaction"
	BulkRequestListParamsResourceTypeExpectedPayment   BulkRequestListParamsResourceType = "expected_payment"
)

func (r BulkRequestListParamsResourceType) IsKnown() bool {
	switch r {
	case BulkRequestListParamsResourceTypePaymentOrder, BulkRequestListParamsResourceTypeLedgerTransaction, BulkRequestListParamsResourceTypeTransaction, BulkRequestListParamsResourceTypeExpectedPayment:
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
