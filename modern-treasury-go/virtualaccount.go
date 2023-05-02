package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type VirtualAccountService struct {
	Options []option.RequestOption
}

func NewVirtualAccountService(opts ...option.RequestOption) (r *VirtualAccountService) {
	r = &VirtualAccountService{}
	r.Options = opts
	return
}

// create virtual_account
func (r *VirtualAccountService) New(ctx context.Context, body VirtualAccountNewParams, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/virtual_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get virtual_account
func (r *VirtualAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// update virtual_account
func (r *VirtualAccountService) Update(ctx context.Context, id string, body VirtualAccountUpdateParams, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of virtual accounts.
func (r *VirtualAccountService) List(ctx context.Context, query VirtualAccountListParams, opts ...option.RequestOption) (res *shared.Page[VirtualAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/virtual_accounts"
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

// Get a list of virtual accounts.
func (r *VirtualAccountService) ListAutoPaging(ctx context.Context, query VirtualAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[VirtualAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete virtual_account
func (r *VirtualAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *VirtualAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/virtual_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type VirtualAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The name of the virtual account.
	Name string `json:"name,required"`
	// An optional free-form description for internal use.
	Description string `json:"description,required,nullable"`
	// The ID of a counterparty that the virtual account belongs to. Optional.
	CounterpartyID string `json:"counterparty_id,required,nullable" format:"uuid"`
	// The ID of the internal account that the virtual account is in.
	InternalAccountID string `json:"internal_account_id,required" format:"uuid"`
	// An array of account detail objects.
	AccountDetails []AccountDetail `json:"account_details,required"`
	// An array of routing detail objects. These will be the routing details of the
	// internal account.
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID string `json:"debit_ledger_account_id,required,nullable" format:"uuid"`
	// The ID of a credit normal ledger account. When money enters the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID string `json:"credit_ledger_account_id,required,nullable" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	JSON     VirtualAccountJSON
}

type VirtualAccountJSON struct {
	ID                    apijson.Metadata
	Object                apijson.Metadata
	LiveMode              apijson.Metadata
	CreatedAt             apijson.Metadata
	UpdatedAt             apijson.Metadata
	DiscardedAt           apijson.Metadata
	Name                  apijson.Metadata
	Description           apijson.Metadata
	CounterpartyID        apijson.Metadata
	InternalAccountID     apijson.Metadata
	AccountDetails        apijson.Metadata
	RoutingDetails        apijson.Metadata
	DebitLedgerAccountID  apijson.Metadata
	CreditLedgerAccountID apijson.Metadata
	Metadata              apijson.Metadata
	raw                   string
	Extras                map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into VirtualAccount using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *VirtualAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualAccountNewParams struct {
	// The name of the virtual account.
	Name field.Field[string] `json:"name,required"`
	// An optional description for internal use.
	Description field.Field[string] `json:"description"`
	// The ID of the counterparty that the virtual account belongs to.
	CounterpartyID field.Field[string] `json:"counterparty_id" format:"uuid"`
	// The ID of the internal account that this virtual account is associated with.
	InternalAccountID field.Field[string] `json:"internal_account_id,required" format:"uuid"`
	// An array of account detail objects.
	AccountDetails field.Field[[]AccountDetailParam] `json:"account_details"`
	// An array of routing detail objects.
	RoutingDetails field.Field[[]RoutingDetailParam] `json:"routing_details"`
	// The ID of a debit normal ledger account. When money enters the virtual account,
	// this ledger account will be debited. Must be accompanied by a
	// credit_ledger_account_id if present.
	DebitLedgerAccountID field.Field[string] `json:"debit_ledger_account_id" format:"uuid"`
	// The ID of a credit normal ledger account. When money leaves the virtual account,
	// this ledger account will be credited. Must be accompanied by a
	// debit_ledger_account_id if present.
	CreditLedgerAccountID field.Field[string] `json:"credit_ledger_account_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes VirtualAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r VirtualAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VirtualAccountUpdateParams struct {
	Name           field.Field[string]            `json:"name,nullable"`
	CounterpartyID field.Field[string]            `json:"counterparty_id" format:"uuid"`
	Metadata       field.Field[map[string]string] `json:"metadata"`
}

// MarshalJSON serializes VirtualAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r VirtualAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VirtualAccountListParams struct {
	AfterCursor       field.Field[string] `query:"after_cursor,nullable"`
	PerPage           field.Field[int64]  `query:"per_page"`
	InternalAccountID field.Field[string] `query:"internal_account_id" format:"uuid"`
	CounterpartyID    field.Field[string] `query:"counterparty_id" format:"uuid"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes VirtualAccountListParams into a url.Values of the query
// parameters associated with this value
func (r VirtualAccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
