package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

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
	ID                pjson.Metadata
	Object            pjson.Metadata
	LiveMode          pjson.Metadata
	CreatedAt         pjson.Metadata
	UpdatedAt         pjson.Metadata
	DiscardedAt       pjson.Metadata
	AccountNumber     pjson.Metadata
	AccountNumberType pjson.Metadata
	AccountNumberSafe pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AccountDetail using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountDetail) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AccountDetailAccountNumberType string

const (
	AccountDetailAccountNumberTypeClabe         AccountDetailAccountNumberType = "clabe"
	AccountDetailAccountNumberTypeIban          AccountDetailAccountNumberType = "iban"
	AccountDetailAccountNumberTypeOther         AccountDetailAccountNumberType = "other"
	AccountDetailAccountNumberTypePan           AccountDetailAccountNumberType = "pan"
	AccountDetailAccountNumberTypeWalletAddress AccountDetailAccountNumberType = "wallet_address"
)
