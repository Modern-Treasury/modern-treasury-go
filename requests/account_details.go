package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type AccountDetail struct {
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

// MarshalJSON serializes AccountDetail into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r AccountDetail) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
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
	return pjson.MarshalRoot(r)
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
	return query.Marshal(r)
}

type AccountDetailDeleteParamsAccountsType string

const (
	AccountDetailDeleteParamsAccountsTypeExternalAccounts AccountDetailDeleteParamsAccountsType = "external_accounts"
)
