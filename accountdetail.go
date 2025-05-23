// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// AccountDetailService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDetailService] method instead.
type AccountDetailService struct {
	Options []option.RequestOption
}

// NewAccountDetailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDetailService(opts ...option.RequestOption) (r *AccountDetailService) {
	r = &AccountDetailService{}
	r.Options = opts
	return
}

// Create an account detail for an external account.
func (r *AccountDetailService) New(ctx context.Context, accountsType AccountDetailNewParamsAccountsType, accountID string, body AccountDetailNewParams, opts ...option.RequestOption) (res *AccountDetail, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/account_details", accountsType, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single account detail for a single internal or external account.
func (r *AccountDetailService) Get(ctx context.Context, accountsType shared.AccountsType, accountID string, id string, opts ...option.RequestOption) (res *AccountDetail, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/account_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of account details for a single internal or external account.
func (r *AccountDetailService) List(ctx context.Context, accountsType shared.AccountsType, accountID string, query AccountDetailListParams, opts ...option.RequestOption) (res *pagination.Page[AccountDetail], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/account_details", accountsType, accountID)
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

// Get a list of account details for a single internal or external account.
func (r *AccountDetailService) ListAutoPaging(ctx context.Context, accountsType shared.AccountsType, accountID string, query AccountDetailListParams, opts ...option.RequestOption) *pagination.PageAutoPager[AccountDetail] {
	return pagination.NewPageAutoPager(r.List(ctx, accountsType, accountID, query, opts...))
}

// Delete a single account detail for an external account.
func (r *AccountDetailService) Delete(ctx context.Context, accountsType AccountDetailDeleteParamsAccountsType, accountID string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/account_details/%s", accountsType, accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccountDetail struct {
	ID string `json:"id,required" format:"uuid"`
	// The last 4 digits of the account_number.
	AccountNumberSafe string `json:"account_number_safe,required"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType AccountDetailAccountNumberType `json:"account_number_type,required"`
	CreatedAt         time.Time                      `json:"created_at,required" format:"date-time"`
	DiscardedAt       time.Time                      `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The account number for the bank account.
	AccountNumber string            `json:"account_number"`
	JSON          accountDetailJSON `json:"-"`
}

// accountDetailJSON contains the JSON metadata for the struct [AccountDetail]
type accountDetailJSON struct {
	ID                apijson.Field
	AccountNumberSafe apijson.Field
	AccountNumberType apijson.Field
	CreatedAt         apijson.Field
	DiscardedAt       apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	UpdatedAt         apijson.Field
	AccountNumber     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDetailJSON) RawJSON() string {
	return r.raw
}

// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
// account number is in a generic format.
type AccountDetailAccountNumberType string

const (
	AccountDetailAccountNumberTypeAuNumber        AccountDetailAccountNumberType = "au_number"
	AccountDetailAccountNumberTypeBaseAddress     AccountDetailAccountNumberType = "base_address"
	AccountDetailAccountNumberTypeClabe           AccountDetailAccountNumberType = "clabe"
	AccountDetailAccountNumberTypeEthereumAddress AccountDetailAccountNumberType = "ethereum_address"
	AccountDetailAccountNumberTypeHkNumber        AccountDetailAccountNumberType = "hk_number"
	AccountDetailAccountNumberTypeIban            AccountDetailAccountNumberType = "iban"
	AccountDetailAccountNumberTypeIDNumber        AccountDetailAccountNumberType = "id_number"
	AccountDetailAccountNumberTypeNzNumber        AccountDetailAccountNumberType = "nz_number"
	AccountDetailAccountNumberTypeOther           AccountDetailAccountNumberType = "other"
	AccountDetailAccountNumberTypePan             AccountDetailAccountNumberType = "pan"
	AccountDetailAccountNumberTypePolygonAddress  AccountDetailAccountNumberType = "polygon_address"
	AccountDetailAccountNumberTypeSgNumber        AccountDetailAccountNumberType = "sg_number"
	AccountDetailAccountNumberTypeSolanaAddress   AccountDetailAccountNumberType = "solana_address"
	AccountDetailAccountNumberTypeWalletAddress   AccountDetailAccountNumberType = "wallet_address"
)

func (r AccountDetailAccountNumberType) IsKnown() bool {
	switch r {
	case AccountDetailAccountNumberTypeAuNumber, AccountDetailAccountNumberTypeBaseAddress, AccountDetailAccountNumberTypeClabe, AccountDetailAccountNumberTypeEthereumAddress, AccountDetailAccountNumberTypeHkNumber, AccountDetailAccountNumberTypeIban, AccountDetailAccountNumberTypeIDNumber, AccountDetailAccountNumberTypeNzNumber, AccountDetailAccountNumberTypeOther, AccountDetailAccountNumberTypePan, AccountDetailAccountNumberTypePolygonAddress, AccountDetailAccountNumberTypeSgNumber, AccountDetailAccountNumberTypeSolanaAddress, AccountDetailAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

type AccountDetailNewParams struct {
	// The account number for the bank account.
	AccountNumber param.Field[string] `json:"account_number,required"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType param.Field[AccountDetailNewParamsAccountNumberType] `json:"account_number_type"`
}

func (r AccountDetailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDetailNewParamsAccountsType string

const (
	AccountDetailNewParamsAccountsTypeExternalAccounts AccountDetailNewParamsAccountsType = "external_accounts"
)

func (r AccountDetailNewParamsAccountsType) IsKnown() bool {
	switch r {
	case AccountDetailNewParamsAccountsTypeExternalAccounts:
		return true
	}
	return false
}

// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
// account number is in a generic format.
type AccountDetailNewParamsAccountNumberType string

const (
	AccountDetailNewParamsAccountNumberTypeAuNumber        AccountDetailNewParamsAccountNumberType = "au_number"
	AccountDetailNewParamsAccountNumberTypeBaseAddress     AccountDetailNewParamsAccountNumberType = "base_address"
	AccountDetailNewParamsAccountNumberTypeClabe           AccountDetailNewParamsAccountNumberType = "clabe"
	AccountDetailNewParamsAccountNumberTypeEthereumAddress AccountDetailNewParamsAccountNumberType = "ethereum_address"
	AccountDetailNewParamsAccountNumberTypeHkNumber        AccountDetailNewParamsAccountNumberType = "hk_number"
	AccountDetailNewParamsAccountNumberTypeIban            AccountDetailNewParamsAccountNumberType = "iban"
	AccountDetailNewParamsAccountNumberTypeIDNumber        AccountDetailNewParamsAccountNumberType = "id_number"
	AccountDetailNewParamsAccountNumberTypeNzNumber        AccountDetailNewParamsAccountNumberType = "nz_number"
	AccountDetailNewParamsAccountNumberTypeOther           AccountDetailNewParamsAccountNumberType = "other"
	AccountDetailNewParamsAccountNumberTypePan             AccountDetailNewParamsAccountNumberType = "pan"
	AccountDetailNewParamsAccountNumberTypePolygonAddress  AccountDetailNewParamsAccountNumberType = "polygon_address"
	AccountDetailNewParamsAccountNumberTypeSgNumber        AccountDetailNewParamsAccountNumberType = "sg_number"
	AccountDetailNewParamsAccountNumberTypeSolanaAddress   AccountDetailNewParamsAccountNumberType = "solana_address"
	AccountDetailNewParamsAccountNumberTypeWalletAddress   AccountDetailNewParamsAccountNumberType = "wallet_address"
)

func (r AccountDetailNewParamsAccountNumberType) IsKnown() bool {
	switch r {
	case AccountDetailNewParamsAccountNumberTypeAuNumber, AccountDetailNewParamsAccountNumberTypeBaseAddress, AccountDetailNewParamsAccountNumberTypeClabe, AccountDetailNewParamsAccountNumberTypeEthereumAddress, AccountDetailNewParamsAccountNumberTypeHkNumber, AccountDetailNewParamsAccountNumberTypeIban, AccountDetailNewParamsAccountNumberTypeIDNumber, AccountDetailNewParamsAccountNumberTypeNzNumber, AccountDetailNewParamsAccountNumberTypeOther, AccountDetailNewParamsAccountNumberTypePan, AccountDetailNewParamsAccountNumberTypePolygonAddress, AccountDetailNewParamsAccountNumberTypeSgNumber, AccountDetailNewParamsAccountNumberTypeSolanaAddress, AccountDetailNewParamsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

type AccountDetailListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [AccountDetailListParams]'s query parameters as
// `url.Values`.
func (r AccountDetailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDetailDeleteParamsAccountsType string

const (
	AccountDetailDeleteParamsAccountsTypeExternalAccounts AccountDetailDeleteParamsAccountsType = "external_accounts"
)

func (r AccountDetailDeleteParamsAccountsType) IsKnown() bool {
	switch r {
	case AccountDetailDeleteParamsAccountsTypeExternalAccounts:
		return true
	}
	return false
}
