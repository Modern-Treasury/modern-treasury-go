// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
	LedgerAccountCategories      *LedgerAccountCategoryService
	LedgerAccounts               *LedgerAccountService
	LedgerAccountBalanceMonitors *LedgerAccountBalanceMonitorService
	LedgerAccountStatements      *LedgerAccountStatementService
	LedgerEntries                *LedgerEntryService
	LedgerTransactions           *LedgerTransactionService
	LineItems                    *LineItemService
	PaymentFlows                 *PaymentFlowService
	PaymentOrders                *PaymentOrderService
	PaymentReferences            *PaymentReferenceService
	Returns                      *ReturnService
	Transactions                 *TransactionService
	Validations                  *ValidationService
	VirtualAccounts              *VirtualAccountService
	BulkRequests                 *BulkRequestService
	BulkResults                  *BulkResultService
	LedgerAccountSettlements     *LedgerAccountSettlementService
	ForeignExchangeQuotes        *ForeignExchangeQuoteService
	ConnectionLegalEntities      *ConnectionLegalEntityService
	LegalEntities                *LegalEntityService
	LegalEntityAssociations      *LegalEntityAssociationService
	PaymentActions               *PaymentActionService
}

// DefaultClientOptions read from the environment (MODERN_TREASURY_API_KEY,
// MODERN_TREASURY_WEBHOOK_KEY, MODERN_TREASURY_ORGANIZATION_ID,
// MODERN_TREASURY_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("MODERN_TREASURY_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("MODERN_TREASURY_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("MODERN_TREASURY_ORGANIZATION_ID"); ok {
		defaults = append(defaults, option.WithOrganizationID(o))
	}
	if o, ok := os.LookupEnv("MODERN_TREASURY_WEBHOOK_KEY"); ok {
		defaults = append(defaults, option.WithWebhookKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (MODERN_TREASURY_API_KEY, MODERN_TREASURY_WEBHOOK_KEY,
// MODERN_TREASURY_ORGANIZATION_ID, MODERN_TREASURY_BASE_URL). The option passed in
// as arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

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
	r.LedgerAccountCategories = NewLedgerAccountCategoryService(opts...)
	r.LedgerAccounts = NewLedgerAccountService(opts...)
	r.LedgerAccountBalanceMonitors = NewLedgerAccountBalanceMonitorService(opts...)
	r.LedgerAccountStatements = NewLedgerAccountStatementService(opts...)
	r.LedgerEntries = NewLedgerEntryService(opts...)
	r.LedgerTransactions = NewLedgerTransactionService(opts...)
	r.LineItems = NewLineItemService(opts...)
	r.PaymentFlows = NewPaymentFlowService(opts...)
	r.PaymentOrders = NewPaymentOrderService(opts...)
	r.PaymentReferences = NewPaymentReferenceService(opts...)
	r.Returns = NewReturnService(opts...)
	r.Transactions = NewTransactionService(opts...)
	r.Validations = NewValidationService(opts...)
	r.VirtualAccounts = NewVirtualAccountService(opts...)
	r.BulkRequests = NewBulkRequestService(opts...)
	r.BulkResults = NewBulkResultService(opts...)
	r.LedgerAccountSettlements = NewLedgerAccountSettlementService(opts...)
	r.ForeignExchangeQuotes = NewForeignExchangeQuoteService(opts...)
	r.ConnectionLegalEntities = NewConnectionLegalEntityService(opts...)
	r.LegalEntities = NewLegalEntityService(opts...)
	r.LegalEntityAssociations = NewLegalEntityAssociationService(opts...)
	r.PaymentActions = NewPaymentActionService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}

// A test endpoint often used to confirm credentials and headers are being passed
// in correctly.
func (r *Client) Ping(ctx context.Context, opts ...option.RequestOption) (res *PingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
