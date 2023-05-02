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

type LineItemService struct {
	Options []option.RequestOption
}

func NewLineItemService(opts ...option.RequestOption) (r *LineItemService) {
	r = &LineItemService{}
	r.Options = opts
	return
}

// Get a single line item
func (r *LineItemService) Get(ctx context.Context, itemizable_type LineItemGetParamsItemizableType, itemizable_id string, id string, opts ...option.RequestOption) (res *LineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/line_items/%s", itemizable_type, itemizable_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update line item
func (r *LineItemService) Update(ctx context.Context, itemizable_type LineItemUpdateParamsItemizableType, itemizable_id string, id string, body LineItemUpdateParams, opts ...option.RequestOption) (res *LineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/%s/%s/line_items/%s", itemizable_type, itemizable_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of line items
func (r *LineItemService) List(ctx context.Context, itemizable_type LineItemListParamsItemizableType, itemizable_id string, query LineItemListParams, opts ...option.RequestOption) (res *shared.Page[LineItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/%s/%s/line_items", itemizable_type, itemizable_id)
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
func (r *LineItemService) ListAutoPaging(ctx context.Context, itemizable_type LineItemListParamsItemizableType, itemizable_id string, query LineItemListParams, opts ...option.RequestOption) *shared.PageAutoPager[LineItem] {
	return shared.NewPageAutoPager(r.List(ctx, itemizable_type, itemizable_id, query, opts...))
}

type LineItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the payment order or expected payment.
	ItemizableID string `json:"itemizable_id,required" format:"uuid"`
	// One of `payment_orders` or `expected_payments`.
	ItemizableType LineItemItemizableType `json:"itemizable_type,required"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount int64 `json:"amount,required"`
	// A free-form description of the line item.
	Description string `json:"description,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata   map[string]string  `json:"metadata,required"`
	Accounting LineItemAccounting `json:"accounting,required"`
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountingCategoryID string `json:"accounting_category_id,required,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	AccountingLedgerClassID string `json:"accounting_ledger_class_id,required,nullable" format:"uuid"`
	JSON                    LineItemJSON
}

type LineItemJSON struct {
	ID                      apijson.Metadata
	Object                  apijson.Metadata
	LiveMode                apijson.Metadata
	CreatedAt               apijson.Metadata
	UpdatedAt               apijson.Metadata
	ItemizableID            apijson.Metadata
	ItemizableType          apijson.Metadata
	Amount                  apijson.Metadata
	Description             apijson.Metadata
	Metadata                apijson.Metadata
	Accounting              apijson.Metadata
	AccountingCategoryID    apijson.Metadata
	AccountingLedgerClassID apijson.Metadata
	raw                     string
	Extras                  map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LineItem using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *LineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LineItemItemizableType string

const (
	LineItemItemizableTypeExpectedPayment LineItemItemizableType = "ExpectedPayment"
	LineItemItemizableTypePaymentOrder    LineItemItemizableType = "PaymentOrder"
)

type LineItemAccounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID string `json:"class_id,nullable" format:"uuid"`
	JSON    LineItemAccountingJSON
}

type LineItemAccountingJSON struct {
	AccountID apijson.Metadata
	ClassID   apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into LineItemAccounting using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *LineItemAccounting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LineItemGetParamsItemizableType string

const (
	LineItemGetParamsItemizableTypeExpectedPayments LineItemGetParamsItemizableType = "expected_payments"
	LineItemGetParamsItemizableTypePaymentOrders    LineItemGetParamsItemizableType = "payment_orders"
)

type LineItemUpdateParams struct {
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes LineItemUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r LineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LineItemUpdateParamsItemizableType string

const (
	LineItemUpdateParamsItemizableTypeExpectedPayments LineItemUpdateParamsItemizableType = "expected_payments"
	LineItemUpdateParamsItemizableTypePaymentOrders    LineItemUpdateParamsItemizableType = "payment_orders"
)

type LineItemListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes LineItemListParams into a url.Values of the query parameters
// associated with this value
func (r LineItemListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type LineItemListParamsItemizableType string

const (
	LineItemListParamsItemizableTypeExpectedPayments LineItemListParamsItemizableType = "expected_payments"
	LineItemListParamsItemizableTypePaymentOrders    LineItemListParamsItemizableType = "payment_orders"
)
