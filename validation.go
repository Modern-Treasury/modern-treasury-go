// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// ValidationService contains methods and other services that help with interacting
// with the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewValidationService] method instead.
type ValidationService struct {
	Options []option.RequestOption
}

// NewValidationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
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
	// The address of the bank.
	BankAddress RoutingNumberLookupRequestBankAddress `json:"bank_address"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number"`
	// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
	// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
	// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
	// respectively.
	RoutingNumberType RoutingNumberLookupRequestRoutingNumberType `json:"routing_number_type"`
	// An object containing key-value pairs, each with a sanctions list as the key and
	// a boolean value representing whether the bank is on that particular sanctions
	// list. Currently, this includes eu_con, uk_hmt, us_ofac, and un sanctions lists.
	Sanctions map[string]interface{} `json:"sanctions"`
	// An array of payment types that are supported for this routing number. This can
	// include `ach`, `wire`, `rtp`, `sepa`, `bacs`, `au_becs` currently.
	SupportedPaymentTypes []RoutingNumberLookupRequestSupportedPaymentType `json:"supported_payment_types"`
	JSON                  routingNumberLookupRequestJSON
}

// routingNumberLookupRequestJSON contains the JSON metadata for the struct
// [RoutingNumberLookupRequest]
type routingNumberLookupRequestJSON struct {
	BankAddress           apijson.Field
	BankName              apijson.Field
	RoutingNumber         apijson.Field
	RoutingNumberType     apijson.Field
	Sanctions             apijson.Field
	SupportedPaymentTypes apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RoutingNumberLookupRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The address of the bank.
type RoutingNumberLookupRequestBankAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,nullable"`
	Line1   string `json:"line1,nullable"`
	Line2   string `json:"line2,nullable"`
	// Locality or City.
	Locality string `json:"locality,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,nullable"`
	// Region or State.
	Region string `json:"region,nullable"`
	JSON   routingNumberLookupRequestBankAddressJSON
}

// routingNumberLookupRequestBankAddressJSON contains the JSON metadata for the
// struct [RoutingNumberLookupRequestBankAddress]
type routingNumberLookupRequestBankAddressJSON struct {
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingNumberLookupRequestBankAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
// respectively.
type RoutingNumberLookupRequestRoutingNumberType string

const (
	RoutingNumberLookupRequestRoutingNumberTypeAba        RoutingNumberLookupRequestRoutingNumberType = "aba"
	RoutingNumberLookupRequestRoutingNumberTypeAuBsb      RoutingNumberLookupRequestRoutingNumberType = "au_bsb"
	RoutingNumberLookupRequestRoutingNumberTypeCaCpa      RoutingNumberLookupRequestRoutingNumberType = "ca_cpa"
	RoutingNumberLookupRequestRoutingNumberTypeGBSortCode RoutingNumberLookupRequestRoutingNumberType = "gb_sort_code"
	RoutingNumberLookupRequestRoutingNumberTypeInIfsc     RoutingNumberLookupRequestRoutingNumberType = "in_ifsc"
	RoutingNumberLookupRequestRoutingNumberTypeSwift      RoutingNumberLookupRequestRoutingNumberType = "swift"
)

type RoutingNumberLookupRequestSupportedPaymentType string

const (
	RoutingNumberLookupRequestSupportedPaymentTypeACH         RoutingNumberLookupRequestSupportedPaymentType = "ach"
	RoutingNumberLookupRequestSupportedPaymentTypeAuBecs      RoutingNumberLookupRequestSupportedPaymentType = "au_becs"
	RoutingNumberLookupRequestSupportedPaymentTypeBacs        RoutingNumberLookupRequestSupportedPaymentType = "bacs"
	RoutingNumberLookupRequestSupportedPaymentTypeBook        RoutingNumberLookupRequestSupportedPaymentType = "book"
	RoutingNumberLookupRequestSupportedPaymentTypeCard        RoutingNumberLookupRequestSupportedPaymentType = "card"
	RoutingNumberLookupRequestSupportedPaymentTypeCheck       RoutingNumberLookupRequestSupportedPaymentType = "check"
	RoutingNumberLookupRequestSupportedPaymentTypeCrossBorder RoutingNumberLookupRequestSupportedPaymentType = "cross_border"
	RoutingNumberLookupRequestSupportedPaymentTypeEft         RoutingNumberLookupRequestSupportedPaymentType = "eft"
	RoutingNumberLookupRequestSupportedPaymentTypeInterac     RoutingNumberLookupRequestSupportedPaymentType = "interac"
	RoutingNumberLookupRequestSupportedPaymentTypeMasav       RoutingNumberLookupRequestSupportedPaymentType = "masav"
	RoutingNumberLookupRequestSupportedPaymentTypeNeft        RoutingNumberLookupRequestSupportedPaymentType = "neft"
	RoutingNumberLookupRequestSupportedPaymentTypeProvxchange RoutingNumberLookupRequestSupportedPaymentType = "provxchange"
	RoutingNumberLookupRequestSupportedPaymentTypeRtp         RoutingNumberLookupRequestSupportedPaymentType = "rtp"
	RoutingNumberLookupRequestSupportedPaymentTypeSen         RoutingNumberLookupRequestSupportedPaymentType = "sen"
	RoutingNumberLookupRequestSupportedPaymentTypeSepa        RoutingNumberLookupRequestSupportedPaymentType = "sepa"
	RoutingNumberLookupRequestSupportedPaymentTypeSignet      RoutingNumberLookupRequestSupportedPaymentType = "signet"
	RoutingNumberLookupRequestSupportedPaymentTypeWire        RoutingNumberLookupRequestSupportedPaymentType = "wire"
)

type ValidationValidateRoutingNumberParams struct {
	// The routing number that is being validated.
	RoutingNumber param.Field[string] `query:"routing_number,required"`
	// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
	// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
	// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
	// respectively.
	RoutingNumberType param.Field[ValidationValidateRoutingNumberParamsRoutingNumberType] `query:"routing_number_type,required"`
}

// URLQuery serializes [ValidationValidateRoutingNumberParams]'s query parameters
// as `url.Values`.
func (r ValidationValidateRoutingNumberParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// One of `aba`, `au_bsb`, `br_codigo`, `ca_cpa`, `cnaps`, `gb_sort_code`,
// `in_ifsc`, `my_branch_code`, or `swift`. In sandbox mode we currently only
// support `aba` and `swift` with routing numbers '123456789' and 'GRINUST0XXX'
// respectively.
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
