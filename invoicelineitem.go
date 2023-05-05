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
	JSON   invoiceLineItemJSON
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	LiveMode    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	Name        apijson.Field
	Description apijson.Field
	Quantity    apijson.Field
	UnitAmount  apijson.Field
	Direction   apijson.Field
	Amount      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemNewParams struct {
	// The name of the line item, typically a product or SKU name.
	Name param.Field[string] `json:"name,required"`
	// An optional free-form description of the line item.
	Description param.Field[string] `json:"description"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity param.Field[int64] `json:"quantity"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit.
	UnitAmount param.Field[int64] `json:"unit_amount,required"`
	// Either `debit` or `credit`. `debit` indicates that a client owes the business
	// money and increases the invoice's `total_amount` due. `credit` has the opposite
	// intention and effect.
	Direction param.Field[string] `json:"direction"`
}

func (r InvoiceLineItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails param.Field[[]InvoiceLineItemUpdateParamsContactDetails] `json:"contact_details"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID param.Field[string] `json:"counterparty_id"`
	// The counterparty's billing address.
	CounterpartyBillingAddress param.Field[InvoiceLineItemUpdateParamsCounterpartyBillingAddress] `json:"counterparty_billing_address,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress param.Field[InvoiceLineItemUpdateParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency param.Field[shared.Currency] `json:"currency,nullable"`
	// A free-form description of the invoice.
	Description param.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate param.Field[time.Time] `json:"due_date" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress param.Field[InvoiceLineItemUpdateParamsInvoicerAddress] `json:"invoicer_address,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID param.Field[string] `json:"originating_account_id"`
}

func (r InvoiceLineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceLineItemUpdateParamsContactDetails struct {
	ID     param.Field[string] `json:"id,required" format:"uuid"`
	Object param.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              param.Field[bool]                                                           `json:"live_mode,required"`
	CreatedAt             param.Field[time.Time]                                                      `json:"created_at,required" format:"date-time"`
	UpdatedAt             param.Field[time.Time]                                                      `json:"updated_at,required" format:"date-time"`
	DiscardedAt           param.Field[time.Time]                                                      `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     param.Field[string]                                                         `json:"contact_identifier,required"`
	ContactIdentifierType param.Field[InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

type InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType string

const (
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeEmail       InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "email"
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceLineItemUpdateParamsContactDetailsContactIdentifierTypeWebsite     InvoiceLineItemUpdateParamsContactDetailsContactIdentifierType = "website"
)

// The counterparty's billing address.
type InvoiceLineItemUpdateParamsCounterpartyBillingAddress struct {
	Line1 param.Field[string] `json:"line1,required"`
	Line2 param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceLineItemUpdateParamsCounterpartyShippingAddress struct {
	Line1 param.Field[string] `json:"line1,required"`
	Line2 param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
}

// The invoice issuer's business address.
type InvoiceLineItemUpdateParamsInvoicerAddress struct {
	Line1 param.Field[string] `json:"line1,required"`
	Line2 param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
}

type InvoiceLineItemListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [InvoiceLineItemListParams]'s query parameters as
// `url.Values`.
func (r InvoiceLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
