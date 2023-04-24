package moderntreasury

import (
	"context"
	"os"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
	"github.com/Modern-Treasury/modern-treasury-go/services"
)

type ModernTreasury struct {
	Options                 []option.RequestOption
	Connections             *services.ConnectionService
	Counterparties          *services.CounterpartyService
	Events                  *services.EventService
	ExpectedPayments        *services.ExpectedPaymentService
	ExternalAccounts        *services.ExternalAccountService
	IncomingPaymentDetails  *services.IncomingPaymentDetailService
	Invoices                *services.InvoiceService
	Documents               *services.DocumentService
	AccountCollectionFlows  *services.AccountCollectionFlowService
	AccountDetails          *services.AccountDetailService
	RoutingDetails          *services.RoutingDetailService
	InternalAccounts        *services.InternalAccountService
	Ledgers                 *services.LedgerService
	LedgerAccountCategories *services.LedgerAccountCategoryService
	LedgerAccounts          *services.LedgerAccountService
	LedgerAccountPayouts    *services.LedgerAccountPayoutService
	LedgerEntries           *services.LedgerEntryService
	LedgerTransactions      *services.LedgerTransactionService
	LineItems               *services.LineItemService
	PaymentFlows            *services.PaymentFlowService
	PaymentOrders           *services.PaymentOrderService
	PaymentReferences       *services.PaymentReferenceService
	Returns                 *services.ReturnService
	Transactions            *services.TransactionService
	Validations             *services.ValidationService
	PaperItems              *services.PaperItemService
	Webhooks                *services.WebhookService
	VirtualAccounts         *services.VirtualAccountService
}

// NewModernTreasury generates a new client with the default option read from the
// environment ("MODERN_TREASURY_API_KEY", "MODERN_TREASURY_ORGANIZATION_ID",
// "MODERN_TREASURY_WEBHOOK_KEY"). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewModernTreasury(opts ...option.RequestOption) (r *ModernTreasury) {
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

	r = &ModernTreasury{Options: opts}

	r.Connections = services.NewConnectionService(opts...)
	r.Counterparties = services.NewCounterpartyService(opts...)
	r.Events = services.NewEventService(opts...)
	r.ExpectedPayments = services.NewExpectedPaymentService(opts...)
	r.ExternalAccounts = services.NewExternalAccountService(opts...)
	r.IncomingPaymentDetails = services.NewIncomingPaymentDetailService(opts...)
	r.Invoices = services.NewInvoiceService(opts...)
	r.Documents = services.NewDocumentService(opts...)
	r.AccountCollectionFlows = services.NewAccountCollectionFlowService(opts...)
	r.AccountDetails = services.NewAccountDetailService(opts...)
	r.RoutingDetails = services.NewRoutingDetailService(opts...)
	r.InternalAccounts = services.NewInternalAccountService(opts...)
	r.Ledgers = services.NewLedgerService(opts...)
	r.LedgerAccountCategories = services.NewLedgerAccountCategoryService(opts...)
	r.LedgerAccounts = services.NewLedgerAccountService(opts...)
	r.LedgerAccountPayouts = services.NewLedgerAccountPayoutService(opts...)
	r.LedgerEntries = services.NewLedgerEntryService(opts...)
	r.LedgerTransactions = services.NewLedgerTransactionService(opts...)
	r.LineItems = services.NewLineItemService(opts...)
	r.PaymentFlows = services.NewPaymentFlowService(opts...)
	r.PaymentOrders = services.NewPaymentOrderService(opts...)
	r.PaymentReferences = services.NewPaymentReferenceService(opts...)
	r.Returns = services.NewReturnService(opts...)
	r.Transactions = services.NewTransactionService(opts...)
	r.Validations = services.NewValidationService(opts...)
	r.PaperItems = services.NewPaperItemService(opts...)
	r.Webhooks = services.NewWebhookService(opts...)
	r.VirtualAccounts = services.NewVirtualAccountService(opts...)

	return
}

// A test endpoint often used to confirm credentials and headers are being passed
// in correctly.
func (r *ModernTreasury) Ping(ctx context.Context, opts ...option.RequestOption) (res *responses.PingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ping"
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
