package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerAccountCategoryNewParams struct {
	// The name of the ledger account category.
	Name field.Field[string] `json:"name,required"`
	// The description of the ledger account category.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The currency of the ledger account category.
	Currency field.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account category.
	CurrencyExponent field.Field[int64] `json:"currency_exponent,nullable"`
	// The id of the ledger that this account category belongs to.
	LedgerID field.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The normal balance of the ledger account category.
	NormalBalance field.Field[LedgerAccountCategoryNewParamsNormalBalance] `json:"normal_balance,required"`
}

// MarshalJSON serializes LedgerAccountCategoryNewParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r LedgerAccountCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerAccountCategoryNewParamsNormalBalance string

const (
	LedgerAccountCategoryNewParamsNormalBalanceCredit LedgerAccountCategoryNewParamsNormalBalance = "credit"
	LedgerAccountCategoryNewParamsNormalBalanceDebit  LedgerAccountCategoryNewParamsNormalBalance = "debit"
)

type LedgerAccountCategoryGetParams struct {
	// For example, if you want the balances as of a particular effective date
	// (YYYY-MM-DD), the encoded query string would be
	// balances%5Bas_of_date%5D=2000-12-31. The balances as of a date are inclusive of
	// entries with that exact date.
	Balances field.Field[LedgerAccountCategoryGetParamsBalances] `query:"balances"`
}

// URLQuery serializes LedgerAccountCategoryGetParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryGetParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountCategoryGetParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountCategoryGetParamsBalances into a url.Values of
// the query parameters associated with this value
func (r LedgerAccountCategoryGetParamsBalances) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountCategoryUpdateParams struct {
	// The name of the ledger account category.
	Name field.Field[string] `json:"name"`
	// The description of the ledger account category.
	Description field.Field[string] `json:"description,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerAccountCategoryUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r LedgerAccountCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

// URLQuery serializes LedgerAccountCategoryUpdateParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryUpdateParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountCategoryUpdateParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountCategoryUpdateParamsBalances into a url.Values
// of the query parameters associated with this value
func (r LedgerAccountCategoryUpdateParamsBalances) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountCategoryListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	Name     field.Field[string]            `query:"name"`
	LedgerID field.Field[string]            `query:"ledger_id"`
	// Query categories that are nested underneath a parent category
	ParentLedgerAccountCategoryID field.Field[string] `query:"parent_ledger_account_category_id"`
}

// URLQuery serializes LedgerAccountCategoryListParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountCategoryDeleteParams struct {
	// For example, if you want the balances as of a particular effective date
	// (YYYY-MM-DD), the encoded query string would be
	// balances%5Bas_of_date%5D=2000-12-31. The balances as of a date are inclusive of
	// entries with that exact date.
	Balances field.Field[LedgerAccountCategoryDeleteParamsBalances] `query:"balances"`
}

// URLQuery serializes LedgerAccountCategoryDeleteParams into a url.Values of the
// query parameters associated with this value
func (r LedgerAccountCategoryDeleteParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type LedgerAccountCategoryDeleteParamsBalances struct {
	AsOfDate    field.Field[time.Time] `query:"as_of_date" format:"date"`
	EffectiveAt field.Field[time.Time] `query:"effective_at" format:"date-time"`
}

// URLQuery serializes LedgerAccountCategoryDeleteParamsBalances into a url.Values
// of the query parameters associated with this value
func (r LedgerAccountCategoryDeleteParamsBalances) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
