// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// LedgerAccountStatementService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLedgerAccountStatementService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	path := "api/ledger_account_statements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single ledger account statement.
func (r *LedgerAccountStatementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LedgerAccountStatementGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
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
	EffectiveAtLowerBound time.Time `json:"effective_at_lower_bound,required" format:"date-time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound time.Time `json:"effective_at_upper_bound,required" format:"date-time"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	EndingBalance shared.LedgerBalances `json:"ending_balance,required"`
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
	StartingBalance shared.LedgerBalances                 `json:"starting_balance,required"`
	UpdatedAt       time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON            ledgerAccountStatementNewResponseJSON `json:"-"`
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

func (r ledgerAccountStatementNewResponseJSON) RawJSON() string {
	return r.raw
}

type LedgerAccountStatementGetResponse struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the ledger account statement.
	Description string `json:"description,required,nullable"`
	// The inclusive lower bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtLowerBound time.Time `json:"effective_at_lower_bound,required" format:"date-time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound time.Time `json:"effective_at_upper_bound,required" format:"date-time"`
	// The pending, posted, and available balances for this ledger account at the
	// `effective_at_upper_bound`. The posted balance is the sum of all posted entries
	// on the account. The pending balance is the sum of all pending and posted entries
	// on the account. The available balance is the posted incoming entries minus the
	// sum of the pending and posted outgoing amounts.
	EndingBalance shared.LedgerBalances `json:"ending_balance,required"`
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
	StartingBalance shared.LedgerBalances                 `json:"starting_balance,required"`
	UpdatedAt       time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON            ledgerAccountStatementGetResponseJSON `json:"-"`
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

func (r ledgerAccountStatementGetResponseJSON) RawJSON() string {
	return r.raw
}

type LedgerAccountStatementNewParams struct {
	// The inclusive lower bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtLowerBound param.Field[time.Time] `json:"effective_at_lower_bound,required" format:"date-time"`
	// The exclusive upper bound of the effective_at timestamp of the ledger entries to
	// be included in the ledger account statement.
	EffectiveAtUpperBound param.Field[time.Time] `json:"effective_at_upper_bound,required" format:"date-time"`
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
