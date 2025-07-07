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
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// ExpectedPaymentService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExpectedPaymentService] method instead.
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update expected payment
func (r *ExpectedPaymentService) Update(ctx context.Context, id string, body ExpectedPaymentUpdateParams, opts ...option.RequestOption) (res *ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list expected_payments
func (r *ExpectedPaymentService) List(ctx context.Context, query ExpectedPaymentListParams, opts ...option.RequestOption) (res *pagination.Page[ExpectedPayment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *ExpectedPaymentService) ListAutoPaging(ctx context.Context, query ExpectedPaymentListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExpectedPayment] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete expected payment
func (r *ExpectedPaymentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ExpectedPayment, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/expected_payments/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ExpectedPayment struct {
	ID string `json:"id,required" format:"uuid"`
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound int64 `json:"amount_lower_bound,required,nullable"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound int64 `json:"amount_upper_bound,required,nullable"`
	// The ID of the counterparty you expect for this payment.
	CounterpartyID string    `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account.
	Currency shared.Currency `json:"currency,required,nullable"`
	// The earliest date the payment may come in. Format: yyyy-mm-dd
	DateLowerBound time.Time `json:"date_lower_bound,required,nullable" format:"date"`
	// The latest date the payment may come in. Format: yyyy-mm-dd
	DateUpperBound time.Time `json:"date_upper_bound,required,nullable" format:"date"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction ExpectedPaymentDirection `json:"direction,required,nullable"`
	// The ID of the Internal Account for the expected payment.
	InternalAccountID string `json:"internal_account_id,required,nullable" format:"uuid"`
	// The ID of the ledger transaction linked to the expected payment.
	LedgerTransactionID string `json:"ledger_transaction_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The reconciliation filters you have for this payment.
	ReconciliationFilters interface{} `json:"reconciliation_filters,required,nullable"`
	// The reconciliation groups you have for this payment.
	ReconciliationGroups interface{} `json:"reconciliation_groups,required,nullable"`
	// One of manual if this expected payment was manually reconciled in the dashboard,
	// automatic if it was automatically reconciled by Modern Treasury, or null if it
	// is unreconciled.
	ReconciliationMethod ExpectedPaymentReconciliationMethod `json:"reconciliation_method,required,nullable"`
	// An array of reconciliation rule variables for this payment.
	ReconciliationRuleVariables []ReconciliationRule `json:"reconciliation_rule_variables,required,nullable"`
	// For `ach`, this field will be passed through on an addenda record. For `wire`
	// payments the field will be passed through as the "Originator to Beneficiary
	// Information", also known as OBI or Fedwire tag 6000.
	RemittanceInformation string `json:"remittance_information,required,nullable"`
	// The statement description you expect to see on the transaction. For ACH
	// payments, this will be the full line item passed from the bank. For wire
	// payments, this will be the OBI field on the wire. For check payments, this will
	// be the memo field.
	StatementDescriptor string `json:"statement_descriptor,required,nullable"`
	// One of unreconciled, partially_reconciled, reconciled, or archived.
	Status ExpectedPaymentStatus `json:"status,required"`
	// The ID of the Transaction this expected payment object has been matched to.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the Transaction Line Item this expected payment has been matched to.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type      ExpectedPaymentType `json:"type,required,nullable"`
	UpdatedAt time.Time           `json:"updated_at,required" format:"date-time"`
	JSON      expectedPaymentJSON `json:"-"`
}

// expectedPaymentJSON contains the JSON metadata for the struct [ExpectedPayment]
type expectedPaymentJSON struct {
	ID                          apijson.Field
	AmountLowerBound            apijson.Field
	AmountUpperBound            apijson.Field
	CounterpartyID              apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	DateLowerBound              apijson.Field
	DateUpperBound              apijson.Field
	Description                 apijson.Field
	Direction                   apijson.Field
	InternalAccountID           apijson.Field
	LedgerTransactionID         apijson.Field
	LiveMode                    apijson.Field
	Metadata                    apijson.Field
	Object                      apijson.Field
	ReconciliationFilters       apijson.Field
	ReconciliationGroups        apijson.Field
	ReconciliationMethod        apijson.Field
	ReconciliationRuleVariables apijson.Field
	RemittanceInformation       apijson.Field
	StatementDescriptor         apijson.Field
	Status                      apijson.Field
	TransactionID               apijson.Field
	TransactionLineItemID       apijson.Field
	Type                        apijson.Field
	UpdatedAt                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ExpectedPayment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r expectedPaymentJSON) RawJSON() string {
	return r.raw
}

func (r ExpectedPayment) implementsBulkResultEntity() {}

// One of credit or debit. When you are receiving money, use credit. When you are
// being charged, use debit.
type ExpectedPaymentDirection string

const (
	ExpectedPaymentDirectionCredit ExpectedPaymentDirection = "credit"
	ExpectedPaymentDirectionDebit  ExpectedPaymentDirection = "debit"
)

func (r ExpectedPaymentDirection) IsKnown() bool {
	switch r {
	case ExpectedPaymentDirectionCredit, ExpectedPaymentDirectionDebit:
		return true
	}
	return false
}

// One of manual if this expected payment was manually reconciled in the dashboard,
// automatic if it was automatically reconciled by Modern Treasury, or null if it
// is unreconciled.
type ExpectedPaymentReconciliationMethod string

const (
	ExpectedPaymentReconciliationMethodAutomatic ExpectedPaymentReconciliationMethod = "automatic"
	ExpectedPaymentReconciliationMethodManual    ExpectedPaymentReconciliationMethod = "manual"
)

func (r ExpectedPaymentReconciliationMethod) IsKnown() bool {
	switch r {
	case ExpectedPaymentReconciliationMethodAutomatic, ExpectedPaymentReconciliationMethodManual:
		return true
	}
	return false
}

// One of unreconciled, partially_reconciled, reconciled, or archived.
type ExpectedPaymentStatus string

const (
	ExpectedPaymentStatusArchived            ExpectedPaymentStatus = "archived"
	ExpectedPaymentStatusPartiallyReconciled ExpectedPaymentStatus = "partially_reconciled"
	ExpectedPaymentStatusReconciled          ExpectedPaymentStatus = "reconciled"
	ExpectedPaymentStatusUnreconciled        ExpectedPaymentStatus = "unreconciled"
)

func (r ExpectedPaymentStatus) IsKnown() bool {
	switch r {
	case ExpectedPaymentStatusArchived, ExpectedPaymentStatusPartiallyReconciled, ExpectedPaymentStatusReconciled, ExpectedPaymentStatusUnreconciled:
		return true
	}
	return false
}

// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
// sepa, signet, wire.
type ExpectedPaymentType string

const (
	ExpectedPaymentTypeACH         ExpectedPaymentType = "ach"
	ExpectedPaymentTypeAuBecs      ExpectedPaymentType = "au_becs"
	ExpectedPaymentTypeBacs        ExpectedPaymentType = "bacs"
	ExpectedPaymentTypeBase        ExpectedPaymentType = "base"
	ExpectedPaymentTypeBook        ExpectedPaymentType = "book"
	ExpectedPaymentTypeCard        ExpectedPaymentType = "card"
	ExpectedPaymentTypeChats       ExpectedPaymentType = "chats"
	ExpectedPaymentTypeCheck       ExpectedPaymentType = "check"
	ExpectedPaymentTypeCrossBorder ExpectedPaymentType = "cross_border"
	ExpectedPaymentTypeDkNets      ExpectedPaymentType = "dk_nets"
	ExpectedPaymentTypeEft         ExpectedPaymentType = "eft"
	ExpectedPaymentTypeEthereum    ExpectedPaymentType = "ethereum"
	ExpectedPaymentTypeHuIcs       ExpectedPaymentType = "hu_ics"
	ExpectedPaymentTypeInterac     ExpectedPaymentType = "interac"
	ExpectedPaymentTypeMasav       ExpectedPaymentType = "masav"
	ExpectedPaymentTypeMxCcen      ExpectedPaymentType = "mx_ccen"
	ExpectedPaymentTypeNeft        ExpectedPaymentType = "neft"
	ExpectedPaymentTypeNics        ExpectedPaymentType = "nics"
	ExpectedPaymentTypeNzBecs      ExpectedPaymentType = "nz_becs"
	ExpectedPaymentTypePlElixir    ExpectedPaymentType = "pl_elixir"
	ExpectedPaymentTypePolygon     ExpectedPaymentType = "polygon"
	ExpectedPaymentTypeProvxchange ExpectedPaymentType = "provxchange"
	ExpectedPaymentTypeRoSent      ExpectedPaymentType = "ro_sent"
	ExpectedPaymentTypeRtp         ExpectedPaymentType = "rtp"
	ExpectedPaymentTypeSeBankgirot ExpectedPaymentType = "se_bankgirot"
	ExpectedPaymentTypeSen         ExpectedPaymentType = "sen"
	ExpectedPaymentTypeSepa        ExpectedPaymentType = "sepa"
	ExpectedPaymentTypeSgGiro      ExpectedPaymentType = "sg_giro"
	ExpectedPaymentTypeSic         ExpectedPaymentType = "sic"
	ExpectedPaymentTypeSignet      ExpectedPaymentType = "signet"
	ExpectedPaymentTypeSknbi       ExpectedPaymentType = "sknbi"
	ExpectedPaymentTypeSolana      ExpectedPaymentType = "solana"
	ExpectedPaymentTypeWire        ExpectedPaymentType = "wire"
	ExpectedPaymentTypeZengin      ExpectedPaymentType = "zengin"
)

func (r ExpectedPaymentType) IsKnown() bool {
	switch r {
	case ExpectedPaymentTypeACH, ExpectedPaymentTypeAuBecs, ExpectedPaymentTypeBacs, ExpectedPaymentTypeBase, ExpectedPaymentTypeBook, ExpectedPaymentTypeCard, ExpectedPaymentTypeChats, ExpectedPaymentTypeCheck, ExpectedPaymentTypeCrossBorder, ExpectedPaymentTypeDkNets, ExpectedPaymentTypeEft, ExpectedPaymentTypeEthereum, ExpectedPaymentTypeHuIcs, ExpectedPaymentTypeInterac, ExpectedPaymentTypeMasav, ExpectedPaymentTypeMxCcen, ExpectedPaymentTypeNeft, ExpectedPaymentTypeNics, ExpectedPaymentTypeNzBecs, ExpectedPaymentTypePlElixir, ExpectedPaymentTypePolygon, ExpectedPaymentTypeProvxchange, ExpectedPaymentTypeRoSent, ExpectedPaymentTypeRtp, ExpectedPaymentTypeSeBankgirot, ExpectedPaymentTypeSen, ExpectedPaymentTypeSepa, ExpectedPaymentTypeSgGiro, ExpectedPaymentTypeSic, ExpectedPaymentTypeSignet, ExpectedPaymentTypeSknbi, ExpectedPaymentTypeSolana, ExpectedPaymentTypeWire, ExpectedPaymentTypeZengin:
		return true
	}
	return false
}

type ReconciliationRule struct {
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound int64 `json:"amount_lower_bound,required"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound int64 `json:"amount_upper_bound,required"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction ReconciliationRuleDirection `json:"direction,required"`
	// The ID of the Internal Account for the expected payment
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// The ID of the counterparty you expect for this payment
	CounterpartyID string `json:"counterparty_id,nullable" format:"uuid"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account
	Currency shared.Currency `json:"currency"`
	// A hash of custom identifiers for this payment
	CustomIdentifiers map[string]string `json:"custom_identifiers,nullable"`
	// The earliest date the payment may come in. Format is yyyy-mm-dd
	DateLowerBound time.Time `json:"date_lower_bound,nullable" format:"date"`
	// The latest date the payment may come in. Format is yyyy-mm-dd
	DateUpperBound time.Time `json:"date_upper_bound,nullable" format:"date"`
	// One of ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet wire
	Type ReconciliationRuleType `json:"type,nullable"`
	JSON reconciliationRuleJSON `json:"-"`
}

// reconciliationRuleJSON contains the JSON metadata for the struct
// [ReconciliationRule]
type reconciliationRuleJSON struct {
	AmountLowerBound  apijson.Field
	AmountUpperBound  apijson.Field
	Direction         apijson.Field
	InternalAccountID apijson.Field
	CounterpartyID    apijson.Field
	Currency          apijson.Field
	CustomIdentifiers apijson.Field
	DateLowerBound    apijson.Field
	DateUpperBound    apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ReconciliationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reconciliationRuleJSON) RawJSON() string {
	return r.raw
}

// One of credit or debit. When you are receiving money, use credit. When you are
// being charged, use debit.
type ReconciliationRuleDirection string

const (
	ReconciliationRuleDirectionCredit ReconciliationRuleDirection = "credit"
	ReconciliationRuleDirectionDebit  ReconciliationRuleDirection = "debit"
)

func (r ReconciliationRuleDirection) IsKnown() bool {
	switch r {
	case ReconciliationRuleDirectionCredit, ReconciliationRuleDirectionDebit:
		return true
	}
	return false
}

// One of ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
// sepa, signet wire
type ReconciliationRuleType string

const (
	ReconciliationRuleTypeACH         ReconciliationRuleType = "ach"
	ReconciliationRuleTypeAuBecs      ReconciliationRuleType = "au_becs"
	ReconciliationRuleTypeBacs        ReconciliationRuleType = "bacs"
	ReconciliationRuleTypeBase        ReconciliationRuleType = "base"
	ReconciliationRuleTypeBook        ReconciliationRuleType = "book"
	ReconciliationRuleTypeCard        ReconciliationRuleType = "card"
	ReconciliationRuleTypeChats       ReconciliationRuleType = "chats"
	ReconciliationRuleTypeCheck       ReconciliationRuleType = "check"
	ReconciliationRuleTypeCrossBorder ReconciliationRuleType = "cross_border"
	ReconciliationRuleTypeDkNets      ReconciliationRuleType = "dk_nets"
	ReconciliationRuleTypeEft         ReconciliationRuleType = "eft"
	ReconciliationRuleTypeEthereum    ReconciliationRuleType = "ethereum"
	ReconciliationRuleTypeHuIcs       ReconciliationRuleType = "hu_ics"
	ReconciliationRuleTypeInterac     ReconciliationRuleType = "interac"
	ReconciliationRuleTypeMasav       ReconciliationRuleType = "masav"
	ReconciliationRuleTypeMxCcen      ReconciliationRuleType = "mx_ccen"
	ReconciliationRuleTypeNeft        ReconciliationRuleType = "neft"
	ReconciliationRuleTypeNics        ReconciliationRuleType = "nics"
	ReconciliationRuleTypeNzBecs      ReconciliationRuleType = "nz_becs"
	ReconciliationRuleTypePlElixir    ReconciliationRuleType = "pl_elixir"
	ReconciliationRuleTypePolygon     ReconciliationRuleType = "polygon"
	ReconciliationRuleTypeProvxchange ReconciliationRuleType = "provxchange"
	ReconciliationRuleTypeRoSent      ReconciliationRuleType = "ro_sent"
	ReconciliationRuleTypeRtp         ReconciliationRuleType = "rtp"
	ReconciliationRuleTypeSeBankgirot ReconciliationRuleType = "se_bankgirot"
	ReconciliationRuleTypeSen         ReconciliationRuleType = "sen"
	ReconciliationRuleTypeSepa        ReconciliationRuleType = "sepa"
	ReconciliationRuleTypeSgGiro      ReconciliationRuleType = "sg_giro"
	ReconciliationRuleTypeSic         ReconciliationRuleType = "sic"
	ReconciliationRuleTypeSignet      ReconciliationRuleType = "signet"
	ReconciliationRuleTypeSknbi       ReconciliationRuleType = "sknbi"
	ReconciliationRuleTypeSolana      ReconciliationRuleType = "solana"
	ReconciliationRuleTypeWire        ReconciliationRuleType = "wire"
	ReconciliationRuleTypeZengin      ReconciliationRuleType = "zengin"
)

func (r ReconciliationRuleType) IsKnown() bool {
	switch r {
	case ReconciliationRuleTypeACH, ReconciliationRuleTypeAuBecs, ReconciliationRuleTypeBacs, ReconciliationRuleTypeBase, ReconciliationRuleTypeBook, ReconciliationRuleTypeCard, ReconciliationRuleTypeChats, ReconciliationRuleTypeCheck, ReconciliationRuleTypeCrossBorder, ReconciliationRuleTypeDkNets, ReconciliationRuleTypeEft, ReconciliationRuleTypeEthereum, ReconciliationRuleTypeHuIcs, ReconciliationRuleTypeInterac, ReconciliationRuleTypeMasav, ReconciliationRuleTypeMxCcen, ReconciliationRuleTypeNeft, ReconciliationRuleTypeNics, ReconciliationRuleTypeNzBecs, ReconciliationRuleTypePlElixir, ReconciliationRuleTypePolygon, ReconciliationRuleTypeProvxchange, ReconciliationRuleTypeRoSent, ReconciliationRuleTypeRtp, ReconciliationRuleTypeSeBankgirot, ReconciliationRuleTypeSen, ReconciliationRuleTypeSepa, ReconciliationRuleTypeSgGiro, ReconciliationRuleTypeSic, ReconciliationRuleTypeSignet, ReconciliationRuleTypeSknbi, ReconciliationRuleTypeSolana, ReconciliationRuleTypeWire, ReconciliationRuleTypeZengin:
		return true
	}
	return false
}

type ReconciliationRuleParam struct {
	// The lowest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountLowerBound param.Field[int64] `json:"amount_lower_bound,required"`
	// The highest amount this expected payment may be equal to. Value in specified
	// currency's smallest unit. e.g. $10 would be represented as 1000.
	AmountUpperBound param.Field[int64] `json:"amount_upper_bound,required"`
	// One of credit or debit. When you are receiving money, use credit. When you are
	// being charged, use debit.
	Direction param.Field[ReconciliationRuleDirection] `json:"direction,required"`
	// The ID of the Internal Account for the expected payment
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// The ID of the counterparty you expect for this payment
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// Must conform to ISO 4217. Defaults to the currency of the internal account
	Currency param.Field[shared.Currency] `json:"currency"`
	// A hash of custom identifiers for this payment
	CustomIdentifiers param.Field[map[string]string] `json:"custom_identifiers"`
	// The earliest date the payment may come in. Format is yyyy-mm-dd
	DateLowerBound param.Field[time.Time] `json:"date_lower_bound" format:"date"`
	// The latest date the payment may come in. Format is yyyy-mm-dd
	DateUpperBound param.Field[time.Time] `json:"date_upper_bound" format:"date"`
	// One of ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet wire
	Type param.Field[ReconciliationRuleType] `json:"type"`
}

func (r ReconciliationRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExpectedPaymentNewParams struct {
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
	Direction param.Field[ExpectedPaymentNewParamsDirection] `json:"direction"`
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
	LedgerTransactionID param.Field[string]                             `json:"ledger_transaction_id" format:"uuid"`
	LineItems           param.Field[[]ExpectedPaymentNewParamsLineItem] `json:"line_items"`
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

func (r ExpectedPaymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of credit or debit. When you are receiving money, use credit. When you are
// being charged, use debit.
type ExpectedPaymentNewParamsDirection string

const (
	ExpectedPaymentNewParamsDirectionCredit ExpectedPaymentNewParamsDirection = "credit"
	ExpectedPaymentNewParamsDirectionDebit  ExpectedPaymentNewParamsDirection = "debit"
)

func (r ExpectedPaymentNewParamsDirection) IsKnown() bool {
	switch r {
	case ExpectedPaymentNewParamsDirectionCredit, ExpectedPaymentNewParamsDirectionDebit:
		return true
	}
	return false
}

type ExpectedPaymentNewParamsLineItem struct {
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

func (r ExpectedPaymentNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExpectedPaymentUpdateParams struct {
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
	Direction param.Field[ExpectedPaymentUpdateParamsDirection] `json:"direction"`
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
	Status param.Field[ExpectedPaymentUpdateParamsStatus] `json:"status"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp, sen,
	// sepa, signet, wire.
	Type param.Field[ExpectedPaymentType] `json:"type"`
}

func (r ExpectedPaymentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of credit or debit. When you are receiving money, use credit. When you are
// being charged, use debit.
type ExpectedPaymentUpdateParamsDirection string

const (
	ExpectedPaymentUpdateParamsDirectionCredit ExpectedPaymentUpdateParamsDirection = "credit"
	ExpectedPaymentUpdateParamsDirectionDebit  ExpectedPaymentUpdateParamsDirection = "debit"
)

func (r ExpectedPaymentUpdateParamsDirection) IsKnown() bool {
	switch r {
	case ExpectedPaymentUpdateParamsDirectionCredit, ExpectedPaymentUpdateParamsDirectionDebit:
		return true
	}
	return false
}

// The Expected Payment's status can be updated from partially_reconciled to
// reconciled.
type ExpectedPaymentUpdateParamsStatus string

const (
	ExpectedPaymentUpdateParamsStatusReconciled ExpectedPaymentUpdateParamsStatus = "reconciled"
)

func (r ExpectedPaymentUpdateParamsStatus) IsKnown() bool {
	switch r {
	case ExpectedPaymentUpdateParamsStatusReconciled:
		return true
	}
	return false
}

type ExpectedPaymentListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Specify counterparty_id to see expected_payments for a specific account.
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// Used to return expected payments created after some datetime
	CreatedAtLowerBound param.Field[time.Time] `query:"created_at_lower_bound" format:"date-time"`
	// Used to return expected payments created before some datetime
	CreatedAtUpperBound param.Field[time.Time] `query:"created_at_upper_bound" format:"date-time"`
	// One of credit, debit
	Direction param.Field[shared.TransactionDirection] `query:"direction"`
	// Specify internal_account_id to see expected_payments for a specific account.
	InternalAccountID param.Field[string] `query:"internal_account_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// One of unreconciled, reconciled, or archived.
	Status param.Field[ExpectedPaymentListParamsStatus] `query:"status"`
	// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp,sen,
	// sepa, signet, wire
	Type param.Field[ExpectedPaymentListParamsType] `query:"type"`
	// Used to return expected payments updated after some datetime
	UpdatedAtLowerBound param.Field[time.Time] `query:"updated_at_lower_bound" format:"date-time"`
	// Used to return expected payments updated before some datetime
	UpdatedAtUpperBound param.Field[time.Time] `query:"updated_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [ExpectedPaymentListParams]'s query parameters as
// `url.Values`.
func (r ExpectedPaymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One of unreconciled, reconciled, or archived.
type ExpectedPaymentListParamsStatus string

const (
	ExpectedPaymentListParamsStatusArchived            ExpectedPaymentListParamsStatus = "archived"
	ExpectedPaymentListParamsStatusPartiallyReconciled ExpectedPaymentListParamsStatus = "partially_reconciled"
	ExpectedPaymentListParamsStatusReconciled          ExpectedPaymentListParamsStatus = "reconciled"
	ExpectedPaymentListParamsStatusUnreconciled        ExpectedPaymentListParamsStatus = "unreconciled"
)

func (r ExpectedPaymentListParamsStatus) IsKnown() bool {
	switch r {
	case ExpectedPaymentListParamsStatusArchived, ExpectedPaymentListParamsStatusPartiallyReconciled, ExpectedPaymentListParamsStatusReconciled, ExpectedPaymentListParamsStatusUnreconciled:
		return true
	}
	return false
}

// One of: ach, au_becs, bacs, book, check, eft, interac, provxchange, rtp,sen,
// sepa, signet, wire
type ExpectedPaymentListParamsType string

const (
	ExpectedPaymentListParamsTypeACH         ExpectedPaymentListParamsType = "ach"
	ExpectedPaymentListParamsTypeAuBecs      ExpectedPaymentListParamsType = "au_becs"
	ExpectedPaymentListParamsTypeBacs        ExpectedPaymentListParamsType = "bacs"
	ExpectedPaymentListParamsTypeBase        ExpectedPaymentListParamsType = "base"
	ExpectedPaymentListParamsTypeBook        ExpectedPaymentListParamsType = "book"
	ExpectedPaymentListParamsTypeCard        ExpectedPaymentListParamsType = "card"
	ExpectedPaymentListParamsTypeChats       ExpectedPaymentListParamsType = "chats"
	ExpectedPaymentListParamsTypeCheck       ExpectedPaymentListParamsType = "check"
	ExpectedPaymentListParamsTypeCrossBorder ExpectedPaymentListParamsType = "cross_border"
	ExpectedPaymentListParamsTypeDkNets      ExpectedPaymentListParamsType = "dk_nets"
	ExpectedPaymentListParamsTypeEft         ExpectedPaymentListParamsType = "eft"
	ExpectedPaymentListParamsTypeEthereum    ExpectedPaymentListParamsType = "ethereum"
	ExpectedPaymentListParamsTypeHuIcs       ExpectedPaymentListParamsType = "hu_ics"
	ExpectedPaymentListParamsTypeInterac     ExpectedPaymentListParamsType = "interac"
	ExpectedPaymentListParamsTypeMasav       ExpectedPaymentListParamsType = "masav"
	ExpectedPaymentListParamsTypeMxCcen      ExpectedPaymentListParamsType = "mx_ccen"
	ExpectedPaymentListParamsTypeNeft        ExpectedPaymentListParamsType = "neft"
	ExpectedPaymentListParamsTypeNics        ExpectedPaymentListParamsType = "nics"
	ExpectedPaymentListParamsTypeNzBecs      ExpectedPaymentListParamsType = "nz_becs"
	ExpectedPaymentListParamsTypePlElixir    ExpectedPaymentListParamsType = "pl_elixir"
	ExpectedPaymentListParamsTypePolygon     ExpectedPaymentListParamsType = "polygon"
	ExpectedPaymentListParamsTypeProvxchange ExpectedPaymentListParamsType = "provxchange"
	ExpectedPaymentListParamsTypeRoSent      ExpectedPaymentListParamsType = "ro_sent"
	ExpectedPaymentListParamsTypeRtp         ExpectedPaymentListParamsType = "rtp"
	ExpectedPaymentListParamsTypeSeBankgirot ExpectedPaymentListParamsType = "se_bankgirot"
	ExpectedPaymentListParamsTypeSen         ExpectedPaymentListParamsType = "sen"
	ExpectedPaymentListParamsTypeSepa        ExpectedPaymentListParamsType = "sepa"
	ExpectedPaymentListParamsTypeSgGiro      ExpectedPaymentListParamsType = "sg_giro"
	ExpectedPaymentListParamsTypeSic         ExpectedPaymentListParamsType = "sic"
	ExpectedPaymentListParamsTypeSignet      ExpectedPaymentListParamsType = "signet"
	ExpectedPaymentListParamsTypeSknbi       ExpectedPaymentListParamsType = "sknbi"
	ExpectedPaymentListParamsTypeSolana      ExpectedPaymentListParamsType = "solana"
	ExpectedPaymentListParamsTypeWire        ExpectedPaymentListParamsType = "wire"
	ExpectedPaymentListParamsTypeZengin      ExpectedPaymentListParamsType = "zengin"
)

func (r ExpectedPaymentListParamsType) IsKnown() bool {
	switch r {
	case ExpectedPaymentListParamsTypeACH, ExpectedPaymentListParamsTypeAuBecs, ExpectedPaymentListParamsTypeBacs, ExpectedPaymentListParamsTypeBase, ExpectedPaymentListParamsTypeBook, ExpectedPaymentListParamsTypeCard, ExpectedPaymentListParamsTypeChats, ExpectedPaymentListParamsTypeCheck, ExpectedPaymentListParamsTypeCrossBorder, ExpectedPaymentListParamsTypeDkNets, ExpectedPaymentListParamsTypeEft, ExpectedPaymentListParamsTypeEthereum, ExpectedPaymentListParamsTypeHuIcs, ExpectedPaymentListParamsTypeInterac, ExpectedPaymentListParamsTypeMasav, ExpectedPaymentListParamsTypeMxCcen, ExpectedPaymentListParamsTypeNeft, ExpectedPaymentListParamsTypeNics, ExpectedPaymentListParamsTypeNzBecs, ExpectedPaymentListParamsTypePlElixir, ExpectedPaymentListParamsTypePolygon, ExpectedPaymentListParamsTypeProvxchange, ExpectedPaymentListParamsTypeRoSent, ExpectedPaymentListParamsTypeRtp, ExpectedPaymentListParamsTypeSeBankgirot, ExpectedPaymentListParamsTypeSen, ExpectedPaymentListParamsTypeSepa, ExpectedPaymentListParamsTypeSgGiro, ExpectedPaymentListParamsTypeSic, ExpectedPaymentListParamsTypeSignet, ExpectedPaymentListParamsTypeSknbi, ExpectedPaymentListParamsTypeSolana, ExpectedPaymentListParamsTypeWire, ExpectedPaymentListParamsTypeZengin:
		return true
	}
	return false
}
