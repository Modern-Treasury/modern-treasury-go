package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

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
	return query.Marshal(r)
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
