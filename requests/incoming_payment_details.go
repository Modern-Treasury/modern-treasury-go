package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type IncomingPaymentDetailUpdateParams struct {
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes IncomingPaymentDetailUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r IncomingPaymentDetailUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type IncomingPaymentDetailListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// One of `credit` or `debit`.
	Direction field.Field[IncomingPaymentDetailListParamsDirection] `query:"direction"`
	// The current status of the incoming payment order. One of `pending`, `completed`,
	// or `returned`.
	Status field.Field[IncomingPaymentDetailListParamsStatus] `query:"status"`
	// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
	// `wire`.
	Type field.Field[IncomingPaymentDetailListParamsType] `query:"type"`
	// Filters incoming payment details with an as_of_date starting on or after the
	// specified date (YYYY-MM-DD).
	AsOfDateStart field.Field[time.Time] `query:"as_of_date_start" format:"date"`
	// Filters incoming payment details with an as_of_date starting on or before the
	// specified date (YYYY-MM-DD).
	AsOfDateEnd field.Field[time.Time] `query:"as_of_date_end" format:"date"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
	// If the incoming payment detail is in a virtual account, the ID of the Virtual
	// Account.
	VirtualAccountID field.Field[string] `query:"virtual_account_id"`
}

// URLQuery serializes IncomingPaymentDetailListParams into a url.Values of the
// query parameters associated with this value
func (r IncomingPaymentDetailListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type IncomingPaymentDetailListParamsDirection string

const (
	IncomingPaymentDetailListParamsDirectionCredit IncomingPaymentDetailListParamsDirection = "credit"
	IncomingPaymentDetailListParamsDirectionDebit  IncomingPaymentDetailListParamsDirection = "debit"
)

type IncomingPaymentDetailListParamsStatus string

const (
	IncomingPaymentDetailListParamsStatusCompleted IncomingPaymentDetailListParamsStatus = "completed"
	IncomingPaymentDetailListParamsStatusPending   IncomingPaymentDetailListParamsStatus = "pending"
	IncomingPaymentDetailListParamsStatusReturned  IncomingPaymentDetailListParamsStatus = "returned"
)

type IncomingPaymentDetailListParamsType string

const (
	IncomingPaymentDetailListParamsTypeACH     IncomingPaymentDetailListParamsType = "ach"
	IncomingPaymentDetailListParamsTypeBook    IncomingPaymentDetailListParamsType = "book"
	IncomingPaymentDetailListParamsTypeCheck   IncomingPaymentDetailListParamsType = "check"
	IncomingPaymentDetailListParamsTypeEft     IncomingPaymentDetailListParamsType = "eft"
	IncomingPaymentDetailListParamsTypeInterac IncomingPaymentDetailListParamsType = "interac"
	IncomingPaymentDetailListParamsTypeRtp     IncomingPaymentDetailListParamsType = "rtp"
	IncomingPaymentDetailListParamsTypeSepa    IncomingPaymentDetailListParamsType = "sepa"
	IncomingPaymentDetailListParamsTypeSignet  IncomingPaymentDetailListParamsType = "signet"
	IncomingPaymentDetailListParamsTypeWire    IncomingPaymentDetailListParamsType = "wire"
)

type IncomingPaymentDetailNewAsyncParams struct {
	// One of `ach`, `wire`, `check`.
	Type field.Field[IncomingPaymentDetailNewAsyncParamsType] `json:"type"`
	// One of `credit`, `debit`.
	Direction field.Field[IncomingPaymentDetailNewAsyncParamsDirection] `json:"direction"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount field.Field[int64] `json:"amount"`
	// Defaults to the currency of the originating account.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// The ID of one of your internal accounts.
	InternalAccountID field.Field[string] `json:"internal_account_id" format:"uuid"`
	// An optional parameter to associate the incoming payment detail to a virtual
	// account.
	VirtualAccountID field.Field[string] `json:"virtual_account_id,nullable" format:"uuid"`
	// Defaults to today.
	AsOfDate field.Field[time.Time] `json:"as_of_date,nullable" format:"date"`
}

// MarshalJSON serializes IncomingPaymentDetailNewAsyncParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r IncomingPaymentDetailNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type IncomingPaymentDetailNewAsyncParamsType string

const (
	IncomingPaymentDetailNewAsyncParamsTypeACH     IncomingPaymentDetailNewAsyncParamsType = "ach"
	IncomingPaymentDetailNewAsyncParamsTypeBook    IncomingPaymentDetailNewAsyncParamsType = "book"
	IncomingPaymentDetailNewAsyncParamsTypeCheck   IncomingPaymentDetailNewAsyncParamsType = "check"
	IncomingPaymentDetailNewAsyncParamsTypeEft     IncomingPaymentDetailNewAsyncParamsType = "eft"
	IncomingPaymentDetailNewAsyncParamsTypeInterac IncomingPaymentDetailNewAsyncParamsType = "interac"
	IncomingPaymentDetailNewAsyncParamsTypeRtp     IncomingPaymentDetailNewAsyncParamsType = "rtp"
	IncomingPaymentDetailNewAsyncParamsTypeSepa    IncomingPaymentDetailNewAsyncParamsType = "sepa"
	IncomingPaymentDetailNewAsyncParamsTypeSignet  IncomingPaymentDetailNewAsyncParamsType = "signet"
	IncomingPaymentDetailNewAsyncParamsTypeWire    IncomingPaymentDetailNewAsyncParamsType = "wire"
)

type IncomingPaymentDetailNewAsyncParamsDirection string

const (
	IncomingPaymentDetailNewAsyncParamsDirectionCredit IncomingPaymentDetailNewAsyncParamsDirection = "credit"
	IncomingPaymentDetailNewAsyncParamsDirectionDebit  IncomingPaymentDetailNewAsyncParamsDirection = "debit"
)
