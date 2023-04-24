package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerAccountNewParams struct {
	// The name of the ledger account.
	Name field.Field[string] `json:"name,required"`
	// The description of the ledger account.
	Description field.Field[string] `json:"description,nullable"`
	// The normal balance of the ledger account.
	NormalBalance field.Field[LedgerAccountNewParamsNormalBalance] `json:"normal_balance,required"`
	// The id of the ledger that this account belongs to.
	LedgerID field.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The currency of the ledger account.
	Currency field.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent field.Field[int64] `json:"currency_exponent,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r LedgerAccountNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerAccountNewParamsNormalBalance string

const (
	LedgerAccountNewParamsNormalBalanceCredit LedgerAccountNewParamsNormalBalance = "credit"
	LedgerAccountNewParamsNormalBalanceDebit  LedgerAccountNewParamsNormalBalance = "debit"
)

type LedgerAccountGetParams struct {
	// For example, if you want the balances as of a particular effective date
	// (YYYY-MM-DD), the encoded query string would be
	// balances%5Bas_of_date%5D=2000-12-31. The balances as of a date are inclusive of
	// entries with that exact date.
	Balances field.Field[LedgerAccountGetParamsBalances] `query:"balances"`
}

// URLQuery serializes LedgerAccountGetParams into a url.Values of the query
// parameters associated with this value
func (r LedgerAccountGetParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountGetParamsBalances struct {
	AsOfDate field.Field[time.Time] `query:"as_of_date" format:"date"`
}

// URLQuery serializes LedgerAccountGetParamsBalances into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountGetParamsBalances) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountUpdateParams struct {
	// The name of the ledger account.
	Name field.Field[string] `json:"name"`
	// The description of the ledger account.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r LedgerAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerAccountListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	ID       field.Field[string]            `query:"id"`
	Name     field.Field[string]            `query:"name"`
	LedgerID field.Field[string]            `query:"ledger_id"`
	// For example, if you want the balances as of a particular effective date
	// (YYYY-MM-DD), the encoded query string would be
	// balances%5Bas_of_date%5D=2000-12-31. The balances as of a date are inclusive of
	// entries with that exact date.
	Balances field.Field[LedgerAccountListParamsBalances] `query:"balances"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt               field.Field[map[string]time.Time] `query:"updated_at"`
	LedgerAccountCategoryID field.Field[string]               `query:"ledger_account_category_id"`
}

// URLQuery serializes LedgerAccountListParams into a url.Values of the query
// parameters associated with this value
func (r LedgerAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountListParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountListParamsBalances into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountListParamsBalances) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
