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
		return
	}
	path := fmt.Sprintf("api/%v/%s/routing_details", accountsType, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single routing detail for a single internal or external account.
func (r *RoutingDetailService) Get(ctx context.Context, accountsType shared.AccountsType, accountID string, id string, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/routing_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) List(ctx context.Context, accountsType shared.AccountsType, accountID string, query RoutingDetailListParams, opts ...option.RequestOption) (res *pagination.Page[RoutingDetail], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
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
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/routing_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type RoutingDetail struct {
	ID          string         `json:"id,required" format:"uuid"`
	BankAddress shared.Address `json:"bank_address,required,nullable"`
	// The name of the bank.
	BankName    string    `json:"bank_name,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType RoutingDetailPaymentType `json:"payment_type,required,nullable"`
	// The routing number of the bank.
	RoutingNumber string `json:"routing_number,required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details.
	RoutingNumberType RoutingDetailRoutingNumberType `json:"routing_number_type,required"`
	UpdatedAt         time.Time                      `json:"updated_at,required" format:"date-time"`
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
	RoutingDetailPaymentTypeBase        RoutingDetailPaymentType = "base"
	RoutingDetailPaymentTypeBook        RoutingDetailPaymentType = "book"
	RoutingDetailPaymentTypeCard        RoutingDetailPaymentType = "card"
	RoutingDetailPaymentTypeChats       RoutingDetailPaymentType = "chats"
	RoutingDetailPaymentTypeCheck       RoutingDetailPaymentType = "check"
	RoutingDetailPaymentTypeCrossBorder RoutingDetailPaymentType = "cross_border"
	RoutingDetailPaymentTypeDkNets      RoutingDetailPaymentType = "dk_nets"
	RoutingDetailPaymentTypeEft         RoutingDetailPaymentType = "eft"
	RoutingDetailPaymentTypeEthereum    RoutingDetailPaymentType = "ethereum"
	RoutingDetailPaymentTypeGBFps       RoutingDetailPaymentType = "gb_fps"
	RoutingDetailPaymentTypeHuIcs       RoutingDetailPaymentType = "hu_ics"
	RoutingDetailPaymentTypeInterac     RoutingDetailPaymentType = "interac"
	RoutingDetailPaymentTypeMasav       RoutingDetailPaymentType = "masav"
	RoutingDetailPaymentTypeMxCcen      RoutingDetailPaymentType = "mx_ccen"
	RoutingDetailPaymentTypeNeft        RoutingDetailPaymentType = "neft"
	RoutingDetailPaymentTypeNics        RoutingDetailPaymentType = "nics"
	RoutingDetailPaymentTypeNzBecs      RoutingDetailPaymentType = "nz_becs"
	RoutingDetailPaymentTypePlElixir    RoutingDetailPaymentType = "pl_elixir"
	RoutingDetailPaymentTypePolygon     RoutingDetailPaymentType = "polygon"
	RoutingDetailPaymentTypeProvxchange RoutingDetailPaymentType = "provxchange"
	RoutingDetailPaymentTypeRoSent      RoutingDetailPaymentType = "ro_sent"
	RoutingDetailPaymentTypeRtp         RoutingDetailPaymentType = "rtp"
	RoutingDetailPaymentTypeSeBankgirot RoutingDetailPaymentType = "se_bankgirot"
	RoutingDetailPaymentTypeSen         RoutingDetailPaymentType = "sen"
	RoutingDetailPaymentTypeSepa        RoutingDetailPaymentType = "sepa"
	RoutingDetailPaymentTypeSgGiro      RoutingDetailPaymentType = "sg_giro"
	RoutingDetailPaymentTypeSic         RoutingDetailPaymentType = "sic"
	RoutingDetailPaymentTypeSignet      RoutingDetailPaymentType = "signet"
	RoutingDetailPaymentTypeSknbi       RoutingDetailPaymentType = "sknbi"
	RoutingDetailPaymentTypeSolana      RoutingDetailPaymentType = "solana"
	RoutingDetailPaymentTypeWire        RoutingDetailPaymentType = "wire"
	RoutingDetailPaymentTypeZengin      RoutingDetailPaymentType = "zengin"
)

func (r RoutingDetailPaymentType) IsKnown() bool {
	switch r {
	case RoutingDetailPaymentTypeACH, RoutingDetailPaymentTypeAuBecs, RoutingDetailPaymentTypeBacs, RoutingDetailPaymentTypeBase, RoutingDetailPaymentTypeBook, RoutingDetailPaymentTypeCard, RoutingDetailPaymentTypeChats, RoutingDetailPaymentTypeCheck, RoutingDetailPaymentTypeCrossBorder, RoutingDetailPaymentTypeDkNets, RoutingDetailPaymentTypeEft, RoutingDetailPaymentTypeEthereum, RoutingDetailPaymentTypeGBFps, RoutingDetailPaymentTypeHuIcs, RoutingDetailPaymentTypeInterac, RoutingDetailPaymentTypeMasav, RoutingDetailPaymentTypeMxCcen, RoutingDetailPaymentTypeNeft, RoutingDetailPaymentTypeNics, RoutingDetailPaymentTypeNzBecs, RoutingDetailPaymentTypePlElixir, RoutingDetailPaymentTypePolygon, RoutingDetailPaymentTypeProvxchange, RoutingDetailPaymentTypeRoSent, RoutingDetailPaymentTypeRtp, RoutingDetailPaymentTypeSeBankgirot, RoutingDetailPaymentTypeSen, RoutingDetailPaymentTypeSepa, RoutingDetailPaymentTypeSgGiro, RoutingDetailPaymentTypeSic, RoutingDetailPaymentTypeSignet, RoutingDetailPaymentTypeSknbi, RoutingDetailPaymentTypeSolana, RoutingDetailPaymentTypeWire, RoutingDetailPaymentTypeZengin:
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
	RoutingDetailRoutingNumberTypeHuInterbankClearingCode RoutingDetailRoutingNumberType = "hu_interbank_clearing_code"
	RoutingDetailRoutingNumberTypeIDSknbiCode             RoutingDetailRoutingNumberType = "id_sknbi_code"
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
	RoutingDetailRoutingNumberTypeZaNationalClearingCode  RoutingDetailRoutingNumberType = "za_national_clearing_code"
)

func (r RoutingDetailRoutingNumberType) IsKnown() bool {
	switch r {
	case RoutingDetailRoutingNumberTypeAba, RoutingDetailRoutingNumberTypeAuBsb, RoutingDetailRoutingNumberTypeBrCodigo, RoutingDetailRoutingNumberTypeCaCpa, RoutingDetailRoutingNumberTypeChips, RoutingDetailRoutingNumberTypeCnaps, RoutingDetailRoutingNumberTypeDkInterbankClearingCode, RoutingDetailRoutingNumberTypeGBSortCode, RoutingDetailRoutingNumberTypeHkInterbankClearingCode, RoutingDetailRoutingNumberTypeHuInterbankClearingCode, RoutingDetailRoutingNumberTypeIDSknbiCode, RoutingDetailRoutingNumberTypeIlBankCode, RoutingDetailRoutingNumberTypeInIfsc, RoutingDetailRoutingNumberTypeJpZenginCode, RoutingDetailRoutingNumberTypeMxBankIdentifier, RoutingDetailRoutingNumberTypeMyBranchCode, RoutingDetailRoutingNumberTypeNzNationalClearingCode, RoutingDetailRoutingNumberTypePlNationalClearingCode, RoutingDetailRoutingNumberTypeSeBankgiroClearingCode, RoutingDetailRoutingNumberTypeSgInterbankClearingCode, RoutingDetailRoutingNumberTypeSwift, RoutingDetailRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type RoutingDetailCreateParam struct {
	// The routing number of the bank.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details.
	RoutingNumberType param.Field[RoutingDetailCreateRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType param.Field[RoutingDetailCreatePaymentType] `json:"payment_type"`
}

func (r RoutingDetailCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details.
type RoutingDetailCreateRoutingNumberType string

const (
	RoutingDetailCreateRoutingNumberTypeAba                     RoutingDetailCreateRoutingNumberType = "aba"
	RoutingDetailCreateRoutingNumberTypeAuBsb                   RoutingDetailCreateRoutingNumberType = "au_bsb"
	RoutingDetailCreateRoutingNumberTypeBrCodigo                RoutingDetailCreateRoutingNumberType = "br_codigo"
	RoutingDetailCreateRoutingNumberTypeCaCpa                   RoutingDetailCreateRoutingNumberType = "ca_cpa"
	RoutingDetailCreateRoutingNumberTypeChips                   RoutingDetailCreateRoutingNumberType = "chips"
	RoutingDetailCreateRoutingNumberTypeCnaps                   RoutingDetailCreateRoutingNumberType = "cnaps"
	RoutingDetailCreateRoutingNumberTypeDkInterbankClearingCode RoutingDetailCreateRoutingNumberType = "dk_interbank_clearing_code"
	RoutingDetailCreateRoutingNumberTypeGBSortCode              RoutingDetailCreateRoutingNumberType = "gb_sort_code"
	RoutingDetailCreateRoutingNumberTypeHkInterbankClearingCode RoutingDetailCreateRoutingNumberType = "hk_interbank_clearing_code"
	RoutingDetailCreateRoutingNumberTypeHuInterbankClearingCode RoutingDetailCreateRoutingNumberType = "hu_interbank_clearing_code"
	RoutingDetailCreateRoutingNumberTypeIDSknbiCode             RoutingDetailCreateRoutingNumberType = "id_sknbi_code"
	RoutingDetailCreateRoutingNumberTypeIlBankCode              RoutingDetailCreateRoutingNumberType = "il_bank_code"
	RoutingDetailCreateRoutingNumberTypeInIfsc                  RoutingDetailCreateRoutingNumberType = "in_ifsc"
	RoutingDetailCreateRoutingNumberTypeJpZenginCode            RoutingDetailCreateRoutingNumberType = "jp_zengin_code"
	RoutingDetailCreateRoutingNumberTypeMxBankIdentifier        RoutingDetailCreateRoutingNumberType = "mx_bank_identifier"
	RoutingDetailCreateRoutingNumberTypeMyBranchCode            RoutingDetailCreateRoutingNumberType = "my_branch_code"
	RoutingDetailCreateRoutingNumberTypeNzNationalClearingCode  RoutingDetailCreateRoutingNumberType = "nz_national_clearing_code"
	RoutingDetailCreateRoutingNumberTypePlNationalClearingCode  RoutingDetailCreateRoutingNumberType = "pl_national_clearing_code"
	RoutingDetailCreateRoutingNumberTypeSeBankgiroClearingCode  RoutingDetailCreateRoutingNumberType = "se_bankgiro_clearing_code"
	RoutingDetailCreateRoutingNumberTypeSgInterbankClearingCode RoutingDetailCreateRoutingNumberType = "sg_interbank_clearing_code"
	RoutingDetailCreateRoutingNumberTypeSwift                   RoutingDetailCreateRoutingNumberType = "swift"
	RoutingDetailCreateRoutingNumberTypeZaNationalClearingCode  RoutingDetailCreateRoutingNumberType = "za_national_clearing_code"
)

func (r RoutingDetailCreateRoutingNumberType) IsKnown() bool {
	switch r {
	case RoutingDetailCreateRoutingNumberTypeAba, RoutingDetailCreateRoutingNumberTypeAuBsb, RoutingDetailCreateRoutingNumberTypeBrCodigo, RoutingDetailCreateRoutingNumberTypeCaCpa, RoutingDetailCreateRoutingNumberTypeChips, RoutingDetailCreateRoutingNumberTypeCnaps, RoutingDetailCreateRoutingNumberTypeDkInterbankClearingCode, RoutingDetailCreateRoutingNumberTypeGBSortCode, RoutingDetailCreateRoutingNumberTypeHkInterbankClearingCode, RoutingDetailCreateRoutingNumberTypeHuInterbankClearingCode, RoutingDetailCreateRoutingNumberTypeIDSknbiCode, RoutingDetailCreateRoutingNumberTypeIlBankCode, RoutingDetailCreateRoutingNumberTypeInIfsc, RoutingDetailCreateRoutingNumberTypeJpZenginCode, RoutingDetailCreateRoutingNumberTypeMxBankIdentifier, RoutingDetailCreateRoutingNumberTypeMyBranchCode, RoutingDetailCreateRoutingNumberTypeNzNationalClearingCode, RoutingDetailCreateRoutingNumberTypePlNationalClearingCode, RoutingDetailCreateRoutingNumberTypeSeBankgiroClearingCode, RoutingDetailCreateRoutingNumberTypeSgInterbankClearingCode, RoutingDetailCreateRoutingNumberTypeSwift, RoutingDetailCreateRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

// If the routing detail is to be used for a specific payment type this field will
// be populated, otherwise null.
type RoutingDetailCreatePaymentType string

const (
	RoutingDetailCreatePaymentTypeACH         RoutingDetailCreatePaymentType = "ach"
	RoutingDetailCreatePaymentTypeAuBecs      RoutingDetailCreatePaymentType = "au_becs"
	RoutingDetailCreatePaymentTypeBacs        RoutingDetailCreatePaymentType = "bacs"
	RoutingDetailCreatePaymentTypeBase        RoutingDetailCreatePaymentType = "base"
	RoutingDetailCreatePaymentTypeBook        RoutingDetailCreatePaymentType = "book"
	RoutingDetailCreatePaymentTypeCard        RoutingDetailCreatePaymentType = "card"
	RoutingDetailCreatePaymentTypeChats       RoutingDetailCreatePaymentType = "chats"
	RoutingDetailCreatePaymentTypeCheck       RoutingDetailCreatePaymentType = "check"
	RoutingDetailCreatePaymentTypeCrossBorder RoutingDetailCreatePaymentType = "cross_border"
	RoutingDetailCreatePaymentTypeDkNets      RoutingDetailCreatePaymentType = "dk_nets"
	RoutingDetailCreatePaymentTypeEft         RoutingDetailCreatePaymentType = "eft"
	RoutingDetailCreatePaymentTypeEthereum    RoutingDetailCreatePaymentType = "ethereum"
	RoutingDetailCreatePaymentTypeGBFps       RoutingDetailCreatePaymentType = "gb_fps"
	RoutingDetailCreatePaymentTypeHuIcs       RoutingDetailCreatePaymentType = "hu_ics"
	RoutingDetailCreatePaymentTypeInterac     RoutingDetailCreatePaymentType = "interac"
	RoutingDetailCreatePaymentTypeMasav       RoutingDetailCreatePaymentType = "masav"
	RoutingDetailCreatePaymentTypeMxCcen      RoutingDetailCreatePaymentType = "mx_ccen"
	RoutingDetailCreatePaymentTypeNeft        RoutingDetailCreatePaymentType = "neft"
	RoutingDetailCreatePaymentTypeNics        RoutingDetailCreatePaymentType = "nics"
	RoutingDetailCreatePaymentTypeNzBecs      RoutingDetailCreatePaymentType = "nz_becs"
	RoutingDetailCreatePaymentTypePlElixir    RoutingDetailCreatePaymentType = "pl_elixir"
	RoutingDetailCreatePaymentTypePolygon     RoutingDetailCreatePaymentType = "polygon"
	RoutingDetailCreatePaymentTypeProvxchange RoutingDetailCreatePaymentType = "provxchange"
	RoutingDetailCreatePaymentTypeRoSent      RoutingDetailCreatePaymentType = "ro_sent"
	RoutingDetailCreatePaymentTypeRtp         RoutingDetailCreatePaymentType = "rtp"
	RoutingDetailCreatePaymentTypeSeBankgirot RoutingDetailCreatePaymentType = "se_bankgirot"
	RoutingDetailCreatePaymentTypeSen         RoutingDetailCreatePaymentType = "sen"
	RoutingDetailCreatePaymentTypeSepa        RoutingDetailCreatePaymentType = "sepa"
	RoutingDetailCreatePaymentTypeSgGiro      RoutingDetailCreatePaymentType = "sg_giro"
	RoutingDetailCreatePaymentTypeSic         RoutingDetailCreatePaymentType = "sic"
	RoutingDetailCreatePaymentTypeSignet      RoutingDetailCreatePaymentType = "signet"
	RoutingDetailCreatePaymentTypeSknbi       RoutingDetailCreatePaymentType = "sknbi"
	RoutingDetailCreatePaymentTypeSolana      RoutingDetailCreatePaymentType = "solana"
	RoutingDetailCreatePaymentTypeWire        RoutingDetailCreatePaymentType = "wire"
	RoutingDetailCreatePaymentTypeZengin      RoutingDetailCreatePaymentType = "zengin"
)

func (r RoutingDetailCreatePaymentType) IsKnown() bool {
	switch r {
	case RoutingDetailCreatePaymentTypeACH, RoutingDetailCreatePaymentTypeAuBecs, RoutingDetailCreatePaymentTypeBacs, RoutingDetailCreatePaymentTypeBase, RoutingDetailCreatePaymentTypeBook, RoutingDetailCreatePaymentTypeCard, RoutingDetailCreatePaymentTypeChats, RoutingDetailCreatePaymentTypeCheck, RoutingDetailCreatePaymentTypeCrossBorder, RoutingDetailCreatePaymentTypeDkNets, RoutingDetailCreatePaymentTypeEft, RoutingDetailCreatePaymentTypeEthereum, RoutingDetailCreatePaymentTypeGBFps, RoutingDetailCreatePaymentTypeHuIcs, RoutingDetailCreatePaymentTypeInterac, RoutingDetailCreatePaymentTypeMasav, RoutingDetailCreatePaymentTypeMxCcen, RoutingDetailCreatePaymentTypeNeft, RoutingDetailCreatePaymentTypeNics, RoutingDetailCreatePaymentTypeNzBecs, RoutingDetailCreatePaymentTypePlElixir, RoutingDetailCreatePaymentTypePolygon, RoutingDetailCreatePaymentTypeProvxchange, RoutingDetailCreatePaymentTypeRoSent, RoutingDetailCreatePaymentTypeRtp, RoutingDetailCreatePaymentTypeSeBankgirot, RoutingDetailCreatePaymentTypeSen, RoutingDetailCreatePaymentTypeSepa, RoutingDetailCreatePaymentTypeSgGiro, RoutingDetailCreatePaymentTypeSic, RoutingDetailCreatePaymentTypeSignet, RoutingDetailCreatePaymentTypeSknbi, RoutingDetailCreatePaymentTypeSolana, RoutingDetailCreatePaymentTypeWire, RoutingDetailCreatePaymentTypeZengin:
		return true
	}
	return false
}

type RoutingDetailNewParams struct {
	RoutingDetailCreate RoutingDetailCreateParam `json:"routing_detail_create,required"`
}

func (r RoutingDetailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RoutingDetailCreate)
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
