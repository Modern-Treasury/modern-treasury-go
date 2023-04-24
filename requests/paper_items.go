package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type PaperItemListParams struct {
	// Specify `lockbox_number` if you wish to see paper items that are associated with
	// a specific lockbox number.
	LockboxNumber field.Field[string] `query:"lockbox_number"`
	// Specify an inclusive start date (YYYY-MM-DD) when filtering by deposit_date
	DepositDateStart field.Field[time.Time] `query:"deposit_date_start" format:"date"`
	// Specify an inclusive end date (YYYY-MM-DD) when filtering by deposit_date
	DepositDateEnd field.Field[time.Time] `query:"deposit_date_end" format:"date"`
	AfterCursor    field.Field[string]    `query:"after_cursor,nullable"`
	PerPage        field.Field[int64]     `query:"per_page"`
}

// URLQuery serializes PaperItemListParams into a url.Values of the query
// parameters associated with this value
func (r PaperItemListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
