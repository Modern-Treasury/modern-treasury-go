package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type LedgerTransaction struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status LedgerTransactionStatus `json:"status,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt time.Time `json:"effective_at,required" format:"date"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerEntry `json:"ledger_entries,required"`
	// The time on which the ledger transaction posted. This is null if the ledger
	// transaction is pending.
	PostedAt string `json:"posted_at,required,nullable" format:"time"`
	// The ID of the ledger this ledger transaction belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType LedgerTransactionLedgerableType `json:"ledgerable_type,required,nullable"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID string `json:"external_id,required,nullable"`
	JSON       LedgerTransactionJSON
}

type LedgerTransactionJSON struct {
	ID             pjson.Metadata
	Object         pjson.Metadata
	LiveMode       pjson.Metadata
	CreatedAt      pjson.Metadata
	UpdatedAt      pjson.Metadata
	Description    pjson.Metadata
	Status         pjson.Metadata
	Metadata       pjson.Metadata
	EffectiveAt    pjson.Metadata
	EffectiveDate  pjson.Metadata
	LedgerEntries  pjson.Metadata
	PostedAt       pjson.Metadata
	LedgerID       pjson.Metadata
	LedgerableType pjson.Metadata
	LedgerableID   pjson.Metadata
	ExternalID     pjson.Metadata
	Raw            []byte
	Extras         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LedgerTransaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LedgerTransaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type LedgerTransactionStatus string

const (
	LedgerTransactionStatusArchived LedgerTransactionStatus = "archived"
	LedgerTransactionStatusPending  LedgerTransactionStatus = "pending"
	LedgerTransactionStatusPosted   LedgerTransactionStatus = "posted"
)

type LedgerTransactionLedgerableType string

const (
	LedgerTransactionLedgerableTypeCounterparty          LedgerTransactionLedgerableType = "counterparty"
	LedgerTransactionLedgerableTypeExpectedPayment       LedgerTransactionLedgerableType = "expected_payment"
	LedgerTransactionLedgerableTypeIncomingPaymentDetail LedgerTransactionLedgerableType = "incoming_payment_detail"
	LedgerTransactionLedgerableTypeInternalAccount       LedgerTransactionLedgerableType = "internal_account"
	LedgerTransactionLedgerableTypeLineItem              LedgerTransactionLedgerableType = "line_item"
	LedgerTransactionLedgerableTypePaperItem             LedgerTransactionLedgerableType = "paper_item"
	LedgerTransactionLedgerableTypePaymentOrder          LedgerTransactionLedgerableType = "payment_order"
	LedgerTransactionLedgerableTypePaymentOrderAttempt   LedgerTransactionLedgerableType = "payment_order_attempt"
	LedgerTransactionLedgerableTypeReturn                LedgerTransactionLedgerableType = "return"
	LedgerTransactionLedgerableTypeReversal              LedgerTransactionLedgerableType = "reversal"
)
