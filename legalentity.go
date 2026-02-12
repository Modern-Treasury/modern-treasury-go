// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// LegalEntityService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegalEntityService] method instead.
type LegalEntityService struct {
	Options []option.RequestOption
}

// NewLegalEntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLegalEntityService(opts ...option.RequestOption) (r *LegalEntityService) {
	r = &LegalEntityService{}
	r.Options = opts
	return
}

// create legal_entity
func (r *LegalEntityService) New(ctx context.Context, body LegalEntityNewParams, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/legal_entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single legal entity.
func (r *LegalEntityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a legal entity.
func (r *LegalEntityService) Update(ctx context.Context, id string, body LegalEntityUpdateParams, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all legal entities.
func (r *LegalEntityService) List(ctx context.Context, query LegalEntityListParams, opts ...option.RequestOption) (res *pagination.Page[LegalEntity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/legal_entities"
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

// Get a list of all legal entities.
func (r *LegalEntityService) ListAutoPaging(ctx context.Context, query LegalEntityListParams, opts ...option.RequestOption) *pagination.PageAutoPager[LegalEntity] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type LegalEntity struct {
	ID string `json:"id,required" format:"uuid"`
	// A list of addresses for the entity.
	Addresses    []LegalEntityAddress           `json:"addresses,required"`
	BankSettings shared.LegalEntityBankSettings `json:"bank_settings,required,nullable"`
	// A description of the business.
	BusinessDescription string `json:"business_description,required,nullable"`
	// The business's legal business name.
	BusinessName string `json:"business_name,required,nullable"`
	// The country of citizenship for an individual.
	CitizenshipCountry string `json:"citizenship_country,required,nullable"`
	// Deprecated: deprecated
	ComplianceDetails interface{} `json:"compliance_details,required,nullable"`
	// The country code where the business is incorporated in the ISO 3166-1 alpha-2 or
	// alpha-3 formats.
	CountryOfIncorporation string    `json:"country_of_incorporation,required,nullable"`
	CreatedAt              time.Time `json:"created_at,required" format:"date-time"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed time.Time `json:"date_formed,required,nullable" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          time.Time  `json:"date_of_birth,required,nullable" format:"date"`
	DiscardedAt          time.Time  `json:"discarded_at,required,nullable" format:"date-time"`
	Documents            []Document `json:"documents,required"`
	DoingBusinessAsNames []string   `json:"doing_business_as_names,required"`
	// The entity's primary email.
	Email string `json:"email,required,nullable"`
	// Monthly expected transaction volume in USD.
	ExpectedActivityVolume int64 `json:"expected_activity_volume,required,nullable"`
	// An individual's first name.
	FirstName string `json:"first_name,required,nullable"`
	// A list of identifications for the legal entity.
	Identifications []LegalEntityIdentification `json:"identifications,required"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications []shared.LegalEntityIndustryClassification `json:"industry_classifications,required"`
	// A description of the intended use of the legal entity.
	IntendedUse string `json:"intended_use,required,nullable"`
	// An individual's last name.
	LastName string `json:"last_name,required,nullable"`
	// The type of legal entity.
	LegalEntityType LegalEntityLegalEntityType `json:"legal_entity_type,required"`
	// The business's legal structure.
	LegalStructure LegalEntityLegalStructure `json:"legal_structure,required,nullable"`
	// ISO 10383 market identifier code.
	ListedExchange string `json:"listed_exchange,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// An individual's middle name.
	MiddleName string `json:"middle_name,required,nullable"`
	Object     string `json:"object,required"`
	// A list of countries where the business operates (ISO 3166-1 alpha-2 or alpha-3
	// codes).
	OperatingJurisdictions []string                 `json:"operating_jurisdictions,required"`
	PhoneNumbers           []LegalEntityPhoneNumber `json:"phone_numbers,required"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson bool `json:"politically_exposed_person,required,nullable"`
	// An individual's preferred name.
	PreferredName string `json:"preferred_name,required,nullable"`
	// An individual's prefix.
	Prefix string `json:"prefix,required,nullable"`
	// A list of primary social media URLs for the business.
	PrimarySocialMediaSites []string `json:"primary_social_media_sites,required"`
	// Array of regulatory bodies overseeing this institution.
	Regulators []LegalEntityRegulator `json:"regulators,required,nullable"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating LegalEntityRiskRating `json:"risk_rating,required,nullable"`
	// The activation status of the legal entity. One of pending, active, suspended, or
	// closed.
	Status LegalEntityStatus `json:"status,required,nullable"`
	// An individual's suffix.
	Suffix string `json:"suffix,required,nullable"`
	// Information describing a third-party verification run by an external vendor.
	ThirdPartyVerification LegalEntityThirdPartyVerification `json:"third_party_verification,required,nullable"`
	// Stock ticker symbol for publicly traded companies.
	TickerSymbol               string                                   `json:"ticker_symbol,required,nullable"`
	UpdatedAt                  time.Time                                `json:"updated_at,required" format:"date-time"`
	WealthAndEmploymentDetails shared.LegalEntityWealthEmploymentDetail `json:"wealth_and_employment_details,required,nullable"`
	// The entity's primary website URL.
	Website string `json:"website,required,nullable"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations []LegalEntityAssociation `json:"legal_entity_associations,nullable"`
	JSON                    legalEntityJSON          `json:"-"`
}

// legalEntityJSON contains the JSON metadata for the struct [LegalEntity]
type legalEntityJSON struct {
	ID                         apijson.Field
	Addresses                  apijson.Field
	BankSettings               apijson.Field
	BusinessDescription        apijson.Field
	BusinessName               apijson.Field
	CitizenshipCountry         apijson.Field
	ComplianceDetails          apijson.Field
	CountryOfIncorporation     apijson.Field
	CreatedAt                  apijson.Field
	DateFormed                 apijson.Field
	DateOfBirth                apijson.Field
	DiscardedAt                apijson.Field
	Documents                  apijson.Field
	DoingBusinessAsNames       apijson.Field
	Email                      apijson.Field
	ExpectedActivityVolume     apijson.Field
	FirstName                  apijson.Field
	Identifications            apijson.Field
	IndustryClassifications    apijson.Field
	IntendedUse                apijson.Field
	LastName                   apijson.Field
	LegalEntityType            apijson.Field
	LegalStructure             apijson.Field
	ListedExchange             apijson.Field
	LiveMode                   apijson.Field
	Metadata                   apijson.Field
	MiddleName                 apijson.Field
	Object                     apijson.Field
	OperatingJurisdictions     apijson.Field
	PhoneNumbers               apijson.Field
	PoliticallyExposedPerson   apijson.Field
	PreferredName              apijson.Field
	Prefix                     apijson.Field
	PrimarySocialMediaSites    apijson.Field
	Regulators                 apijson.Field
	RiskRating                 apijson.Field
	Status                     apijson.Field
	Suffix                     apijson.Field
	ThirdPartyVerification     apijson.Field
	TickerSymbol               apijson.Field
	UpdatedAt                  apijson.Field
	WealthAndEmploymentDetails apijson.Field
	Website                    apijson.Field
	LegalEntityAssociations    apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *LegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityJSON) RawJSON() string {
	return r.raw
}

type LegalEntityAddress struct {
	ID string `json:"id,required" format:"uuid"`
	// The types of this address.
	AddressTypes []LegalEntityAddressesAddressType `json:"address_types,required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country     string    `json:"country,required,nullable"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	Line1       string    `json:"line1,required,nullable"`
	Line2       string    `json:"line2,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	Object   string `json:"object,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Region or State.
	Region    string                 `json:"region,required,nullable"`
	UpdatedAt time.Time              `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityAddressJSON `json:"-"`
}

// legalEntityAddressJSON contains the JSON metadata for the struct
// [LegalEntityAddress]
type legalEntityAddressJSON struct {
	ID           apijson.Field
	AddressTypes apijson.Field
	Country      apijson.Field
	CreatedAt    apijson.Field
	DiscardedAt  apijson.Field
	Line1        apijson.Field
	Line2        apijson.Field
	LiveMode     apijson.Field
	Locality     apijson.Field
	Object       apijson.Field
	PostalCode   apijson.Field
	Region       apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LegalEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityAddressJSON) RawJSON() string {
	return r.raw
}

type LegalEntityAddressesAddressType string

const (
	LegalEntityAddressesAddressTypeBusiness    LegalEntityAddressesAddressType = "business"
	LegalEntityAddressesAddressTypeMailing     LegalEntityAddressesAddressType = "mailing"
	LegalEntityAddressesAddressTypeOther       LegalEntityAddressesAddressType = "other"
	LegalEntityAddressesAddressTypePoBox       LegalEntityAddressesAddressType = "po_box"
	LegalEntityAddressesAddressTypeResidential LegalEntityAddressesAddressType = "residential"
)

func (r LegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityAddressesAddressTypeBusiness, LegalEntityAddressesAddressTypeMailing, LegalEntityAddressesAddressTypeOther, LegalEntityAddressesAddressTypePoBox, LegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityIdentification struct {
	ID          string     `json:"id,required" format:"uuid"`
	CreatedAt   time.Time  `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time  `json:"discarded_at,required,nullable" format:"date-time"`
	Documents   []Document `json:"documents,required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate time.Time `json:"expiration_date,required,nullable" format:"date"`
	// The type of ID number.
	IDType LegalEntityIdentificationsIDType `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country,required,nullable"`
	// The region in which the identifcation was issued.
	IssuingRegion string `json:"issuing_region,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                          `json:"live_mode,required"`
	Object    string                        `json:"object,required"`
	UpdatedAt time.Time                     `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityIdentificationJSON `json:"-"`
}

// legalEntityIdentificationJSON contains the JSON metadata for the struct
// [LegalEntityIdentification]
type legalEntityIdentificationJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	DiscardedAt    apijson.Field
	Documents      apijson.Field
	ExpirationDate apijson.Field
	IDType         apijson.Field
	IssuingCountry apijson.Field
	IssuingRegion  apijson.Field
	LiveMode       apijson.Field
	Object         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LegalEntityIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityIdentificationJSON) RawJSON() string {
	return r.raw
}

// The type of ID number.
type LegalEntityIdentificationsIDType string

const (
	LegalEntityIdentificationsIDTypeArCuil         LegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityIdentificationsIDTypeArCuit         LegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityIdentificationsIDTypeBrCnpj         LegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityIdentificationsIDTypeBrCpf          LegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityIdentificationsIDTypeClRun          LegalEntityIdentificationsIDType = "cl_run"
	LegalEntityIdentificationsIDTypeClRut          LegalEntityIdentificationsIDType = "cl_rut"
	LegalEntityIdentificationsIDTypeCoCedulas      LegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityIdentificationsIDTypeCoNit          LegalEntityIdentificationsIDType = "co_nit"
	LegalEntityIdentificationsIDTypeDriversLicense LegalEntityIdentificationsIDType = "drivers_license"
	LegalEntityIdentificationsIDTypeHnID           LegalEntityIdentificationsIDType = "hn_id"
	LegalEntityIdentificationsIDTypeHnRtn          LegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityIdentificationsIDTypeInLei          LegalEntityIdentificationsIDType = "in_lei"
	LegalEntityIdentificationsIDTypeKrBrn          LegalEntityIdentificationsIDType = "kr_brn"
	LegalEntityIdentificationsIDTypeKrCrn          LegalEntityIdentificationsIDType = "kr_crn"
	LegalEntityIdentificationsIDTypeKrRrn          LegalEntityIdentificationsIDType = "kr_rrn"
	LegalEntityIdentificationsIDTypePassport       LegalEntityIdentificationsIDType = "passport"
	LegalEntityIdentificationsIDTypeSaTin          LegalEntityIdentificationsIDType = "sa_tin"
	LegalEntityIdentificationsIDTypeSaVat          LegalEntityIdentificationsIDType = "sa_vat"
	LegalEntityIdentificationsIDTypeUsEin          LegalEntityIdentificationsIDType = "us_ein"
	LegalEntityIdentificationsIDTypeUsItin         LegalEntityIdentificationsIDType = "us_itin"
	LegalEntityIdentificationsIDTypeUsSsn          LegalEntityIdentificationsIDType = "us_ssn"
	LegalEntityIdentificationsIDTypeVnTin          LegalEntityIdentificationsIDType = "vn_tin"
)

func (r LegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityIdentificationsIDTypeArCuil, LegalEntityIdentificationsIDTypeArCuit, LegalEntityIdentificationsIDTypeBrCnpj, LegalEntityIdentificationsIDTypeBrCpf, LegalEntityIdentificationsIDTypeClRun, LegalEntityIdentificationsIDTypeClRut, LegalEntityIdentificationsIDTypeCoCedulas, LegalEntityIdentificationsIDTypeCoNit, LegalEntityIdentificationsIDTypeDriversLicense, LegalEntityIdentificationsIDTypeHnID, LegalEntityIdentificationsIDTypeHnRtn, LegalEntityIdentificationsIDTypeInLei, LegalEntityIdentificationsIDTypeKrBrn, LegalEntityIdentificationsIDTypeKrCrn, LegalEntityIdentificationsIDTypeKrRrn, LegalEntityIdentificationsIDTypePassport, LegalEntityIdentificationsIDTypeSaTin, LegalEntityIdentificationsIDTypeSaVat, LegalEntityIdentificationsIDTypeUsEin, LegalEntityIdentificationsIDTypeUsItin, LegalEntityIdentificationsIDTypeUsSsn, LegalEntityIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

// The type of legal entity.
type LegalEntityLegalEntityType string

const (
	LegalEntityLegalEntityTypeBusiness   LegalEntityLegalEntityType = "business"
	LegalEntityLegalEntityTypeIndividual LegalEntityLegalEntityType = "individual"
	LegalEntityLegalEntityTypeJoint      LegalEntityLegalEntityType = "joint"
)

func (r LegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityLegalEntityTypeBusiness, LegalEntityLegalEntityTypeIndividual, LegalEntityLegalEntityTypeJoint:
		return true
	}
	return false
}

// The business's legal structure.
type LegalEntityLegalStructure string

const (
	LegalEntityLegalStructureCorporation        LegalEntityLegalStructure = "corporation"
	LegalEntityLegalStructureLlc                LegalEntityLegalStructure = "llc"
	LegalEntityLegalStructureNonProfit          LegalEntityLegalStructure = "non_profit"
	LegalEntityLegalStructurePartnership        LegalEntityLegalStructure = "partnership"
	LegalEntityLegalStructureSoleProprietorship LegalEntityLegalStructure = "sole_proprietorship"
	LegalEntityLegalStructureTrust              LegalEntityLegalStructure = "trust"
)

func (r LegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityLegalStructureCorporation, LegalEntityLegalStructureLlc, LegalEntityLegalStructureNonProfit, LegalEntityLegalStructurePartnership, LegalEntityLegalStructureSoleProprietorship, LegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityPhoneNumber struct {
	PhoneNumber string                     `json:"phone_number"`
	JSON        legalEntityPhoneNumberJSON `json:"-"`
}

// legalEntityPhoneNumberJSON contains the JSON metadata for the struct
// [LegalEntityPhoneNumber]
type legalEntityPhoneNumberJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalEntityPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type LegalEntityRegulator struct {
	// The country code where the regulator operates in the ISO 3166-1 alpha-2 format
	// (e.g., "US", "CA", "GB").
	Jurisdiction string `json:"jurisdiction,required"`
	// Full name of the regulatory body.
	Name string `json:"name,required"`
	// Registration or identification number with the regulator.
	RegistrationNumber string                   `json:"registration_number,required"`
	JSON               legalEntityRegulatorJSON `json:"-"`
}

// legalEntityRegulatorJSON contains the JSON metadata for the struct
// [LegalEntityRegulator]
type legalEntityRegulatorJSON struct {
	Jurisdiction       apijson.Field
	Name               apijson.Field
	RegistrationNumber apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LegalEntityRegulator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityRegulatorJSON) RawJSON() string {
	return r.raw
}

// The risk rating of the legal entity. One of low, medium, high.
type LegalEntityRiskRating string

const (
	LegalEntityRiskRatingLow    LegalEntityRiskRating = "low"
	LegalEntityRiskRatingMedium LegalEntityRiskRating = "medium"
	LegalEntityRiskRatingHigh   LegalEntityRiskRating = "high"
)

func (r LegalEntityRiskRating) IsKnown() bool {
	switch r {
	case LegalEntityRiskRatingLow, LegalEntityRiskRatingMedium, LegalEntityRiskRatingHigh:
		return true
	}
	return false
}

// The activation status of the legal entity. One of pending, active, suspended, or
// closed.
type LegalEntityStatus string

const (
	LegalEntityStatusActive    LegalEntityStatus = "active"
	LegalEntityStatusClosed    LegalEntityStatus = "closed"
	LegalEntityStatusPending   LegalEntityStatus = "pending"
	LegalEntityStatusSuspended LegalEntityStatus = "suspended"
)

func (r LegalEntityStatus) IsKnown() bool {
	switch r {
	case LegalEntityStatusActive, LegalEntityStatusClosed, LegalEntityStatusPending, LegalEntityStatusSuspended:
		return true
	}
	return false
}

// Information describing a third-party verification run by an external vendor.
type LegalEntityThirdPartyVerification struct {
	// The vendor that performed the verification, e.g. `persona`.
	Vendor LegalEntityThirdPartyVerificationVendor `json:"vendor,required"`
	// The identification of the third party verification in `vendor`'s system.
	VendorVerificationID string                                `json:"vendor_verification_id,required"`
	JSON                 legalEntityThirdPartyVerificationJSON `json:"-"`
}

// legalEntityThirdPartyVerificationJSON contains the JSON metadata for the struct
// [LegalEntityThirdPartyVerification]
type legalEntityThirdPartyVerificationJSON struct {
	Vendor               apijson.Field
	VendorVerificationID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LegalEntityThirdPartyVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityThirdPartyVerificationJSON) RawJSON() string {
	return r.raw
}

// The vendor that performed the verification, e.g. `persona`.
type LegalEntityThirdPartyVerificationVendor string

const (
	LegalEntityThirdPartyVerificationVendorPersona LegalEntityThirdPartyVerificationVendor = "persona"
)

func (r LegalEntityThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case LegalEntityThirdPartyVerificationVendorPersona:
		return true
	}
	return false
}

type LegalEntityNewParams struct {
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityType] `json:"legal_entity_type,required"`
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[shared.LegalEntityBankSettingsParam]           `json:"bank_settings"`
	// A description of the business.
	BusinessDescription param.Field[string] `json:"business_description"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// The country of citizenship for an individual.
	CitizenshipCountry param.Field[string] `json:"citizenship_country"`
	// The connection ID for the connection the legal entity is associated with.
	// Defaults to the id of the connection designated with an is_default value of true
	// or the id of an existing operational connection if only one is available. Pass
	// in a value of null to prevent the connection from being associated with the
	// legal entity.
	ConnectionID param.Field[string] `json:"connection_id"`
	// The country code where the business is incorporated in the ISO 3166-1 alpha-2 or
	// alpha-3 formats.
	CountryOfIncorporation param.Field[string] `json:"country_of_incorporation"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed param.Field[time.Time] `json:"date_formed" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// Monthly expected transaction volume in USD.
	ExpectedActivityVolume param.Field[int64] `json:"expected_activity_volume"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// A list of identifications for the legal entity.
	Identifications param.Field[[]shared.IdentificationCreateRequestParam] `json:"identifications"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications param.Field[[]shared.LegalEntityIndustryClassificationParam] `json:"industry_classifications"`
	// A description of the intended use of the legal entity.
	IntendedUse param.Field[string] `json:"intended_use"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations param.Field[[]shared.LegalEntityAssociationInlineCreateParam] `json:"legal_entity_associations"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityNewParamsLegalStructure] `json:"legal_structure"`
	// ISO 10383 market identifier code.
	ListedExchange param.Field[string] `json:"listed_exchange"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName param.Field[string] `json:"middle_name"`
	// A list of countries where the business operates (ISO 3166-1 alpha-2 or alpha-3
	// codes).
	OperatingJurisdictions param.Field[[]string]                          `json:"operating_jurisdictions"`
	PhoneNumbers           param.Field[[]LegalEntityNewParamsPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// A list of primary social media URLs for the business.
	PrimarySocialMediaSites param.Field[[]string] `json:"primary_social_media_sites"`
	// Array of regulatory bodies overseeing this institution.
	Regulators param.Field[[]LegalEntityNewParamsRegulator] `json:"regulators"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[LegalEntityNewParamsRiskRating] `json:"risk_rating"`
	// The activation status of the legal entity. One of pending, active, suspended, or
	// closed.
	Status param.Field[LegalEntityNewParamsStatus] `json:"status"`
	// An individual's suffix.
	Suffix param.Field[string] `json:"suffix"`
	// Information describing a third-party verification run by an external vendor.
	ThirdPartyVerification param.Field[LegalEntityNewParamsThirdPartyVerification] `json:"third_party_verification"`
	// Stock ticker symbol for publicly traded companies.
	TickerSymbol               param.Field[string]                                        `json:"ticker_symbol"`
	WealthAndEmploymentDetails param.Field[shared.LegalEntityWealthEmploymentDetailParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of legal entity.
type LegalEntityNewParamsLegalEntityType string

const (
	LegalEntityNewParamsLegalEntityTypeBusiness   LegalEntityNewParamsLegalEntityType = "business"
	LegalEntityNewParamsLegalEntityTypeIndividual LegalEntityNewParamsLegalEntityType = "individual"
)

func (r LegalEntityNewParamsLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityTypeBusiness, LegalEntityNewParamsLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type LegalEntityNewParamsLegalStructure string

const (
	LegalEntityNewParamsLegalStructureCorporation        LegalEntityNewParamsLegalStructure = "corporation"
	LegalEntityNewParamsLegalStructureLlc                LegalEntityNewParamsLegalStructure = "llc"
	LegalEntityNewParamsLegalStructureNonProfit          LegalEntityNewParamsLegalStructure = "non_profit"
	LegalEntityNewParamsLegalStructurePartnership        LegalEntityNewParamsLegalStructure = "partnership"
	LegalEntityNewParamsLegalStructureSoleProprietorship LegalEntityNewParamsLegalStructure = "sole_proprietorship"
	LegalEntityNewParamsLegalStructureTrust              LegalEntityNewParamsLegalStructure = "trust"
)

func (r LegalEntityNewParamsLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalStructureCorporation, LegalEntityNewParamsLegalStructureLlc, LegalEntityNewParamsLegalStructureNonProfit, LegalEntityNewParamsLegalStructurePartnership, LegalEntityNewParamsLegalStructureSoleProprietorship, LegalEntityNewParamsLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityNewParamsPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityNewParamsPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsRegulator struct {
	// The country code where the regulator operates in the ISO 3166-1 alpha-2 format
	// (e.g., "US", "CA", "GB").
	Jurisdiction param.Field[string] `json:"jurisdiction,required"`
	// Full name of the regulatory body.
	Name param.Field[string] `json:"name,required"`
	// Registration or identification number with the regulator.
	RegistrationNumber param.Field[string] `json:"registration_number,required"`
}

func (r LegalEntityNewParamsRegulator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The risk rating of the legal entity. One of low, medium, high.
type LegalEntityNewParamsRiskRating string

const (
	LegalEntityNewParamsRiskRatingLow    LegalEntityNewParamsRiskRating = "low"
	LegalEntityNewParamsRiskRatingMedium LegalEntityNewParamsRiskRating = "medium"
	LegalEntityNewParamsRiskRatingHigh   LegalEntityNewParamsRiskRating = "high"
)

func (r LegalEntityNewParamsRiskRating) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsRiskRatingLow, LegalEntityNewParamsRiskRatingMedium, LegalEntityNewParamsRiskRatingHigh:
		return true
	}
	return false
}

// The activation status of the legal entity. One of pending, active, suspended, or
// closed.
type LegalEntityNewParamsStatus string

const (
	LegalEntityNewParamsStatusActive    LegalEntityNewParamsStatus = "active"
	LegalEntityNewParamsStatusClosed    LegalEntityNewParamsStatus = "closed"
	LegalEntityNewParamsStatusPending   LegalEntityNewParamsStatus = "pending"
	LegalEntityNewParamsStatusSuspended LegalEntityNewParamsStatus = "suspended"
)

func (r LegalEntityNewParamsStatus) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsStatusActive, LegalEntityNewParamsStatusClosed, LegalEntityNewParamsStatusPending, LegalEntityNewParamsStatusSuspended:
		return true
	}
	return false
}

// Information describing a third-party verification run by an external vendor.
type LegalEntityNewParamsThirdPartyVerification struct {
	// The vendor that performed the verification, e.g. `persona`.
	Vendor param.Field[LegalEntityNewParamsThirdPartyVerificationVendor] `json:"vendor,required"`
	// The identification of the third party verification in `vendor`'s system.
	VendorVerificationID param.Field[string] `json:"vendor_verification_id,required"`
}

func (r LegalEntityNewParamsThirdPartyVerification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The vendor that performed the verification, e.g. `persona`.
type LegalEntityNewParamsThirdPartyVerificationVendor string

const (
	LegalEntityNewParamsThirdPartyVerificationVendorPersona LegalEntityNewParamsThirdPartyVerificationVendor = "persona"
)

func (r LegalEntityNewParamsThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsThirdPartyVerificationVendorPersona:
		return true
	}
	return false
}

type LegalEntityUpdateParams struct {
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[shared.LegalEntityBankSettingsParam]           `json:"bank_settings"`
	// A description of the business.
	BusinessDescription param.Field[string] `json:"business_description"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// The country of citizenship for an individual.
	CitizenshipCountry param.Field[string] `json:"citizenship_country"`
	// The country code where the business is incorporated in the ISO 3166-1 alpha-2 or
	// alpha-3 formats.
	CountryOfIncorporation param.Field[string] `json:"country_of_incorporation"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed param.Field[time.Time] `json:"date_formed" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// Monthly expected transaction volume in USD.
	ExpectedActivityVolume param.Field[int64] `json:"expected_activity_volume"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// A list of identifications for the legal entity.
	Identifications param.Field[[]shared.IdentificationCreateRequestParam] `json:"identifications"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications param.Field[[]shared.LegalEntityIndustryClassificationParam] `json:"industry_classifications"`
	// A description of the intended use of the legal entity.
	IntendedUse param.Field[string] `json:"intended_use"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityUpdateParamsLegalStructure] `json:"legal_structure"`
	// ISO 10383 market identifier code.
	ListedExchange param.Field[string] `json:"listed_exchange"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName param.Field[string] `json:"middle_name"`
	// A list of countries where the business operates (ISO 3166-1 alpha-2 or alpha-3
	// codes).
	OperatingJurisdictions param.Field[[]string]                             `json:"operating_jurisdictions"`
	PhoneNumbers           param.Field[[]LegalEntityUpdateParamsPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// A list of primary social media URLs for the business.
	PrimarySocialMediaSites param.Field[[]string] `json:"primary_social_media_sites"`
	// Array of regulatory bodies overseeing this institution.
	Regulators param.Field[[]LegalEntityUpdateParamsRegulator] `json:"regulators"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[LegalEntityUpdateParamsRiskRating] `json:"risk_rating"`
	// The activation status of the legal entity. One of pending, active, suspended, or
	// closed.
	Status param.Field[LegalEntityUpdateParamsStatus] `json:"status"`
	// An individual's suffix.
	Suffix param.Field[string] `json:"suffix"`
	// Information describing a third-party verification run by an external vendor.
	ThirdPartyVerification param.Field[LegalEntityUpdateParamsThirdPartyVerification] `json:"third_party_verification"`
	// Stock ticker symbol for publicly traded companies.
	TickerSymbol               param.Field[string]                                        `json:"ticker_symbol"`
	WealthAndEmploymentDetails param.Field[shared.LegalEntityWealthEmploymentDetailParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The business's legal structure.
type LegalEntityUpdateParamsLegalStructure string

const (
	LegalEntityUpdateParamsLegalStructureCorporation        LegalEntityUpdateParamsLegalStructure = "corporation"
	LegalEntityUpdateParamsLegalStructureLlc                LegalEntityUpdateParamsLegalStructure = "llc"
	LegalEntityUpdateParamsLegalStructureNonProfit          LegalEntityUpdateParamsLegalStructure = "non_profit"
	LegalEntityUpdateParamsLegalStructurePartnership        LegalEntityUpdateParamsLegalStructure = "partnership"
	LegalEntityUpdateParamsLegalStructureSoleProprietorship LegalEntityUpdateParamsLegalStructure = "sole_proprietorship"
	LegalEntityUpdateParamsLegalStructureTrust              LegalEntityUpdateParamsLegalStructure = "trust"
)

func (r LegalEntityUpdateParamsLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsLegalStructureCorporation, LegalEntityUpdateParamsLegalStructureLlc, LegalEntityUpdateParamsLegalStructureNonProfit, LegalEntityUpdateParamsLegalStructurePartnership, LegalEntityUpdateParamsLegalStructureSoleProprietorship, LegalEntityUpdateParamsLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityUpdateParamsPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityUpdateParamsPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityUpdateParamsRegulator struct {
	// The country code where the regulator operates in the ISO 3166-1 alpha-2 format
	// (e.g., "US", "CA", "GB").
	Jurisdiction param.Field[string] `json:"jurisdiction,required"`
	// Full name of the regulatory body.
	Name param.Field[string] `json:"name,required"`
	// Registration or identification number with the regulator.
	RegistrationNumber param.Field[string] `json:"registration_number,required"`
}

func (r LegalEntityUpdateParamsRegulator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The risk rating of the legal entity. One of low, medium, high.
type LegalEntityUpdateParamsRiskRating string

const (
	LegalEntityUpdateParamsRiskRatingLow    LegalEntityUpdateParamsRiskRating = "low"
	LegalEntityUpdateParamsRiskRatingMedium LegalEntityUpdateParamsRiskRating = "medium"
	LegalEntityUpdateParamsRiskRatingHigh   LegalEntityUpdateParamsRiskRating = "high"
)

func (r LegalEntityUpdateParamsRiskRating) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsRiskRatingLow, LegalEntityUpdateParamsRiskRatingMedium, LegalEntityUpdateParamsRiskRatingHigh:
		return true
	}
	return false
}

// The activation status of the legal entity. One of pending, active, suspended, or
// closed.
type LegalEntityUpdateParamsStatus string

const (
	LegalEntityUpdateParamsStatusActive    LegalEntityUpdateParamsStatus = "active"
	LegalEntityUpdateParamsStatusClosed    LegalEntityUpdateParamsStatus = "closed"
	LegalEntityUpdateParamsStatusPending   LegalEntityUpdateParamsStatus = "pending"
	LegalEntityUpdateParamsStatusSuspended LegalEntityUpdateParamsStatus = "suspended"
)

func (r LegalEntityUpdateParamsStatus) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsStatusActive, LegalEntityUpdateParamsStatusClosed, LegalEntityUpdateParamsStatusPending, LegalEntityUpdateParamsStatusSuspended:
		return true
	}
	return false
}

// Information describing a third-party verification run by an external vendor.
type LegalEntityUpdateParamsThirdPartyVerification struct {
	// The vendor that performed the verification, e.g. `persona`.
	Vendor param.Field[LegalEntityUpdateParamsThirdPartyVerificationVendor] `json:"vendor,required"`
	// The identification of the third party verification in `vendor`'s system.
	VendorVerificationID param.Field[string] `json:"vendor_verification_id,required"`
}

func (r LegalEntityUpdateParamsThirdPartyVerification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The vendor that performed the verification, e.g. `persona`.
type LegalEntityUpdateParamsThirdPartyVerificationVendor string

const (
	LegalEntityUpdateParamsThirdPartyVerificationVendorPersona LegalEntityUpdateParamsThirdPartyVerificationVendor = "persona"
)

func (r LegalEntityUpdateParamsThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsThirdPartyVerificationVendorPersona:
		return true
	}
	return false
}

type LegalEntityListParams struct {
	AfterCursor     param.Field[string]                               `query:"after_cursor"`
	LegalEntityType param.Field[LegalEntityListParamsLegalEntityType] `query:"legal_entity_type"`
	// For example, if you want to query for records with metadata key `Type` and value
	// `Loan`, the query would be `metadata%5BType%5D=Loan`. This encodes the query
	// parameters.
	Metadata    param.Field[map[string]string]           `query:"metadata"`
	PerPage     param.Field[int64]                       `query:"per_page"`
	ShowDeleted param.Field[string]                      `query:"show_deleted"`
	Status      param.Field[LegalEntityListParamsStatus] `query:"status"`
}

// URLQuery serializes [LegalEntityListParams]'s query parameters as `url.Values`.
func (r LegalEntityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LegalEntityListParamsLegalEntityType string

const (
	LegalEntityListParamsLegalEntityTypeBusiness   LegalEntityListParamsLegalEntityType = "business"
	LegalEntityListParamsLegalEntityTypeIndividual LegalEntityListParamsLegalEntityType = "individual"
)

func (r LegalEntityListParamsLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityListParamsLegalEntityTypeBusiness, LegalEntityListParamsLegalEntityTypeIndividual:
		return true
	}
	return false
}

type LegalEntityListParamsStatus string

const (
	LegalEntityListParamsStatusPending   LegalEntityListParamsStatus = "pending"
	LegalEntityListParamsStatusActive    LegalEntityListParamsStatus = "active"
	LegalEntityListParamsStatusSuspended LegalEntityListParamsStatus = "suspended"
	LegalEntityListParamsStatusClosed    LegalEntityListParamsStatus = "closed"
)

func (r LegalEntityListParamsStatus) IsKnown() bool {
	switch r {
	case LegalEntityListParamsStatusPending, LegalEntityListParamsStatusActive, LegalEntityListParamsStatusSuspended, LegalEntityListParamsStatusClosed:
		return true
	}
	return false
}
