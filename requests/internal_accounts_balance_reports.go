package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
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
	return query.Marshal(r)
}

type BalanceReportListParamsBalanceReportType string

const (
	BalanceReportListParamsBalanceReportTypeIntraday    BalanceReportListParamsBalanceReportType = "intraday"
	BalanceReportListParamsBalanceReportTypeOther       BalanceReportListParamsBalanceReportType = "other"
	BalanceReportListParamsBalanceReportTypePreviousDay BalanceReportListParamsBalanceReportType = "previous_day"
	BalanceReportListParamsBalanceReportTypeRealTime    BalanceReportListParamsBalanceReportType = "real_time"
)
