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

// Create a legal entity. All country fields use ISO 3166-1 alpha-2 (e.g. US).
func (r *LegalEntityService) New(ctx context.Context, body LegalEntityNewParams, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/legal_entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get details on a single legal entity.
func (r *LegalEntityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a legal entity.
func (r *LegalEntityService) Update(ctx context.Context, id string, body LegalEntityUpdateParams, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
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

// Update Legal Entity Status (sandbox only)
func (r *LegalEntityService) UpdateStatus(ctx context.Context, id string, body LegalEntityUpdateStatusParams, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/simulations/legal_entities/%s/update_status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type LegalEntity struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// A list of addresses for the entity.
	Addresses    []LegalEntityAddress           `json:"addresses" api:"required"`
	BankSettings shared.LegalEntityBankSettings `json:"bank_settings" api:"required,nullable"`
	// A description of the business.
	BusinessDescription string `json:"business_description" api:"required,nullable"`
	// Legal designation associated with the business.
	BusinessDesignation LegalEntityBusinessDesignation `json:"business_designation" api:"required,nullable"`
	// The business's legal business name.
	BusinessName string `json:"business_name" api:"required,nullable"`
	// The country of citizenship for an individual.
	CitizenshipCountry string `json:"citizenship_country" api:"required,nullable"`
	// Deprecated: deprecated
	ComplianceDetails interface{} `json:"compliance_details" api:"required,nullable"`
	// The country where the business is incorporated, as an ISO 3166-1 alpha-2 country
	// code (e.g. US).
	CountryOfIncorporation string    `json:"country_of_incorporation" api:"required,nullable"`
	CreatedAt              time.Time `json:"created_at" api:"required" format:"date-time"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed time.Time `json:"date_formed" api:"required,nullable" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          time.Time  `json:"date_of_birth" api:"required,nullable" format:"date"`
	DiscardedAt          time.Time  `json:"discarded_at" api:"required,nullable" format:"date-time"`
	Documents            []Document `json:"documents" api:"required"`
	DoingBusinessAsNames []string   `json:"doing_business_as_names" api:"required"`
	// The entity's primary email.
	Email string `json:"email" api:"required,nullable"`
	// Monthly expected transaction volume in USD.
	ExpectedActivityVolume int64 `json:"expected_activity_volume" api:"required,nullable"`
	// An optional user-defined 180 character unique identifier.
	ExternalID string `json:"external_id" api:"required,nullable"`
	// An individual's first name.
	FirstName string `json:"first_name" api:"required,nullable"`
	// A list of identifications for the legal entity.
	Identifications []Identification `json:"identifications" api:"required"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications []shared.LegalEntityIndustryClassification `json:"industry_classifications" api:"required"`
	// A description of the intended use of the legal entity.
	IntendedUse string `json:"intended_use" api:"required,nullable"`
	// An individual's last name.
	LastName string `json:"last_name" api:"required,nullable"`
	// The type of legal entity.
	LegalEntityType LegalEntityLegalEntityType `json:"legal_entity_type" api:"required"`
	// The business's legal structure.
	LegalStructure LegalEntityLegalStructure `json:"legal_structure" api:"required,nullable"`
	// ISO 10383 market identifier code.
	ListedExchange string `json:"listed_exchange" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode" api:"required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata" api:"required"`
	// An individual's middle name.
	MiddleName string `json:"middle_name" api:"required,nullable"`
	Object     string `json:"object" api:"required"`
	// A list of countries where the business operates, as ISO 3166-1 alpha-2 country
	// codes (e.g. ["US", "CA"]).
	OperatingJurisdictions []string                 `json:"operating_jurisdictions" api:"required"`
	PhoneNumbers           []LegalEntityPhoneNumber `json:"phone_numbers" api:"required"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson bool `json:"politically_exposed_person" api:"required,nullable"`
	// An individual's preferred name.
	PreferredName string `json:"preferred_name" api:"required,nullable"`
	// An individual's prefix.
	Prefix string `json:"prefix" api:"required,nullable"`
	// A list of primary social media URLs for the business.
	PrimarySocialMediaSites []string `json:"primary_social_media_sites" api:"required"`
	// Array of regulatory bodies overseeing this institution.
	Regulators []LegalEntityRegulator `json:"regulators" api:"required,nullable"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating LegalEntityRiskRating `json:"risk_rating" api:"required,nullable"`
	// The UUID of the parent legal entity in the service provider tree.
	ServiceProviderLegalEntityID string `json:"service_provider_legal_entity_id" api:"required,nullable" format:"uuid"`
	// The activation status of the legal entity. One of pending, active, suspended, or
	// denied.
	Status LegalEntityStatus `json:"status" api:"required,nullable"`
	// An individual's suffix.
	Suffix string `json:"suffix" api:"required,nullable"`
	// Acceptance of terms of use by the legal entity.
	TermsOfUse LegalEntityTermsOfUse `json:"terms_of_use" api:"required,nullable"`
	// Deprecated. Use `third_party_verifications` instead.
	//
	// Deprecated: deprecated
	ThirdPartyVerification shared.ThirdPartyVerification `json:"third_party_verification" api:"required,nullable"`
	// A list of third-party verifications run by external vendors.
	ThirdPartyVerifications []shared.ThirdPartyVerification `json:"third_party_verifications" api:"required"`
	// Stock ticker symbol for publicly traded companies.
	TickerSymbol               string                                   `json:"ticker_symbol" api:"required,nullable"`
	UpdatedAt                  time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	WealthAndEmploymentDetails shared.LegalEntityWealthEmploymentDetail `json:"wealth_and_employment_details" api:"required,nullable"`
	// The entity's primary website URL.
	Website string `json:"website" api:"required,nullable"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations []LegalEntityAssociation `json:"legal_entity_associations" api:"nullable"`
	JSON                    legalEntityJSON          `json:"-"`
}

// legalEntityJSON contains the JSON metadata for the struct [LegalEntity]
type legalEntityJSON struct {
	ID                           apijson.Field
	Addresses                    apijson.Field
	BankSettings                 apijson.Field
	BusinessDescription          apijson.Field
	BusinessDesignation          apijson.Field
	BusinessName                 apijson.Field
	CitizenshipCountry           apijson.Field
	ComplianceDetails            apijson.Field
	CountryOfIncorporation       apijson.Field
	CreatedAt                    apijson.Field
	DateFormed                   apijson.Field
	DateOfBirth                  apijson.Field
	DiscardedAt                  apijson.Field
	Documents                    apijson.Field
	DoingBusinessAsNames         apijson.Field
	Email                        apijson.Field
	ExpectedActivityVolume       apijson.Field
	ExternalID                   apijson.Field
	FirstName                    apijson.Field
	Identifications              apijson.Field
	IndustryClassifications      apijson.Field
	IntendedUse                  apijson.Field
	LastName                     apijson.Field
	LegalEntityType              apijson.Field
	LegalStructure               apijson.Field
	ListedExchange               apijson.Field
	LiveMode                     apijson.Field
	Metadata                     apijson.Field
	MiddleName                   apijson.Field
	Object                       apijson.Field
	OperatingJurisdictions       apijson.Field
	PhoneNumbers                 apijson.Field
	PoliticallyExposedPerson     apijson.Field
	PreferredName                apijson.Field
	Prefix                       apijson.Field
	PrimarySocialMediaSites      apijson.Field
	Regulators                   apijson.Field
	RiskRating                   apijson.Field
	ServiceProviderLegalEntityID apijson.Field
	Status                       apijson.Field
	Suffix                       apijson.Field
	TermsOfUse                   apijson.Field
	ThirdPartyVerification       apijson.Field
	ThirdPartyVerifications      apijson.Field
	TickerSymbol                 apijson.Field
	UpdatedAt                    apijson.Field
	WealthAndEmploymentDetails   apijson.Field
	Website                      apijson.Field
	LegalEntityAssociations      apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *LegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityJSON) RawJSON() string {
	return r.raw
}

type LegalEntityAddress struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The types of this address.
	AddressTypes []LegalEntityAddressesAddressType `json:"address_types" api:"required"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country     string    `json:"country" api:"required,nullable"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at" api:"required,nullable" format:"date-time"`
	Line1       string    `json:"line1" api:"required,nullable"`
	Line2       string    `json:"line2" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode" api:"required"`
	// Locality or City. Use the full city name rather than an abbreviation (e.g. San
	// Francisco).
	Locality string `json:"locality" api:"required,nullable"`
	Object   string `json:"object" api:"required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code" api:"required,nullable"`
	// Whether this address is the primary address for the legal entity. Optional; when
	// omitted it is inferred from the address types.
	Primary bool `json:"primary" api:"required,nullable"`
	// Region or State. This field is free-form; for US states, we recommend a
	// two-letter code (e.g. CA). Full state names are also accepted.
	Region    string                 `json:"region" api:"required,nullable"`
	UpdatedAt time.Time              `json:"updated_at" api:"required" format:"date-time"`
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
	Primary      apijson.Field
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
	LegalEntityAddressesAddressTypeBusiness           LegalEntityAddressesAddressType = "business"
	LegalEntityAddressesAddressTypeBusinessPhysical   LegalEntityAddressesAddressType = "business_physical"
	LegalEntityAddressesAddressTypeBusinessRegistered LegalEntityAddressesAddressType = "business_registered"
	LegalEntityAddressesAddressTypeMailing            LegalEntityAddressesAddressType = "mailing"
	LegalEntityAddressesAddressTypeOther              LegalEntityAddressesAddressType = "other"
	LegalEntityAddressesAddressTypePoBox              LegalEntityAddressesAddressType = "po_box"
	LegalEntityAddressesAddressTypeResidential        LegalEntityAddressesAddressType = "residential"
)

func (r LegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityAddressesAddressTypeBusiness, LegalEntityAddressesAddressTypeBusinessPhysical, LegalEntityAddressesAddressTypeBusinessRegistered, LegalEntityAddressesAddressTypeMailing, LegalEntityAddressesAddressTypeOther, LegalEntityAddressesAddressTypePoBox, LegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

// Legal designation associated with the business.
type LegalEntityBusinessDesignation string

const (
	LegalEntityBusinessDesignationExemptFinancialInstitution LegalEntityBusinessDesignation = "exempt_financial_institution"
	LegalEntityBusinessDesignationNonOperatingBusiness       LegalEntityBusinessDesignation = "non_operating_business"
)

func (r LegalEntityBusinessDesignation) IsKnown() bool {
	switch r {
	case LegalEntityBusinessDesignationExemptFinancialInstitution, LegalEntityBusinessDesignationNonOperatingBusiness:
		return true
	}
	return false
}

// The type of legal entity.
type LegalEntityLegalEntityType string

const (
	LegalEntityLegalEntityTypeBusiness   LegalEntityLegalEntityType = "business"
	LegalEntityLegalEntityTypeIndividual LegalEntityLegalEntityType = "individual"
)

func (r LegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityLegalEntityTypeBusiness, LegalEntityLegalEntityTypeIndividual:
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
	// A phone number in E.164 format. This format is strictly validated: include a
	// leading + and country code, followed by digits only (no spaces or dashes), e.g.
	// +12025551234.
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
	Jurisdiction string `json:"jurisdiction" api:"required"`
	// Full name of the regulatory body.
	Name string `json:"name" api:"required"`
	// Registration or identification number with the regulator.
	RegistrationNumber string                   `json:"registration_number" api:"required"`
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
// denied.
type LegalEntityStatus string

const (
	LegalEntityStatusActive    LegalEntityStatus = "active"
	LegalEntityStatusDenied    LegalEntityStatus = "denied"
	LegalEntityStatusPending   LegalEntityStatus = "pending"
	LegalEntityStatusSuspended LegalEntityStatus = "suspended"
)

func (r LegalEntityStatus) IsKnown() bool {
	switch r {
	case LegalEntityStatusActive, LegalEntityStatusDenied, LegalEntityStatusPending, LegalEntityStatusSuspended:
		return true
	}
	return false
}

// Acceptance of terms of use by the legal entity.
type LegalEntityTermsOfUse struct {
	// The ISO 8601 timestamp indicating when the terms of use were accepted.
	AcceptedAt time.Time `json:"accepted_at" format:"date-time"`
	// The IP address from which the terms of use were accepted. Supports both IPv4 and
	// IPv6 formats.
	IPAddress string                    `json:"ip_address"`
	JSON      legalEntityTermsOfUseJSON `json:"-"`
}

// legalEntityTermsOfUseJSON contains the JSON metadata for the struct
// [LegalEntityTermsOfUse]
type legalEntityTermsOfUseJSON struct {
	AcceptedAt  apijson.Field
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalEntityTermsOfUse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityTermsOfUseJSON) RawJSON() string {
	return r.raw
}

type LegalEntityNewParams struct {
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityType] `json:"legal_entity_type" api:"required"`
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[shared.LegalEntityBankSettingsParam]           `json:"bank_settings"`
	// A description of the business.
	BusinessDescription param.Field[string] `json:"business_description"`
	// Legal designation associated with the business.
	BusinessDesignation param.Field[LegalEntityNewParamsBusinessDesignation] `json:"business_designation"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// The country of citizenship for an individual.
	CitizenshipCountry param.Field[string]      `json:"citizenship_country"`
	ComplianceDetails  param.Field[interface{}] `json:"compliance_details"`
	// The connection ID for the connection the legal entity is associated with.
	// Defaults to the id of the connection designated with an is_default value of true
	// or the id of an existing operational connection if only one is available. Pass
	// in a value of null to prevent the connection from being associated with the
	// legal entity.
	ConnectionID param.Field[string] `json:"connection_id"`
	// The country where the business is incorporated, as an ISO 3166-1 alpha-2 country
	// code (e.g. US).
	CountryOfIncorporation param.Field[string] `json:"country_of_incorporation"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed param.Field[time.Time] `json:"date_formed" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth param.Field[time.Time] `json:"date_of_birth" format:"date"`
	// A list of documents to attach to the legal entity (e.g. articles of
	// incorporation, certificate of good standing, proof of address).
	Documents            param.Field[[]LegalEntityNewParamsDocument] `json:"documents"`
	DoingBusinessAsNames param.Field[[]string]                       `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// Monthly expected transaction volume in USD.
	ExpectedActivityVolume param.Field[int64] `json:"expected_activity_volume"`
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
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
	// A list of countries where the business operates, as ISO 3166-1 alpha-2 country
	// codes (e.g. ["US", "CA"]).
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
	// The UUID of the parent legal entity in the service provider tree.
	ServiceProviderLegalEntityID param.Field[string] `json:"service_provider_legal_entity_id" format:"uuid"`
	// An individual's suffix.
	Suffix param.Field[string] `json:"suffix"`
	// Acceptance of terms of use by the legal entity.
	TermsOfUse param.Field[LegalEntityNewParamsTermsOfUse] `json:"terms_of_use"`
	// Deprecated. Use `third_party_verifications` instead.
	ThirdPartyVerification param.Field[shared.ThirdPartyVerificationParam] `json:"third_party_verification"`
	// A list of third-party verifications run by external vendors.
	ThirdPartyVerifications param.Field[[]shared.ThirdPartyVerificationParam] `json:"third_party_verifications"`
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

// Legal designation associated with the business.
type LegalEntityNewParamsBusinessDesignation string

const (
	LegalEntityNewParamsBusinessDesignationExemptFinancialInstitution LegalEntityNewParamsBusinessDesignation = "exempt_financial_institution"
	LegalEntityNewParamsBusinessDesignationNonOperatingBusiness       LegalEntityNewParamsBusinessDesignation = "non_operating_business"
)

func (r LegalEntityNewParamsBusinessDesignation) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsBusinessDesignationExemptFinancialInstitution, LegalEntityNewParamsBusinessDesignationNonOperatingBusiness:
		return true
	}
	return false
}

type LegalEntityNewParamsDocument struct {
	// A category given to the document, can be `null`.
	DocumentType param.Field[LegalEntityNewParamsDocumentsDocumentType] `json:"document_type" api:"required"`
	// Base64-encoded file content for the document.
	FileData param.Field[string] `json:"file_data" api:"required"`
	// The original filename of the document.
	Filename param.Field[string] `json:"filename"`
}

func (r LegalEntityNewParamsDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category given to the document, can be `null`.
type LegalEntityNewParamsDocumentsDocumentType string

const (
	LegalEntityNewParamsDocumentsDocumentTypeArticlesOfIncorporation   LegalEntityNewParamsDocumentsDocumentType = "articles_of_incorporation"
	LegalEntityNewParamsDocumentsDocumentTypeCertificateOfGoodStanding LegalEntityNewParamsDocumentsDocumentType = "certificate_of_good_standing"
	LegalEntityNewParamsDocumentsDocumentTypeEinLetter                 LegalEntityNewParamsDocumentsDocumentType = "ein_letter"
	LegalEntityNewParamsDocumentsDocumentTypeGeneric                   LegalEntityNewParamsDocumentsDocumentType = "generic"
	LegalEntityNewParamsDocumentsDocumentTypeIdentificationBack        LegalEntityNewParamsDocumentsDocumentType = "identification_back"
	LegalEntityNewParamsDocumentsDocumentTypeIdentificationFront       LegalEntityNewParamsDocumentsDocumentType = "identification_front"
	LegalEntityNewParamsDocumentsDocumentTypeProofOfAddress            LegalEntityNewParamsDocumentsDocumentType = "proof_of_address"
)

func (r LegalEntityNewParamsDocumentsDocumentType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsDocumentsDocumentTypeArticlesOfIncorporation, LegalEntityNewParamsDocumentsDocumentTypeCertificateOfGoodStanding, LegalEntityNewParamsDocumentsDocumentTypeEinLetter, LegalEntityNewParamsDocumentsDocumentTypeGeneric, LegalEntityNewParamsDocumentsDocumentTypeIdentificationBack, LegalEntityNewParamsDocumentsDocumentTypeIdentificationFront, LegalEntityNewParamsDocumentsDocumentTypeProofOfAddress:
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
	// A phone number in E.164 format. This format is strictly validated: include a
	// leading + and country code, followed by digits only (no spaces or dashes), e.g.
	// +12025551234.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityNewParamsPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsRegulator struct {
	// The country code where the regulator operates in the ISO 3166-1 alpha-2 format
	// (e.g., "US", "CA", "GB").
	Jurisdiction param.Field[string] `json:"jurisdiction" api:"required"`
	// Full name of the regulatory body.
	Name param.Field[string] `json:"name" api:"required"`
	// Registration or identification number with the regulator.
	RegistrationNumber param.Field[string] `json:"registration_number" api:"required"`
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

// Acceptance of terms of use by the legal entity.
type LegalEntityNewParamsTermsOfUse struct {
	// The ISO 8601 timestamp indicating when the terms of use were accepted.
	AcceptedAt param.Field[time.Time] `json:"accepted_at" format:"date-time"`
	// The IP address from which the terms of use were accepted. Supports both IPv4 and
	// IPv6 formats.
	IPAddress param.Field[string] `json:"ip_address"`
}

func (r LegalEntityNewParamsTermsOfUse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityUpdateParams struct {
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[shared.LegalEntityBankSettingsParam]           `json:"bank_settings"`
	// A description of the business.
	BusinessDescription param.Field[string] `json:"business_description"`
	// Legal designation associated with the business.
	BusinessDesignation param.Field[LegalEntityUpdateParamsBusinessDesignation] `json:"business_designation"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// The country of citizenship for an individual.
	CitizenshipCountry param.Field[string] `json:"citizenship_country"`
	// The country where the business is incorporated, as an ISO 3166-1 alpha-2 country
	// code (e.g. US).
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
	// An optional user-defined 180 character unique identifier.
	ExternalID param.Field[string] `json:"external_id"`
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
	// A list of countries where the business operates, as ISO 3166-1 alpha-2 country
	// codes (e.g. ["US", "CA"]).
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
	// The UUID of the parent legal entity in the service provider tree.
	ServiceProviderLegalEntityID param.Field[string] `json:"service_provider_legal_entity_id" format:"uuid"`
	// An individual's suffix.
	Suffix param.Field[string] `json:"suffix"`
	// Acceptance of terms of use by the legal entity.
	TermsOfUse param.Field[LegalEntityUpdateParamsTermsOfUse] `json:"terms_of_use"`
	// Deprecated. Use `third_party_verifications` instead.
	ThirdPartyVerification param.Field[shared.ThirdPartyVerificationParam] `json:"third_party_verification"`
	// A list of third-party verifications run by external vendors.
	ThirdPartyVerifications param.Field[[]shared.ThirdPartyVerificationParam] `json:"third_party_verifications"`
	// Stock ticker symbol for publicly traded companies.
	TickerSymbol               param.Field[string]                                        `json:"ticker_symbol"`
	WealthAndEmploymentDetails param.Field[shared.LegalEntityWealthEmploymentDetailParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Legal designation associated with the business.
type LegalEntityUpdateParamsBusinessDesignation string

const (
	LegalEntityUpdateParamsBusinessDesignationExemptFinancialInstitution LegalEntityUpdateParamsBusinessDesignation = "exempt_financial_institution"
	LegalEntityUpdateParamsBusinessDesignationNonOperatingBusiness       LegalEntityUpdateParamsBusinessDesignation = "non_operating_business"
)

func (r LegalEntityUpdateParamsBusinessDesignation) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsBusinessDesignationExemptFinancialInstitution, LegalEntityUpdateParamsBusinessDesignationNonOperatingBusiness:
		return true
	}
	return false
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
	// A phone number in E.164 format. This format is strictly validated: include a
	// leading + and country code, followed by digits only (no spaces or dashes), e.g.
	// +12025551234.
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityUpdateParamsPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityUpdateParamsRegulator struct {
	// The country code where the regulator operates in the ISO 3166-1 alpha-2 format
	// (e.g., "US", "CA", "GB").
	Jurisdiction param.Field[string] `json:"jurisdiction" api:"required"`
	// Full name of the regulatory body.
	Name param.Field[string] `json:"name" api:"required"`
	// Registration or identification number with the regulator.
	RegistrationNumber param.Field[string] `json:"registration_number" api:"required"`
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

// Acceptance of terms of use by the legal entity.
type LegalEntityUpdateParamsTermsOfUse struct {
	// The ISO 8601 timestamp indicating when the terms of use were accepted.
	AcceptedAt param.Field[time.Time] `json:"accepted_at" format:"date-time"`
	// The IP address from which the terms of use were accepted. Supports both IPv4 and
	// IPv6 formats.
	IPAddress param.Field[string] `json:"ip_address"`
}

func (r LegalEntityUpdateParamsTermsOfUse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// An optional user-defined 180 character unique identifier.
	ExternalID      param.Field[string]                               `query:"external_id"`
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
	LegalEntityListParamsStatusDenied    LegalEntityListParamsStatus = "denied"
)

func (r LegalEntityListParamsStatus) IsKnown() bool {
	switch r {
	case LegalEntityListParamsStatusPending, LegalEntityListParamsStatusActive, LegalEntityListParamsStatusSuspended, LegalEntityListParamsStatusDenied:
		return true
	}
	return false
}

type LegalEntityUpdateStatusParams struct {
	// The target status for the legal entity. One of `active`, `suspended`, or
	// `denied`. Valid transitions depend on the current status.
	Status param.Field[LegalEntityUpdateStatusParamsStatus] `json:"status" api:"required"`
}

func (r LegalEntityUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target status for the legal entity. One of `active`, `suspended`, or
// `denied`. Valid transitions depend on the current status.
type LegalEntityUpdateStatusParamsStatus string

const (
	LegalEntityUpdateStatusParamsStatusActive    LegalEntityUpdateStatusParamsStatus = "active"
	LegalEntityUpdateStatusParamsStatusSuspended LegalEntityUpdateStatusParamsStatus = "suspended"
	LegalEntityUpdateStatusParamsStatusDenied    LegalEntityUpdateStatusParamsStatus = "denied"
)

func (r LegalEntityUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case LegalEntityUpdateStatusParamsStatusActive, LegalEntityUpdateStatusParamsStatusSuspended, LegalEntityUpdateStatusParamsStatusDenied:
		return true
	}
	return false
}
