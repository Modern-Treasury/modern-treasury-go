package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerTransactionVersionListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// created_at timestamp. For example, for all dates after Jan 1 2000 12:00 UTC, use
	// created_at%5Bgt%5D=2000-01-01T12:00:00Z.
	CreatedAt field.Field[map[string]time.Time] `query:"created_at"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// version. For example, for all versions after 2, use version%5Bgt%5D=2.
	Version field.Field[map[string]int64] `query:"version"`
}

// URLQuery serializes LedgerTransactionVersionListParams into a url.Values of the
// query parameters associated with this value
func (r LedgerTransactionVersionListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
