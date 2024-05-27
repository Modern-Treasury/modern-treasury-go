// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// ForeignExchangeQuoteService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewForeignExchangeQuoteService] method instead.
type ForeignExchangeQuoteService struct {
	Options []option.RequestOption
}

// NewForeignExchangeQuoteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewForeignExchangeQuoteService(opts ...option.RequestOption) (r *ForeignExchangeQuoteService) {
	r = &ForeignExchangeQuoteService{}
	r.Options = opts
	return
}

// create foreign_exchange_quote
func (r *ForeignExchangeQuoteService) New(ctx context.Context, body ForeignExchangeQuoteNewParams, opts ...option.RequestOption) (res *ForeignExchangeQuote, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/foreign_exchange_quotes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get foreign_exchange_quote
func (r *ForeignExchangeQuoteService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ForeignExchangeQuote, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/foreign_exchange_quotes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// list foreign_exchange_quotes
func (r *ForeignExchangeQuoteService) List(ctx context.Context, query ForeignExchangeQuoteListParams, opts ...option.RequestOption) (res *pagination.Page[ForeignExchangeQuote], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/foreign_exchange_quotes"
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

// list foreign_exchange_quotes
func (r *ForeignExchangeQuoteService) ListAutoPaging(ctx context.Context, query ForeignExchangeQuoteListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ForeignExchangeQuote] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type ForeignExchangeQuote struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The timestamp until when the quoted rate is valid.
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// The timestamp until which the quote must be booked by.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Either `fixed_to_variable` if the `base_amount` was specified, or
	// `variable_to_fixed` if the `target_amount` was specified when requesting the
	// quote.
	ForeignExchangeIndicator string `json:"foreign_exchange_indicator,required"`
	// The serialized rate information represented by this quote.
	ForeignExchangeRate ForeignExchangeQuoteForeignExchangeRate `json:"foreign_exchange_rate,required"`
	// The ID for the `InternalAccount` this quote is associated with.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata  map[string]string `json:"metadata,required"`
	Object    string            `json:"object,required"`
	UpdatedAt time.Time         `json:"updated_at,required" format:"date-time"`
	// This vendor assigned ID for this quote.
	VendorID string                   `json:"vendor_id"`
	JSON     foreignExchangeQuoteJSON `json:"-"`
}

// foreignExchangeQuoteJSON contains the JSON metadata for the struct
// [ForeignExchangeQuote]
type foreignExchangeQuoteJSON struct {
	ID                       apijson.Field
	CreatedAt                apijson.Field
	EffectiveAt              apijson.Field
	ExpiresAt                apijson.Field
	ForeignExchangeIndicator apijson.Field
	ForeignExchangeRate      apijson.Field
	InternalAccountID        apijson.Field
	LiveMode                 apijson.Field
	Metadata                 apijson.Field
	Object                   apijson.Field
	UpdatedAt                apijson.Field
	VendorID                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ForeignExchangeQuote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r foreignExchangeQuoteJSON) RawJSON() string {
	return r.raw
}

// The serialized rate information represented by this quote.
type ForeignExchangeQuoteForeignExchangeRate struct {
	// Amount in the lowest denomination of the `base_currency` to convert, often
	// called the "sell" amount.
	BaseAmount int64 `json:"base_amount,required"`
	// Currency to convert, often called the "sell" currency.
	BaseCurrency shared.Currency `json:"base_currency,required,nullable"`
	// The exponent component of the rate. The decimal is calculated as `value` / (10 ^
	// `exponent`).
	Exponent int64 `json:"exponent,required"`
	// A string representation of the rate.
	RateString string `json:"rate_string,required"`
	// Amount in the lowest denomination of the `target_currency`, often called the
	// "buy" amount.
	TargetAmount int64 `json:"target_amount,required"`
	// Currency to convert the `base_currency` to, often called the "buy" currency.
	TargetCurrency shared.Currency `json:"target_currency,required,nullable"`
	// The whole number component of the rate. The decimal is calculated as `value` /
	// (10 ^ `exponent`).
	Value int64                                       `json:"value,required"`
	JSON  foreignExchangeQuoteForeignExchangeRateJSON `json:"-"`
}

// foreignExchangeQuoteForeignExchangeRateJSON contains the JSON metadata for the
// struct [ForeignExchangeQuoteForeignExchangeRate]
type foreignExchangeQuoteForeignExchangeRateJSON struct {
	BaseAmount     apijson.Field
	BaseCurrency   apijson.Field
	Exponent       apijson.Field
	RateString     apijson.Field
	TargetAmount   apijson.Field
	TargetCurrency apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ForeignExchangeQuoteForeignExchangeRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r foreignExchangeQuoteForeignExchangeRateJSON) RawJSON() string {
	return r.raw
}

type ForeignExchangeQuoteNewParams struct {
	// The ID for the `InternalAccount` this quote is associated with.
	InternalAccountID param.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// Currency to convert the `base_currency` to, often called the "buy" currency.
	TargetCurrency param.Field[shared.Currency] `json:"target_currency,required"`
	// Amount in the lowest denomination of the `base_currency` to convert, often
	// called the "sell" amount.
	BaseAmount param.Field[int64] `json:"base_amount"`
	// Currency to convert, often called the "sell" currency.
	BaseCurrency param.Field[shared.Currency] `json:"base_currency"`
	// The timestamp until when the quoted rate is valid.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// Amount in the lowest denomination of the `target_currency`, often called the
	// "buy" amount.
	TargetAmount param.Field[int64] `json:"target_amount"`
}

func (r ForeignExchangeQuoteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ForeignExchangeQuoteListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Currency to convert, often called the "sell" currency.
	BaseCurrency param.Field[string] `query:"base_currency"`
	// An inclusive upper bound for searching effective_at
	EffectiveAtEnd param.Field[time.Time] `query:"effective_at_end" format:"date"`
	// An inclusive lower bound for searching effective_at
	EffectiveAtStart param.Field[time.Time] `query:"effective_at_start" format:"date"`
	// The timestamp until which the quote must be booked by.
	ExpiresAt param.Field[time.Time] `query:"expires_at" format:"date-time"`
	// The ID for the `InternalAccount` this quote is associated with.
	InternalAccountID param.Field[string] `query:"internal_account_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	PerPage  param.Field[int64]             `query:"per_page"`
	// Currency to convert the `base_currency` to, often called the "buy" currency.
	TargetCurrency param.Field[string] `query:"target_currency"`
}

// URLQuery serializes [ForeignExchangeQuoteListParams]'s query parameters as
// `url.Values`.
func (r ForeignExchangeQuoteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
