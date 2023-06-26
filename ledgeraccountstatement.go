// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// LedgerAccountStatementService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLedgerAccountStatementService] method instead.
type LedgerAccountStatementService struct {
	Options []option.RequestOption
}

// NewLedgerAccountStatementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLedgerAccountStatementService(opts ...option.RequestOption) (r *LedgerAccountStatementService) {
	r = &LedgerAccountStatementService{}
	r.Options = opts
	return
}

// Create a ledger account statement.
func (r *LedgerAccountStatementService) New(ctx context.Context, params LedgerAccountStatementNewParams, opts ...option.RequestOption) (res *LedgerAccountStatementCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_statements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details on a single ledger account statement.
func (r *LedgerAccountStatementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountStatementCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_statements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LedgerAccountStatementCreateResponse struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The id of the ledger that this ledger account statement belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// The description of the ledger account statement.
	Description string `json:"description,required,nullable"`
	// The id of the ledger account whose ledger entries are queried against, and its
	// balances are computed as a result.
	LedgerAccountID string `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account at the time of statement generation.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required"`
	// The normal balance of the ledger account.
	LedgerAccountNormalBalance LedgerAccountStatementCreateResponseLedgerAccountNormalBalance `json:"ledger_account_normal_balance,required"`
	// The inclusive lower bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtLowerBound string `json:"effective_at_lower_bound,required" format:"time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound string `json:"effective_at_upper_bound,required" format:"time"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_lower_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	StartingBalance LedgerAccountStatementCreateResponseStartingBalance `json:"starting_balance,required"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	EndingBalance LedgerAccountStatementCreateResponseEndingBalance `json:"ending_balance,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     ledgerAccountStatementCreateResponseJSON
}

// ledgerAccountStatementCreateResponseJSON contains the JSON metadata for the
// struct [LedgerAccountStatementCreateResponse]
type ledgerAccountStatementCreateResponseJSON struct {
	ID                         apijson.Field
	Object                     apijson.Field
	LiveMode                   apijson.Field
	CreatedAt                  apijson.Field
	UpdatedAt                  apijson.Field
	LedgerID                   apijson.Field
	Description                apijson.Field
	LedgerAccountID            apijson.Field
	LedgerAccountLockVersion   apijson.Field
	LedgerAccountNormalBalance apijson.Field
	EffectiveAtLowerBound      apijson.Field
	EffectiveAtUpperBound      apijson.Field
	StartingBalance            apijson.Field
	EndingBalance              apijson.Field
	Metadata                   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The normal balance of the ledger account.
type LedgerAccountStatementCreateResponseLedgerAccountNormalBalance string

const (
	LedgerAccountStatementCreateResponseLedgerAccountNormalBalanceCredit LedgerAccountStatementCreateResponseLedgerAccountNormalBalance = "credit"
	LedgerAccountStatementCreateResponseLedgerAccountNormalBalanceDebit  LedgerAccountStatementCreateResponseLedgerAccountNormalBalance = "debit"
)

// The pending, posted, and available balances for this ledger account at the
// `effective_at_lower_bound`. The posted balance is the sum of all posted entries
// on the account. The pending balance is the sum of all pending and posted entries
// on the account. The available balance is the posted incoming entries minus the
// sum of the pending and posted outgoing amounts.
type LedgerAccountStatementCreateResponseStartingBalance struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountStatementCreateResponseStartingBalancePendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountStatementCreateResponseStartingBalancePostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountStatementCreateResponseStartingBalanceAvailableBalance `json:"available_balance,required"`
	JSON             ledgerAccountStatementCreateResponseStartingBalanceJSON
}

// ledgerAccountStatementCreateResponseStartingBalanceJSON contains the JSON
// metadata for the struct [LedgerAccountStatementCreateResponseStartingBalance]
type ledgerAccountStatementCreateResponseStartingBalanceJSON struct {
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	AvailableBalance apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseStartingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountStatementCreateResponseStartingBalancePendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountStatementCreateResponseStartingBalancePendingBalanceJSON
}

// ledgerAccountStatementCreateResponseStartingBalancePendingBalanceJSON contains
// the JSON metadata for the struct
// [LedgerAccountStatementCreateResponseStartingBalancePendingBalance]
type ledgerAccountStatementCreateResponseStartingBalancePendingBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseStartingBalancePendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountStatementCreateResponseStartingBalancePostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountStatementCreateResponseStartingBalancePostedBalanceJSON
}

// ledgerAccountStatementCreateResponseStartingBalancePostedBalanceJSON contains
// the JSON metadata for the struct
// [LedgerAccountStatementCreateResponseStartingBalancePostedBalance]
type ledgerAccountStatementCreateResponseStartingBalancePostedBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseStartingBalancePostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountStatementCreateResponseStartingBalanceAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountStatementCreateResponseStartingBalanceAvailableBalanceJSON
}

// ledgerAccountStatementCreateResponseStartingBalanceAvailableBalanceJSON contains
// the JSON metadata for the struct
// [LedgerAccountStatementCreateResponseStartingBalanceAvailableBalance]
type ledgerAccountStatementCreateResponseStartingBalanceAvailableBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseStartingBalanceAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account at the
// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
// on the account. The pending balance is the sum of all pending and posted entries
// on the account. The available balance is the posted incoming entries minus the
// sum of the pending and posted outgoing amounts.
type LedgerAccountStatementCreateResponseEndingBalance struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountStatementCreateResponseEndingBalancePendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountStatementCreateResponseEndingBalancePostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountStatementCreateResponseEndingBalanceAvailableBalance `json:"available_balance,required"`
	JSON             ledgerAccountStatementCreateResponseEndingBalanceJSON
}

// ledgerAccountStatementCreateResponseEndingBalanceJSON contains the JSON metadata
// for the struct [LedgerAccountStatementCreateResponseEndingBalance]
type ledgerAccountStatementCreateResponseEndingBalanceJSON struct {
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	AvailableBalance apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseEndingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountStatementCreateResponseEndingBalancePendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountStatementCreateResponseEndingBalancePendingBalanceJSON
}

// ledgerAccountStatementCreateResponseEndingBalancePendingBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementCreateResponseEndingBalancePendingBalance]
type ledgerAccountStatementCreateResponseEndingBalancePendingBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseEndingBalancePendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountStatementCreateResponseEndingBalancePostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountStatementCreateResponseEndingBalancePostedBalanceJSON
}

// ledgerAccountStatementCreateResponseEndingBalancePostedBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementCreateResponseEndingBalancePostedBalance]
type ledgerAccountStatementCreateResponseEndingBalancePostedBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseEndingBalancePostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountStatementCreateResponseEndingBalanceAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerAccountStatementCreateResponseEndingBalanceAvailableBalanceJSON
}

// ledgerAccountStatementCreateResponseEndingBalanceAvailableBalanceJSON contains
// the JSON metadata for the struct
// [LedgerAccountStatementCreateResponseEndingBalanceAvailableBalance]
type ledgerAccountStatementCreateResponseEndingBalanceAvailableBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementCreateResponseEndingBalanceAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountStatementNewParams struct {
	// The inclusive lower bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtLowerBound param.Field[string] `json:"effective_at_lower_bound,required" format:"time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound param.Field[string] `json:"effective_at_upper_bound,required" format:"time"`
	// The id of the ledger account whose ledger entries are queried against, and its
	// balances are computed as a result.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// The description of the ledger account statement.
	Description param.Field[string] `json:"description"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata       param.Field[map[string]string] `json:"metadata"`
	IdempotencyKey param.Field[string]            `header:"Idempotency-Key"`
}

func (r LedgerAccountStatementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
