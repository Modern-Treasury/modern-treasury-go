// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
	"github.com/tidwall/gjson"
)

// ExternalAccountService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalAccountService] method instead.
type ExternalAccountService struct {
	Options []option.RequestOption
}

// NewExternalAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExternalAccountService(opts ...option.RequestOption) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Options = opts
	return
}

// create external account
func (r *ExternalAccountService) New(ctx context.Context, body ExternalAccountNewParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/external_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// show external account
func (r *ExternalAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update external account
func (r *ExternalAccountService) Update(ctx context.Context, id string, body ExternalAccountUpdateParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list external accounts
func (r *ExternalAccountService) List(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) (res *pagination.Page[ExternalAccount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/external_accounts"
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

// list external accounts
func (r *ExternalAccountService) ListAutoPaging(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ExternalAccount] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete external account
func (r *ExternalAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// complete verification of external account
func (r *ExternalAccountService) CompleteVerification(ctx context.Context, id string, body ExternalAccountCompleteVerificationParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/external_accounts/%s/complete_verification", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// verify external account
func (r *ExternalAccountService) Verify(ctx context.Context, id string, body ExternalAccountVerifyParams, opts ...option.RequestOption) (res *ExternalAccountVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/external_accounts/%s/verify", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExternalAccount struct {
	ID             string          `json:"id,required" format:"uuid"`
	AccountDetails []AccountDetail `json:"account_details,required"`
	// Can be `checking`, `savings` or `other`.
	AccountType    ExternalAccountType    `json:"account_type,required"`
	ContactDetails []shared.ContactDetail `json:"contact_details,required"`
	CounterpartyID string                 `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time              `json:"created_at,required" format:"date-time"`
	DiscardedAt    time.Time              `json:"discarded_at,required,nullable" format:"date-time"`
	// An optional user-defined 180 character unique identifier.
	ExternalID string `json:"external_id,required,nullable"`
	// If the external account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name   string `json:"name,required,nullable"`
	Object string `json:"object,required"`
	// The address associated with the owner or `null`.
	PartyAddress shared.Address `json:"party_address,required,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name,required"`
	// Either `individual` or `business`.
	PartyType          ExternalAccountPartyType          `json:"party_type,required,nullable"`
	RoutingDetails     []RoutingDetail                   `json:"routing_details,required"`
	UpdatedAt          time.Time                         `json:"updated_at,required" format:"date-time"`
	VerificationSource ExternalAccountVerificationSource `json:"verification_source,required,nullable"`
	VerificationStatus ExternalAccountVerificationStatus `json:"verification_status,required"`
	JSON               externalAccountJSON               `json:"-"`
}

// externalAccountJSON contains the JSON metadata for the struct [ExternalAccount]
type externalAccountJSON struct {
	ID                 apijson.Field
	AccountDetails     apijson.Field
	AccountType        apijson.Field
	ContactDetails     apijson.Field
	CounterpartyID     apijson.Field
	CreatedAt          apijson.Field
	DiscardedAt        apijson.Field
	ExternalID         apijson.Field
	LedgerAccountID    apijson.Field
	LiveMode           apijson.Field
	Metadata           apijson.Field
	Name               apijson.Field
	Object             apijson.Field
	PartyAddress       apijson.Field
	PartyName          apijson.Field
	PartyType          apijson.Field
	RoutingDetails     apijson.Field
	UpdatedAt          apijson.Field
	VerificationSource apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalAccountJSON) RawJSON() string {
	return r.raw
}

func (r ExternalAccount) implementsExternalAccountVerifyResponse() {}

// Either `individual` or `business`.
type ExternalAccountPartyType string

const (
	ExternalAccountPartyTypeBusiness   ExternalAccountPartyType = "business"
	ExternalAccountPartyTypeIndividual ExternalAccountPartyType = "individual"
)

func (r ExternalAccountPartyType) IsKnown() bool {
	switch r {
	case ExternalAccountPartyTypeBusiness, ExternalAccountPartyTypeIndividual:
		return true
	}
	return false
}

type ExternalAccountVerificationSource string

const (
	ExternalAccountVerificationSourceACHPrenote    ExternalAccountVerificationSource = "ach_prenote"
	ExternalAccountVerificationSourceMicrodeposits ExternalAccountVerificationSource = "microdeposits"
	ExternalAccountVerificationSourcePlaid         ExternalAccountVerificationSource = "plaid"
)

func (r ExternalAccountVerificationSource) IsKnown() bool {
	switch r {
	case ExternalAccountVerificationSourceACHPrenote, ExternalAccountVerificationSourceMicrodeposits, ExternalAccountVerificationSourcePlaid:
		return true
	}
	return false
}

type ExternalAccountVerificationStatus string

const (
	ExternalAccountVerificationStatusPendingVerification ExternalAccountVerificationStatus = "pending_verification"
	ExternalAccountVerificationStatusUnverified          ExternalAccountVerificationStatus = "unverified"
	ExternalAccountVerificationStatusVerified            ExternalAccountVerificationStatus = "verified"
)

func (r ExternalAccountVerificationStatus) IsKnown() bool {
	switch r {
	case ExternalAccountVerificationStatusPendingVerification, ExternalAccountVerificationStatusUnverified, ExternalAccountVerificationStatusVerified:
		return true
	}
	return false
}

// Can be `checking`, `savings` or `other`.
type ExternalAccountType string

const (
	ExternalAccountTypeBaseWallet     ExternalAccountType = "base_wallet"
	ExternalAccountTypeCash           ExternalAccountType = "cash"
	ExternalAccountTypeChecking       ExternalAccountType = "checking"
	ExternalAccountTypeCryptoWallet   ExternalAccountType = "crypto_wallet"
	ExternalAccountTypeEthereumWallet ExternalAccountType = "ethereum_wallet"
	ExternalAccountTypeGeneralLedger  ExternalAccountType = "general_ledger"
	ExternalAccountTypeLoan           ExternalAccountType = "loan"
	ExternalAccountTypeNonResident    ExternalAccountType = "non_resident"
	ExternalAccountTypeOther          ExternalAccountType = "other"
	ExternalAccountTypeOverdraft      ExternalAccountType = "overdraft"
	ExternalAccountTypePolygonWallet  ExternalAccountType = "polygon_wallet"
	ExternalAccountTypeSavings        ExternalAccountType = "savings"
	ExternalAccountTypeSolanaWallet   ExternalAccountType = "solana_wallet"
)

func (r ExternalAccountType) IsKnown() bool {
	switch r {
	case ExternalAccountTypeBaseWallet, ExternalAccountTypeCash, ExternalAccountTypeChecking, ExternalAccountTypeCryptoWallet, ExternalAccountTypeEthereumWallet, ExternalAccountTypeGeneralLedger, ExternalAccountTypeLoan, ExternalAccountTypeNonResident, ExternalAccountTypeOther, ExternalAccountTypeOverdraft, ExternalAccountTypePolygonWallet, ExternalAccountTypeSavings, ExternalAccountTypeSolanaWallet:
		return true
	}
	return false
}

type ExternalAccountVerifyResponse struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// This field can have the runtime type of [[]AccountDetail].
	AccountDetails interface{} `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType ExternalAccountType `json:"account_type"`
	// This field can have the runtime type of [[]shared.ContactDetail].
	ContactDetails interface{} `json:"contact_details"`
	CounterpartyID string      `json:"counterparty_id,nullable" format:"uuid"`
	DiscardedAt    time.Time   `json:"discarded_at,nullable" format:"date-time"`
	// The ID of the external account.
	ExternalAccountID string `json:"external_account_id" format:"uuid"`
	// An optional user-defined 180 character unique identifier.
	ExternalID string `json:"external_id,nullable"`
	// If the external account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,nullable" format:"uuid"`
	// This field can have the runtime type of [map[string]string].
	Metadata interface{} `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name string `json:"name,nullable"`
	// The ID of the internal account where the micro-deposits originate from.
	OriginatingAccountID string `json:"originating_account_id" format:"uuid"`
	// The address associated with the owner or `null`.
	PartyAddress shared.Address `json:"party_address,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name"`
	// Either `individual` or `business`.
	PartyType ExternalAccountVerifyResponsePartyType `json:"party_type,nullable"`
	// The type of payment that can be made to this account. Can be `ach`, `eft`, or
	// `rtp`.
	PaymentType ExternalAccountVerifyResponsePaymentType `json:"payment_type"`
	// The priority of the payment. Can be `normal` or `high`.
	Priority ExternalAccountVerifyResponsePriority `json:"priority,nullable"`
	// This field can have the runtime type of [[]RoutingDetail].
	RoutingDetails interface{} `json:"routing_details"`
	// The status of the verification attempt. Can be `pending_verification`,
	// `verified`, `failed`, or `cancelled`.
	Status             ExternalAccountVerifyResponseStatus             `json:"status"`
	VerificationSource ExternalAccountVerifyResponseVerificationSource `json:"verification_source,nullable"`
	VerificationStatus ExternalAccountVerifyResponseVerificationStatus `json:"verification_status"`
	JSON               externalAccountVerifyResponseJSON               `json:"-"`
	union              ExternalAccountVerifyResponseUnion
}

// externalAccountVerifyResponseJSON contains the JSON metadata for the struct
// [ExternalAccountVerifyResponse]
type externalAccountVerifyResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	LiveMode             apijson.Field
	Object               apijson.Field
	UpdatedAt            apijson.Field
	AccountDetails       apijson.Field
	AccountType          apijson.Field
	ContactDetails       apijson.Field
	CounterpartyID       apijson.Field
	DiscardedAt          apijson.Field
	ExternalAccountID    apijson.Field
	ExternalID           apijson.Field
	LedgerAccountID      apijson.Field
	Metadata             apijson.Field
	Name                 apijson.Field
	OriginatingAccountID apijson.Field
	PartyAddress         apijson.Field
	PartyName            apijson.Field
	PartyType            apijson.Field
	PaymentType          apijson.Field
	Priority             apijson.Field
	RoutingDetails       apijson.Field
	Status               apijson.Field
	VerificationSource   apijson.Field
	VerificationStatus   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r externalAccountVerifyResponseJSON) RawJSON() string {
	return r.raw
}

func (r *ExternalAccountVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	*r = ExternalAccountVerifyResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ExternalAccountVerifyResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [ExternalAccount],
// [ExternalAccountVerifyResponseExternalAccountVerificationAttempt].
func (r ExternalAccountVerifyResponse) AsUnion() ExternalAccountVerifyResponseUnion {
	return r.union
}

// Union satisfied by [ExternalAccount] or
// [ExternalAccountVerifyResponseExternalAccountVerificationAttempt].
type ExternalAccountVerifyResponseUnion interface {
	implementsExternalAccountVerifyResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExternalAccountVerifyResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ExternalAccountVerifyResponseExternalAccountVerificationAttempt{}),
		},
	)
}

type ExternalAccountVerifyResponseExternalAccountVerificationAttempt struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the external account.
	ExternalAccountID string `json:"external_account_id,required" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The ID of the internal account where the micro-deposits originate from.
	OriginatingAccountID string `json:"originating_account_id,required" format:"uuid"`
	// The type of payment that can be made to this account. Can be `ach`, `eft`, or
	// `rtp`.
	PaymentType ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType `json:"payment_type,required"`
	// The priority of the payment. Can be `normal` or `high`.
	Priority ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriority `json:"priority,required,nullable"`
	// The status of the verification attempt. Can be `pending_verification`,
	// `verified`, `failed`, or `cancelled`.
	Status    ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus `json:"status,required"`
	UpdatedAt time.Time                                                             `json:"updated_at,required" format:"date-time"`
	JSON      externalAccountVerifyResponseExternalAccountVerificationAttemptJSON   `json:"-"`
}

// externalAccountVerifyResponseExternalAccountVerificationAttemptJSON contains the
// JSON metadata for the struct
// [ExternalAccountVerifyResponseExternalAccountVerificationAttempt]
type externalAccountVerifyResponseExternalAccountVerificationAttemptJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ExternalAccountID    apijson.Field
	LiveMode             apijson.Field
	Object               apijson.Field
	OriginatingAccountID apijson.Field
	PaymentType          apijson.Field
	Priority             apijson.Field
	Status               apijson.Field
	UpdatedAt            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ExternalAccountVerifyResponseExternalAccountVerificationAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalAccountVerifyResponseExternalAccountVerificationAttemptJSON) RawJSON() string {
	return r.raw
}

func (r ExternalAccountVerifyResponseExternalAccountVerificationAttempt) implementsExternalAccountVerifyResponse() {
}

// The type of payment that can be made to this account. Can be `ach`, `eft`, or
// `rtp`.
type ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType string

const (
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeACH         ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "ach"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeAuBecs      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "au_becs"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeBacs        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "bacs"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeBase        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "base"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeBook        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "book"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeCard        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "card"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeChats       ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "chats"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeCheck       ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "check"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeCrossBorder ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "cross_border"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeDkNets      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "dk_nets"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeEft         ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "eft"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeEthereum    ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "ethereum"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeGBFps       ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "gb_fps"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeHuIcs       ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "hu_ics"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeInterac     ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "interac"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeMasav       ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "masav"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeMxCcen      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "mx_ccen"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeNeft        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "neft"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeNics        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "nics"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeNzBecs      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "nz_becs"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypePlElixir    ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "pl_elixir"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypePolygon     ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "polygon"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeProvxchange ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "provxchange"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeRoSent      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "ro_sent"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeRtp         ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "rtp"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSeBankgirot ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "se_bankgirot"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSen         ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "sen"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSepa        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "sepa"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSgGiro      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "sg_giro"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSic         ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "sic"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSignet      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "signet"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSknbi       ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "sknbi"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSolana      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "solana"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeWire        ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "wire"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeZengin      ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType = "zengin"
)

func (r ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentType) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeACH, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeAuBecs, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeBacs, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeBase, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeBook, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeCard, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeChats, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeCheck, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeCrossBorder, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeDkNets, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeEft, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeEthereum, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeGBFps, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeHuIcs, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeInterac, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeMasav, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeMxCcen, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeNeft, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeNics, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeNzBecs, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypePlElixir, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypePolygon, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeProvxchange, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeRoSent, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeRtp, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSeBankgirot, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSen, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSepa, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSgGiro, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSic, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSignet, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSknbi, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeSolana, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeWire, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPaymentTypeZengin:
		return true
	}
	return false
}

// The priority of the payment. Can be `normal` or `high`.
type ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriority string

const (
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriorityHigh   ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriority = "high"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriorityNormal ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriority = "normal"
)

func (r ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriority) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriorityHigh, ExternalAccountVerifyResponseExternalAccountVerificationAttemptPriorityNormal:
		return true
	}
	return false
}

// The status of the verification attempt. Can be `pending_verification`,
// `verified`, `failed`, or `cancelled`.
type ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus string

const (
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusCancelled           ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus = "cancelled"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusFailed              ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus = "failed"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusPendingVerification ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus = "pending_verification"
	ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusVerified            ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus = "verified"
)

func (r ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatus) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusCancelled, ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusFailed, ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusPendingVerification, ExternalAccountVerifyResponseExternalAccountVerificationAttemptStatusVerified:
		return true
	}
	return false
}

// Either `individual` or `business`.
type ExternalAccountVerifyResponsePartyType string

const (
	ExternalAccountVerifyResponsePartyTypeBusiness   ExternalAccountVerifyResponsePartyType = "business"
	ExternalAccountVerifyResponsePartyTypeIndividual ExternalAccountVerifyResponsePartyType = "individual"
)

func (r ExternalAccountVerifyResponsePartyType) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponsePartyTypeBusiness, ExternalAccountVerifyResponsePartyTypeIndividual:
		return true
	}
	return false
}

// The type of payment that can be made to this account. Can be `ach`, `eft`, or
// `rtp`.
type ExternalAccountVerifyResponsePaymentType string

const (
	ExternalAccountVerifyResponsePaymentTypeACH         ExternalAccountVerifyResponsePaymentType = "ach"
	ExternalAccountVerifyResponsePaymentTypeAuBecs      ExternalAccountVerifyResponsePaymentType = "au_becs"
	ExternalAccountVerifyResponsePaymentTypeBacs        ExternalAccountVerifyResponsePaymentType = "bacs"
	ExternalAccountVerifyResponsePaymentTypeBase        ExternalAccountVerifyResponsePaymentType = "base"
	ExternalAccountVerifyResponsePaymentTypeBook        ExternalAccountVerifyResponsePaymentType = "book"
	ExternalAccountVerifyResponsePaymentTypeCard        ExternalAccountVerifyResponsePaymentType = "card"
	ExternalAccountVerifyResponsePaymentTypeChats       ExternalAccountVerifyResponsePaymentType = "chats"
	ExternalAccountVerifyResponsePaymentTypeCheck       ExternalAccountVerifyResponsePaymentType = "check"
	ExternalAccountVerifyResponsePaymentTypeCrossBorder ExternalAccountVerifyResponsePaymentType = "cross_border"
	ExternalAccountVerifyResponsePaymentTypeDkNets      ExternalAccountVerifyResponsePaymentType = "dk_nets"
	ExternalAccountVerifyResponsePaymentTypeEft         ExternalAccountVerifyResponsePaymentType = "eft"
	ExternalAccountVerifyResponsePaymentTypeEthereum    ExternalAccountVerifyResponsePaymentType = "ethereum"
	ExternalAccountVerifyResponsePaymentTypeGBFps       ExternalAccountVerifyResponsePaymentType = "gb_fps"
	ExternalAccountVerifyResponsePaymentTypeHuIcs       ExternalAccountVerifyResponsePaymentType = "hu_ics"
	ExternalAccountVerifyResponsePaymentTypeInterac     ExternalAccountVerifyResponsePaymentType = "interac"
	ExternalAccountVerifyResponsePaymentTypeMasav       ExternalAccountVerifyResponsePaymentType = "masav"
	ExternalAccountVerifyResponsePaymentTypeMxCcen      ExternalAccountVerifyResponsePaymentType = "mx_ccen"
	ExternalAccountVerifyResponsePaymentTypeNeft        ExternalAccountVerifyResponsePaymentType = "neft"
	ExternalAccountVerifyResponsePaymentTypeNics        ExternalAccountVerifyResponsePaymentType = "nics"
	ExternalAccountVerifyResponsePaymentTypeNzBecs      ExternalAccountVerifyResponsePaymentType = "nz_becs"
	ExternalAccountVerifyResponsePaymentTypePlElixir    ExternalAccountVerifyResponsePaymentType = "pl_elixir"
	ExternalAccountVerifyResponsePaymentTypePolygon     ExternalAccountVerifyResponsePaymentType = "polygon"
	ExternalAccountVerifyResponsePaymentTypeProvxchange ExternalAccountVerifyResponsePaymentType = "provxchange"
	ExternalAccountVerifyResponsePaymentTypeRoSent      ExternalAccountVerifyResponsePaymentType = "ro_sent"
	ExternalAccountVerifyResponsePaymentTypeRtp         ExternalAccountVerifyResponsePaymentType = "rtp"
	ExternalAccountVerifyResponsePaymentTypeSeBankgirot ExternalAccountVerifyResponsePaymentType = "se_bankgirot"
	ExternalAccountVerifyResponsePaymentTypeSen         ExternalAccountVerifyResponsePaymentType = "sen"
	ExternalAccountVerifyResponsePaymentTypeSepa        ExternalAccountVerifyResponsePaymentType = "sepa"
	ExternalAccountVerifyResponsePaymentTypeSgGiro      ExternalAccountVerifyResponsePaymentType = "sg_giro"
	ExternalAccountVerifyResponsePaymentTypeSic         ExternalAccountVerifyResponsePaymentType = "sic"
	ExternalAccountVerifyResponsePaymentTypeSignet      ExternalAccountVerifyResponsePaymentType = "signet"
	ExternalAccountVerifyResponsePaymentTypeSknbi       ExternalAccountVerifyResponsePaymentType = "sknbi"
	ExternalAccountVerifyResponsePaymentTypeSolana      ExternalAccountVerifyResponsePaymentType = "solana"
	ExternalAccountVerifyResponsePaymentTypeWire        ExternalAccountVerifyResponsePaymentType = "wire"
	ExternalAccountVerifyResponsePaymentTypeZengin      ExternalAccountVerifyResponsePaymentType = "zengin"
)

func (r ExternalAccountVerifyResponsePaymentType) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponsePaymentTypeACH, ExternalAccountVerifyResponsePaymentTypeAuBecs, ExternalAccountVerifyResponsePaymentTypeBacs, ExternalAccountVerifyResponsePaymentTypeBase, ExternalAccountVerifyResponsePaymentTypeBook, ExternalAccountVerifyResponsePaymentTypeCard, ExternalAccountVerifyResponsePaymentTypeChats, ExternalAccountVerifyResponsePaymentTypeCheck, ExternalAccountVerifyResponsePaymentTypeCrossBorder, ExternalAccountVerifyResponsePaymentTypeDkNets, ExternalAccountVerifyResponsePaymentTypeEft, ExternalAccountVerifyResponsePaymentTypeEthereum, ExternalAccountVerifyResponsePaymentTypeGBFps, ExternalAccountVerifyResponsePaymentTypeHuIcs, ExternalAccountVerifyResponsePaymentTypeInterac, ExternalAccountVerifyResponsePaymentTypeMasav, ExternalAccountVerifyResponsePaymentTypeMxCcen, ExternalAccountVerifyResponsePaymentTypeNeft, ExternalAccountVerifyResponsePaymentTypeNics, ExternalAccountVerifyResponsePaymentTypeNzBecs, ExternalAccountVerifyResponsePaymentTypePlElixir, ExternalAccountVerifyResponsePaymentTypePolygon, ExternalAccountVerifyResponsePaymentTypeProvxchange, ExternalAccountVerifyResponsePaymentTypeRoSent, ExternalAccountVerifyResponsePaymentTypeRtp, ExternalAccountVerifyResponsePaymentTypeSeBankgirot, ExternalAccountVerifyResponsePaymentTypeSen, ExternalAccountVerifyResponsePaymentTypeSepa, ExternalAccountVerifyResponsePaymentTypeSgGiro, ExternalAccountVerifyResponsePaymentTypeSic, ExternalAccountVerifyResponsePaymentTypeSignet, ExternalAccountVerifyResponsePaymentTypeSknbi, ExternalAccountVerifyResponsePaymentTypeSolana, ExternalAccountVerifyResponsePaymentTypeWire, ExternalAccountVerifyResponsePaymentTypeZengin:
		return true
	}
	return false
}

// The priority of the payment. Can be `normal` or `high`.
type ExternalAccountVerifyResponsePriority string

const (
	ExternalAccountVerifyResponsePriorityHigh   ExternalAccountVerifyResponsePriority = "high"
	ExternalAccountVerifyResponsePriorityNormal ExternalAccountVerifyResponsePriority = "normal"
)

func (r ExternalAccountVerifyResponsePriority) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponsePriorityHigh, ExternalAccountVerifyResponsePriorityNormal:
		return true
	}
	return false
}

// The status of the verification attempt. Can be `pending_verification`,
// `verified`, `failed`, or `cancelled`.
type ExternalAccountVerifyResponseStatus string

const (
	ExternalAccountVerifyResponseStatusCancelled           ExternalAccountVerifyResponseStatus = "cancelled"
	ExternalAccountVerifyResponseStatusFailed              ExternalAccountVerifyResponseStatus = "failed"
	ExternalAccountVerifyResponseStatusPendingVerification ExternalAccountVerifyResponseStatus = "pending_verification"
	ExternalAccountVerifyResponseStatusVerified            ExternalAccountVerifyResponseStatus = "verified"
)

func (r ExternalAccountVerifyResponseStatus) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponseStatusCancelled, ExternalAccountVerifyResponseStatusFailed, ExternalAccountVerifyResponseStatusPendingVerification, ExternalAccountVerifyResponseStatusVerified:
		return true
	}
	return false
}

type ExternalAccountVerifyResponseVerificationSource string

const (
	ExternalAccountVerifyResponseVerificationSourceACHPrenote    ExternalAccountVerifyResponseVerificationSource = "ach_prenote"
	ExternalAccountVerifyResponseVerificationSourceMicrodeposits ExternalAccountVerifyResponseVerificationSource = "microdeposits"
	ExternalAccountVerifyResponseVerificationSourcePlaid         ExternalAccountVerifyResponseVerificationSource = "plaid"
)

func (r ExternalAccountVerifyResponseVerificationSource) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponseVerificationSourceACHPrenote, ExternalAccountVerifyResponseVerificationSourceMicrodeposits, ExternalAccountVerifyResponseVerificationSourcePlaid:
		return true
	}
	return false
}

type ExternalAccountVerifyResponseVerificationStatus string

const (
	ExternalAccountVerifyResponseVerificationStatusPendingVerification ExternalAccountVerifyResponseVerificationStatus = "pending_verification"
	ExternalAccountVerifyResponseVerificationStatusUnverified          ExternalAccountVerifyResponseVerificationStatus = "unverified"
	ExternalAccountVerifyResponseVerificationStatusVerified            ExternalAccountVerifyResponseVerificationStatus = "verified"
)

func (r ExternalAccountVerifyResponseVerificationStatus) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyResponseVerificationStatusPendingVerification, ExternalAccountVerifyResponseVerificationStatusUnverified, ExternalAccountVerifyResponseVerificationStatusVerified:
		return true
	}
	return false
}

type ExternalAccountNewParams struct {
	CounterpartyID param.Field[string]                                  `json:"counterparty_id,required" format:"uuid"`
	AccountDetails param.Field[[]ExternalAccountNewParamsAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]               `json:"account_type"`
	ContactDetails param.Field[[]ContactDetailCreateRequestParam] `json:"contact_details"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[shared.LedgerAccountCreateRequestParam] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[shared.AddressRequestParam] `json:"party_address"`
	PartyIdentifier param.Field[string]                     `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[ExternalAccountNewParamsPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                  `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]ExternalAccountNewParamsRoutingDetail] `json:"routing_details"`
}

func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountNewParamsAccountDetail struct {
	AccountNumber     param.Field[string]                                                  `json:"account_number,required"`
	AccountNumberType param.Field[ExternalAccountNewParamsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r ExternalAccountNewParamsAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountNewParamsAccountDetailsAccountNumberType string

const (
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeAuNumber        ExternalAccountNewParamsAccountDetailsAccountNumberType = "au_number"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeBaseAddress     ExternalAccountNewParamsAccountDetailsAccountNumberType = "base_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeClabe           ExternalAccountNewParamsAccountDetailsAccountNumberType = "clabe"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeEthereumAddress ExternalAccountNewParamsAccountDetailsAccountNumberType = "ethereum_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeHkNumber        ExternalAccountNewParamsAccountDetailsAccountNumberType = "hk_number"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban            ExternalAccountNewParamsAccountDetailsAccountNumberType = "iban"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeIDNumber        ExternalAccountNewParamsAccountDetailsAccountNumberType = "id_number"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeNzNumber        ExternalAccountNewParamsAccountDetailsAccountNumberType = "nz_number"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeOther           ExternalAccountNewParamsAccountDetailsAccountNumberType = "other"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypePan             ExternalAccountNewParamsAccountDetailsAccountNumberType = "pan"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypePolygonAddress  ExternalAccountNewParamsAccountDetailsAccountNumberType = "polygon_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeSgNumber        ExternalAccountNewParamsAccountDetailsAccountNumberType = "sg_number"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeSolanaAddress   ExternalAccountNewParamsAccountDetailsAccountNumberType = "solana_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeWalletAddress   ExternalAccountNewParamsAccountDetailsAccountNumberType = "wallet_address"
)

func (r ExternalAccountNewParamsAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case ExternalAccountNewParamsAccountDetailsAccountNumberTypeAuNumber, ExternalAccountNewParamsAccountDetailsAccountNumberTypeBaseAddress, ExternalAccountNewParamsAccountDetailsAccountNumberTypeClabe, ExternalAccountNewParamsAccountDetailsAccountNumberTypeEthereumAddress, ExternalAccountNewParamsAccountDetailsAccountNumberTypeHkNumber, ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban, ExternalAccountNewParamsAccountDetailsAccountNumberTypeIDNumber, ExternalAccountNewParamsAccountDetailsAccountNumberTypeNzNumber, ExternalAccountNewParamsAccountDetailsAccountNumberTypeOther, ExternalAccountNewParamsAccountDetailsAccountNumberTypePan, ExternalAccountNewParamsAccountDetailsAccountNumberTypePolygonAddress, ExternalAccountNewParamsAccountDetailsAccountNumberTypeSgNumber, ExternalAccountNewParamsAccountDetailsAccountNumberTypeSolanaAddress, ExternalAccountNewParamsAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type ExternalAccountNewParamsPartyType string

const (
	ExternalAccountNewParamsPartyTypeBusiness   ExternalAccountNewParamsPartyType = "business"
	ExternalAccountNewParamsPartyTypeIndividual ExternalAccountNewParamsPartyType = "individual"
)

func (r ExternalAccountNewParamsPartyType) IsKnown() bool {
	switch r {
	case ExternalAccountNewParamsPartyTypeBusiness, ExternalAccountNewParamsPartyTypeIndividual:
		return true
	}
	return false
}

type ExternalAccountNewParamsRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                  `json:"routing_number,required"`
	RoutingNumberType param.Field[ExternalAccountNewParamsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[ExternalAccountNewParamsRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r ExternalAccountNewParamsRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountNewParamsRoutingDetailsRoutingNumberType string

const (
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba                     ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "aba"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAuBsb                   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "au_bsb"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeBrCodigo                ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "br_codigo"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCaCpa                   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "ca_cpa"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeChips                   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "chips"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCnaps                   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "cnaps"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeDkInterbankClearingCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeGBSortCode              ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "gb_sort_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeHkInterbankClearingCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeHuInterbankClearingCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeIDSknbiCode             ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "id_sknbi_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeIlBankCode              ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "il_bank_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeInIfsc                  ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "in_ifsc"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeJpZenginCode            ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "jp_zengin_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMyBranchCode            ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "my_branch_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMxBankIdentifier        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeNzNationalClearingCode  ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypePlNationalClearingCode  ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSgInterbankClearingCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSwift                   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "swift"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeZaNationalClearingCode  ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r ExternalAccountNewParamsRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAuBsb, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeBrCodigo, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCaCpa, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeChips, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCnaps, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeGBSortCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeIDSknbiCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeIlBankCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeInIfsc, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeJpZenginCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMyBranchCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMxBankIdentifier, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeNzNationalClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypePlNationalClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSwift, ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type ExternalAccountNewParamsRoutingDetailsPaymentType string

const (
	ExternalAccountNewParamsRoutingDetailsPaymentTypeACH         ExternalAccountNewParamsRoutingDetailsPaymentType = "ach"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeAuBecs      ExternalAccountNewParamsRoutingDetailsPaymentType = "au_becs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBacs        ExternalAccountNewParamsRoutingDetailsPaymentType = "bacs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBase        ExternalAccountNewParamsRoutingDetailsPaymentType = "base"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBook        ExternalAccountNewParamsRoutingDetailsPaymentType = "book"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCard        ExternalAccountNewParamsRoutingDetailsPaymentType = "card"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeChats       ExternalAccountNewParamsRoutingDetailsPaymentType = "chats"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCheck       ExternalAccountNewParamsRoutingDetailsPaymentType = "check"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCrossBorder ExternalAccountNewParamsRoutingDetailsPaymentType = "cross_border"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeDkNets      ExternalAccountNewParamsRoutingDetailsPaymentType = "dk_nets"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeEft         ExternalAccountNewParamsRoutingDetailsPaymentType = "eft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeEthereum    ExternalAccountNewParamsRoutingDetailsPaymentType = "ethereum"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeGBFps       ExternalAccountNewParamsRoutingDetailsPaymentType = "gb_fps"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeHuIcs       ExternalAccountNewParamsRoutingDetailsPaymentType = "hu_ics"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeInterac     ExternalAccountNewParamsRoutingDetailsPaymentType = "interac"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeMasav       ExternalAccountNewParamsRoutingDetailsPaymentType = "masav"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeMxCcen      ExternalAccountNewParamsRoutingDetailsPaymentType = "mx_ccen"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeNeft        ExternalAccountNewParamsRoutingDetailsPaymentType = "neft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeNics        ExternalAccountNewParamsRoutingDetailsPaymentType = "nics"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeNzBecs      ExternalAccountNewParamsRoutingDetailsPaymentType = "nz_becs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypePlElixir    ExternalAccountNewParamsRoutingDetailsPaymentType = "pl_elixir"
	ExternalAccountNewParamsRoutingDetailsPaymentTypePolygon     ExternalAccountNewParamsRoutingDetailsPaymentType = "polygon"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeProvxchange ExternalAccountNewParamsRoutingDetailsPaymentType = "provxchange"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeRoSent      ExternalAccountNewParamsRoutingDetailsPaymentType = "ro_sent"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeRtp         ExternalAccountNewParamsRoutingDetailsPaymentType = "rtp"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSeBankgirot ExternalAccountNewParamsRoutingDetailsPaymentType = "se_bankgirot"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSen         ExternalAccountNewParamsRoutingDetailsPaymentType = "sen"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSepa        ExternalAccountNewParamsRoutingDetailsPaymentType = "sepa"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSgGiro      ExternalAccountNewParamsRoutingDetailsPaymentType = "sg_giro"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSic         ExternalAccountNewParamsRoutingDetailsPaymentType = "sic"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSignet      ExternalAccountNewParamsRoutingDetailsPaymentType = "signet"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSknbi       ExternalAccountNewParamsRoutingDetailsPaymentType = "sknbi"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSolana      ExternalAccountNewParamsRoutingDetailsPaymentType = "solana"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeWire        ExternalAccountNewParamsRoutingDetailsPaymentType = "wire"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeZengin      ExternalAccountNewParamsRoutingDetailsPaymentType = "zengin"
)

func (r ExternalAccountNewParamsRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case ExternalAccountNewParamsRoutingDetailsPaymentTypeACH, ExternalAccountNewParamsRoutingDetailsPaymentTypeAuBecs, ExternalAccountNewParamsRoutingDetailsPaymentTypeBacs, ExternalAccountNewParamsRoutingDetailsPaymentTypeBase, ExternalAccountNewParamsRoutingDetailsPaymentTypeBook, ExternalAccountNewParamsRoutingDetailsPaymentTypeCard, ExternalAccountNewParamsRoutingDetailsPaymentTypeChats, ExternalAccountNewParamsRoutingDetailsPaymentTypeCheck, ExternalAccountNewParamsRoutingDetailsPaymentTypeCrossBorder, ExternalAccountNewParamsRoutingDetailsPaymentTypeDkNets, ExternalAccountNewParamsRoutingDetailsPaymentTypeEft, ExternalAccountNewParamsRoutingDetailsPaymentTypeEthereum, ExternalAccountNewParamsRoutingDetailsPaymentTypeGBFps, ExternalAccountNewParamsRoutingDetailsPaymentTypeHuIcs, ExternalAccountNewParamsRoutingDetailsPaymentTypeInterac, ExternalAccountNewParamsRoutingDetailsPaymentTypeMasav, ExternalAccountNewParamsRoutingDetailsPaymentTypeMxCcen, ExternalAccountNewParamsRoutingDetailsPaymentTypeNeft, ExternalAccountNewParamsRoutingDetailsPaymentTypeNics, ExternalAccountNewParamsRoutingDetailsPaymentTypeNzBecs, ExternalAccountNewParamsRoutingDetailsPaymentTypePlElixir, ExternalAccountNewParamsRoutingDetailsPaymentTypePolygon, ExternalAccountNewParamsRoutingDetailsPaymentTypeProvxchange, ExternalAccountNewParamsRoutingDetailsPaymentTypeRoSent, ExternalAccountNewParamsRoutingDetailsPaymentTypeRtp, ExternalAccountNewParamsRoutingDetailsPaymentTypeSeBankgirot, ExternalAccountNewParamsRoutingDetailsPaymentTypeSen, ExternalAccountNewParamsRoutingDetailsPaymentTypeSepa, ExternalAccountNewParamsRoutingDetailsPaymentTypeSgGiro, ExternalAccountNewParamsRoutingDetailsPaymentTypeSic, ExternalAccountNewParamsRoutingDetailsPaymentTypeSignet, ExternalAccountNewParamsRoutingDetailsPaymentTypeSknbi, ExternalAccountNewParamsRoutingDetailsPaymentTypeSolana, ExternalAccountNewParamsRoutingDetailsPaymentTypeWire, ExternalAccountNewParamsRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

type ExternalAccountUpdateParams struct {
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType] `json:"account_type"`
	CounterpartyID param.Field[string]              `json:"counterparty_id" format:"uuid"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name         param.Field[string]                     `json:"name"`
	PartyAddress param.Field[shared.AddressRequestParam] `json:"party_address"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[ExternalAccountUpdateParamsPartyType] `json:"party_type"`
}

func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `individual` or `business`.
type ExternalAccountUpdateParamsPartyType string

const (
	ExternalAccountUpdateParamsPartyTypeBusiness   ExternalAccountUpdateParamsPartyType = "business"
	ExternalAccountUpdateParamsPartyTypeIndividual ExternalAccountUpdateParamsPartyType = "individual"
)

func (r ExternalAccountUpdateParamsPartyType) IsKnown() bool {
	switch r {
	case ExternalAccountUpdateParamsPartyTypeBusiness, ExternalAccountUpdateParamsPartyTypeIndividual:
		return true
	}
	return false
}

type ExternalAccountListParams struct {
	AfterCursor    param.Field[string] `query:"after_cursor"`
	CounterpartyID param.Field[string] `query:"counterparty_id"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `query:"external_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Searches the ExternalAccount's party_name AND the Counterparty's party_name
	PartyName param.Field[string] `query:"party_name"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [ExternalAccountListParams]'s query parameters as
// `url.Values`.
func (r ExternalAccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalAccountCompleteVerificationParams struct {
	Amounts param.Field[[]int64] `json:"amounts"`
}

func (r ExternalAccountCompleteVerificationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountVerifyParams struct {
	// The ID of the internal account where the micro-deposits originate from. Both
	// credit and debit capabilities must be enabled.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Can be `ach`, `eft`, or `rtp`.
	PaymentType param.Field[ExternalAccountVerifyParamsPaymentType] `json:"payment_type,required"`
	// Defaults to the currency of the originating account.
	Currency param.Field[shared.Currency] `json:"currency"`
	// A payment type to fallback to if the original type is not valid for the
	// receiving account. Currently, this only supports falling back from RTP to ACH
	// (payment_type=rtp and fallback_type=ach)
	FallbackType param.Field[ExternalAccountVerifyParamsFallbackType] `json:"fallback_type"`
	// Either `normal` or `high`. For ACH payments, `high` represents a same-day ACH
	// transfer. This will apply to both `payment_type` and `fallback_type`.
	Priority param.Field[ExternalAccountVerifyParamsPriority] `json:"priority"`
}

func (r ExternalAccountVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Can be `ach`, `eft`, or `rtp`.
type ExternalAccountVerifyParamsPaymentType string

const (
	ExternalAccountVerifyParamsPaymentTypeACH         ExternalAccountVerifyParamsPaymentType = "ach"
	ExternalAccountVerifyParamsPaymentTypeAuBecs      ExternalAccountVerifyParamsPaymentType = "au_becs"
	ExternalAccountVerifyParamsPaymentTypeBacs        ExternalAccountVerifyParamsPaymentType = "bacs"
	ExternalAccountVerifyParamsPaymentTypeBase        ExternalAccountVerifyParamsPaymentType = "base"
	ExternalAccountVerifyParamsPaymentTypeBook        ExternalAccountVerifyParamsPaymentType = "book"
	ExternalAccountVerifyParamsPaymentTypeCard        ExternalAccountVerifyParamsPaymentType = "card"
	ExternalAccountVerifyParamsPaymentTypeChats       ExternalAccountVerifyParamsPaymentType = "chats"
	ExternalAccountVerifyParamsPaymentTypeCheck       ExternalAccountVerifyParamsPaymentType = "check"
	ExternalAccountVerifyParamsPaymentTypeCrossBorder ExternalAccountVerifyParamsPaymentType = "cross_border"
	ExternalAccountVerifyParamsPaymentTypeDkNets      ExternalAccountVerifyParamsPaymentType = "dk_nets"
	ExternalAccountVerifyParamsPaymentTypeEft         ExternalAccountVerifyParamsPaymentType = "eft"
	ExternalAccountVerifyParamsPaymentTypeEthereum    ExternalAccountVerifyParamsPaymentType = "ethereum"
	ExternalAccountVerifyParamsPaymentTypeGBFps       ExternalAccountVerifyParamsPaymentType = "gb_fps"
	ExternalAccountVerifyParamsPaymentTypeHuIcs       ExternalAccountVerifyParamsPaymentType = "hu_ics"
	ExternalAccountVerifyParamsPaymentTypeInterac     ExternalAccountVerifyParamsPaymentType = "interac"
	ExternalAccountVerifyParamsPaymentTypeMasav       ExternalAccountVerifyParamsPaymentType = "masav"
	ExternalAccountVerifyParamsPaymentTypeMxCcen      ExternalAccountVerifyParamsPaymentType = "mx_ccen"
	ExternalAccountVerifyParamsPaymentTypeNeft        ExternalAccountVerifyParamsPaymentType = "neft"
	ExternalAccountVerifyParamsPaymentTypeNics        ExternalAccountVerifyParamsPaymentType = "nics"
	ExternalAccountVerifyParamsPaymentTypeNzBecs      ExternalAccountVerifyParamsPaymentType = "nz_becs"
	ExternalAccountVerifyParamsPaymentTypePlElixir    ExternalAccountVerifyParamsPaymentType = "pl_elixir"
	ExternalAccountVerifyParamsPaymentTypePolygon     ExternalAccountVerifyParamsPaymentType = "polygon"
	ExternalAccountVerifyParamsPaymentTypeProvxchange ExternalAccountVerifyParamsPaymentType = "provxchange"
	ExternalAccountVerifyParamsPaymentTypeRoSent      ExternalAccountVerifyParamsPaymentType = "ro_sent"
	ExternalAccountVerifyParamsPaymentTypeRtp         ExternalAccountVerifyParamsPaymentType = "rtp"
	ExternalAccountVerifyParamsPaymentTypeSeBankgirot ExternalAccountVerifyParamsPaymentType = "se_bankgirot"
	ExternalAccountVerifyParamsPaymentTypeSen         ExternalAccountVerifyParamsPaymentType = "sen"
	ExternalAccountVerifyParamsPaymentTypeSepa        ExternalAccountVerifyParamsPaymentType = "sepa"
	ExternalAccountVerifyParamsPaymentTypeSgGiro      ExternalAccountVerifyParamsPaymentType = "sg_giro"
	ExternalAccountVerifyParamsPaymentTypeSic         ExternalAccountVerifyParamsPaymentType = "sic"
	ExternalAccountVerifyParamsPaymentTypeSignet      ExternalAccountVerifyParamsPaymentType = "signet"
	ExternalAccountVerifyParamsPaymentTypeSknbi       ExternalAccountVerifyParamsPaymentType = "sknbi"
	ExternalAccountVerifyParamsPaymentTypeSolana      ExternalAccountVerifyParamsPaymentType = "solana"
	ExternalAccountVerifyParamsPaymentTypeWire        ExternalAccountVerifyParamsPaymentType = "wire"
	ExternalAccountVerifyParamsPaymentTypeZengin      ExternalAccountVerifyParamsPaymentType = "zengin"
)

func (r ExternalAccountVerifyParamsPaymentType) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyParamsPaymentTypeACH, ExternalAccountVerifyParamsPaymentTypeAuBecs, ExternalAccountVerifyParamsPaymentTypeBacs, ExternalAccountVerifyParamsPaymentTypeBase, ExternalAccountVerifyParamsPaymentTypeBook, ExternalAccountVerifyParamsPaymentTypeCard, ExternalAccountVerifyParamsPaymentTypeChats, ExternalAccountVerifyParamsPaymentTypeCheck, ExternalAccountVerifyParamsPaymentTypeCrossBorder, ExternalAccountVerifyParamsPaymentTypeDkNets, ExternalAccountVerifyParamsPaymentTypeEft, ExternalAccountVerifyParamsPaymentTypeEthereum, ExternalAccountVerifyParamsPaymentTypeGBFps, ExternalAccountVerifyParamsPaymentTypeHuIcs, ExternalAccountVerifyParamsPaymentTypeInterac, ExternalAccountVerifyParamsPaymentTypeMasav, ExternalAccountVerifyParamsPaymentTypeMxCcen, ExternalAccountVerifyParamsPaymentTypeNeft, ExternalAccountVerifyParamsPaymentTypeNics, ExternalAccountVerifyParamsPaymentTypeNzBecs, ExternalAccountVerifyParamsPaymentTypePlElixir, ExternalAccountVerifyParamsPaymentTypePolygon, ExternalAccountVerifyParamsPaymentTypeProvxchange, ExternalAccountVerifyParamsPaymentTypeRoSent, ExternalAccountVerifyParamsPaymentTypeRtp, ExternalAccountVerifyParamsPaymentTypeSeBankgirot, ExternalAccountVerifyParamsPaymentTypeSen, ExternalAccountVerifyParamsPaymentTypeSepa, ExternalAccountVerifyParamsPaymentTypeSgGiro, ExternalAccountVerifyParamsPaymentTypeSic, ExternalAccountVerifyParamsPaymentTypeSignet, ExternalAccountVerifyParamsPaymentTypeSknbi, ExternalAccountVerifyParamsPaymentTypeSolana, ExternalAccountVerifyParamsPaymentTypeWire, ExternalAccountVerifyParamsPaymentTypeZengin:
		return true
	}
	return false
}

// A payment type to fallback to if the original type is not valid for the
// receiving account. Currently, this only supports falling back from RTP to ACH
// (payment_type=rtp and fallback_type=ach)
type ExternalAccountVerifyParamsFallbackType string

const (
	ExternalAccountVerifyParamsFallbackTypeACH ExternalAccountVerifyParamsFallbackType = "ach"
)

func (r ExternalAccountVerifyParamsFallbackType) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyParamsFallbackTypeACH:
		return true
	}
	return false
}

// Either `normal` or `high`. For ACH payments, `high` represents a same-day ACH
// transfer. This will apply to both `payment_type` and `fallback_type`.
type ExternalAccountVerifyParamsPriority string

const (
	ExternalAccountVerifyParamsPriorityHigh   ExternalAccountVerifyParamsPriority = "high"
	ExternalAccountVerifyParamsPriorityNormal ExternalAccountVerifyParamsPriority = "normal"
)

func (r ExternalAccountVerifyParamsPriority) IsKnown() bool {
	switch r {
	case ExternalAccountVerifyParamsPriorityHigh, ExternalAccountVerifyParamsPriorityNormal:
		return true
	}
	return false
}
