// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
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
func (r *RoutingDetailService) New(ctx context.Context, accountsType RoutingDetailNewParamsAccountsType, accountID string, body RoutingDetailNewParams, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%v/%s/routing_details", accountsType, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single routing detail for a single internal or external account.
func (r *RoutingDetailService) Get(ctx context.Context, accountsType shared.AccountsType, accountID string, id string, opts ...option.RequestOption) (res *RoutingDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%v/%s/routing_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of routing details for a single internal or external account.
func (r *RoutingDetailService) List(ctx context.Context, accountsType shared.AccountsType, accountID string, query RoutingDetailListParams, opts ...option.RequestOption) (res *shared.Page[RoutingDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *RoutingDetailService) ListAutoPaging(ctx context.Context, accountsType shared.AccountsType, accountID string, query RoutingDetailListParams, opts ...option.RequestOption) *shared.PageAutoPager[RoutingDetail] {
	return shared.NewPageAutoPager(r.List(ctx, accountsType, accountID, query, opts...))
}

// Delete a routing detail for a single external account.
func (r *RoutingDetailService) Delete(ctx context.Context, accountsType RoutingDetailDeleteParamsAccountsType, accountID string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/%v/%s/routing_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type RoutingDetail struct {
	ID          string                   `json:"id,required" format:"uuid"`
	BankAddress RoutingDetailBankAddress `json:"bank_address,required,nullable"`
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

type RoutingDetailBankAddress struct {
	ID string `json:"id,required" format:"uuid"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country   string    `json:"country,required,nullable"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	Object   string `json:"object,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Region or State.
	Region    string                       `json:"region,required,nullable"`
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	JSON      routingDetailBankAddressJSON `json:"-"`
}

// routingDetailBankAddressJSON contains the JSON metadata for the struct
// [RoutingDetailBankAddress]
type routingDetailBankAddressJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	CreatedAt   apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	LiveMode    apijson.Field
	Locality    apijson.Field
	Object      apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingDetailBankAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingDetailBankAddressJSON) RawJSON() string {
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
	RoutingDetailPaymentTypeHuIcs       RoutingDetailPaymentType = "hu_ics"
	RoutingDetailPaymentTypeInterac     RoutingDetailPaymentType = "interac"
	RoutingDetailPaymentTypeMasav       RoutingDetailPaymentType = "masav"
	RoutingDetailPaymentTypeMxCcen      RoutingDetailPaymentType = "mx_ccen"
	RoutingDetailPaymentTypeNeft        RoutingDetailPaymentType = "neft"
	RoutingDetailPaymentTypeNics        RoutingDetailPaymentType = "nics"
	RoutingDetailPaymentTypeNzBecs      RoutingDetailPaymentType = "nz_becs"
	RoutingDetailPaymentTypePlElixir    RoutingDetailPaymentType = "pl_elixir"
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
	RoutingDetailPaymentTypeWire        RoutingDetailPaymentType = "wire"
	RoutingDetailPaymentTypeZengin      RoutingDetailPaymentType = "zengin"
)

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
	RoutingDetailRoutingNumberTypeInIfsc                  RoutingDetailRoutingNumberType = "in_ifsc"
	RoutingDetailRoutingNumberTypeJpZenginCode            RoutingDetailRoutingNumberType = "jp_zengin_code"
	RoutingDetailRoutingNumberTypeMxBankIdentifier        RoutingDetailRoutingNumberType = "mx_bank_identifier"
	RoutingDetailRoutingNumberTypeMyBranchCode            RoutingDetailRoutingNumberType = "my_branch_code"
	RoutingDetailRoutingNumberTypeNzNationalClearingCode  RoutingDetailRoutingNumberType = "nz_national_clearing_code"
	RoutingDetailRoutingNumberTypePlNationalClearingCode  RoutingDetailRoutingNumberType = "pl_national_clearing_code"
	RoutingDetailRoutingNumberTypeSeBankgiroClearingCode  RoutingDetailRoutingNumberType = "se_bankgiro_clearing_code"
	RoutingDetailRoutingNumberTypeSwift                   RoutingDetailRoutingNumberType = "swift"
)

type RoutingDetailNewParams struct {
	// The routing number of the bank.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details.
	RoutingNumberType param.Field[RoutingDetailNewParamsRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType param.Field[RoutingDetailNewParamsPaymentType] `json:"payment_type"`
}

func (r RoutingDetailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoutingDetailNewParamsAccountsType string

const (
	RoutingDetailNewParamsAccountsTypeExternalAccounts RoutingDetailNewParamsAccountsType = "external_accounts"
)

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details.
type RoutingDetailNewParamsRoutingNumberType string

const (
	RoutingDetailNewParamsRoutingNumberTypeAba                     RoutingDetailNewParamsRoutingNumberType = "aba"
	RoutingDetailNewParamsRoutingNumberTypeAuBsb                   RoutingDetailNewParamsRoutingNumberType = "au_bsb"
	RoutingDetailNewParamsRoutingNumberTypeBrCodigo                RoutingDetailNewParamsRoutingNumberType = "br_codigo"
	RoutingDetailNewParamsRoutingNumberTypeCaCpa                   RoutingDetailNewParamsRoutingNumberType = "ca_cpa"
	RoutingDetailNewParamsRoutingNumberTypeChips                   RoutingDetailNewParamsRoutingNumberType = "chips"
	RoutingDetailNewParamsRoutingNumberTypeCnaps                   RoutingDetailNewParamsRoutingNumberType = "cnaps"
	RoutingDetailNewParamsRoutingNumberTypeDkInterbankClearingCode RoutingDetailNewParamsRoutingNumberType = "dk_interbank_clearing_code"
	RoutingDetailNewParamsRoutingNumberTypeGBSortCode              RoutingDetailNewParamsRoutingNumberType = "gb_sort_code"
	RoutingDetailNewParamsRoutingNumberTypeHkInterbankClearingCode RoutingDetailNewParamsRoutingNumberType = "hk_interbank_clearing_code"
	RoutingDetailNewParamsRoutingNumberTypeHuInterbankClearingCode RoutingDetailNewParamsRoutingNumberType = "hu_interbank_clearing_code"
	RoutingDetailNewParamsRoutingNumberTypeIDSknbiCode             RoutingDetailNewParamsRoutingNumberType = "id_sknbi_code"
	RoutingDetailNewParamsRoutingNumberTypeInIfsc                  RoutingDetailNewParamsRoutingNumberType = "in_ifsc"
	RoutingDetailNewParamsRoutingNumberTypeJpZenginCode            RoutingDetailNewParamsRoutingNumberType = "jp_zengin_code"
	RoutingDetailNewParamsRoutingNumberTypeMxBankIdentifier        RoutingDetailNewParamsRoutingNumberType = "mx_bank_identifier"
	RoutingDetailNewParamsRoutingNumberTypeMyBranchCode            RoutingDetailNewParamsRoutingNumberType = "my_branch_code"
	RoutingDetailNewParamsRoutingNumberTypeNzNationalClearingCode  RoutingDetailNewParamsRoutingNumberType = "nz_national_clearing_code"
	RoutingDetailNewParamsRoutingNumberTypePlNationalClearingCode  RoutingDetailNewParamsRoutingNumberType = "pl_national_clearing_code"
	RoutingDetailNewParamsRoutingNumberTypeSeBankgiroClearingCode  RoutingDetailNewParamsRoutingNumberType = "se_bankgiro_clearing_code"
	RoutingDetailNewParamsRoutingNumberTypeSwift                   RoutingDetailNewParamsRoutingNumberType = "swift"
)

// If the routing detail is to be used for a specific payment type this field will
// be populated, otherwise null.
type RoutingDetailNewParamsPaymentType string

const (
	RoutingDetailNewParamsPaymentTypeACH         RoutingDetailNewParamsPaymentType = "ach"
	RoutingDetailNewParamsPaymentTypeAuBecs      RoutingDetailNewParamsPaymentType = "au_becs"
	RoutingDetailNewParamsPaymentTypeBacs        RoutingDetailNewParamsPaymentType = "bacs"
	RoutingDetailNewParamsPaymentTypeBook        RoutingDetailNewParamsPaymentType = "book"
	RoutingDetailNewParamsPaymentTypeCard        RoutingDetailNewParamsPaymentType = "card"
	RoutingDetailNewParamsPaymentTypeChats       RoutingDetailNewParamsPaymentType = "chats"
	RoutingDetailNewParamsPaymentTypeCheck       RoutingDetailNewParamsPaymentType = "check"
	RoutingDetailNewParamsPaymentTypeCrossBorder RoutingDetailNewParamsPaymentType = "cross_border"
	RoutingDetailNewParamsPaymentTypeDkNets      RoutingDetailNewParamsPaymentType = "dk_nets"
	RoutingDetailNewParamsPaymentTypeEft         RoutingDetailNewParamsPaymentType = "eft"
	RoutingDetailNewParamsPaymentTypeHuIcs       RoutingDetailNewParamsPaymentType = "hu_ics"
	RoutingDetailNewParamsPaymentTypeInterac     RoutingDetailNewParamsPaymentType = "interac"
	RoutingDetailNewParamsPaymentTypeMasav       RoutingDetailNewParamsPaymentType = "masav"
	RoutingDetailNewParamsPaymentTypeMxCcen      RoutingDetailNewParamsPaymentType = "mx_ccen"
	RoutingDetailNewParamsPaymentTypeNeft        RoutingDetailNewParamsPaymentType = "neft"
	RoutingDetailNewParamsPaymentTypeNics        RoutingDetailNewParamsPaymentType = "nics"
	RoutingDetailNewParamsPaymentTypeNzBecs      RoutingDetailNewParamsPaymentType = "nz_becs"
	RoutingDetailNewParamsPaymentTypePlElixir    RoutingDetailNewParamsPaymentType = "pl_elixir"
	RoutingDetailNewParamsPaymentTypeProvxchange RoutingDetailNewParamsPaymentType = "provxchange"
	RoutingDetailNewParamsPaymentTypeRoSent      RoutingDetailNewParamsPaymentType = "ro_sent"
	RoutingDetailNewParamsPaymentTypeRtp         RoutingDetailNewParamsPaymentType = "rtp"
	RoutingDetailNewParamsPaymentTypeSeBankgirot RoutingDetailNewParamsPaymentType = "se_bankgirot"
	RoutingDetailNewParamsPaymentTypeSen         RoutingDetailNewParamsPaymentType = "sen"
	RoutingDetailNewParamsPaymentTypeSepa        RoutingDetailNewParamsPaymentType = "sepa"
	RoutingDetailNewParamsPaymentTypeSgGiro      RoutingDetailNewParamsPaymentType = "sg_giro"
	RoutingDetailNewParamsPaymentTypeSic         RoutingDetailNewParamsPaymentType = "sic"
	RoutingDetailNewParamsPaymentTypeSignet      RoutingDetailNewParamsPaymentType = "signet"
	RoutingDetailNewParamsPaymentTypeSknbi       RoutingDetailNewParamsPaymentType = "sknbi"
	RoutingDetailNewParamsPaymentTypeWire        RoutingDetailNewParamsPaymentType = "wire"
	RoutingDetailNewParamsPaymentTypeZengin      RoutingDetailNewParamsPaymentType = "zengin"
)

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
