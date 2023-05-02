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

type InvoiceLineItemService struct {
	Options []option.RequestOption
}

func NewInvoiceLineItemService(opts ...option.RequestOption) (r *InvoiceLineItemService) {
	r = &InvoiceLineItemService{}
	r.Options = opts
	return
}

// create invoice_line_item
func (r *InvoiceLineItemService) New(ctx context.Context, invoice_id string, body InvoiceLineItemNewParams, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items", invoice_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get invoice_line_item
func (r *InvoiceLineItemService) Get(ctx context.Context, invoice_id string, id string, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoice_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update invoice_line_item
func (r *InvoiceLineItemService) Update(ctx context.Context, invoice_id string, id string, body InvoiceLineItemUpdateParams, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoice_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list invoice_line_items
func (r *InvoiceLineItemService) List(ctx context.Context, invoice_id string, query InvoiceLineItemListParams, opts ...option.RequestOption) (res *shared.Page[InvoiceLineItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items", invoice_id)
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
func (r *InvoiceLineItemService) ListAutoPaging(ctx context.Context, invoice_id string, query InvoiceLineItemListParams, opts ...option.RequestOption) *shared.PageAutoPager[InvoiceLineItem] {
	return shared.NewPageAutoPager(r.List(ctx, invoice_id, query, opts...))
}

// delete invoice_line_item
func (r *InvoiceLineItemService) Delete(ctx context.Context, invoice_id string, id string, opts ...option.RequestOption) (res *InvoiceLineItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s/invoice_line_items/%s", invoice_id, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type InvoiceLineItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The name of the line item, typically a product or SKU name.
	Name string `json:"name,required"`
	// An optional free-form description of the line item.
	Description string `json:"description,required"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity int64 `json:"quantity,required"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit.
	UnitAmount int64 `json:"unit_amount,required"`
	// Either `debit` or `credit`. `debit` indicates that a client owes the business
	// money and increases the invoice's `total_amount` due. `credit` has the opposite
	// intention and effect.
	Direction string `json:"direction,required"`
	// The total amount for this line item specified in the invoice currency's smallest
	// unit.
	Amount int64 `json:"amount,required"`
	JSON   InvoiceLineItemJSON
}

type InvoiceLineItemJSON struct {
	ID          apijson.Metadata
	Object      apijson.Metadata
	LiveMode    apijson.Metadata
	CreatedAt   apijson.Metadata
	UpdatedAt   apijson.Metadata
	Name        apijson.Metadata
	Description apijson.Metadata
	Quantity    apijson.Metadata
	UnitAmount  apijson.Metadata
	Direction   apijson.Metadata
	Amount      apijson.Metadata
	raw         string
	Extras      map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InvoiceLineItem using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemNewParams struct {
	// The name of the line item, typically a product or SKU name.
	Name field.Field[string] `json:"name,required"`
	// An optional free-form description of the line item.
	Description field.Field[string] `json:"description"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity field.Field[int64] `json:"quantity"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit.
	UnitAmount field.Field[int64] `json:"unit_amount,required"`
	// Either `debit` or `credit`. `debit` indicates that a client owes the business
	// money and increases the invoice's `total_amount` due. `credit` has the opposite
	// intention and effect.
	Direction field.Field[string] `json:"direction"`
}

// MarshalJSON serializes InvoiceLineItemNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InvoiceLineItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails field.Field[[]InvoiceLineItemUpdateParamsContactDetails] `json:"contact_details"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
	// The counterparty's billing address.
	CounterpartyBillingAddress field.Field[InvoiceLineItemUpdateParamsCounterpartyBillingAddress] `json:"counterparty_billing_address,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress field.Field[InvoiceLineItemUpdateParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency field.Field[shared.Currency] `json:"currency,nullable"`
	// A free-form description of the invoice.
	Description field.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate field.Field[time.Time] `json:"due_date" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress field.Field[InvoiceLineItemUpdateParamsInvoicerAddress] `json:"invoicer_address,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID field.Field[string] `json:"originating_account_id"`
}

// MarshalJSON serializes InvoiceLineItemUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InvoiceLineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParamsContactDetails struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              field.Field[bool]                                                           `json:"live_mode,required"`
	CreatedAt             field.Field[time.Time]                                                      `json:"created_at,required" format:"date-time"`
	UpdatedAt             field.Field[time.Time]                                                      `json:"updated_at,required" format:"date-time"`
	DiscardedAt           field.Field[time.Time]                                                      `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     field.Field[string]                                                         `json:"contact_identifier,required"`
	ContactIdentifierType field.Field[InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

type InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType string

const (
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail       InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "email"
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeWebsite     InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "website"
)

type InvoiceLineItemUpdateParamsCounterpartyBillingAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceLineItemUpdateParamsCounterpartyShippingAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceLineItemUpdateParamsInvoicerAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InvoiceLineItemListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes InvoiceLineItemListParams into a url.Values of the query
// parameters associated with this value
func (r InvoiceLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
