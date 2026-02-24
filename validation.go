// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// ValidationService contains methods and other services that help with interacting
// with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewValidationService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	path := "api/validations/routing_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RoutingNumberLookupRequest struct {
	// The address of the bank.
	BankAddress shared.AddressRequest `json:"bank_address"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details. In sandbox mode we currently only support `aba` and `swift` with
	// routing numbers '123456789' and 'GRINUST0XXX' respectively.
	RoutingNumberType RoutingNumberLookupRequestRoutingNumberType `json:"routing_number_type"`
	// An object containing key-value pairs, each with a sanctions list as the key and
	// a boolean value representing whether the bank is on that particular sanctions
	// list. Currently, this includes eu_con, uk_hmt, us_ofac, and un sanctions lists.
	Sanctions map[string]interface{} `json:"sanctions"`
	// An array of payment types that are supported for this routing number. This can
	// include `ach`, `wire`, `rtp`, `sepa`, `bacs`, `au_becs`, and 'fednow' currently.
	SupportedPaymentTypes []RoutingNumberLookupRequestSupportedPaymentType `json:"supported_payment_types"`
	JSON                  routingNumberLookupRequestJSON                   `json:"-"`
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

func (r routingNumberLookupRequestJSON) RawJSON() string {
	return r.raw
}

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details. In sandbox mode we currently only support `aba` and `swift` with
// routing numbers '123456789' and 'GRINUST0XXX' respectively.
type RoutingNumberLookupRequestRoutingNumberType string

const (
	RoutingNumberLookupRequestRoutingNumberTypeAba                    RoutingNumberLookupRequestRoutingNumberType = "aba"
	RoutingNumberLookupRequestRoutingNumberTypeAuBsb                  RoutingNumberLookupRequestRoutingNumberType = "au_bsb"
	RoutingNumberLookupRequestRoutingNumberTypeCaCpa                  RoutingNumberLookupRequestRoutingNumberType = "ca_cpa"
	RoutingNumberLookupRequestRoutingNumberTypeGBSortCode             RoutingNumberLookupRequestRoutingNumberType = "gb_sort_code"
	RoutingNumberLookupRequestRoutingNumberTypeInIfsc                 RoutingNumberLookupRequestRoutingNumberType = "in_ifsc"
	RoutingNumberLookupRequestRoutingNumberTypeNzNationalClearingCode RoutingNumberLookupRequestRoutingNumberType = "nz_national_clearing_code"
	RoutingNumberLookupRequestRoutingNumberTypeSeBankgiroClearingCode RoutingNumberLookupRequestRoutingNumberType = "se_bankgiro_clearing_code"
	RoutingNumberLookupRequestRoutingNumberTypeSwift                  RoutingNumberLookupRequestRoutingNumberType = "swift"
	RoutingNumberLookupRequestRoutingNumberTypeZaNationalClearingCode RoutingNumberLookupRequestRoutingNumberType = "za_national_clearing_code"
)

func (r RoutingNumberLookupRequestRoutingNumberType) IsKnown() bool {
	switch r {
	case RoutingNumberLookupRequestRoutingNumberTypeAba, RoutingNumberLookupRequestRoutingNumberTypeAuBsb, RoutingNumberLookupRequestRoutingNumberTypeCaCpa, RoutingNumberLookupRequestRoutingNumberTypeGBSortCode, RoutingNumberLookupRequestRoutingNumberTypeInIfsc, RoutingNumberLookupRequestRoutingNumberTypeNzNationalClearingCode, RoutingNumberLookupRequestRoutingNumberTypeSeBankgiroClearingCode, RoutingNumberLookupRequestRoutingNumberTypeSwift, RoutingNumberLookupRequestRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type RoutingNumberLookupRequestSupportedPaymentType string

const (
	RoutingNumberLookupRequestSupportedPaymentTypeACH         RoutingNumberLookupRequestSupportedPaymentType = "ach"
	RoutingNumberLookupRequestSupportedPaymentTypeAuBecs      RoutingNumberLookupRequestSupportedPaymentType = "au_becs"
	RoutingNumberLookupRequestSupportedPaymentTypeBacs        RoutingNumberLookupRequestSupportedPaymentType = "bacs"
	RoutingNumberLookupRequestSupportedPaymentTypeBase        RoutingNumberLookupRequestSupportedPaymentType = "base"
	RoutingNumberLookupRequestSupportedPaymentTypeBook        RoutingNumberLookupRequestSupportedPaymentType = "book"
	RoutingNumberLookupRequestSupportedPaymentTypeCard        RoutingNumberLookupRequestSupportedPaymentType = "card"
	RoutingNumberLookupRequestSupportedPaymentTypeChats       RoutingNumberLookupRequestSupportedPaymentType = "chats"
	RoutingNumberLookupRequestSupportedPaymentTypeCheck       RoutingNumberLookupRequestSupportedPaymentType = "check"
	RoutingNumberLookupRequestSupportedPaymentTypeCrossBorder RoutingNumberLookupRequestSupportedPaymentType = "cross_border"
	RoutingNumberLookupRequestSupportedPaymentTypeDkNets      RoutingNumberLookupRequestSupportedPaymentType = "dk_nets"
	RoutingNumberLookupRequestSupportedPaymentTypeEft         RoutingNumberLookupRequestSupportedPaymentType = "eft"
	RoutingNumberLookupRequestSupportedPaymentTypeEthereum    RoutingNumberLookupRequestSupportedPaymentType = "ethereum"
	RoutingNumberLookupRequestSupportedPaymentTypeGBFps       RoutingNumberLookupRequestSupportedPaymentType = "gb_fps"
	RoutingNumberLookupRequestSupportedPaymentTypeHuIcs       RoutingNumberLookupRequestSupportedPaymentType = "hu_ics"
	RoutingNumberLookupRequestSupportedPaymentTypeInterac     RoutingNumberLookupRequestSupportedPaymentType = "interac"
	RoutingNumberLookupRequestSupportedPaymentTypeMasav       RoutingNumberLookupRequestSupportedPaymentType = "masav"
	RoutingNumberLookupRequestSupportedPaymentTypeMxCcen      RoutingNumberLookupRequestSupportedPaymentType = "mx_ccen"
	RoutingNumberLookupRequestSupportedPaymentTypeNeft        RoutingNumberLookupRequestSupportedPaymentType = "neft"
	RoutingNumberLookupRequestSupportedPaymentTypeNics        RoutingNumberLookupRequestSupportedPaymentType = "nics"
	RoutingNumberLookupRequestSupportedPaymentTypeNzBecs      RoutingNumberLookupRequestSupportedPaymentType = "nz_becs"
	RoutingNumberLookupRequestSupportedPaymentTypePlElixir    RoutingNumberLookupRequestSupportedPaymentType = "pl_elixir"
	RoutingNumberLookupRequestSupportedPaymentTypePolygon     RoutingNumberLookupRequestSupportedPaymentType = "polygon"
	RoutingNumberLookupRequestSupportedPaymentTypeProvxchange RoutingNumberLookupRequestSupportedPaymentType = "provxchange"
	RoutingNumberLookupRequestSupportedPaymentTypeRoSent      RoutingNumberLookupRequestSupportedPaymentType = "ro_sent"
	RoutingNumberLookupRequestSupportedPaymentTypeRtp         RoutingNumberLookupRequestSupportedPaymentType = "rtp"
	RoutingNumberLookupRequestSupportedPaymentTypeSeBankgirot RoutingNumberLookupRequestSupportedPaymentType = "se_bankgirot"
	RoutingNumberLookupRequestSupportedPaymentTypeSen         RoutingNumberLookupRequestSupportedPaymentType = "sen"
	RoutingNumberLookupRequestSupportedPaymentTypeSepa        RoutingNumberLookupRequestSupportedPaymentType = "sepa"
	RoutingNumberLookupRequestSupportedPaymentTypeSgGiro      RoutingNumberLookupRequestSupportedPaymentType = "sg_giro"
	RoutingNumberLookupRequestSupportedPaymentTypeSic         RoutingNumberLookupRequestSupportedPaymentType = "sic"
	RoutingNumberLookupRequestSupportedPaymentTypeSignet      RoutingNumberLookupRequestSupportedPaymentType = "signet"
	RoutingNumberLookupRequestSupportedPaymentTypeSknbi       RoutingNumberLookupRequestSupportedPaymentType = "sknbi"
	RoutingNumberLookupRequestSupportedPaymentTypeSolana      RoutingNumberLookupRequestSupportedPaymentType = "solana"
	RoutingNumberLookupRequestSupportedPaymentTypeWire        RoutingNumberLookupRequestSupportedPaymentType = "wire"
	RoutingNumberLookupRequestSupportedPaymentTypeZengin      RoutingNumberLookupRequestSupportedPaymentType = "zengin"
)

func (r RoutingNumberLookupRequestSupportedPaymentType) IsKnown() bool {
	switch r {
	case RoutingNumberLookupRequestSupportedPaymentTypeACH, RoutingNumberLookupRequestSupportedPaymentTypeAuBecs, RoutingNumberLookupRequestSupportedPaymentTypeBacs, RoutingNumberLookupRequestSupportedPaymentTypeBase, RoutingNumberLookupRequestSupportedPaymentTypeBook, RoutingNumberLookupRequestSupportedPaymentTypeCard, RoutingNumberLookupRequestSupportedPaymentTypeChats, RoutingNumberLookupRequestSupportedPaymentTypeCheck, RoutingNumberLookupRequestSupportedPaymentTypeCrossBorder, RoutingNumberLookupRequestSupportedPaymentTypeDkNets, RoutingNumberLookupRequestSupportedPaymentTypeEft, RoutingNumberLookupRequestSupportedPaymentTypeEthereum, RoutingNumberLookupRequestSupportedPaymentTypeGBFps, RoutingNumberLookupRequestSupportedPaymentTypeHuIcs, RoutingNumberLookupRequestSupportedPaymentTypeInterac, RoutingNumberLookupRequestSupportedPaymentTypeMasav, RoutingNumberLookupRequestSupportedPaymentTypeMxCcen, RoutingNumberLookupRequestSupportedPaymentTypeNeft, RoutingNumberLookupRequestSupportedPaymentTypeNics, RoutingNumberLookupRequestSupportedPaymentTypeNzBecs, RoutingNumberLookupRequestSupportedPaymentTypePlElixir, RoutingNumberLookupRequestSupportedPaymentTypePolygon, RoutingNumberLookupRequestSupportedPaymentTypeProvxchange, RoutingNumberLookupRequestSupportedPaymentTypeRoSent, RoutingNumberLookupRequestSupportedPaymentTypeRtp, RoutingNumberLookupRequestSupportedPaymentTypeSeBankgirot, RoutingNumberLookupRequestSupportedPaymentTypeSen, RoutingNumberLookupRequestSupportedPaymentTypeSepa, RoutingNumberLookupRequestSupportedPaymentTypeSgGiro, RoutingNumberLookupRequestSupportedPaymentTypeSic, RoutingNumberLookupRequestSupportedPaymentTypeSignet, RoutingNumberLookupRequestSupportedPaymentTypeSknbi, RoutingNumberLookupRequestSupportedPaymentTypeSolana, RoutingNumberLookupRequestSupportedPaymentTypeWire, RoutingNumberLookupRequestSupportedPaymentTypeZengin:
		return true
	}
	return false
}

type ValidationValidateRoutingNumberParams struct {
	// The routing number that is being validated.
	RoutingNumber param.Field[string] `query:"routing_number" api:"required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details. In sandbox mode we currently only support `aba` and `swift` with
	// routing numbers '123456789' and 'GRINUST0XXX' respectively.
	RoutingNumberType param.Field[ValidationValidateRoutingNumberParamsRoutingNumberType] `query:"routing_number_type" api:"required"`
}

// URLQuery serializes [ValidationValidateRoutingNumberParams]'s query parameters
// as `url.Values`.
func (r ValidationValidateRoutingNumberParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details. In sandbox mode we currently only support `aba` and `swift` with
// routing numbers '123456789' and 'GRINUST0XXX' respectively.
type ValidationValidateRoutingNumberParamsRoutingNumberType string

const (
	ValidationValidateRoutingNumberParamsRoutingNumberTypeAba                     ValidationValidateRoutingNumberParamsRoutingNumberType = "aba"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeAuBsb                   ValidationValidateRoutingNumberParamsRoutingNumberType = "au_bsb"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeBrCodigo                ValidationValidateRoutingNumberParamsRoutingNumberType = "br_codigo"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeCaCpa                   ValidationValidateRoutingNumberParamsRoutingNumberType = "ca_cpa"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeChips                   ValidationValidateRoutingNumberParamsRoutingNumberType = "chips"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeCnaps                   ValidationValidateRoutingNumberParamsRoutingNumberType = "cnaps"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeDkInterbankClearingCode ValidationValidateRoutingNumberParamsRoutingNumberType = "dk_interbank_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeGBSortCode              ValidationValidateRoutingNumberParamsRoutingNumberType = "gb_sort_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeHkInterbankClearingCode ValidationValidateRoutingNumberParamsRoutingNumberType = "hk_interbank_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeHuInterbankClearingCode ValidationValidateRoutingNumberParamsRoutingNumberType = "hu_interbank_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeIDSknbiCode             ValidationValidateRoutingNumberParamsRoutingNumberType = "id_sknbi_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeIlBankCode              ValidationValidateRoutingNumberParamsRoutingNumberType = "il_bank_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeInIfsc                  ValidationValidateRoutingNumberParamsRoutingNumberType = "in_ifsc"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeJpZenginCode            ValidationValidateRoutingNumberParamsRoutingNumberType = "jp_zengin_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeMxBankIdentifier        ValidationValidateRoutingNumberParamsRoutingNumberType = "mx_bank_identifier"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeMyBranchCode            ValidationValidateRoutingNumberParamsRoutingNumberType = "my_branch_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeNzNationalClearingCode  ValidationValidateRoutingNumberParamsRoutingNumberType = "nz_national_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypePlNationalClearingCode  ValidationValidateRoutingNumberParamsRoutingNumberType = "pl_national_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeSeBankgiroClearingCode  ValidationValidateRoutingNumberParamsRoutingNumberType = "se_bankgiro_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeSgInterbankClearingCode ValidationValidateRoutingNumberParamsRoutingNumberType = "sg_interbank_clearing_code"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeSwift                   ValidationValidateRoutingNumberParamsRoutingNumberType = "swift"
	ValidationValidateRoutingNumberParamsRoutingNumberTypeZaNationalClearingCode  ValidationValidateRoutingNumberParamsRoutingNumberType = "za_national_clearing_code"
)

func (r ValidationValidateRoutingNumberParamsRoutingNumberType) IsKnown() bool {
	switch r {
	case ValidationValidateRoutingNumberParamsRoutingNumberTypeAba, ValidationValidateRoutingNumberParamsRoutingNumberTypeAuBsb, ValidationValidateRoutingNumberParamsRoutingNumberTypeBrCodigo, ValidationValidateRoutingNumberParamsRoutingNumberTypeCaCpa, ValidationValidateRoutingNumberParamsRoutingNumberTypeChips, ValidationValidateRoutingNumberParamsRoutingNumberTypeCnaps, ValidationValidateRoutingNumberParamsRoutingNumberTypeDkInterbankClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeGBSortCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeHkInterbankClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeHuInterbankClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeIDSknbiCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeIlBankCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeInIfsc, ValidationValidateRoutingNumberParamsRoutingNumberTypeJpZenginCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeMxBankIdentifier, ValidationValidateRoutingNumberParamsRoutingNumberTypeMyBranchCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeNzNationalClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypePlNationalClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeSeBankgiroClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeSgInterbankClearingCode, ValidationValidateRoutingNumberParamsRoutingNumberTypeSwift, ValidationValidateRoutingNumberParamsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}
