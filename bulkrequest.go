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

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiform"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// BulkRequestService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewBulkRequestService]
// method instead.
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
	path := fmt.Sprintf("api/bulk_requests/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// list bulk_requests
func (r *BulkRequestService) List(ctx context.Context, query BulkRequestListParams, opts ...option.RequestOption) (res *shared.Page[BulkRequest], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *BulkRequestService) ListAutoPaging(ctx context.Context, query BulkRequestListParams, opts ...option.RequestOption) *shared.PageAutoPager[BulkRequest] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
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
	TotalResourceCount int64     `json:"total_resource_count,required"`
	UpdatedAt          time.Time `json:"updated_at,required" format:"date-time"`
	JSON               bulkRequestJSON
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

// One of create, or update.
type BulkRequestActionType string

const (
	BulkRequestActionTypeCreate BulkRequestActionType = "create"
	BulkRequestActionTypeUpdate BulkRequestActionType = "update"
)

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestResourceType string

const (
	BulkRequestResourceTypePaymentOrder      BulkRequestResourceType = "payment_order"
	BulkRequestResourceTypeLedgerTransaction BulkRequestResourceType = "ledger_transaction"
	BulkRequestResourceTypeExpectedPayment   BulkRequestResourceType = "expected_payment"
)

// One of pending, processing, or completed.
type BulkRequestStatus string

const (
	BulkRequestStatusPending    BulkRequestStatus = "pending"
	BulkRequestStatusProcessing BulkRequestStatus = "processing"
	BulkRequestStatusCompleted  BulkRequestStatus = "completed"
)

type BulkRequestNewParams struct {
	// One of create, or update.
	ActionType param.Field[BulkRequestNewParamsActionType] `json:"action_type,required"`
	// One of payment_order, expected_payment, or ledger_transaction.
	ResourceType param.Field[BulkRequestNewParamsResourceType] `json:"resource_type,required"`
	// An array of objects where each object contains the input params for a single
	// `action_type` request on a `resource_type` resource
	Resources param.Field[[]BulkRequestNewParamsResource] `json:"resources,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// One of create, or update.
type BulkRequestNewParamsActionType string

const (
	BulkRequestNewParamsActionTypeCreate BulkRequestNewParamsActionType = "create"
	BulkRequestNewParamsActionTypeUpdate BulkRequestNewParamsActionType = "update"
)

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestNewParamsResourceType string

const (
	BulkRequestNewParamsResourceTypePaymentOrder      BulkRequestNewParamsResourceType = "payment_order"
	BulkRequestNewParamsResourceTypeLedgerTransaction BulkRequestNewParamsResourceType = "ledger_transaction"
	BulkRequestNewParamsResourceTypeExpectedPayment   BulkRequestNewParamsResourceType = "expected_payment"
)

// Satisfied by [BulkRequestNewParamsResourcesPaymentOrderCreateRequest],
// [BulkRequestNewParamsResourcesExpectedPaymentCreateRequest],
// [BulkRequestNewParamsResourcesLedgerTransactionCreateRequest],
// [BulkRequestNewParamsResourcePaymentOrderUpdateRequestWithID],
// [BulkRequestNewParamsResourceExpectedPaymentUpdateRequestWithID],
// [BulkRequestNewParamsResourceLedgerTransactionUpdateRequestWithID].
type BulkRequestNewParamsResource interface {
	implementsBulkRequestNewParamsResource()
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequest struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented as
	// 1000 (cents). For RTP, the maximum amount allowed by the network is $100,000.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirection] `json:"direction,required"`
	// The ID of one of your organization's internal accounts.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// One of `ach`, `bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`,
	// `bacs`, `au_becs`, `interac`, `neft`, `nics`, `nz_national_clearing_code`,
	// `sic`, `signet`, `provexchange`, `zengin`.
	Type       param.Field[PaymentOrderType]                                                 `json:"type,required"`
	Accounting param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestAccounting] `json:"accounting"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id" format:"uuid"`
	// The ID of one of your accounting ledger classes. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingLedgerClassID param.Field[string] `json:"accounting_ledger_class_id" format:"uuid"`
	// The party that will pay the fees for the payment order. Only applies to wire
	// payment orders. Can be one of shared, sender, or receiver, which correspond
	// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
	ChargeBearer param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearer] `json:"charge_bearer"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// An array of documents to be attached to the payment order. Note that if you
	// attach documents, the request's content type must be `multipart/form-data`.
	Documents param.Field[[]BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocument] `json:"documents"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// Specifies a ledger transaction object that will be created with the payment
	// order. If the ledger transaction cannot be created, then the payment order
	// creation will fail. The resulting ledger transaction will mirror the status of
	// the payment order.
	LedgerTransaction param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransaction] `json:"ledger_transaction"`
	// Either ledger_transaction or ledger_transaction_id can be provided. Only a
	// pending ledger transaction can be attached upon payment order creation. Once the
	// payment order is created, the status of the ledger transaction tracks the
	// payment order automatically.
	LedgerTransactionID param.Field[string] `json:"ledger_transaction_id" format:"uuid"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]BulkRequestNewParamsResourcesPaymentOrderCreateRequestLineItem] `json:"line_items"`
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
	Priority param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriority] `json:"priority"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose param.Field[string] `json:"purpose"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccount] `json:"receiving_account"`
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

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequest) implementsBulkRequestNewParamsResource() {
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirection string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirectionCredit BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirection = "credit"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirectionDebit  BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirection = "debit"
)

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. Only applies to wire
// payment orders. Can be one of shared, sender, or receiver, which correspond
// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearer string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearerShared   BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearer = "shared"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearerSender   BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearer = "sender"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearerReceiver BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearer = "receiver"
)

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocument struct {
	// The unique identifier for the associated object.
	DocumentableID   param.Field[string]                                                                          `json:"documentable_id,required"`
	DocumentableType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType] `json:"documentable_type,required"`
	File             param.Field[io.Reader]                                                                       `json:"file,required" format:"binary"`
	// A category given to the document, can be `null`.
	DocumentType param.Field[string] `json:"document_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases                  BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "cases"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCounterparties         BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "counterparties"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeExpectedPayments       BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "expected_payments"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeExternalAccounts       BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "external_accounts"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeIncomingPaymentDetails BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "incoming_payment_details"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeInternalAccounts       BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "internal_accounts"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeOrganizations          BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "organizations"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypePaperItems             BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "paper_items"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypePaymentOrders          BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "payment_orders"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeTransactions           BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "transactions"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeDecisions              BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "decisions"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeConnections            BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableType = "connections"
)

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackTypeACH BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackType = "ach"
)

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicator string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicatorFixedToVariable BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicator = "fixed_to_variable"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicatorVariableToFixed BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicator = "variable_to_fixed"
)

// Specifies a ledger transaction object that will be created with the payment
// order. If the ledger transaction cannot be created, then the payment order
// creation will fail. The resulting ledger transaction will mirror the status of
// the payment order.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransaction struct {
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerEntry] `json:"ledger_entries,required"`
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
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatus] `json:"status"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransaction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerEntry struct {
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

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeCounterparty          BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "counterparty"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeInternalAccount       BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "internal_account"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeLineItem              BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "line_item"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypePaperItem             BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypePaymentOrderAttempt   BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "payment_order_attempt"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeReturn                BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "return"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeReversal              BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableType = "reversal"
)

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatus string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatusArchived BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatus = "archived"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatusPending  BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatus = "pending"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatusPosted   BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatus = "posted"
)

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestLineItem struct {
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

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriority string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriorityHigh   BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriority = "high"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriorityNormal BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriority = "normal"
)

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccount struct {
	AccountDetails param.Field[[]BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]                                                                   `json:"account_type"`
	ContactDetails param.Field[[]BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetail] `json:"contact_details"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccount] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyAddress] `json:"party_address"`
	PartyIdentifier param.Field[string]                                                                             `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                                                `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                                                                `json:"account_number,required"`
	AccountNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban          BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType = "iban"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeClabe         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType = "clabe"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeWalletAddress BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypePan           BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType = "pan"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeOther         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberType = "other"
)

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetail struct {
	ContactIdentifier     param.Field[string]                                                                                                    `json:"contact_identifier"`
	ContactIdentifierType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail       BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierType = "email"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypePhoneNumber BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeWebsite     BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierType = "website"
)

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccount struct {
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
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableTypeExternalAccount BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableType = "external_account"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableTypeInternalAccount BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableType = "internal_account"
)

// Required if receiving wire payments.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyAddress struct {
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

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `individual` or `business`.
type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyTypeBusiness   BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyType = "business"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyTypeIndividual BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyType = "individual"
)

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                                                                `json:"routing_number,required"`
	RoutingNumberType param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba                     BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeChips                   BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
)

type BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType string

const (
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "ach"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeAuBecs      BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "au_becs"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeBacs        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "bacs"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeBook        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "book"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeCard        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "card"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeChats       BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "chats"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeCheck       BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "check"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeCrossBorder BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "cross_border"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeEft         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "eft"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeInterac     BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "interac"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeMasav       BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "masav"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeNeft        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "neft"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeNics        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "nics"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeNzBecs      BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeProvxchange BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "provxchange"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeRtp         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "rtp"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeSgGiro      BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeSeBankgirot BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeSen         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "sen"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeSepa        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "sepa"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeSic         BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "sic"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeSignet      BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "signet"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeWire        BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "wire"
	BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeZengin      BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentType = "zengin"
)

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
	Description param.Field[string]                                                              `json:"description"`
	LineItems   param.Field[[]BulkRequestNewParamsResourcesExpectedPaymentCreateRequestLineItem] `json:"line_items"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The reconciliation filters you have for this payment.
	ReconciliationFilters param.Field[interface{}] `json:"reconciliation_filters"`
	// The reconciliation groups you have for this payment.
	ReconciliationGroups param.Field[interface{}] `json:"reconciliation_groups"`
	// An array of reconciliation rule variables for this payment.
	ReconciliationRuleVariables param.Field[[]interface{}] `json:"reconciliation_rule_variables"`
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

func (r BulkRequestNewParamsResourcesExpectedPaymentCreateRequest) implementsBulkRequestNewParamsResource() {
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
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
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

func (r BulkRequestNewParamsResourcesLedgerTransactionCreateRequest) implementsBulkRequestNewParamsResource() {
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
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType string

const (
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeCounterparty          BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "counterparty"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeExpectedPayment       BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "expected_payment"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "incoming_payment_detail"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeInternalAccount       BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "internal_account"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeLineItem              BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "line_item"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaperItem             BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "paper_item"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaymentOrder          BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "payment_order"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypePaymentOrderAttempt   BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "payment_order_attempt"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeReturn                BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "return"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableTypeReversal              BulkRequestNewParamsResourcesLedgerTransactionCreateRequestLedgerableType = "reversal"
)

// To post a ledger transaction at creation, use `posted`.
type BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus string

const (
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusArchived BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus = "archived"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusPending  BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus = "pending"
	BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatusPosted   BulkRequestNewParamsResourcesLedgerTransactionCreateRequestStatus = "posted"
)

type BulkRequestNewParamsResourcePaymentOrderUpdateRequestWithID struct {
	ID         param.Field[string]                                        `json:"id" format:"uuid"`
	Accounting param.Field[BulkRequestNewParamsResourcesObjectAccounting] `json:"accounting"`
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
	ChargeBearer param.Field[BulkRequestNewParamsResourcesObjectChargeBearer] `json:"charge_bearer"`
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
	Direction param.Field[BulkRequestNewParamsResourcesObjectDirection] `json:"direction"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// RFP payments require an expires_at. This value must be past the effective_date.
	ExpiresAt param.Field[time.Time] `json:"expires_at" format:"date-time"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (type=rtp and fallback_type=ach)
	FallbackType param.Field[BulkRequestNewParamsResourcesObjectFallbackType] `json:"fallback_type"`
	// If present, indicates a specific foreign exchange contract number that has been
	// generated by your financial institution.
	ForeignExchangeContract param.Field[string] `json:"foreign_exchange_contract"`
	// Indicates the type of FX transfer to initiate, can be either
	// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
	// currency matches the originating account currency.
	ForeignExchangeIndicator param.Field[BulkRequestNewParamsResourcesObjectForeignExchangeIndicator] `json:"foreign_exchange_indicator"`
	// An array of line items that must sum up to the amount of the payment order.
	LineItems param.Field[[]BulkRequestNewParamsResourcesObjectLineItem] `json:"line_items"`
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
	Priority param.Field[BulkRequestNewParamsResourcesObjectPriority] `json:"priority"`
	// For `wire`, this is usually the purpose which is transmitted via the
	// "InstrForDbtrAgt" field in the ISO20022 file. If you are using Currencycloud,
	// this is the `payment.purpose_code` field. For `eft`, this field is the 3 digit
	// CPA Code that will be attached to the payment.
	Purpose param.Field[string] `json:"purpose"`
	// Either `receiving_account` or `receiving_account_id` must be present. When using
	// `receiving_account_id`, you may pass the id of an external account or an
	// internal account.
	ReceivingAccount param.Field[BulkRequestNewParamsResourcesObjectReceivingAccount] `json:"receiving_account"`
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
	Status param.Field[BulkRequestNewParamsResourcesObjectStatus] `json:"status"`
	// An additional layer of classification for the type of payment order you are
	// doing. This field is only used for `ach` payment orders currently. For `ach`
	// payment orders, the `subtype` represents the SEC code. We currently support
	// `CCD`, `PPD`, `IAT`, `CTX`, `WEB`, `CIE`, and `TEL`.
	Subtype param.Field[PaymentOrderSubtype] `json:"subtype"`
	// One of `ach`, `bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`, `sepa`,
	// `bacs`, `au_becs`, `interac`, `neft`, `nics`, `nz_national_clearing_code`,
	// `sic`, `signet`, `provexchange`, `zengin`.
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

func (r BulkRequestNewParamsResourcePaymentOrderUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourcePaymentOrderUpdateRequestWithID) implementsBulkRequestNewParamsResource() {
}

type BulkRequestNewParamsResourcesObjectAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r BulkRequestNewParamsResourcesObjectAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The party that will pay the fees for the payment order. Only applies to wire
// payment orders. Can be one of shared, sender, or receiver, which correspond
// respectively with the SWIFT 71A values `SHA`, `OUR`, `BEN`.
type BulkRequestNewParamsResourcesObjectChargeBearer string

const (
	BulkRequestNewParamsResourcesObjectChargeBearerShared   BulkRequestNewParamsResourcesObjectChargeBearer = "shared"
	BulkRequestNewParamsResourcesObjectChargeBearerSender   BulkRequestNewParamsResourcesObjectChargeBearer = "sender"
	BulkRequestNewParamsResourcesObjectChargeBearerReceiver BulkRequestNewParamsResourcesObjectChargeBearer = "receiver"
)

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type BulkRequestNewParamsResourcesObjectDirection string

const (
	BulkRequestNewParamsResourcesObjectDirectionCredit BulkRequestNewParamsResourcesObjectDirection = "credit"
	BulkRequestNewParamsResourcesObjectDirectionDebit  BulkRequestNewParamsResourcesObjectDirection = "debit"
)

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (type=rtp and fallback_type=ach)
type BulkRequestNewParamsResourcesObjectFallbackType string

const (
	BulkRequestNewParamsResourcesObjectFallbackTypeACH BulkRequestNewParamsResourcesObjectFallbackType = "ach"
)

// Indicates the type of FX transfer to initiate, can be either
// `variable_to_fixed`, `fixed_to_variable`, or `null` if the payment order
// currency matches the originating account currency.
type BulkRequestNewParamsResourcesObjectForeignExchangeIndicator string

const (
	BulkRequestNewParamsResourcesObjectForeignExchangeIndicatorFixedToVariable BulkRequestNewParamsResourcesObjectForeignExchangeIndicator = "fixed_to_variable"
	BulkRequestNewParamsResourcesObjectForeignExchangeIndicatorVariableToFixed BulkRequestNewParamsResourcesObjectForeignExchangeIndicator = "variable_to_fixed"
)

type BulkRequestNewParamsResourcesObjectLineItem struct {
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

func (r BulkRequestNewParamsResourcesObjectLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `normal` or `high`. For ACH and EFT payments, `high` represents a
// same-day ACH or EFT transfer, respectively. For check payments, `high` can mean
// an overnight check rather than standard mail.
type BulkRequestNewParamsResourcesObjectPriority string

const (
	BulkRequestNewParamsResourcesObjectPriorityHigh   BulkRequestNewParamsResourcesObjectPriority = "high"
	BulkRequestNewParamsResourcesObjectPriorityNormal BulkRequestNewParamsResourcesObjectPriority = "normal"
)

// Either `receiving_account` or `receiving_account_id` must be present. When using
// `receiving_account_id`, you may pass the id of an external account or an
// internal account.
type BulkRequestNewParamsResourcesObjectReceivingAccount struct {
	AccountDetails param.Field[[]BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]                                                `json:"account_type"`
	ContactDetails param.Field[[]BulkRequestNewParamsResourcesObjectReceivingAccountContactDetail] `json:"contact_details"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccount] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountPartyAddress] `json:"party_address"`
	PartyIdentifier param.Field[string]                                                          `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                                             `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetail] `json:"routing_details"`
}

func (r BulkRequestNewParamsResourcesObjectReceivingAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetail struct {
	AccountNumber     param.Field[string]                                                                             `json:"account_number,required"`
	AccountNumberType param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType string

const (
	BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberTypeIban          BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType = "iban"
	BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberTypeClabe         BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType = "clabe"
	BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberTypeWalletAddress BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType = "wallet_address"
	BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberTypePan           BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType = "pan"
	BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberTypeOther         BulkRequestNewParamsResourcesObjectReceivingAccountAccountDetailsAccountNumberType = "other"
)

type BulkRequestNewParamsResourcesObjectReceivingAccountContactDetail struct {
	ContactIdentifier     param.Field[string]                                                                                 `json:"contact_identifier"`
	ContactIdentifierType param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r BulkRequestNewParamsResourcesObjectReceivingAccountContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierType string

const (
	BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierTypeEmail       BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierType = "email"
	BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierTypePhoneNumber BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierType = "phone_number"
	BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierTypeWebsite     BulkRequestNewParamsResourcesObjectReceivingAccountContactDetailsContactIdentifierType = "website"
)

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccount struct {
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
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccountLedgerableType string

const (
	BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccountLedgerableTypeExternalAccount BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccountLedgerableType = "external_account"
	BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccountLedgerableTypeInternalAccount BulkRequestNewParamsResourcesObjectReceivingAccountLedgerAccountLedgerableType = "internal_account"
)

// Required if receiving wire payments.
type BulkRequestNewParamsResourcesObjectReceivingAccountPartyAddress struct {
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

func (r BulkRequestNewParamsResourcesObjectReceivingAccountPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `individual` or `business`.
type BulkRequestNewParamsResourcesObjectReceivingAccountPartyType string

const (
	BulkRequestNewParamsResourcesObjectReceivingAccountPartyTypeBusiness   BulkRequestNewParamsResourcesObjectReceivingAccountPartyType = "business"
	BulkRequestNewParamsResourcesObjectReceivingAccountPartyTypeIndividual BulkRequestNewParamsResourcesObjectReceivingAccountPartyType = "individual"
)

type BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                                             `json:"routing_number,required"`
	RoutingNumberType param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType string

const (
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeAba                     BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "aba"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeAuBsb                   BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "au_bsb"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeBrCodigo                BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "br_codigo"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeCaCpa                   BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "ca_cpa"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeChips                   BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "chips"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeCnaps                   BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "cnaps"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeGBSortCode              BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "gb_sort_code"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeHkInterbankClearingCode BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeInIfsc                  BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "in_ifsc"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeMyBranchCode            BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "my_branch_code"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeNzNationalClearingCode  BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeSwift                   BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "swift"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberTypeJpZenginCode            BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsRoutingNumberType = "jp_zengin_code"
)

type BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType string

const (
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeACH         BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "ach"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeAuBecs      BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "au_becs"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeBacs        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "bacs"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeBook        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "book"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeCard        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "card"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeChats       BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "chats"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeCheck       BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "check"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeCrossBorder BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "cross_border"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeEft         BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "eft"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeInterac     BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "interac"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeMasav       BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "masav"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeNeft        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "neft"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeNics        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "nics"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeNzBecs      BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "nz_becs"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeProvxchange BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "provxchange"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeRtp         BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "rtp"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeSgGiro      BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "sg_giro"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeSeBankgirot BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "se_bankgirot"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeSen         BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "sen"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeSepa        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "sepa"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeSic         BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "sic"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeSignet      BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "signet"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeWire        BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "wire"
	BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentTypeZengin      BulkRequestNewParamsResourcesObjectReceivingAccountRoutingDetailsPaymentType = "zengin"
)

// To cancel a payment order, use `cancelled`. To redraft a returned payment order,
// use `approved`. To undo approval on a denied or approved payment order, use
// `needs_approval`.
type BulkRequestNewParamsResourcesObjectStatus string

const (
	BulkRequestNewParamsResourcesObjectStatusApproved      BulkRequestNewParamsResourcesObjectStatus = "approved"
	BulkRequestNewParamsResourcesObjectStatusCancelled     BulkRequestNewParamsResourcesObjectStatus = "cancelled"
	BulkRequestNewParamsResourcesObjectStatusCompleted     BulkRequestNewParamsResourcesObjectStatus = "completed"
	BulkRequestNewParamsResourcesObjectStatusDenied        BulkRequestNewParamsResourcesObjectStatus = "denied"
	BulkRequestNewParamsResourcesObjectStatusFailed        BulkRequestNewParamsResourcesObjectStatus = "failed"
	BulkRequestNewParamsResourcesObjectStatusNeedsApproval BulkRequestNewParamsResourcesObjectStatus = "needs_approval"
	BulkRequestNewParamsResourcesObjectStatusPending       BulkRequestNewParamsResourcesObjectStatus = "pending"
	BulkRequestNewParamsResourcesObjectStatusProcessing    BulkRequestNewParamsResourcesObjectStatus = "processing"
	BulkRequestNewParamsResourcesObjectStatusReturned      BulkRequestNewParamsResourcesObjectStatus = "returned"
	BulkRequestNewParamsResourcesObjectStatusReversed      BulkRequestNewParamsResourcesObjectStatus = "reversed"
	BulkRequestNewParamsResourcesObjectStatusSent          BulkRequestNewParamsResourcesObjectStatus = "sent"
)

type BulkRequestNewParamsResourceExpectedPaymentUpdateRequestWithID struct {
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
	ReconciliationRuleVariables param.Field[[]interface{}] `json:"reconciliation_rule_variables"`
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

func (r BulkRequestNewParamsResourceExpectedPaymentUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourceExpectedPaymentUpdateRequestWithID) implementsBulkRequestNewParamsResource() {
}

type BulkRequestNewParamsResourceLedgerTransactionUpdateRequestWithID struct {
	ID param.Field[string] `json:"id" format:"uuid"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]BulkRequestNewParamsResourcesObjectLedgerEntry] `json:"ledger_entries"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[BulkRequestNewParamsResourcesObjectStatus] `json:"status"`
}

func (r BulkRequestNewParamsResourceLedgerTransactionUpdateRequestWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BulkRequestNewParamsResourceLedgerTransactionUpdateRequestWithID) implementsBulkRequestNewParamsResource() {
}

type BulkRequestNewParamsResourcesObjectLedgerEntry struct {
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

func (r BulkRequestNewParamsResourcesObjectLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// One of payment_order, expected_payment, or ledger_transaction.
type BulkRequestListParamsResourceType string

const (
	BulkRequestListParamsResourceTypePaymentOrder      BulkRequestListParamsResourceType = "payment_order"
	BulkRequestListParamsResourceTypeLedgerTransaction BulkRequestListParamsResourceType = "ledger_transaction"
	BulkRequestListParamsResourceTypeExpectedPayment   BulkRequestListParamsResourceType = "expected_payment"
)

// One of pending, processing, or completed.
type BulkRequestListParamsStatus string

const (
	BulkRequestListParamsStatusPending    BulkRequestListParamsStatus = "pending"
	BulkRequestListParamsStatusProcessing BulkRequestListParamsStatus = "processing"
	BulkRequestListParamsStatusCompleted  BulkRequestListParamsStatus = "completed"
)