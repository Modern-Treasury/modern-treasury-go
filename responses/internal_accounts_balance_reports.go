package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type BalanceReport struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The specific type of balance report. One of `intraday`, `previous_day`,
	// `real_time`, or `other`.
	BalanceReportType BalanceReportBalanceReportType `json:"balance_report_type,required"`
	// The date of the balance report in local time.
	AsOfDate time.Time `json:"as_of_date,required" format:"date"`
	// The time (24-hour clock) of the balance report in local time.
	AsOfTime string `json:"as_of_time,required,nullable" format:"time"`
	// An array of `Balance` objects.
	Balances []BalanceReportBalances `json:"balances,required"`
	// The ID of one of your organization's Internal Accounts.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	JSON              BalanceReportJSON
}

type BalanceReportJSON struct {
	ID                pjson.Metadata
	Object            pjson.Metadata
	LiveMode          pjson.Metadata
	CreatedAt         pjson.Metadata
	UpdatedAt         pjson.Metadata
	BalanceReportType pjson.Metadata
	AsOfDate          pjson.Metadata
	AsOfTime          pjson.Metadata
	Balances          pjson.Metadata
	InternalAccountID pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceReport using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BalanceReport) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BalanceReportBalanceReportType string

const (
	BalanceReportBalanceReportTypeIntraday    BalanceReportBalanceReportType = "intraday"
	BalanceReportBalanceReportTypeOther       BalanceReportBalanceReportType = "other"
	BalanceReportBalanceReportTypePreviousDay BalanceReportBalanceReportType = "previous_day"
	BalanceReportBalanceReportTypeRealTime    BalanceReportBalanceReportType = "real_time"
)

type BalanceReportBalances struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The balance amount.
	Amount int64 `json:"amount,required"`
	// The currency of the balance.
	Currency Currency `json:"currency,required,nullable"`
	// The specific type of balance reported. One of `opening_ledger`,
	// `closing_ledger`, `current_ledger`, `opening_available`,
	// `opening_available_next_business_day`, `closing_available`, `current_available`,
	// or `other`.
	BalanceType BalanceReportBalancesBalanceType `json:"balance_type,required"`
	// The code used by the bank when reporting this specific balance.
	VendorCode string `json:"vendor_code,required"`
	// The code used by the bank when reporting this specific balance.
	VendorCodeType BalanceReportBalancesVendorCodeType `json:"vendor_code_type,required,nullable"`
	JSON           BalanceReportBalancesJSON
}

type BalanceReportBalancesJSON struct {
	ID             pjson.Metadata
	Object         pjson.Metadata
	LiveMode       pjson.Metadata
	CreatedAt      pjson.Metadata
	UpdatedAt      pjson.Metadata
	Amount         pjson.Metadata
	Currency       pjson.Metadata
	BalanceType    pjson.Metadata
	VendorCode     pjson.Metadata
	VendorCodeType pjson.Metadata
	Raw            []byte
	Extras         map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceReportBalances using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BalanceReportBalances) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BalanceReportBalancesBalanceType string

const (
	BalanceReportBalancesBalanceTypeClosingAvailable                BalanceReportBalancesBalanceType = "closing_available"
	BalanceReportBalancesBalanceTypeClosingLedger                   BalanceReportBalancesBalanceType = "closing_ledger"
	BalanceReportBalancesBalanceTypeCurrentAvailable                BalanceReportBalancesBalanceType = "current_available"
	BalanceReportBalancesBalanceTypeCurrentLedger                   BalanceReportBalancesBalanceType = "current_ledger"
	BalanceReportBalancesBalanceTypeOpeningAvailable                BalanceReportBalancesBalanceType = "opening_available"
	BalanceReportBalancesBalanceTypeOpeningAvailableNextBusinessDay BalanceReportBalancesBalanceType = "opening_available_next_business_day"
	BalanceReportBalancesBalanceTypeOpeningLedger                   BalanceReportBalancesBalanceType = "opening_ledger"
	BalanceReportBalancesBalanceTypeOther                           BalanceReportBalancesBalanceType = "other"
)

type BalanceReportBalancesVendorCodeType string

const (
	BalanceReportBalancesVendorCodeTypeBai2          BalanceReportBalancesVendorCodeType = "bai2"
	BalanceReportBalancesVendorCodeTypeBankprov      BalanceReportBalancesVendorCodeType = "bankprov"
	BalanceReportBalancesVendorCodeTypeBnkDev        BalanceReportBalancesVendorCodeType = "bnk_dev"
	BalanceReportBalancesVendorCodeTypeCleartouch    BalanceReportBalancesVendorCodeType = "cleartouch"
	BalanceReportBalancesVendorCodeTypeColumn        BalanceReportBalancesVendorCodeType = "column"
	BalanceReportBalancesVendorCodeTypeCrossRiver    BalanceReportBalancesVendorCodeType = "cross_river"
	BalanceReportBalancesVendorCodeTypeCurrencycloud BalanceReportBalancesVendorCodeType = "currencycloud"
	BalanceReportBalancesVendorCodeTypeDcBank        BalanceReportBalancesVendorCodeType = "dc_bank"
	BalanceReportBalancesVendorCodeTypeDwolla        BalanceReportBalancesVendorCodeType = "dwolla"
	BalanceReportBalancesVendorCodeTypeEvolve        BalanceReportBalancesVendorCodeType = "evolve"
	BalanceReportBalancesVendorCodeTypeGoldmanSachs  BalanceReportBalancesVendorCodeType = "goldman_sachs"
	BalanceReportBalancesVendorCodeTypeIso20022      BalanceReportBalancesVendorCodeType = "iso20022"
	BalanceReportBalancesVendorCodeTypeJpmc          BalanceReportBalancesVendorCodeType = "jpmc"
	BalanceReportBalancesVendorCodeTypeMx            BalanceReportBalancesVendorCodeType = "mx"
	BalanceReportBalancesVendorCodeTypeRspecVendor   BalanceReportBalancesVendorCodeType = "rspec_vendor"
	BalanceReportBalancesVendorCodeTypeSignet        BalanceReportBalancesVendorCodeType = "signet"
	BalanceReportBalancesVendorCodeTypeSilvergate    BalanceReportBalancesVendorCodeType = "silvergate"
	BalanceReportBalancesVendorCodeTypeSwift         BalanceReportBalancesVendorCodeType = "swift"
	BalanceReportBalancesVendorCodeTypeUsBank        BalanceReportBalancesVendorCodeType = "us_bank"
)
