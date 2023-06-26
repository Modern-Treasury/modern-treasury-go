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

// TransactionLineItemService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewTransactionLineItemService] method instead.
type TransactionLineItemService struct {
	Options []option.RequestOption
}

// NewTransactionLineItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionLineItemService(opts ...option.RequestOption) (r *TransactionLineItemService) {
	r = &TransactionLineItemService{}
	r.Options = opts
	return
}

// list transaction_line_items
func (r *TransactionLineItemService) List(ctx context.Context, transactionID string, query TransactionLineItemListParams, opts ...option.RequestOption) (res *shared.Page[TransactionLineItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/transactions/%s/line_items", transactionID)
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

// list transaction_line_items
func (r *TransactionLineItemService) ListAutoPaging(ctx context.Context, transactionID string, query TransactionLineItemListParams, opts ...option.RequestOption) *shared.PageAutoPager[TransactionLineItem] {
	return shared.NewPageAutoPager(r.List(ctx, transactionID, query, opts...))
}

type TransactionLineItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment, or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Indicates whether the line item is `originating` or `receiving` (see
	// https://www.moderntreasury.com/journal/beginners-guide-to-ach for more).
	Type TransactionLineItemType `json:"type,required"`
	// If a matching object exists in Modern Treasury, the type will be populated here,
	// otherwise `null`.
	TransactableType TransactionLineItemTransactableType `json:"transactable_type,required,nullable"`
	// If a matching object exists in Modern Treasury, the ID will be populated here,
	// otherwise `null`.
	TransactableID string `json:"transactable_id,required,nullable"`
	// If a matching object exists in Modern Treasury, `amount` will be populated.
	// Value in specified currency's smallest unit (taken from parent Transaction).
	Amount int64 `json:"amount,required"`
	// If no matching object is found, `description` will be a free-form text field
	// describing the line item. This field may contain personally identifiable
	// information (PII) and is not included in API responses by default. Learn more
	// about changing your settings at
	// https://docs.moderntreasury.com/reference/personally-identifiable-information.
	Description string `json:"description,required"`
	// The ID for the counterparty for this transaction line item.
	CounterpartyID string `json:"counterparty_id,required"`
	// The ID of the reconciled Expected Payment, otherwise `null`.
	ExpectedPaymentID string `json:"expected_payment_id,required"`
	JSON              transactionLineItemJSON
}

// transactionLineItemJSON contains the JSON metadata for the struct
// [TransactionLineItem]
type transactionLineItemJSON struct {
	ID                apijson.Field
	Object            apijson.Field
	LiveMode          apijson.Field
	CreatedAt         apijson.Field
	UpdatedAt         apijson.Field
	DiscardedAt       apijson.Field
	Type              apijson.Field
	TransactableType  apijson.Field
	TransactableID    apijson.Field
	Amount            apijson.Field
	Description       apijson.Field
	CounterpartyID    apijson.Field
	ExpectedPaymentID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates whether the line item is `originating` or `receiving` (see
// https://www.moderntreasury.com/journal/beginners-guide-to-ach for more).
type TransactionLineItemType string

const (
	TransactionLineItemTypeOriginating TransactionLineItemType = "originating"
	TransactionLineItemTypeReceiving   TransactionLineItemType = "receiving"
)

// If a matching object exists in Modern Treasury, the type will be populated here,
// otherwise `null`.
type TransactionLineItemTransactableType string

const (
	TransactionLineItemTransactableTypeIncomingPaymentDetail TransactionLineItemTransactableType = "incoming_payment_detail"
	TransactionLineItemTransactableTypePaperItem             TransactionLineItemTransactableType = "paper_item"
	TransactionLineItemTransactableTypePaymentOrder          TransactionLineItemTransactableType = "payment_order"
	TransactionLineItemTransactableTypePaymentOrderAttempt   TransactionLineItemTransactableType = "payment_order_attempt"
	TransactionLineItemTransactableTypeReturn                TransactionLineItemTransactableType = "return"
	TransactionLineItemTransactableTypeReversal              TransactionLineItemTransactableType = "reversal"
)

type TransactionLineItemListParams struct {
	AfterCursor param.Field[string]                            `query:"after_cursor"`
	PerPage     param.Field[int64]                             `query:"per_page"`
	Type        param.Field[TransactionLineItemListParamsType] `query:"type"`
}

// URLQuery serializes [TransactionLineItemListParams]'s query parameters as
// `url.Values`.
func (r TransactionLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionLineItemListParamsType string

const (
	TransactionLineItemListParamsTypeOriginating TransactionLineItemListParamsType = "originating"
	TransactionLineItemListParamsTypeReceiving   TransactionLineItemListParamsType = "receiving"
)
