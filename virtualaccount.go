// File generated from our OpenAPI spec by Stainless.

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

// VirtualAccountService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewVirtualAccountService] method instead.
type VirtualAccountService struct {
	Options []option.RequestOption
}

// NewVirtualAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVirtualAccountService(opts ...option.RequestOption) (r *VirtualAccountService) {
	r = &VirtualAccountService{}
	r.Options = opts
	return
}

// create virtual_account
func (r *VirtualAccountService) New(ctx context.Context, body VirtualAccountNewParams, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/virtual_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get virtual_account
func (r *VirtualAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update virtual_account
func (r *VirtualAccountService) Update(ctx context.Context, id string, body VirtualAccountUpdateParams, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of virtual accounts.
func (r *VirtualAccountService) List(ctx context.Context, query VirtualAccountListParams, opts ...option.RequestOption) (res *shared.Page[VirtualAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/virtual_accounts"
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

// Get a list of virtual accounts.
func (r *VirtualAccountService) ListAutoPaging(ctx context.Context, query VirtualAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[VirtualAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete virtual_account
func (r *VirtualAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type VirtualAccount struct {
	ID string `json:"id,required" format:"uuid"`
	// An array of account detail objects.
	AccountDetails []AccountDetail `json:"account_details,required"`
	// The ID of a counterparty that the virtual account belongs to. Optional.
	CounterpartyID string    `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// The ID of a credit normal ledger account. When money enters the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID string `json:"credit_ledger_account_id,required,nullable" format:"uuid"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID string `json:"debit_ledger_account_id,required,nullable" format:"uuid"`
	// An optional free-form description for internal use.
	Description string    `json:"description,required,nullable"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The ID of the internal account that the virtual account is in.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The name of the virtual account.
	Name   string `json:"name,required"`
	Object string `json:"object,required"`
	// An array of routing detail objects. These will be the routing details of the
	// internal account.
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	UpdatedAt      time.Time       `json:"updated_at,required" format:"date-time"`
	JSON           virtualAccountJSON
}

// virtualAccountJSON contains the JSON metadata for the struct [VirtualAccount]
type virtualAccountJSON struct {
	ID                    apijson.Field
	AccountDetails        apijson.Field
	CounterpartyID        apijson.Field
	CreatedAt             apijson.Field
	CreditLedgerAccountID apijson.Field
	DebitLedgerAccountID  apijson.Field
	Description           apijson.Field
	DiscardedAt           apijson.Field
	InternalAccountID     apijson.Field
	LiveMode              apijson.Field
	Metadata              apijson.Field
	Name                  apijson.Field
	Object                apijson.Field
	RoutingDetails        apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *VirtualAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r VirtualAccount) implementsPaymentOrderUltimateOriginatingAccount() {}

type VirtualAccountNewParams struct {
	// The ID of the internal account that this virtual account is associated with.
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// The name of the virtual account.
	Name param.Field[string] `json:"name,required"`
	// An array of account detail objects.
	AccountDetails param.Field[[]VirtualAccountNewParamsAccountDetail] `json:"account_details"`
	// The ID of the counterparty that the virtual account belongs to.
	CounterpartyID param.Field[string] `json:"counterparty_id" format:"uuid"`
	// The ID of a credit normal ledger account. When money leaves the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID param.Field[string] `json:"credit_ledger_account_id" format:"uuid"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID param.Field[string] `json:"debit_ledger_account_id" format:"uuid"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An array of routing detail objects.
	RoutingDetails param.Field[[]VirtualAccountNewParamsRoutingDetail] `json:"routing_details"`
}

func (r VirtualAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VirtualAccountNewParamsAccountDetail struct {
	// The account number for the bank account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType param.Field[VirtualAccountNewParamsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r VirtualAccountNewParamsAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
// account number is in a generic format.
type VirtualAccountNewParamsAccountDetailsAccountNumberType string

const (
	VirtualAccountNewParamsAccountDetailsAccountNumberTypeClabe         VirtualAccountNewParamsAccountDetailsAccountNumberType = "clabe"
	VirtualAccountNewParamsAccountDetailsAccountNumberTypeIban          VirtualAccountNewParamsAccountDetailsAccountNumberType = "iban"
	VirtualAccountNewParamsAccountDetailsAccountNumberTypeOther         VirtualAccountNewParamsAccountDetailsAccountNumberType = "other"
	VirtualAccountNewParamsAccountDetailsAccountNumberTypePan           VirtualAccountNewParamsAccountDetailsAccountNumberType = "pan"
	VirtualAccountNewParamsAccountDetailsAccountNumberTypeWalletAddress VirtualAccountNewParamsAccountDetailsAccountNumberType = "wallet_address"
)

type VirtualAccountNewParamsRoutingDetail struct {
	// The routing number of the bank.
	RoutingNumber param.Field[string] `json:"routing_number,required"`
	// The type of routing number. See
	// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
	// more details.
	RoutingNumberType param.Field[VirtualAccountNewParamsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	// If the routing detail is to be used for a specific payment type this field will
	// be populated, otherwise null.
	PaymentType param.Field[VirtualAccountNewParamsRoutingDetailsPaymentType] `json:"payment_type"`
}

func (r VirtualAccountNewParamsRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of routing number. See
// https://docs.moderntreasury.com/platform/reference/routing-detail-object for
// more details.
type VirtualAccountNewParamsRoutingDetailsRoutingNumberType string

const (
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAba                     VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "aba"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeAuBsb                   VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "au_bsb"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeBrCodigo                VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "br_codigo"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeCaCpa                   VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "ca_cpa"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeChips                   VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "chips"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeCnaps                   VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "cnaps"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeGBSortCode              VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "gb_sort_code"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeHkInterbankClearingCode VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeInIfsc                  VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "in_ifsc"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeJpZenginCode            VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "jp_zengin_code"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeMyBranchCode            VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "my_branch_code"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeNzNationalClearingCode  VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	VirtualAccountNewParamsRoutingDetailsRoutingNumberTypeSwift                   VirtualAccountNewParamsRoutingDetailsRoutingNumberType = "swift"
)

// If the routing detail is to be used for a specific payment type this field will
// be populated, otherwise null.
type VirtualAccountNewParamsRoutingDetailsPaymentType string

const (
	VirtualAccountNewParamsRoutingDetailsPaymentTypeACH         VirtualAccountNewParamsRoutingDetailsPaymentType = "ach"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeAuBecs      VirtualAccountNewParamsRoutingDetailsPaymentType = "au_becs"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeBacs        VirtualAccountNewParamsRoutingDetailsPaymentType = "bacs"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeBook        VirtualAccountNewParamsRoutingDetailsPaymentType = "book"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeCard        VirtualAccountNewParamsRoutingDetailsPaymentType = "card"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeChats       VirtualAccountNewParamsRoutingDetailsPaymentType = "chats"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeCheck       VirtualAccountNewParamsRoutingDetailsPaymentType = "check"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeCrossBorder VirtualAccountNewParamsRoutingDetailsPaymentType = "cross_border"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeEft         VirtualAccountNewParamsRoutingDetailsPaymentType = "eft"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeInterac     VirtualAccountNewParamsRoutingDetailsPaymentType = "interac"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeMasav       VirtualAccountNewParamsRoutingDetailsPaymentType = "masav"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeNeft        VirtualAccountNewParamsRoutingDetailsPaymentType = "neft"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeNics        VirtualAccountNewParamsRoutingDetailsPaymentType = "nics"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeNzBecs      VirtualAccountNewParamsRoutingDetailsPaymentType = "nz_becs"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeProvxchange VirtualAccountNewParamsRoutingDetailsPaymentType = "provxchange"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeRtp         VirtualAccountNewParamsRoutingDetailsPaymentType = "rtp"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeSeBankgirot VirtualAccountNewParamsRoutingDetailsPaymentType = "se_bankgirot"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeSen         VirtualAccountNewParamsRoutingDetailsPaymentType = "sen"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeSepa        VirtualAccountNewParamsRoutingDetailsPaymentType = "sepa"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeSic         VirtualAccountNewParamsRoutingDetailsPaymentType = "sic"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeSignet      VirtualAccountNewParamsRoutingDetailsPaymentType = "signet"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeWire        VirtualAccountNewParamsRoutingDetailsPaymentType = "wire"
	VirtualAccountNewParamsRoutingDetailsPaymentTypeZengin      VirtualAccountNewParamsRoutingDetailsPaymentType = "zengin"
)

type VirtualAccountUpdateParams struct {
	CounterpartyID param.Field[string]            `json:"counterparty_id" format:"uuid"`
	Metadata       param.Field[map[string]string] `json:"metadata"`
	Name           param.Field[string]            `json:"name"`
}

func (r VirtualAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VirtualAccountListParams struct {
	AfterCursor       param.Field[string] `query:"after_cursor"`
	CounterpartyID    param.Field[string] `query:"counterparty_id"`
	InternalAccountID param.Field[string] `query:"internal_account_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
}

// URLQuery serializes [VirtualAccountListParams]'s query parameters as
// `url.Values`.
func (r VirtualAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
