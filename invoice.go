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

// InvoiceService contains methods and other services that help with interacting
// with the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options   []option.RequestOption
	LineItems *InvoiceLineItemService
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	r.LineItems = NewInvoiceLineItemService(opts...)
	return
}

// create invoice
func (r *InvoiceService) New(ctx context.Context, params InvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// get invoice
func (r *InvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update invoice
func (r *InvoiceService) Update(ctx context.Context, id string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list invoices
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *shared.Page[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/invoices"
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

// list invoices
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *shared.PageAutoPager[Invoice] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Invoice struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails []InvoiceContactDetails `json:"contact_details,required"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID string `json:"counterparty_id,required"`
	// The counterparty's billing address.
	CounterpartyBillingAddress InvoiceCounterpartyBillingAddress `json:"counterparty_billing_address,required,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress InvoiceCounterpartyShippingAddress `json:"counterparty_shipping_address,required,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency shared.Currency `json:"currency,required,nullable"`
	// A free-form description of the invoice.
	Description string `json:"description,required"`
	// A future date by when the invoice needs to be paid.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress InvoiceInvoicerAddress `json:"invoicer_address,required,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID string `json:"originating_account_id,required"`
	// The URL of the hosted web UI where the invoice can be viewed.
	HostedURL string `json:"hosted_url,required"`
	// A unique record number assigned to each invoice that is issued.
	Number string `json:"number,required"`
	// The payment orders created for paying the invoice through the invoice payment
	// UI.
	PaymentOrders []PaymentOrder `json:"payment_orders,required"`
	// The URL where the invoice PDF can be downloaded.
	PdfURL string `json:"pdf_url,required,nullable"`
	// The status of the invoice.
	Status InvoiceStatus `json:"status,required"`
	// Total amount due in specified currency's smallest unit, e.g., $10 USD would be
	// represented as 1000.
	TotalAmount int64 `json:"total_amount,required"`
	JSON        invoiceJSON
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID                          apijson.Field
	Object                      apijson.Field
	LiveMode                    apijson.Field
	CreatedAt                   apijson.Field
	UpdatedAt                   apijson.Field
	ContactDetails              apijson.Field
	CounterpartyID              apijson.Field
	CounterpartyBillingAddress  apijson.Field
	CounterpartyShippingAddress apijson.Field
	Currency                    apijson.Field
	Description                 apijson.Field
	DueDate                     apijson.Field
	InvoicerAddress             apijson.Field
	OriginatingAccountID        apijson.Field
	HostedURL                   apijson.Field
	Number                      apijson.Field
	PaymentOrders               apijson.Field
	PdfURL                      apijson.Field
	Status                      apijson.Field
	TotalAmount                 apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceContactDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              bool                                       `json:"live_mode,required"`
	CreatedAt             time.Time                                  `json:"created_at,required" format:"date-time"`
	UpdatedAt             time.Time                                  `json:"updated_at,required" format:"date-time"`
	DiscardedAt           time.Time                                  `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     string                                     `json:"contact_identifier,required"`
	ContactIdentifierType InvoiceContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	JSON                  invoiceContactDetailsJSON
}

// invoiceContactDetailsJSON contains the JSON metadata for the struct
// [InvoiceContactDetails]
type invoiceContactDetailsJSON struct {
	ID                    apijson.Field
	Object                apijson.Field
	LiveMode              apijson.Field
	CreatedAt             apijson.Field
	UpdatedAt             apijson.Field
	DiscardedAt           apijson.Field
	ContactIdentifier     apijson.Field
	ContactIdentifierType apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceContactDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceContactDetailsContactIdentifierType string

const (
	InvoiceContactDetailsContactIdentifierTypeEmail       InvoiceContactDetailsContactIdentifierType = "email"
	InvoiceContactDetailsContactIdentifierTypePhoneNumber InvoiceContactDetailsContactIdentifierType = "phone_number"
	InvoiceContactDetailsContactIdentifierTypeWebsite     InvoiceContactDetailsContactIdentifierType = "website"
)

// The counterparty's billing address.
type InvoiceCounterpartyBillingAddress struct {
	Line1 string `json:"line1,required"`
	Line2 string `json:"line2"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// Region or State.
	Region string `json:"region,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	JSON    invoiceCounterpartyBillingAddressJSON
}

// invoiceCounterpartyBillingAddressJSON contains the JSON metadata for the struct
// [InvoiceCounterpartyBillingAddress]
type invoiceCounterpartyBillingAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	Region      apijson.Field
	PostalCode  apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCounterpartyBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceCounterpartyShippingAddress struct {
	Line1 string `json:"line1,required"`
	Line2 string `json:"line2"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// Region or State.
	Region string `json:"region,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	JSON    invoiceCounterpartyShippingAddressJSON
}

// invoiceCounterpartyShippingAddressJSON contains the JSON metadata for the struct
// [InvoiceCounterpartyShippingAddress]
type invoiceCounterpartyShippingAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	Region      apijson.Field
	PostalCode  apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCounterpartyShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The invoice issuer's business address.
type InvoiceInvoicerAddress struct {
	Line1 string `json:"line1,required"`
	Line2 string `json:"line2"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// Region or State.
	Region string `json:"region,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	JSON    invoiceInvoicerAddressJSON
}

// invoiceInvoicerAddressJSON contains the JSON metadata for the struct
// [InvoiceInvoicerAddress]
type invoiceInvoicerAddressJSON struct {
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	Region      apijson.Field
	PostalCode  apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceInvoicerAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusPaid           InvoiceStatus = "paid"
	InvoiceStatusPaymentPending InvoiceStatus = "payment_pending"
	InvoiceStatusUnpaid         InvoiceStatus = "unpaid"
	InvoiceStatusVoided         InvoiceStatus = "voided"
)

type InvoiceNewParams struct {
	// The ID of the counterparty receiving the invoice.
	CounterpartyID param.Field[string] `json:"counterparty_id,required"`
	// A future date by when the invoice needs to be paid.
	DueDate param.Field[time.Time] `json:"due_date,required" format:"date-time"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required"`
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails param.Field[[]InvoiceNewParamsContactDetails] `json:"contact_details"`
	// The counterparty's billing address.
	CounterpartyBillingAddress param.Field[InvoiceNewParamsCounterpartyBillingAddress] `json:"counterparty_billing_address"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress param.Field[InvoiceNewParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency param.Field[shared.Currency] `json:"currency"`
	// A free-form description of the invoice.
	Description param.Field[string] `json:"description"`
	// The invoice issuer's business address.
	InvoicerAddress param.Field[InvoiceNewParamsInvoicerAddress] `json:"invoicer_address"`
	IdempotencyKey  param.Field[string]                          `header:"Idempotency-Key"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsContactDetails struct {
	ID     param.Field[string] `json:"id,required" format:"uuid"`
	Object param.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              param.Field[bool]                                                `json:"live_mode,required"`
	CreatedAt             param.Field[time.Time]                                           `json:"created_at,required" format:"date-time"`
	UpdatedAt             param.Field[time.Time]                                           `json:"updated_at,required" format:"date-time"`
	DiscardedAt           param.Field[time.Time]                                           `json:"discarded_at,required" format:"date-time"`
	ContactIdentifier     param.Field[string]                                              `json:"contact_identifier,required"`
	ContactIdentifierType param.Field[InvoiceNewParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

func (r InvoiceNewParamsContactDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsContactDetailsContactIdentifierType string

const (
	InvoiceNewParamsContactDetailsContactIdentifierTypeEmail       InvoiceNewParamsContactDetailsContactIdentifierType = "email"
	InvoiceNewParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceNewParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceNewParamsContactDetailsContactIdentifierTypeWebsite     InvoiceNewParamsContactDetailsContactIdentifierType = "website"
)

// The counterparty's billing address.
type InvoiceNewParamsCounterpartyBillingAddress struct {
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

func (r InvoiceNewParamsCounterpartyBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceNewParamsCounterpartyShippingAddress struct {
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

func (r InvoiceNewParamsCounterpartyShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice issuer's business address.
type InvoiceNewParamsInvoicerAddress struct {
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

func (r InvoiceNewParamsInvoicerAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails param.Field[[]InvoiceUpdateParamsContactDetails] `json:"contact_details"`
	// The counterparty's billing address.
	CounterpartyBillingAddress param.Field[InvoiceUpdateParamsCounterpartyBillingAddress] `json:"counterparty_billing_address"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID param.Field[string] `json:"counterparty_id"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress param.Field[InvoiceUpdateParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency param.Field[shared.Currency] `json:"currency"`
	// A free-form description of the invoice.
	Description param.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate param.Field[time.Time] `json:"due_date" format:"date-time"`
	// When opening an invoice, whether to show the embedded payment UI with the
	// invoice. Default true.
	IncludePaymentUi param.Field[bool] `json:"include_payment_ui"`
	// The invoice issuer's business address.
	InvoicerAddress param.Field[InvoiceUpdateParamsInvoicerAddress] `json:"invoicer_address"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID param.Field[string] `json:"originating_account_id"`
	// Invoice status must be updated in a `PATCH` request that does not modify any
	// other invoice attributes. Valid state transitions are `draft` to `unpaid` and
	// `draft` or `unpaid` to `voided`.
	Status param.Field[string] `json:"status"`
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateParamsContactDetails struct {
	ID     param.Field[string] `json:"id,required" format:"uuid"`
	Object param.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              param.Field[bool]                                                   `json:"live_mode,required"`
	CreatedAt             param.Field[time.Time]                                              `json:"created_at,required" format:"date-time"`
	UpdatedAt             param.Field[time.Time]                                              `json:"updated_at,required" format:"date-time"`
	DiscardedAt           param.Field[time.Time]                                              `json:"discarded_at,required" format:"date-time"`
	ContactIdentifier     param.Field[string]                                                 `json:"contact_identifier,required"`
	ContactIdentifierType param.Field[InvoiceUpdateParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

func (r InvoiceUpdateParamsContactDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateParamsContactDetailsContactIdentifierType string

const (
	InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail       InvoiceUpdateParamsContactDetailsContactIdentifierType = "email"
	InvoiceUpdateParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceUpdateParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceUpdateParamsContactDetailsContactIdentifierTypeWebsite     InvoiceUpdateParamsContactDetailsContactIdentifierType = "website"
)

// The counterparty's billing address.
type InvoiceUpdateParamsCounterpartyBillingAddress struct {
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

func (r InvoiceUpdateParamsCounterpartyBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceUpdateParamsCounterpartyShippingAddress struct {
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

func (r InvoiceUpdateParamsCounterpartyShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice issuer's business address.
type InvoiceUpdateParamsInvoicerAddress struct {
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

func (r InvoiceUpdateParamsInvoicerAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
