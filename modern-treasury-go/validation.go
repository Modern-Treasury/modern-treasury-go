package moderntreasury

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type ValidationService struct {
	Options []option.RequestOption
}

func NewValidationService(opts ...option.RequestOption) (r *ValidationService) {
	r = &ValidationService{}
	r.Options = opts
	return
}

// Validates the routing number information supplied without creating a routing
// detail
func (r *ValidationService) ValidateRoutingNumber(ctx context.Context, query ValidationValidateRoutingNumberParams, opts ...option.RequestOption) (res *RoutingNumberLookupRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/validations/routing_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RoutingNumberLookupRequest struct {
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number"`
	// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
	// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
	// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
	// respectively.
	RoutingNumberType RoutingNumberLookupRequestRoutingNumberType `json:"routing_number_type"`
	// An array of payment types that are supported for this routing number. This can
	// include `ach`, `wire`, `rtp`, `sepa`, `bacs`, `au_becs` currently.
	SupportedPaymentTypes []RoutingNumberLookupRequestSupportedPaymentTypes `json:"supported_payment_types"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The address of the bank.
	BankAddress RoutingNumberLookupRequestBankAddress `json:"bank_address"`
	// An object containing key-value pairs, each with a sanctions list as the key and
	// a boolean value representing whether the bank is on that particular sanctions
	// list. Currently, this includes eu_con, uk_hmt, us_ofac, and un sanctions lists.
	Sanctions map[string]interface{} `json:"sanctions"`
	JSON      RoutingNumberLookupRequestJSON
}

type RoutingNumberLookupRequestJSON struct {
	RoutingNumber         apijson.Metadata
	RoutingNumberType     apijson.Metadata
	SupportedPaymentTypes apijson.Metadata
	BankName              apijson.Metadata
	BankAddress           apijson.Metadata
	Sanctions             apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingNumberLookupRequest
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RoutingNumberLookupRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutingNumberLookupRequestRoutingNumberType string

const (
	RoutingNumberLookupRequestRoutingNumberTypeAba        RoutingNumberLookupRequestRoutingNumberType = "aba"
	RoutingNumberLookupRequestRoutingNumberTypeAuBsb      RoutingNumberLookupRequestRoutingNumberType = "au_bsb"
	RoutingNumberLookupRequestRoutingNumberTypeCaCpa      RoutingNumberLookupRequestRoutingNumberType = "ca_cpa"
	RoutingNumberLookupRequestRoutingNumberTypeGBSortCode RoutingNumberLookupRequestRoutingNumberType = "gb_sort_code"
	RoutingNumberLookupRequestRoutingNumberTypeInIfsc     RoutingNumberLookupRequestRoutingNumberType = "in_ifsc"
	RoutingNumberLookupRequestRoutingNumberTypeSwift      RoutingNumberLookupRequestRoutingNumberType = "swift"
)

type RoutingNumberLookupRequestSupportedPaymentTypes string

const (
	RoutingNumberLookupRequestSupportedPaymentTypesACH         RoutingNumberLookupRequestSupportedPaymentTypes = "ach"
	RoutingNumberLookupRequestSupportedPaymentTypesAuBecs      RoutingNumberLookupRequestSupportedPaymentTypes = "au_becs"
	RoutingNumberLookupRequestSupportedPaymentTypesBacs        RoutingNumberLookupRequestSupportedPaymentTypes = "bacs"
	RoutingNumberLookupRequestSupportedPaymentTypesBook        RoutingNumberLookupRequestSupportedPaymentTypes = "book"
	RoutingNumberLookupRequestSupportedPaymentTypesCard        RoutingNumberLookupRequestSupportedPaymentTypes = "card"
	RoutingNumberLookupRequestSupportedPaymentTypesCheck       RoutingNumberLookupRequestSupportedPaymentTypes = "check"
	RoutingNumberLookupRequestSupportedPaymentTypesCrossBorder RoutingNumberLookupRequestSupportedPaymentTypes = "cross_border"
	RoutingNumberLookupRequestSupportedPaymentTypesEft         RoutingNumberLookupRequestSupportedPaymentTypes = "eft"
	RoutingNumberLookupRequestSupportedPaymentTypesInterac     RoutingNumberLookupRequestSupportedPaymentTypes = "interac"
	RoutingNumberLookupRequestSupportedPaymentTypesMasav       RoutingNumberLookupRequestSupportedPaymentTypes = "masav"
	RoutingNumberLookupRequestSupportedPaymentTypesNeft        RoutingNumberLookupRequestSupportedPaymentTypes = "neft"
	RoutingNumberLookupRequestSupportedPaymentTypesProvxchange RoutingNumberLookupRequestSupportedPaymentTypes = "provxchange"
	RoutingNumberLookupRequestSupportedPaymentTypesRtp         RoutingNumberLookupRequestSupportedPaymentTypes = "rtp"
	RoutingNumberLookupRequestSupportedPaymentTypesSen         RoutingNumberLookupRequestSupportedPaymentTypes = "sen"
	RoutingNumberLookupRequestSupportedPaymentTypesSepa        RoutingNumberLookupRequestSupportedPaymentTypes = "sepa"
	RoutingNumberLookupRequestSupportedPaymentTypesSignet      RoutingNumberLookupRequestSupportedPaymentTypes = "signet"
	RoutingNumberLookupRequestSupportedPaymentTypesWire        RoutingNumberLookupRequestSupportedPaymentTypes = "wire"
)

type RoutingNumberLookupRequestBankAddress struct {
	Line1 string `json:"line1,nullable"`
	Line2 string `json:"line2,nullable"`
	// Locality or City.
	Locality string `json:"locality,nullable"`
	// Region or State.
	Region string `json:"region,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,nullable"`
	JSON    RoutingNumberLookupRequestBankAddressJSON
}

type RoutingNumberLookupRequestBankAddressJSON struct {
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// RoutingNumberLookupRequestBankAddress using the internal json library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *RoutingNumberLookupRequestBankAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ValidationValidateRoutingNumberParams struct {
	// The routing number that is being validated.
	RoutingNumber field.Field[string] `query:"routing_number,required"`
	// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
	// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
	// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
	// respectively.
	RoutingNumberType field.Field[ValidationValidateRoutingNumberParamsRoutingNumberType] `query:"routing_number_type,required"`
}

// URLQuery serializes ValidationValidateRoutingNumberParams into a url.Values of
// the query parameters associated with this value
func (r ValidationValidateRoutingNumberParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type ValidationValidateRoutingNumberParamsRoutingNumberType string

const (
	ValidationValidateRoutingNumberParamsRoutingNumberTypeAba          ValidationValidateRoutingNumberParamsRoutingNumberType = "aba"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeAuBsb        ValidationValidateRoutingNumberParamsRoutingNumberType = "au_bsb"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeBrCodigo     ValidationValidateRoutingNumberParamsRoutingNumberType = "br_codigo"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeCaCpa        ValidationValidateRoutingNumberParamsRoutingNumberType = "ca_cpa"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeChips        ValidationValidateRoutingNumberParamsRoutingNumberType = "chips"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeCnaps        ValidationValidateRoutingNumberParamsRoutingNumberType = "cnaps"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeGBSortCode   ValidationValidateRoutingNumberParamsRoutingNumberType = "gb_sort_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeInIfsc       ValidationValidateRoutingNumberParamsRoutingNumberType = "in_ifsc"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeMyBranchCode ValidationValidateRoutingNumberParamsRoutingNumberType = "my_branch_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeSwift        ValidationValidateRoutingNumberParamsRoutingNumberType = "swift"
)
