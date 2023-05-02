package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type RoutingDetailService struct {
	Options []option.RequestOption
}

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
	JSON        RoutingDetailJSON
}

type RoutingDetailJSON struct {
	ID                apijson.Metadata
	Object            apijson.Metadata
	LiveMode          apijson.Metadata
	CreatedAt         apijson.Metadata
	UpdatedAt         apijson.Metadata
	DiscardedAt       apijson.Metadata
	RoutingNumber     apijson.Metadata
	RoutingNumberType apijson.Metadata
	PaymentType       apijson.Metadata
	BankName          apijson.Metadata
	BankAddress       apijson.Metadata
	raw               string
	Extras            map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingDetail using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RoutingDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutingDetailParam struct {
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
	BankName    field.Field[string]                        `json:"bank_name,required"`
	BankAddress field.Field[RoutingDetailBankAddressParam] `json:"bank_address,required,nullable"`
}

// MarshalJSON serializes RoutingDetailParam into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r RoutingDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ID         apijson.Metadata
	Object     apijson.Metadata
	LiveMode   apijson.Metadata
	CreatedAt  apijson.Metadata
	UpdatedAt  apijson.Metadata
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RoutingDetailBankAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *RoutingDetailBankAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RoutingDetailBankAddressParam struct {
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

// MarshalJSON serializes RoutingDetailBankAddressParam into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r RoutingDetailBankAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes RoutingDetailListParams into a url.Values of the query
// parameters associated with this value
func (r RoutingDetailListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type RoutingDetailDeleteParamsAccountsType string

const (
	RoutingDetailDeleteParamsAccountsTypeExternalAccounts RoutingDetailDeleteParamsAccountsType = "external_accounts"
)
