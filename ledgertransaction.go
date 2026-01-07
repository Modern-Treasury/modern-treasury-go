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

// LedgerTransactionService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerTransactionService] method instead.
type LedgerTransactionService struct {
	Options  []option.RequestOption
	Versions *LedgerTransactionVersionService
}

// NewLedgerTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLedgerTransactionService(opts ...option.RequestOption) (r *LedgerTransactionService) {
	r = &LedgerTransactionService{}
	r.Options = opts
	r.Versions = NewLedgerTransactionVersionService(opts...)
	return
}

// Create a ledger transaction.
func (r *LedgerTransactionService) New(ctx context.Context, body LedgerTransactionNewParams, opts ...option.RequestOption) (res *LedgerTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/ledger_transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger transaction.
func (r *LedgerTransactionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the details of a ledger transaction.
func (r *LedgerTransactionService) Update(ctx context.Context, id string, body LedgerTransactionUpdateParams, opts ...option.RequestOption) (res *LedgerTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_transactions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of ledger transactions.
func (r *LedgerTransactionService) List(ctx context.Context, query LedgerTransactionListParams, opts ...option.RequestOption) (res *pagination.Page[LedgerTransaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_transactions"
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

// Get a list of ledger transactions.
func (r *LedgerTransactionService) ListAutoPaging(ctx context.Context, query LedgerTransactionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LedgerTransaction] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Create a ledger transaction that partially posts another ledger transaction.
func (r *LedgerTransactionService) NewPartialPost(ctx context.Context, id string, body LedgerTransactionNewPartialPostParams, opts ...option.RequestOption) (res *LedgerTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_transactions/%s/partial_post", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a ledger transaction reversal.
func (r *LedgerTransactionService) NewReversal(ctx context.Context, id string, body LedgerTransactionNewReversalParams, opts ...option.RequestOption) (res *LedgerTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/ledger_transactions/%s/reversal", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LedgerTransaction struct {
	ID string `json:"id,required" format:"uuid"`
	// System-set reason why the ledger transaction was archived; currently only
	// 'balance_lock_failure' for transactions that violated balance constraints. Only
	// populated when archive_on_balance_lock_failure is true and a balance lock
	// violation occurs, otherwise null.
	ArchivedReason string    `json:"archived_reason,required,nullable"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt time.Time `json:"effective_at,required" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID string `json:"external_id,required,nullable"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerEntry `json:"ledger_entries,required"`
	// The ID of the ledger this ledger transaction belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType LedgerTransactionLedgerableType `json:"ledgerable_type,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The ID of the ledger transaction that this ledger transaction partially posts.
	PartiallyPostsLedgerTransactionID string `json:"partially_posts_ledger_transaction_id,required,nullable"`
	// The time on which the ledger transaction posted. This is null if the ledger
	// transaction is pending.
	PostedAt time.Time `json:"posted_at,required,nullable" format:"date-time"`
	// The ID of the ledger transaction that reversed this ledger transaction.
	ReversedByLedgerTransactionID string `json:"reversed_by_ledger_transaction_id,required,nullable"`
	// The ID of the original ledger transaction that this ledger transaction reverses.
	ReversesLedgerTransactionID string `json:"reverses_ledger_transaction_id,required,nullable"`
	// To post a ledger transaction at creation, use `posted`.
	Status    LedgerTransactionStatus `json:"status,required"`
	UpdatedAt time.Time               `json:"updated_at,required" format:"date-time"`
	JSON      ledgerTransactionJSON   `json:"-"`
}

// ledgerTransactionJSON contains the JSON metadata for the struct
// [LedgerTransaction]
type ledgerTransactionJSON struct {
	ID                                apijson.Field
	ArchivedReason                    apijson.Field
	CreatedAt                         apijson.Field
	Description                       apijson.Field
	EffectiveAt                       apijson.Field
	EffectiveDate                     apijson.Field
	ExternalID                        apijson.Field
	LedgerEntries                     apijson.Field
	LedgerID                          apijson.Field
	LedgerableID                      apijson.Field
	LedgerableType                    apijson.Field
	LiveMode                          apijson.Field
	Metadata                          apijson.Field
	Object                            apijson.Field
	PartiallyPostsLedgerTransactionID apijson.Field
	PostedAt                          apijson.Field
	ReversedByLedgerTransactionID     apijson.Field
	ReversesLedgerTransactionID       apijson.Field
	Status                            apijson.Field
	UpdatedAt                         apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *LedgerTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerTransactionJSON) RawJSON() string {
	return r.raw
}

func (r LedgerTransaction) implementsBulkResultEntity() {}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type LedgerTransactionLedgerableType string

const (
	LedgerTransactionLedgerableTypeExpectedPayment       LedgerTransactionLedgerableType = "expected_payment"
	LedgerTransactionLedgerableTypeIncomingPaymentDetail LedgerTransactionLedgerableType = "incoming_payment_detail"
	LedgerTransactionLedgerableTypePaymentOrder          LedgerTransactionLedgerableType = "payment_order"
	LedgerTransactionLedgerableTypeReturn                LedgerTransactionLedgerableType = "return"
	LedgerTransactionLedgerableTypeReversal              LedgerTransactionLedgerableType = "reversal"
)

func (r LedgerTransactionLedgerableType) IsKnown() bool {
	switch r {
	case LedgerTransactionLedgerableTypeExpectedPayment, LedgerTransactionLedgerableTypeIncomingPaymentDetail, LedgerTransactionLedgerableTypePaymentOrder, LedgerTransactionLedgerableTypeReturn, LedgerTransactionLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type LedgerTransactionStatus string

const (
	LedgerTransactionStatusArchived LedgerTransactionStatus = "archived"
	LedgerTransactionStatusPending  LedgerTransactionStatus = "pending"
	LedgerTransactionStatusPosted   LedgerTransactionStatus = "posted"
)

func (r LedgerTransactionStatus) IsKnown() bool {
	switch r {
	case LedgerTransactionStatusArchived, LedgerTransactionStatusPending, LedgerTransactionStatusPosted:
		return true
	}
	return false
}

type LedgerTransactionNewParams struct {
	LedgerTransactionCreateRequest shared.LedgerTransactionCreateRequestParam `json:"ledger_transaction_create_request,required"`
}

func (r LedgerTransactionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LedgerTransactionCreateRequest)
}

type LedgerTransactionUpdateParams struct {
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]shared.LedgerEntryCreateRequestParam] `json:"ledger_entries"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType param.Field[LedgerTransactionUpdateParamsLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[LedgerTransactionUpdateParamsStatus] `json:"status"`
}

func (r LedgerTransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
type LedgerTransactionUpdateParamsLedgerableType string

const (
	LedgerTransactionUpdateParamsLedgerableTypeExpectedPayment       LedgerTransactionUpdateParamsLedgerableType = "expected_payment"
	LedgerTransactionUpdateParamsLedgerableTypeIncomingPaymentDetail LedgerTransactionUpdateParamsLedgerableType = "incoming_payment_detail"
	LedgerTransactionUpdateParamsLedgerableTypePaymentOrder          LedgerTransactionUpdateParamsLedgerableType = "payment_order"
	LedgerTransactionUpdateParamsLedgerableTypeReturn                LedgerTransactionUpdateParamsLedgerableType = "return"
	LedgerTransactionUpdateParamsLedgerableTypeReversal              LedgerTransactionUpdateParamsLedgerableType = "reversal"
)

func (r LedgerTransactionUpdateParamsLedgerableType) IsKnown() bool {
	switch r {
	case LedgerTransactionUpdateParamsLedgerableTypeExpectedPayment, LedgerTransactionUpdateParamsLedgerableTypeIncomingPaymentDetail, LedgerTransactionUpdateParamsLedgerableTypePaymentOrder, LedgerTransactionUpdateParamsLedgerableTypeReturn, LedgerTransactionUpdateParamsLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type LedgerTransactionUpdateParamsStatus string

const (
	LedgerTransactionUpdateParamsStatusArchived LedgerTransactionUpdateParamsStatus = "archived"
	LedgerTransactionUpdateParamsStatusPending  LedgerTransactionUpdateParamsStatus = "pending"
	LedgerTransactionUpdateParamsStatusPosted   LedgerTransactionUpdateParamsStatus = "posted"
)

func (r LedgerTransactionUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LedgerTransactionUpdateParamsStatusArchived, LedgerTransactionUpdateParamsStatusPending, LedgerTransactionUpdateParamsStatusPosted:
		return true
	}
	return false
}

type LedgerTransactionListParams struct {
	// If you have specific IDs to retrieve in bulk, you can pass them as query
	// parameters delimited with `id[]=`, for example `?id[]=123&id[]=abc`.
	ID          param.Field[[]string] `query:"id"`
	AfterCursor param.Field[string]   `query:"after_cursor"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by amount.
	Amount param.Field[LedgerTransactionListParamsAmount] `query:"amount"`
	// Use "gt" (>), "gte" (>=), "lt" (<), "lte" (<=), or "eq" (=) to filter by
	// effective at. For example, for all transactions after Jan 1 2000, use
	// effective_at%5Bgt%5D=2000-01-01T00:00:00:00.000Z.
	EffectiveAt param.Field[map[string]time.Time] `query:"effective_at" format:"date-time"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by
	// effective date. For example, for all dates after Jan 1 2000, use
	// effective_date%5Bgt%5D=2000-01-01.
	EffectiveDate             param.Field[map[string]time.Time]                      `query:"effective_date" format:"date-time"`
	ExternalID                param.Field[string]                                    `query:"external_id"`
	LedgerAccountCategoryID   param.Field[string]                                    `query:"ledger_account_category_id"`
	LedgerAccountID           param.Field[string]                                    `query:"ledger_account_id"`
	LedgerAccountSettlementID param.Field[string]                                    `query:"ledger_account_settlement_id"`
	LedgerID                  param.Field[string]                                    `query:"ledger_id"`
	LedgerableID              param.Field[string]                                    `query:"ledgerable_id"`
	LedgerableType            param.Field[LedgerTransactionListParamsLedgerableType] `query:"ledgerable_type"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
	// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
	// by only one field at a time is supported.
	OrderBy                           param.Field[LedgerTransactionListParamsOrderBy] `query:"order_by"`
	PartiallyPostsLedgerTransactionID param.Field[string]                             `query:"partially_posts_ledger_transaction_id"`
	PerPage                           param.Field[int64]                              `query:"per_page"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// posted_at%5Bgt%5D=2000-01-01T12:00:00Z.
	PostedAt                    param.Field[map[string]time.Time]                   `query:"posted_at" format:"date-time"`
	ReversesLedgerTransactionID param.Field[string]                                 `query:"reverses_ledger_transaction_id"`
	Status                      param.Field[LedgerTransactionListParamsStatusUnion] `query:"status"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// posted at timestamp. For example, for all times after Jan 1 2000 12:00 UTC, use
	// updated_at%5Bgt%5D=2000-01-01T12:00:00Z.
	UpdatedAt param.Field[map[string]time.Time] `query:"updated_at" format:"date-time"`
}

// URLQuery serializes [LedgerTransactionListParams]'s query parameters as
// `url.Values`.
func (r LedgerTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by amount.
type LedgerTransactionListParamsAmount struct {
	Eq  param.Field[int64] `query:"eq"`
	Gt  param.Field[int64] `query:"gt"`
	Gte param.Field[int64] `query:"gte"`
	Lt  param.Field[int64] `query:"lt"`
	Lte param.Field[int64] `query:"lte"`
}

// URLQuery serializes [LedgerTransactionListParamsAmount]'s query parameters as
// `url.Values`.
func (r LedgerTransactionListParamsAmount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerTransactionListParamsLedgerableType string

const (
	LedgerTransactionListParamsLedgerableTypeExpectedPayment       LedgerTransactionListParamsLedgerableType = "expected_payment"
	LedgerTransactionListParamsLedgerableTypeIncomingPaymentDetail LedgerTransactionListParamsLedgerableType = "incoming_payment_detail"
	LedgerTransactionListParamsLedgerableTypePaymentOrder          LedgerTransactionListParamsLedgerableType = "payment_order"
	LedgerTransactionListParamsLedgerableTypeReturn                LedgerTransactionListParamsLedgerableType = "return"
	LedgerTransactionListParamsLedgerableTypeReversal              LedgerTransactionListParamsLedgerableType = "reversal"
)

func (r LedgerTransactionListParamsLedgerableType) IsKnown() bool {
	switch r {
	case LedgerTransactionListParamsLedgerableTypeExpectedPayment, LedgerTransactionListParamsLedgerableTypeIncomingPaymentDetail, LedgerTransactionListParamsLedgerableTypePaymentOrder, LedgerTransactionListParamsLedgerableTypeReturn, LedgerTransactionListParamsLedgerableTypeReversal:
		return true
	}
	return false
}

// Order by `created_at` or `effective_at` in `asc` or `desc` order. For example,
// to order by `effective_at asc`, use `order_by%5Beffective_at%5D=asc`. Ordering
// by only one field at a time is supported.
type LedgerTransactionListParamsOrderBy struct {
	CreatedAt   param.Field[LedgerTransactionListParamsOrderByCreatedAt]   `query:"created_at"`
	EffectiveAt param.Field[LedgerTransactionListParamsOrderByEffectiveAt] `query:"effective_at"`
}

// URLQuery serializes [LedgerTransactionListParamsOrderBy]'s query parameters as
// `url.Values`.
func (r LedgerTransactionListParamsOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LedgerTransactionListParamsOrderByCreatedAt string

const (
	LedgerTransactionListParamsOrderByCreatedAtAsc  LedgerTransactionListParamsOrderByCreatedAt = "asc"
	LedgerTransactionListParamsOrderByCreatedAtDesc LedgerTransactionListParamsOrderByCreatedAt = "desc"
)

func (r LedgerTransactionListParamsOrderByCreatedAt) IsKnown() bool {
	switch r {
	case LedgerTransactionListParamsOrderByCreatedAtAsc, LedgerTransactionListParamsOrderByCreatedAtDesc:
		return true
	}
	return false
}

type LedgerTransactionListParamsOrderByEffectiveAt string

const (
	LedgerTransactionListParamsOrderByEffectiveAtAsc  LedgerTransactionListParamsOrderByEffectiveAt = "asc"
	LedgerTransactionListParamsOrderByEffectiveAtDesc LedgerTransactionListParamsOrderByEffectiveAt = "desc"
)

func (r LedgerTransactionListParamsOrderByEffectiveAt) IsKnown() bool {
	switch r {
	case LedgerTransactionListParamsOrderByEffectiveAtAsc, LedgerTransactionListParamsOrderByEffectiveAtDesc:
		return true
	}
	return false
}

// Satisfied by [LedgerTransactionListParamsStatusString],
// [LedgerTransactionListParamsStatusArray].
type LedgerTransactionListParamsStatusUnion interface {
	implementsLedgerTransactionListParamsStatusUnion()
}

type LedgerTransactionListParamsStatusString string

const (
	LedgerTransactionListParamsStatusStringPending  LedgerTransactionListParamsStatusString = "pending"
	LedgerTransactionListParamsStatusStringPosted   LedgerTransactionListParamsStatusString = "posted"
	LedgerTransactionListParamsStatusStringArchived LedgerTransactionListParamsStatusString = "archived"
)

func (r LedgerTransactionListParamsStatusString) IsKnown() bool {
	switch r {
	case LedgerTransactionListParamsStatusStringPending, LedgerTransactionListParamsStatusStringPosted, LedgerTransactionListParamsStatusStringArchived:
		return true
	}
	return false
}

func (r LedgerTransactionListParamsStatusString) implementsLedgerTransactionListParamsStatusUnion() {}

type LedgerTransactionListParamsStatusArray []LedgerTransactionListParamsStatusArrayItem

func (r LedgerTransactionListParamsStatusArray) implementsLedgerTransactionListParamsStatusUnion() {}

type LedgerTransactionListParamsStatusArrayItem string

const (
	LedgerTransactionListParamsStatusArrayItemPending  LedgerTransactionListParamsStatusArrayItem = "pending"
	LedgerTransactionListParamsStatusArrayItemPosted   LedgerTransactionListParamsStatusArrayItem = "posted"
	LedgerTransactionListParamsStatusArrayItemArchived LedgerTransactionListParamsStatusArrayItem = "archived"
)

func (r LedgerTransactionListParamsStatusArrayItem) IsKnown() bool {
	switch r {
	case LedgerTransactionListParamsStatusArrayItemPending, LedgerTransactionListParamsStatusArrayItemPosted, LedgerTransactionListParamsStatusArrayItemArchived:
		return true
	}
	return false
}

type LedgerTransactionNewPartialPostParams struct {
	// An array of ledger entry objects to be set on the posted ledger transaction.
	// There must be one entry for each of the existing entries with a lesser amount
	// than the existing entry.
	PostedLedgerEntries param.Field[[]LedgerTransactionNewPartialPostParamsPostedLedgerEntry] `json:"posted_ledger_entries,required"`
	// An optional free-form description for the posted ledger transaction. Maximum of
	// 1000 characters allowed.
	Description param.Field[string] `json:"description"`
	// The timestamp (IS08601 format) at which the posted ledger transaction happened
	// for reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerTransactionNewPartialPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerTransactionNewPartialPostParamsPostedLedgerEntry struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerTransactionNewPartialPostParamsPostedLedgerEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of `credit`, `debit`. Describes the direction money is flowing in the
// transaction. A `credit` moves money from your account to someone else's. A
// `debit` pulls money from someone else's account to your own. Note that wire,
// rtp, and check payments will always be `credit`.
type LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirection string

const (
	LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirectionCredit LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirection = "credit"
	LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirectionDebit  LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirection = "debit"
)

func (r LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirection) IsKnown() bool {
	switch r {
	case LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirectionCredit, LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirectionDebit:
		return true
	}
	return false
}

type LedgerTransactionNewReversalParams struct {
	// An optional free-form description for the reversal ledger transaction. Maximum
	// of 1000 characters allowed.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the reversal ledger transaction happened
	// for reporting purposes. It defaults to the `effective_at` of the original ledger
	// transaction if not provided.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// Must be unique within the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// Specify this if you'd like to link the reversal ledger transaction to a Payment
	// object like Return or Reversal.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// Specify this if you'd like to link the reversal ledger transaction to a Payment
	// object like Return or Reversal.
	LedgerableType param.Field[LedgerTransactionNewReversalParamsLedgerableType] `json:"ledgerable_type"`
	// Additional data to be added to the reversal ledger transaction as key-value
	// pairs. Both the key and value must be strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Status of the reversal ledger transaction. It defaults to `posted` if not
	// provided.
	Status param.Field[LedgerTransactionNewReversalParamsStatus] `json:"status"`
}

func (r LedgerTransactionNewReversalParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify this if you'd like to link the reversal ledger transaction to a Payment
// object like Return or Reversal.
type LedgerTransactionNewReversalParamsLedgerableType string

const (
	LedgerTransactionNewReversalParamsLedgerableTypeExpectedPayment       LedgerTransactionNewReversalParamsLedgerableType = "expected_payment"
	LedgerTransactionNewReversalParamsLedgerableTypeIncomingPaymentDetail LedgerTransactionNewReversalParamsLedgerableType = "incoming_payment_detail"
	LedgerTransactionNewReversalParamsLedgerableTypePaymentOrder          LedgerTransactionNewReversalParamsLedgerableType = "payment_order"
	LedgerTransactionNewReversalParamsLedgerableTypeReturn                LedgerTransactionNewReversalParamsLedgerableType = "return"
	LedgerTransactionNewReversalParamsLedgerableTypeReversal              LedgerTransactionNewReversalParamsLedgerableType = "reversal"
)

func (r LedgerTransactionNewReversalParamsLedgerableType) IsKnown() bool {
	switch r {
	case LedgerTransactionNewReversalParamsLedgerableTypeExpectedPayment, LedgerTransactionNewReversalParamsLedgerableTypeIncomingPaymentDetail, LedgerTransactionNewReversalParamsLedgerableTypePaymentOrder, LedgerTransactionNewReversalParamsLedgerableTypeReturn, LedgerTransactionNewReversalParamsLedgerableTypeReversal:
		return true
	}
	return false
}

// Status of the reversal ledger transaction. It defaults to `posted` if not
// provided.
type LedgerTransactionNewReversalParamsStatus string

const (
	LedgerTransactionNewReversalParamsStatusArchived LedgerTransactionNewReversalParamsStatus = "archived"
	LedgerTransactionNewReversalParamsStatusPending  LedgerTransactionNewReversalParamsStatus = "pending"
	LedgerTransactionNewReversalParamsStatusPosted   LedgerTransactionNewReversalParamsStatus = "posted"
)

func (r LedgerTransactionNewReversalParamsStatus) IsKnown() bool {
	switch r {
	case LedgerTransactionNewReversalParamsStatusArchived, LedgerTransactionNewReversalParamsStatusPending, LedgerTransactionNewReversalParamsStatusPosted:
		return true
	}
	return false
}
