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

// InternalAccountBalanceReportService contains methods and other services that
// help with interacting with the Modern Treasury API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewInternalAccountBalanceReportService] method instead.
type InternalAccountBalanceReportService struct {
	Options []option.RequestOption
}

// NewInternalAccountBalanceReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInternalAccountBalanceReportService(opts ...option.RequestOption) (r *InternalAccountBalanceReportService) {
	r = &InternalAccountBalanceReportService{}
	r.Options = opts
	return
}

// create balance reports
func (r *InternalAccountBalanceReportService) New(ctx context.Context, internalAccountID string, body BalanceReportNewParams, opts ...option.RequestOption) (res *BalanceReport, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports", internalAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a single balance report for a given internal account.
func (r *InternalAccountBalanceReportService) Get(ctx context.Context, internalAccountID string, id string, opts ...option.RequestOption) (res *BalanceReport, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports/%s", internalAccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all balance reports for a given internal account.
func (r *InternalAccountBalanceReportService) List(ctx context.Context, internalAccountID string, query BalanceReportListParams, opts ...option.RequestOption) (res *shared.Page[BalanceReport], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports", internalAccountID)
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
func (r *InternalAccountBalanceReportService) ListAutoPaging(ctx context.Context, internalAccountID string, query BalanceReportListParams, opts ...option.RequestOption) *shared.PageAutoPager[BalanceReport] {
	return shared.NewPageAutoPager(r.List(ctx, internalAccountID, query, opts...))
}

// Deletes a given balance report.
func (r *InternalAccountBalanceReportService) Delete(ctx context.Context, internalAccountID string, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/internal_accounts/%s/balance_reports/%s", internalAccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BalanceReport struct {
	ID string `json:"id,required" format:"uuid"`
	// The date of the balance report in local time.
	AsOfDate time.Time `json:"as_of_date,required" format:"date"`
	// The time (24-hour clock) of the balance report in local time.
	AsOfTime string `json:"as_of_time,required,nullable" format:"time"`
	// The specific type of balance report. One of `intraday`, `previous_day`,
	// `real_time`, or `other`.
	BalanceReportType BalanceReportBalanceReportType `json:"balance_report_type,required"`
	// An array of `Balance` objects.
	Balances  []BalanceReportBalance `json:"balances,required"`
	CreatedAt time.Time              `json:"created_at,required" format:"date-time"`
	// The ID of one of your organization's Internal Accounts.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool              `json:"live_mode,required"`
	Object    string            `json:"object,required"`
	UpdatedAt time.Time         `json:"updated_at,required" format:"date-time"`
	JSON      balanceReportJSON `json:"-"`
}

// balanceReportJSON contains the JSON metadata for the struct [BalanceReport]
type balanceReportJSON struct {
	ID                apijson.Field
	AsOfDate          apijson.Field
	AsOfTime          apijson.Field
	BalanceReportType apijson.Field
	Balances          apijson.Field
	CreatedAt         apijson.Field
	InternalAccountID apijson.Field
	LiveMode          apijson.Field
	Object            apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BalanceReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceReportJSON) RawJSON() string {
	return r.raw
}

// The specific type of balance report. One of `intraday`, `previous_day`,
// `real_time`, or `other`.
type BalanceReportBalanceReportType string

const (
	BalanceReportBalanceReportTypeIntraday    BalanceReportBalanceReportType = "intraday"
	BalanceReportBalanceReportTypeOther       BalanceReportBalanceReportType = "other"
	BalanceReportBalanceReportTypePreviousDay BalanceReportBalanceReportType = "previous_day"
	BalanceReportBalanceReportTypeRealTime    BalanceReportBalanceReportType = "real_time"
)

type BalanceReportBalance struct {
	ID string `json:"id,required" format:"uuid"`
	// The balance amount.
	Amount int64 `json:"amount,required"`
	// The specific type of balance reported. One of `opening_ledger`,
	// `closing_ledger`, `current_ledger`, `opening_available`,
	// `opening_available_next_business_day`, `closing_available`, `current_available`,
	// or `other`.
	BalanceType BalanceReportBalancesBalanceType `json:"balance_type,required"`
	CreatedAt   time.Time                        `json:"created_at,required" format:"date-time"`
	// The currency of the balance.
	Currency shared.Currency `json:"currency,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The code used by the bank when reporting this specific balance.
	VendorCode string `json:"vendor_code,required"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, or `us_bank`.
	VendorCodeType string                   `json:"vendor_code_type,required,nullable"`
	JSON           balanceReportBalanceJSON `json:"-"`
}

// balanceReportBalanceJSON contains the JSON metadata for the struct
// [BalanceReportBalance]
type balanceReportBalanceJSON struct {
	ID             apijson.Field
	Amount         apijson.Field
	BalanceType    apijson.Field
	CreatedAt      apijson.Field
	Currency       apijson.Field
	LiveMode       apijson.Field
	Object         apijson.Field
	UpdatedAt      apijson.Field
	VendorCode     apijson.Field
	VendorCodeType apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BalanceReportBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceReportBalanceJSON) RawJSON() string {
	return r.raw
}

// The specific type of balance reported. One of `opening_ledger`,
// `closing_ledger`, `current_ledger`, `opening_available`,
// `opening_available_next_business_day`, `closing_available`, `current_available`,
// or `other`.
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

type BalanceReportNewParams struct {
	// The date of the balance report in local time.
	AsOfDate param.Field[time.Time] `json:"as_of_date,required" format:"date"`
	// The time (24-hour clock) of the balance report in local time.
	AsOfTime param.Field[string] `json:"as_of_time,required"`
	// The specific type of balance report. One of `intraday`, `previous_day`,
	// `real_time`, or `other`.
	BalanceReportType param.Field[BalanceReportNewParamsBalanceReportType] `json:"balance_report_type,required"`
	// An array of `Balance` objects.
	Balances param.Field[[]BalanceReportNewParamsBalance] `json:"balances,required"`
}

func (r BalanceReportNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The specific type of balance report. One of `intraday`, `previous_day`,
// `real_time`, or `other`.
type BalanceReportNewParamsBalanceReportType string

const (
	BalanceReportNewParamsBalanceReportTypeIntraday    BalanceReportNewParamsBalanceReportType = "intraday"
	BalanceReportNewParamsBalanceReportTypeOther       BalanceReportNewParamsBalanceReportType = "other"
	BalanceReportNewParamsBalanceReportTypePreviousDay BalanceReportNewParamsBalanceReportType = "previous_day"
	BalanceReportNewParamsBalanceReportTypeRealTime    BalanceReportNewParamsBalanceReportType = "real_time"
)

type BalanceReportNewParamsBalance struct {
	// The balance amount.
	Amount param.Field[int64] `json:"amount,required"`
	// The specific type of balance reported. One of `opening_ledger`,
	// `closing_ledger`, `current_ledger`, `opening_available`,
	// `opening_available_next_business_day`, `closing_available`, `current_available`,
	// or `other`.
	BalanceType param.Field[BalanceReportNewParamsBalancesBalanceType] `json:"balance_type,required"`
	// The code used by the bank when reporting this specific balance.
	VendorCode param.Field[string] `json:"vendor_code,required"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, or `us_bank`.
	VendorCodeType param.Field[string] `json:"vendor_code_type,required"`
}

func (r BalanceReportNewParamsBalance) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The specific type of balance reported. One of `opening_ledger`,
// `closing_ledger`, `current_ledger`, `opening_available`,
// `opening_available_next_business_day`, `closing_available`, `current_available`,
// or `other`.
type BalanceReportNewParamsBalancesBalanceType string

const (
	BalanceReportNewParamsBalancesBalanceTypeClosingAvailable                BalanceReportNewParamsBalancesBalanceType = "closing_available"
	BalanceReportNewParamsBalancesBalanceTypeClosingLedger                   BalanceReportNewParamsBalancesBalanceType = "closing_ledger"
	BalanceReportNewParamsBalancesBalanceTypeCurrentAvailable                BalanceReportNewParamsBalancesBalanceType = "current_available"
	BalanceReportNewParamsBalancesBalanceTypeCurrentLedger                   BalanceReportNewParamsBalancesBalanceType = "current_ledger"
	BalanceReportNewParamsBalancesBalanceTypeOpeningAvailable                BalanceReportNewParamsBalancesBalanceType = "opening_available"
	BalanceReportNewParamsBalancesBalanceTypeOpeningAvailableNextBusinessDay BalanceReportNewParamsBalancesBalanceType = "opening_available_next_business_day"
	BalanceReportNewParamsBalancesBalanceTypeOpeningLedger                   BalanceReportNewParamsBalancesBalanceType = "opening_ledger"
	BalanceReportNewParamsBalancesBalanceTypeOther                           BalanceReportNewParamsBalancesBalanceType = "other"
)

type BalanceReportListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// The date of the balance report in local time.
	AsOfDate param.Field[time.Time] `query:"as_of_date" format:"date"`
	// The specific type of balance report. One of `intraday`, `previous_day`,
	// `real_time`, or `other`.
	BalanceReportType param.Field[BalanceReportListParamsBalanceReportType] `query:"balance_report_type"`
	PerPage           param.Field[int64]                                    `query:"per_page"`
}

// URLQuery serializes [BalanceReportListParams]'s query parameters as
// `url.Values`.
func (r BalanceReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The specific type of balance report. One of `intraday`, `previous_day`,
// `real_time`, or `other`.
type BalanceReportListParamsBalanceReportType string

const (
	BalanceReportListParamsBalanceReportTypeIntraday    BalanceReportListParamsBalanceReportType = "intraday"
	BalanceReportListParamsBalanceReportTypeOther       BalanceReportListParamsBalanceReportType = "other"
	BalanceReportListParamsBalanceReportTypePreviousDay BalanceReportListParamsBalanceReportType = "previous_day"
	BalanceReportListParamsBalanceReportTypeRealTime    BalanceReportListParamsBalanceReportType = "real_time"
)
