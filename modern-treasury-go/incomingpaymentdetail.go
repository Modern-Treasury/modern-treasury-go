package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type IncomingPaymentDetailService struct {
	Options []option.RequestOption
}

func NewIncomingPaymentDetailService(opts ...option.RequestOption) (r *IncomingPaymentDetailService) {
	r = &IncomingPaymentDetailService{}
	r.Options = opts
	return
}

// Get an existing Incoming Payment Detail.
func (r *IncomingPaymentDetailService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IncomingPaymentDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/incoming_payment_details/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing Incoming Payment Detail.
func (r *IncomingPaymentDetailService) Update(ctx context.Context, id string, body IncomingPaymentDetailUpdateParams, opts ...option.RequestOption) (res *IncomingPaymentDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/incoming_payment_details/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of Incoming Payment Details.
func (r *IncomingPaymentDetailService) List(ctx context.Context, query IncomingPaymentDetailListParams, opts ...option.RequestOption) (res *shared.Page[IncomingPaymentDetail], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/incoming_payment_details"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of Incoming Payment Details.
func (r *IncomingPaymentDetailService) ListAutoPaging(ctx context.Context, query IncomingPaymentDetailListParams, opts ...option.RequestOption) *shared.PageAutoPager[IncomingPaymentDetail] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Simulate Incoming Payment Detail
func (r *IncomingPaymentDetailService) NewAsync(ctx context.Context, body IncomingPaymentDetailNewAsyncParams, opts ...option.RequestOption) (res *shared.AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/simulations/incoming_payment_details/create_async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IncomingPaymentDetail struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the Internal Account for the incoming payment detail. This is always
	// present.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// If the incoming payment detail is in a virtual account, the ID of the Virtual
	// Account.
	VirtualAccountID string `json:"virtual_account_id,required,nullable" format:"uuid"`
	// If the incoming payment detail is in a virtual account, the serialized virtual
	// account object.
	VirtualAccount VirtualAccount `json:"virtual_account,required,nullable"`
	// The ID of the reconciled Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The ID of the reconciled Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// One of: `ach`, `book`, `check`, `eft`, `interac`, `rtp`, `sepa`, `signet`, or
	// `wire`.
	Type IncomingPaymentDetailType `json:"type,required"`
	// The raw data from the payment pre-notification file that we get from the bank.
	Data map[string]interface{} `json:"data,required"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// The currency of the incoming payment detail.
	Currency shared.Currency `json:"currency,required,nullable"`
	// One of `credit` or `debit`.
	Direction IncomingPaymentDetailDirection `json:"direction,required"`
	// The current status of the incoming payment order. One of `pending`, `completed`,
	// or `returned`.
	Status IncomingPaymentDetailStatus `json:"status,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The date on which the corresponding transaction will occur.
	AsOfDate time.Time `json:"as_of_date,required" format:"date"`
	// The identifier of the vendor bank.
	VendorID string `json:"vendor_id,required,nullable" format:"uuid"`
	JSON     IncomingPaymentDetailJSON
}

type IncomingPaymentDetailJSON struct {
	ID                    apijson.Metadata
	Object                apijson.Metadata
	LiveMode              apijson.Metadata
	CreatedAt             apijson.Metadata
	UpdatedAt             apijson.Metadata
	InternalAccountID     apijson.Metadata
	VirtualAccountID      apijson.Metadata
	VirtualAccount        apijson.Metadata
	TransactionLineItemID apijson.Metadata
	TransactionID         apijson.Metadata
	Type                  apijson.Metadata
	Data                  apijson.Metadata
	Amount                apijson.Metadata
	Currency              apijson.Metadata
	Direction             apijson.Metadata
	Status                apijson.Metadata
	Metadata              apijson.Metadata
	AsOfDate              apijson.Metadata
	VendorID              apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into IncomingPaymentDetail using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *IncomingPaymentDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IncomingPaymentDetailType string

const (
	IncomingPaymentDetailTypeACH     IncomingPaymentDetailType = "ach"
	IncomingPaymentDetailTypeBook    IncomingPaymentDetailType = "book"
	IncomingPaymentDetailTypeCheck   IncomingPaymentDetailType = "check"
	IncomingPaymentDetailTypeEft     IncomingPaymentDetailType = "eft"
	IncomingPaymentDetailTypeInterac IncomingPaymentDetailType = "interac"
	IncomingPaymentDetailTypeRtp     IncomingPaymentDetailType = "rtp"
	IncomingPaymentDetailTypeSepa    IncomingPaymentDetailType = "sepa"
	IncomingPaymentDetailTypeSignet  IncomingPaymentDetailType = "signet"
	IncomingPaymentDetailTypeWire    IncomingPaymentDetailType = "wire"
)

type IncomingPaymentDetailDirection string

const (
	IncomingPaymentDetailDirectionCredit IncomingPaymentDetailDirection = "credit"
	IncomingPaymentDetailDirectionDebit  IncomingPaymentDetailDirection = "debit"
)

type IncomingPaymentDetailStatus string

const (
	IncomingPaymentDetailStatusCompleted IncomingPaymentDetailStatus = "completed"
	IncomingPaymentDetailStatusPending   IncomingPaymentDetailStatus = "pending"
	IncomingPaymentDetailStatusReturned  IncomingPaymentDetailStatus = "returned"
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
	return apijson.MarshalRoot(r)
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
	return apiquery.Marshal(r)
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
	Currency field.Field[shared.Currency] `json:"currency,nullable"`
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
	return apijson.MarshalRoot(r)
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
