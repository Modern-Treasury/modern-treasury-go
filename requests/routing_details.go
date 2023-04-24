package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type RoutingDetail struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    field.Field[bool]      `json:"live_mode,required"`
	CreatedAt   field.Field[time.Time] `json:"created_at,required" format:"date-time"`
	UpdatedAt   field.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	DiscardedAt field.Field[time.Time] `json:"discarded_at,required,nullable" format:"date-time"`
	// The routing number of the bank.
	RoutingNumber field.Field[string] `json:"routing_number,required"`
	// One of `aba`, `swift`, `ca_cpa`, `au_bsb`, `gb_sort_code`, `in_ifsc`, `cnaps`.
	RoutingNumberType field.Field[RoutingDetailRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType field.Field[RoutingDetailPaymentType] `json:"payment_type,required,nullable"`
	// The name of the bank.
	BankName    field.Field[string]                   `json:"bank_name,required"`
	BankAddress field.Field[RoutingDetailBankAddress] `json:"bank_address,required,nullable"`
}

// MarshalJSON serializes RoutingDetail into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r RoutingDetail) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
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
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  field.Field[bool]      `json:"live_mode,required"`
	CreatedAt field.Field[time.Time] `json:"created_at,required" format:"date-time"`
	UpdatedAt field.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	Line1     field.Field[string]    `json:"line1,required,nullable"`
	Line2     field.Field[string]    `json:"line2,required,nullable"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required,nullable"`
	// Region or State.
	Region field.Field[string] `json:"region,required,nullable"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required,nullable"`
}

// MarshalJSON serializes RoutingDetailBankAddress into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r RoutingDetailBankAddress) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type RoutingDetailNewParams struct {
	// The routing number of the bank.
	RoutingNumber field.Field[string] `json:"routing_number,required"`
	// One of `aba`, `swift`, `ca_cpa`, `au_bsb`, `gb_sort_code`, `in_ifsc`, `cnaps`.
	RoutingNumberType field.Field[RoutingDetailNewParamsRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType field.Field[RoutingDetailNewParamsPaymentType] `json:"payment_type,nullable"`
}

// MarshalJSON serializes RoutingDetailNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r RoutingDetailNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type RoutingDetailNewParamsRoutingNumberType string

const (
	RoutingDetailNewParamsRoutingNumberTypeAba          RoutingDetailNewParamsRoutingNumberType = "aba"
	RoutingDetailNewParamsRoutingNumberTypeAuBsb        RoutingDetailNewParamsRoutingNumberType = "au_bsb"
	RoutingDetailNewParamsRoutingNumberTypeBrCodigo     RoutingDetailNewParamsRoutingNumberType = "br_codigo"
	RoutingDetailNewParamsRoutingNumberTypeCaCpa        RoutingDetailNewParamsRoutingNumberType = "ca_cpa"
	RoutingDetailNewParamsRoutingNumberTypeChips        RoutingDetailNewParamsRoutingNumberType = "chips"
	RoutingDetailNewParamsRoutingNumberTypeCnaps        RoutingDetailNewParamsRoutingNumberType = "cnaps"
	RoutingDetailNewParamsRoutingNumberTypeGBSortCode   RoutingDetailNewParamsRoutingNumberType = "gb_sort_code"
	RoutingDetailNewParamsRoutingNumberTypeInIfsc       RoutingDetailNewParamsRoutingNumberType = "in_ifsc"
	RoutingDetailNewParamsRoutingNumberTypeMyBranchCode RoutingDetailNewParamsRoutingNumberType = "my_branch_code"
	RoutingDetailNewParamsRoutingNumberTypeSwift        RoutingDetailNewParamsRoutingNumberType = "swift"
)

type RoutingDetailNewParamsPaymentType string

const (
	RoutingDetailNewParamsPaymentTypeACH         RoutingDetailNewParamsPaymentType = "ach"
	RoutingDetailNewParamsPaymentTypeAuBecs      RoutingDetailNewParamsPaymentType = "au_becs"
	RoutingDetailNewParamsPaymentTypeBacs        RoutingDetailNewParamsPaymentType = "bacs"
	RoutingDetailNewParamsPaymentTypeBook        RoutingDetailNewParamsPaymentType = "book"
	RoutingDetailNewParamsPaymentTypeCard        RoutingDetailNewParamsPaymentType = "card"
	RoutingDetailNewParamsPaymentTypeCheck       RoutingDetailNewParamsPaymentType = "check"
	RoutingDetailNewParamsPaymentTypeCrossBorder RoutingDetailNewParamsPaymentType = "cross_border"
	RoutingDetailNewParamsPaymentTypeEft         RoutingDetailNewParamsPaymentType = "eft"
	RoutingDetailNewParamsPaymentTypeInterac     RoutingDetailNewParamsPaymentType = "interac"
	RoutingDetailNewParamsPaymentTypeMasav       RoutingDetailNewParamsPaymentType = "masav"
	RoutingDetailNewParamsPaymentTypeNeft        RoutingDetailNewParamsPaymentType = "neft"
	RoutingDetailNewParamsPaymentTypeProvxchange RoutingDetailNewParamsPaymentType = "provxchange"
	RoutingDetailNewParamsPaymentTypeRtp         RoutingDetailNewParamsPaymentType = "rtp"
	RoutingDetailNewParamsPaymentTypeSen         RoutingDetailNewParamsPaymentType = "sen"
	RoutingDetailNewParamsPaymentTypeSepa        RoutingDetailNewParamsPaymentType = "sepa"
	RoutingDetailNewParamsPaymentTypeSignet      RoutingDetailNewParamsPaymentType = "signet"
	RoutingDetailNewParamsPaymentTypeWire        RoutingDetailNewParamsPaymentType = "wire"
)

type RoutingDetailNewParamsAccountsType string

const (
	RoutingDetailNewParamsAccountsTypeExternalAccounts RoutingDetailNewParamsAccountsType = "external_accounts"
)

type RoutingDetailListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes RoutingDetailListParams into a url.Values of the query
// parameters associated with this value
func (r RoutingDetailListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type RoutingDetailDeleteParamsAccountsType string

const (
	RoutingDetailDeleteParamsAccountsTypeExternalAccounts RoutingDetailDeleteParamsAccountsType = "external_accounts"
)
