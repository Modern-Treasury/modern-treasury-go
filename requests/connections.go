package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type ConnectionListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// An identifier assigned by the vendor to your organization.
	VendorCustomerID field.Field[string] `query:"vendor_customer_id" format:"uuid"`
	// A string code representing the vendor (i.e. bank).
	Entity field.Field[string] `query:"entity"`
}

// URLQuery serializes ConnectionListParams into a url.Values of the query
// parameters associated with this value
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
