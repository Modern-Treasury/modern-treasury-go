// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// RoutingDetailService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutingDetailService] method instead.
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
func (r *RoutingDetailService) New(ctx context.Context, accountsType RoutingDetailNewParamsAccountsType, accountID string, body RoutingDetailNewParams, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/%v/%s/routing_details", accountsType, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a single routing detail for a single internal or external account.
func (r *RoutingDetailService) Get(ctx context.Context, accountsType shared.AccountsType, accountID string, id string, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/%v/%s/routing_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) List(ctx context.Context, accountsType shared.AccountsType, accountID string, query RoutingDetailListParams, opts ...option.RequestOption) (res *pagination.Page[RoutingDetail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/%v/%s/routing_details", accountsType, accountID)
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
func (r *RoutingDetailService) ListAutoPaging(ctx context.Context, accountsType shared.AccountsType, accountID string, query RoutingDetailListParams, opts ...option.RequestOption) *pagination.PageAutoPager[RoutingDetail] {
	return pagination.NewPageAutoPager(r.List(ctx, accountsType, accountID, query, opts...))
}

// Delete a routing detail for a single external account.
func (r *RoutingDetailService) Delete(ctx context.Context, accountsType RoutingDetailDeleteParamsAccountsType, accountID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("api/%v/%s/routing_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type RoutingDetail struct {
	ID          string         `json:"id" api:"required" format:"uuid"`
	BankAddress shared.Address `json:"bank_address" api:"required,nullable"`
	// The name of the bank.
	BankName    string    `json:"bank_name" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at" api:"required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType RoutingDetailPaymentType `json:"payment_type" api:"required,nullable"`
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number" api:"required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details.
	RoutingNumberType RoutingDetailRoutingNumberType `json:"routing_number_type" api:"required"`
	UpdatedAt         time.Time                      `json:"updated_at" api:"required" format:"date-time"`
	JSON              routingDetailJSON              `json:"-"`
}

// routingDetailJSON contains the JSON metadata for the struct [RoutingDetail]
type routingDetailJSON struct {
	ID                apijson.Field
	BankAddress       apijson.Field
	BankName          apijson.Field
	CreatedAt         apijson.Field
	DiscardedAt       apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	PaymentType       apijson.Field
	RoutingNumber     apijson.Field
	RoutingNumberType apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RoutingDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDetailJSON) RawJSON() string {
	return r.raw
}

// If the routing detail is to be used for a specific payment type this field will
// be populated, otherwise null.
type RoutingDetailPaymentType string

const (
	RoutingDetailPaymentTypeACH         RoutingDetailPaymentType = "ach"
	RoutingDetailPaymentTypeAuBecs      RoutingDetailPaymentType = "au_becs"
	RoutingDetailPaymentTypeBacs        RoutingDetailPaymentType = "bacs"
	RoutingDetailPaymentTypeBook        RoutingDetailPaymentType = "book"
	RoutingDetailPaymentTypeCard        RoutingDetailPaymentType = "card"
	RoutingDetailPaymentTypeChats       RoutingDetailPaymentType = "chats"
	RoutingDetailPaymentTypeCheck       RoutingDetailPaymentType = "check"
	RoutingDetailPaymentTypeCrossBorder RoutingDetailPaymentType = "cross_border"
	RoutingDetailPaymentTypeDkNets      RoutingDetailPaymentType = "dk_nets"
	RoutingDetailPaymentTypeEft         RoutingDetailPaymentType = "eft"
	RoutingDetailPaymentTypeGBFps       RoutingDetailPaymentType = "gb_fps"
	RoutingDetailPaymentTypeMasav       RoutingDetailPaymentType = "masav"
	RoutingDetailPaymentTypeMxCcen      RoutingDetailPaymentType = "mx_ccen"
	RoutingDetailPaymentTypeNeft        RoutingDetailPaymentType = "neft"
	RoutingDetailPaymentTypeNics        RoutingDetailPaymentType = "nics"
	RoutingDetailPaymentTypeNzBecs      RoutingDetailPaymentType = "nz_becs"
	RoutingDetailPaymentTypePlElixir    RoutingDetailPaymentType = "pl_elixir"
	RoutingDetailPaymentTypeRtp         RoutingDetailPaymentType = "rtp"
	RoutingDetailPaymentTypeSeBankgirot RoutingDetailPaymentType = "se_bankgirot"
	RoutingDetailPaymentTypeSepa        RoutingDetailPaymentType = "sepa"
	RoutingDetailPaymentTypeSgGiro      RoutingDetailPaymentType = "sg_giro"
	RoutingDetailPaymentTypeSic         RoutingDetailPaymentType = "sic"
	RoutingDetailPaymentTypeStablecoin  RoutingDetailPaymentType = "stablecoin"
	RoutingDetailPaymentTypeWire        RoutingDetailPaymentType = "wire"
	RoutingDetailPaymentTypeZengin      RoutingDetailPaymentType = "zengin"
)

func (r RoutingDetailPaymentType) IsKnown() bool {
	switch r {
	case RoutingDetailPaymentTypeACH, RoutingDetailPaymentTypeAuBecs, RoutingDetailPaymentTypeBacs, RoutingDetailPaymentTypeBook, RoutingDetailPaymentTypeCard, RoutingDetailPaymentTypeChats, RoutingDetailPaymentTypeCheck, RoutingDetailPaymentTypeCrossBorder, RoutingDetailPaymentTypeDkNets, RoutingDetailPaymentTypeEft, RoutingDetailPaymentTypeGBFps, RoutingDetailPaymentTypeMasav, RoutingDetailPaymentTypeMxCcen, RoutingDetailPaymentTypeNeft, RoutingDetailPaymentTypeNics, RoutingDetailPaymentTypeNzBecs, RoutingDetailPaymentTypePlElixir, RoutingDetailPaymentTypeRtp, RoutingDetailPaymentTypeSeBankgirot, RoutingDetailPaymentTypeSepa, RoutingDetailPaymentTypeSgGiro, RoutingDetailPaymentTypeSic, RoutingDetailPaymentTypeStablecoin, RoutingDetailPaymentTypeWire, RoutingDetailPaymentTypeZengin:
		return true
	}
	return false
}

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details.
type RoutingDetailRoutingNumberType string

const (
	RoutingDetailRoutingNumberTypeAba                     RoutingDetailRoutingNumberType = "aba"
	RoutingDetailRoutingNumberTypeAuBsb                   RoutingDetailRoutingNumberType = "au_bsb"
	RoutingDetailRoutingNumberTypeBrCodigo                RoutingDetailRoutingNumberType = "br_codigo"
	RoutingDetailRoutingNumberTypeCaCpa                   RoutingDetailRoutingNumberType = "ca_cpa"
	RoutingDetailRoutingNumberTypeChips                   RoutingDetailRoutingNumberType = "chips"
	RoutingDetailRoutingNumberTypeCnaps                   RoutingDetailRoutingNumberType = "cnaps"
	RoutingDetailRoutingNumberTypeDkInterbankClearingCode RoutingDetailRoutingNumberType = "dk_interbank_clearing_code"
	RoutingDetailRoutingNumberTypeGBSortCode              RoutingDetailRoutingNumberType = "gb_sort_code"
	RoutingDetailRoutingNumberTypeHkInterbankClearingCode RoutingDetailRoutingNumberType = "hk_interbank_clearing_code"
	RoutingDetailRoutingNumberTypeIlBankCode              RoutingDetailRoutingNumberType = "il_bank_code"
	RoutingDetailRoutingNumberTypeInIfsc                  RoutingDetailRoutingNumberType = "in_ifsc"
	RoutingDetailRoutingNumberTypeJpZenginCode            RoutingDetailRoutingNumberType = "jp_zengin_code"
	RoutingDetailRoutingNumberTypeMxBankIdentifier        RoutingDetailRoutingNumberType = "mx_bank_identifier"
	RoutingDetailRoutingNumberTypeMyBranchCode            RoutingDetailRoutingNumberType = "my_branch_code"
	RoutingDetailRoutingNumberTypeNzNationalClearingCode  RoutingDetailRoutingNumberType = "nz_national_clearing_code"
	RoutingDetailRoutingNumberTypePlNationalClearingCode  RoutingDetailRoutingNumberType = "pl_national_clearing_code"
	RoutingDetailRoutingNumberTypeSeBankgiroClearingCode  RoutingDetailRoutingNumberType = "se_bankgiro_clearing_code"
	RoutingDetailRoutingNumberTypeSgInterbankClearingCode RoutingDetailRoutingNumberType = "sg_interbank_clearing_code"
	RoutingDetailRoutingNumberTypeSwift                   RoutingDetailRoutingNumberType = "swift"
	RoutingDetailRoutingNumberTypeVasp                    RoutingDetailRoutingNumberType = "vasp"
	RoutingDetailRoutingNumberTypeZaNationalClearingCode  RoutingDetailRoutingNumberType = "za_national_clearing_code"
)

func (r RoutingDetailRoutingNumberType) IsKnown() bool {
	switch r {
	case RoutingDetailRoutingNumberTypeAba, RoutingDetailRoutingNumberTypeAuBsb, RoutingDetailRoutingNumberTypeBrCodigo, RoutingDetailRoutingNumberTypeCaCpa, RoutingDetailRoutingNumberTypeChips, RoutingDetailRoutingNumberTypeCnaps, RoutingDetailRoutingNumberTypeDkInterbankClearingCode, RoutingDetailRoutingNumberTypeGBSortCode, RoutingDetailRoutingNumberTypeHkInterbankClearingCode, RoutingDetailRoutingNumberTypeIlBankCode, RoutingDetailRoutingNumberTypeInIfsc, RoutingDetailRoutingNumberTypeJpZenginCode, RoutingDetailRoutingNumberTypeMxBankIdentifier, RoutingDetailRoutingNumberTypeMyBranchCode, RoutingDetailRoutingNumberTypeNzNationalClearingCode, RoutingDetailRoutingNumberTypePlNationalClearingCode, RoutingDetailRoutingNumberTypeSeBankgiroClearingCode, RoutingDetailRoutingNumberTypeSgInterbankClearingCode, RoutingDetailRoutingNumberTypeSwift, RoutingDetailRoutingNumberTypeVasp, RoutingDetailRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type RoutingDetailCreateRequestParam struct {
	// The routing number of the bank.
	RoutingNumber param.Field[string] `json:"routing_number" api:"required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details.
	RoutingNumberType param.Field[RoutingDetailCreateRequestRoutingNumberType] `json:"routing_number_type" api:"required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType param.Field[RoutingDetailCreateRequestPaymentType] `json:"payment_type"`
}

func (r RoutingDetailCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details.
type RoutingDetailCreateRequestRoutingNumberType string

const (
	RoutingDetailCreateRequestRoutingNumberTypeAba                     RoutingDetailCreateRequestRoutingNumberType = "aba"
	RoutingDetailCreateRequestRoutingNumberTypeAuBsb                   RoutingDetailCreateRequestRoutingNumberType = "au_bsb"
	RoutingDetailCreateRequestRoutingNumberTypeBrCodigo                RoutingDetailCreateRequestRoutingNumberType = "br_codigo"
	RoutingDetailCreateRequestRoutingNumberTypeCaCpa                   RoutingDetailCreateRequestRoutingNumberType = "ca_cpa"
	RoutingDetailCreateRequestRoutingNumberTypeChips                   RoutingDetailCreateRequestRoutingNumberType = "chips"
	RoutingDetailCreateRequestRoutingNumberTypeCnaps                   RoutingDetailCreateRequestRoutingNumberType = "cnaps"
	RoutingDetailCreateRequestRoutingNumberTypeDkInterbankClearingCode RoutingDetailCreateRequestRoutingNumberType = "dk_interbank_clearing_code"
	RoutingDetailCreateRequestRoutingNumberTypeGBSortCode              RoutingDetailCreateRequestRoutingNumberType = "gb_sort_code"
	RoutingDetailCreateRequestRoutingNumberTypeHkInterbankClearingCode RoutingDetailCreateRequestRoutingNumberType = "hk_interbank_clearing_code"
	RoutingDetailCreateRequestRoutingNumberTypeIlBankCode              RoutingDetailCreateRequestRoutingNumberType = "il_bank_code"
	RoutingDetailCreateRequestRoutingNumberTypeInIfsc                  RoutingDetailCreateRequestRoutingNumberType = "in_ifsc"
	RoutingDetailCreateRequestRoutingNumberTypeJpZenginCode            RoutingDetailCreateRequestRoutingNumberType = "jp_zengin_code"
	RoutingDetailCreateRequestRoutingNumberTypeMxBankIdentifier        RoutingDetailCreateRequestRoutingNumberType = "mx_bank_identifier"
	RoutingDetailCreateRequestRoutingNumberTypeMyBranchCode            RoutingDetailCreateRequestRoutingNumberType = "my_branch_code"
	RoutingDetailCreateRequestRoutingNumberTypeNzNationalClearingCode  RoutingDetailCreateRequestRoutingNumberType = "nz_national_clearing_code"
	RoutingDetailCreateRequestRoutingNumberTypePlNationalClearingCode  RoutingDetailCreateRequestRoutingNumberType = "pl_national_clearing_code"
	RoutingDetailCreateRequestRoutingNumberTypeSeBankgiroClearingCode  RoutingDetailCreateRequestRoutingNumberType = "se_bankgiro_clearing_code"
	RoutingDetailCreateRequestRoutingNumberTypeSgInterbankClearingCode RoutingDetailCreateRequestRoutingNumberType = "sg_interbank_clearing_code"
	RoutingDetailCreateRequestRoutingNumberTypeSwift                   RoutingDetailCreateRequestRoutingNumberType = "swift"
	RoutingDetailCreateRequestRoutingNumberTypeVasp                    RoutingDetailCreateRequestRoutingNumberType = "vasp"
	RoutingDetailCreateRequestRoutingNumberTypeZaNationalClearingCode  RoutingDetailCreateRequestRoutingNumberType = "za_national_clearing_code"
)

func (r RoutingDetailCreateRequestRoutingNumberType) IsKnown() bool {
	switch r {
	case RoutingDetailCreateRequestRoutingNumberTypeAba, RoutingDetailCreateRequestRoutingNumberTypeAuBsb, RoutingDetailCreateRequestRoutingNumberTypeBrCodigo, RoutingDetailCreateRequestRoutingNumberTypeCaCpa, RoutingDetailCreateRequestRoutingNumberTypeChips, RoutingDetailCreateRequestRoutingNumberTypeCnaps, RoutingDetailCreateRequestRoutingNumberTypeDkInterbankClearingCode, RoutingDetailCreateRequestRoutingNumberTypeGBSortCode, RoutingDetailCreateRequestRoutingNumberTypeHkInterbankClearingCode, RoutingDetailCreateRequestRoutingNumberTypeIlBankCode, RoutingDetailCreateRequestRoutingNumberTypeInIfsc, RoutingDetailCreateRequestRoutingNumberTypeJpZenginCode, RoutingDetailCreateRequestRoutingNumberTypeMxBankIdentifier, RoutingDetailCreateRequestRoutingNumberTypeMyBranchCode, RoutingDetailCreateRequestRoutingNumberTypeNzNationalClearingCode, RoutingDetailCreateRequestRoutingNumberTypePlNationalClearingCode, RoutingDetailCreateRequestRoutingNumberTypeSeBankgiroClearingCode, RoutingDetailCreateRequestRoutingNumberTypeSgInterbankClearingCode, RoutingDetailCreateRequestRoutingNumberTypeSwift, RoutingDetailCreateRequestRoutingNumberTypeVasp, RoutingDetailCreateRequestRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

// If the routing detail is to be used for a specific payment type this field will
// be populated, otherwise null.
type RoutingDetailCreateRequestPaymentType string

const (
	RoutingDetailCreateRequestPaymentTypeACH         RoutingDetailCreateRequestPaymentType = "ach"
	RoutingDetailCreateRequestPaymentTypeAuBecs      RoutingDetailCreateRequestPaymentType = "au_becs"
	RoutingDetailCreateRequestPaymentTypeBacs        RoutingDetailCreateRequestPaymentType = "bacs"
	RoutingDetailCreateRequestPaymentTypeBook        RoutingDetailCreateRequestPaymentType = "book"
	RoutingDetailCreateRequestPaymentTypeCard        RoutingDetailCreateRequestPaymentType = "card"
	RoutingDetailCreateRequestPaymentTypeChats       RoutingDetailCreateRequestPaymentType = "chats"
	RoutingDetailCreateRequestPaymentTypeCheck       RoutingDetailCreateRequestPaymentType = "check"
	RoutingDetailCreateRequestPaymentTypeCrossBorder RoutingDetailCreateRequestPaymentType = "cross_border"
	RoutingDetailCreateRequestPaymentTypeDkNets      RoutingDetailCreateRequestPaymentType = "dk_nets"
	RoutingDetailCreateRequestPaymentTypeEft         RoutingDetailCreateRequestPaymentType = "eft"
	RoutingDetailCreateRequestPaymentTypeGBFps       RoutingDetailCreateRequestPaymentType = "gb_fps"
	RoutingDetailCreateRequestPaymentTypeMasav       RoutingDetailCreateRequestPaymentType = "masav"
	RoutingDetailCreateRequestPaymentTypeMxCcen      RoutingDetailCreateRequestPaymentType = "mx_ccen"
	RoutingDetailCreateRequestPaymentTypeNeft        RoutingDetailCreateRequestPaymentType = "neft"
	RoutingDetailCreateRequestPaymentTypeNics        RoutingDetailCreateRequestPaymentType = "nics"
	RoutingDetailCreateRequestPaymentTypeNzBecs      RoutingDetailCreateRequestPaymentType = "nz_becs"
	RoutingDetailCreateRequestPaymentTypePlElixir    RoutingDetailCreateRequestPaymentType = "pl_elixir"
	RoutingDetailCreateRequestPaymentTypeRtp         RoutingDetailCreateRequestPaymentType = "rtp"
	RoutingDetailCreateRequestPaymentTypeSeBankgirot RoutingDetailCreateRequestPaymentType = "se_bankgirot"
	RoutingDetailCreateRequestPaymentTypeSepa        RoutingDetailCreateRequestPaymentType = "sepa"
	RoutingDetailCreateRequestPaymentTypeSgGiro      RoutingDetailCreateRequestPaymentType = "sg_giro"
	RoutingDetailCreateRequestPaymentTypeSic         RoutingDetailCreateRequestPaymentType = "sic"
	RoutingDetailCreateRequestPaymentTypeStablecoin  RoutingDetailCreateRequestPaymentType = "stablecoin"
	RoutingDetailCreateRequestPaymentTypeWire        RoutingDetailCreateRequestPaymentType = "wire"
	RoutingDetailCreateRequestPaymentTypeZengin      RoutingDetailCreateRequestPaymentType = "zengin"
)

func (r RoutingDetailCreateRequestPaymentType) IsKnown() bool {
	switch r {
	case RoutingDetailCreateRequestPaymentTypeACH, RoutingDetailCreateRequestPaymentTypeAuBecs, RoutingDetailCreateRequestPaymentTypeBacs, RoutingDetailCreateRequestPaymentTypeBook, RoutingDetailCreateRequestPaymentTypeCard, RoutingDetailCreateRequestPaymentTypeChats, RoutingDetailCreateRequestPaymentTypeCheck, RoutingDetailCreateRequestPaymentTypeCrossBorder, RoutingDetailCreateRequestPaymentTypeDkNets, RoutingDetailCreateRequestPaymentTypeEft, RoutingDetailCreateRequestPaymentTypeGBFps, RoutingDetailCreateRequestPaymentTypeMasav, RoutingDetailCreateRequestPaymentTypeMxCcen, RoutingDetailCreateRequestPaymentTypeNeft, RoutingDetailCreateRequestPaymentTypeNics, RoutingDetailCreateRequestPaymentTypeNzBecs, RoutingDetailCreateRequestPaymentTypePlElixir, RoutingDetailCreateRequestPaymentTypeRtp, RoutingDetailCreateRequestPaymentTypeSeBankgirot, RoutingDetailCreateRequestPaymentTypeSepa, RoutingDetailCreateRequestPaymentTypeSgGiro, RoutingDetailCreateRequestPaymentTypeSic, RoutingDetailCreateRequestPaymentTypeStablecoin, RoutingDetailCreateRequestPaymentTypeWire, RoutingDetailCreateRequestPaymentTypeZengin:
		return true
	}
	return false
}

type RoutingDetailNewParams struct {
	RoutingDetailCreateRequest RoutingDetailCreateRequestParam `json:"routing_detail_create_request" api:"required"`
}

func (r RoutingDetailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RoutingDetailCreateRequest)
}

type RoutingDetailNewParamsAccountsType string

const (
	RoutingDetailNewParamsAccountsTypeExternalAccounts RoutingDetailNewParamsAccountsType = "external_accounts"
)

func (r RoutingDetailNewParamsAccountsType) IsKnown() bool {
	switch r {
	case RoutingDetailNewParamsAccountsTypeExternalAccounts:
		return true
	}
	return false
}

type RoutingDetailListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [RoutingDetailListParams]'s query parameters as
// `url.Values`.
func (r RoutingDetailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingDetailDeleteParamsAccountsType string

const (
	RoutingDetailDeleteParamsAccountsTypeExternalAccounts RoutingDetailDeleteParamsAccountsType = "external_accounts"
)

func (r RoutingDetailDeleteParamsAccountsType) IsKnown() bool {
	switch r {
	case RoutingDetailDeleteParamsAccountsTypeExternalAccounts:
		return true
	}
	return false
}
