// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// TransactionService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewTransactionService]
// method instead.
type TransactionService struct {
	Options   []option.RequestOption
	LineItems *TransactionLineItemService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.LineItems = NewTransactionLineItemService(opts...)
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
	ID string `json:"id,required" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// The date on which the transaction occurred.
	AsOfDate time.Time `json:"as_of_date,required,nullable" format:"date"`
	// The time on which the transaction occurred. Depending on the granularity of the
	// timestamp information received from the bank, it may be `null`.
	AsOfTime  string    `json:"as_of_time,required,nullable" format:"time"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Currency that this transaction is denominated in.
	Currency shared.Currency `json:"currency,required,nullable"`
	// This field contains additional information that the bank provided about the
	// transaction. This is structured data. Some of the data in here might overlap
	// with what is in the `vendor_description`. For example, the OBI could be a part
	// of the vendor description, and it would also be included in here. The attributes
	// that are passed through the details field will vary based on your banking
	// partner. Currently, the following keys may be in the details object:
	// `originator_name`, `originator_to_beneficiary_information`.
	Details map[string]string `json:"details,required"`
	// Either `credit` or `debit`.
	Direction   string    `json:"direction,required"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The ID of the relevant Internal Account.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// This field will be `true` if the transaction has posted to the account.
	Posted bool `json:"posted,required"`
	// This field will be `true` if a transaction is reconciled by the Modern Treasury
	// system. This means that it has transaction line items that sum up to the
	// transaction's amount.
	Reconciled bool `json:"reconciled,required"`
	// The type of the transaction. Can be one of `ach`, `wire`, `check`, `rtp`,
	// `book`, or `sen`.
	Type      TransactionType `json:"type,required"`
	UpdatedAt time.Time       `json:"updated_at,required" format:"date-time"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode string `json:"vendor_code,required,nullable"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, or `us_bank`.
	VendorCodeType TransactionVendorCodeType `json:"vendor_code_type,required,nullable"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription string `json:"vendor_description,required,nullable"`
	// An identifier given to this transaction by the bank.
	VendorID string `json:"vendor_id,required,nullable"`
	JSON     transactionJSON
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	AsOfDate          apijson.Field
	AsOfTime          apijson.Field
	CreatedAt         apijson.Field
	Currency          apijson.Field
	Details           apijson.Field
	Direction         apijson.Field
	DiscardedAt       apijson.Field
	InternalAccountID apijson.Field
	LiveMode          apijson.Field
	Metadata          apijson.Field
	Object            apijson.Field
	Posted            apijson.Field
	Reconciled        apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	VendorCode        apijson.Field
	VendorCodeType    apijson.Field
	VendorCustomerID  apijson.Field
	VendorDescription apijson.Field
	VendorID          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the transaction. Can be one of `ach`, `wire`, `check`, `rtp`,
// `book`, or `sen`.
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
	TransactionTypeNics        TransactionType = "nics"
	TransactionTypeProvxchange TransactionType = "provxchange"
	TransactionTypeRtp         TransactionType = "rtp"
	TransactionTypeSeBankgirot TransactionType = "se_bankgirot"
	TransactionTypeSen         TransactionType = "sen"
	TransactionTypeSepa        TransactionType = "sepa"
	TransactionTypeSic         TransactionType = "sic"
	TransactionTypeSignet      TransactionType = "signet"
	TransactionTypeWire        TransactionType = "wire"
	TransactionTypeZengin      TransactionType = "zengin"
)

// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
// `swift`, or `us_bank`.
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

type TransactionUpdateParams struct {
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Filters transactions with an `as_of_date` starting on or before the specified
	// date (YYYY-MM-DD).
	AsOfDateEnd param.Field[time.Time] `query:"as_of_date_end" format:"date"`
	// Filters transactions with an `as_of_date` starting on or after the specified
	// date (YYYY-MM-DD).
	AsOfDateStart  param.Field[time.Time] `query:"as_of_date_start" format:"date"`
	CounterpartyID param.Field[string]    `query:"counterparty_id" format:"uuid"`
	// Filters for transactions including the queried string in the description.
	Description param.Field[string] `query:"description"`
	Direction   param.Field[string] `query:"direction"`
	// Specify `internal_account_id` if you wish to see transactions to/from a specific
	// account.
	InternalAccountID param.Field[string] `query:"internal_account_id" format:"uuid"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata    param.Field[map[string]string] `query:"metadata"`
	PaymentType param.Field[string]            `query:"payment_type"`
	PerPage     param.Field[int64]             `query:"per_page"`
	// Either `true` or `false`.
	Posted           param.Field[bool]   `query:"posted"`
	TransactableType param.Field[string] `query:"transactable_type"`
	// Filters for transactions including the queried vendor id (an identifier given to
	// transactions by the bank).
	VendorID         param.Field[string] `query:"vendor_id"`
	VirtualAccountID param.Field[string] `query:"virtual_account_id" format:"uuid"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
