package moderntreasury

import (
	"context"
	"fmt"
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

// ExpectedPaymentService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewExpectedPaymentService] method instead.
type ExpectedPaymentService struct {
	Options []option.RequestOption
}

// NewExpectedPaymentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExpectedPaymentService(opts ...option.RequestOption) (r *ExpectedPaymentService) {
	r = &ExpectedPaymentService{}
	r.Options = opts
	return
}

// create expected payment
func (r *ExpectedPaymentService) New(ctx context.Context, body ExpectedPaymentNewParams, opts ...option.RequestOption) (res *ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/expected_payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get expected payment
func (r *ExpectedPaymentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update expected payment
func (r *ExpectedPaymentService) Update(ctx context.Context, id string, body ExpectedPaymentUpdateParams, opts ...option.RequestOption) (res *ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list expected_payments
func (r *ExpectedPaymentService) List(ctx context.Context, query ExpectedPaymentListParams, opts ...option.RequestOption) (res *shared.Page[ExpectedPayment], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/expected_payments"
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

// list expected_payments
func (r *ExpectedPaymentService) ListAutoPaging(ctx context.Context, query ExpectedPaymentListParams, opts ...option.RequestOption) *shared.PageAutoPager[ExpectedPayment] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete expected payment
func (r *ExpectedPaymentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ExpectedPayment struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound int64 `json:"amount_upper_bound,required"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound int64 `json:"amount_lower_bound,required"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction ExpectedPaymentDirection `json:"direction,required"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type ExpectedPaymentType `json:"type,required,nullable"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account.
	Currency shared.Currency `json:"currency,required,nullable"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound time.Time `json:"date_upper_bound,required,nullable" format:"date"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound time.Time `json:"date_lower_bound,required,nullable" format:"date"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// The statement description you expect to see on the transaction. For ACH
	// payments, this will be the full line item passed from the bank. For wire
	// payments, this will be the OBI field on the wire. For check payments, this will
	// be the memo field.
	StatementDescriptor string `json:"statement_descriptor,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID string `json:"counterparty_id,required,nullable" format:"uuid"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The ID of the Transaction this expected payment object has been matched to.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the Transaction Line Item this expected payment has been matched to.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// One of unreconciled, reconciled, or archived.
	Status ExpectedPaymentStatus `json:"status,required"`
	// One of manual if this expected payment was manually reconciled in the dashboard,
	// automatic if it was automatically reconciled by Modern Treasury, or null if it
	// is unreconciled.
	ReconciliationMethod ExpectedPaymentReconciliationMethod `json:"reconciliation_method,required,nullable"`
	// The ID of the ledger transaction linked to the expected payment.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	JSON                expectedPaymentJSON
}

// expectedPaymentJSON contains the JSON metadata for the struct [ExpectedPayment]
type expectedPaymentJSON struct {
	ID                    apijson.Field
	Object                apijson.Field
	LiveMode              apijson.Field
	CreatedAt             apijson.Field
	UpdatedAt             apijson.Field
	AmountUpperBound      apijson.Field
	AmountLowerBound      apijson.Field
	Direction             apijson.Field
	InternalAccountID     apijson.Field
	Type                  apijson.Field
	Currency              apijson.Field
	DateUpperBound        apijson.Field
	DateLowerBound        apijson.Field
	Description           apijson.Field
	StatementDescriptor   apijson.Field
	Metadata              apijson.Field
	CounterpartyID        apijson.Field
	RemittanceInformation apijson.Field
	TransactionID         apijson.Field
	TransactionLineItemID apijson.Field
	Status                apijson.Field
	ReconciliationMethod  apijson.Field
	LedgerTransactionID   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ExpectedPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExpectedPaymentDirection string

const (
	ExpectedPaymentDirectionCredit ExpectedPaymentDirection = "credit"
	ExpectedPaymentDirectionDebit  ExpectedPaymentDirection = "debit"
)

type ExpectedPaymentType string

const (
	ExpectedPaymentTypeACH         ExpectedPaymentType = "ach"
	ExpectedPaymentTypeAuBecs      ExpectedPaymentType = "au_becs"
	ExpectedPaymentTypeBacs        ExpectedPaymentType = "bacs"
	ExpectedPaymentTypeBook        ExpectedPaymentType = "book"
	ExpectedPaymentTypeCard        ExpectedPaymentType = "card"
	ExpectedPaymentTypeCheck       ExpectedPaymentType = "check"
	ExpectedPaymentTypeCrossBorder ExpectedPaymentType = "cross_border"
	ExpectedPaymentTypeEft         ExpectedPaymentType = "eft"
	ExpectedPaymentTypeInterac     ExpectedPaymentType = "interac"
	ExpectedPaymentTypeMasav       ExpectedPaymentType = "masav"
	ExpectedPaymentTypeNeft        ExpectedPaymentType = "neft"
	ExpectedPaymentTypeProvxchange ExpectedPaymentType = "provxchange"
	ExpectedPaymentTypeRtp         ExpectedPaymentType = "rtp"
	ExpectedPaymentTypeSen         ExpectedPaymentType = "sen"
	ExpectedPaymentTypeSepa        ExpectedPaymentType = "sepa"
	ExpectedPaymentTypeSignet      ExpectedPaymentType = "signet"
	ExpectedPaymentTypeWire        ExpectedPaymentType = "wire"
)

type ExpectedPaymentStatus string

const (
	ExpectedPaymentStatusArchived     ExpectedPaymentStatus = "archived"
	ExpectedPaymentStatusReconciled   ExpectedPaymentStatus = "reconciled"
	ExpectedPaymentStatusUnreconciled ExpectedPaymentStatus = "unreconciled"
)

type ExpectedPaymentReconciliationMethod string

const (
	ExpectedPaymentReconciliationMethodAutomatic ExpectedPaymentReconciliationMethod = "automatic"
	ExpectedPaymentReconciliationMethodManual    ExpectedPaymentReconciliationMethod = "manual"
)

type ExpectedPaymentNewParams struct {
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound,required"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound,required"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction param.Field[ExpectedPaymentNewParamsDirection] `json:"direction,required"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type param.Field[ExpectedPaymentType] `json:"type,nullable"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account.
	Currency param.Field[shared.Currency] `json:"currency,nullable"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound,nullable" format:"date"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound,nullable" format:"date"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description,nullable"`
	// The statement description you expect to see on the transaction. For ACH
	// payments, this will be the full line item passed from the bank. For wire
	// payments, this will be the OBI field on the wire. For check payments, this will
	// be the memo field.
	StatementDescriptor param.Field[string] `json:"statement_descriptor,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID param.Field[string] `json:"counterparty_id,nullable" format:"uuid"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string]                              `json:"remittance_information,nullable"`
	LineItems             param.Field[[]ExpectedPaymentNewParamsLineItems] `json:"line_items"`
}

func (r ExpectedPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExpectedPaymentNewParamsDirection string

const (
	ExpectedPaymentNewParamsDirectionCredit ExpectedPaymentNewParamsDirection = "credit"
	ExpectedPaymentNewParamsDirectionDebit  ExpectedPaymentNewParamsDirection = "debit"
)

type ExpectedPaymentNewParamsLineItems struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A free-form description of the line item.
	Description param.Field[string] `json:"description,nullable"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID param.Field[string] `json:"accounting_category_id,nullable"`
}

type ExpectedPaymentUpdateParams struct {
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction param.Field[ExpectedPaymentUpdateParamsDirection] `json:"direction"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID param.Field[string] `json:"internal_account_id" format:"uuid"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type param.Field[ExpectedPaymentType] `json:"type,nullable"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account.
	Currency param.Field[shared.Currency] `json:"currency,nullable"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound,nullable" format:"date"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound,nullable" format:"date"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description,nullable"`
	// The statement description you expect to see on the transaction. For ACH
	// payments, this will be the full line item passed from the bank. For wire
	// payments, this will be the OBI field on the wire. For check payments, this will
	// be the memo field.
	StatementDescriptor param.Field[string] `json:"statement_descriptor,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID param.Field[string] `json:"counterparty_id,nullable" format:"uuid"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation param.Field[string] `json:"remittance_information,nullable"`
}

func (r ExpectedPaymentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExpectedPaymentUpdateParamsDirection string

const (
	ExpectedPaymentUpdateParamsDirectionCredit ExpectedPaymentUpdateParamsDirection = "credit"
	ExpectedPaymentUpdateParamsDirectionDebit  ExpectedPaymentUpdateParamsDirection = "debit"
)

type ExpectedPaymentListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
	// One of unreconciled, reconciled, or archived.
	Status param.Field[ExpectedPaymentListParamsStatus] `query:"status"`
	// Specify internal_account_id to see expected_payments for a specific account.
	InternalAccountID param.Field[string] `query:"internal_account_id"`
	// One of credit, debit
	Direction param.Field[ExpectedPaymentListParamsDirection] `query:"direction"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp,sen,
	// sepa, signet, wire
	Type param.Field[ExpectedPaymentListParamsType] `query:"type"`
	// Specify counterparty_id to see expected_payments for a specific account.
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Used to return expected payments created after some datetime
	CreatedAtLowerBound param.Field[time.Time] `query:"created_at_lower_bound" format:"date-time"`
	// Used to return expected payments created before some datetime
	CreatedAtUpperBound param.Field[time.Time] `query:"created_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [ExpectedPaymentListParams]'s query parameters as
// `url.Values`.
func (r ExpectedPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ExpectedPaymentListParamsStatus string

const (
	ExpectedPaymentListParamsStatusArchived     ExpectedPaymentListParamsStatus = "archived"
	ExpectedPaymentListParamsStatusReconciled   ExpectedPaymentListParamsStatus = "reconciled"
	ExpectedPaymentListParamsStatusUnreconciled ExpectedPaymentListParamsStatus = "unreconciled"
)

type ExpectedPaymentListParamsDirection string

const (
	ExpectedPaymentListParamsDirectionCredit ExpectedPaymentListParamsDirection = "credit"
	ExpectedPaymentListParamsDirectionDebit  ExpectedPaymentListParamsDirection = "debit"
)

type ExpectedPaymentListParamsType string

const (
	ExpectedPaymentListParamsTypeACH         ExpectedPaymentListParamsType = "ach"
	ExpectedPaymentListParamsTypeAuBecs      ExpectedPaymentListParamsType = "au_becs"
	ExpectedPaymentListParamsTypeBacs        ExpectedPaymentListParamsType = "bacs"
	ExpectedPaymentListParamsTypeBook        ExpectedPaymentListParamsType = "book"
	ExpectedPaymentListParamsTypeCard        ExpectedPaymentListParamsType = "card"
	ExpectedPaymentListParamsTypeCheck       ExpectedPaymentListParamsType = "check"
	ExpectedPaymentListParamsTypeCrossBorder ExpectedPaymentListParamsType = "cross_border"
	ExpectedPaymentListParamsTypeEft         ExpectedPaymentListParamsType = "eft"
	ExpectedPaymentListParamsTypeInterac     ExpectedPaymentListParamsType = "interac"
	ExpectedPaymentListParamsTypeMasav       ExpectedPaymentListParamsType = "masav"
	ExpectedPaymentListParamsTypeNeft        ExpectedPaymentListParamsType = "neft"
	ExpectedPaymentListParamsTypeProvxchange ExpectedPaymentListParamsType = "provxchange"
	ExpectedPaymentListParamsTypeRtp         ExpectedPaymentListParamsType = "rtp"
	ExpectedPaymentListParamsTypeSen         ExpectedPaymentListParamsType = "sen"
	ExpectedPaymentListParamsTypeSepa        ExpectedPaymentListParamsType = "sepa"
	ExpectedPaymentListParamsTypeSignet      ExpectedPaymentListParamsType = "signet"
	ExpectedPaymentListParamsTypeWire        ExpectedPaymentListParamsType = "wire"
)
