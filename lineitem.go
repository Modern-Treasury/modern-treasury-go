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
)

// LineItemService contains methods and other services that help with interacting
// with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLineItemService] method instead.
type LineItemService struct {
	Options []option.RequestOption
}

// NewLineItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLineItemService(opts ...option.RequestOption) (r *LineItemService) {
	r = &LineItemService{}
	r.Options = opts
	return
}

// Get a single line item
func (r *LineItemService) Get(ctx context.Context, itemizableType LineItemGetParamsItemizableType, itemizableID string, id string, opts ...option.RequestOption) (res *LineItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemizableID == "" {
		err = errors.New("missing required itemizable_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/line_items/%s", itemizableType, itemizableID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update line item
func (r *LineItemService) Update(ctx context.Context, itemizableType LineItemUpdateParamsItemizableType, itemizableID string, id string, body LineItemUpdateParams, opts ...option.RequestOption) (res *LineItem, err error) {
	opts = slices.Concat(r.Options, opts)
	if itemizableID == "" {
		err = errors.New("missing required itemizable_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/line_items/%s", itemizableType, itemizableID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of line items
func (r *LineItemService) List(ctx context.Context, itemizableType LineItemListParamsItemizableType, itemizableID string, query LineItemListParams, opts ...option.RequestOption) (res *pagination.Page[LineItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if itemizableID == "" {
		err = errors.New("missing required itemizable_id parameter")
		return
	}
	path := fmt.Sprintf("api/%v/%s/line_items", itemizableType, itemizableID)
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

// Get a list of line items
func (r *LineItemService) ListAutoPaging(ctx context.Context, itemizableType LineItemListParamsItemizableType, itemizableID string, query LineItemListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LineItem] {
	return pagination.NewPageAutoPager(r.List(ctx, itemizableType, itemizableID, query, opts...))
}

type LineItem struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Accounting LineItemAccounting `json:"accounting,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingCategoryID string `json:"accounting_category_id,required,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountingLedgerClassID string `json:"accounting_ledger_class_id,required,nullable" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount    int64     `json:"amount,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A free-form description of the line item.
	Description string `json:"description,required,nullable"`
	// The ID of the payment order or expected payment.
	ItemizableID string `json:"itemizable_id,required" format:"uuid"`
	// One of `payment_orders` or `expected_payments`.
	ItemizableType LineItemItemizableType `json:"itemizable_type,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata  map[string]string `json:"metadata,required"`
	Object    string            `json:"object,required"`
	UpdatedAt time.Time         `json:"updated_at,required" format:"date-time"`
	JSON      lineItemJSON      `json:"-"`
}

// lineItemJSON contains the JSON metadata for the struct [LineItem]
type lineItemJSON struct {
	ID                      apijson.Field
	Accounting              apijson.Field
	AccountingCategoryID    apijson.Field
	AccountingLedgerClassID apijson.Field
	Amount                  apijson.Field
	CreatedAt               apijson.Field
	Description             apijson.Field
	ItemizableID            apijson.Field
	ItemizableType          apijson.Field
	LiveMode                apijson.Field
	Metadata                apijson.Field
	Object                  apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lineItemJSON) RawJSON() string {
	return r.raw
}

// Deprecated: deprecated
type LineItemAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	//
	// Deprecated: deprecated
	ClassID string                 `json:"class_id,nullable" format:"uuid"`
	JSON    lineItemAccountingJSON `json:"-"`
}

// lineItemAccountingJSON contains the JSON metadata for the struct
// [LineItemAccounting]
type lineItemAccountingJSON struct {
	AccountID   apijson.Field
	ClassID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LineItemAccounting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lineItemAccountingJSON) RawJSON() string {
	return r.raw
}

// One of `payment_orders` or `expected_payments`.
type LineItemItemizableType string

const (
	LineItemItemizableTypeExpectedPayment LineItemItemizableType = "ExpectedPayment"
	LineItemItemizableTypePaymentOrder    LineItemItemizableType = "PaymentOrder"
)

func (r LineItemItemizableType) IsKnown() bool {
	switch r {
	case LineItemItemizableTypeExpectedPayment, LineItemItemizableTypePaymentOrder:
		return true
	}
	return false
}

type LineItemGetParamsItemizableType string

const (
	LineItemGetParamsItemizableTypeExpectedPayments LineItemGetParamsItemizableType = "expected_payments"
	LineItemGetParamsItemizableTypePaymentOrders    LineItemGetParamsItemizableType = "payment_orders"
)

func (r LineItemGetParamsItemizableType) IsKnown() bool {
	switch r {
	case LineItemGetParamsItemizableTypeExpectedPayments, LineItemGetParamsItemizableTypePaymentOrders:
		return true
	}
	return false
}

type LineItemUpdateParams struct {
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LineItemUpdateParamsItemizableType string

const (
	LineItemUpdateParamsItemizableTypeExpectedPayments LineItemUpdateParamsItemizableType = "expected_payments"
	LineItemUpdateParamsItemizableTypePaymentOrders    LineItemUpdateParamsItemizableType = "payment_orders"
)

func (r LineItemUpdateParamsItemizableType) IsKnown() bool {
	switch r {
	case LineItemUpdateParamsItemizableTypeExpectedPayments, LineItemUpdateParamsItemizableTypePaymentOrders:
		return true
	}
	return false
}

type LineItemListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [LineItemListParams]'s query parameters as `url.Values`.
func (r LineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LineItemListParamsItemizableType string

const (
	LineItemListParamsItemizableTypeExpectedPayments LineItemListParamsItemizableType = "expected_payments"
	LineItemListParamsItemizableTypePaymentOrders    LineItemListParamsItemizableType = "payment_orders"
)

func (r LineItemListParamsItemizableType) IsKnown() bool {
	switch r {
	case LineItemListParamsItemizableTypeExpectedPayments, LineItemListParamsItemizableTypePaymentOrders:
		return true
	}
	return false
}
