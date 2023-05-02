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

type InternalAccountService struct {
	Options        []option.RequestOption
	BalanceReports *InternalAccountBalanceReportService
}

func NewInternalAccountService(opts ...option.RequestOption) (r *InternalAccountService) {
	r = &InternalAccountService{}
	r.Options = opts
	r.BalanceReports = NewInternalAccountBalanceReportService(opts...)
	return
}

// create internal account
func (r *InternalAccountService) New(ctx context.Context, body InternalAccountNewParams, opts ...option.RequestOption) (res *InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/internal_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get internal account
func (r *InternalAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update internal account
func (r *InternalAccountService) Update(ctx context.Context, id string, body InternalAccountUpdateParams, opts ...option.RequestOption) (res *InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/internal_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list internal accounts
func (r *InternalAccountService) List(ctx context.Context, query InternalAccountListParams, opts ...option.RequestOption) (res *shared.Page[InternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/internal_accounts"
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

// list internal accounts
func (r *InternalAccountService) ListAutoPaging(ctx context.Context, query InternalAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[InternalAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type InternalAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Can be checking, savings or other.
	AccountType InternalAccountAccountType `json:"account_type,required,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name,required"`
	// Either individual or business.
	PartyType InternalAccountPartyType `json:"party_type,required,nullable"`
	// The address associated with the owner or null.
	PartyAddress InternalAccountPartyAddress `json:"party_address,required,nullable"`
	// A nickname for the account.
	Name string `json:"name,required,nullable"`
	// An array of account detail objects.
	AccountDetails []AccountDetail `json:"account_details,required"`
	// An array of routing detail objects.
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	// Specifies which financial institution the accounts belong to.
	Connection Connection `json:"connection,required"`
	// The currency of the account.
	Currency shared.Currency `json:"currency,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The parent InternalAccount of this account.
	ParentAccountID string `json:"parent_account_id,required,nullable" format:"uuid"`
	// The Counterparty associated to this account.
	CounterpartyID string `json:"counterparty_id,required,nullable" format:"uuid"`
	// If the internal account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,required,nullable" format:"uuid"`
	JSON            InternalAccountJSON
}

type InternalAccountJSON struct {
	ID              apijson.Metadata
	Object          apijson.Metadata
	LiveMode        apijson.Metadata
	CreatedAt       apijson.Metadata
	UpdatedAt       apijson.Metadata
	AccountType     apijson.Metadata
	PartyName       apijson.Metadata
	PartyType       apijson.Metadata
	PartyAddress    apijson.Metadata
	Name            apijson.Metadata
	AccountDetails  apijson.Metadata
	RoutingDetails  apijson.Metadata
	Connection      apijson.Metadata
	Currency        apijson.Metadata
	Metadata        apijson.Metadata
	ParentAccountID apijson.Metadata
	CounterpartyID  apijson.Metadata
	LedgerAccountID apijson.Metadata
	raw             string
	Extras          map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InternalAccount using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InternalAccountAccountType string

const (
	InternalAccountAccountTypeCash        InternalAccountAccountType = "cash"
	InternalAccountAccountTypeChecking    InternalAccountAccountType = "checking"
	InternalAccountAccountTypeLoan        InternalAccountAccountType = "loan"
	InternalAccountAccountTypeNonResident InternalAccountAccountType = "non_resident"
	InternalAccountAccountTypeOther       InternalAccountAccountType = "other"
	InternalAccountAccountTypeOverdraft   InternalAccountAccountType = "overdraft"
	InternalAccountAccountTypeSavings     InternalAccountAccountType = "savings"
)

type InternalAccountPartyType string

const (
	InternalAccountPartyTypeBusiness   InternalAccountPartyType = "business"
	InternalAccountPartyTypeIndividual InternalAccountPartyType = "individual"
)

type InternalAccountPartyAddress struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	// Region or State.
	Region string `json:"region,required,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required,nullable"`
	JSON    InternalAccountPartyAddressJSON
}

type InternalAccountPartyAddressJSON struct {
	ID         apijson.Metadata
	Object     apijson.Metadata
	LiveMode   apijson.Metadata
	CreatedAt  apijson.Metadata
	UpdatedAt  apijson.Metadata
	Line1      apijson.Metadata
	Line2      apijson.Metadata
	Locality   apijson.Metadata
	Region     apijson.Metadata
	PostalCode apijson.Metadata
	Country    apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InternalAccountPartyAddress
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InternalAccountPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InternalAccountNewParams struct {
	// The identifier of the financial institution the account belongs to.
	ConnectionID field.Field[string] `json:"connection_id,required"`
	// The nickname of the account.
	Name field.Field[string] `json:"name,required"`
	// The legal name of the entity which owns the account.
	PartyName field.Field[string] `json:"party_name,required"`
	// The address associated with the owner or null.
	PartyAddress field.Field[InternalAccountNewParamsPartyAddress] `json:"party_address"`
	// Either "USD" or "CAD". Internal accounts created at Increase only supports
	// "USD".
	Currency field.Field[InternalAccountNewParamsCurrency] `json:"currency,required"`
	// The identifier of the entity at Increase which owns the account.
	EntityID field.Field[string] `json:"entity_id"`
	// The parent internal account of this new account.
	ParentAccountID field.Field[string] `json:"parent_account_id"`
	// The Counterparty associated to this account.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
}

// MarshalJSON serializes InternalAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InternalAccountNewParamsPartyAddress struct {
	Line1 field.Field[string] `json:"line1,required"`
	Line2 field.Field[string] `json:"line2"`
	// Locality or City.
	Locality field.Field[string] `json:"locality,required"`
	// Region or State.
	Region field.Field[string] `json:"region,required"`
	// The postal code of the address.
	PostalCode field.Field[string] `json:"postal_code,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country field.Field[string] `json:"country,required"`
}

type InternalAccountNewParamsCurrency string

const (
	InternalAccountNewParamsCurrencyUsd InternalAccountNewParamsCurrency = "USD"
	InternalAccountNewParamsCurrencyCad InternalAccountNewParamsCurrency = "CAD"
)

type InternalAccountUpdateParams struct {
	// The nickname for the internal account.
	Name field.Field[string] `json:"name"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata field.Field[map[string]string] `json:"metadata"`
	// The parent internal account for this account.
	ParentAccountID field.Field[string] `json:"parent_account_id"`
	// The Counterparty associated to this account.
	CounterpartyID field.Field[string] `json:"counterparty_id"`
}

// MarshalJSON serializes InternalAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r InternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InternalAccountListParams struct {
	AfterCursor field.Field[string] `query:"after_cursor,nullable"`
	PerPage     field.Field[int64]  `query:"per_page"`
	// The currency associated with the internal account.
	Currency field.Field[shared.Currency] `json:"currency,nullable"`
	// The counterparty associated with the internal account.
	CounterpartyID field.Field[string] `query:"counterparty_id"`
	// The type of payment that can be made by the internal account.
	PaymentType field.Field[InternalAccountListParamsPaymentType] `query:"payment_type"`
	// The direction of payments that can be made by internal account.
	PaymentDirection field.Field[InternalAccountListParamsPaymentDirection] `query:"payment_direction"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata field.Field[map[string]string] `query:"metadata"`
}

// URLQuery serializes InternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r InternalAccountListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}

type InternalAccountListParamsPaymentType string

const (
	InternalAccountListParamsPaymentTypeACH         InternalAccountListParamsPaymentType = "ach"
	InternalAccountListParamsPaymentTypeAuBecs      InternalAccountListParamsPaymentType = "au_becs"
	InternalAccountListParamsPaymentTypeBacs        InternalAccountListParamsPaymentType = "bacs"
	InternalAccountListParamsPaymentTypeBook        InternalAccountListParamsPaymentType = "book"
	InternalAccountListParamsPaymentTypeCard        InternalAccountListParamsPaymentType = "card"
	InternalAccountListParamsPaymentTypeCheck       InternalAccountListParamsPaymentType = "check"
	InternalAccountListParamsPaymentTypeCrossBorder InternalAccountListParamsPaymentType = "cross_border"
	InternalAccountListParamsPaymentTypeEft         InternalAccountListParamsPaymentType = "eft"
	InternalAccountListParamsPaymentTypeInterac     InternalAccountListParamsPaymentType = "interac"
	InternalAccountListParamsPaymentTypeMasav       InternalAccountListParamsPaymentType = "masav"
	InternalAccountListParamsPaymentTypeNeft        InternalAccountListParamsPaymentType = "neft"
	InternalAccountListParamsPaymentTypeProvxchange InternalAccountListParamsPaymentType = "provxchange"
	InternalAccountListParamsPaymentTypeRtp         InternalAccountListParamsPaymentType = "rtp"
	InternalAccountListParamsPaymentTypeSen         InternalAccountListParamsPaymentType = "sen"
	InternalAccountListParamsPaymentTypeSepa        InternalAccountListParamsPaymentType = "sepa"
	InternalAccountListParamsPaymentTypeSignet      InternalAccountListParamsPaymentType = "signet"
	InternalAccountListParamsPaymentTypeWire        InternalAccountListParamsPaymentType = "wire"
)

type InternalAccountListParamsPaymentDirection string

const (
	InternalAccountListParamsPaymentDirectionCredit InternalAccountListParamsPaymentDirection = "credit"
	InternalAccountListParamsPaymentDirectionDebit  InternalAccountListParamsPaymentDirection = "debit"
)
