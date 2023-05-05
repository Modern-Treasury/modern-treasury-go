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

// AccountDetailService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewAccountDetailService]
// method instead.
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
func (r *AccountDetailService) New(ctx context.Context, accounts_type AccountDetailNewParamsAccountsType, account_id string, body AccountDetailNewParams, opts ...option.RequestOption) (res *AccountDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/account_details", accounts_type, account_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single account detail for a single internal or external account.
func (r *AccountDetailService) Get(ctx context.Context, accounts_type shared.AccountsType, account_id string, id string, opts ...option.RequestOption) (res *AccountDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/account_details/%s", accounts_type, account_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of account details for a single internal or external account.
func (r *AccountDetailService) List(ctx context.Context, accounts_type shared.AccountsType, account_id string, query AccountDetailListParams, opts ...option.RequestOption) (res *shared.Page[AccountDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/account_details", accounts_type, account_id)
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
func (r *AccountDetailService) ListAutoPaging(ctx context.Context, accounts_type shared.AccountsType, account_id string, query AccountDetailListParams, opts ...option.RequestOption) *shared.PageAutoPager[AccountDetail] {
	return shared.NewPageAutoPager(r.List(ctx, accounts_type, account_id, query, opts...))
}

// Delete a single account detail for an external account.
func (r *AccountDetailService) Delete(ctx context.Context, accounts_type AccountDetailDeleteParamsAccountsType, account_id string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/%s/%s/account_details/%s", accounts_type, account_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccountDetail struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The account number for the bank account.
	AccountNumber string `json:"account_number"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType AccountDetailAccountNumberType `json:"account_number_type,required"`
	// The last 4 digits of the account_number.
	AccountNumberSafe string `json:"account_number_safe,required"`
	JSON              accountDetailJSON
}

// accountDetailJSON contains the JSON metadata for the struct [AccountDetail]
type accountDetailJSON struct {
	ID                apijson.Field
	Object            apijson.Field
	LiveMode          apijson.Field
	CreatedAt         apijson.Field
	UpdatedAt         apijson.Field
	DiscardedAt       apijson.Field
	AccountNumber     apijson.Field
	AccountNumberType apijson.Field
	AccountNumberSafe apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDetailParam struct {
	ID     param.Field[string] `json:"id,required" format:"uuid"`
	Object param.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    param.Field[bool]      `json:"live_mode,required"`
	CreatedAt   param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	DiscardedAt param.Field[time.Time] `json:"discarded_at,required,nullable" format:"date-time"`
	// The account number for the bank account.
	AccountNumber param.Field[string] `json:"account_number"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType param.Field[AccountDetailAccountNumberType] `json:"account_number_type,required"`
	// The last 4 digits of the account_number.
	AccountNumberSafe param.Field[string] `json:"account_number_safe,required"`
}

func (r AccountDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDetailAccountNumberType string

const (
	AccountDetailAccountNumberTypeClabe         AccountDetailAccountNumberType = "clabe"
	AccountDetailAccountNumberTypeIban          AccountDetailAccountNumberType = "iban"
	AccountDetailAccountNumberTypeOther         AccountDetailAccountNumberType = "other"
	AccountDetailAccountNumberTypePan           AccountDetailAccountNumberType = "pan"
	AccountDetailAccountNumberTypeWalletAddress AccountDetailAccountNumberType = "wallet_address"
)

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

type AccountDetailNewParamsAccountNumberType string

const (
	AccountDetailNewParamsAccountNumberTypeClabe         AccountDetailNewParamsAccountNumberType = "clabe"
	AccountDetailNewParamsAccountNumberTypeIban          AccountDetailNewParamsAccountNumberType = "iban"
	AccountDetailNewParamsAccountNumberTypeOther         AccountDetailNewParamsAccountNumberType = "other"
	AccountDetailNewParamsAccountNumberTypePan           AccountDetailNewParamsAccountNumberType = "pan"
	AccountDetailNewParamsAccountNumberTypeWalletAddress AccountDetailNewParamsAccountNumberType = "wallet_address"
)

type AccountDetailNewParamsAccountsType string

const (
	AccountDetailNewParamsAccountsTypeExternalAccounts AccountDetailNewParamsAccountsType = "external_accounts"
)

type AccountDetailListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [AccountDetailListParams]'s query parameters as
// `url.Values`.
func (r AccountDetailListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type AccountDetailDeleteParamsAccountsType string

const (
	AccountDetailDeleteParamsAccountsTypeExternalAccounts AccountDetailDeleteParamsAccountsType = "external_accounts"
)
