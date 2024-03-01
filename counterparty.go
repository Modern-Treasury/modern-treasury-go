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

// CounterpartyService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewCounterpartyService]
// method instead.
type CounterpartyService struct {
	Options []option.RequestOption
}

// NewCounterpartyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCounterpartyService(opts ...option.RequestOption) (r *CounterpartyService) {
	r = &CounterpartyService{}
	r.Options = opts
	return
}

// Create a new counterparty.
func (r *CounterpartyService) New(ctx context.Context, body CounterpartyNewParams, opts ...option.RequestOption) (res *Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/counterparties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single counterparty.
func (r *CounterpartyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a given counterparty with new information.
func (r *CounterpartyService) Update(ctx context.Context, id string, body CounterpartyUpdateParams, opts ...option.RequestOption) (res *Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of all counterparties.
func (r *CounterpartyService) List(ctx context.Context, query CounterpartyListParams, opts ...option.RequestOption) (res *shared.Page[Counterparty], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/counterparties"
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

// Get a paginated list of all counterparties.
func (r *CounterpartyService) ListAutoPaging(ctx context.Context, query CounterpartyListParams, opts ...option.RequestOption) *shared.PageAutoPager[Counterparty] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a given counterparty.
func (r *CounterpartyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Send an email requesting account details.
func (r *CounterpartyService) CollectAccount(ctx context.Context, id string, body CounterpartyCollectAccountParams, opts ...option.RequestOption) (res *CounterpartyCollectAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/counterparties/%s/collect_account", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Counterparty struct {
	ID string `json:"id,required" format:"uuid"`
	// The accounts for this counterparty.
	Accounts    []CounterpartyAccount `json:"accounts,required"`
	CreatedAt   time.Time             `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time             `json:"discarded_at,required,nullable" format:"date-time"`
	// The counterparty's email.
	Email string `json:"email,required,nullable" format:"email"`
	// The id of the legal entity.
	LegalEntityID string `json:"legal_entity_id,required,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// A human friendly name for this counterparty.
	Name   string `json:"name,required,nullable"`
	Object string `json:"object,required"`
	// Send an email to the counterparty whenever an associated payment order is sent
	// to the bank.
	SendRemittanceAdvice bool      `json:"send_remittance_advice,required"`
	UpdatedAt            time.Time `json:"updated_at,required" format:"date-time"`
	// The verification status of the counterparty.
	VerificationStatus CounterpartyVerificationStatus `json:"verification_status,required"`
	JSON               counterpartyJSON               `json:"-"`
}

// counterpartyJSON contains the JSON metadata for the struct [Counterparty]
type counterpartyJSON struct {
	ID                   apijson.Field
	Accounts             apijson.Field
	CreatedAt            apijson.Field
	DiscardedAt          apijson.Field
	Email                apijson.Field
	LegalEntityID        apijson.Field
	LiveMode             apijson.Field
	Metadata             apijson.Field
	Name                 apijson.Field
	Object               apijson.Field
	SendRemittanceAdvice apijson.Field
	UpdatedAt            apijson.Field
	VerificationStatus   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Counterparty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccount struct {
	ID             string          `json:"id" format:"uuid"`
	AccountDetails []AccountDetail `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    ExternalAccountType                 `json:"account_type"`
	ContactDetails []CounterpartyAccountsContactDetail `json:"contact_details"`
	CreatedAt      time.Time                           `json:"created_at" format:"date-time"`
	DiscardedAt    time.Time                           `json:"discarded_at,nullable" format:"date-time"`
	// If the external account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID string `json:"ledger_account_id,nullable" format:"uuid"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name   string `json:"name,nullable"`
	Object string `json:"object"`
	// The address associated with the owner or `null`.
	PartyAddress CounterpartyAccountsPartyAddress `json:"party_address,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name"`
	// Either `individual` or `business`.
	PartyType          CounterpartyAccountsPartyType          `json:"party_type,nullable"`
	RoutingDetails     []RoutingDetail                        `json:"routing_details"`
	UpdatedAt          time.Time                              `json:"updated_at" format:"date-time"`
	VerificationStatus CounterpartyAccountsVerificationStatus `json:"verification_status"`
	JSON               counterpartyAccountJSON                `json:"-"`
}

// counterpartyAccountJSON contains the JSON metadata for the struct
// [CounterpartyAccount]
type counterpartyAccountJSON struct {
	ID                 apijson.Field
	AccountDetails     apijson.Field
	AccountType        apijson.Field
	ContactDetails     apijson.Field
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

func (r *CounterpartyAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsContactDetail struct {
	ID                    string                                                  `json:"id,required" format:"uuid"`
	ContactIdentifier     string                                                  `json:"contact_identifier,required"`
	ContactIdentifierType CounterpartyAccountsContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	CreatedAt             time.Time                                               `json:"created_at,required" format:"date-time"`
	DiscardedAt           time.Time                                               `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                                  `json:"live_mode,required"`
	Object    string                                `json:"object,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON      counterpartyAccountsContactDetailJSON `json:"-"`
}

// counterpartyAccountsContactDetailJSON contains the JSON metadata for the struct
// [CounterpartyAccountsContactDetail]
type counterpartyAccountsContactDetailJSON struct {
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

func (r *CounterpartyAccountsContactDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsContactDetailsContactIdentifierType string

const (
	CounterpartyAccountsContactDetailsContactIdentifierTypeEmail       CounterpartyAccountsContactDetailsContactIdentifierType = "email"
	CounterpartyAccountsContactDetailsContactIdentifierTypePhoneNumber CounterpartyAccountsContactDetailsContactIdentifierType = "phone_number"
	CounterpartyAccountsContactDetailsContactIdentifierTypeWebsite     CounterpartyAccountsContactDetailsContactIdentifierType = "website"
)

// The address associated with the owner or `null`.
type CounterpartyAccountsPartyAddress struct {
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
	Region    string                               `json:"region,required,nullable"`
	UpdatedAt time.Time                            `json:"updated_at,required" format:"date-time"`
	JSON      counterpartyAccountsPartyAddressJSON `json:"-"`
}

// counterpartyAccountsPartyAddressJSON contains the JSON metadata for the struct
// [CounterpartyAccountsPartyAddress]
type counterpartyAccountsPartyAddressJSON struct {
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

func (r *CounterpartyAccountsPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Either `individual` or `business`.
type CounterpartyAccountsPartyType string

const (
	CounterpartyAccountsPartyTypeBusiness   CounterpartyAccountsPartyType = "business"
	CounterpartyAccountsPartyTypeIndividual CounterpartyAccountsPartyType = "individual"
)

type CounterpartyAccountsVerificationStatus string

const (
	CounterpartyAccountsVerificationStatusPendingVerification CounterpartyAccountsVerificationStatus = "pending_verification"
	CounterpartyAccountsVerificationStatusUnverified          CounterpartyAccountsVerificationStatus = "unverified"
	CounterpartyAccountsVerificationStatusVerified            CounterpartyAccountsVerificationStatus = "verified"
)

// The verification status of the counterparty.
type CounterpartyVerificationStatus string

const (
	CounterpartyVerificationStatusDenied        CounterpartyVerificationStatus = "denied"
	CounterpartyVerificationStatusNeedsApproval CounterpartyVerificationStatus = "needs_approval"
	CounterpartyVerificationStatusUnverified    CounterpartyVerificationStatus = "unverified"
	CounterpartyVerificationStatusVerified      CounterpartyVerificationStatus = "verified"
)

type CounterpartyCollectAccountResponse struct {
	// The id of the existing counterparty.
	ID string `json:"id,required"`
	// This is the link to the secure Modern Treasury form. By default, Modern Treasury
	// will send an email to your counterparty that includes a link to this form.
	// However, if `send_email` is passed as `false` in the body then Modern Treasury
	// will not send the email and you can send it to the counterparty directly.
	FormLink string `json:"form_link,required" format:"uri"`
	// This field will be `true` if an email requesting account details has already
	// been sent to this counterparty.
	IsResend bool                                   `json:"is_resend,required"`
	JSON     counterpartyCollectAccountResponseJSON `json:"-"`
}

// counterpartyCollectAccountResponseJSON contains the JSON metadata for the struct
// [CounterpartyCollectAccountResponse]
type counterpartyCollectAccountResponseJSON struct {
	ID          apijson.Field
	FormLink    apijson.Field
	IsResend    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CounterpartyCollectAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyNewParams struct {
	// A human friendly name for this counterparty.
	Name       param.Field[string]                          `json:"name,required"`
	Accounting param.Field[CounterpartyNewParamsAccounting] `json:"accounting"`
	// The accounts for this counterparty.
	Accounts param.Field[[]CounterpartyNewParamsAccount] `json:"accounts"`
	// The counterparty's email.
	Email param.Field[string] `json:"email" format:"email"`
	// An optional type to auto-sync the counterparty to your ledger. Either `customer`
	// or `vendor`.
	LedgerType  param.Field[CounterpartyNewParamsLedgerType]  `json:"ledger_type"`
	LegalEntity param.Field[CounterpartyNewParamsLegalEntity] `json:"legal_entity"`
	// The id of the legal entity.
	LegalEntityID param.Field[string] `json:"legal_entity_id" format:"uuid"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Send an email to the counterparty whenever an associated payment order is sent
	// to the bank.
	SendRemittanceAdvice param.Field[bool] `json:"send_remittance_advice"`
	// Either a valid SSN or EIN.
	TaxpayerIdentifier param.Field[string] `json:"taxpayer_identifier"`
	// The verification status of the counterparty.
	VerificationStatus param.Field[CounterpartyNewParamsVerificationStatus] `json:"verification_status"`
}

func (r CounterpartyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsAccounting struct {
	// An optional type to auto-sync the counterparty to your ledger. Either `customer`
	// or `vendor`.
	Type param.Field[CounterpartyNewParamsAccountingType] `json:"type"`
}

func (r CounterpartyNewParamsAccounting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An optional type to auto-sync the counterparty to your ledger. Either `customer`
// or `vendor`.
type CounterpartyNewParamsAccountingType string

const (
	CounterpartyNewParamsAccountingTypeCustomer CounterpartyNewParamsAccountingType = "customer"
	CounterpartyNewParamsAccountingTypeVendor   CounterpartyNewParamsAccountingType = "vendor"
)

type CounterpartyNewParamsAccount struct {
	AccountDetails param.Field[[]CounterpartyNewParamsAccountsAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]                          `json:"account_type"`
	ContactDetails param.Field[[]CounterpartyNewParamsAccountsContactDetail] `json:"contact_details"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[CounterpartyNewParamsAccountsLedgerAccount] `json:"ledger_account"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name param.Field[string] `json:"name"`
	// Required if receiving wire payments.
	PartyAddress    param.Field[CounterpartyNewParamsAccountsPartyAddress] `json:"party_address"`
	PartyIdentifier param.Field[string]                                    `json:"party_identifier"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName param.Field[string] `json:"party_name"`
	// Either `individual` or `business`.
	PartyType param.Field[CounterpartyNewParamsAccountsPartyType] `json:"party_type"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                       `json:"plaid_processor_token"`
	RoutingDetails      param.Field[[]CounterpartyNewParamsAccountsRoutingDetail] `json:"routing_details"`
}

func (r CounterpartyNewParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsAccountsAccountDetail struct {
	AccountNumber     param.Field[string]                                                       `json:"account_number,required"`
	AccountNumberType param.Field[CounterpartyNewParamsAccountsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

func (r CounterpartyNewParamsAccountsAccountDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsAccountsAccountDetailsAccountNumberType string

const (
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban          CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "iban"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeHkNumber      CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "hk_number"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeClabe         CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "clabe"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeWalletAddress CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "wallet_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePan           CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "pan"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeOther         CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "other"
)

type CounterpartyNewParamsAccountsContactDetail struct {
	ContactIdentifier     param.Field[string]                                                           `json:"contact_identifier"`
	ContactIdentifierType param.Field[CounterpartyNewParamsAccountsContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

func (r CounterpartyNewParamsAccountsContactDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsAccountsContactDetailsContactIdentifierType string

const (
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail       CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "email"
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypePhoneNumber CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "phone_number"
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeWebsite     CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "website"
)

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://docs.moderntreasury.com/docs/linking-to-other-modern-treasury-objects
// for more details.
type CounterpartyNewParamsAccountsLedgerAccount struct {
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[shared.TransactionDirection] `json:"normal_balance,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The array of ledger account category ids that this ledger account should be a
	// child of.
	LedgerAccountCategoryIDs param.Field[[]string] `json:"ledger_account_category_ids" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[CounterpartyNewParamsAccountsLedgerAccountLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r CounterpartyNewParamsAccountsLedgerAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type CounterpartyNewParamsAccountsLedgerAccountLedgerableType string

const (
	CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeExternalAccount CounterpartyNewParamsAccountsLedgerAccountLedgerableType = "external_account"
	CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeInternalAccount CounterpartyNewParamsAccountsLedgerAccountLedgerableType = "internal_account"
	CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeVirtualAccount  CounterpartyNewParamsAccountsLedgerAccountLedgerableType = "virtual_account"
)

// Required if receiving wire payments.
type CounterpartyNewParamsAccountsPartyAddress struct {
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

func (r CounterpartyNewParamsAccountsPartyAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Either `individual` or `business`.
type CounterpartyNewParamsAccountsPartyType string

const (
	CounterpartyNewParamsAccountsPartyTypeBusiness   CounterpartyNewParamsAccountsPartyType = "business"
	CounterpartyNewParamsAccountsPartyTypeIndividual CounterpartyNewParamsAccountsPartyType = "individual"
)

type CounterpartyNewParamsAccountsRoutingDetail struct {
	RoutingNumber     param.Field[string]                                                       `json:"routing_number,required"`
	RoutingNumberType param.Field[CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[CounterpartyNewParamsAccountsRoutingDetailsPaymentType]       `json:"payment_type"`
}

func (r CounterpartyNewParamsAccountsRoutingDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba                     CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "aba"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAuBsb                   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "au_bsb"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeBrCodigo                CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "br_codigo"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCaCpa                   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "ca_cpa"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeChips                   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "chips"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCnaps                   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "cnaps"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeDkInterbankClearingCode CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "dk_interbank_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeGBSortCode              CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "gb_sort_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeHkInterbankClearingCode CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "hk_interbank_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeHuInterbankClearingCode CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "hu_interbank_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeIDSknbiCode             CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "id_sknbi_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeInIfsc                  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "in_ifsc"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeJpZenginCode            CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "jp_zengin_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMyBranchCode            CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "my_branch_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMxBankIdentifier        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeNzNationalClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypePlNationalClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSwift                   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "swift"
)

type CounterpartyNewParamsAccountsRoutingDetailsPaymentType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ach"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeAuBecs      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "au_becs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBacs        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "bacs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBook        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "book"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCard        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "card"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeChats       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "chats"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCheck       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "check"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCrossBorder CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "cross_border"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeDkNets      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "dk_nets"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEft         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "eft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeHuIcs       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "hu_ics"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeInterac     CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "interac"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMasav       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "masav"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMxCcen      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "mx_ccen"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNeft        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "neft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNics        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "nics"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNzBecs      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "nz_becs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypePlElixir    CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "pl_elixir"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeProvxchange CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "provxchange"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRoSent      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ro_sent"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRtp         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "rtp"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSgGiro      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sg_giro"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSeBankgirot CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "se_bankgirot"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSen         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sen"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSepa        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sepa"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSic         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sic"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSignet      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "signet"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSknbi       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sknbi"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeWire        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "wire"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeZengin      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "zengin"
)

// An optional type to auto-sync the counterparty to your ledger. Either `customer`
// or `vendor`.
type CounterpartyNewParamsLedgerType string

const (
	CounterpartyNewParamsLedgerTypeCustomer CounterpartyNewParamsLedgerType = "customer"
	CounterpartyNewParamsLedgerTypeVendor   CounterpartyNewParamsLedgerType = "vendor"
)

type CounterpartyNewParamsLegalEntity struct {
	// The type of legal entity.
	LegalEntityType param.Field[CounterpartyNewParamsLegalEntityLegalEntityType] `json:"legal_entity_type,required"`
	// A list of addresses for the entity.
	Addresses param.Field[[]CounterpartyNewParamsLegalEntityAddress] `json:"addresses"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// An individual's data of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// A list of identifications for the legal entity.
	Identifications param.Field[[]CounterpartyNewParamsLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The business's legal structure.
	LegalStructure param.Field[CounterpartyNewParamsLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                             `json:"metadata"`
	PhoneNumbers param.Field[[]CounterpartyNewParamsLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r CounterpartyNewParamsLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of legal entity.
type CounterpartyNewParamsLegalEntityLegalEntityType string

const (
	CounterpartyNewParamsLegalEntityLegalEntityTypeBusiness   CounterpartyNewParamsLegalEntityLegalEntityType = "business"
	CounterpartyNewParamsLegalEntityLegalEntityTypeIndividual CounterpartyNewParamsLegalEntityLegalEntityType = "individual"
)

type CounterpartyNewParamsLegalEntityAddress struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
	Line1   param.Field[string] `json:"line1,required"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	// The types of this address.
	AddressTypes param.Field[[]CounterpartyNewParamsLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                 `json:"line2"`
}

func (r CounterpartyNewParamsLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsLegalEntityAddressesAddressType string

const (
	CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness    CounterpartyNewParamsLegalEntityAddressesAddressType = "business"
	CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing     CounterpartyNewParamsLegalEntityAddressesAddressType = "mailing"
	CounterpartyNewParamsLegalEntityAddressesAddressTypeOther       CounterpartyNewParamsLegalEntityAddressesAddressType = "other"
	CounterpartyNewParamsLegalEntityAddressesAddressTypePoBox       CounterpartyNewParamsLegalEntityAddressesAddressType = "po_box"
	CounterpartyNewParamsLegalEntityAddressesAddressTypeResidential CounterpartyNewParamsLegalEntityAddressesAddressType = "residential"
)

type CounterpartyNewParamsLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[CounterpartyNewParamsLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r CounterpartyNewParamsLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type CounterpartyNewParamsLegalEntityIdentificationsIDType string

const (
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil    CounterpartyNewParamsLegalEntityIdentificationsIDType = "ar_cuil"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuit    CounterpartyNewParamsLegalEntityIdentificationsIDType = "ar_cuit"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeBrCnpj    CounterpartyNewParamsLegalEntityIdentificationsIDType = "br_cnpj"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeBrCpf     CounterpartyNewParamsLegalEntityIdentificationsIDType = "br_cpf"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeClNut     CounterpartyNewParamsLegalEntityIdentificationsIDType = "cl_nut"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeCoCedulas CounterpartyNewParamsLegalEntityIdentificationsIDType = "co_cedulas"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeCoNit     CounterpartyNewParamsLegalEntityIdentificationsIDType = "co_nit"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeHnID      CounterpartyNewParamsLegalEntityIdentificationsIDType = "hn_id"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeHnRtn     CounterpartyNewParamsLegalEntityIdentificationsIDType = "hn_rtn"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypePassport  CounterpartyNewParamsLegalEntityIdentificationsIDType = "passport"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeUsEin     CounterpartyNewParamsLegalEntityIdentificationsIDType = "us_ein"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeUsItin    CounterpartyNewParamsLegalEntityIdentificationsIDType = "us_itin"
	CounterpartyNewParamsLegalEntityIdentificationsIDTypeUsSsn     CounterpartyNewParamsLegalEntityIdentificationsIDType = "us_ssn"
)

// The business's legal structure.
type CounterpartyNewParamsLegalEntityLegalStructure string

const (
	CounterpartyNewParamsLegalEntityLegalStructureCorporation        CounterpartyNewParamsLegalEntityLegalStructure = "corporation"
	CounterpartyNewParamsLegalEntityLegalStructureLlc                CounterpartyNewParamsLegalEntityLegalStructure = "llc"
	CounterpartyNewParamsLegalEntityLegalStructureNonProfit          CounterpartyNewParamsLegalEntityLegalStructure = "non_profit"
	CounterpartyNewParamsLegalEntityLegalStructurePartnership        CounterpartyNewParamsLegalEntityLegalStructure = "partnership"
	CounterpartyNewParamsLegalEntityLegalStructureSoleProprietorship CounterpartyNewParamsLegalEntityLegalStructure = "sole_proprietorship"
	CounterpartyNewParamsLegalEntityLegalStructureTrust              CounterpartyNewParamsLegalEntityLegalStructure = "trust"
)

// A list of phone numbers in E.164 format.
type CounterpartyNewParamsLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r CounterpartyNewParamsLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The verification status of the counterparty.
type CounterpartyNewParamsVerificationStatus string

const (
	CounterpartyNewParamsVerificationStatusDenied        CounterpartyNewParamsVerificationStatus = "denied"
	CounterpartyNewParamsVerificationStatusNeedsApproval CounterpartyNewParamsVerificationStatus = "needs_approval"
	CounterpartyNewParamsVerificationStatusUnverified    CounterpartyNewParamsVerificationStatus = "unverified"
	CounterpartyNewParamsVerificationStatusVerified      CounterpartyNewParamsVerificationStatus = "verified"
)

type CounterpartyUpdateParams struct {
	// A new email for the counterparty.
	Email param.Field[string] `json:"email" format:"email"`
	// The id of the legal entity.
	LegalEntityID param.Field[string] `json:"legal_entity_id" format:"uuid"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A new name for the counterparty. Will only update if passed.
	Name param.Field[string] `json:"name"`
	// If this is `true`, Modern Treasury will send an email to the counterparty
	// whenever an associated payment order is sent to the bank.
	SendRemittanceAdvice param.Field[bool] `json:"send_remittance_advice"`
	// Either a valid SSN or EIN.
	TaxpayerIdentifier param.Field[string] `json:"taxpayer_identifier"`
}

func (r CounterpartyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Used to return counterparties created after some datetime.
	CreatedAtLowerBound param.Field[time.Time] `query:"created_at_lower_bound" format:"date-time"`
	// Used to return counterparties created before some datetime.
	CreatedAtUpperBound param.Field[time.Time] `query:"created_at_upper_bound" format:"date-time"`
	// Performs a partial string match of the email field. This is also case
	// insensitive.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters for counterparties with the given legal entity ID.
	LegalEntityID param.Field[string] `query:"legal_entity_id"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Performs a partial string match of the name field. This is also case
	// insensitive.
	Name    param.Field[string] `query:"name"`
	PerPage param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [CounterpartyListParams]'s query parameters as `url.Values`.
func (r CounterpartyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CounterpartyCollectAccountParams struct {
	// One of `credit` or `debit`. Use `credit` when you want to pay a counterparty.
	// Use `debit` when you need to charge a counterparty. This field helps us send a
	// more tailored email to your counterparties."
	Direction param.Field[shared.TransactionDirection] `json:"direction,required"`
	// The URL you want your customer to visit upon filling out the form. By default,
	// they will be sent to a Modern Treasury landing page. This must be a valid HTTPS
	// URL if set.
	CustomRedirect param.Field[string] `json:"custom_redirect" format:"uri"`
	// The list of fields you want on the form. This field is optional and if it is not
	// set, will default to [\"nameOnAccount\", \"accountType\", \"accountNumber\",
	// \"routingNumber\", \"address\"]. The full list of options is [\"name\",
	// \"nameOnAccount\", \"taxpayerIdentifier\", \"accountType\", \"accountNumber\",
	// \"routingNumber\", \"address\", \"ibanNumber\", \"swiftCode\"].
	Fields param.Field[[]CounterpartyCollectAccountParamsField] `json:"fields"`
	// By default, Modern Treasury will send an email to your counterparty that
	// includes a link to the form they must fill out. However, if you would like to
	// send the counterparty the link, you can set this parameter to `false`. The JSON
	// body will include the link to the secure Modern Treasury form.
	SendEmail param.Field[bool] `json:"send_email"`
}

func (r CounterpartyCollectAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyCollectAccountParamsField string

const (
	CounterpartyCollectAccountParamsFieldName                    CounterpartyCollectAccountParamsField = "name"
	CounterpartyCollectAccountParamsFieldNameOnAccount           CounterpartyCollectAccountParamsField = "nameOnAccount"
	CounterpartyCollectAccountParamsFieldTaxpayerIdentifier      CounterpartyCollectAccountParamsField = "taxpayerIdentifier"
	CounterpartyCollectAccountParamsFieldAccountType             CounterpartyCollectAccountParamsField = "accountType"
	CounterpartyCollectAccountParamsFieldAccountNumber           CounterpartyCollectAccountParamsField = "accountNumber"
	CounterpartyCollectAccountParamsFieldIbanNumber              CounterpartyCollectAccountParamsField = "ibanNumber"
	CounterpartyCollectAccountParamsFieldClabeNumber             CounterpartyCollectAccountParamsField = "clabeNumber"
	CounterpartyCollectAccountParamsFieldWalletAddress           CounterpartyCollectAccountParamsField = "walletAddress"
	CounterpartyCollectAccountParamsFieldPanNumber               CounterpartyCollectAccountParamsField = "panNumber"
	CounterpartyCollectAccountParamsFieldRoutingNumber           CounterpartyCollectAccountParamsField = "routingNumber"
	CounterpartyCollectAccountParamsFieldAbaWireRoutingNumber    CounterpartyCollectAccountParamsField = "abaWireRoutingNumber"
	CounterpartyCollectAccountParamsFieldSwiftCode               CounterpartyCollectAccountParamsField = "swiftCode"
	CounterpartyCollectAccountParamsFieldAuBsb                   CounterpartyCollectAccountParamsField = "auBsb"
	CounterpartyCollectAccountParamsFieldCaCpa                   CounterpartyCollectAccountParamsField = "caCpa"
	CounterpartyCollectAccountParamsFieldCnaps                   CounterpartyCollectAccountParamsField = "cnaps"
	CounterpartyCollectAccountParamsFieldGBSortCode              CounterpartyCollectAccountParamsField = "gbSortCode"
	CounterpartyCollectAccountParamsFieldInIfsc                  CounterpartyCollectAccountParamsField = "inIfsc"
	CounterpartyCollectAccountParamsFieldMyBranchCode            CounterpartyCollectAccountParamsField = "myBranchCode"
	CounterpartyCollectAccountParamsFieldBrCodigo                CounterpartyCollectAccountParamsField = "brCodigo"
	CounterpartyCollectAccountParamsFieldRoutingNumberType       CounterpartyCollectAccountParamsField = "routingNumberType"
	CounterpartyCollectAccountParamsFieldAddress                 CounterpartyCollectAccountParamsField = "address"
	CounterpartyCollectAccountParamsFieldJpZenginCode            CounterpartyCollectAccountParamsField = "jp_zengin_code"
	CounterpartyCollectAccountParamsFieldSeBankgiroClearingCode  CounterpartyCollectAccountParamsField = "se_bankgiro_clearing_code"
	CounterpartyCollectAccountParamsFieldNzNationalClearingCode  CounterpartyCollectAccountParamsField = "nz_national_clearing_code"
	CounterpartyCollectAccountParamsFieldHkInterbankClearingCode CounterpartyCollectAccountParamsField = "hk_interbank_clearing_code"
	CounterpartyCollectAccountParamsFieldHuInterbankClearingCode CounterpartyCollectAccountParamsField = "hu_interbank_clearing_code"
	CounterpartyCollectAccountParamsFieldDkInterbankClearingCode CounterpartyCollectAccountParamsField = "dk_interbank_clearing_code"
	CounterpartyCollectAccountParamsFieldIDSknbiCode             CounterpartyCollectAccountParamsField = "id_sknbi_code"
)
