// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// TransactionService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
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

// create transaction
func (r *TransactionService) New(ctx context.Context, body TransactionNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single transaction.
func (r *TransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a single transaction.
func (r *TransactionService) Update(ctx context.Context, id string, body TransactionUpdateParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all transactions.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *pagination.Page[Transaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
func (r *TransactionService) ListAutoPaging(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Transaction] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete transaction
func (r *TransactionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
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
	AsOfTime string `json:"as_of_time,required,nullable" format:"time"`
	// The timezone in which the `as_of_time` is represented. Can be `null` if the bank
	// does not provide timezone info.
	AsOfTimezone string    `json:"as_of_timezone,required,nullable"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	// Currency that this transaction is denominated in.
	Currency shared.Currency `json:"currency,required"`
	// An object containing key-value pairs, each with a custom identifier as the key
	// and a string value.
	CustomIdentifiers map[string]string `json:"custom_identifiers,required"`
	// Either `credit` or `debit`.
	Direction   string    `json:"direction,required"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Associated serialized foreign exchange rate information.
	ForeignExchangeRate shared.ForeignExchangeRate `json:"foreign_exchange_rate,required,nullable"`
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
	// The type of the transaction. Examples could be
	// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
	Type      TransactionType `json:"type,required"`
	UpdatedAt time.Time       `json:"updated_at,required" format:"date-time"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode string `json:"vendor_code,required,nullable"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, `us_bank`, or others.
	VendorCodeType TransactionVendorCodeType `json:"vendor_code_type,required,nullable"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID string `json:"vendor_customer_id,required,nullable"`
	// An identifier given to this transaction by the bank.
	VendorID string `json:"vendor_id,required,nullable"`
	// This field contains additional information that the bank provided about the
	// transaction. This is structured data. Some of the data in here might overlap
	// with what is in the `vendor_description`. For example, the OBI could be a part
	// of the vendor description, and it would also be included in here. The attributes
	// that are passed through the details field will vary based on your banking
	// partner. Currently, the following keys may be in the details object:
	// `originator_name`, `originator_to_beneficiary_information`.
	Details map[string]string `json:"details"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription string          `json:"vendor_description,nullable"`
	JSON              transactionJSON `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID                  apijson.Field
	Amount              apijson.Field
	AsOfDate            apijson.Field
	AsOfTime            apijson.Field
	AsOfTimezone        apijson.Field
	CreatedAt           apijson.Field
	Currency            apijson.Field
	CustomIdentifiers   apijson.Field
	Direction           apijson.Field
	DiscardedAt         apijson.Field
	ForeignExchangeRate apijson.Field
	InternalAccountID   apijson.Field
	LiveMode            apijson.Field
	Metadata            apijson.Field
	Object              apijson.Field
	Posted              apijson.Field
	Reconciled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	VendorCode          apijson.Field
	VendorCodeType      apijson.Field
	VendorCustomerID    apijson.Field
	VendorID            apijson.Field
	Details             apijson.Field
	VendorDescription   apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

func (r Transaction) implementsBulkResultEntity() {}

// The type of the transaction. Examples could be
// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
type TransactionType string

const (
	TransactionTypeACH         TransactionType = "ach"
	TransactionTypeAuBecs      TransactionType = "au_becs"
	TransactionTypeBacs        TransactionType = "bacs"
	TransactionTypeBase        TransactionType = "base"
	TransactionTypeBook        TransactionType = "book"
	TransactionTypeCard        TransactionType = "card"
	TransactionTypeChats       TransactionType = "chats"
	TransactionTypeCheck       TransactionType = "check"
	TransactionTypeCrossBorder TransactionType = "cross_border"
	TransactionTypeDkNets      TransactionType = "dk_nets"
	TransactionTypeEft         TransactionType = "eft"
	TransactionTypeEthereum    TransactionType = "ethereum"
	TransactionTypeGBFps       TransactionType = "gb_fps"
	TransactionTypeHuIcs       TransactionType = "hu_ics"
	TransactionTypeInterac     TransactionType = "interac"
	TransactionTypeMasav       TransactionType = "masav"
	TransactionTypeMxCcen      TransactionType = "mx_ccen"
	TransactionTypeNeft        TransactionType = "neft"
	TransactionTypeNics        TransactionType = "nics"
	TransactionTypeNzBecs      TransactionType = "nz_becs"
	TransactionTypePlElixir    TransactionType = "pl_elixir"
	TransactionTypePolygon     TransactionType = "polygon"
	TransactionTypeProvxchange TransactionType = "provxchange"
	TransactionTypeRoSent      TransactionType = "ro_sent"
	TransactionTypeRtp         TransactionType = "rtp"
	TransactionTypeSeBankgirot TransactionType = "se_bankgirot"
	TransactionTypeSen         TransactionType = "sen"
	TransactionTypeSepa        TransactionType = "sepa"
	TransactionTypeSgGiro      TransactionType = "sg_giro"
	TransactionTypeSic         TransactionType = "sic"
	TransactionTypeSignet      TransactionType = "signet"
	TransactionTypeSknbi       TransactionType = "sknbi"
	TransactionTypeSolana      TransactionType = "solana"
	TransactionTypeWire        TransactionType = "wire"
	TransactionTypeZengin      TransactionType = "zengin"
	TransactionTypeOther       TransactionType = "other"
)

func (r TransactionType) IsKnown() bool {
	switch r {
	case TransactionTypeACH, TransactionTypeAuBecs, TransactionTypeBacs, TransactionTypeBase, TransactionTypeBook, TransactionTypeCard, TransactionTypeChats, TransactionTypeCheck, TransactionTypeCrossBorder, TransactionTypeDkNets, TransactionTypeEft, TransactionTypeEthereum, TransactionTypeGBFps, TransactionTypeHuIcs, TransactionTypeInterac, TransactionTypeMasav, TransactionTypeMxCcen, TransactionTypeNeft, TransactionTypeNics, TransactionTypeNzBecs, TransactionTypePlElixir, TransactionTypePolygon, TransactionTypeProvxchange, TransactionTypeRoSent, TransactionTypeRtp, TransactionTypeSeBankgirot, TransactionTypeSen, TransactionTypeSepa, TransactionTypeSgGiro, TransactionTypeSic, TransactionTypeSignet, TransactionTypeSknbi, TransactionTypeSolana, TransactionTypeWire, TransactionTypeZengin, TransactionTypeOther:
		return true
	}
	return false
}

// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
// `swift`, `us_bank`, or others.
type TransactionVendorCodeType string

const (
	TransactionVendorCodeTypeBai2            TransactionVendorCodeType = "bai2"
	TransactionVendorCodeTypeBankingCircle   TransactionVendorCodeType = "banking_circle"
	TransactionVendorCodeTypeBankprov        TransactionVendorCodeType = "bankprov"
	TransactionVendorCodeTypeBnkDev          TransactionVendorCodeType = "bnk_dev"
	TransactionVendorCodeTypeBrale           TransactionVendorCodeType = "brale"
	TransactionVendorCodeTypeCleartouch      TransactionVendorCodeType = "cleartouch"
	TransactionVendorCodeTypeColumn          TransactionVendorCodeType = "column"
	TransactionVendorCodeTypeCrossRiver      TransactionVendorCodeType = "cross_river"
	TransactionVendorCodeTypeCurrencycloud   TransactionVendorCodeType = "currencycloud"
	TransactionVendorCodeTypeDcBank          TransactionVendorCodeType = "dc_bank"
	TransactionVendorCodeTypeDwolla          TransactionVendorCodeType = "dwolla"
	TransactionVendorCodeTypeEvolve          TransactionVendorCodeType = "evolve"
	TransactionVendorCodeTypeFakeVendor      TransactionVendorCodeType = "fake_vendor"
	TransactionVendorCodeTypeGoldmanSachs    TransactionVendorCodeType = "goldman_sachs"
	TransactionVendorCodeTypeHifi            TransactionVendorCodeType = "hifi"
	TransactionVendorCodeTypeIso20022        TransactionVendorCodeType = "iso20022"
	TransactionVendorCodeTypeJpmc            TransactionVendorCodeType = "jpmc"
	TransactionVendorCodeTypeMtFof           TransactionVendorCodeType = "mt_fof"
	TransactionVendorCodeTypeMx              TransactionVendorCodeType = "mx"
	TransactionVendorCodeTypePaypal          TransactionVendorCodeType = "paypal"
	TransactionVendorCodeTypePlaid           TransactionVendorCodeType = "plaid"
	TransactionVendorCodeTypePnc             TransactionVendorCodeType = "pnc"
	TransactionVendorCodeTypeSignet          TransactionVendorCodeType = "signet"
	TransactionVendorCodeTypeSilvergate      TransactionVendorCodeType = "silvergate"
	TransactionVendorCodeTypeSwift           TransactionVendorCodeType = "swift"
	TransactionVendorCodeTypeUsBank          TransactionVendorCodeType = "us_bank"
	TransactionVendorCodeTypeUser            TransactionVendorCodeType = "user"
	TransactionVendorCodeTypeWesternAlliance TransactionVendorCodeType = "western_alliance"
)

func (r TransactionVendorCodeType) IsKnown() bool {
	switch r {
	case TransactionVendorCodeTypeBai2, TransactionVendorCodeTypeBankingCircle, TransactionVendorCodeTypeBankprov, TransactionVendorCodeTypeBnkDev, TransactionVendorCodeTypeBrale, TransactionVendorCodeTypeCleartouch, TransactionVendorCodeTypeColumn, TransactionVendorCodeTypeCrossRiver, TransactionVendorCodeTypeCurrencycloud, TransactionVendorCodeTypeDcBank, TransactionVendorCodeTypeDwolla, TransactionVendorCodeTypeEvolve, TransactionVendorCodeTypeFakeVendor, TransactionVendorCodeTypeGoldmanSachs, TransactionVendorCodeTypeHifi, TransactionVendorCodeTypeIso20022, TransactionVendorCodeTypeJpmc, TransactionVendorCodeTypeMtFof, TransactionVendorCodeTypeMx, TransactionVendorCodeTypePaypal, TransactionVendorCodeTypePlaid, TransactionVendorCodeTypePnc, TransactionVendorCodeTypeSignet, TransactionVendorCodeTypeSilvergate, TransactionVendorCodeTypeSwift, TransactionVendorCodeTypeUsBank, TransactionVendorCodeTypeUser, TransactionVendorCodeTypeWesternAlliance:
		return true
	}
	return false
}

type TransactionNewParams struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// The date on which the transaction occurred.
	AsOfDate param.Field[time.Time] `json:"as_of_date,required" format:"date"`
	// Either `credit` or `debit`.
	Direction param.Field[string] `json:"direction,required"`
	// The ID of the relevant Internal Account.
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// When applicable, the bank-given code that determines the transaction's category.
	// For most banks this is the BAI2/BTRS transaction code.
	VendorCode param.Field[string] `json:"vendor_code,required"`
	// The type of `vendor_code` being reported. Can be one of `bai2`, `bankprov`,
	// `bnk_dev`, `cleartouch`, `currencycloud`, `cross_river`, `dc_bank`, `dwolla`,
	// `evolve`, `goldman_sachs`, `iso20022`, `jpmc`, `mx`, `signet`, `silvergate`,
	// `swift`, `us_bank`, or others.
	VendorCodeType param.Field[string] `json:"vendor_code_type,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// This field will be `true` if the transaction has posted to the account.
	Posted param.Field[bool] `json:"posted"`
	// The type of the transaction. Examples could be
	// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
	Type param.Field[TransactionNewParamsType] `json:"type"`
	// An identifier given to this transaction by the bank, often `null`.
	VendorCustomerID param.Field[string] `json:"vendor_customer_id"`
	// The transaction detail text that often appears in on your bank statement and in
	// your banking portal.
	VendorDescription param.Field[string] `json:"vendor_description"`
}

func (r TransactionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the transaction. Examples could be
// `card, `ach`, `wire`, `check`, `rtp`, `book`, or `sen`.
type TransactionNewParamsType string

const (
	TransactionNewParamsTypeACH         TransactionNewParamsType = "ach"
	TransactionNewParamsTypeAuBecs      TransactionNewParamsType = "au_becs"
	TransactionNewParamsTypeBacs        TransactionNewParamsType = "bacs"
	TransactionNewParamsTypeBase        TransactionNewParamsType = "base"
	TransactionNewParamsTypeBook        TransactionNewParamsType = "book"
	TransactionNewParamsTypeCard        TransactionNewParamsType = "card"
	TransactionNewParamsTypeChats       TransactionNewParamsType = "chats"
	TransactionNewParamsTypeCheck       TransactionNewParamsType = "check"
	TransactionNewParamsTypeCrossBorder TransactionNewParamsType = "cross_border"
	TransactionNewParamsTypeDkNets      TransactionNewParamsType = "dk_nets"
	TransactionNewParamsTypeEft         TransactionNewParamsType = "eft"
	TransactionNewParamsTypeEthereum    TransactionNewParamsType = "ethereum"
	TransactionNewParamsTypeGBFps       TransactionNewParamsType = "gb_fps"
	TransactionNewParamsTypeHuIcs       TransactionNewParamsType = "hu_ics"
	TransactionNewParamsTypeInterac     TransactionNewParamsType = "interac"
	TransactionNewParamsTypeMasav       TransactionNewParamsType = "masav"
	TransactionNewParamsTypeMxCcen      TransactionNewParamsType = "mx_ccen"
	TransactionNewParamsTypeNeft        TransactionNewParamsType = "neft"
	TransactionNewParamsTypeNics        TransactionNewParamsType = "nics"
	TransactionNewParamsTypeNzBecs      TransactionNewParamsType = "nz_becs"
	TransactionNewParamsTypePlElixir    TransactionNewParamsType = "pl_elixir"
	TransactionNewParamsTypePolygon     TransactionNewParamsType = "polygon"
	TransactionNewParamsTypeProvxchange TransactionNewParamsType = "provxchange"
	TransactionNewParamsTypeRoSent      TransactionNewParamsType = "ro_sent"
	TransactionNewParamsTypeRtp         TransactionNewParamsType = "rtp"
	TransactionNewParamsTypeSeBankgirot TransactionNewParamsType = "se_bankgirot"
	TransactionNewParamsTypeSen         TransactionNewParamsType = "sen"
	TransactionNewParamsTypeSepa        TransactionNewParamsType = "sepa"
	TransactionNewParamsTypeSgGiro      TransactionNewParamsType = "sg_giro"
	TransactionNewParamsTypeSic         TransactionNewParamsType = "sic"
	TransactionNewParamsTypeSignet      TransactionNewParamsType = "signet"
	TransactionNewParamsTypeSknbi       TransactionNewParamsType = "sknbi"
	TransactionNewParamsTypeSolana      TransactionNewParamsType = "solana"
	TransactionNewParamsTypeWire        TransactionNewParamsType = "wire"
	TransactionNewParamsTypeZengin      TransactionNewParamsType = "zengin"
	TransactionNewParamsTypeOther       TransactionNewParamsType = "other"
)

func (r TransactionNewParamsType) IsKnown() bool {
	switch r {
	case TransactionNewParamsTypeACH, TransactionNewParamsTypeAuBecs, TransactionNewParamsTypeBacs, TransactionNewParamsTypeBase, TransactionNewParamsTypeBook, TransactionNewParamsTypeCard, TransactionNewParamsTypeChats, TransactionNewParamsTypeCheck, TransactionNewParamsTypeCrossBorder, TransactionNewParamsTypeDkNets, TransactionNewParamsTypeEft, TransactionNewParamsTypeEthereum, TransactionNewParamsTypeGBFps, TransactionNewParamsTypeHuIcs, TransactionNewParamsTypeInterac, TransactionNewParamsTypeMasav, TransactionNewParamsTypeMxCcen, TransactionNewParamsTypeNeft, TransactionNewParamsTypeNics, TransactionNewParamsTypeNzBecs, TransactionNewParamsTypePlElixir, TransactionNewParamsTypePolygon, TransactionNewParamsTypeProvxchange, TransactionNewParamsTypeRoSent, TransactionNewParamsTypeRtp, TransactionNewParamsTypeSeBankgirot, TransactionNewParamsTypeSen, TransactionNewParamsTypeSepa, TransactionNewParamsTypeSgGiro, TransactionNewParamsTypeSic, TransactionNewParamsTypeSignet, TransactionNewParamsTypeSknbi, TransactionNewParamsTypeSolana, TransactionNewParamsTypeWire, TransactionNewParamsTypeZengin, TransactionNewParamsTypeOther:
		return true
	}
	return false
}

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
	CounterpartyID param.Field[string]    `query:"counterparty_id"`
	// Filters for transactions including the queried string in the description.
	Description param.Field[string] `query:"description"`
	Direction   param.Field[string] `query:"direction"`
	// Specify `internal_account_id` if you wish to see transactions to/from a specific
	// account.
	InternalAccountID param.Field[string] `query:"internal_account_id"`
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
	VirtualAccountID param.Field[string] `query:"virtual_account_id"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
