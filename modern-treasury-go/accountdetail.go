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

type AccountDetailService struct {
	Options []option.RequestOption
}

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
	JSON              AccountDetailJSON
}

type AccountDetailJSON struct {
	ID                apijson.Metadata
	Object            apijson.Metadata
	LiveMode          apijson.Metadata
	CreatedAt         apijson.Metadata
	UpdatedAt         apijson.Metadata
	DiscardedAt       apijson.Metadata
	AccountNumber     apijson.Metadata
	AccountNumberType apijson.Metadata
	AccountNumberSafe apijson.Metadata
	raw               string
	Extras            map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountDetail using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDetailParam struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    field.Field[bool]      `json:"live_mode,required"`
	CreatedAt   field.Field[time.Time] `json:"created_at,required" format:"date-time"`
	UpdatedAt   field.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	DiscardedAt field.Field[time.Time] `json:"discarded_at,required,nullable" format:"date-time"`
	// The account number for the bank account.
	AccountNumber field.Field[string] `json:"account_number"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType field.Field[AccountDetailAccountNumberType] `json:"account_number_type,required"`
	// The last 4 digits of the account_number.
	AccountNumberSafe field.Field[string] `json:"account_number_safe,required"`
}

// MarshalJSON serializes AccountDetailParam into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
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
	AccountNumber field.Field[string] `json:"account_number,required"`
	// One of `iban`, `clabe`, `wallet_address`, or `other`. Use `other` if the bank
	// account number is in a generic format.
	AccountNumberType field.Field[AccountDetailNewParamsAccountNumberType] `json:"account_number_type"`
}

// MarshalJSON serializes AccountDetailNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
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
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes AccountDetailListParams into a url.Values of the query
// parameters associated with this value
func (r AccountDetailListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type AccountDetailDeleteParamsAccountsType string

const (
	AccountDetailDeleteParamsAccountsTypeExternalAccounts AccountDetailDeleteParamsAccountsType = "external_accounts"
)
