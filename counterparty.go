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

// CounterpartyService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCounterpartyService] method instead.
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a given counterparty with new information.
func (r *CounterpartyService) Update(ctx context.Context, id string, body CounterpartyUpdateParams, opts ...option.RequestOption) (res *Counterparty, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of all counterparties.
func (r *CounterpartyService) List(ctx context.Context, query CounterpartyListParams, opts ...option.RequestOption) (res *pagination.Page[Counterparty], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *CounterpartyService) ListAutoPaging(ctx context.Context, query CounterpartyListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Counterparty] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a given counterparty.
func (r *CounterpartyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/counterparties/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Send an email requesting account details.
func (r *CounterpartyService) CollectAccount(ctx context.Context, id string, body CounterpartyCollectAccountParams, opts ...option.RequestOption) (res *CounterpartyCollectAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
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

func (r counterpartyJSON) RawJSON() string {
	return r.raw
}

type CounterpartyAccount struct {
	ID             string          `json:"id" format:"uuid"`
	AccountDetails []AccountDetail `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    ExternalAccountType    `json:"account_type"`
	ContactDetails []shared.ContactDetail `json:"contact_details"`
	CreatedAt      time.Time              `json:"created_at" format:"date-time"`
	DiscardedAt    time.Time              `json:"discarded_at,nullable" format:"date-time"`
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
	PartyAddress shared.Address `json:"party_address,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name"`
	// Either `individual` or `business`.
	PartyType          CounterpartyAccountsPartyType          `json:"party_type,nullable"`
	RoutingDetails     []RoutingDetail                        `json:"routing_details"`
	UpdatedAt          time.Time                              `json:"updated_at" format:"date-time"`
	VerificationSource CounterpartyAccountsVerificationSource `json:"verification_source,nullable"`
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
	VerificationSource apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CounterpartyAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r counterpartyAccountJSON) RawJSON() string {
	return r.raw
}

// Either `individual` or `business`.
type CounterpartyAccountsPartyType string

const (
	CounterpartyAccountsPartyTypeBusiness   CounterpartyAccountsPartyType = "business"
	CounterpartyAccountsPartyTypeIndividual CounterpartyAccountsPartyType = "individual"
)

func (r CounterpartyAccountsPartyType) IsKnown() bool {
	switch r {
	case CounterpartyAccountsPartyTypeBusiness, CounterpartyAccountsPartyTypeIndividual:
		return true
	}
	return false
}

type CounterpartyAccountsVerificationSource string

const (
	CounterpartyAccountsVerificationSourceACHPrenote    CounterpartyAccountsVerificationSource = "ach_prenote"
	CounterpartyAccountsVerificationSourceMicrodeposits CounterpartyAccountsVerificationSource = "microdeposits"
	CounterpartyAccountsVerificationSourcePlaid         CounterpartyAccountsVerificationSource = "plaid"
)

func (r CounterpartyAccountsVerificationSource) IsKnown() bool {
	switch r {
	case CounterpartyAccountsVerificationSourceACHPrenote, CounterpartyAccountsVerificationSourceMicrodeposits, CounterpartyAccountsVerificationSourcePlaid:
		return true
	}
	return false
}

type CounterpartyAccountsVerificationStatus string

const (
	CounterpartyAccountsVerificationStatusPendingVerification CounterpartyAccountsVerificationStatus = "pending_verification"
	CounterpartyAccountsVerificationStatusUnverified          CounterpartyAccountsVerificationStatus = "unverified"
	CounterpartyAccountsVerificationStatusVerified            CounterpartyAccountsVerificationStatus = "verified"
)

func (r CounterpartyAccountsVerificationStatus) IsKnown() bool {
	switch r {
	case CounterpartyAccountsVerificationStatusPendingVerification, CounterpartyAccountsVerificationStatusUnverified, CounterpartyAccountsVerificationStatusVerified:
		return true
	}
	return false
}

// The verification status of the counterparty.
type CounterpartyVerificationStatus string

const (
	CounterpartyVerificationStatusDenied        CounterpartyVerificationStatus = "denied"
	CounterpartyVerificationStatusNeedsApproval CounterpartyVerificationStatus = "needs_approval"
	CounterpartyVerificationStatusUnverified    CounterpartyVerificationStatus = "unverified"
	CounterpartyVerificationStatusVerified      CounterpartyVerificationStatus = "verified"
)

func (r CounterpartyVerificationStatus) IsKnown() bool {
	switch r {
	case CounterpartyVerificationStatusDenied, CounterpartyVerificationStatusNeedsApproval, CounterpartyVerificationStatusUnverified, CounterpartyVerificationStatusVerified:
		return true
	}
	return false
}

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

func (r counterpartyCollectAccountResponseJSON) RawJSON() string {
	return r.raw
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

func (r CounterpartyNewParamsAccountingType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsAccountingTypeCustomer, CounterpartyNewParamsAccountingTypeVendor:
		return true
	}
	return false
}

type CounterpartyNewParamsAccount struct {
	AccountDetails param.Field[[]CounterpartyNewParamsAccountsAccountDetail] `json:"account_details"`
	// Can be `checking`, `savings` or `other`.
	AccountType    param.Field[ExternalAccountType]               `json:"account_type"`
	ContactDetails param.Field[[]ContactDetailCreateRequestParam] `json:"contact_details"`
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
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber        CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "au_number"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeBaseAddress     CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "base_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeClabe           CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "clabe"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeEthereumAddress CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "ethereum_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeHkNumber        CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "hk_number"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban            CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "iban"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIDNumber        CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "id_number"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeNzNumber        CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "nz_number"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeOther           CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "other"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePan             CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "pan"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePolygonAddress  CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "polygon_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeSgNumber        CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "sg_number"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeSolanaAddress   CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "solana_address"
	CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeWalletAddress   CounterpartyNewParamsAccountsAccountDetailsAccountNumberType = "wallet_address"
)

func (r CounterpartyNewParamsAccountsAccountDetailsAccountNumberType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeBaseAddress, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeClabe, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeEthereumAddress, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeHkNumber, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIDNumber, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeNzNumber, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeOther, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePan, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypePolygonAddress, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeSgNumber, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeSolanaAddress, CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeWalletAddress:
		return true
	}
	return false
}

// Either `individual` or `business`.
type CounterpartyNewParamsAccountsPartyType string

const (
	CounterpartyNewParamsAccountsPartyTypeBusiness   CounterpartyNewParamsAccountsPartyType = "business"
	CounterpartyNewParamsAccountsPartyTypeIndividual CounterpartyNewParamsAccountsPartyType = "individual"
)

func (r CounterpartyNewParamsAccountsPartyType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsAccountsPartyTypeBusiness, CounterpartyNewParamsAccountsPartyTypeIndividual:
		return true
	}
	return false
}

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
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeIlBankCode              CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "il_bank_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeInIfsc                  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "in_ifsc"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeJpZenginCode            CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "jp_zengin_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMyBranchCode            CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "my_branch_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMxBankIdentifier        CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "mx_bank_identifier"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeNzNationalClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "nz_national_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypePlNationalClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "pl_national_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "se_bankgiro_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSgInterbankClearingCode CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "sg_interbank_clearing_code"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSwift                   CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "swift"
	CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeZaNationalClearingCode  CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType = "za_national_clearing_code"
)

func (r CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAuBsb, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeBrCodigo, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCaCpa, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeChips, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeCnaps, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeDkInterbankClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeGBSortCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeHkInterbankClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeHuInterbankClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeIDSknbiCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeIlBankCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeInIfsc, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeJpZenginCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMyBranchCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeMxBankIdentifier, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeNzNationalClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypePlNationalClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSeBankgiroClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSgInterbankClearingCode, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeSwift, CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeZaNationalClearingCode:
		return true
	}
	return false
}

type CounterpartyNewParamsAccountsRoutingDetailsPaymentType string

const (
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ach"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeAuBecs      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "au_becs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBacs        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "bacs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBase        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "base"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBook        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "book"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCard        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "card"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeChats       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "chats"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCheck       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "check"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCrossBorder CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "cross_border"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeDkNets      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "dk_nets"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEft         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "eft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEthereum    CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ethereum"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeHuIcs       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "hu_ics"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeInterac     CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "interac"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMasav       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "masav"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMxCcen      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "mx_ccen"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNeft        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "neft"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNics        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "nics"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNzBecs      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "nz_becs"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypePlElixir    CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "pl_elixir"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypePolygon     CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "polygon"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeProvxchange CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "provxchange"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRoSent      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "ro_sent"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRtp         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "rtp"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSeBankgirot CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "se_bankgirot"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSen         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sen"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSepa        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sepa"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSgGiro      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sg_giro"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSic         CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sic"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSignet      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "signet"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSknbi       CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "sknbi"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSolana      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "solana"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeWire        CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "wire"
	CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeZengin      CounterpartyNewParamsAccountsRoutingDetailsPaymentType = "zengin"
)

func (r CounterpartyNewParamsAccountsRoutingDetailsPaymentType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeAuBecs, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBacs, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBase, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeBook, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCard, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeChats, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCheck, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeCrossBorder, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeDkNets, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEft, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeEthereum, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeHuIcs, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeInterac, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMasav, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeMxCcen, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNeft, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNics, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeNzBecs, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypePlElixir, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypePolygon, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeProvxchange, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRoSent, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeRtp, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSeBankgirot, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSen, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSepa, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSgGiro, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSic, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSignet, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSknbi, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeSolana, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeWire, CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeZengin:
		return true
	}
	return false
}

// An optional type to auto-sync the counterparty to your ledger. Either `customer`
// or `vendor`.
type CounterpartyNewParamsLedgerType string

const (
	CounterpartyNewParamsLedgerTypeCustomer CounterpartyNewParamsLedgerType = "customer"
	CounterpartyNewParamsLedgerTypeVendor   CounterpartyNewParamsLedgerType = "vendor"
)

func (r CounterpartyNewParamsLedgerType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsLedgerTypeCustomer, CounterpartyNewParamsLedgerTypeVendor:
		return true
	}
	return false
}

type CounterpartyNewParamsLegalEntity struct {
	// The type of legal entity.
	LegalEntityType param.Field[CounterpartyNewParamsLegalEntityLegalEntityType] `json:"legal_entity_type,required"`
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[BankSettingsParam]                             `json:"bank_settings"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// The country of citizenship for an individual.
	CitizenshipCountry param.Field[string]                                  `json:"citizenship_country"`
	ComplianceDetails  param.Field[shared.LegalEntityComplianceDetailParam] `json:"compliance_details"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed param.Field[time.Time] `json:"date_formed" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// A list of identifications for the legal entity.
	Identifications param.Field[[]shared.IdentificationCreateRequestParam] `json:"identifications"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications param.Field[[]shared.LegalEntityIndustryClassificationParam] `json:"industry_classifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations param.Field[[]CounterpartyNewParamsLegalEntityLegalEntityAssociation] `json:"legal_entity_associations"`
	// The business's legal structure.
	LegalStructure param.Field[CounterpartyNewParamsLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName   param.Field[string]                                        `json:"middle_name"`
	PhoneNumbers param.Field[[]CounterpartyNewParamsLegalEntityPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[CounterpartyNewParamsLegalEntityRiskRating] `json:"risk_rating"`
	// An individual's suffix.
	Suffix                     param.Field[string]                          `json:"suffix"`
	WealthAndEmploymentDetails param.Field[WealthAndEmploymentDetailsParam] `json:"wealth_and_employment_details"`
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

func (r CounterpartyNewParamsLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsLegalEntityLegalEntityTypeBusiness, CounterpartyNewParamsLegalEntityLegalEntityTypeIndividual:
		return true
	}
	return false
}

type CounterpartyNewParamsLegalEntityLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[shared.ChildLegalEntityCreateParam] `json:"child_legal_entity"`
	// The ID of the child legal entity.
	ChildLegalEntityID param.Field[string] `json:"child_legal_entity_id"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the child entity at the parent entity.
	Title param.Field[string] `json:"title"`
}

func (r CounterpartyNewParamsLegalEntityLegalEntityAssociation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the child entity relates to parent entity.
type CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType string

const (
	CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "beneficial_owner"
	CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson   CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "control_person"
)

func (r CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson:
		return true
	}
	return false
}

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

func (r CounterpartyNewParamsLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsLegalEntityLegalStructureCorporation, CounterpartyNewParamsLegalEntityLegalStructureLlc, CounterpartyNewParamsLegalEntityLegalStructureNonProfit, CounterpartyNewParamsLegalEntityLegalStructurePartnership, CounterpartyNewParamsLegalEntityLegalStructureSoleProprietorship, CounterpartyNewParamsLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type CounterpartyNewParamsLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r CounterpartyNewParamsLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The risk rating of the legal entity. One of low, medium, high.
type CounterpartyNewParamsLegalEntityRiskRating string

const (
	CounterpartyNewParamsLegalEntityRiskRatingLow    CounterpartyNewParamsLegalEntityRiskRating = "low"
	CounterpartyNewParamsLegalEntityRiskRatingMedium CounterpartyNewParamsLegalEntityRiskRating = "medium"
	CounterpartyNewParamsLegalEntityRiskRatingHigh   CounterpartyNewParamsLegalEntityRiskRating = "high"
)

func (r CounterpartyNewParamsLegalEntityRiskRating) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsLegalEntityRiskRatingLow, CounterpartyNewParamsLegalEntityRiskRatingMedium, CounterpartyNewParamsLegalEntityRiskRatingHigh:
		return true
	}
	return false
}

// The verification status of the counterparty.
type CounterpartyNewParamsVerificationStatus string

const (
	CounterpartyNewParamsVerificationStatusDenied        CounterpartyNewParamsVerificationStatus = "denied"
	CounterpartyNewParamsVerificationStatusNeedsApproval CounterpartyNewParamsVerificationStatus = "needs_approval"
	CounterpartyNewParamsVerificationStatusUnverified    CounterpartyNewParamsVerificationStatus = "unverified"
	CounterpartyNewParamsVerificationStatusVerified      CounterpartyNewParamsVerificationStatus = "verified"
)

func (r CounterpartyNewParamsVerificationStatus) IsKnown() bool {
	switch r {
	case CounterpartyNewParamsVerificationStatusDenied, CounterpartyNewParamsVerificationStatusNeedsApproval, CounterpartyNewParamsVerificationStatusUnverified, CounterpartyNewParamsVerificationStatusVerified:
		return true
	}
	return false
}

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
	CounterpartyCollectAccountParamsFieldJpZenginCode            CounterpartyCollectAccountParamsField = "jpZenginCode"
	CounterpartyCollectAccountParamsFieldSeBankgiroClearingCode  CounterpartyCollectAccountParamsField = "seBankgiroClearingCode"
	CounterpartyCollectAccountParamsFieldNzNationalClearingCode  CounterpartyCollectAccountParamsField = "nzNationalClearingCode"
	CounterpartyCollectAccountParamsFieldHkInterbankClearingCode CounterpartyCollectAccountParamsField = "hkInterbankClearingCode"
	CounterpartyCollectAccountParamsFieldHuInterbankClearingCode CounterpartyCollectAccountParamsField = "huInterbankClearingCode"
	CounterpartyCollectAccountParamsFieldDkInterbankClearingCode CounterpartyCollectAccountParamsField = "dkInterbankClearingCode"
	CounterpartyCollectAccountParamsFieldIDSknbiCode             CounterpartyCollectAccountParamsField = "idSknbiCode"
	CounterpartyCollectAccountParamsFieldZaNationalClearingCode  CounterpartyCollectAccountParamsField = "zaNationalClearingCode"
)

func (r CounterpartyCollectAccountParamsField) IsKnown() bool {
	switch r {
	case CounterpartyCollectAccountParamsFieldName, CounterpartyCollectAccountParamsFieldNameOnAccount, CounterpartyCollectAccountParamsFieldTaxpayerIdentifier, CounterpartyCollectAccountParamsFieldAccountType, CounterpartyCollectAccountParamsFieldAccountNumber, CounterpartyCollectAccountParamsFieldIbanNumber, CounterpartyCollectAccountParamsFieldClabeNumber, CounterpartyCollectAccountParamsFieldWalletAddress, CounterpartyCollectAccountParamsFieldPanNumber, CounterpartyCollectAccountParamsFieldRoutingNumber, CounterpartyCollectAccountParamsFieldAbaWireRoutingNumber, CounterpartyCollectAccountParamsFieldSwiftCode, CounterpartyCollectAccountParamsFieldAuBsb, CounterpartyCollectAccountParamsFieldCaCpa, CounterpartyCollectAccountParamsFieldCnaps, CounterpartyCollectAccountParamsFieldGBSortCode, CounterpartyCollectAccountParamsFieldInIfsc, CounterpartyCollectAccountParamsFieldMyBranchCode, CounterpartyCollectAccountParamsFieldBrCodigo, CounterpartyCollectAccountParamsFieldRoutingNumberType, CounterpartyCollectAccountParamsFieldAddress, CounterpartyCollectAccountParamsFieldJpZenginCode, CounterpartyCollectAccountParamsFieldSeBankgiroClearingCode, CounterpartyCollectAccountParamsFieldNzNationalClearingCode, CounterpartyCollectAccountParamsFieldHkInterbankClearingCode, CounterpartyCollectAccountParamsFieldHuInterbankClearingCode, CounterpartyCollectAccountParamsFieldDkInterbankClearingCode, CounterpartyCollectAccountParamsFieldIDSknbiCode, CounterpartyCollectAccountParamsFieldZaNationalClearingCode:
		return true
	}
	return false
}
