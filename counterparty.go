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
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// A human friendly name for this counterparty.
	Name string `json:"name,required,nullable"`
	// The accounts for this counterparty.
	Accounts []CounterpartyAccounts `json:"accounts,required"`
	// The counterparty's email.
	Email string `json:"email,required,nullable" format:"email"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// Send an email to the counterparty whenever an associated payment order is sent
	// to the bank.
	SendRemittanceAdvice bool `json:"send_remittance_advice,required"`
	JSON                 counterpartyJSON
}

// counterpartyJSON contains the JSON metadata for the struct [Counterparty]
type counterpartyJSON struct {
	ID                   apijson.Field
	Object               apijson.Field
	LiveMode             apijson.Field
	CreatedAt            apijson.Field
	UpdatedAt            apijson.Field
	DiscardedAt          apijson.Field
	Name                 apijson.Field
	Accounts             apijson.Field
	Email                apijson.Field
	Metadata             apijson.Field
	SendRemittanceAdvice apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Counterparty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccounts struct {
	ID     string `json:"id" format:"uuid"`
	Object string `json:"object"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,nullable" format:"date-time"`
	// Can be `checking`, `savings` or `other`.
	AccountType ExternalAccountType `json:"account_type"`
	// Either `individual` or `business`.
	PartyType CounterpartyAccountsPartyType `json:"party_type,nullable"`
	// The address associated with the owner or `null`.
	PartyAddress CounterpartyAccountsPartyAddress `json:"party_address,nullable"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           string          `json:"name,nullable"`
	AccountDetails []AccountDetail `json:"account_details"`
	RoutingDetails []RoutingDetail `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata"`
	// The legal name of the entity which owns the account.
	PartyName      string                               `json:"party_name"`
	ContactDetails []CounterpartyAccountsContactDetails `json:"contact_details"`
	// If the external account links to a ledger account in Modern Treasury, the id of
	// the ledger account will be populated here.
	LedgerAccountID    string                                 `json:"ledger_account_id,nullable" format:"uuid"`
	VerificationStatus CounterpartyAccountsVerificationStatus `json:"verification_status"`
	JSON               counterpartyAccountsJSON
}

// counterpartyAccountsJSON contains the JSON metadata for the struct
// [CounterpartyAccounts]
type counterpartyAccountsJSON struct {
	ID                 apijson.Field
	Object             apijson.Field
	LiveMode           apijson.Field
	CreatedAt          apijson.Field
	UpdatedAt          apijson.Field
	DiscardedAt        apijson.Field
	AccountType        apijson.Field
	PartyType          apijson.Field
	PartyAddress       apijson.Field
	Name               apijson.Field
	AccountDetails     apijson.Field
	RoutingDetails     apijson.Field
	Metadata           apijson.Field
	PartyName          apijson.Field
	ContactDetails     apijson.Field
	LedgerAccountID    apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CounterpartyAccounts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsPartyType string

const (
	CounterpartyAccountsPartyTypeBusiness   CounterpartyAccountsPartyType = "business"
	CounterpartyAccountsPartyTypeIndividual CounterpartyAccountsPartyType = "individual"
)

// The address associated with the owner or `null`.
type CounterpartyAccountsPartyAddress struct {
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
	JSON    counterpartyAccountsPartyAddressJSON
}

// counterpartyAccountsPartyAddressJSON contains the JSON metadata for the struct
// [CounterpartyAccountsPartyAddress]
type counterpartyAccountsPartyAddressJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	LiveMode    apijson.Field
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	Region      apijson.Field
	PostalCode  apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CounterpartyAccountsPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsContactDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              bool                                                    `json:"live_mode,required"`
	CreatedAt             time.Time                                               `json:"created_at,required" format:"date-time"`
	UpdatedAt             time.Time                                               `json:"updated_at,required" format:"date-time"`
	DiscardedAt           time.Time                                               `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     string                                                  `json:"contact_identifier,required"`
	ContactIdentifierType CounterpartyAccountsContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	JSON                  counterpartyAccountsContactDetailsJSON
}

// counterpartyAccountsContactDetailsJSON contains the JSON metadata for the struct
// [CounterpartyAccountsContactDetails]
type counterpartyAccountsContactDetailsJSON struct {
	ID                    apijson.Field
	Object                apijson.Field
	LiveMode              apijson.Field
	CreatedAt             apijson.Field
	UpdatedAt             apijson.Field
	DiscardedAt           apijson.Field
	ContactIdentifier     apijson.Field
	ContactIdentifierType apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CounterpartyAccountsContactDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsContactDetailsContactIdentifierType string

const (
	CounterpartyAccountsContactDetailsContactIdentifierTypeEmail       CounterpartyAccountsContactDetailsContactIdentifierType = "email"
	CounterpartyAccountsContactDetailsContactIdentifierTypePhoneNumber CounterpartyAccountsContactDetailsContactIdentifierType = "phone_number"
	CounterpartyAccountsContactDetailsContactIdentifierTypeWebsite     CounterpartyAccountsContactDetailsContactIdentifierType = "website"
)

type CounterpartyAccountsVerificationStatus string

const (
	CounterpartyAccountsVerificationStatusPendingVerification CounterpartyAccountsVerificationStatus = "pending_verification"
	CounterpartyAccountsVerificationStatusUnverified          CounterpartyAccountsVerificationStatus = "unverified"
	CounterpartyAccountsVerificationStatusVerified            CounterpartyAccountsVerificationStatus = "verified"
)

type CounterpartyCollectAccountResponse struct {
	// The id of the existing counterparty.
	ID string `json:"id,required"`
	// This field will be `true` if an email requesting account details has already
	// been sent to this counterparty.
	IsResend bool `json:"is_resend,required"`
	// This is the link to the secure Modern Treasury form. By default, Modern Treasury
	// will send an email to your counterparty that includes a link to this form.
	// However, if `send_email` is passed as `false` in the body then Modern Treasury
	// will not send the email and you can send it to the counterparty directly.
	FormLink string `json:"form_link,required" format:"uri"`
	JSON     counterpartyCollectAccountResponseJSON
}

// counterpartyCollectAccountResponseJSON contains the JSON metadata for the struct
// [CounterpartyCollectAccountResponse]
type counterpartyCollectAccountResponseJSON struct {
	ID          apijson.Field
	IsResend    apijson.Field
	FormLink    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CounterpartyCollectAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CounterpartyNewParams struct {
	// A human friendly name for this counterparty.
	Name param.Field[string] `json:"name,required,nullable"`
	// The accounts for this counterparty.
	Accounts param.Field[[]CounterpartyNewParamsAccounts] `json:"accounts"`
	// The counterparty's email.
	Email param.Field[string] `json:"email,nullable" format:"email"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Send an email to the counterparty whenever an associated payment order is sent
	// to the bank.
	SendRemittanceAdvice param.Field[bool]                            `json:"send_remittance_advice"`
	Accounting           param.Field[CounterpartyNewParamsAccounting] `json:"accounting"`
	// An optional type to auto-sync the counterparty to your ledger. Either `customer`
	// or `vendor`.
	LedgerType param.Field[CounterpartyNewParamsLedgerType] `json:"ledger_type"`
	// Either a valid SSN or EIN.
	TaxpayerIdentifier param.Field[string] `json:"taxpayer_identifier"`
}

func (r CounterpartyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyNewParamsAccounts struct {
	// Can be `checking`, `savings` or `other`.
	AccountType param.Field[ExternalAccountType] `json:"account_type"`
	// Either `individual` or `business`.
	PartyType param.Field[CounterpartyNewParamsAccountsPartyType] `json:"party_type,nullable"`
	// Required if receiving wire payments.
	PartyAddress param.Field[CounterpartyNewParamsAccountsPartyAddress] `json:"party_address"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           param.Field[string]                                        `json:"name,nullable"`
	AccountDetails param.Field[[]CounterpartyNewParamsAccountsAccountDetails] `json:"account_details"`
	RoutingDetails param.Field[[]CounterpartyNewParamsAccountsRoutingDetails] `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// If this value isn't provided, it will be inherited from the counterparty's name.
	PartyName       param.Field[string] `json:"party_name"`
	PartyIdentifier param.Field[string] `json:"party_identifier"`
	// Specifies a ledger account object that will be created with the external
	// account. The resulting ledger account is linked to the external account for
	// auto-ledgering Payment objects. See
	// https://dash.readme.com/project/modern-treasury/v1.1/docs/linking-to-other-modern-treasury-objects
	// for more details.
	LedgerAccount param.Field[CounterpartyNewParamsAccountsLedgerAccount] `json:"ledger_account"`
	// If you've enabled the Modern Treasury + Plaid integration in your Plaid account,
	// you can pass the processor token in this field.
	PlaidProcessorToken param.Field[string]                                        `json:"plaid_processor_token"`
	ContactDetails      param.Field[[]CounterpartyNewParamsAccountsContactDetails] `json:"contact_details"`
}

type CounterpartyNewParamsAccountsPartyType string

const (
	CounterpartyNewParamsAccountsPartyTypeBusiness   CounterpartyNewParamsAccountsPartyType = "business"
	CounterpartyNewParamsAccountsPartyTypeIndividual CounterpartyNewParamsAccountsPartyType = "individual"
)

// Required if receiving wire payments.
type CounterpartyNewParamsAccountsPartyAddress struct {
	Line1 param.Field[string] `json:"line1,nullable"`
	Line2 param.Field[string] `json:"line2,nullable"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,nullable"`
	// Region or State.
	Region param.Field[string] `json:"region,nullable"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,nullable"`
}

type CounterpartyNewParamsAccountsAccountDetails struct {
	AccountNumber     param.Field[string]                                                       `json:"account_number,required"`
	AccountNumberType param.Field[CounterpartyNewParamsAccountsAccountDetailsAccountNumberType] `json:"account_number_type"`
}

type CounterpartyNewParamsAccountsAccountDetailsAccountNumberType string

const (
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban          CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "iban"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeClabe         CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "clabe"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeWalletAddress CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "wallet_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePan           CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "pan"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeOther         CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "other"
)

type CounterpartyNewParamsAccountsRoutingDetails struct {
	RoutingNumber     param.Field[string]                                                       `json:"routing_number,required"`
	RoutingNumberType param.Field[CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType] `json:"routing_number_type,required"`
	PaymentType       param.Field[CounterpartyNewParamsAccountsRoutingDetailsPaymentType]       `json:"payment_type"`
}

type CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba          CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "aba"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAuBsb        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "au_bsb"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeBrCodigo     CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "br_codigo"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCaCpa        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "ca_cpa"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeChips        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "chips"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCnaps        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "cnaps"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeGBSortCode   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "gb_sort_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeInIfsc       CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "in_ifsc"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMyBranchCode CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "my_branch_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSwift        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "swift"
)

type CounterpartyNewParamsAccountsRoutingDetailsPaymentType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ach"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeAuBecs      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "au_becs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBacs        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "bacs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBook        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "book"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCard        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "card"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCheck       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "check"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEft         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "eft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCrossBorder CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "cross_border"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeInterac     CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "interac"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMasav       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "masav"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNeft        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "neft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeProvxchange CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "provxchange"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRtp         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "rtp"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSen         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sen"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSepa        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sepa"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSignet      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "signet"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeWire        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "wire"
)

// Specifies a ledger account object that will be created with the external
// account. The resulting ledger account is linked to the external account for
// auto-ledgering Payment objects. See
// https://dash.readme.com/project/modern-treasury/v1.1/docs/linking-to-other-modern-treasury-objects
// for more details.
type CounterpartyNewParamsAccountsLedgerAccount struct {
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description,nullable"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[CounterpartyNewParamsAccountsLedgerAccountNormalBalance] `json:"normal_balance,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent,nullable"`
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

type CounterpartyNewParamsAccountsLedgerAccountNormalBalance string

const (
	CounterpartyNewParamsAccountsLedgerAccountNormalBalanceCredit CounterpartyNewParamsAccountsLedgerAccountNormalBalance = "credit"
	CounterpartyNewParamsAccountsLedgerAccountNormalBalanceDebit  CounterpartyNewParamsAccountsLedgerAccountNormalBalance = "debit"
)

type CounterpartyNewParamsAccountsLedgerAccountLedgerableType string

const (
	CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeExternalAccount CounterpartyNewParamsAccountsLedgerAccountLedgerableType = "external_account"
	CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeInternalAccount CounterpartyNewParamsAccountsLedgerAccountLedgerableType = "internal_account"
)

type CounterpartyNewParamsAccountsContactDetails struct {
	ContactIdentifier     param.Field[string]                                                           `json:"contact_identifier"`
	ContactIdentifierType param.Field[CounterpartyNewParamsAccountsContactDetailsContactIdentifierType] `json:"contact_identifier_type"`
}

type CounterpartyNewParamsAccountsContactDetailsContactIdentifierType string

const (
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail       CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "email"
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypePhoneNumber CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "phone_number"
	CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeWebsite     CounterpartyNewParamsAccountsContactDetailsContactIdentifierType = "website"
)

type CounterpartyNewParamsAccounting struct {
	// An optional type to auto-sync the counterparty to your ledger. Either `customer`
	// or `vendor`.
	Type param.Field[CounterpartyNewParamsAccountingType] `json:"type"`
}

type CounterpartyNewParamsAccountingType string

const (
	CounterpartyNewParamsAccountingTypeCustomer CounterpartyNewParamsAccountingType = "customer"
	CounterpartyNewParamsAccountingTypeVendor   CounterpartyNewParamsAccountingType = "vendor"
)

type CounterpartyNewParamsLedgerType string

const (
	CounterpartyNewParamsLedgerTypeCustomer CounterpartyNewParamsLedgerType = "customer"
	CounterpartyNewParamsLedgerTypeVendor   CounterpartyNewParamsLedgerType = "vendor"
)

type CounterpartyUpdateParams struct {
	// A new name for the counterparty. Will only update if passed.
	Name param.Field[string] `json:"name"`
	// A new email for the counterparty.
	Email param.Field[string] `json:"email" format:"email"`
	// Additional data in the form of key-value pairs. Pairs can be removed by passing
	// an empty string or `null` as the value.
	Metadata param.Field[map[string]string] `json:"metadata"`
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
	AfterCursor param.Field[string] `query:"after_cursor,nullable"`
	PerPage     param.Field[int64]  `query:"per_page"`
	// Performs a partial string match of the name field. This is also case
	// insensitive.
	Name param.Field[string] `query:"name"`
	// Performs a partial string match of the email field. This is also case
	// insensitive.
	Email param.Field[string] `query:"email" format:"email"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata param.Field[map[string]string] `query:"metadata"`
	// Used to return counterparties created after some datetime.
	CreatedAtLowerBound param.Field[time.Time] `query:"created_at_lower_bound" format:"date-time"`
	// Used to return counterparties created before some datetime.
	CreatedAtUpperBound param.Field[time.Time] `query:"created_at_upper_bound" format:"date-time"`
}

// URLQuery serializes [CounterpartyListParams]'s query parameters as `url.Values`.
func (r CounterpartyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CounterpartyCollectAccountParams struct {
	// One of `credit` or `debit`. Use `credit` when you want to pay a counterparty.
	// Use `debit` when you need to charge a counterparty. This field helps us send a
	// more tailored email to your counterparties."
	Direction param.Field[CounterpartyCollectAccountParamsDirection] `json:"direction,required"`
	// By default, Modern Treasury will send an email to your counterparty that
	// includes a link to the form they must fill out. However, if you would like to
	// send the counterparty the link, you can set this parameter to `false`. The JSON
	// body will include the link to the secure Modern Treasury form.
	SendEmail param.Field[bool] `json:"send_email"`
	// The list of fields you want on the form. This field is optional and if it is not
	// set, will default to [\"nameOnAccount\", \"accountType\", \"accountNumber\",
	// \"routingNumber\", \"address\"]. The full list of options is [\"name\",
	// \"nameOnAccount\", \"taxpayerIdentifier\", \"accountType\", \"accountNumber\",
	// \"routingNumber\", \"address\", \"ibanNumber\", \"swiftCode\"].
	Fields param.Field[[]CounterpartyCollectAccountParamsFields] `json:"fields"`
	// The URL you want your customer to visit upon filling out the form. By default,
	// they will be sent to a Modern Treasury landing page. This must be a valid HTTPS
	// URL if set.
	CustomRedirect param.Field[string] `json:"custom_redirect" format:"uri"`
}

func (r CounterpartyCollectAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterpartyCollectAccountParamsDirection string

const (
	CounterpartyCollectAccountParamsDirectionCredit CounterpartyCollectAccountParamsDirection = "credit"
	CounterpartyCollectAccountParamsDirectionDebit  CounterpartyCollectAccountParamsDirection = "debit"
)

type CounterpartyCollectAccountParamsFields string

const (
	CounterpartyCollectAccountParamsFieldsName                 CounterpartyCollectAccountParamsFields = "name"
	CounterpartyCollectAccountParamsFieldsNameOnAccount        CounterpartyCollectAccountParamsFields = "nameOnAccount"
	CounterpartyCollectAccountParamsFieldsTaxpayerIdentifier   CounterpartyCollectAccountParamsFields = "taxpayerIdentifier"
	CounterpartyCollectAccountParamsFieldsAccountType          CounterpartyCollectAccountParamsFields = "accountType"
	CounterpartyCollectAccountParamsFieldsAccountNumber        CounterpartyCollectAccountParamsFields = "accountNumber"
	CounterpartyCollectAccountParamsFieldsIbanNumber           CounterpartyCollectAccountParamsFields = "ibanNumber"
	CounterpartyCollectAccountParamsFieldsClabeNumber          CounterpartyCollectAccountParamsFields = "clabeNumber"
	CounterpartyCollectAccountParamsFieldsWalletAddress        CounterpartyCollectAccountParamsFields = "walletAddress"
	CounterpartyCollectAccountParamsFieldsPanNumber            CounterpartyCollectAccountParamsFields = "panNumber"
	CounterpartyCollectAccountParamsFieldsRoutingNumber        CounterpartyCollectAccountParamsFields = "routingNumber"
	CounterpartyCollectAccountParamsFieldsAbaWireRoutingNumber CounterpartyCollectAccountParamsFields = "abaWireRoutingNumber"
	CounterpartyCollectAccountParamsFieldsSwiftCode            CounterpartyCollectAccountParamsFields = "swiftCode"
	CounterpartyCollectAccountParamsFieldsAuBsb                CounterpartyCollectAccountParamsFields = "auBsb"
	CounterpartyCollectAccountParamsFieldsCaCpa                CounterpartyCollectAccountParamsFields = "caCpa"
	CounterpartyCollectAccountParamsFieldsCnaps                CounterpartyCollectAccountParamsFields = "cnaps"
	CounterpartyCollectAccountParamsFieldsGBSortCode           CounterpartyCollectAccountParamsFields = "gbSortCode"
	CounterpartyCollectAccountParamsFieldsInIfsc               CounterpartyCollectAccountParamsFields = "inIfsc"
	CounterpartyCollectAccountParamsFieldsMyBranchCode         CounterpartyCollectAccountParamsFields = "myBranchCode"
	CounterpartyCollectAccountParamsFieldsBrCodigo             CounterpartyCollectAccountParamsFields = "brCodigo"
	CounterpartyCollectAccountParamsFieldsRoutingNumberType    CounterpartyCollectAccountParamsFields = "routingNumberType"
	CounterpartyCollectAccountParamsFieldsAddress              CounterpartyCollectAccountParamsFields = "address"
)
