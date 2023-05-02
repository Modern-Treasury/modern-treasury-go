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

type TransactionService struct {
	Options []option.RequestOption
}

func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	return
}

// Get details on a single transaction.
func (r *TransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a single transaction.
func (r *TransactionService) Update(ctx context.Context, id string, body TransactionUpdateParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all transactions.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *shared.Page[Transaction], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/transactions"
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

// Get a list of all transactions.
func (r *TransactionService) ListAutoPaging(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) *shared.PageAutoPager[Transaction] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Transaction struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// Currency that this transaction is denominated in.
	Currency shared.Currency `json:"currency,required,nullable"`
	// Either `credit` or `debit`.
	Direction string `json:"direction,required"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription string `json:"vendor_description,required,nullable"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode string `json:"vendor_code,required,nullable"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, or `us_bank`.
	VendorCodeType TransactionVendorCodeType `json:"vendor_code_type,required,nullable"`
	// An identifier given to this transaction by the bank.
	VendorID string `json:"vendor_id,required,nullable"`
	// The date on which the transaction occurred.
	AsOfDate time.Time `json:"as_of_date,required,nullable" format:"date"`
	// The time on which the transaction occurred. Depending on the granularity of the
	// timestamp information received from the bank, it may be `null`.
	AsOfTime string `json:"as_of_time,required,nullable" format:"time"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// This field will be `true` if the transaction has posted to the account.
	Posted bool `json:"posted,required"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable"`
	// This field will be `true` if a transaction is reconciled by the Modern Treasury
	// system. This means that it has transaction line items that sum up to the
	// transaction's amount.
	Reconciled bool `json:"reconciled,required"`
	// This field contains additional information that the bank provided about the
	// transaction. This is structured data. Some of the data in here might overlap
	// with what is in the `vendor_description`. For example, the OBI could be a part
	// of the vendor description, and it would also be included in here. The attributes
	// that are passed through the details field will vary based on your banking
	// partner. Currently, the following keys may be in the details object:
	// `originator_name`, `originator_to_beneficiary_information`.
	Details map[string]string `json:"details,required"`
	// The type of the transaction. Can be one of `ach`, `wire`, `check`, `rtp`,
	// `book`, or `sen`.
	Type TransactionType `json:"type,required"`
	JSON TransactionJSON
}

type TransactionJSON struct {
	ID                apijson.Metadata
	Object            apijson.Metadata
	LiveMode          apijson.Metadata
	CreatedAt         apijson.Metadata
	UpdatedAt         apijson.Metadata
	DiscardedAt       apijson.Metadata
	Amount            apijson.Metadata
	Currency          apijson.Metadata
	Direction         apijson.Metadata
	VendorDescription apijson.Metadata
	VendorCode        apijson.Metadata
	VendorCodeType    apijson.Metadata
	VendorID          apijson.Metadata
	AsOfDate          apijson.Metadata
	AsOfTime          apijson.Metadata
	InternalAccountID apijson.Metadata
	Metadata          apijson.Metadata
	Posted            apijson.Metadata
	VendorCustomerID  apijson.Metadata
	Reconciled        apijson.Metadata
	Details           apijson.Metadata
	Type              apijson.Metadata
	raw               string
	Extras            map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Transaction using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionVendorCodeType string

const (
	TransactionVendorCodeTypeBai2          TransactionVendorCodeType = "bai2"
	TransactionVendorCodeTypeBankprov      TransactionVendorCodeType = "bankprov"
	TransactionVendorCodeTypeBnkDev        TransactionVendorCodeType = "bnk_dev"
	TransactionVendorCodeTypeCleartouch    TransactionVendorCodeType = "cleartouch"
	TransactionVendorCodeTypeColumn        TransactionVendorCodeType = "column"
	TransactionVendorCodeTypeCrossRiver    TransactionVendorCodeType = "cross_river"
	TransactionVendorCodeTypeCurrencycloud TransactionVendorCodeType = "currencycloud"
	TransactionVendorCodeTypeDcBank        TransactionVendorCodeType = "dc_bank"
	TransactionVendorCodeTypeDwolla        TransactionVendorCodeType = "dwolla"
	TransactionVendorCodeTypeEvolve        TransactionVendorCodeType = "evolve"
	TransactionVendorCodeTypeGoldmanSachs  TransactionVendorCodeType = "goldman_sachs"
	TransactionVendorCodeTypeIso20022      TransactionVendorCodeType = "iso20022"
	TransactionVendorCodeTypeJpmc          TransactionVendorCodeType = "jpmc"
	TransactionVendorCodeTypeMx            TransactionVendorCodeType = "mx"
	TransactionVendorCodeTypePlaid         TransactionVendorCodeType = "plaid"
	TransactionVendorCodeTypeRspecVendor   TransactionVendorCodeType = "rspec_vendor"
	TransactionVendorCodeTypeSignet        TransactionVendorCodeType = "signet"
	TransactionVendorCodeTypeSilvergate    TransactionVendorCodeType = "silvergate"
	TransactionVendorCodeTypeSwift         TransactionVendorCodeType = "swift"
	TransactionVendorCodeTypeUsBank        TransactionVendorCodeType = "us_bank"
)

type TransactionType string

const (
	TransactionTypeACH         TransactionType = "ach"
	TransactionTypeAuBecs      TransactionType = "au_becs"
	TransactionTypeBacs        TransactionType = "bacs"
	TransactionTypeBook        TransactionType = "book"
	TransactionTypeCard        TransactionType = "card"
	TransactionTypeCheck       TransactionType = "check"
	TransactionTypeCrossBorder TransactionType = "cross_border"
	TransactionTypeEft         TransactionType = "eft"
	TransactionTypeInterac     TransactionType = "interac"
	TransactionTypeMasav       TransactionType = "masav"
	TransactionTypeNeft        TransactionType = "neft"
	TransactionTypeProvxchange TransactionType = "provxchange"
	TransactionTypeRtp         TransactionType = "rtp"
	TransactionTypeSen         TransactionType = "sen"
	TransactionTypeSepa        TransactionType = "sepa"
	TransactionTypeSignet      TransactionType = "signet"
	TransactionTypeWire        TransactionType = "wire"
)

type TransactionUpdateParams struct {
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes TransactionUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// Specify `internal_account_id` if you wish to see transactions to/from a specific
	// account.
	InternalAccountID field.Field[string] `query:"internal_account_id" format:"uuid"`
	VirtualAccountID  field.Field[string] `query:"virtual_account_id" format:"uuid"`
	// Either `true` or `false`.
	Posted field.Field[bool] `query:"posted"`
	// Filters transactions with an `as_of_date` starting on or after the specified
	// date (YYYY-MM-DD).
	AsOfDateStart field.Field[time.Time] `query:"as_of_date_start" format:"date"`
	// Filters transactions with an `as_of_date` starting on or before the specified
	// date (YYYY-MM-DD).
	AsOfDateEnd      field.Field[time.Time] `query:"as_of_date_end" format:"date"`
	Direction        field.Field[string]    `query:"direction"`
	CounterpartyID   field.Field[string]    `query:"counterparty_id" format:"uuid"`
	PaymentType      field.Field[string]    `query:"payment_type"`
	TransactableType field.Field[string]    `query:"transactable_type"`
	// Filters for transactions including the queried string in the description.
	Description field.Field[string] `query:"description"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes TransactionListParams into a url.Values of the query
// parameters associated with this value
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
