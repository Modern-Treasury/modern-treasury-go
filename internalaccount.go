// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// InternalAccountService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInternalAccountService] method instead.
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/internal_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update internal account
func (r *InternalAccountService) Update(ctx context.Context, id string, body InternalAccountUpdateParams, opts ...option.RequestOption) (res *InternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/internal_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list internal accounts
func (r *InternalAccountService) List(ctx context.Context, query InternalAccountListParams, opts ...option.RequestOption) (res *pagination.Page[InternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *InternalAccountService) ListAutoPaging(ctx context.Context, query InternalAccountListParams, opts ...option.RequestOption) *pagination.PageAutoPager[InternalAccount] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type InternalAccount struct {
	ID string `json:"id,required" format:"uuid"`
	// An array of AccountCapability objects that list the originating abilities of the
	// internal account and any relevant information for them.
	AccountCapabilities []InternalAccountAccountCapability `json:"account_capabilities,required"`
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
	Currency shared.Currency `json:"currency,required"`
	// If the internal account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,required,nullable" format:"uuid"`
	// The Legal Entity associated to this account
	LegalEntityID string `json:"legal_entity_id,required,nullable" format:"uuid"`
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
	ID                  apijson.Field
	AccountCapabilities apijson.Field
	AccountDetails      apijson.Field
	AccountType         apijson.Field
	Connection          apijson.Field
	CounterpartyID      apijson.Field
	CreatedAt           apijson.Field
	Currency            apijson.Field
	LedgerAccountID     apijson.Field
	LegalEntityID       apijson.Field
	LiveMode            apijson.Field
	Metadata            apijson.Field
	Name                apijson.Field
	Object              apijson.Field
	ParentAccountID     apijson.Field
	PartyAddress        apijson.Field
	PartyName           apijson.Field
	PartyType           apijson.Field
	RoutingDetails      apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internalAccountJSON) RawJSON() string {
	return r.raw
}

func (r InternalAccount) implementsPaymentOrderUltimateOriginatingAccount() {}

type InternalAccountAccountCapability struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// One of `debit` or `credit`. Indicates the direction of money movement this
	// capability is responsible for.
	Direction   shared.TransactionDirection `json:"direction,required"`
	DiscardedAt time.Time                   `json:"discarded_at,required,nullable" format:"date-time"`
	// A unique reference assigned by your bank for tracking and recognizing payment
	// files. It is important this is formatted exactly how the bank assigned it.
	Identifier string `json:"identifier,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// Indicates the the type of payment this capability is responsible for
	// originating.
	PaymentType InternalAccountAccountCapabilitiesPaymentType `json:"payment_type,required"`
	UpdatedAt   time.Time                                     `json:"updated_at,required" format:"date-time"`
	ExtraFields map[string]interface{}                        `json:"-,extras"`
	JSON        internalAccountAccountCapabilityJSON          `json:"-"`
}

// internalAccountAccountCapabilityJSON contains the JSON metadata for the struct
// [InternalAccountAccountCapability]
type internalAccountAccountCapabilityJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Direction   apijson.Field
	DiscardedAt apijson.Field
	Identifier  apijson.Field
	LiveMode    apijson.Field
	Object      apijson.Field
	PaymentType apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InternalAccountAccountCapability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internalAccountAccountCapabilityJSON) RawJSON() string {
	return r.raw
}

// Indicates the the type of payment this capability is responsible for
// originating.
type InternalAccountAccountCapabilitiesPaymentType string

const (
	InternalAccountAccountCapabilitiesPaymentTypeACH         InternalAccountAccountCapabilitiesPaymentType = "ach"
	InternalAccountAccountCapabilitiesPaymentTypeAuBecs      InternalAccountAccountCapabilitiesPaymentType = "au_becs"
	InternalAccountAccountCapabilitiesPaymentTypeBacs        InternalAccountAccountCapabilitiesPaymentType = "bacs"
	InternalAccountAccountCapabilitiesPaymentTypeBase        InternalAccountAccountCapabilitiesPaymentType = "base"
	InternalAccountAccountCapabilitiesPaymentTypeBook        InternalAccountAccountCapabilitiesPaymentType = "book"
	InternalAccountAccountCapabilitiesPaymentTypeCard        InternalAccountAccountCapabilitiesPaymentType = "card"
	InternalAccountAccountCapabilitiesPaymentTypeChats       InternalAccountAccountCapabilitiesPaymentType = "chats"
	InternalAccountAccountCapabilitiesPaymentTypeCheck       InternalAccountAccountCapabilitiesPaymentType = "check"
	InternalAccountAccountCapabilitiesPaymentTypeCrossBorder InternalAccountAccountCapabilitiesPaymentType = "cross_border"
	InternalAccountAccountCapabilitiesPaymentTypeDkNets      InternalAccountAccountCapabilitiesPaymentType = "dk_nets"
	InternalAccountAccountCapabilitiesPaymentTypeEft         InternalAccountAccountCapabilitiesPaymentType = "eft"
	InternalAccountAccountCapabilitiesPaymentTypeEthereum    InternalAccountAccountCapabilitiesPaymentType = "ethereum"
	InternalAccountAccountCapabilitiesPaymentTypeHuIcs       InternalAccountAccountCapabilitiesPaymentType = "hu_ics"
	InternalAccountAccountCapabilitiesPaymentTypeInterac     InternalAccountAccountCapabilitiesPaymentType = "interac"
	InternalAccountAccountCapabilitiesPaymentTypeMasav       InternalAccountAccountCapabilitiesPaymentType = "masav"
	InternalAccountAccountCapabilitiesPaymentTypeMxCcen      InternalAccountAccountCapabilitiesPaymentType = "mx_ccen"
	InternalAccountAccountCapabilitiesPaymentTypeNeft        InternalAccountAccountCapabilitiesPaymentType = "neft"
	InternalAccountAccountCapabilitiesPaymentTypeNics        InternalAccountAccountCapabilitiesPaymentType = "nics"
	InternalAccountAccountCapabilitiesPaymentTypeNzBecs      InternalAccountAccountCapabilitiesPaymentType = "nz_becs"
	InternalAccountAccountCapabilitiesPaymentTypePlElixir    InternalAccountAccountCapabilitiesPaymentType = "pl_elixir"
	InternalAccountAccountCapabilitiesPaymentTypePolygon     InternalAccountAccountCapabilitiesPaymentType = "polygon"
	InternalAccountAccountCapabilitiesPaymentTypeProvxchange InternalAccountAccountCapabilitiesPaymentType = "provxchange"
	InternalAccountAccountCapabilitiesPaymentTypeRoSent      InternalAccountAccountCapabilitiesPaymentType = "ro_sent"
	InternalAccountAccountCapabilitiesPaymentTypeRtp         InternalAccountAccountCapabilitiesPaymentType = "rtp"
	InternalAccountAccountCapabilitiesPaymentTypeSeBankgirot InternalAccountAccountCapabilitiesPaymentType = "se_bankgirot"
	InternalAccountAccountCapabilitiesPaymentTypeSen         InternalAccountAccountCapabilitiesPaymentType = "sen"
	InternalAccountAccountCapabilitiesPaymentTypeSepa        InternalAccountAccountCapabilitiesPaymentType = "sepa"
	InternalAccountAccountCapabilitiesPaymentTypeSgGiro      InternalAccountAccountCapabilitiesPaymentType = "sg_giro"
	InternalAccountAccountCapabilitiesPaymentTypeSic         InternalAccountAccountCapabilitiesPaymentType = "sic"
	InternalAccountAccountCapabilitiesPaymentTypeSignet      InternalAccountAccountCapabilitiesPaymentType = "signet"
	InternalAccountAccountCapabilitiesPaymentTypeSknbi       InternalAccountAccountCapabilitiesPaymentType = "sknbi"
	InternalAccountAccountCapabilitiesPaymentTypeSolana      InternalAccountAccountCapabilitiesPaymentType = "solana"
	InternalAccountAccountCapabilitiesPaymentTypeWire        InternalAccountAccountCapabilitiesPaymentType = "wire"
	InternalAccountAccountCapabilitiesPaymentTypeZengin      InternalAccountAccountCapabilitiesPaymentType = "zengin"
)

func (r InternalAccountAccountCapabilitiesPaymentType) IsKnown() bool {
	switch r {
	case InternalAccountAccountCapabilitiesPaymentTypeACH, InternalAccountAccountCapabilitiesPaymentTypeAuBecs, InternalAccountAccountCapabilitiesPaymentTypeBacs, InternalAccountAccountCapabilitiesPaymentTypeBase, InternalAccountAccountCapabilitiesPaymentTypeBook, InternalAccountAccountCapabilitiesPaymentTypeCard, InternalAccountAccountCapabilitiesPaymentTypeChats, InternalAccountAccountCapabilitiesPaymentTypeCheck, InternalAccountAccountCapabilitiesPaymentTypeCrossBorder, InternalAccountAccountCapabilitiesPaymentTypeDkNets, InternalAccountAccountCapabilitiesPaymentTypeEft, InternalAccountAccountCapabilitiesPaymentTypeEthereum, InternalAccountAccountCapabilitiesPaymentTypeHuIcs, InternalAccountAccountCapabilitiesPaymentTypeInterac, InternalAccountAccountCapabilitiesPaymentTypeMasav, InternalAccountAccountCapabilitiesPaymentTypeMxCcen, InternalAccountAccountCapabilitiesPaymentTypeNeft, InternalAccountAccountCapabilitiesPaymentTypeNics, InternalAccountAccountCapabilitiesPaymentTypeNzBecs, InternalAccountAccountCapabilitiesPaymentTypePlElixir, InternalAccountAccountCapabilitiesPaymentTypePolygon, InternalAccountAccountCapabilitiesPaymentTypeProvxchange, InternalAccountAccountCapabilitiesPaymentTypeRoSent, InternalAccountAccountCapabilitiesPaymentTypeRtp, InternalAccountAccountCapabilitiesPaymentTypeSeBankgirot, InternalAccountAccountCapabilitiesPaymentTypeSen, InternalAccountAccountCapabilitiesPaymentTypeSepa, InternalAccountAccountCapabilitiesPaymentTypeSgGiro, InternalAccountAccountCapabilitiesPaymentTypeSic, InternalAccountAccountCapabilitiesPaymentTypeSignet, InternalAccountAccountCapabilitiesPaymentTypeSknbi, InternalAccountAccountCapabilitiesPaymentTypeSolana, InternalAccountAccountCapabilitiesPaymentTypeWire, InternalAccountAccountCapabilitiesPaymentTypeZengin:
		return true
	}
	return false
}

// Can be checking, savings or other.
type InternalAccountAccountType string

const (
	InternalAccountAccountTypeBaseWallet     InternalAccountAccountType = "base_wallet"
	InternalAccountAccountTypeCash           InternalAccountAccountType = "cash"
	InternalAccountAccountTypeChecking       InternalAccountAccountType = "checking"
	InternalAccountAccountTypeCryptoWallet   InternalAccountAccountType = "crypto_wallet"
	InternalAccountAccountTypeEthereumWallet InternalAccountAccountType = "ethereum_wallet"
	InternalAccountAccountTypeGeneralLedger  InternalAccountAccountType = "general_ledger"
	InternalAccountAccountTypeLoan           InternalAccountAccountType = "loan"
	InternalAccountAccountTypeNonResident    InternalAccountAccountType = "non_resident"
	InternalAccountAccountTypeOther          InternalAccountAccountType = "other"
	InternalAccountAccountTypeOverdraft      InternalAccountAccountType = "overdraft"
	InternalAccountAccountTypePolygonWallet  InternalAccountAccountType = "polygon_wallet"
	InternalAccountAccountTypeSavings        InternalAccountAccountType = "savings"
	InternalAccountAccountTypeSolanaWallet   InternalAccountAccountType = "solana_wallet"
)

func (r InternalAccountAccountType) IsKnown() bool {
	switch r {
	case InternalAccountAccountTypeBaseWallet, InternalAccountAccountTypeCash, InternalAccountAccountTypeChecking, InternalAccountAccountTypeCryptoWallet, InternalAccountAccountTypeEthereumWallet, InternalAccountAccountTypeGeneralLedger, InternalAccountAccountTypeLoan, InternalAccountAccountTypeNonResident, InternalAccountAccountTypeOther, InternalAccountAccountTypeOverdraft, InternalAccountAccountTypePolygonWallet, InternalAccountAccountTypeSavings, InternalAccountAccountTypeSolanaWallet:
		return true
	}
	return false
}

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

func (r internalAccountPartyAddressJSON) RawJSON() string {
	return r.raw
}

// Either individual or business.
type InternalAccountPartyType string

const (
	InternalAccountPartyTypeBusiness   InternalAccountPartyType = "business"
	InternalAccountPartyTypeIndividual InternalAccountPartyType = "individual"
)

func (r InternalAccountPartyType) IsKnown() bool {
	switch r {
	case InternalAccountPartyTypeBusiness, InternalAccountPartyTypeIndividual:
		return true
	}
	return false
}

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
	// An array of AccountCapability objects that list the originating abilities of the
	// internal account and any relevant information for them.
	AccountCapabilities param.Field[[]InternalAccountNewParamsAccountCapability] `json:"account_capabilities"`
	// The account type, used to provision the appropriate account at the financial
	// institution.
	AccountType param.Field[InternalAccountNewParamsAccountType] `json:"account_type"`
	// The Counterparty associated to this account.
	CounterpartyID param.Field[string] `json:"counterparty_id"`
	// The LegalEntity associated to this account.
	LegalEntityID param.Field[string] `json:"legal_entity_id"`
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

func (r InternalAccountNewParamsCurrency) IsKnown() bool {
	switch r {
	case InternalAccountNewParamsCurrencyUsd, InternalAccountNewParamsCurrencyCad:
		return true
	}
	return false
}

type InternalAccountNewParamsAccountCapability struct {
	ID        param.Field[string]    `json:"id,required" format:"uuid"`
	CreatedAt param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	// One of `debit` or `credit`. Indicates the direction of money movement this
	// capability is responsible for.
	Direction   param.Field[shared.TransactionDirection] `json:"direction,required"`
	DiscardedAt param.Field[time.Time]                   `json:"discarded_at,required" format:"date-time"`
	// A unique reference assigned by your bank for tracking and recognizing payment
	// files. It is important this is formatted exactly how the bank assigned it.
	Identifier param.Field[string] `json:"identifier,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// Indicates the the type of payment this capability is responsible for
	// originating.
	PaymentType param.Field[InternalAccountNewParamsAccountCapabilitiesPaymentType] `json:"payment_type,required"`
	UpdatedAt   param.Field[time.Time]                                              `json:"updated_at,required" format:"date-time"`
	ExtraFields map[string]interface{}                                              `json:"-,extras"`
}

func (r InternalAccountNewParamsAccountCapability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates the the type of payment this capability is responsible for
// originating.
type InternalAccountNewParamsAccountCapabilitiesPaymentType string

const (
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeACH         InternalAccountNewParamsAccountCapabilitiesPaymentType = "ach"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeAuBecs      InternalAccountNewParamsAccountCapabilitiesPaymentType = "au_becs"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeBacs        InternalAccountNewParamsAccountCapabilitiesPaymentType = "bacs"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeBase        InternalAccountNewParamsAccountCapabilitiesPaymentType = "base"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeBook        InternalAccountNewParamsAccountCapabilitiesPaymentType = "book"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeCard        InternalAccountNewParamsAccountCapabilitiesPaymentType = "card"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeChats       InternalAccountNewParamsAccountCapabilitiesPaymentType = "chats"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeCheck       InternalAccountNewParamsAccountCapabilitiesPaymentType = "check"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeCrossBorder InternalAccountNewParamsAccountCapabilitiesPaymentType = "cross_border"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeDkNets      InternalAccountNewParamsAccountCapabilitiesPaymentType = "dk_nets"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeEft         InternalAccountNewParamsAccountCapabilitiesPaymentType = "eft"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeEthereum    InternalAccountNewParamsAccountCapabilitiesPaymentType = "ethereum"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeHuIcs       InternalAccountNewParamsAccountCapabilitiesPaymentType = "hu_ics"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeInterac     InternalAccountNewParamsAccountCapabilitiesPaymentType = "interac"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeMasav       InternalAccountNewParamsAccountCapabilitiesPaymentType = "masav"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeMxCcen      InternalAccountNewParamsAccountCapabilitiesPaymentType = "mx_ccen"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeNeft        InternalAccountNewParamsAccountCapabilitiesPaymentType = "neft"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeNics        InternalAccountNewParamsAccountCapabilitiesPaymentType = "nics"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeNzBecs      InternalAccountNewParamsAccountCapabilitiesPaymentType = "nz_becs"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypePlElixir    InternalAccountNewParamsAccountCapabilitiesPaymentType = "pl_elixir"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypePolygon     InternalAccountNewParamsAccountCapabilitiesPaymentType = "polygon"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeProvxchange InternalAccountNewParamsAccountCapabilitiesPaymentType = "provxchange"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeRoSent      InternalAccountNewParamsAccountCapabilitiesPaymentType = "ro_sent"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeRtp         InternalAccountNewParamsAccountCapabilitiesPaymentType = "rtp"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSeBankgirot InternalAccountNewParamsAccountCapabilitiesPaymentType = "se_bankgirot"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSen         InternalAccountNewParamsAccountCapabilitiesPaymentType = "sen"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSepa        InternalAccountNewParamsAccountCapabilitiesPaymentType = "sepa"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSgGiro      InternalAccountNewParamsAccountCapabilitiesPaymentType = "sg_giro"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSic         InternalAccountNewParamsAccountCapabilitiesPaymentType = "sic"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSignet      InternalAccountNewParamsAccountCapabilitiesPaymentType = "signet"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSknbi       InternalAccountNewParamsAccountCapabilitiesPaymentType = "sknbi"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeSolana      InternalAccountNewParamsAccountCapabilitiesPaymentType = "solana"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeWire        InternalAccountNewParamsAccountCapabilitiesPaymentType = "wire"
	InternalAccountNewParamsAccountCapabilitiesPaymentTypeZengin      InternalAccountNewParamsAccountCapabilitiesPaymentType = "zengin"
)

func (r InternalAccountNewParamsAccountCapabilitiesPaymentType) IsKnown() bool {
	switch r {
	case InternalAccountNewParamsAccountCapabilitiesPaymentTypeACH, InternalAccountNewParamsAccountCapabilitiesPaymentTypeAuBecs, InternalAccountNewParamsAccountCapabilitiesPaymentTypeBacs, InternalAccountNewParamsAccountCapabilitiesPaymentTypeBase, InternalAccountNewParamsAccountCapabilitiesPaymentTypeBook, InternalAccountNewParamsAccountCapabilitiesPaymentTypeCard, InternalAccountNewParamsAccountCapabilitiesPaymentTypeChats, InternalAccountNewParamsAccountCapabilitiesPaymentTypeCheck, InternalAccountNewParamsAccountCapabilitiesPaymentTypeCrossBorder, InternalAccountNewParamsAccountCapabilitiesPaymentTypeDkNets, InternalAccountNewParamsAccountCapabilitiesPaymentTypeEft, InternalAccountNewParamsAccountCapabilitiesPaymentTypeEthereum, InternalAccountNewParamsAccountCapabilitiesPaymentTypeHuIcs, InternalAccountNewParamsAccountCapabilitiesPaymentTypeInterac, InternalAccountNewParamsAccountCapabilitiesPaymentTypeMasav, InternalAccountNewParamsAccountCapabilitiesPaymentTypeMxCcen, InternalAccountNewParamsAccountCapabilitiesPaymentTypeNeft, InternalAccountNewParamsAccountCapabilitiesPaymentTypeNics, InternalAccountNewParamsAccountCapabilitiesPaymentTypeNzBecs, InternalAccountNewParamsAccountCapabilitiesPaymentTypePlElixir, InternalAccountNewParamsAccountCapabilitiesPaymentTypePolygon, InternalAccountNewParamsAccountCapabilitiesPaymentTypeProvxchange, InternalAccountNewParamsAccountCapabilitiesPaymentTypeRoSent, InternalAccountNewParamsAccountCapabilitiesPaymentTypeRtp, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSeBankgirot, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSen, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSepa, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSgGiro, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSic, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSignet, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSknbi, InternalAccountNewParamsAccountCapabilitiesPaymentTypeSolana, InternalAccountNewParamsAccountCapabilitiesPaymentTypeWire, InternalAccountNewParamsAccountCapabilitiesPaymentTypeZengin:
		return true
	}
	return false
}

// The account type, used to provision the appropriate account at the financial
// institution.
type InternalAccountNewParamsAccountType string

const (
	InternalAccountNewParamsAccountTypeBaseWallet     InternalAccountNewParamsAccountType = "base_wallet"
	InternalAccountNewParamsAccountTypeCash           InternalAccountNewParamsAccountType = "cash"
	InternalAccountNewParamsAccountTypeChecking       InternalAccountNewParamsAccountType = "checking"
	InternalAccountNewParamsAccountTypeCryptoWallet   InternalAccountNewParamsAccountType = "crypto_wallet"
	InternalAccountNewParamsAccountTypeEthereumWallet InternalAccountNewParamsAccountType = "ethereum_wallet"
	InternalAccountNewParamsAccountTypeGeneralLedger  InternalAccountNewParamsAccountType = "general_ledger"
	InternalAccountNewParamsAccountTypeLoan           InternalAccountNewParamsAccountType = "loan"
	InternalAccountNewParamsAccountTypeNonResident    InternalAccountNewParamsAccountType = "non_resident"
	InternalAccountNewParamsAccountTypeOther          InternalAccountNewParamsAccountType = "other"
	InternalAccountNewParamsAccountTypeOverdraft      InternalAccountNewParamsAccountType = "overdraft"
	InternalAccountNewParamsAccountTypePolygonWallet  InternalAccountNewParamsAccountType = "polygon_wallet"
	InternalAccountNewParamsAccountTypeSavings        InternalAccountNewParamsAccountType = "savings"
	InternalAccountNewParamsAccountTypeSolanaWallet   InternalAccountNewParamsAccountType = "solana_wallet"
)

func (r InternalAccountNewParamsAccountType) IsKnown() bool {
	switch r {
	case InternalAccountNewParamsAccountTypeBaseWallet, InternalAccountNewParamsAccountTypeCash, InternalAccountNewParamsAccountTypeChecking, InternalAccountNewParamsAccountTypeCryptoWallet, InternalAccountNewParamsAccountTypeEthereumWallet, InternalAccountNewParamsAccountTypeGeneralLedger, InternalAccountNewParamsAccountTypeLoan, InternalAccountNewParamsAccountTypeNonResident, InternalAccountNewParamsAccountTypeOther, InternalAccountNewParamsAccountTypeOverdraft, InternalAccountNewParamsAccountTypePolygonWallet, InternalAccountNewParamsAccountTypeSavings, InternalAccountNewParamsAccountTypeSolanaWallet:
		return true
	}
	return false
}

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
	// Only return internal accounts associated with this counterparty.
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// Only return internal accounts with this currency.
	Currency param.Field[shared.Currency] `query:"currency"`
	// Only return internal accounts associated with this legal entity.
	LegalEntityID param.Field[string] `query:"legal_entity_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Only return internal accounts that can originate payments with this direction.
	PaymentDirection param.Field[shared.TransactionDirection] `query:"payment_direction"`
	// Only return internal accounts that can make this type of payment.
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

// Only return internal accounts that can make this type of payment.
type InternalAccountListParamsPaymentType string

const (
	InternalAccountListParamsPaymentTypeACH         InternalAccountListParamsPaymentType = "ach"
	InternalAccountListParamsPaymentTypeAuBecs      InternalAccountListParamsPaymentType = "au_becs"
	InternalAccountListParamsPaymentTypeBacs        InternalAccountListParamsPaymentType = "bacs"
	InternalAccountListParamsPaymentTypeBase        InternalAccountListParamsPaymentType = "base"
	InternalAccountListParamsPaymentTypeBook        InternalAccountListParamsPaymentType = "book"
	InternalAccountListParamsPaymentTypeCard        InternalAccountListParamsPaymentType = "card"
	InternalAccountListParamsPaymentTypeChats       InternalAccountListParamsPaymentType = "chats"
	InternalAccountListParamsPaymentTypeCheck       InternalAccountListParamsPaymentType = "check"
	InternalAccountListParamsPaymentTypeCrossBorder InternalAccountListParamsPaymentType = "cross_border"
	InternalAccountListParamsPaymentTypeDkNets      InternalAccountListParamsPaymentType = "dk_nets"
	InternalAccountListParamsPaymentTypeEft         InternalAccountListParamsPaymentType = "eft"
	InternalAccountListParamsPaymentTypeEthereum    InternalAccountListParamsPaymentType = "ethereum"
	InternalAccountListParamsPaymentTypeHuIcs       InternalAccountListParamsPaymentType = "hu_ics"
	InternalAccountListParamsPaymentTypeInterac     InternalAccountListParamsPaymentType = "interac"
	InternalAccountListParamsPaymentTypeMasav       InternalAccountListParamsPaymentType = "masav"
	InternalAccountListParamsPaymentTypeMxCcen      InternalAccountListParamsPaymentType = "mx_ccen"
	InternalAccountListParamsPaymentTypeNeft        InternalAccountListParamsPaymentType = "neft"
	InternalAccountListParamsPaymentTypeNics        InternalAccountListParamsPaymentType = "nics"
	InternalAccountListParamsPaymentTypeNzBecs      InternalAccountListParamsPaymentType = "nz_becs"
	InternalAccountListParamsPaymentTypePlElixir    InternalAccountListParamsPaymentType = "pl_elixir"
	InternalAccountListParamsPaymentTypePolygon     InternalAccountListParamsPaymentType = "polygon"
	InternalAccountListParamsPaymentTypeProvxchange InternalAccountListParamsPaymentType = "provxchange"
	InternalAccountListParamsPaymentTypeRoSent      InternalAccountListParamsPaymentType = "ro_sent"
	InternalAccountListParamsPaymentTypeRtp         InternalAccountListParamsPaymentType = "rtp"
	InternalAccountListParamsPaymentTypeSeBankgirot InternalAccountListParamsPaymentType = "se_bankgirot"
	InternalAccountListParamsPaymentTypeSen         InternalAccountListParamsPaymentType = "sen"
	InternalAccountListParamsPaymentTypeSepa        InternalAccountListParamsPaymentType = "sepa"
	InternalAccountListParamsPaymentTypeSgGiro      InternalAccountListParamsPaymentType = "sg_giro"
	InternalAccountListParamsPaymentTypeSic         InternalAccountListParamsPaymentType = "sic"
	InternalAccountListParamsPaymentTypeSignet      InternalAccountListParamsPaymentType = "signet"
	InternalAccountListParamsPaymentTypeSknbi       InternalAccountListParamsPaymentType = "sknbi"
	InternalAccountListParamsPaymentTypeSolana      InternalAccountListParamsPaymentType = "solana"
	InternalAccountListParamsPaymentTypeWire        InternalAccountListParamsPaymentType = "wire"
	InternalAccountListParamsPaymentTypeZengin      InternalAccountListParamsPaymentType = "zengin"
)

func (r InternalAccountListParamsPaymentType) IsKnown() bool {
	switch r {
	case InternalAccountListParamsPaymentTypeACH, InternalAccountListParamsPaymentTypeAuBecs, InternalAccountListParamsPaymentTypeBacs, InternalAccountListParamsPaymentTypeBase, InternalAccountListParamsPaymentTypeBook, InternalAccountListParamsPaymentTypeCard, InternalAccountListParamsPaymentTypeChats, InternalAccountListParamsPaymentTypeCheck, InternalAccountListParamsPaymentTypeCrossBorder, InternalAccountListParamsPaymentTypeDkNets, InternalAccountListParamsPaymentTypeEft, InternalAccountListParamsPaymentTypeEthereum, InternalAccountListParamsPaymentTypeHuIcs, InternalAccountListParamsPaymentTypeInterac, InternalAccountListParamsPaymentTypeMasav, InternalAccountListParamsPaymentTypeMxCcen, InternalAccountListParamsPaymentTypeNeft, InternalAccountListParamsPaymentTypeNics, InternalAccountListParamsPaymentTypeNzBecs, InternalAccountListParamsPaymentTypePlElixir, InternalAccountListParamsPaymentTypePolygon, InternalAccountListParamsPaymentTypeProvxchange, InternalAccountListParamsPaymentTypeRoSent, InternalAccountListParamsPaymentTypeRtp, InternalAccountListParamsPaymentTypeSeBankgirot, InternalAccountListParamsPaymentTypeSen, InternalAccountListParamsPaymentTypeSepa, InternalAccountListParamsPaymentTypeSgGiro, InternalAccountListParamsPaymentTypeSic, InternalAccountListParamsPaymentTypeSignet, InternalAccountListParamsPaymentTypeSknbi, InternalAccountListParamsPaymentTypeSolana, InternalAccountListParamsPaymentTypeWire, InternalAccountListParamsPaymentTypeZengin:
		return true
	}
	return false
}
