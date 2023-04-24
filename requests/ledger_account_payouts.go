package requests

import (
	"net/url"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type LedgerAccountPayoutNewParams struct {
	// The description of the ledger account payout.
	Description field.Field[string] `json:"description,nullable"`
	// The status of the ledger account payout. It is set to `pending` by default. To
	// post a ledger account payout at creation, use `posted`.
	Status field.Field[LedgerAccountPayoutNewParamsStatus] `json:"status,nullable"`
	// The id of the payout ledger account whose ledger entries are queried against,
	// and its balance is reduced as a result.
	PayoutLedgerAccountID field.Field[string] `json:"payout_ledger_account_id,required" format:"uuid"`
	// The id of the funding ledger account that sends to or receives funds from the
	// payout ledger account.
	FundingLedgerAccountID field.Field[string] `json:"funding_ledger_account_id,required" format:"uuid"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account payout. The default value is the created_at
	// timestamp of the ledger account payout.
	EffectiveAtUpperBound field.Field[string] `json:"effective_at_upper_bound,nullable" format:"time"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerAccountPayoutNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r LedgerAccountPayoutNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerAccountPayoutNewParamsStatus string

const (
	LedgerAccountPayoutNewParamsStatusPending LedgerAccountPayoutNewParamsStatus = "pending"
	LedgerAccountPayoutNewParamsStatusPosted  LedgerAccountPayoutNewParamsStatus = "posted"
)

type LedgerAccountPayoutUpdateParams struct {
	// The description of the ledger account payout.
	Description field.Field[string] `json:"description,nullable"`
	// To post a pending ledger account payout, use `posted`. To archive a pending
	// ledger transaction, use `archived`.
	Status field.Field[LedgerAccountPayoutUpdateParamsStatus] `json:"status"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LedgerAccountPayoutUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r LedgerAccountPayoutUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type LedgerAccountPayoutUpdateParamsStatus string

const (
	LedgerAccountPayoutUpdateParamsStatusPosted   LedgerAccountPayoutUpdateParamsStatus = "posted"
	LedgerAccountPayoutUpdateParamsStatusArchived LedgerAccountPayoutUpdateParamsStatus = "archived"
)

type LedgerAccountPayoutListParams struct {
	AfterCursor           field.Field[string] `query:"after_cursor,nullable"`
	PerPage               field.Field[int64]  `query:"per_page"`
	PayoutLedgerAccountID field.Field[string] `query:"payout_ledger_account_id"`
}

// URLQuery serializes LedgerAccountPayoutListParams into a url.Values of the query
// parameters associated with this value
func (r LedgerAccountPayoutListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
