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
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
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
func (r *LedgerAccountStatementService) New(ctx context.Context, body LedgerAccountStatementNewParams, opts ...option.RequestOption) (res *LedgerAccountStatementNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_statements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger account statement.
func (r *LedgerAccountStatementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountStatementGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_statements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LedgerAccountStatementNewResponse struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the ledger account statement.
	Description string `json:"description,required,nullable"`
	// The inclusive lower bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtLowerBound string `json:"effective_at_lower_bound,required" format:"time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound string `json:"effective_at_upper_bound,required" format:"time"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	EndingBalance LedgerAccountStatementNewResponseEndingBalance `json:"ending_balance,required"`
	// The id of the ledger account whose ledger entries are queried against, and its
	// balances are computed as a result.
	LedgerAccountID string `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account at the time of statement generation.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required"`
	// The normal balance of the ledger account.
	LedgerAccountNormalBalance shared.TransactionDirection `json:"ledger_account_normal_balance,required"`
	// The id of the ledger that this ledger account statement belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_lower_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	StartingBalance LedgerAccountStatementNewResponseStartingBalance `json:"starting_balance,required"`
	UpdatedAt       time.Time                                        `json:"updated_at,required" format:"date-time"`
	JSON            ledgerAccountStatementNewResponseJSON
}

// ledgerAccountStatementNewResponseJSON contains the JSON metadata for the struct
// [LedgerAccountStatementNewResponse]
type ledgerAccountStatementNewResponseJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	Description                apijson.Field
	EffectiveAtLowerBound      apijson.Field
	EffectiveAtUpperBound      apijson.Field
	EndingBalance              apijson.Field
	LedgerAccountID            apijson.Field
	LedgerAccountLockVersion   apijson.Field
	LedgerAccountNormalBalance apijson.Field
	LedgerID                   apijson.Field
	LiveMode                   apijson.Field
	Metadata                   apijson.Field
	Object                     apijson.Field
	StartingBalance            apijson.Field
	UpdatedAt                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account at the
// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
// on the account. The pending balance is the sum of all pending and posted entries
// on the account. The available balance is the posted incoming entries minus the
// sum of the pending and posted outgoing amounts.
type LedgerAccountStatementNewResponseEndingBalance struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountStatementNewResponseEndingBalanceAvailableBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountStatementNewResponseEndingBalancePendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountStatementNewResponseEndingBalancePostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountStatementNewResponseEndingBalanceJSON
}

// ledgerAccountStatementNewResponseEndingBalanceJSON contains the JSON metadata
// for the struct [LedgerAccountStatementNewResponseEndingBalance]
type ledgerAccountStatementNewResponseEndingBalanceJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseEndingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountStatementNewResponseEndingBalanceAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementNewResponseEndingBalanceAvailableBalanceJSON
}

// ledgerAccountStatementNewResponseEndingBalanceAvailableBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementNewResponseEndingBalanceAvailableBalance]
type ledgerAccountStatementNewResponseEndingBalanceAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseEndingBalanceAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountStatementNewResponseEndingBalancePendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementNewResponseEndingBalancePendingBalanceJSON
}

// ledgerAccountStatementNewResponseEndingBalancePendingBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementNewResponseEndingBalancePendingBalance]
type ledgerAccountStatementNewResponseEndingBalancePendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseEndingBalancePendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountStatementNewResponseEndingBalancePostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementNewResponseEndingBalancePostedBalanceJSON
}

// ledgerAccountStatementNewResponseEndingBalancePostedBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementNewResponseEndingBalancePostedBalance]
type ledgerAccountStatementNewResponseEndingBalancePostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseEndingBalancePostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account at the
// `effective_at_lower_bound`. The posted balance is the sum of all posted entries
// on the account. The pending balance is the sum of all pending and posted entries
// on the account. The available balance is the posted incoming entries minus the
// sum of the pending and posted outgoing amounts.
type LedgerAccountStatementNewResponseStartingBalance struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountStatementNewResponseStartingBalanceAvailableBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountStatementNewResponseStartingBalancePendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountStatementNewResponseStartingBalancePostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountStatementNewResponseStartingBalanceJSON
}

// ledgerAccountStatementNewResponseStartingBalanceJSON contains the JSON metadata
// for the struct [LedgerAccountStatementNewResponseStartingBalance]
type ledgerAccountStatementNewResponseStartingBalanceJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseStartingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountStatementNewResponseStartingBalanceAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementNewResponseStartingBalanceAvailableBalanceJSON
}

// ledgerAccountStatementNewResponseStartingBalanceAvailableBalanceJSON contains
// the JSON metadata for the struct
// [LedgerAccountStatementNewResponseStartingBalanceAvailableBalance]
type ledgerAccountStatementNewResponseStartingBalanceAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseStartingBalanceAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountStatementNewResponseStartingBalancePendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementNewResponseStartingBalancePendingBalanceJSON
}

// ledgerAccountStatementNewResponseStartingBalancePendingBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementNewResponseStartingBalancePendingBalance]
type ledgerAccountStatementNewResponseStartingBalancePendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseStartingBalancePendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountStatementNewResponseStartingBalancePostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementNewResponseStartingBalancePostedBalanceJSON
}

// ledgerAccountStatementNewResponseStartingBalancePostedBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementNewResponseStartingBalancePostedBalance]
type ledgerAccountStatementNewResponseStartingBalancePostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementNewResponseStartingBalancePostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerAccountStatementGetResponse struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the ledger account statement.
	Description string `json:"description,required,nullable"`
	// The inclusive lower bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtLowerBound string `json:"effective_at_lower_bound,required" format:"time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound string `json:"effective_at_upper_bound,required" format:"time"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	EndingBalance LedgerAccountStatementGetResponseEndingBalance `json:"ending_balance,required"`
	// The id of the ledger account whose ledger entries are queried against, and its
	// balances are computed as a result.
	LedgerAccountID string `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account at the time of statement generation.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required"`
	// The normal balance of the ledger account.
	LedgerAccountNormalBalance shared.TransactionDirection `json:"ledger_account_normal_balance,required"`
	// The id of the ledger that this ledger account statement belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	Object   string            `json:"object,required"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_lower_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	StartingBalance LedgerAccountStatementGetResponseStartingBalance `json:"starting_balance,required"`
	UpdatedAt       time.Time                                        `json:"updated_at,required" format:"date-time"`
	JSON            ledgerAccountStatementGetResponseJSON
}

// ledgerAccountStatementGetResponseJSON contains the JSON metadata for the struct
// [LedgerAccountStatementGetResponse]
type ledgerAccountStatementGetResponseJSON struct {
	ID                         apijson.Field
	CreatedAt                  apijson.Field
	Description                apijson.Field
	EffectiveAtLowerBound      apijson.Field
	EffectiveAtUpperBound      apijson.Field
	EndingBalance              apijson.Field
	LedgerAccountID            apijson.Field
	LedgerAccountLockVersion   apijson.Field
	LedgerAccountNormalBalance apijson.Field
	LedgerID                   apijson.Field
	LiveMode                   apijson.Field
	Metadata                   apijson.Field
	Object                     apijson.Field
	StartingBalance            apijson.Field
	UpdatedAt                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account at the
// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
// on the account. The pending balance is the sum of all pending and posted entries
// on the account. The available balance is the posted incoming entries minus the
// sum of the pending and posted outgoing amounts.
type LedgerAccountStatementGetResponseEndingBalance struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountStatementGetResponseEndingBalanceAvailableBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountStatementGetResponseEndingBalancePendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountStatementGetResponseEndingBalancePostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountStatementGetResponseEndingBalanceJSON
}

// ledgerAccountStatementGetResponseEndingBalanceJSON contains the JSON metadata
// for the struct [LedgerAccountStatementGetResponseEndingBalance]
type ledgerAccountStatementGetResponseEndingBalanceJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseEndingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountStatementGetResponseEndingBalanceAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementGetResponseEndingBalanceAvailableBalanceJSON
}

// ledgerAccountStatementGetResponseEndingBalanceAvailableBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementGetResponseEndingBalanceAvailableBalance]
type ledgerAccountStatementGetResponseEndingBalanceAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseEndingBalanceAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountStatementGetResponseEndingBalancePendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementGetResponseEndingBalancePendingBalanceJSON
}

// ledgerAccountStatementGetResponseEndingBalancePendingBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementGetResponseEndingBalancePendingBalance]
type ledgerAccountStatementGetResponseEndingBalancePendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseEndingBalancePendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountStatementGetResponseEndingBalancePostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementGetResponseEndingBalancePostedBalanceJSON
}

// ledgerAccountStatementGetResponseEndingBalancePostedBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementGetResponseEndingBalancePostedBalance]
type ledgerAccountStatementGetResponseEndingBalancePostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseEndingBalancePostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending, posted, and available balances for this ledger account at the
// `effective_at_lower_bound`. The posted balance is the sum of all posted entries
// on the account. The pending balance is the sum of all pending and posted entries
// on the account. The available balance is the posted incoming entries minus the
// sum of the pending and posted outgoing amounts.
type LedgerAccountStatementGetResponseStartingBalance struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerAccountStatementGetResponseStartingBalanceAvailableBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerAccountStatementGetResponseStartingBalancePendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerAccountStatementGetResponseStartingBalancePostedBalance `json:"posted_balance,required"`
	JSON          ledgerAccountStatementGetResponseStartingBalanceJSON
}

// ledgerAccountStatementGetResponseStartingBalanceJSON contains the JSON metadata
// for the struct [LedgerAccountStatementGetResponseStartingBalance]
type ledgerAccountStatementGetResponseStartingBalanceJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseStartingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerAccountStatementGetResponseStartingBalanceAvailableBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementGetResponseStartingBalanceAvailableBalanceJSON
}

// ledgerAccountStatementGetResponseStartingBalanceAvailableBalanceJSON contains
// the JSON metadata for the struct
// [LedgerAccountStatementGetResponseStartingBalanceAvailableBalance]
type ledgerAccountStatementGetResponseStartingBalanceAvailableBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseStartingBalanceAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerAccountStatementGetResponseStartingBalancePendingBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementGetResponseStartingBalancePendingBalanceJSON
}

// ledgerAccountStatementGetResponseStartingBalancePendingBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementGetResponseStartingBalancePendingBalance]
type ledgerAccountStatementGetResponseStartingBalancePendingBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseStartingBalancePendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerAccountStatementGetResponseStartingBalancePostedBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	Debits           int64 `json:"debits,required"`
	JSON             ledgerAccountStatementGetResponseStartingBalancePostedBalanceJSON
}

// ledgerAccountStatementGetResponseStartingBalancePostedBalanceJSON contains the
// JSON metadata for the struct
// [LedgerAccountStatementGetResponseStartingBalancePostedBalance]
type ledgerAccountStatementGetResponseStartingBalancePostedBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerAccountStatementGetResponseStartingBalancePostedBalance) UnmarshalJSON(data []byte) (err error) {
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
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountStatementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
