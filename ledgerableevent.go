// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// LedgerableEventService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLedgerableEventService] method instead.
type LedgerableEventService struct {
	Options []option.RequestOption
}

// NewLedgerableEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLedgerableEventService(opts ...option.RequestOption) (r *LedgerableEventService) {
	r = &LedgerableEventService{}
	r.Options = opts
	return
}

// Translation missing:
// en.openapi.descriptions.ledger.operations.create_ledgerable_event
func (r *LedgerableEventService) New(ctx context.Context, body LedgerableEventNewParams, opts ...option.RequestOption) (res *LedgerableEvent, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledgerable_events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledgerable event.
func (r *LedgerableEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerableEvent, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledgerable_events/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LedgerableEvent struct {
	ID string `json:"id,required" format:"uuid"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount    int64     `json:"amount,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// An ISO 4217 conformed currency or a custom currency.
	Currency string `json:"currency,required"`
	// Must be included if currency is a custom currency. The currency_exponent cannot
	// exceed 30.
	CurrencyExponent int64 `json:"currency_exponent,required,nullable"`
	// Additionally data to be used by the Ledger Event Handler.
	CustomData interface{} `json:"custom_data,required,nullable"`
	// Description of the ledgerable event.
	Description string `json:"description,required,nullable"`
	// One of `credit`, `debit`.
	Direction string `json:"direction,required,nullable"`
	// Id of the ledger event handler that is used to create a ledger transaction.
	LedgerEventHandlerID string `json:"ledger_event_handler_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required,nullable"`
	// Name of the ledgerable event.
	Name      string    `json:"name,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	JSON      ledgerableEventJSON
}

// ledgerableEventJSON contains the JSON metadata for the struct [LedgerableEvent]
type ledgerableEventJSON struct {
	ID                   apijson.Field
	Amount               apijson.Field
	CreatedAt            apijson.Field
	Currency             apijson.Field
	CurrencyExponent     apijson.Field
	CustomData           apijson.Field
	Description          apijson.Field
	Direction            apijson.Field
	LedgerEventHandlerID apijson.Field
	LiveMode             apijson.Field
	Metadata             apijson.Field
	Name                 apijson.Field
	Object               apijson.Field
	UpdatedAt            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LedgerableEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerableEventNewParams struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000.
	Amount param.Field[int64] `json:"amount,required"`
	// Name of the ledgerable event.
	Name param.Field[string] `json:"name,required"`
	// An ISO 4217 conformed currency or a custom currency.
	Currency param.Field[string] `json:"currency"`
	// Must be included if currency is a custom currency. The currency_exponent cannot
	// exceed 30.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// Additionally data to be used by the Ledger Event Handler.
	CustomData param.Field[interface{}] `json:"custom_data"`
	// Description of the ledgerable event.
	Description param.Field[string] `json:"description"`
	// One of `credit`, `debit`.
	Direction param.Field[string] `json:"direction"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerableEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
