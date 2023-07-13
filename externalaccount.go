// File generated from our OpenAPI spec by Stainless.

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

// ExternalAccountService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewExternalAccountService] method instead.
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
func (r *ExternalAccountService) New(ctx context.Context, params ExternalAccountNewParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/external_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// show external account
func (r *ExternalAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// update external account
func (r *ExternalAccountService) Update(ctx context.Context, id string, body ExternalAccountUpdateParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// list external accounts
func (r *ExternalAccountService) List(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) (res *shared.Page[ExternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *ExternalAccountService) ListAutoPaging(ctx context.Context, query ExternalAccountListParams, opts ...option.RequestOption) *shared.PageAutoPager[ExternalAccount] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete external account
func (r *ExternalAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// complete verification of external account
func (r *ExternalAccountService) CompleteVerification(ctx context.Context, id string, params ExternalAccountCompleteVerificationParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s/complete_verification", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// verify external account
func (r *ExternalAccountService) Verify(ctx context.Context, id string, params ExternalAccountVerifyParams, opts ...option.RequestOption) (res *ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s/verify", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ExternalAccount struct {
	ID             string          `json:"id,required" format:"uuid"`
	AccountDetails []AccountDetail `json:"account_details,required"`
	// Can be `checking`, `savings` or `other`.
	AccountType    ExternalAccountType            `json:"account_type,required"`
	ContactDetails []ExternalAccountContactDetail `json:"contact_details,required"`
	CounterpartyID string                         `json:"counterparty_id,required,nullable" format:"uuid"`
	CreatedAt      time.Time                      `json:"created_at,required" format:"date-time"`
	DiscardedAt    time.Time                      `json:"discarded_at,required,nullable" format:"date-time"`
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
	PartyAddress ExternalAccountPartyAddress `json:"party_address,required,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name,required"`
	// Either `individual` or `business`.
	PartyType          ExternalAccountPartyType          `json:"party_type,required,nullable"`
	RoutingDetails     []RoutingDetail                   `json:"routing_details,required"`
	UpdatedAt          time.Time                         `json:"updated_at,required" format:"date-time"`
	VerificationStatus ExternalAccountVerificationStatus `json:"verification_status,required"`
	JSON               externalAccountJSON
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
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Can be `checking`, `savings` or `other`.
type ExternalAccountType string

const (
	ExternalAccountTypeCash        ExternalAccountType = "cash"
	ExternalAccountTypeChecking    ExternalAccountType = "checking"
	ExternalAccountTypeLoan        ExternalAccountType = "loan"
	ExternalAccountTypeNonResident ExternalAccountType = "non_resident"
	ExternalAccountTypeOther       ExternalAccountType = "other"
	ExternalAccountTypeOverdraft   ExternalAccountType = "overdraft"
	ExternalAccountTypeSavings     ExternalAccountType = "savings"
)

type ExternalAccountContactDetail struct {
	ID                    string                                             `json:"id,required" format:"uuid"`
	ContactIdentifier     string                                             `json:"contact_identifier,required"`
	ContactIdentifierType ExternalAccountContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	CreatedAt             time.Time                                          `json:"created_at,required" format:"date-time"`
	DiscardedAt           time.Time                                          `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	Object    string    `json:"object,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	JSON      externalAccountContactDetailJSON
}

// externalAccountContactDetailJSON contains the JSON metadata for the struct
// [ExternalAccountContactDetail]
type externalAccountContactDetailJSON struct {
	ID                    apijson.Field
	ContactIdentifier     apijson.Field
	ContactIdentifierType apijson.Field
	CreatedAt             apijson.Field
	DiscardedAt           apijson.Field
	LiveMode              apijson.Field
	Object                apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ExternalAccountContactDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalAccountContactDetailsContactIdentifierType string

const (
	ExternalAccountContactDetailsContactIdentifierTypeEmail       ExternalAccountContactDetailsContactIdentifierType = "email"
	ExternalAccountContactDetailsContactIdentifierTypePhoneNumber ExternalAccountContactDetailsContactIdentifierType = "phone_number"
	ExternalAccountContactDetailsContactIdentifierTypeWebsite     ExternalAccountContactDetailsContactIdentifierType = "website"
)

// The address associated with the owner or `null`.
type ExternalAccountPartyAddress struct {
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
	Region    string    `json:"region,required,nullable"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	JSON      externalAccountPartyAddressJSON
}

// externalAccountPartyAddressJSON contains the JSON metadata for the struct
// [ExternalAccountPartyAddress]
type externalAccountPartyAddressJSON struct {
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

func (r *ExternalAccountPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Either `individual` or `business`.
type ExternalAccountPartyType string

const (
	ExternalAccountPartyTypeBusiness   ExternalAccountPartyType = "business"
	ExternalAccountPartyTypeIndividual ExternalAccountPartyType = "individual"
)

type ExternalAccountVerificationStatus string

const (
	ExternalAccountVerificationStatusPendingVerification ExternalAccountVerificationStatus = "pending_verification"
	ExternalAccountVerificationStatusUnverified          ExternalAccountVerificationStatus = "unverified"
	ExternalAccountVerificationStatusVerified            ExternalAccountVerificationStatus = "verified"
)

type ExternalAccountNewParams struct {
	CounterpartyID param.Field[string]                                  `json:"counterparty_id,required" format:"uuid"`
	AccountDetails param.Field[[]ExternalAccountNewParamsAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]                     `json:"account_type"`
	ContactDetails param.Field[[]ExternalAccountNewParamsContactDetail] `json:"contact_details"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[ExternalAccountNewParamsLedgerAccount] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[ExternalAccountNewParamsPartyAddress] `json:"party_address"`
	PartyIdentifier param.Field[string]                               `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[ExternalAccountNewParamsPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                  `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]ExternalAccountNewParamsRoutingDetail] `json:"routing_details"`
	IdempotencyKey      param.Field[string]                                  `header:"Idempotency-Key"`
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
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeIban          ExternalAccountNewParamsAccountDetailsAccountNumberType = "iban"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeClabe         ExternalAccountNewParamsAccountDetailsAccountNumberType = "clabe"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeWalletAddress ExternalAccountNewParamsAccountDetailsAccountNumberType = "wallet_address"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypePan           ExternalAccountNewParamsAccountDetailsAccountNumberType = "pan"
	ExternalAccountNewParamsAccountDetailsAccountNumberTypeOther         ExternalAccountNewParamsAccountDetailsAccountNumberType = "other"
)

type ExternalAccountNewParamsContactDetail struct {
	ContactIdentifier     param.Field[string]                                                      `json:"contact_identifier"`
	ContactIdentifierType param.Field[ExternalAccountNewParamsContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r ExternalAccountNewParamsContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountNewParamsContactDetailsContactIdentifierType string

const (
	ExternalAccountNewParamsContactDetailsContactIdentifierTypeEmail       ExternalAccountNewParamsContactDetailsContactIdentifierType = "email"
	ExternalAccountNewParamsContactDetailsContactIdentifierTypePhoneNumber ExternalAccountNewParamsContactDetailsContactIdentifierType = "phone_number"
	ExternalAccountNewParamsContactDetailsContactIdentifierTypeWebsite     ExternalAccountNewParamsContactDetailsContactIdentifierType = "website"
)

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type ExternalAccountNewParamsLedgerAccount struct {
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[ExternalAccountNewParamsLedgerAccountNormalBalance] `json:"normal_balance,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[ExternalAccountNewParamsLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r ExternalAccountNewParamsLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The normal balance of the ledger account.
type ExternalAccountNewParamsLedgerAccountNormalBalance string

const (
	ExternalAccountNewParamsLedgerAccountNormalBalanceCredit ExternalAccountNewParamsLedgerAccountNormalBalance = "credit"
	ExternalAccountNewParamsLedgerAccountNormalBalanceDebit  ExternalAccountNewParamsLedgerAccountNormalBalance = "debit"
)

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type ExternalAccountNewParamsLedgerAccountLedgerableType string

const (
	ExternalAccountNewParamsLedgerAccountLedgerableTypeExternalAccount ExternalAccountNewParamsLedgerAccountLedgerableType = "external_account"
	ExternalAccountNewParamsLedgerAccountLedgerableTypeInternalAccount ExternalAccountNewParamsLedgerAccountLedgerableType = "internal_account"
)

// Required if receiving wire payments.
type ExternalAccountNewParamsPartyAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
	Line1   param.Field[string] `json:"line1"`
	Line2   param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Region or State.
	Region param.Field[string] `json:"region"`
}

func (r ExternalAccountNewParamsPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `individual` or `business`.
type ExternalAccountNewParamsPartyType string

const (
	ExternalAccountNewParamsPartyTypeBusiness   ExternalAccountNewParamsPartyType = "business"
	ExternalAccountNewParamsPartyTypeIndividual ExternalAccountNewParamsPartyType = "individual"
)

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
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAba          ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "aba"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeAuBsb        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "au_bsb"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeBrCodigo     ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "br_codigo"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCaCpa        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "ca_cpa"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeChips        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "chips"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeCnaps        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "cnaps"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeGBSortCode   ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "gb_sort_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeInIfsc       ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "in_ifsc"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeMyBranchCode ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "my_branch_code"
	ExternalAccountNewParamsRoutingDetailsRoutingNumberTypeSwift        ExternalAccountNewParamsRoutingDetailsRoutingNumberType = "swift"
)

type ExternalAccountNewParamsRoutingDetailsPaymentType string

const (
	ExternalAccountNewParamsRoutingDetailsPaymentTypeACH         ExternalAccountNewParamsRoutingDetailsPaymentType = "ach"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeAuBecs      ExternalAccountNewParamsRoutingDetailsPaymentType = "au_becs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBacs        ExternalAccountNewParamsRoutingDetailsPaymentType = "bacs"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeBook        ExternalAccountNewParamsRoutingDetailsPaymentType = "book"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCard        ExternalAccountNewParamsRoutingDetailsPaymentType = "card"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCheck       ExternalAccountNewParamsRoutingDetailsPaymentType = "check"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeEft         ExternalAccountNewParamsRoutingDetailsPaymentType = "eft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeCrossBorder ExternalAccountNewParamsRoutingDetailsPaymentType = "cross_border"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeInterac     ExternalAccountNewParamsRoutingDetailsPaymentType = "interac"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeMasav       ExternalAccountNewParamsRoutingDetailsPaymentType = "masav"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeNeft        ExternalAccountNewParamsRoutingDetailsPaymentType = "neft"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeProvxchange ExternalAccountNewParamsRoutingDetailsPaymentType = "provxchange"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeRtp         ExternalAccountNewParamsRoutingDetailsPaymentType = "rtp"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSen         ExternalAccountNewParamsRoutingDetailsPaymentType = "sen"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSepa        ExternalAccountNewParamsRoutingDetailsPaymentType = "sepa"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeSignet      ExternalAccountNewParamsRoutingDetailsPaymentType = "signet"
	ExternalAccountNewParamsRoutingDetailsPaymentTypeWire        ExternalAccountNewParamsRoutingDetailsPaymentType = "wire"
)

type ExternalAccountUpdateParams struct {
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType] `json:"account_type"`
	CounterpartyID param.Field[string]              `json:"counterparty_id" format:"uuid"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name         param.Field[string]                                  `json:"name"`
	PartyAddress param.Field[ExternalAccountUpdateParamsPartyAddress] `json:"party_address"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[ExternalAccountUpdateParamsPartyType] `json:"party_type"`
}

func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountUpdateParamsPartyAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
	Line1   param.Field[string] `json:"line1"`
	Line2   param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Region or State.
	Region param.Field[string] `json:"region"`
}

func (r ExternalAccountUpdateParamsPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `individual` or `business`.
type ExternalAccountUpdateParamsPartyType string

const (
	ExternalAccountUpdateParamsPartyTypeBusiness   ExternalAccountUpdateParamsPartyType = "business"
	ExternalAccountUpdateParamsPartyTypeIndividual ExternalAccountUpdateParamsPartyType = "individual"
)

type ExternalAccountListParams struct {
	AfterCursor    param.Field[string] `query:"after_cursor"`
	CounterpartyID param.Field[string] `query:"counterparty_id"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalAccountCompleteVerificationParams struct {
	Amounts        param.Field[[]int64] `json:"amounts"`
	IdempotencyKey param.Field[string]  `header:"Idempotency-Key"`
}

func (r ExternalAccountCompleteVerificationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountVerifyParams struct {
	// The ID of the internal account where the micro-deposits originate from. Both
	// credit and debit capabilities must be enabled.
	OriginatingAccountID param.Field[string] `json:"originating_account_id,required" format:"uuid"`
	// Both ach and eft are supported payment types.
	PaymentType param.Field[ExternalAccountVerifyParamsPaymentType] `json:"payment_type,required"`
	// Defaults to the currency of the originating account.
	Currency       param.Field[shared.Currency] `json:"currency"`
	IdempotencyKey param.Field[string]          `header:"Idempotency-Key"`
}

func (r ExternalAccountVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Both ach and eft are supported payment types.
type ExternalAccountVerifyParamsPaymentType string

const (
	ExternalAccountVerifyParamsPaymentTypeACH         ExternalAccountVerifyParamsPaymentType = "ach"
	ExternalAccountVerifyParamsPaymentTypeAuBecs      ExternalAccountVerifyParamsPaymentType = "au_becs"
	ExternalAccountVerifyParamsPaymentTypeBacs        ExternalAccountVerifyParamsPaymentType = "bacs"
	ExternalAccountVerifyParamsPaymentTypeBook        ExternalAccountVerifyParamsPaymentType = "book"
	ExternalAccountVerifyParamsPaymentTypeCard        ExternalAccountVerifyParamsPaymentType = "card"
	ExternalAccountVerifyParamsPaymentTypeCheck       ExternalAccountVerifyParamsPaymentType = "check"
	ExternalAccountVerifyParamsPaymentTypeCrossBorder ExternalAccountVerifyParamsPaymentType = "cross_border"
	ExternalAccountVerifyParamsPaymentTypeEft         ExternalAccountVerifyParamsPaymentType = "eft"
	ExternalAccountVerifyParamsPaymentTypeInterac     ExternalAccountVerifyParamsPaymentType = "interac"
	ExternalAccountVerifyParamsPaymentTypeMasav       ExternalAccountVerifyParamsPaymentType = "masav"
	ExternalAccountVerifyParamsPaymentTypeNeft        ExternalAccountVerifyParamsPaymentType = "neft"
	ExternalAccountVerifyParamsPaymentTypeProvxchange ExternalAccountVerifyParamsPaymentType = "provxchange"
	ExternalAccountVerifyParamsPaymentTypeRtp         ExternalAccountVerifyParamsPaymentType = "rtp"
	ExternalAccountVerifyParamsPaymentTypeSen         ExternalAccountVerifyParamsPaymentType = "sen"
	ExternalAccountVerifyParamsPaymentTypeSepa        ExternalAccountVerifyParamsPaymentType = "sepa"
	ExternalAccountVerifyParamsPaymentTypeSignet      ExternalAccountVerifyParamsPaymentType = "signet"
	ExternalAccountVerifyParamsPaymentTypeWire        ExternalAccountVerifyParamsPaymentType = "wire"
)
