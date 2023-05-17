package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// RoutingDetailService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewRoutingDetailService]
// method instead.
type RoutingDetailService struct {
	Options []option.RequestOption
}

// NewRoutingDetailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingDetailService(opts ...option.RequestOption) (r *RoutingDetailService) {
	r = &RoutingDetailService{}
	r.Options = opts
	return
}

// Create a routing detail for a single external account.
func (r *RoutingDetailService) New(ctx context.Context, accounts_type RoutingDetailNewParamsAccountsType, account_id string, body RoutingDetailNewParams, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details", accounts_type, account_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single routing detail for a single internal or external account.
func (r *RoutingDetailService) Get(ctx context.Context, accounts_type shared.AccountsType, account_id string, id string, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details/%s", accounts_type, account_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) List(ctx context.Context, accounts_type shared.AccountsType, account_id string, query RoutingDetailListParams, opts ...option.RequestOption) (res *shared.Page[RoutingDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details", accounts_type, account_id)
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

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) ListAutoPaging(ctx context.Context, accounts_type shared.AccountsType, account_id string, query RoutingDetailListParams, opts ...option.RequestOption) *shared.PageAutoPager[RoutingDetail] {
	return shared.NewPageAutoPager(r.List(ctx, accounts_type, account_id, query, opts...))
}

// Delete a routing detail for a single external account.
func (r *RoutingDetailService) Delete(ctx context.Context, accounts_type RoutingDetailDeleteParamsAccountsType, account_id string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/%s/%s/routing_details/%s", accounts_type, account_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

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
	JSON        routingDetailJSON
}

// routingDetailJSON contains the JSON metadata for the struct [RoutingDetail]
type routingDetailJSON struct {
	ID                apijson.Field
	Object            apijson.Field
	LiveMode          apijson.Field
	CreatedAt         apijson.Field
	UpdatedAt         apijson.Field
	DiscardedAt       apijson.Field
	RoutingNumber     apijson.Field
	RoutingNumberType apijson.Field
	PaymentType       apijson.Field
	BankName          apijson.Field
	BankAddress       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RoutingDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	JSON    routingDetailBankAddressJSON
}

// routingDetailBankAddressJSON contains the JSON metadata for the struct
// [RoutingDetailBankAddress]
type routingDetailBankAddressJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	LiveMode    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	Region      apijson.Field
	PostalCode  apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDetailBankAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutingDetailParam struct {
	ID     param.Field[string] `json:"id,required" format:"uuid"`
	Object param.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    param.Field[bool]      `json:"live_mode,required"`
	CreatedAt   param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	DiscardedAt param.Field[time.Time] `json:"discarded_at,required,nullable" format:"date-time"`
	// The routing number of the bank.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// One of `aba`, `swift`, `ca_cpa`, `au_bsb`, `gb_sort_code`, `in_ifsc`, `cnaps`.
	RoutingNumberType param.Field[RoutingDetailRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType param.Field[RoutingDetailPaymentType] `json:"payment_type,required,nullable"`
	// The name of the bank.
	BankName    param.Field[string]                        `json:"bank_name,required"`
	BankAddress param.Field[RoutingDetailBankAddressParam] `json:"bank_address,required,nullable"`
}

func (r RoutingDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoutingDetailBankAddressParam struct {
	ID     param.Field[string] `json:"id,required" format:"uuid"`
	Object param.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  param.Field[bool]      `json:"live_mode,required"`
	CreatedAt param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	UpdatedAt param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	Line1     param.Field[string]    `json:"line1,required,nullable"`
	Line2     param.Field[string]    `json:"line2,required,nullable"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required,nullable"`
	// Region or State.
	Region param.Field[string] `json:"region,required,nullable"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required,nullable"`
}

func (r RoutingDetailBankAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoutingDetailNewParams struct {
	// The routing number of the bank.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// One of `aba`, `swift`, `ca_cpa`, `au_bsb`, `gb_sort_code`, `in_ifsc`, `cnaps`.
	RoutingNumberType param.Field[RoutingDetailNewParamsRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType param.Field[RoutingDetailNewParamsPaymentType] `json:"payment_type,nullable"`
}

func (r RoutingDetailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [RoutingDetailListParams]'s query parameters as
// `url.Values`.
func (r RoutingDetailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingDetailDeleteParamsAccountsType string

const (
	RoutingDetailDeleteParamsAccountsTypeExternalAccounts RoutingDetailDeleteParamsAccountsType = "external_accounts"
)
