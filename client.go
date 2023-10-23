// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"net/http"
	"os"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the Modern Treasury API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                      []option.RequestOption
	Connections                  *ConnectionService
	Counterparties               *CounterpartyService
	Events                       *EventService
	ExpectedPayments             *ExpectedPaymentService
	ExternalAccounts             *ExternalAccountService
	IncomingPaymentDetails       *IncomingPaymentDetailService
	Invoices                     *InvoiceService
	Documents                    *DocumentService
	AccountCollectionFlows       *AccountCollectionFlowService
	AccountDetails               *AccountDetailService
	RoutingDetails               *RoutingDetailService
	InternalAccounts             *InternalAccountService
	Ledgers                      *LedgerService
	LedgerableEvents             *LedgerableEventService
	LedgerAccountCategories      *LedgerAccountCategoryService
	LedgerAccounts               *LedgerAccountService
	LedgerAccountBalanceMonitors *LedgerAccountBalanceMonitorService
	LedgerAccountPayouts         *LedgerAccountPayoutService
	LedgerAccountStatements      *LedgerAccountStatementService
	LedgerEntries                *LedgerEntryService
	LedgerEventHandlers          *LedgerEventHandlerService
	LedgerTransactions           *LedgerTransactionService
	LineItems                    *LineItemService
	PaymentFlows                 *PaymentFlowService
	PaymentOrders                *PaymentOrderService
	PaymentReferences            *PaymentReferenceService
	Returns                      *ReturnService
	Transactions                 *TransactionService
	Validations                  *ValidationService
	PaperItems                   *PaperItemService
	Webhooks                     *WebhookService
	VirtualAccounts              *VirtualAccountService
}

// NewClient generates a new client with the default option read from the
// environment (MODERN_TREASURY_API_KEY, MODERN_TREASURY_ORGANIZATION_ID,
// MODERN_TREASURY_WEBHOOK_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("MODERN_TREASURY_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("MODERN_TREASURY_ORGANIZATION_ID"); ok {
		defaults = append(defaults, option.WithOrganizationID(o))
	}
	if o, ok := os.LookupEnv("MODERN_TREASURY_WEBHOOK_KEY"); ok {
		defaults = append(defaults, option.WithWebhookKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Connections = NewConnectionService(opts...)
	r.Counterparties = NewCounterpartyService(opts...)
	r.Events = NewEventService(opts...)
	r.ExpectedPayments = NewExpectedPaymentService(opts...)
	r.ExternalAccounts = NewExternalAccountService(opts...)
	r.IncomingPaymentDetails = NewIncomingPaymentDetailService(opts...)
	r.Invoices = NewInvoiceService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.AccountCollectionFlows = NewAccountCollectionFlowService(opts...)
	r.AccountDetails = NewAccountDetailService(opts...)
	r.RoutingDetails = NewRoutingDetailService(opts...)
	r.InternalAccounts = NewInternalAccountService(opts...)
	r.Ledgers = NewLedgerService(opts...)
	r.LedgerableEvents = NewLedgerableEventService(opts...)
	r.LedgerAccountCategories = NewLedgerAccountCategoryService(opts...)
	r.LedgerAccounts = NewLedgerAccountService(opts...)
	r.LedgerAccountBalanceMonitors = NewLedgerAccountBalanceMonitorService(opts...)
	r.LedgerAccountPayouts = NewLedgerAccountPayoutService(opts...)
	r.LedgerAccountStatements = NewLedgerAccountStatementService(opts...)
	r.LedgerEntries = NewLedgerEntryService(opts...)
	r.LedgerEventHandlers = NewLedgerEventHandlerService(opts...)
	r.LedgerTransactions = NewLedgerTransactionService(opts...)
	r.LineItems = NewLineItemService(opts...)
	r.PaymentFlows = NewPaymentFlowService(opts...)
	r.PaymentOrders = NewPaymentOrderService(opts...)
	r.PaymentReferences = NewPaymentReferenceService(opts...)
	r.Returns = NewReturnService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.Validations = NewValidationService(opts...)
	r.PaperItems = NewPaperItemService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.VirtualAccounts = NewVirtualAccountService(opts...)

	return
}

// A test endpoint often used to confirm credentials and headers are being passed
// in correctly.
func (r *Client) Ping(ctx context.Context, opts ...option.RequestOption) (res *PingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
