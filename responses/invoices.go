package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

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
	Currency Currency `json:"currency,required,nullable"`
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
	// One of `draft`, `unpaid`, `payment_pending`, `paid`, `void`, or `closed`.
	Status string `json:"status,required"`
	// Total amount due in specified currency's smallest unit, e.g., $10 USD would be
	// represented as 1000.
	TotalAmount int64 `json:"total_amount,required"`
	JSON        InvoiceJSON
}

type InvoiceJSON struct {
	ID                          pjson.Metadata
	Object                      pjson.Metadata
	LiveMode                    pjson.Metadata
	CreatedAt                   pjson.Metadata
	UpdatedAt                   pjson.Metadata
	ContactDetails              pjson.Metadata
	CounterpartyID              pjson.Metadata
	CounterpartyBillingAddress  pjson.Metadata
	CounterpartyShippingAddress pjson.Metadata
	Currency                    pjson.Metadata
	Description                 pjson.Metadata
	DueDate                     pjson.Metadata
	InvoicerAddress             pjson.Metadata
	OriginatingAccountID        pjson.Metadata
	HostedURL                   pjson.Metadata
	Number                      pjson.Metadata
	PaymentOrders               pjson.Metadata
	PdfURL                      pjson.Metadata
	Status                      pjson.Metadata
	TotalAmount                 pjson.Metadata
	Raw                         []byte
	Extras                      map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Invoice using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
	JSON                  InvoiceContactDetailsJSON
}

type InvoiceContactDetailsJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	DiscardedAt           pjson.Metadata
	ContactIdentifier     pjson.Metadata
	ContactIdentifierType pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InvoiceContactDetails using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InvoiceContactDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InvoiceContactDetailsContactIdentifierType string

const (
	InvoiceContactDetailsContactIdentifierTypeEmail       InvoiceContactDetailsContactIdentifierType = "email"
	InvoiceContactDetailsContactIdentifierTypePhoneNumber InvoiceContactDetailsContactIdentifierType = "phone_number"
	InvoiceContactDetailsContactIdentifierTypeWebsite     InvoiceContactDetailsContactIdentifierType = "website"
)

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
	JSON    InvoiceCounterpartyBillingAddressJSON
}

type InvoiceCounterpartyBillingAddressJSON struct {
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	Locality   pjson.Metadata
	Region     pjson.Metadata
	PostalCode pjson.Metadata
	Country    pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InvoiceCounterpartyBillingAddress using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *InvoiceCounterpartyBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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
	JSON    InvoiceCounterpartyShippingAddressJSON
}

type InvoiceCounterpartyShippingAddressJSON struct {
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	Locality   pjson.Metadata
	Region     pjson.Metadata
	PostalCode pjson.Metadata
	Country    pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// InvoiceCounterpartyShippingAddress using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *InvoiceCounterpartyShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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
	JSON    InvoiceInvoicerAddressJSON
}

type InvoiceInvoicerAddressJSON struct {
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	Locality   pjson.Metadata
	Region     pjson.Metadata
	PostalCode pjson.Metadata
	Country    pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InvoiceInvoicerAddress using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InvoiceInvoicerAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
