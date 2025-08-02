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
	opts = append(r.Options[:], opts...)
	path := "api/legal_entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single legal entity.
func (r *LegalEntityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	Addresses    []LegalEntityAddress `json:"addresses,required"`
	BankSettings LegalEntity          `json:"bank_settings,required,nullable"`
	// The business's legal business name.
	BusinessName string `json:"business_name,required,nullable"`
	// The country of citizenship for an individual.
	CitizenshipCountry string                             `json:"citizenship_country,required,nullable"`
	ComplianceDetails  shared.LegalEntityComplianceDetail `json:"compliance_details,required,nullable"`
	CreatedAt          time.Time                          `json:"created_at,required" format:"date-time"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed time.Time `json:"date_formed,required,nullable" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          time.Time `json:"date_of_birth,required,nullable" format:"date"`
	DiscardedAt          time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	DoingBusinessAsNames []string  `json:"doing_business_as_names,required"`
	// The entity's primary email.
	Email string `json:"email,required,nullable"`
	// An individual's first name.
	FirstName string `json:"first_name,required,nullable"`
	// A list of identifications for the legal entity.
	Identifications []LegalEntityIdentification `json:"identifications,required"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications []shared.LegalEntityIndustryClassification `json:"industry_classifications,required"`
	// An individual's last name.
	LastName string `json:"last_name,required,nullable"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations []LegalEntityAssociation `json:"legal_entity_associations,required,nullable"`
	// The type of legal entity.
	LegalEntityType LegalEntityLegalEntityType `json:"legal_entity_type,required"`
	// The business's legal structure.
	LegalStructure LegalEntityLegalStructure `json:"legal_structure,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// An individual's middle name.
	MiddleName   string                   `json:"middle_name,required,nullable"`
	Object       string                   `json:"object,required"`
	PhoneNumbers []LegalEntityPhoneNumber `json:"phone_numbers,required"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson bool `json:"politically_exposed_person,required,nullable"`
	// An individual's preferred name.
	PreferredName string `json:"preferred_name,required,nullable"`
	// An individual's prefix.
	Prefix string `json:"prefix,required,nullable"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating LegalEntityRiskRating `json:"risk_rating,required,nullable"`
	// An individual's suffix.
	Suffix                     string      `json:"suffix,required,nullable"`
	UpdatedAt                  time.Time   `json:"updated_at,required" format:"date-time"`
	WealthAndEmploymentDetails LegalEntity `json:"wealth_and_employment_details,required,nullable"`
	// The entity's primary website URL.
	Website string          `json:"website,required,nullable"`
	JSON    legalEntityJSON `json:"-"`
}

// legalEntityJSON contains the JSON metadata for the struct [LegalEntity]
type legalEntityJSON struct {
	ID                         apijson.Field
	Addresses                  apijson.Field
	BankSettings               apijson.Field
	BusinessName               apijson.Field
	CitizenshipCountry         apijson.Field
	ComplianceDetails          apijson.Field
	CreatedAt                  apijson.Field
	DateFormed                 apijson.Field
	DateOfBirth                apijson.Field
	DiscardedAt                apijson.Field
	DoingBusinessAsNames       apijson.Field
	Email                      apijson.Field
	FirstName                  apijson.Field
	Identifications            apijson.Field
	IndustryClassifications    apijson.Field
	LastName                   apijson.Field
	LegalEntityAssociations    apijson.Field
	LegalEntityType            apijson.Field
	LegalStructure             apijson.Field
	LiveMode                   apijson.Field
	Metadata                   apijson.Field
	MiddleName                 apijson.Field
	Object                     apijson.Field
	PhoneNumbers               apijson.Field
	PoliticallyExposedPerson   apijson.Field
	PreferredName              apijson.Field
	Prefix                     apijson.Field
	RiskRating                 apijson.Field
	Suffix                     apijson.Field
	UpdatedAt                  apijson.Field
	WealthAndEmploymentDetails apijson.Field
	Website                    apijson.Field
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
	ID          string    `json:"id,required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
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

type LegalEntityNewParams struct {
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityType] `json:"legal_entity_type,required"`
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[LegalEntityNewParamsBankSettings]              `json:"bank_settings"`
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
	LegalEntityAssociations param.Field[[]LegalEntityNewParamsLegalEntityAssociation] `json:"legal_entity_associations"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityNewParamsLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName   param.Field[string]                            `json:"middle_name"`
	PhoneNumbers param.Field[[]LegalEntityNewParamsPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[LegalEntityNewParamsRiskRating] `json:"risk_rating"`
	// An individual's suffix.
	Suffix                     param.Field[string]                                         `json:"suffix"`
	WealthAndEmploymentDetails param.Field[LegalEntityNewParamsWealthAndEmploymentDetails] `json:"wealth_and_employment_details"`
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

type LegalEntityNewParamsBankSettings struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// The percentage of backup withholding to apply to the legal entity.
	BackupWithholdingPercentage param.Field[int64]     `json:"backup_withholding_percentage,required"`
	CreatedAt                   param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	DiscardedAt                 param.Field[time.Time] `json:"discarded_at,required" format:"date-time"`
	// Whether backup withholding is enabled. See more here -
	// https://www.irs.gov/businesses/small-businesses-self-employed/backup-withholding.
	EnableBackupWithholding param.Field[bool] `json:"enable_backup_withholding,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// Cross River Bank specific setting to opt out of privacy policy.
	PrivacyOptOut param.Field[bool] `json:"privacy_opt_out,required"`
	// It covers, among other types of insider loans, extensions of credit by a member
	// bank to an executive officer, director, or principal shareholder of the member
	// bank; a bank holding company of which the member bank is a subsidiary; and any
	// other subsidiary of that bank holding company.
	RegulationO param.Field[bool]      `json:"regulation_o,required"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
}

func (r LegalEntityNewParamsBankSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]LegalEntityNewParamsLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[shared.ChildLegalEntityCreateParam] `json:"child_legal_entity"`
	// The ID of the child legal entity.
	ChildLegalEntityID param.Field[string] `json:"child_legal_entity_id"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the child entity at the parent entity.
	Title param.Field[string] `json:"title"`
}

func (r LegalEntityNewParamsLegalEntityAssociation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the child entity relates to parent entity.
type LegalEntityNewParamsLegalEntityAssociationsRelationshipType string

const (
	LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeAuthorizedSigner LegalEntityNewParamsLegalEntityAssociationsRelationshipType = "authorized_signer"
	LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner  LegalEntityNewParamsLegalEntityAssociationsRelationshipType = "beneficial_owner"
	LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson    LegalEntityNewParamsLegalEntityAssociationsRelationshipType = "control_person"
)

func (r LegalEntityNewParamsLegalEntityAssociationsRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeAuthorizedSigner, LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson:
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

type LegalEntityNewParamsWealthAndEmploymentDetails struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// The annual income of the individual.
	AnnualIncome param.Field[int64]     `json:"annual_income,required"`
	CreatedAt    param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	DiscardedAt  param.Field[time.Time] `json:"discarded_at,required" format:"date-time"`
	// The country in which the employer is located.
	EmployerCountry param.Field[string] `json:"employer_country,required"`
	// The name of the employer.
	EmployerName param.Field[string] `json:"employer_name,required"`
	// The state in which the employer is located.
	EmployerState param.Field[string] `json:"employer_state,required"`
	// The employment status of the individual.
	EmploymentStatus param.Field[LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus] `json:"employment_status,required"`
	// The country in which the individual's income is earned.
	IncomeCountry param.Field[string] `json:"income_country,required"`
	// The source of the individual's income.
	IncomeSource param.Field[LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource] `json:"income_source,required"`
	// The state in which the individual's income is earned.
	IncomeState param.Field[string] `json:"income_state,required"`
	// The industry of the individual.
	Industry param.Field[LegalEntityNewParamsWealthAndEmploymentDetailsIndustry] `json:"industry,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// The occupation of the individual.
	Occupation param.Field[LegalEntityNewParamsWealthAndEmploymentDetailsOccupation] `json:"occupation,required"`
	// The source of the individual's funds.
	SourceOfFunds param.Field[LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds] `json:"source_of_funds,required"`
	UpdatedAt     param.Field[time.Time]                                                   `json:"updated_at,required" format:"date-time"`
	// The source of the individual's wealth.
	WealthSource param.Field[LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource] `json:"wealth_source,required"`
}

func (r LegalEntityNewParamsWealthAndEmploymentDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The employment status of the individual.
type LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus string

const (
	LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusEmployed     LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus = "employed"
	LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusRetired      LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus = "retired"
	LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusSelfEmployed LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus = "self_employed"
	LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusStudent      LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus = "student"
	LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusUnemployed   LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus = "unemployed"
)

func (r LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatus) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusEmployed, LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusRetired, LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusSelfEmployed, LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusStudent, LegalEntityNewParamsWealthAndEmploymentDetailsEmploymentStatusUnemployed:
		return true
	}
	return false
}

// The source of the individual's income.
type LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource string

const (
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceFamilySupport      LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "family_support"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceGovernmentBenefits LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "government_benefits"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceInheritance        LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "inheritance"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceInvestments        LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "investments"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceRentalIncome       LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "rental_income"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceRetirement         LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "retirement"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceSalary             LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "salary"
	LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceSelfEmployed       LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource = "self_employed"
)

func (r LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSource) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceFamilySupport, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceGovernmentBenefits, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceInheritance, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceInvestments, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceRentalIncome, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceRetirement, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceSalary, LegalEntityNewParamsWealthAndEmploymentDetailsIncomeSourceSelfEmployed:
		return true
	}
	return false
}

// The industry of the individual.
type LegalEntityNewParamsWealthAndEmploymentDetailsIndustry string

const (
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryAccounting            LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "accounting"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryAgriculture           LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "agriculture"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryAutomotive            LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "automotive"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryChemicalManufacturing LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "chemical_manufacturing"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryConstruction          LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "construction"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryEducationalMedical    LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "educational_medical"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryFoodService           LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "food_service"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryFinance               LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "finance"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryGasoline              LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "gasoline"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryHealthStores          LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "health_stores"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryLaundry               LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "laundry"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryMaintenance           LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "maintenance"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryManufacturing         LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "manufacturing"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryMerchantWholesale     LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "merchant_wholesale"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryMining                LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "mining"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryPerformingArts        LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "performing_arts"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryProfessionalNonLegal  LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "professional_non_legal"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryPublicAdministration  LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "public_administration"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryPublishing            LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "publishing"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRealEstate            LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "real_estate"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRecreationGambling    LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "recreation_gambling"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryReligiousCharity      LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "religious_charity"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRentalServices        LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "rental_services"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailClothing        LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_clothing"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailElectronics     LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_electronics"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailFood            LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_food"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailFurnishing      LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_furnishing"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailHome            LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_home"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailNonStore        LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_non_store"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailSporting        LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "retail_sporting"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryTransportation        LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "transportation"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryTravel                LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "travel"
	LegalEntityNewParamsWealthAndEmploymentDetailsIndustryUtilities             LegalEntityNewParamsWealthAndEmploymentDetailsIndustry = "utilities"
)

func (r LegalEntityNewParamsWealthAndEmploymentDetailsIndustry) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsWealthAndEmploymentDetailsIndustryAccounting, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryAgriculture, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryAutomotive, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryChemicalManufacturing, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryConstruction, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryEducationalMedical, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryFoodService, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryFinance, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryGasoline, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryHealthStores, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryLaundry, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryMaintenance, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryManufacturing, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryMerchantWholesale, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryMining, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryPerformingArts, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryProfessionalNonLegal, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryPublicAdministration, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryPublishing, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRealEstate, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRecreationGambling, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryReligiousCharity, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRentalServices, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailClothing, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailElectronics, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailFood, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailFurnishing, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailHome, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailNonStore, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryRetailSporting, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryTransportation, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryTravel, LegalEntityNewParamsWealthAndEmploymentDetailsIndustryUtilities:
		return true
	}
	return false
}

// The occupation of the individual.
type LegalEntityNewParamsWealthAndEmploymentDetailsOccupation string

const (
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationConsulting         LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "consulting"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationExecutive          LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "executive"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationFinanceAccounting  LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "finance_accounting"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationFoodServices       LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "food_services"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationGovernment         LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "government"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationHealthcare         LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "healthcare"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationLegalServices      LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "legal_services"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationManufacturing      LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "manufacturing"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationOther              LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "other"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationSales              LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "sales"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationScienceEngineering LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "science_engineering"
	LegalEntityNewParamsWealthAndEmploymentDetailsOccupationTechnology         LegalEntityNewParamsWealthAndEmploymentDetailsOccupation = "technology"
)

func (r LegalEntityNewParamsWealthAndEmploymentDetailsOccupation) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsWealthAndEmploymentDetailsOccupationConsulting, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationExecutive, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationFinanceAccounting, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationFoodServices, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationGovernment, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationHealthcare, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationLegalServices, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationManufacturing, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationOther, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationSales, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationScienceEngineering, LegalEntityNewParamsWealthAndEmploymentDetailsOccupationTechnology:
		return true
	}
	return false
}

// The source of the individual's funds.
type LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds string

const (
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsAlimony            LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "alimony"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsAnnuity            LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "annuity"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsBusinessOwner      LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "business_owner"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsGeneralEmployee    LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "general_employee"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsGovernmentBenefits LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "government_benefits"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsHomemaker          LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "homemaker"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsInheritanceGift    LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "inheritance_gift"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsInvestment         LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "investment"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsLegalSettlement    LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "legal_settlement"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsLottery            LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "lottery"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsRealEstate         LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "real_estate"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsRetired            LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "retired"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsRetirement         LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "retirement"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsSalary             LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "salary"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsSelfEmployed       LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "self_employed"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsSeniorExecutive    LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "senior_executive"
	LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsTrustIncome        LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds = "trust_income"
)

func (r LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFunds) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsAlimony, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsAnnuity, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsBusinessOwner, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsGeneralEmployee, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsGovernmentBenefits, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsHomemaker, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsInheritanceGift, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsInvestment, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsLegalSettlement, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsLottery, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsRealEstate, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsRetired, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsRetirement, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsSalary, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsSelfEmployed, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsSeniorExecutive, LegalEntityNewParamsWealthAndEmploymentDetailsSourceOfFundsTrustIncome:
		return true
	}
	return false
}

// The source of the individual's wealth.
type LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource string

const (
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceBusinessSale       LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "business_sale"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceFamilySupport      LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "family_support"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceGovernmentBenefits LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "government_benefits"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceInheritance        LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "inheritance"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceInvestments        LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "investments"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceOther              LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "other"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceRentalIncome       LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "rental_income"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceRetirement         LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "retirement"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceSalary             LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "salary"
	LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceSelfEmployed       LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource = "self_employed"
)

func (r LegalEntityNewParamsWealthAndEmploymentDetailsWealthSource) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceBusinessSale, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceFamilySupport, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceGovernmentBenefits, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceInheritance, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceInvestments, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceOther, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceRentalIncome, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceRetirement, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceSalary, LegalEntityNewParamsWealthAndEmploymentDetailsWealthSourceSelfEmployed:
		return true
	}
	return false
}

type LegalEntityUpdateParams struct {
	// A list of addresses for the entity.
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[LegalEntityUpdateParamsBankSettings]           `json:"bank_settings"`
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
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityUpdateParamsLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName   param.Field[string]                               `json:"middle_name"`
	PhoneNumbers param.Field[[]LegalEntityUpdateParamsPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[LegalEntityUpdateParamsRiskRating] `json:"risk_rating"`
	// An individual's suffix.
	Suffix                     param.Field[string]                                            `json:"suffix"`
	WealthAndEmploymentDetails param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetails] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityUpdateParamsBankSettings struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// The percentage of backup withholding to apply to the legal entity.
	BackupWithholdingPercentage param.Field[int64]     `json:"backup_withholding_percentage,required"`
	CreatedAt                   param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	DiscardedAt                 param.Field[time.Time] `json:"discarded_at,required" format:"date-time"`
	// Whether backup withholding is enabled. See more here -
	// https://www.irs.gov/businesses/small-businesses-self-employed/backup-withholding.
	EnableBackupWithholding param.Field[bool] `json:"enable_backup_withholding,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// Cross River Bank specific setting to opt out of privacy policy.
	PrivacyOptOut param.Field[bool] `json:"privacy_opt_out,required"`
	// It covers, among other types of insider loans, extensions of credit by a member
	// bank to an executive officer, director, or principal shareholder of the member
	// bank; a bank holding company of which the member bank is a subsidiary; and any
	// other subsidiary of that bank holding company.
	RegulationO param.Field[bool]      `json:"regulation_o,required"`
	UpdatedAt   param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
}

func (r LegalEntityUpdateParamsBankSettings) MarshalJSON() (data []byte, err error) {
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

type LegalEntityUpdateParamsWealthAndEmploymentDetails struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// The annual income of the individual.
	AnnualIncome param.Field[int64]     `json:"annual_income,required"`
	CreatedAt    param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	DiscardedAt  param.Field[time.Time] `json:"discarded_at,required" format:"date-time"`
	// The country in which the employer is located.
	EmployerCountry param.Field[string] `json:"employer_country,required"`
	// The name of the employer.
	EmployerName param.Field[string] `json:"employer_name,required"`
	// The state in which the employer is located.
	EmployerState param.Field[string] `json:"employer_state,required"`
	// The employment status of the individual.
	EmploymentStatus param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus] `json:"employment_status,required"`
	// The country in which the individual's income is earned.
	IncomeCountry param.Field[string] `json:"income_country,required"`
	// The source of the individual's income.
	IncomeSource param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource] `json:"income_source,required"`
	// The state in which the individual's income is earned.
	IncomeState param.Field[string] `json:"income_state,required"`
	// The industry of the individual.
	Industry param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry] `json:"industry,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// The occupation of the individual.
	Occupation param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation] `json:"occupation,required"`
	// The source of the individual's funds.
	SourceOfFunds param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds] `json:"source_of_funds,required"`
	UpdatedAt     param.Field[time.Time]                                                      `json:"updated_at,required" format:"date-time"`
	// The source of the individual's wealth.
	WealthSource param.Field[LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource] `json:"wealth_source,required"`
}

func (r LegalEntityUpdateParamsWealthAndEmploymentDetails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The employment status of the individual.
type LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus string

const (
	LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusEmployed     LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus = "employed"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusRetired      LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus = "retired"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusSelfEmployed LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus = "self_employed"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusStudent      LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus = "student"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusUnemployed   LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus = "unemployed"
)

func (r LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatus) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusEmployed, LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusRetired, LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusSelfEmployed, LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusStudent, LegalEntityUpdateParamsWealthAndEmploymentDetailsEmploymentStatusUnemployed:
		return true
	}
	return false
}

// The source of the individual's income.
type LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource string

const (
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceFamilySupport      LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "family_support"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceGovernmentBenefits LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "government_benefits"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceInheritance        LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "inheritance"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceInvestments        LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "investments"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceRentalIncome       LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "rental_income"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceRetirement         LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "retirement"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceSalary             LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "salary"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceSelfEmployed       LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource = "self_employed"
)

func (r LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSource) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceFamilySupport, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceGovernmentBenefits, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceInheritance, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceInvestments, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceRentalIncome, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceRetirement, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceSalary, LegalEntityUpdateParamsWealthAndEmploymentDetailsIncomeSourceSelfEmployed:
		return true
	}
	return false
}

// The industry of the individual.
type LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry string

const (
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryAccounting            LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "accounting"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryAgriculture           LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "agriculture"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryAutomotive            LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "automotive"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryChemicalManufacturing LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "chemical_manufacturing"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryConstruction          LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "construction"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryEducationalMedical    LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "educational_medical"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryFoodService           LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "food_service"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryFinance               LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "finance"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryGasoline              LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "gasoline"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryHealthStores          LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "health_stores"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryLaundry               LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "laundry"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryMaintenance           LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "maintenance"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryManufacturing         LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "manufacturing"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryMerchantWholesale     LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "merchant_wholesale"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryMining                LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "mining"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryPerformingArts        LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "performing_arts"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryProfessionalNonLegal  LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "professional_non_legal"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryPublicAdministration  LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "public_administration"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryPublishing            LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "publishing"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRealEstate            LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "real_estate"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRecreationGambling    LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "recreation_gambling"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryReligiousCharity      LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "religious_charity"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRentalServices        LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "rental_services"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailClothing        LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_clothing"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailElectronics     LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_electronics"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailFood            LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_food"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailFurnishing      LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_furnishing"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailHome            LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_home"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailNonStore        LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_non_store"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailSporting        LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "retail_sporting"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryTransportation        LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "transportation"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryTravel                LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "travel"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryUtilities             LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry = "utilities"
)

func (r LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustry) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryAccounting, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryAgriculture, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryAutomotive, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryChemicalManufacturing, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryConstruction, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryEducationalMedical, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryFoodService, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryFinance, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryGasoline, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryHealthStores, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryLaundry, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryMaintenance, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryManufacturing, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryMerchantWholesale, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryMining, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryPerformingArts, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryProfessionalNonLegal, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryPublicAdministration, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryPublishing, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRealEstate, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRecreationGambling, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryReligiousCharity, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRentalServices, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailClothing, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailElectronics, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailFood, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailFurnishing, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailHome, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailNonStore, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryRetailSporting, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryTransportation, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryTravel, LegalEntityUpdateParamsWealthAndEmploymentDetailsIndustryUtilities:
		return true
	}
	return false
}

// The occupation of the individual.
type LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation string

const (
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationConsulting         LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "consulting"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationExecutive          LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "executive"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationFinanceAccounting  LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "finance_accounting"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationFoodServices       LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "food_services"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationGovernment         LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "government"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationHealthcare         LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "healthcare"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationLegalServices      LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "legal_services"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationManufacturing      LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "manufacturing"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationOther              LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "other"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationSales              LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "sales"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationScienceEngineering LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "science_engineering"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationTechnology         LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation = "technology"
)

func (r LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupation) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationConsulting, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationExecutive, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationFinanceAccounting, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationFoodServices, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationGovernment, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationHealthcare, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationLegalServices, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationManufacturing, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationOther, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationSales, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationScienceEngineering, LegalEntityUpdateParamsWealthAndEmploymentDetailsOccupationTechnology:
		return true
	}
	return false
}

// The source of the individual's funds.
type LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds string

const (
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsAlimony            LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "alimony"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsAnnuity            LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "annuity"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsBusinessOwner      LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "business_owner"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsGeneralEmployee    LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "general_employee"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsGovernmentBenefits LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "government_benefits"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsHomemaker          LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "homemaker"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsInheritanceGift    LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "inheritance_gift"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsInvestment         LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "investment"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsLegalSettlement    LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "legal_settlement"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsLottery            LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "lottery"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsRealEstate         LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "real_estate"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsRetired            LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "retired"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsRetirement         LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "retirement"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsSalary             LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "salary"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsSelfEmployed       LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "self_employed"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsSeniorExecutive    LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "senior_executive"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsTrustIncome        LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds = "trust_income"
)

func (r LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFunds) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsAlimony, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsAnnuity, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsBusinessOwner, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsGeneralEmployee, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsGovernmentBenefits, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsHomemaker, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsInheritanceGift, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsInvestment, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsLegalSettlement, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsLottery, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsRealEstate, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsRetired, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsRetirement, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsSalary, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsSelfEmployed, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsSeniorExecutive, LegalEntityUpdateParamsWealthAndEmploymentDetailsSourceOfFundsTrustIncome:
		return true
	}
	return false
}

// The source of the individual's wealth.
type LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource string

const (
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceBusinessSale       LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "business_sale"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceFamilySupport      LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "family_support"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceGovernmentBenefits LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "government_benefits"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceInheritance        LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "inheritance"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceInvestments        LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "investments"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceOther              LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "other"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceRentalIncome       LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "rental_income"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceRetirement         LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "retirement"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceSalary             LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "salary"
	LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceSelfEmployed       LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource = "self_employed"
)

func (r LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSource) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceBusinessSale, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceFamilySupport, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceGovernmentBenefits, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceInheritance, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceInvestments, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceOther, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceRentalIncome, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceRetirement, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceSalary, LegalEntityUpdateParamsWealthAndEmploymentDetailsWealthSourceSelfEmployed:
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
	Metadata    param.Field[map[string]string] `query:"metadata"`
	PerPage     param.Field[int64]             `query:"per_page"`
	ShowDeleted param.Field[string]            `query:"show_deleted"`
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
