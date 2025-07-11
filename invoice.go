// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// InvoiceService contains methods and other services that help with interacting
// with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceService] method instead.
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
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get invoice
func (r *InvoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update invoice
func (r *InvoiceService) Update(ctx context.Context, id string, body InvoiceUpdateParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/invoices/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list invoices
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *pagination.Page[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Invoice] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Add a payment order to an invoice.
func (r *InvoiceService) AddPaymentOrder(ctx context.Context, id string, paymentOrderID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if paymentOrderID == "" {
		err = errors.New("missing required payment_order_id parameter")
		return
	}
	path := fmt.Sprintf("api/invoices/%s/payment_orders/%s", id, paymentOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

type Invoice struct {
	ID string `json:"id,required" format:"uuid"`
	// Amount paid on the invoice in specified currency's smallest unit, e.g., $10 USD
	// would be represented as 1000.
	AmountPaid int64 `json:"amount_paid,required"`
	// Amount remaining due on the invoice in specified currency's smallest unit, e.g.,
	// $10 USD would be represented as 1000.
	AmountRemaining int64 `json:"amount_remaining,required"`
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails []shared.ContactDetail `json:"contact_details,required"`
	// The counterparty's billing address.
	CounterpartyBillingAddress InvoiceCounterpartyBillingAddress `json:"counterparty_billing_address,required,nullable"`
	// The ID of the counterparty receiving the invoice.
	CounterpartyID string `json:"counterparty_id,required"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress InvoiceCounterpartyShippingAddress `json:"counterparty_shipping_address,required,nullable"`
	CreatedAt                   time.Time                          `json:"created_at,required" format:"date-time"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency shared.Currency `json:"currency,required"`
	// An optional free-form description of the invoice.
	Description string `json:"description,required"`
	// A future date by when the invoice needs to be paid.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// The expected payments created for an unpaid invoice.
	ExpectedPayments []ExpectedPayment `json:"expected_payments,required"`
	// When payment_method is automatic, the fallback payment method to use when an
	// automatic payment fails. One of `manual` or `ui`.
	FallbackPaymentMethod string `json:"fallback_payment_method,required,nullable"`
	// The URL of the hosted web UI where the invoice can be viewed.
	HostedURL string `json:"hosted_url,required"`
	// The invoice issuer's business address.
	InvoicerAddress InvoiceInvoicerAddress `json:"invoicer_address,required,nullable"`
	// The name of the issuer for the invoice. Defaults to the name of the
	// Organization.
	InvoicerName string `json:"invoicer_name,required,nullable"`
	// The ledger account settlement object linked to the invoice.
	LedgerAccountSettlementID string `json:"ledger_account_settlement_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required,nullable"`
	// Emails in addition to the counterparty email to send invoice status
	// notifications to. At least one email is required if notifications are enabled
	// and the counterparty doesn't have an email.
	NotificationEmailAddresses []string `json:"notification_email_addresses,required,nullable"`
	// If true, the invoice will send email notifications to the invoice recipients
	// about invoice status changes.
	NotificationsEnabled bool `json:"notifications_enabled,required"`
	// A unique record number assigned to each invoice that is issued.
	Number string `json:"number,required"`
	Object string `json:"object,required"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID string `json:"originating_account_id,required"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	PaymentEffectiveDate time.Time `json:"payment_effective_date,required,nullable" format:"date"`
	// When opening an invoice, whether to show the embedded payment UI , automatically
	// debit the recipient, or rely on manual payment from the recipient.
	PaymentMethod InvoicePaymentMethod `json:"payment_method,required,nullable"`
	// The payment orders created for paying the invoice through the invoice payment
	// UI.
	PaymentOrders []PaymentOrder `json:"payment_orders,required"`
	// One of `ach` or `eft`.
	PaymentType InvoicePaymentType `json:"payment_type,required,nullable"`
	// The URL where the invoice PDF can be downloaded.
	PdfURL string `json:"pdf_url,required,nullable"`
	// The receiving account ID. Can be an `internal_account`.
	ReceivingAccountID string `json:"receiving_account_id,required,nullable" format:"uuid"`
	// The email of the recipient of the invoice. Leaving this value as null will
	// fallback to using the counterparty's name.
	RecipientEmail string `json:"recipient_email,required,nullable"`
	// The name of the recipient of the invoice. Leaving this value as null will
	// fallback to using the counterparty's name.
	RecipientName string `json:"recipient_name,required,nullable"`
	// Number of days after due date when overdue reminder emails will be sent out to
	// invoice recipients.
	RemindAfterOverdueDays []int64 `json:"remind_after_overdue_days,required,nullable"`
	// The status of the invoice.
	Status InvoiceStatus `json:"status,required"`
	// Total amount due in specified currency's smallest unit, e.g., $10 USD would be
	// represented as 1000.
	TotalAmount int64 `json:"total_amount,required"`
	// IDs of transaction line items associated with an invoice.
	TransactionLineItemIDs []string  `json:"transaction_line_item_ids,required" format:"uuid"`
	UpdatedAt              time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the virtual account the invoice should be paid to.
	VirtualAccountID string      `json:"virtual_account_id,required,nullable" format:"uuid"`
	JSON             invoiceJSON `json:"-"`
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID                          apijson.Field
	AmountPaid                  apijson.Field
	AmountRemaining             apijson.Field
	ContactDetails              apijson.Field
	CounterpartyBillingAddress  apijson.Field
	CounterpartyID              apijson.Field
	CounterpartyShippingAddress apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	Description                 apijson.Field
	DueDate                     apijson.Field
	ExpectedPayments            apijson.Field
	FallbackPaymentMethod       apijson.Field
	HostedURL                   apijson.Field
	InvoicerAddress             apijson.Field
	InvoicerName                apijson.Field
	LedgerAccountSettlementID   apijson.Field
	LiveMode                    apijson.Field
	Metadata                    apijson.Field
	NotificationEmailAddresses  apijson.Field
	NotificationsEnabled        apijson.Field
	Number                      apijson.Field
	Object                      apijson.Field
	OriginatingAccountID        apijson.Field
	PaymentEffectiveDate        apijson.Field
	PaymentMethod               apijson.Field
	PaymentOrders               apijson.Field
	PaymentType                 apijson.Field
	PdfURL                      apijson.Field
	ReceivingAccountID          apijson.Field
	RecipientEmail              apijson.Field
	RecipientName               apijson.Field
	RemindAfterOverdueDays      apijson.Field
	Status                      apijson.Field
	TotalAmount                 apijson.Field
	TransactionLineItemIDs      apijson.Field
	UpdatedAt                   apijson.Field
	VirtualAccountID            apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceJSON) RawJSON() string {
	return r.raw
}

// The counterparty's billing address.
type InvoiceCounterpartyBillingAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	Line1   string `json:"line1,required"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Region or State.
	Region string                                `json:"region,required"`
	Line2  string                                `json:"line2"`
	JSON   invoiceCounterpartyBillingAddressJSON `json:"-"`
}

// invoiceCounterpartyBillingAddressJSON contains the JSON metadata for the struct
// [InvoiceCounterpartyBillingAddress]
type invoiceCounterpartyBillingAddressJSON struct {
	Country     apijson.Field
	Line1       apijson.Field
	Locality    apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	Line2       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCounterpartyBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCounterpartyBillingAddressJSON) RawJSON() string {
	return r.raw
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceCounterpartyShippingAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	Line1   string `json:"line1,required"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Region or State.
	Region string                                 `json:"region,required"`
	Line2  string                                 `json:"line2"`
	JSON   invoiceCounterpartyShippingAddressJSON `json:"-"`
}

// invoiceCounterpartyShippingAddressJSON contains the JSON metadata for the struct
// [InvoiceCounterpartyShippingAddress]
type invoiceCounterpartyShippingAddressJSON struct {
	Country     apijson.Field
	Line1       apijson.Field
	Locality    apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	Line2       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCounterpartyShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceCounterpartyShippingAddressJSON) RawJSON() string {
	return r.raw
}

// The invoice issuer's business address.
type InvoiceInvoicerAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required"`
	Line1   string `json:"line1,required"`
	// Locality or City.
	Locality string `json:"locality,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required"`
	// Region or State.
	Region string                     `json:"region,required"`
	Line2  string                     `json:"line2"`
	JSON   invoiceInvoicerAddressJSON `json:"-"`
}

// invoiceInvoicerAddressJSON contains the JSON metadata for the struct
// [InvoiceInvoicerAddress]
type invoiceInvoicerAddressJSON struct {
	Country     apijson.Field
	Line1       apijson.Field
	Locality    apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	Line2       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceInvoicerAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invoiceInvoicerAddressJSON) RawJSON() string {
	return r.raw
}

// When opening an invoice, whether to show the embedded payment UI , automatically
// debit the recipient, or rely on manual payment from the recipient.
type InvoicePaymentMethod string

const (
	InvoicePaymentMethodUi        InvoicePaymentMethod = "ui"
	InvoicePaymentMethodManual    InvoicePaymentMethod = "manual"
	InvoicePaymentMethodAutomatic InvoicePaymentMethod = "automatic"
)

func (r InvoicePaymentMethod) IsKnown() bool {
	switch r {
	case InvoicePaymentMethodUi, InvoicePaymentMethodManual, InvoicePaymentMethodAutomatic:
		return true
	}
	return false
}

// One of `ach` or `eft`.
type InvoicePaymentType string

const (
	InvoicePaymentTypeEft InvoicePaymentType = "eft"
	InvoicePaymentTypeACH InvoicePaymentType = "ach"
)

func (r InvoicePaymentType) IsKnown() bool {
	switch r {
	case InvoicePaymentTypeEft, InvoicePaymentTypeACH:
		return true
	}
	return false
}

// The status of the invoice.
type InvoiceStatus string

const (
	InvoiceStatusDraft          InvoiceStatus = "draft"
	InvoiceStatusPaid           InvoiceStatus = "paid"
	InvoiceStatusPartiallyPaid  InvoiceStatus = "partially_paid"
	InvoiceStatusPaymentPending InvoiceStatus = "payment_pending"
	InvoiceStatusUnpaid         InvoiceStatus = "unpaid"
	InvoiceStatusVoided         InvoiceStatus = "voided"
)

func (r InvoiceStatus) IsKnown() bool {
	switch r {
	case InvoiceStatusDraft, InvoiceStatusPaid, InvoiceStatusPartiallyPaid, InvoiceStatusPaymentPending, InvoiceStatusUnpaid, InvoiceStatusVoided:
		return true
	}
	return false
}

type InvoiceNewParams struct {
	// The ID of the counterparty receiving the invoice.
	CounterpartyID param.Field[string] `json:"counterparty_id,required"`
	// A future date by when the invoice needs to be paid.
	DueDate param.Field[time.Time] `json:"due_date,required" format:"date-time"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required"`
	// When true, the invoice will progress to unpaid automatically and cannot be
	// edited after entering that state. If the invoice fails to progress to unpaid,
	// the errors will be returned and the invoice will not be created.
	AutoAdvance param.Field[bool] `json:"auto_advance"`
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails param.Field[[]shared.ContactDetailParam] `json:"contact_details"`
	// The counterparty's billing address.
	CounterpartyBillingAddress param.Field[InvoiceNewParamsCounterpartyBillingAddress] `json:"counterparty_billing_address"`
	// The counterparty's shipping address where physical goods should be delivered.
	CounterpartyShippingAddress param.Field[InvoiceNewParamsCounterpartyShippingAddress] `json:"counterparty_shipping_address"`
	// Currency that the invoice is denominated in. Defaults to `USD` if not provided.
	Currency param.Field[shared.Currency] `json:"currency"`
	// A free-form description of the invoice.
	Description param.Field[string] `json:"description"`
	// When payment_method is automatic, the fallback payment method to use when an
	// automatic payment fails. One of `manual` or `ui`.
	FallbackPaymentMethod param.Field[string] `json:"fallback_payment_method"`
	// Whether to ingest the ledger_entries to populate the invoice line items. If this
	// is false, then a line item must be provided. If this is true, line_items must be
	// empty. Ignored if ledger_account_settlement_id is empty.
	IngestLedgerEntries param.Field[bool] `json:"ingest_ledger_entries"`
	// An array of invoice line items. The API supports a maximum of 50 invoice line
	// items per invoice. If a greater number of invoice line items is required, please
	// contact support.
	InvoiceLineItems param.Field[[]InvoiceNewParamsInvoiceLineItem] `json:"invoice_line_items"`
	// The invoice issuer's business address.
	InvoicerAddress param.Field[InvoiceNewParamsInvoicerAddress] `json:"invoicer_address"`
	// The ID of the virtual account the invoice should be paid to.
	LedgerAccountSettlementID param.Field[string] `json:"ledger_account_settlement_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Emails in addition to the counterparty email to send invoice status
	// notifications to. At least one email is required if notifications are enabled
	// and the counterparty doesn't have an email.
	NotificationEmailAddresses param.Field[[]string] `json:"notification_email_addresses"`
	// If true, the invoice will send email notifications to the invoice recipients
	// about invoice status changes.
	NotificationsEnabled param.Field[bool] `json:"notifications_enabled"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	PaymentEffectiveDate param.Field[time.Time] `json:"payment_effective_date" format:"date"`
	// The method by which the invoice can be paid. `ui` will show the embedded payment
	// collection flow. `automatic` will automatically initiate payment based upon the
	// account details of the receiving_account id.\nIf the invoice amount is positive,
	// the automatically initiated payment order's direction will be debit. If the
	// invoice amount is negative, the automatically initiated payment order's
	// direction will be credit. One of `manual`, `ui`, or `automatic`.
	PaymentMethod param.Field[InvoiceNewParamsPaymentMethod] `json:"payment_method"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	PaymentType param.Field[PaymentOrderType] `json:"payment_type"`
	// The receiving account ID. Can be an `external_account`.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// The email of the recipient of the invoice. Leaving this value as null will
	// fallback to using the counterparty's name.
	RecipientEmail param.Field[string] `json:"recipient_email"`
	// The name of the recipient of the invoice. Leaving this value as null will
	// fallback to using the counterparty's name.
	RecipientName param.Field[string] `json:"recipient_name"`
	// Number of days after due date when overdue reminder emails will be sent out to
	// invoice recipients.
	RemindAfterOverdueDays param.Field[[]int64] `json:"remind_after_overdue_days"`
	// The ID of the virtual account the invoice should be paid to.
	VirtualAccountID param.Field[string] `json:"virtual_account_id" format:"uuid"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's billing address.
type InvoiceNewParamsCounterpartyBillingAddress struct {
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

func (r InvoiceNewParamsCounterpartyBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceNewParamsCounterpartyShippingAddress struct {
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

func (r InvoiceNewParamsCounterpartyShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsInvoiceLineItem struct {
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
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity param.Field[int64] `json:"quantity"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit. Accepts decimal strings with
	// up to 12 decimals
	UnitAmountDecimal param.Field[string] `json:"unit_amount_decimal"`
}

func (r InvoiceNewParamsInvoiceLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice issuer's business address.
type InvoiceNewParamsInvoicerAddress struct {
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

func (r InvoiceNewParamsInvoicerAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The method by which the invoice can be paid. `ui` will show the embedded payment
// collection flow. `automatic` will automatically initiate payment based upon the
// account details of the receiving_account id.\nIf the invoice amount is positive,
// the automatically initiated payment order's direction will be debit. If the
// invoice amount is negative, the automatically initiated payment order's
// direction will be credit. One of `manual`, `ui`, or `automatic`.
type InvoiceNewParamsPaymentMethod string

const (
	InvoiceNewParamsPaymentMethodUi        InvoiceNewParamsPaymentMethod = "ui"
	InvoiceNewParamsPaymentMethodManual    InvoiceNewParamsPaymentMethod = "manual"
	InvoiceNewParamsPaymentMethodAutomatic InvoiceNewParamsPaymentMethod = "automatic"
)

func (r InvoiceNewParamsPaymentMethod) IsKnown() bool {
	switch r {
	case InvoiceNewParamsPaymentMethodUi, InvoiceNewParamsPaymentMethodManual, InvoiceNewParamsPaymentMethodAutomatic:
		return true
	}
	return false
}

type InvoiceUpdateParams struct {
	// The invoicer's contact details displayed at the top of the invoice.
	ContactDetails param.Field[[]shared.ContactDetailParam] `json:"contact_details"`
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
	// When payment_method is automatic, the fallback payment method to use when an
	// automatic payment fails. One of `manual` or `ui`.
	FallbackPaymentMethod param.Field[string] `json:"fallback_payment_method"`
	// Whether to ingest the ledger_entries to populate the invoice line items. If this
	// is false, then a line item must be provided. If this is true, line_items must be
	// empty. Ignored if ledger_account_settlement_id is empty.
	IngestLedgerEntries param.Field[bool] `json:"ingest_ledger_entries"`
	// An array of invoice line items. The API supports a maximum of 50 invoice line
	// items per invoice. If a greater number of invoice line items is required, please
	// contact support.
	InvoiceLineItems param.Field[[]InvoiceUpdateParamsInvoiceLineItem] `json:"invoice_line_items"`
	// The invoice issuer's business address.
	InvoicerAddress param.Field[InvoiceUpdateParamsInvoicerAddress] `json:"invoicer_address"`
	// The ID of the virtual account the invoice should be paid to.
	LedgerAccountSettlementID param.Field[string] `json:"ledger_account_settlement_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Emails in addition to the counterparty email to send invoice status
	// notifications to. At least one email is required if notifications are enabled
	// and the counterparty doesn't have an email.
	NotificationEmailAddresses param.Field[[]string] `json:"notification_email_addresses"`
	// If true, the invoice will send email notifications to the invoice recipients
	// about invoice status changes.
	NotificationsEnabled param.Field[bool] `json:"notifications_enabled"`
	// The ID of the internal account the invoice should be paid to.
	OriginatingAccountID param.Field[string] `json:"originating_account_id"`
	// Date transactions are to be posted to the participants' account. Defaults to the
	// current business day or the next business day if the current day is a bank
	// holiday or weekend. Format: yyyy-mm-dd.
	PaymentEffectiveDate param.Field[time.Time] `json:"payment_effective_date" format:"date"`
	// The method by which the invoice can be paid. `ui` will show the embedded payment
	// collection flow. `automatic` will automatically initiate payment based upon the
	// account details of the receiving_account id.\nIf the invoice amount is positive,
	// the automatically initiated payment order's direction will be debit. If the
	// invoice amount is negative, the automatically initiated payment order's
	// direction will be credit. One of `manual`, `ui`, or `automatic`.
	PaymentMethod param.Field[InvoiceUpdateParamsPaymentMethod] `json:"payment_method"`
	// One of `ach`, `se_bankgirot`, `eft`, `wire`, `check`, `sen`, `book`, `rtp`,
	// `sepa`, `bacs`, `au_becs`, `interac`, `neft`, `nics`,
	// `nz_national_clearing_code`, `sic`, `signet`, `provexchange`, `zengin`.
	PaymentType param.Field[PaymentOrderType] `json:"payment_type"`
	// The receiving account ID. Can be an `external_account`.
	ReceivingAccountID param.Field[string] `json:"receiving_account_id" format:"uuid"`
	// The email of the recipient of the invoice. Leaving this value as null will
	// fallback to using the counterparty's name.
	RecipientEmail param.Field[string] `json:"recipient_email"`
	// The name of the recipient of the invoice. Leaving this value as null will
	// fallback to using the counterparty's name.
	RecipientName param.Field[string] `json:"recipient_name"`
	// Number of days after due date when overdue reminder emails will be sent out to
	// invoice recipients.
	RemindAfterOverdueDays param.Field[[]int64] `json:"remind_after_overdue_days"`
	// Invoice status must be updated in a `PATCH` request that does not modify any
	// other invoice attributes. Valid state transitions are `draft` to `unpaid`,
	// `draft` or `unpaid` to `voided`, and `draft` or `unpaid` to `paid`.
	Status param.Field[string] `json:"status"`
	// The ID of the virtual account the invoice should be paid to.
	VirtualAccountID param.Field[string] `json:"virtual_account_id" format:"uuid"`
}

func (r InvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's billing address.
type InvoiceUpdateParamsCounterpartyBillingAddress struct {
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

func (r InvoiceUpdateParamsCounterpartyBillingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The counterparty's shipping address where physical goods should be delivered.
type InvoiceUpdateParamsCounterpartyShippingAddress struct {
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

func (r InvoiceUpdateParamsCounterpartyShippingAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceUpdateParamsInvoiceLineItem struct {
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
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The number of units of a product or service that this line item is for. Must be
	// a whole number. Defaults to 1 if not provided.
	Quantity param.Field[int64] `json:"quantity"`
	// The cost per unit of the product or service that this line item is for,
	// specified in the invoice currency's smallest unit. Accepts decimal strings with
	// up to 12 decimals
	UnitAmountDecimal param.Field[string] `json:"unit_amount_decimal"`
}

func (r InvoiceUpdateParamsInvoiceLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The invoice issuer's business address.
type InvoiceUpdateParamsInvoicerAddress struct {
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

func (r InvoiceUpdateParamsInvoicerAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The method by which the invoice can be paid. `ui` will show the embedded payment
// collection flow. `automatic` will automatically initiate payment based upon the
// account details of the receiving_account id.\nIf the invoice amount is positive,
// the automatically initiated payment order's direction will be debit. If the
// invoice amount is negative, the automatically initiated payment order's
// direction will be credit. One of `manual`, `ui`, or `automatic`.
type InvoiceUpdateParamsPaymentMethod string

const (
	InvoiceUpdateParamsPaymentMethodUi        InvoiceUpdateParamsPaymentMethod = "ui"
	InvoiceUpdateParamsPaymentMethodManual    InvoiceUpdateParamsPaymentMethod = "manual"
	InvoiceUpdateParamsPaymentMethodAutomatic InvoiceUpdateParamsPaymentMethod = "automatic"
)

func (r InvoiceUpdateParamsPaymentMethod) IsKnown() bool {
	switch r {
	case InvoiceUpdateParamsPaymentMethodUi, InvoiceUpdateParamsPaymentMethodManual, InvoiceUpdateParamsPaymentMethodAutomatic:
		return true
	}
	return false
}

type InvoiceListParams struct {
	AfterCursor    param.Field[string] `query:"after_cursor"`
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// An inclusive upper bound for searching due_date
	DueDateEnd param.Field[time.Time] `query:"due_date_end" format:"date"`
	// An inclusive lower bound for searching due_date
	DueDateStart      param.Field[time.Time] `query:"due_date_start" format:"date"`
	ExpectedPaymentID param.Field[string]    `query:"expected_payment_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// A unique record number assigned to each invoice that is issued.
	Number               param.Field[string]                  `query:"number"`
	OriginatingAccountID param.Field[string]                  `query:"originating_account_id"`
	PaymentOrderID       param.Field[string]                  `query:"payment_order_id"`
	PerPage              param.Field[int64]                   `query:"per_page"`
	Status               param.Field[InvoiceListParamsStatus] `query:"status"`
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceListParamsStatus string

const (
	InvoiceListParamsStatusDraft          InvoiceListParamsStatus = "draft"
	InvoiceListParamsStatusPaid           InvoiceListParamsStatus = "paid"
	InvoiceListParamsStatusPartiallyPaid  InvoiceListParamsStatus = "partially_paid"
	InvoiceListParamsStatusPaymentPending InvoiceListParamsStatus = "payment_pending"
	InvoiceListParamsStatusUnpaid         InvoiceListParamsStatus = "unpaid"
	InvoiceListParamsStatusVoided         InvoiceListParamsStatus = "voided"
)

func (r InvoiceListParamsStatus) IsKnown() bool {
	switch r {
	case InvoiceListParamsStatusDraft, InvoiceListParamsStatusPaid, InvoiceListParamsStatusPartiallyPaid, InvoiceListParamsStatusPaymentPending, InvoiceListParamsStatusUnpaid, InvoiceListParamsStatusVoided:
		return true
	}
	return false
}
