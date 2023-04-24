package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type ReturnNewParams struct {
	// The ID of the object being returned or `null`.
	ReturnableID field.Field[string] `json:"returnable_id,required,nullable" format:"uuid"`
	// The return code. For ACH returns, this is the required ACH return code.
	Code field.Field[ReturnNewParamsCode] `json:"code,nullable"`
	// An optional description of the reason for the return. This is for internal usage
	// and will not be transmitted to the bank.”
	Reason field.Field[string] `json:"reason,nullable"`
	// If the return code is `R14` or `R15` this is the date the deceased counterparty
	// passed away.
	DateOfDeath field.Field[time.Time] `json:"date_of_death,nullable" format:"date"`
	// Some returns may include additional information from the bank. In these cases,
	// this string will be present.
	AdditionalInformation field.Field[string] `json:"additional_information,nullable"`
	// The type of object being returned. Currently, this may only be
	// incoming_payment_detail.
	ReturnableType field.Field[ReturnNewParamsReturnableType] `json:"returnable_type,required"`
}

// MarshalJSON serializes ReturnNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r ReturnNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type ReturnNewParamsCode string

const (
	ReturnNewParamsCode901           ReturnNewParamsCode = "901"
	ReturnNewParamsCode902           ReturnNewParamsCode = "902"
	ReturnNewParamsCode903           ReturnNewParamsCode = "903"
	ReturnNewParamsCode904           ReturnNewParamsCode = "904"
	ReturnNewParamsCode905           ReturnNewParamsCode = "905"
	ReturnNewParamsCode907           ReturnNewParamsCode = "907"
	ReturnNewParamsCode908           ReturnNewParamsCode = "908"
	ReturnNewParamsCode909           ReturnNewParamsCode = "909"
	ReturnNewParamsCode910           ReturnNewParamsCode = "910"
	ReturnNewParamsCode911           ReturnNewParamsCode = "911"
	ReturnNewParamsCode912           ReturnNewParamsCode = "912"
	ReturnNewParamsCode914           ReturnNewParamsCode = "914"
	ReturnNewParamsCodeR01           ReturnNewParamsCode = "R01"
	ReturnNewParamsCodeR02           ReturnNewParamsCode = "R02"
	ReturnNewParamsCodeR03           ReturnNewParamsCode = "R03"
	ReturnNewParamsCodeR04           ReturnNewParamsCode = "R04"
	ReturnNewParamsCodeR05           ReturnNewParamsCode = "R05"
	ReturnNewParamsCodeR06           ReturnNewParamsCode = "R06"
	ReturnNewParamsCodeR07           ReturnNewParamsCode = "R07"
	ReturnNewParamsCodeR08           ReturnNewParamsCode = "R08"
	ReturnNewParamsCodeR09           ReturnNewParamsCode = "R09"
	ReturnNewParamsCodeR10           ReturnNewParamsCode = "R10"
	ReturnNewParamsCodeR11           ReturnNewParamsCode = "R11"
	ReturnNewParamsCodeR12           ReturnNewParamsCode = "R12"
	ReturnNewParamsCodeR14           ReturnNewParamsCode = "R14"
	ReturnNewParamsCodeR15           ReturnNewParamsCode = "R15"
	ReturnNewParamsCodeR16           ReturnNewParamsCode = "R16"
	ReturnNewParamsCodeR17           ReturnNewParamsCode = "R17"
	ReturnNewParamsCodeR20           ReturnNewParamsCode = "R20"
	ReturnNewParamsCodeR21           ReturnNewParamsCode = "R21"
	ReturnNewParamsCodeR22           ReturnNewParamsCode = "R22"
	ReturnNewParamsCodeR23           ReturnNewParamsCode = "R23"
	ReturnNewParamsCodeR24           ReturnNewParamsCode = "R24"
	ReturnNewParamsCodeR29           ReturnNewParamsCode = "R29"
	ReturnNewParamsCodeR31           ReturnNewParamsCode = "R31"
	ReturnNewParamsCodeR33           ReturnNewParamsCode = "R33"
	ReturnNewParamsCodeR37           ReturnNewParamsCode = "R37"
	ReturnNewParamsCodeR38           ReturnNewParamsCode = "R38"
	ReturnNewParamsCodeR39           ReturnNewParamsCode = "R39"
	ReturnNewParamsCodeR51           ReturnNewParamsCode = "R51"
	ReturnNewParamsCodeR52           ReturnNewParamsCode = "R52"
	ReturnNewParamsCodeR53           ReturnNewParamsCode = "R53"
	ReturnNewParamsCodeCurrencycloud ReturnNewParamsCode = "currencycloud"
)

type ReturnNewParamsReturnableType string

const (
	ReturnNewParamsReturnableTypeIncomingPaymentDetail ReturnNewParamsReturnableType = "incoming_payment_detail"
)

type ReturnListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Specify `internal_account_id` if you wish to see returns to/from a specific
	// account.
	InternalAccountID field.Field[string] `query:"internal_account_id"`
	// Specify `counterparty_id` if you wish to see returns that occurred with a
	// specific counterparty.
	CounterpartyID field.Field[string] `query:"counterparty_id"`
	// The ID of a valid returnable. Must be accompanied by `returnable_type`.
	ReturnableID field.Field[string] `query:"returnable_id"`
	// One of `payment_order`, `paper_item`, `reversal`, or `incoming_payment_detail`.
	// Must be accompanied by `returnable_id`.
	ReturnableType field.Field[ReturnListParamsReturnableType] `query:"returnable_type"`
}

// URLQuery serializes ReturnListParams into a url.Values of the query parameters
// associated with this value
func (r ReturnListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type ReturnListParamsReturnableType string

const (
	ReturnListParamsReturnableTypeIncomingPaymentDetail ReturnListParamsReturnableType = "incoming_payment_detail"
	ReturnListParamsReturnableTypePaperItem             ReturnListParamsReturnableType = "paper_item"
	ReturnListParamsReturnableTypePaymentOrder          ReturnListParamsReturnableType = "payment_order"
	ReturnListParamsReturnableTypeReturn                ReturnListParamsReturnableType = "return"
	ReturnListParamsReturnableTypeReversal              ReturnListParamsReturnableType = "reversal"
)
