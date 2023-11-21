// File generated from our OpenAPI spec by Stainless.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// InternalAccountService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewInternalAccountService] method instead.
type InternalAccountService struct {
	Options        []option.RequestOption
	BalanceReports *InternalAccountBalanceReportService
}

// NewInternalAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
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
	ID string `json:"id,required" format:"uuid"`
	// An array of account detail objects.
	AccountDetails []AccountDetail `json:"account_details,required"`
	// Can be checking, savings or other.
	AccountType InternalAccountAccountType `json:"account_type,required,nullable"`
	// Specifies which financial institution the accounts belong to.
	Connection Connection `json:"connection,required"`
	// The Counterparty associated to this account.
	CounterpartyID string    `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the account.
	Currency shared.Currency `json:"currency,required,nullable"`
	// If the internal account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// A nickname for the account.
	Name   string `json:"name,required,nullable"`
	Object string `json:"object,required"`
	// The parent InternalAccount of this account.
	ParentAccountID string `json:"parent_account_id,required,nullable" format:"uuid"`
	// The address associated with the owner or null.
	PartyAddress InternalAccountPartyAddress `json:"party_address,required,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name,required"`
	// Either individual or business.
	PartyType InternalAccountPartyType `json:"party_type,required,nullable"`
	// An array of routing detail objects.
	RoutingDetails []RoutingDetail     `json:"routing_details,required"`
	UpdatedAt      time.Time           `json:"updated_at,required" format:"date-time"`
	JSON           internalAccountJSON `json:"-"`
}

// internalAccountJSON contains the JSON metadata for the struct [InternalAccount]
type internalAccountJSON struct {
	ID              apijson.Field
	AccountDetails  apijson.Field
	AccountType     apijson.Field
	Connection      apijson.Field
	CounterpartyID  apijson.Field
	CreatedAt       apijson.Field
	Currency        apijson.Field
	LedgerAccountID apijson.Field
	LiveMode        apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Object          apijson.Field
	ParentAccountID apijson.Field
	PartyAddress    apijson.Field
	PartyName       apijson.Field
	PartyType       apijson.Field
	RoutingDetails  apijson.Field
	UpdatedAt       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InternalAccount) implementsPaymentOrderUltimateOriginatingAccount() {}

// Can be checking, savings or other.
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

// The address associated with the owner or null.
type InternalAccountPartyAddress struct {
	ID string `json:"id,required" format:"uuid"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country   string    `json:"country,required,nullable"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	Object   string `json:"object,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Region or State.
	Region    string                          `json:"region,required,nullable"`
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	JSON      internalAccountPartyAddressJSON `json:"-"`
}

// internalAccountPartyAddressJSON contains the JSON metadata for the struct
// [InternalAccountPartyAddress]
type internalAccountPartyAddressJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	CreatedAt   apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	LiveMode    apijson.Field
	Locality    apijson.Field
	Object      apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InternalAccountPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Either individual or business.
type InternalAccountPartyType string

const (
	InternalAccountPartyTypeBusiness   InternalAccountPartyType = "business"
	InternalAccountPartyTypeIndividual InternalAccountPartyType = "individual"
)

type InternalAccountNewParams struct {
	// The identifier of the financial institution the account belongs to.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// Either "USD" or "CAD". Internal accounts created at Increase only supports
	// "USD".
	Currency param.Field[InternalAccountNewParamsCurrency] `json:"currency,required"`
	// The nickname of the account.
	Name param.Field[string] `json:"name,required"`
	// The legal name of the entity which owns the account.
	PartyName param.Field[string] `json:"party_name,required"`
	// The Counterparty associated to this account.
	CounterpartyID param.Field[string] `json:"counterparty_id"`
	// The parent internal account of this new account.
	ParentAccountID param.Field[string] `json:"parent_account_id"`
	// The address associated with the owner or null.
	PartyAddress param.Field[InternalAccountNewParamsPartyAddress] `json:"party_address"`
	// A hash of vendor specific attributes that will be used when creating the account
	// at the vendor specified by the given connection.
	VendorAttributes param.Field[map[string]string] `json:"vendor_attributes"`
}

func (r InternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either "USD" or "CAD". Internal accounts created at Increase only supports
// "USD".
type InternalAccountNewParamsCurrency string

const (
	InternalAccountNewParamsCurrencyUsd InternalAccountNewParamsCurrency = "USD"
	InternalAccountNewParamsCurrencyCad InternalAccountNewParamsCurrency = "CAD"
)

// The address associated with the owner or null.
type InternalAccountNewParamsPartyAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
	Line1   param.Field[string] `json:"line1,required"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	Line2  param.Field[string] `json:"line2"`
}

func (r InternalAccountNewParamsPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InternalAccountUpdateParams struct {
	// The Counterparty associated to this account.
	CounterpartyID param.Field[string] `json:"counterparty_id"`
	// The Ledger Account associated to this account.
	LedgerAccountID param.Field[string] `json:"ledger_account_id"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The nickname for the internal account.
	Name param.Field[string] `json:"name"`
	// The parent internal account for this account.
	ParentAccountID param.Field[string] `json:"parent_account_id"`
}

func (r InternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InternalAccountListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// The counterparty associated with the internal account.
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// The currency associated with the internal account.
	Currency param.Field[shared.Currency] `query:"currency"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// The direction of payments that can be made by internal account.
	PaymentDirection param.Field[shared.TransactionDirection] `query:"payment_direction"`
	// The type of payment that can be made by the internal account.
	PaymentType param.Field[InternalAccountListParamsPaymentType] `query:"payment_type"`
	PerPage     param.Field[int64]                                `query:"per_page"`
}

// URLQuery serializes [InternalAccountListParams]'s query parameters as
// `url.Values`.
func (r InternalAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of payment that can be made by the internal account.
type InternalAccountListParamsPaymentType string

const (
	InternalAccountListParamsPaymentTypeACH         InternalAccountListParamsPaymentType = "ach"
	InternalAccountListParamsPaymentTypeAuBecs      InternalAccountListParamsPaymentType = "au_becs"
	InternalAccountListParamsPaymentTypeBacs        InternalAccountListParamsPaymentType = "bacs"
	InternalAccountListParamsPaymentTypeBook        InternalAccountListParamsPaymentType = "book"
	InternalAccountListParamsPaymentTypeCard        InternalAccountListParamsPaymentType = "card"
	InternalAccountListParamsPaymentTypeChats       InternalAccountListParamsPaymentType = "chats"
	InternalAccountListParamsPaymentTypeCheck       InternalAccountListParamsPaymentType = "check"
	InternalAccountListParamsPaymentTypeCrossBorder InternalAccountListParamsPaymentType = "cross_border"
	InternalAccountListParamsPaymentTypeDkNets      InternalAccountListParamsPaymentType = "dk_nets"
	InternalAccountListParamsPaymentTypeEft         InternalAccountListParamsPaymentType = "eft"
	InternalAccountListParamsPaymentTypeHuIcs       InternalAccountListParamsPaymentType = "hu_ics"
	InternalAccountListParamsPaymentTypeInterac     InternalAccountListParamsPaymentType = "interac"
	InternalAccountListParamsPaymentTypeMasav       InternalAccountListParamsPaymentType = "masav"
	InternalAccountListParamsPaymentTypeNeft        InternalAccountListParamsPaymentType = "neft"
	InternalAccountListParamsPaymentTypeNics        InternalAccountListParamsPaymentType = "nics"
	InternalAccountListParamsPaymentTypeNzBecs      InternalAccountListParamsPaymentType = "nz_becs"
	InternalAccountListParamsPaymentTypeProvxchange InternalAccountListParamsPaymentType = "provxchange"
	InternalAccountListParamsPaymentTypeRoSent      InternalAccountListParamsPaymentType = "ro_sent"
	InternalAccountListParamsPaymentTypeRtp         InternalAccountListParamsPaymentType = "rtp"
	InternalAccountListParamsPaymentTypeSeBankgirot InternalAccountListParamsPaymentType = "se_bankgirot"
	InternalAccountListParamsPaymentTypeSen         InternalAccountListParamsPaymentType = "sen"
	InternalAccountListParamsPaymentTypeSepa        InternalAccountListParamsPaymentType = "sepa"
	InternalAccountListParamsPaymentTypeSgGiro      InternalAccountListParamsPaymentType = "sg_giro"
	InternalAccountListParamsPaymentTypeSic         InternalAccountListParamsPaymentType = "sic"
	InternalAccountListParamsPaymentTypeSignet      InternalAccountListParamsPaymentType = "signet"
	InternalAccountListParamsPaymentTypeWire        InternalAccountListParamsPaymentType = "wire"
	InternalAccountListParamsPaymentTypeZengin      InternalAccountListParamsPaymentType = "zengin"
)
