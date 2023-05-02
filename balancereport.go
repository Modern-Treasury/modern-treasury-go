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

type InternalAccountBalanceReportService struct {
	Options []option.RequestOption
}

func NewInternalAccountBalanceReportService(opts ...option.RequestOption) (r *InternalAccountBalanceReportService) {
	r = &InternalAccountBalanceReportService{}
	r.Options = opts
	return
}

// Get a single balance report for a given internal account.
func (r *InternalAccountBalanceReportService) Get(ctx context.Context, internal_account_id string, id string, opts ...option.RequestOption) (res *BalanceReport, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports/%s", internal_account_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all balance reports for a given internal account.
func (r *InternalAccountBalanceReportService) List(ctx context.Context, internal_account_id string, query BalanceReportListParams, opts ...option.RequestOption) (res *shared.Page[BalanceReport], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports", internal_account_id)
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

// Get all balance reports for a given internal account.
func (r *InternalAccountBalanceReportService) ListAutoPaging(ctx context.Context, internal_account_id string, query BalanceReportListParams, opts ...option.RequestOption) *shared.PageAutoPager[BalanceReport] {
	return shared.NewPageAutoPager(r.List(ctx, internal_account_id, query, opts...))
}

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
	ID                apijson.Metadata
	Object            apijson.Metadata
	LiveMode          apijson.Metadata
	CreatedAt         apijson.Metadata
	UpdatedAt         apijson.Metadata
	BalanceReportType apijson.Metadata
	AsOfDate          apijson.Metadata
	AsOfTime          apijson.Metadata
	Balances          apijson.Metadata
	InternalAccountID apijson.Metadata
	raw               string
	Extras            map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceReport using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BalanceReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	Currency shared.Currency `json:"currency,required,nullable"`
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
	ID             apijson.Metadata
	Object         apijson.Metadata
	LiveMode       apijson.Metadata
	CreatedAt      apijson.Metadata
	UpdatedAt      apijson.Metadata
	Amount         apijson.Metadata
	Currency       apijson.Metadata
	BalanceType    apijson.Metadata
	VendorCode     apijson.Metadata
	VendorCodeType apijson.Metadata
	raw            string
	Extras         map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceReportBalances using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BalanceReportBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
	BalanceReportBalancesVendorCodeTypePlaid         BalanceReportBalancesVendorCodeType = "plaid"
	BalanceReportBalancesVendorCodeTypeRspecVendor   BalanceReportBalancesVendorCodeType = "rspec_vendor"
	BalanceReportBalancesVendorCodeTypeSignet        BalanceReportBalancesVendorCodeType = "signet"
	BalanceReportBalancesVendorCodeTypeSilvergate    BalanceReportBalancesVendorCodeType = "silvergate"
	BalanceReportBalancesVendorCodeTypeSwift         BalanceReportBalancesVendorCodeType = "swift"
	BalanceReportBalancesVendorCodeTypeUsBank        BalanceReportBalancesVendorCodeType = "us_bank"
)

type BalanceReportListParams struct {
	// The date of the balance report in local time.
	AsOfDate field.Field[time.Time] `query:"as_of_date" format:"date"`
	// The specific type of balance report. One of `intraday`, `previous_day`,
	// `real_time`, or `other`.
	BalanceReportType field.Field[BalanceReportListParamsBalanceReportType] `query:"balance_report_type"`
	AfterCursor       field.Field[string]                                   `query:"after_cursor,nullable"`
	PerPage           field.Field[int64]                                    `query:"per_page"`
}

// URLQuery serializes BalanceReportListParams into a url.Values of the query
// parameters associated with this value
func (r BalanceReportListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type BalanceReportListParamsBalanceReportType string

const (
	BalanceReportListParamsBalanceReportTypeIntraday    BalanceReportListParamsBalanceReportType = "intraday"
	BalanceReportListParamsBalanceReportTypeOther       BalanceReportListParamsBalanceReportType = "other"
	BalanceReportListParamsBalanceReportTypePreviousDay BalanceReportListParamsBalanceReportType = "previous_day"
	BalanceReportListParamsBalanceReportTypeRealTime    BalanceReportListParamsBalanceReportType = "real_time"
)
