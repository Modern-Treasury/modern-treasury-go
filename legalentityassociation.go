// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// LegalEntityAssociationService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegalEntityAssociationService] method instead.
type LegalEntityAssociationService struct {
	Options []option.RequestOption
}

// NewLegalEntityAssociationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLegalEntityAssociationService(opts ...option.RequestOption) (r *LegalEntityAssociationService) {
	r = &LegalEntityAssociationService{}
	r.Options = opts
	return
}

// create legal_entity_association
func (r *LegalEntityAssociationService) New(ctx context.Context, body LegalEntityAssociationNewParams, opts ...option.RequestOption) (res *LegalEntityAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/legal_entity_associations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChildLegalEntity struct {
	ID string `json:"id,required" format:"uuid"`
	// A list of addresses for the entity.
	Addresses    []ChildLegalEntityAddress      `json:"addresses,required"`
	BankSettings shared.LegalEntityBankSettings `json:"bank_settings,required,nullable"`
	// A description of the business.
	BusinessDescription string `json:"business_description,required,nullable"`
	// The business's legal business name.
	BusinessName string `json:"business_name,required,nullable"`
	// The country of citizenship for an individual.
	CitizenshipCountry string                             `json:"citizenship_country,required,nullable"`
	ComplianceDetails  shared.LegalEntityComplianceDetail `json:"compliance_details,required,nullable"`
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
	Identifications []ChildLegalEntityIdentification `json:"identifications,required"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications []shared.LegalEntityIndustryClassification `json:"industry_classifications,required"`
	// A description of the intended use of the legal entity.
	IntendedUse string `json:"intended_use,required,nullable"`
	// An individual's last name.
	LastName string `json:"last_name,required,nullable"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations []LegalEntityAssociation `json:"legal_entity_associations,required,nullable"`
	// The type of legal entity.
	LegalEntityType ChildLegalEntityLegalEntityType `json:"legal_entity_type,required"`
	// The business's legal structure.
	LegalStructure ChildLegalEntityLegalStructure `json:"legal_structure,required,nullable"`
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
	OperatingJurisdictions []string                      `json:"operating_jurisdictions,required"`
	PhoneNumbers           []ChildLegalEntityPhoneNumber `json:"phone_numbers,required"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson bool `json:"politically_exposed_person,required,nullable"`
	// An individual's preferred name.
	PreferredName string `json:"preferred_name,required,nullable"`
	// An individual's prefix.
	Prefix string `json:"prefix,required,nullable"`
	// A list of primary social media URLs for the business.
	PrimarySocialMediaSites []string `json:"primary_social_media_sites,required"`
	// Array of regulatory bodies overseeing this institution.
	Regulators []ChildLegalEntityRegulator `json:"regulators,required,nullable"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating ChildLegalEntityRiskRating `json:"risk_rating,required,nullable"`
	// An individual's suffix.
	Suffix string `json:"suffix,required,nullable"`
	// Information describing a third-party verification run by an external vendor.
	ThirdPartyVerification ChildLegalEntityThirdPartyVerification `json:"third_party_verification,required,nullable"`
	// Stock ticker symbol for publicly traded companies.
	TickerSymbol               string                                   `json:"ticker_symbol,required,nullable"`
	UpdatedAt                  time.Time                                `json:"updated_at,required" format:"date-time"`
	WealthAndEmploymentDetails shared.LegalEntityWealthEmploymentDetail `json:"wealth_and_employment_details,required,nullable"`
	// The entity's primary website URL.
	Website string               `json:"website,required,nullable"`
	JSON    childLegalEntityJSON `json:"-"`
}

// childLegalEntityJSON contains the JSON metadata for the struct
// [ChildLegalEntity]
type childLegalEntityJSON struct {
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
	LegalEntityAssociations    apijson.Field
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
	Suffix                     apijson.Field
	ThirdPartyVerification     apijson.Field
	TickerSymbol               apijson.Field
	UpdatedAt                  apijson.Field
	WealthAndEmploymentDetails apijson.Field
	Website                    apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ChildLegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityJSON) RawJSON() string {
	return r.raw
}

type ChildLegalEntityAddress struct {
	ID string `json:"id,required" format:"uuid"`
	// The types of this address.
	AddressTypes []ChildLegalEntityAddressesAddressType `json:"address_types,required"`
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
	Region    string                      `json:"region,required,nullable"`
	UpdatedAt time.Time                   `json:"updated_at,required" format:"date-time"`
	JSON      childLegalEntityAddressJSON `json:"-"`
}

// childLegalEntityAddressJSON contains the JSON metadata for the struct
// [ChildLegalEntityAddress]
type childLegalEntityAddressJSON struct {
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

func (r *ChildLegalEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityAddressJSON) RawJSON() string {
	return r.raw
}

type ChildLegalEntityAddressesAddressType string

const (
	ChildLegalEntityAddressesAddressTypeBusiness    ChildLegalEntityAddressesAddressType = "business"
	ChildLegalEntityAddressesAddressTypeMailing     ChildLegalEntityAddressesAddressType = "mailing"
	ChildLegalEntityAddressesAddressTypeOther       ChildLegalEntityAddressesAddressType = "other"
	ChildLegalEntityAddressesAddressTypePoBox       ChildLegalEntityAddressesAddressType = "po_box"
	ChildLegalEntityAddressesAddressTypeResidential ChildLegalEntityAddressesAddressType = "residential"
)

func (r ChildLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case ChildLegalEntityAddressesAddressTypeBusiness, ChildLegalEntityAddressesAddressTypeMailing, ChildLegalEntityAddressesAddressTypeOther, ChildLegalEntityAddressesAddressTypePoBox, ChildLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type ChildLegalEntityIdentification struct {
	ID          string     `json:"id,required" format:"uuid"`
	CreatedAt   time.Time  `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time  `json:"discarded_at,required,nullable" format:"date-time"`
	Documents   []Document `json:"documents,required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate time.Time `json:"expiration_date,required,nullable" format:"date"`
	// The type of ID number.
	IDType ChildLegalEntityIdentificationsIDType `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country,required,nullable"`
	// The region in which the identifcation was issued.
	IssuingRegion string `json:"issuing_region,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                               `json:"live_mode,required"`
	Object    string                             `json:"object,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	JSON      childLegalEntityIdentificationJSON `json:"-"`
}

// childLegalEntityIdentificationJSON contains the JSON metadata for the struct
// [ChildLegalEntityIdentification]
type childLegalEntityIdentificationJSON struct {
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

func (r *ChildLegalEntityIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityIdentificationJSON) RawJSON() string {
	return r.raw
}

// The type of ID number.
type ChildLegalEntityIdentificationsIDType string

const (
	ChildLegalEntityIdentificationsIDTypeArCuil         ChildLegalEntityIdentificationsIDType = "ar_cuil"
	ChildLegalEntityIdentificationsIDTypeArCuit         ChildLegalEntityIdentificationsIDType = "ar_cuit"
	ChildLegalEntityIdentificationsIDTypeBrCnpj         ChildLegalEntityIdentificationsIDType = "br_cnpj"
	ChildLegalEntityIdentificationsIDTypeBrCpf          ChildLegalEntityIdentificationsIDType = "br_cpf"
	ChildLegalEntityIdentificationsIDTypeClRun          ChildLegalEntityIdentificationsIDType = "cl_run"
	ChildLegalEntityIdentificationsIDTypeClRut          ChildLegalEntityIdentificationsIDType = "cl_rut"
	ChildLegalEntityIdentificationsIDTypeCoCedulas      ChildLegalEntityIdentificationsIDType = "co_cedulas"
	ChildLegalEntityIdentificationsIDTypeCoNit          ChildLegalEntityIdentificationsIDType = "co_nit"
	ChildLegalEntityIdentificationsIDTypeDriversLicense ChildLegalEntityIdentificationsIDType = "drivers_license"
	ChildLegalEntityIdentificationsIDTypeHnID           ChildLegalEntityIdentificationsIDType = "hn_id"
	ChildLegalEntityIdentificationsIDTypeHnRtn          ChildLegalEntityIdentificationsIDType = "hn_rtn"
	ChildLegalEntityIdentificationsIDTypeInLei          ChildLegalEntityIdentificationsIDType = "in_lei"
	ChildLegalEntityIdentificationsIDTypeKrBrn          ChildLegalEntityIdentificationsIDType = "kr_brn"
	ChildLegalEntityIdentificationsIDTypeKrCrn          ChildLegalEntityIdentificationsIDType = "kr_crn"
	ChildLegalEntityIdentificationsIDTypeKrRrn          ChildLegalEntityIdentificationsIDType = "kr_rrn"
	ChildLegalEntityIdentificationsIDTypePassport       ChildLegalEntityIdentificationsIDType = "passport"
	ChildLegalEntityIdentificationsIDTypeSaTin          ChildLegalEntityIdentificationsIDType = "sa_tin"
	ChildLegalEntityIdentificationsIDTypeSaVat          ChildLegalEntityIdentificationsIDType = "sa_vat"
	ChildLegalEntityIdentificationsIDTypeUsEin          ChildLegalEntityIdentificationsIDType = "us_ein"
	ChildLegalEntityIdentificationsIDTypeUsItin         ChildLegalEntityIdentificationsIDType = "us_itin"
	ChildLegalEntityIdentificationsIDTypeUsSsn          ChildLegalEntityIdentificationsIDType = "us_ssn"
	ChildLegalEntityIdentificationsIDTypeVnTin          ChildLegalEntityIdentificationsIDType = "vn_tin"
)

func (r ChildLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case ChildLegalEntityIdentificationsIDTypeArCuil, ChildLegalEntityIdentificationsIDTypeArCuit, ChildLegalEntityIdentificationsIDTypeBrCnpj, ChildLegalEntityIdentificationsIDTypeBrCpf, ChildLegalEntityIdentificationsIDTypeClRun, ChildLegalEntityIdentificationsIDTypeClRut, ChildLegalEntityIdentificationsIDTypeCoCedulas, ChildLegalEntityIdentificationsIDTypeCoNit, ChildLegalEntityIdentificationsIDTypeDriversLicense, ChildLegalEntityIdentificationsIDTypeHnID, ChildLegalEntityIdentificationsIDTypeHnRtn, ChildLegalEntityIdentificationsIDTypeInLei, ChildLegalEntityIdentificationsIDTypeKrBrn, ChildLegalEntityIdentificationsIDTypeKrCrn, ChildLegalEntityIdentificationsIDTypeKrRrn, ChildLegalEntityIdentificationsIDTypePassport, ChildLegalEntityIdentificationsIDTypeSaTin, ChildLegalEntityIdentificationsIDTypeSaVat, ChildLegalEntityIdentificationsIDTypeUsEin, ChildLegalEntityIdentificationsIDTypeUsItin, ChildLegalEntityIdentificationsIDTypeUsSsn, ChildLegalEntityIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

// The type of legal entity.
type ChildLegalEntityLegalEntityType string

const (
	ChildLegalEntityLegalEntityTypeBusiness   ChildLegalEntityLegalEntityType = "business"
	ChildLegalEntityLegalEntityTypeIndividual ChildLegalEntityLegalEntityType = "individual"
	ChildLegalEntityLegalEntityTypeJoint      ChildLegalEntityLegalEntityType = "joint"
)

func (r ChildLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case ChildLegalEntityLegalEntityTypeBusiness, ChildLegalEntityLegalEntityTypeIndividual, ChildLegalEntityLegalEntityTypeJoint:
		return true
	}
	return false
}

// The business's legal structure.
type ChildLegalEntityLegalStructure string

const (
	ChildLegalEntityLegalStructureCorporation        ChildLegalEntityLegalStructure = "corporation"
	ChildLegalEntityLegalStructureLlc                ChildLegalEntityLegalStructure = "llc"
	ChildLegalEntityLegalStructureNonProfit          ChildLegalEntityLegalStructure = "non_profit"
	ChildLegalEntityLegalStructurePartnership        ChildLegalEntityLegalStructure = "partnership"
	ChildLegalEntityLegalStructureSoleProprietorship ChildLegalEntityLegalStructure = "sole_proprietorship"
	ChildLegalEntityLegalStructureTrust              ChildLegalEntityLegalStructure = "trust"
)

func (r ChildLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case ChildLegalEntityLegalStructureCorporation, ChildLegalEntityLegalStructureLlc, ChildLegalEntityLegalStructureNonProfit, ChildLegalEntityLegalStructurePartnership, ChildLegalEntityLegalStructureSoleProprietorship, ChildLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type ChildLegalEntityPhoneNumber struct {
	PhoneNumber string                          `json:"phone_number"`
	JSON        childLegalEntityPhoneNumberJSON `json:"-"`
}

// childLegalEntityPhoneNumberJSON contains the JSON metadata for the struct
// [ChildLegalEntityPhoneNumber]
type childLegalEntityPhoneNumberJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChildLegalEntityPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityPhoneNumberJSON) RawJSON() string {
	return r.raw
}

type ChildLegalEntityRegulator struct {
	// The country code where the regulator operates in the ISO 3166-1 alpha-2 format
	// (e.g., "US", "CA", "GB").
	Jurisdiction string `json:"jurisdiction,required"`
	// Full name of the regulatory body.
	Name string `json:"name,required"`
	// Registration or identification number with the regulator.
	RegistrationNumber string                        `json:"registration_number,required"`
	JSON               childLegalEntityRegulatorJSON `json:"-"`
}

// childLegalEntityRegulatorJSON contains the JSON metadata for the struct
// [ChildLegalEntityRegulator]
type childLegalEntityRegulatorJSON struct {
	Jurisdiction       apijson.Field
	Name               apijson.Field
	RegistrationNumber apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ChildLegalEntityRegulator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityRegulatorJSON) RawJSON() string {
	return r.raw
}

// The risk rating of the legal entity. One of low, medium, high.
type ChildLegalEntityRiskRating string

const (
	ChildLegalEntityRiskRatingLow    ChildLegalEntityRiskRating = "low"
	ChildLegalEntityRiskRatingMedium ChildLegalEntityRiskRating = "medium"
	ChildLegalEntityRiskRatingHigh   ChildLegalEntityRiskRating = "high"
)

func (r ChildLegalEntityRiskRating) IsKnown() bool {
	switch r {
	case ChildLegalEntityRiskRatingLow, ChildLegalEntityRiskRatingMedium, ChildLegalEntityRiskRatingHigh:
		return true
	}
	return false
}

// Information describing a third-party verification run by an external vendor.
type ChildLegalEntityThirdPartyVerification struct {
	// The vendor that performed the verification, e.g. `persona`.
	Vendor ChildLegalEntityThirdPartyVerificationVendor `json:"vendor,required"`
	// The identification of the third party verification in `vendor`'s system.
	VendorVerificationID string                                     `json:"vendor_verification_id,required"`
	JSON                 childLegalEntityThirdPartyVerificationJSON `json:"-"`
}

// childLegalEntityThirdPartyVerificationJSON contains the JSON metadata for the
// struct [ChildLegalEntityThirdPartyVerification]
type childLegalEntityThirdPartyVerificationJSON struct {
	Vendor               apijson.Field
	VendorVerificationID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ChildLegalEntityThirdPartyVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityThirdPartyVerificationJSON) RawJSON() string {
	return r.raw
}

// The vendor that performed the verification, e.g. `persona`.
type ChildLegalEntityThirdPartyVerificationVendor string

const (
	ChildLegalEntityThirdPartyVerificationVendorPersona ChildLegalEntityThirdPartyVerificationVendor = "persona"
)

func (r ChildLegalEntityThirdPartyVerificationVendor) IsKnown() bool {
	switch r {
	case ChildLegalEntityThirdPartyVerificationVendorPersona:
		return true
	}
	return false
}

type LegalEntityAssociation struct {
	ID string `json:"id,required" format:"uuid"`
	// The child legal entity.
	ChildLegalEntity *ChildLegalEntity `json:"child_legal_entity,required"`
	CreatedAt        time.Time         `json:"created_at,required" format:"date-time"`
	DiscardedAt      time.Time         `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage int64 `json:"ownership_percentage,required,nullable"`
	// The ID of the parent legal entity. This must be a business or joint legal
	// entity.
	ParentLegalEntityID string                                   `json:"parent_legal_entity_id,required"`
	RelationshipTypes   []LegalEntityAssociationRelationshipType `json:"relationship_types,required"`
	// The job title of the child entity at the parent entity.
	Title     string                     `json:"title,required,nullable"`
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityAssociationJSON `json:"-"`
}

// legalEntityAssociationJSON contains the JSON metadata for the struct
// [LegalEntityAssociation]
type legalEntityAssociationJSON struct {
	ID                  apijson.Field
	ChildLegalEntity    apijson.Field
	CreatedAt           apijson.Field
	DiscardedAt         apijson.Field
	LiveMode            apijson.Field
	Object              apijson.Field
	OwnershipPercentage apijson.Field
	ParentLegalEntityID apijson.Field
	RelationshipTypes   apijson.Field
	Title               apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LegalEntityAssociation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityAssociationJSON) RawJSON() string {
	return r.raw
}

// A list of relationship types for how the child entity relates to parent entity.
type LegalEntityAssociationRelationshipType string

const (
	LegalEntityAssociationRelationshipTypeAuthorizedSigner LegalEntityAssociationRelationshipType = "authorized_signer"
	LegalEntityAssociationRelationshipTypeBeneficialOwner  LegalEntityAssociationRelationshipType = "beneficial_owner"
	LegalEntityAssociationRelationshipTypeControlPerson    LegalEntityAssociationRelationshipType = "control_person"
)

func (r LegalEntityAssociationRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationRelationshipTypeAuthorizedSigner, LegalEntityAssociationRelationshipTypeBeneficialOwner, LegalEntityAssociationRelationshipTypeControlPerson:
		return true
	}
	return false
}

type LegalEntityAssociationNewParams struct {
	// The ID of the parent legal entity. This must be a business or joint legal
	// entity.
	ParentLegalEntityID param.Field[string]                                            `json:"parent_legal_entity_id,required"`
	RelationshipTypes   param.Field[[]LegalEntityAssociationNewParamsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[shared.ChildLegalEntityCreateParam] `json:"child_legal_entity"`
	// The ID of the child legal entity.
	ChildLegalEntityID param.Field[string] `json:"child_legal_entity_id"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the child entity at the parent entity.
	Title param.Field[string] `json:"title"`
}

func (r LegalEntityAssociationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the child entity relates to parent entity.
type LegalEntityAssociationNewParamsRelationshipType string

const (
	LegalEntityAssociationNewParamsRelationshipTypeAuthorizedSigner LegalEntityAssociationNewParamsRelationshipType = "authorized_signer"
	LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner  LegalEntityAssociationNewParamsRelationshipType = "beneficial_owner"
	LegalEntityAssociationNewParamsRelationshipTypeControlPerson    LegalEntityAssociationNewParamsRelationshipType = "control_person"
)

func (r LegalEntityAssociationNewParamsRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsRelationshipTypeAuthorizedSigner, LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner, LegalEntityAssociationNewParamsRelationshipTypeControlPerson:
		return true
	}
	return false
}
