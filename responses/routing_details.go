package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type RoutingDetail struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number,required"`
	// One of `aba`, `swift`, `ca_cpa`, `au_bsb`, `gb_sort_code`, `in_ifsc`, `cnaps`.
	RoutingNumberType RoutingDetailRoutingNumberType `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType RoutingDetailPaymentType `json:"payment_type,required,nullable"`
	// The name of the bank.
	BankName    string                   `json:"bank_name,required"`
	BankAddress RoutingDetailBankAddress `json:"bank_address,required,nullable"`
	JSON        RoutingDetailJSON
}

type RoutingDetailJSON struct {
	ID                pjson.Metadata
	Object            pjson.Metadata
	LiveMode          pjson.Metadata
	CreatedAt         pjson.Metadata
	UpdatedAt         pjson.Metadata
	DiscardedAt       pjson.Metadata
	RoutingNumber     pjson.Metadata
	RoutingNumberType pjson.Metadata
	PaymentType       pjson.Metadata
	BankName          pjson.Metadata
	BankAddress       pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingDetail using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingDetail) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RoutingDetailRoutingNumberType string

const (
	RoutingDetailRoutingNumberTypeAba          RoutingDetailRoutingNumberType = "aba"
	RoutingDetailRoutingNumberTypeAuBsb        RoutingDetailRoutingNumberType = "au_bsb"
	RoutingDetailRoutingNumberTypeBrCodigo     RoutingDetailRoutingNumberType = "br_codigo"
	RoutingDetailRoutingNumberTypeCaCpa        RoutingDetailRoutingNumberType = "ca_cpa"
	RoutingDetailRoutingNumberTypeChips        RoutingDetailRoutingNumberType = "chips"
	RoutingDetailRoutingNumberTypeCnaps        RoutingDetailRoutingNumberType = "cnaps"
	RoutingDetailRoutingNumberTypeGBSortCode   RoutingDetailRoutingNumberType = "gb_sort_code"
	RoutingDetailRoutingNumberTypeInIfsc       RoutingDetailRoutingNumberType = "in_ifsc"
	RoutingDetailRoutingNumberTypeMyBranchCode RoutingDetailRoutingNumberType = "my_branch_code"
	RoutingDetailRoutingNumberTypeSwift        RoutingDetailRoutingNumberType = "swift"
)

type RoutingDetailPaymentType string

const (
	RoutingDetailPaymentTypeACH         RoutingDetailPaymentType = "ach"
	RoutingDetailPaymentTypeAuBecs      RoutingDetailPaymentType = "au_becs"
	RoutingDetailPaymentTypeBacs        RoutingDetailPaymentType = "bacs"
	RoutingDetailPaymentTypeBook        RoutingDetailPaymentType = "book"
	RoutingDetailPaymentTypeCard        RoutingDetailPaymentType = "card"
	RoutingDetailPaymentTypeCheck       RoutingDetailPaymentType = "check"
	RoutingDetailPaymentTypeCrossBorder RoutingDetailPaymentType = "cross_border"
	RoutingDetailPaymentTypeEft         RoutingDetailPaymentType = "eft"
	RoutingDetailPaymentTypeInterac     RoutingDetailPaymentType = "interac"
	RoutingDetailPaymentTypeMasav       RoutingDetailPaymentType = "masav"
	RoutingDetailPaymentTypeNeft        RoutingDetailPaymentType = "neft"
	RoutingDetailPaymentTypeProvxchange RoutingDetailPaymentType = "provxchange"
	RoutingDetailPaymentTypeRtp         RoutingDetailPaymentType = "rtp"
	RoutingDetailPaymentTypeSen         RoutingDetailPaymentType = "sen"
	RoutingDetailPaymentTypeSepa        RoutingDetailPaymentType = "sepa"
	RoutingDetailPaymentTypeSignet      RoutingDetailPaymentType = "signet"
	RoutingDetailPaymentTypeWire        RoutingDetailPaymentType = "wire"
)

type RoutingDetailBankAddress struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	// Region or State.
	Region string `json:"region,required,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required,nullable"`
	JSON    RoutingDetailBankAddressJSON
}

type RoutingDetailBankAddressJSON struct {
	ID         pjson.Metadata
	Object     pjson.Metadata
	LiveMode   pjson.Metadata
	CreatedAt  pjson.Metadata
	UpdatedAt  pjson.Metadata
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	Locality   pjson.Metadata
	Region     pjson.Metadata
	PostalCode pjson.Metadata
	Country    pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingDetailBankAddress
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RoutingDetailBankAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
