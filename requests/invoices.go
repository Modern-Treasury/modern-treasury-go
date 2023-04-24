package requests

import (
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/core/field"
	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/core/query"
)

type InvoiceNewParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails field.Field[[]InvoiceNewParamsContactDetails] `json:"contact_details"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID field.Field[string] `json:"counterparty_id,required"`
	// The counterparty's billing address.
	CounterpartyBillingAddress field.Field[InvoiceNewParamsCounterpartyBillingAddress] `json:"counterparty_billing_address,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress field.Field[InvoiceNewParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// A free-form description of the invoice.
	Description field.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate field.Field[time.Time] `json:"due_date,required" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress field.Field[InvoiceNewParamsInvoicerAddress] `json:"invoicer_address,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID field.Field[string] `json:"originating_account_id,required"`
}

// MarshalJSON serializes InvoiceNewParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type InvoiceNewParamsContactDetails struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              field.Field[bool]                                                `json:"live_mode,required"`
	CreatedAt             field.Field[time.Time]                                           `json:"created_at,required" format:"date-time"`
	UpdatedAt             field.Field[time.Time]                                           `json:"updated_at,required" format:"date-time"`
	DiscardedAt           field.Field[time.Time]                                           `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     field.Field[string]                                              `json:"contact_identifier,required"`
	ContactIdentifierType field.Field[InvoiceNewParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

type InvoiceNewParamsContactDetailsContactIdentifierType string

const (
	InvoiceNewParamsContactDetailsContactIdentifierTypeEmail       InvoiceNewParamsContactDetailsContactIdentifierType = "email"
	InvoiceNewParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceNewParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceNewParamsContactDetailsContactIdentifierTypeWebsite     InvoiceNewParamsContactDetailsContactIdentifierType = "website"
)

type InvoiceNewParamsCounterpartyBillingAddress struct {
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

type InvoiceNewParamsCounterpartyShippingAddress struct {
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

type InvoiceNewParamsInvoicerAddress struct {
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

type InvoiceUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails field.Field[[]InvoiceUpdateParamsContactDetails] `json:"contact_details"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
	// The counterparty's billing address.
	CounterpartyBillingAddress field.Field[InvoiceUpdateParamsCounterpartyBillingAddress] `json:"counterparty_billing_address,nullable"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress field.Field[InvoiceUpdateParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address,nullable"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency field.Field[Currency] `json:"currency,nullable"`
	// A free-form description of the invoice.
	Description field.Field[string] `json:"description"`
	// A future date by when the invoice needs to be paid.
	DueDate field.Field[time.Time] `json:"due_date" format:"date-time"`
	// The invoice issuer's business address.
	InvoicerAddress field.Field[InvoiceUpdateParamsInvoicerAddress] `json:"invoicer_address,nullable"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID field.Field[string] `json:"originating_account_id"`
	// When opening an invoice, whether to show the embedded payment UI with the
	// invoice. Default true.
	IncludePaymentUi field.Field[bool] `json:"include_payment_ui"`
	// Invoice status must be updated in a `PATCH` request that does not modify any
	// other invoice attributes. Valid state transitions are `draft` to `unpaid`,
	// `draft` or `unpaid` to `void` and `unpaid` to `closed`.
	Status field.Field[string] `json:"status"`
}

// MarshalJSON serializes InvoiceUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

type InvoiceUpdateParamsContactDetails struct {
	ID     field.Field[string] `json:"id,required" format:"uuid"`
	Object field.Field[string] `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              field.Field[bool]                                                   `json:"live_mode,required"`
	CreatedAt             field.Field[time.Time]                                              `json:"created_at,required" format:"date-time"`
	UpdatedAt             field.Field[time.Time]                                              `json:"updated_at,required" format:"date-time"`
	DiscardedAt           field.Field[time.Time]                                              `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     field.Field[string]                                                 `json:"contact_identifier,required"`
	ContactIdentifierType field.Field[InvoiceUpdateParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type,required"`
}

type InvoiceUpdateParamsContactDetailsContactIdentifierType string

const (
	InvoiceUpdateParamsContactDetailsContactIdentifierTypeEmail       InvoiceUpdateParamsContactDetailsContactIdentifierType = "email"
	InvoiceUpdateParamsContactDetailsContactIdentifierTypePhoneNumber InvoiceUpdateParamsContactDetailsContactIdentifierType = "phone_number"
	InvoiceUpdateParamsContactDetailsContactIdentifierTypeWebsite     InvoiceUpdateParamsContactDetailsContactIdentifierType = "website"
)

type InvoiceUpdateParamsCounterpartyBillingAddress struct {
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

type InvoiceUpdateParamsCounterpartyShippingAddress struct {
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

type InvoiceUpdateParamsInvoicerAddress struct {
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

type InvoiceListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
}

// URLQuery serializes InvoiceListParams into a url.Values of the query parameters
// associated with this value
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
