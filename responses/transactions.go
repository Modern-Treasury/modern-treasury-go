package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type Transaction struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// Currency that this transaction is denominated in.
	Currency Currency `json:"currency,required,nullable"`
	// Either `credit` or `debit`.
	Direction string `json:"direction,required"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription string `json:"vendor_description,required,nullable"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode string `json:"vendor_code,required,nullable"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, or `us_bank`.
	VendorCodeType TransactionVendorCodeType `json:"vendor_code_type,required,nullable"`
	// An identifier given to this transaction by the bank.
	VendorID string `json:"vendor_id,required,nullable"`
	// The date on which the transaction occurred.
	AsOfDate time.Time `json:"as_of_date,required,nullable" format:"date"`
	// The time on which the transaction occurred. Depending on the granularity of the
	// timestamp information received from the bank, it may be `null`.
	AsOfTime string `json:"as_of_time,required,nullable" format:"time"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// This field will be `true` if the transaction has posted to the account.
	Posted bool `json:"posted,required"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable"`
	// This field will be `true` if a transaction is reconciled by the Modern Treasury
	// system. This means that it has transaction line items that sum up to the
	// transaction's amount.
	Reconciled bool `json:"reconciled,required"`
	// This field contains additional information that the bank provided about the
	// transaction. This is structured data. Some of the data in here might overlap
	// with what is in the `vendor_description`. For example, the OBI could be a part
	// of the vendor description, and it would also be included in here. The attributes
	// that are passed through the details field will vary based on your banking
	// partner. Currently, the following keys may be in the details object:
	// `originator_name`, `originator_to_beneficiary_information`.
	Details map[string]string `json:"details,required"`
	// The type of the transaction. Can be one of `ach`, `wire`, `check`, `rtp`,
	// `book`, or `sen`.
	Type TransactionType `json:"type,required"`
	JSON TransactionJSON
}

type TransactionJSON struct {
	ID                pjson.Metadata
	Object            pjson.Metadata
	LiveMode          pjson.Metadata
	CreatedAt         pjson.Metadata
	UpdatedAt         pjson.Metadata
	DiscardedAt       pjson.Metadata
	Amount            pjson.Metadata
	Currency          pjson.Metadata
	Direction         pjson.Metadata
	VendorDescription pjson.Metadata
	VendorCode        pjson.Metadata
	VendorCodeType    pjson.Metadata
	VendorID          pjson.Metadata
	AsOfDate          pjson.Metadata
	AsOfTime          pjson.Metadata
	InternalAccountID pjson.Metadata
	Metadata          pjson.Metadata
	Posted            pjson.Metadata
	VendorCustomerID  pjson.Metadata
	Reconciled        pjson.Metadata
	Details           pjson.Metadata
	Type              pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Transaction using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TransactionVendorCodeType string

const (
	TransactionVendorCodeTypeBai2          TransactionVendorCodeType = "bai2"
	TransactionVendorCodeTypeBankprov      TransactionVendorCodeType = "bankprov"
	TransactionVendorCodeTypeBnkDev        TransactionVendorCodeType = "bnk_dev"
	TransactionVendorCodeTypeCleartouch    TransactionVendorCodeType = "cleartouch"
	TransactionVendorCodeTypeColumn        TransactionVendorCodeType = "column"
	TransactionVendorCodeTypeCrossRiver    TransactionVendorCodeType = "cross_river"
	TransactionVendorCodeTypeCurrencycloud TransactionVendorCodeType = "currencycloud"
	TransactionVendorCodeTypeDcBank        TransactionVendorCodeType = "dc_bank"
	TransactionVendorCodeTypeDwolla        TransactionVendorCodeType = "dwolla"
	TransactionVendorCodeTypeEvolve        TransactionVendorCodeType = "evolve"
	TransactionVendorCodeTypeGoldmanSachs  TransactionVendorCodeType = "goldman_sachs"
	TransactionVendorCodeTypeIso20022      TransactionVendorCodeType = "iso20022"
	TransactionVendorCodeTypeJpmc          TransactionVendorCodeType = "jpmc"
	TransactionVendorCodeTypeMx            TransactionVendorCodeType = "mx"
	TransactionVendorCodeTypeRspecVendor   TransactionVendorCodeType = "rspec_vendor"
	TransactionVendorCodeTypeSignet        TransactionVendorCodeType = "signet"
	TransactionVendorCodeTypeSilvergate    TransactionVendorCodeType = "silvergate"
	TransactionVendorCodeTypeSwift         TransactionVendorCodeType = "swift"
	TransactionVendorCodeTypeUsBank        TransactionVendorCodeType = "us_bank"
)

type TransactionType string

const (
	TransactionTypeACH         TransactionType = "ach"
	TransactionTypeAuBecs      TransactionType = "au_becs"
	TransactionTypeBacs        TransactionType = "bacs"
	TransactionTypeBook        TransactionType = "book"
	TransactionTypeCard        TransactionType = "card"
	TransactionTypeCheck       TransactionType = "check"
	TransactionTypeCrossBorder TransactionType = "cross_border"
	TransactionTypeEft         TransactionType = "eft"
	TransactionTypeInterac     TransactionType = "interac"
	TransactionTypeMasav       TransactionType = "masav"
	TransactionTypeNeft        TransactionType = "neft"
	TransactionTypeProvxchange TransactionType = "provxchange"
	TransactionTypeRtp         TransactionType = "rtp"
	TransactionTypeSen         TransactionType = "sen"
	TransactionTypeSepa        TransactionType = "sepa"
	TransactionTypeSignet      TransactionType = "signet"
	TransactionTypeWire        TransactionType = "wire"
)
