package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerNewParams struct {
	// The name of the ledger.
	Name field.Field[string] `json:"name,required"`
	// An optional free-form description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r LedgerNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerUpdateParams struct {
	// The name of the ledger.
	Name field.Field[string] `json:"name"`
	// An optional free-form description for internal use.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r LedgerUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt field.Field[map[string]time.Time] `query:"updated_at"`
}

// URLQuery serializes LedgerListParams into a url.Values of the query parameters
// associated with this value
func (r LedgerListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
