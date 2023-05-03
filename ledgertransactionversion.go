package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// LedgerTransactionVersionService contains methods and other services that help
// with interacting with the Modern Treasury API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewLedgerTransactionVersionService] method instead.
type LedgerTransactionVersionService struct {
	Options []option.RequestOption
}

// NewLedgerTransactionVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLedgerTransactionVersionService(opts ...option.RequestOption) (r *LedgerTransactionVersionService) {
	r = &LedgerTransactionVersionService{}
	r.Options = opts
	return
}

// Get a list of ledger transaction versions.
func (r *LedgerTransactionVersionService) List(ctx context.Context, id string, query LedgerTransactionVersionListParams, opts ...option.RequestOption) (res *shared.Page[LedgerTransactionVersion], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("api/ledger_transactions/%s/versions", id)
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

// Get a list of ledger transaction versions.
func (r *LedgerTransactionVersionService) ListAutoPaging(ctx context.Context, id string, query LedgerTransactionVersionListParams, opts ...option.RequestOption) *shared.PageAutoPager[LedgerTransactionVersion] {
	return shared.NewPageAutoPager(r.List(ctx, id, query, opts...))
}

type LedgerTransactionVersion struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the ledger transaction
	LedgerTransactionID string `json:"ledger_transaction_id,required" format:"uuid"`
	// An optional description for internal use.
	Description string `json:"description,required,nullable"`
	// One of `pending`, `posted`, or `archived`
	Status LedgerTransactionVersionStatus `json:"status,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// An array of ledger entry objects.
	LedgerEntries []LedgerTransactionVersionLedgerEntries `json:"ledger_entries,required"`
	// The time on which the ledger transaction posted. This is null if the ledger
	// transaction is pending.
	PostedAt string `json:"posted_at,required,nullable" format:"time"`
	// The ID of the ledger this ledger transaction belongs to.
	LedgerID string `json:"ledger_id,required" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
	LedgerableType LedgerTransactionVersionLedgerableType `json:"ledgerable_type,required,nullable"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID string `json:"ledgerable_id,required,nullable" format:"uuid"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID string `json:"external_id,required,nullable"`
	// Version number of the ledger transaction.
	Version int64 `json:"version,required"`
	JSON    ledgerTransactionVersionJSON
}

// ledgerTransactionVersionJSON contains the JSON metadata for the struct
// [LedgerTransactionVersion]
type ledgerTransactionVersionJSON struct {
	ID                  apijson.Field
	Object              apijson.Field
	LiveMode            apijson.Field
	CreatedAt           apijson.Field
	LedgerTransactionID apijson.Field
	Description         apijson.Field
	Status              apijson.Field
	Metadata            apijson.Field
	EffectiveDate       apijson.Field
	LedgerEntries       apijson.Field
	PostedAt            apijson.Field
	LedgerID            apijson.Field
	LedgerableType      apijson.Field
	LedgerableID        apijson.Field
	ExternalID          apijson.Field
	Version             apijson.Field
	raw                 string
	Extras              map[string]apijson.Field
}

func (r *LedgerTransactionVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionStatus string

const (
	LedgerTransactionVersionStatusArchived LedgerTransactionVersionStatus = "archived"
	LedgerTransactionVersionStatusPending  LedgerTransactionVersionStatus = "pending"
	LedgerTransactionVersionStatusPosted   LedgerTransactionVersionStatus = "posted"
)

type LedgerTransactionVersionLedgerEntries struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount int64 `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction LedgerTransactionVersionLedgerEntriesDirection `json:"direction,required"`
	// Equal to the state of the ledger transaction when the ledger entry was created.
	// One of `pending`, `posted`, or `archived`.
	Status LedgerTransactionVersionLedgerEntriesStatus `json:"status,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID string `json:"ledger_account_id,required" format:"uuid"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LedgerAccountLockVersion int64 `json:"ledger_account_lock_version,required,nullable"`
	// The currency of the ledger account.
	LedgerAccountCurrency string `json:"ledger_account_currency,required"`
	// The currency exponent of the ledger account.
	LedgerAccountCurrencyExponent int64 `json:"ledger_account_currency_exponent,required"`
	// The ledger transaction that this ledger entry is associated with.
	LedgerTransactionID string `json:"ledger_transaction_id,required"`
	// The pending, posted, and available balances for this ledger entry's ledger
	// account. The posted balance is the sum of all posted entries on the account. The
	// pending balance is the sum of all pending and posted entries on the account. The
	// available balance is the posted incoming entries minus the sum of the pending
	// and posted outgoing amounts. Please see
	// https://docs.moderntreasury.com/docs/transaction-status-and-balances for more
	// details.
	ResultingLedgerAccountBalances LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances `json:"resulting_ledger_account_balances,required,nullable"`
	JSON                           ledgerTransactionVersionLedgerEntriesJSON
}

// ledgerTransactionVersionLedgerEntriesJSON contains the JSON metadata for the
// struct [LedgerTransactionVersionLedgerEntries]
type ledgerTransactionVersionLedgerEntriesJSON struct {
	ID                             apijson.Field
	Object                         apijson.Field
	LiveMode                       apijson.Field
	CreatedAt                      apijson.Field
	Amount                         apijson.Field
	Direction                      apijson.Field
	Status                         apijson.Field
	LedgerAccountID                apijson.Field
	LedgerAccountLockVersion       apijson.Field
	LedgerAccountCurrency          apijson.Field
	LedgerAccountCurrencyExponent  apijson.Field
	LedgerTransactionID            apijson.Field
	ResultingLedgerAccountBalances apijson.Field
	raw                            string
	Extras                         map[string]apijson.Field
}

func (r *LedgerTransactionVersionLedgerEntries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerEntriesDirection string

const (
	LedgerTransactionVersionLedgerEntriesDirectionCredit LedgerTransactionVersionLedgerEntriesDirection = "credit"
	LedgerTransactionVersionLedgerEntriesDirectionDebit  LedgerTransactionVersionLedgerEntriesDirection = "debit"
)

type LedgerTransactionVersionLedgerEntriesStatus string

const (
	LedgerTransactionVersionLedgerEntriesStatusArchived LedgerTransactionVersionLedgerEntriesStatus = "archived"
	LedgerTransactionVersionLedgerEntriesStatusPending  LedgerTransactionVersionLedgerEntriesStatus = "pending"
	LedgerTransactionVersionLedgerEntriesStatusPosted   LedgerTransactionVersionLedgerEntriesStatus = "posted"
)

// The pending, posted, and available balances for this ledger entry's ledger
// account. The posted balance is the sum of all posted entries on the account. The
// pending balance is the sum of all pending and posted entries on the account. The
// available balance is the posted incoming entries minus the sum of the pending
// and posted outgoing amounts. Please see
// https://docs.moderntreasury.com/docs/transaction-status-and-balances for more
// details.
type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances struct {
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance `json:"posted_balance,required"`
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance `json:"available_balance,required"`
	JSON             ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesJSON
}

// ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesJSON contains
// the JSON metadata for the struct
// [LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances]
type ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesJSON struct {
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	AvailableBalance apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The pending_balance is the sum of all pending and posted entries.
type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalanceJSON
}

// ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalanceJSON
// contains the JSON metadata for the struct
// [LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance]
type ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPendingBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The posted_balance is the sum of all posted entries.
type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalanceJSON
}

// ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalanceJSON
// contains the JSON metadata for the struct
// [LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance]
type ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesPostedBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The available_balance is the sum of all posted inbound entries and pending
// outbound entries. For credit normal, available_amount = posted_credits -
// pending_debits; for debit normal, available_amount = posted_debits -
// pending_credits.
type LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance struct {
	Credits int64 `json:"credits,required"`
	Debits  int64 `json:"debits,required"`
	Amount  int64 `json:"amount,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64 `json:"currency_exponent,required"`
	JSON             ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalanceJSON
}

// ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalanceJSON
// contains the JSON metadata for the struct
// [LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance]
type ledgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalanceJSON struct {
	Credits          apijson.Field
	Debits           apijson.Field
	Amount           apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	raw              string
	Extras           map[string]apijson.Field
}

func (r *LedgerTransactionVersionLedgerEntriesResultingLedgerAccountBalancesAvailableBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LedgerTransactionVersionLedgerableType string

const (
	LedgerTransactionVersionLedgerableTypeCounterparty          LedgerTransactionVersionLedgerableType = "counterparty"
	LedgerTransactionVersionLedgerableTypeExpectedPayment       LedgerTransactionVersionLedgerableType = "expected_payment"
	LedgerTransactionVersionLedgerableTypeIncomingPaymentDetail LedgerTransactionVersionLedgerableType = "incoming_payment_detail"
	LedgerTransactionVersionLedgerableTypeInternalAccount       LedgerTransactionVersionLedgerableType = "internal_account"
	LedgerTransactionVersionLedgerableTypeLineItem              LedgerTransactionVersionLedgerableType = "line_item"
	LedgerTransactionVersionLedgerableTypePaperItem             LedgerTransactionVersionLedgerableType = "paper_item"
	LedgerTransactionVersionLedgerableTypePaymentOrder          LedgerTransactionVersionLedgerableType = "payment_order"
	LedgerTransactionVersionLedgerableTypePaymentOrderAttempt   LedgerTransactionVersionLedgerableType = "payment_order_attempt"
	LedgerTransactionVersionLedgerableTypeReturn                LedgerTransactionVersionLedgerableType = "return"
	LedgerTransactionVersionLedgerableTypeReversal              LedgerTransactionVersionLedgerableType = "reversal"
)

type LedgerTransactionVersionListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// created_at timestamp. For example, for all dates after Jan 1 2000 12:00 UTC, use
	// created_at%5Bgt%5D=2000-01-01T12:00:00Z.
	CreatedAt param.Field[map[string]time.Time] `query:"created_at" format:"date-time"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to filter by the
	// version. For example, for all versions after 2, use version%5Bgt%5D=2.
	Version param.Field[map[string]int64] `query:"version"`
}

// URLQuery serializes [LedgerTransactionVersionListParams]'s query parameters as
// `url.Values`.
func (r LedgerTransactionVersionListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
