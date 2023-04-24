package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

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
	Currency Currency `json:"currency,required,nullable"`
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
	JSON                ExpectedPaymentJSON
}

type ExpectedPaymentJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	AmountUpperBound      pjson.Metadata
	AmountLowerBound      pjson.Metadata
	Direction             pjson.Metadata
	InternalAccountID     pjson.Metadata
	Type                  pjson.Metadata
	Currency              pjson.Metadata
	DateUpperBound        pjson.Metadata
	DateLowerBound        pjson.Metadata
	Description           pjson.Metadata
	StatementDescriptor   pjson.Metadata
	Metadata              pjson.Metadata
	CounterpartyID        pjson.Metadata
	RemittanceInformation pjson.Metadata
	TransactionID         pjson.Metadata
	TransactionLineItemID pjson.Metadata
	Status                pjson.Metadata
	ReconciliationMethod  pjson.Metadata
	LedgerTransactionID   pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExpectedPayment using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExpectedPayment) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
