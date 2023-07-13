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

// InvoiceLineItemService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewInvoiceLineItemService] method instead.
type InvoiceLineItemService struct {
	Options []option.RequestOption
}

// NewInvoiceLineItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvoiceLineItemService(opts ...option.RequestOption) (r *InvoiceLineItemService) {
	r = &InvoiceLineItemService{}
	r.Options = opts
	return
}

// create invoice_line_item
func (r *InvoiceLineItemService) New(ctx context.Context, invoiceID string, params InvoiceLineItemNewParams, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// get invoice_line_item
func (r *InvoiceLineItemService) Get(ctx context.Context, invoiceID string, id string, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoiceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update invoice_line_item
func (r *InvoiceLineItemService) Update(ctx context.Context, invoiceID string, id string, body InvoiceLineItemUpdateParams, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoiceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list invoice_line_items
func (r *InvoiceLineItemService) List(ctx context.Context, invoiceID string, query InvoiceLineItemListParams, opts ...option.RequestOption) (res *shared.Page[InvoiceLineItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items", invoiceID)
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

// list invoice_line_items
func (r *InvoiceLineItemService) ListAutoPaging(ctx context.Context, invoiceID string, query InvoiceLineItemListParams, opts ...option.RequestOption) *shared.PageAutoPager[InvoiceLineItem] {
	return shared.NewPageAutoPager(r.List(ctx, invoiceID, query, opts...))
}

// delete invoice_line_item
func (r *InvoiceLineItemService) Delete(ctx context.Context, invoiceID string, id string, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoiceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type InvoiceLineItem struct {
	ID string `json:"id,required" format:"uuid"`
	// The total amount for this line item specified in the invoice currency's smallest
	// unit.
	Amount    int64     `json:"amount,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// An optional free-form description of the line item.
	Description string `json:"description,required"`
	// Either `debit` or `credit`. `debit` indicates that a client owes the business
	// money and increases the invoice's `total_amount` due. `credit` has the opposite
	// intention and effect.
	Direction string `json:"direction,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// The name of the line item, typically a product or SKU name.
	Name   string `json:"name,required"`
	Object string `json:"object,required"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity int64 `json:"quantity,required"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit.
	UnitAmount int64     `json:"unit_amount,required"`
	UpdatedAt  time.Time `json:"updated_at,required" format:"date-time"`
	JSON       invoiceLineItemJSON
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Direction   apijson.Field
	LiveMode    apijson.Field
	Name        apijson.Field
	Object      apijson.Field
	Quantity    apijson.Field
	UnitAmount  apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemNewParams struct {
	// The name of the line item, typically a product or SKU name.
	Name param.Field[string] `json:"name,required"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit.
	UnitAmount param.Field[int64] `json:"unit_amount,required"`
	// An optional free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// Either `debit` or `credit`. `debit` indicates that a client owes the business
	// money and increases the invoice's `total_amount` due. `credit` has the opposite
	// intention and effect.
	Direction param.Field[string] `json:"direction"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity       param.Field[int64]  `json:"quantity"`
	IdempotencyKey param.Field[string] `header:"Idempotency-Key"`
}

func (r InvoiceLineItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails param.Field[[]InvoiceLineItemUpdateParamsContactDetail] `json:"contact_details"`
	// The counterparty's billing address.
	CounterpartyBillingAddress param.Field[InvoiceLineItemUpdateParamsCounterpartyBillingAddress] `json:"counterparty_billing_address"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID param.Field[string] `json:"counterparty_id"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress param.Field[InvoiceLineItemUpdateParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency param.Field[shared.Currency] `json:"currency"`
	// A free-form description of the invoice.
	Description param.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate param.Field[time.Time] `json:"due_date" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress param.Field[InvoiceLineItemUpdateParamsInvoicerAddress] `json:"invoicer_address"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID param.Field[string] `json:"originating_account_id"`
}

func (r InvoiceLineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParamsContactDetail struct {
	ID                    param.Field[string]                                                         `json:"id,required" format:"uuid"`
	ContactIdentifier     param.Field[string]                                                         `json:"contact_identifier,required"`
	ContactIdentifierType param.Field[InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
	CreatedAt             param.Field[time.Time]                                                      `json:"created_at,required" format:"date-time"`
	DiscardedAt           param.Field[time.Time]                                                      `json:"discarded_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  param.Field[bool]      `json:"live_mode,required"`
	Object    param.Field[string]    `json:"object,required"`
	UpdatedAt param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
}

func (r InvoiceLineItemUpdateParamsContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType string

const (
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail       InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "email"
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeWebsite     InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "website"
)

// The counterparty's billing address.
type InvoiceLineItemUpdateParamsCounterpartyBillingAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
	Line1   param.Field[string] `json:"line1,required"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	Line2  param.Field[string] `json:"line2"`
}

func (r InvoiceLineItemUpdateParamsCounterpartyBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceLineItemUpdateParamsCounterpartyShippingAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
	Line1   param.Field[string] `json:"line1,required"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	Line2  param.Field[string] `json:"line2"`
}

func (r InvoiceLineItemUpdateParamsCounterpartyShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice issuer's business address.
type InvoiceLineItemUpdateParamsInvoicerAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
	Line1   param.Field[string] `json:"line1,required"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	Line2  param.Field[string] `json:"line2"`
}

func (r InvoiceLineItemUpdateParamsInvoicerAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [InvoiceLineItemListParams]'s query parameters as
// `url.Values`.
func (r InvoiceLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
