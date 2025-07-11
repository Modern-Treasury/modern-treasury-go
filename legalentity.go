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

type BankSettings struct {
	ID string `json:"id,required" format:"uuid"`
	// The percentage of backup withholding to apply to the legal entity.
	BackupWithholdingPercentage int64     `json:"backup_withholding_percentage,required,nullable"`
	CreatedAt                   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt                 time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Whether backup withholding is enabled. See more here -
	// https://www.irs.gov/businesses/small-businesses-self-employed/backup-withholding.
	EnableBackupWithholding bool `json:"enable_backup_withholding,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// Cross River Bank specific setting to opt out of privacy policy.
	PrivacyOptOut bool `json:"privacy_opt_out,required,nullable"`
	// It covers, among other types of insider loans, extensions of credit by a member
	// bank to an executive officer, director, or principal shareholder of the member
	// bank; a bank holding company of which the member bank is a subsidiary; and any
	// other subsidiary of that bank holding company.
	RegulationO bool             `json:"regulation_o,required,nullable"`
	UpdatedAt   time.Time        `json:"updated_at,required" format:"date-time"`
	JSON        bankSettingsJSON `json:"-"`
}

// bankSettingsJSON contains the JSON metadata for the struct [BankSettings]
type bankSettingsJSON struct {
	ID                          apijson.Field
	BackupWithholdingPercentage apijson.Field
	CreatedAt                   apijson.Field
	DiscardedAt                 apijson.Field
	EnableBackupWithholding     apijson.Field
	LiveMode                    apijson.Field
	Object                      apijson.Field
	PrivacyOptOut               apijson.Field
	RegulationO                 apijson.Field
	UpdatedAt                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *BankSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bankSettingsJSON) RawJSON() string {
	return r.raw
}

type BankSettingsParam struct {
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

func (r BankSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntity struct {
	ID string `json:"id,required" format:"uuid"`
	// A list of addresses for the entity.
	Addresses    []LegalEntityAddress `json:"addresses,required"`
	BankSettings BankSettings         `json:"bank_settings,required,nullable"`
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
	Suffix                     string                     `json:"suffix,required,nullable"`
	UpdatedAt                  time.Time                  `json:"updated_at,required" format:"date-time"`
	WealthAndEmploymentDetails WealthAndEmploymentDetails `json:"wealth_and_employment_details,required,nullable"`
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

type WealthAndEmploymentDetails struct {
	ID string `json:"id,required" format:"uuid"`
	// The annual income of the individual.
	AnnualIncome int64     `json:"annual_income,required,nullable"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt  time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The country in which the employer is located.
	EmployerCountry string `json:"employer_country,required,nullable"`
	// The name of the employer.
	EmployerName string `json:"employer_name,required,nullable"`
	// The state in which the employer is located.
	EmployerState string `json:"employer_state,required,nullable"`
	// The employment status of the individual.
	EmploymentStatus WealthAndEmploymentDetailsEmploymentStatus `json:"employment_status,required,nullable"`
	// The country in which the individual's income is earned.
	IncomeCountry string `json:"income_country,required,nullable"`
	// The source of the individual's income.
	IncomeSource WealthAndEmploymentDetailsIncomeSource `json:"income_source,required,nullable"`
	// The state in which the individual's income is earned.
	IncomeState string `json:"income_state,required,nullable"`
	// The industry of the individual.
	Industry WealthAndEmploymentDetailsIndustry `json:"industry,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The occupation of the individual.
	Occupation WealthAndEmploymentDetailsOccupation `json:"occupation,required,nullable"`
	// The source of the individual's funds.
	SourceOfFunds WealthAndEmploymentDetailsSourceOfFunds `json:"source_of_funds,required,nullable"`
	UpdatedAt     time.Time                               `json:"updated_at,required" format:"date-time"`
	// The source of the individual's wealth.
	WealthSource WealthAndEmploymentDetailsWealthSource `json:"wealth_source,required,nullable"`
	JSON         wealthAndEmploymentDetailsJSON         `json:"-"`
}

// wealthAndEmploymentDetailsJSON contains the JSON metadata for the struct
// [WealthAndEmploymentDetails]
type wealthAndEmploymentDetailsJSON struct {
	ID               apijson.Field
	AnnualIncome     apijson.Field
	CreatedAt        apijson.Field
	DiscardedAt      apijson.Field
	EmployerCountry  apijson.Field
	EmployerName     apijson.Field
	EmployerState    apijson.Field
	EmploymentStatus apijson.Field
	IncomeCountry    apijson.Field
	IncomeSource     apijson.Field
	IncomeState      apijson.Field
	Industry         apijson.Field
	LiveMode         apijson.Field
	Object           apijson.Field
	Occupation       apijson.Field
	SourceOfFunds    apijson.Field
	UpdatedAt        apijson.Field
	WealthSource     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WealthAndEmploymentDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wealthAndEmploymentDetailsJSON) RawJSON() string {
	return r.raw
}

// The employment status of the individual.
type WealthAndEmploymentDetailsEmploymentStatus string

const (
	WealthAndEmploymentDetailsEmploymentStatusEmployed     WealthAndEmploymentDetailsEmploymentStatus = "employed"
	WealthAndEmploymentDetailsEmploymentStatusRetired      WealthAndEmploymentDetailsEmploymentStatus = "retired"
	WealthAndEmploymentDetailsEmploymentStatusSelfEmployed WealthAndEmploymentDetailsEmploymentStatus = "self_employed"
	WealthAndEmploymentDetailsEmploymentStatusStudent      WealthAndEmploymentDetailsEmploymentStatus = "student"
	WealthAndEmploymentDetailsEmploymentStatusUnemployed   WealthAndEmploymentDetailsEmploymentStatus = "unemployed"
)

func (r WealthAndEmploymentDetailsEmploymentStatus) IsKnown() bool {
	switch r {
	case WealthAndEmploymentDetailsEmploymentStatusEmployed, WealthAndEmploymentDetailsEmploymentStatusRetired, WealthAndEmploymentDetailsEmploymentStatusSelfEmployed, WealthAndEmploymentDetailsEmploymentStatusStudent, WealthAndEmploymentDetailsEmploymentStatusUnemployed:
		return true
	}
	return false
}

// The source of the individual's income.
type WealthAndEmploymentDetailsIncomeSource string

const (
	WealthAndEmploymentDetailsIncomeSourceFamilySupport      WealthAndEmploymentDetailsIncomeSource = "family_support"
	WealthAndEmploymentDetailsIncomeSourceGovernmentBenefits WealthAndEmploymentDetailsIncomeSource = "government_benefits"
	WealthAndEmploymentDetailsIncomeSourceInheritance        WealthAndEmploymentDetailsIncomeSource = "inheritance"
	WealthAndEmploymentDetailsIncomeSourceInvestments        WealthAndEmploymentDetailsIncomeSource = "investments"
	WealthAndEmploymentDetailsIncomeSourceRentalIncome       WealthAndEmploymentDetailsIncomeSource = "rental_income"
	WealthAndEmploymentDetailsIncomeSourceRetirement         WealthAndEmploymentDetailsIncomeSource = "retirement"
	WealthAndEmploymentDetailsIncomeSourceSalary             WealthAndEmploymentDetailsIncomeSource = "salary"
	WealthAndEmploymentDetailsIncomeSourceSelfEmployed       WealthAndEmploymentDetailsIncomeSource = "self_employed"
)

func (r WealthAndEmploymentDetailsIncomeSource) IsKnown() bool {
	switch r {
	case WealthAndEmploymentDetailsIncomeSourceFamilySupport, WealthAndEmploymentDetailsIncomeSourceGovernmentBenefits, WealthAndEmploymentDetailsIncomeSourceInheritance, WealthAndEmploymentDetailsIncomeSourceInvestments, WealthAndEmploymentDetailsIncomeSourceRentalIncome, WealthAndEmploymentDetailsIncomeSourceRetirement, WealthAndEmploymentDetailsIncomeSourceSalary, WealthAndEmploymentDetailsIncomeSourceSelfEmployed:
		return true
	}
	return false
}

// The industry of the individual.
type WealthAndEmploymentDetailsIndustry string

const (
	WealthAndEmploymentDetailsIndustryAccounting            WealthAndEmploymentDetailsIndustry = "accounting"
	WealthAndEmploymentDetailsIndustryAgriculture           WealthAndEmploymentDetailsIndustry = "agriculture"
	WealthAndEmploymentDetailsIndustryAutomotive            WealthAndEmploymentDetailsIndustry = "automotive"
	WealthAndEmploymentDetailsIndustryChemicalManufacturing WealthAndEmploymentDetailsIndustry = "chemical_manufacturing"
	WealthAndEmploymentDetailsIndustryConstruction          WealthAndEmploymentDetailsIndustry = "construction"
	WealthAndEmploymentDetailsIndustryEducationalMedical    WealthAndEmploymentDetailsIndustry = "educational_medical"
	WealthAndEmploymentDetailsIndustryFoodService           WealthAndEmploymentDetailsIndustry = "food_service"
	WealthAndEmploymentDetailsIndustryFinance               WealthAndEmploymentDetailsIndustry = "finance"
	WealthAndEmploymentDetailsIndustryGasoline              WealthAndEmploymentDetailsIndustry = "gasoline"
	WealthAndEmploymentDetailsIndustryHealthStores          WealthAndEmploymentDetailsIndustry = "health_stores"
	WealthAndEmploymentDetailsIndustryLaundry               WealthAndEmploymentDetailsIndustry = "laundry"
	WealthAndEmploymentDetailsIndustryMaintenance           WealthAndEmploymentDetailsIndustry = "maintenance"
	WealthAndEmploymentDetailsIndustryManufacturing         WealthAndEmploymentDetailsIndustry = "manufacturing"
	WealthAndEmploymentDetailsIndustryMerchantWholesale     WealthAndEmploymentDetailsIndustry = "merchant_wholesale"
	WealthAndEmploymentDetailsIndustryMining                WealthAndEmploymentDetailsIndustry = "mining"
	WealthAndEmploymentDetailsIndustryPerformingArts        WealthAndEmploymentDetailsIndustry = "performing_arts"
	WealthAndEmploymentDetailsIndustryProfessionalNonLegal  WealthAndEmploymentDetailsIndustry = "professional_non_legal"
	WealthAndEmploymentDetailsIndustryPublicAdministration  WealthAndEmploymentDetailsIndustry = "public_administration"
	WealthAndEmploymentDetailsIndustryPublishing            WealthAndEmploymentDetailsIndustry = "publishing"
	WealthAndEmploymentDetailsIndustryRealEstate            WealthAndEmploymentDetailsIndustry = "real_estate"
	WealthAndEmploymentDetailsIndustryRecreationGambling    WealthAndEmploymentDetailsIndustry = "recreation_gambling"
	WealthAndEmploymentDetailsIndustryReligiousCharity      WealthAndEmploymentDetailsIndustry = "religious_charity"
	WealthAndEmploymentDetailsIndustryRentalServices        WealthAndEmploymentDetailsIndustry = "rental_services"
	WealthAndEmploymentDetailsIndustryRetailClothing        WealthAndEmploymentDetailsIndustry = "retail_clothing"
	WealthAndEmploymentDetailsIndustryRetailElectronics     WealthAndEmploymentDetailsIndustry = "retail_electronics"
	WealthAndEmploymentDetailsIndustryRetailFood            WealthAndEmploymentDetailsIndustry = "retail_food"
	WealthAndEmploymentDetailsIndustryRetailFurnishing      WealthAndEmploymentDetailsIndustry = "retail_furnishing"
	WealthAndEmploymentDetailsIndustryRetailHome            WealthAndEmploymentDetailsIndustry = "retail_home"
	WealthAndEmploymentDetailsIndustryRetailNonStore        WealthAndEmploymentDetailsIndustry = "retail_non_store"
	WealthAndEmploymentDetailsIndustryRetailSporting        WealthAndEmploymentDetailsIndustry = "retail_sporting"
	WealthAndEmploymentDetailsIndustryTransportation        WealthAndEmploymentDetailsIndustry = "transportation"
	WealthAndEmploymentDetailsIndustryTravel                WealthAndEmploymentDetailsIndustry = "travel"
	WealthAndEmploymentDetailsIndustryUtilities             WealthAndEmploymentDetailsIndustry = "utilities"
)

func (r WealthAndEmploymentDetailsIndustry) IsKnown() bool {
	switch r {
	case WealthAndEmploymentDetailsIndustryAccounting, WealthAndEmploymentDetailsIndustryAgriculture, WealthAndEmploymentDetailsIndustryAutomotive, WealthAndEmploymentDetailsIndustryChemicalManufacturing, WealthAndEmploymentDetailsIndustryConstruction, WealthAndEmploymentDetailsIndustryEducationalMedical, WealthAndEmploymentDetailsIndustryFoodService, WealthAndEmploymentDetailsIndustryFinance, WealthAndEmploymentDetailsIndustryGasoline, WealthAndEmploymentDetailsIndustryHealthStores, WealthAndEmploymentDetailsIndustryLaundry, WealthAndEmploymentDetailsIndustryMaintenance, WealthAndEmploymentDetailsIndustryManufacturing, WealthAndEmploymentDetailsIndustryMerchantWholesale, WealthAndEmploymentDetailsIndustryMining, WealthAndEmploymentDetailsIndustryPerformingArts, WealthAndEmploymentDetailsIndustryProfessionalNonLegal, WealthAndEmploymentDetailsIndustryPublicAdministration, WealthAndEmploymentDetailsIndustryPublishing, WealthAndEmploymentDetailsIndustryRealEstate, WealthAndEmploymentDetailsIndustryRecreationGambling, WealthAndEmploymentDetailsIndustryReligiousCharity, WealthAndEmploymentDetailsIndustryRentalServices, WealthAndEmploymentDetailsIndustryRetailClothing, WealthAndEmploymentDetailsIndustryRetailElectronics, WealthAndEmploymentDetailsIndustryRetailFood, WealthAndEmploymentDetailsIndustryRetailFurnishing, WealthAndEmploymentDetailsIndustryRetailHome, WealthAndEmploymentDetailsIndustryRetailNonStore, WealthAndEmploymentDetailsIndustryRetailSporting, WealthAndEmploymentDetailsIndustryTransportation, WealthAndEmploymentDetailsIndustryTravel, WealthAndEmploymentDetailsIndustryUtilities:
		return true
	}
	return false
}

// The occupation of the individual.
type WealthAndEmploymentDetailsOccupation string

const (
	WealthAndEmploymentDetailsOccupationConsulting         WealthAndEmploymentDetailsOccupation = "consulting"
	WealthAndEmploymentDetailsOccupationExecutive          WealthAndEmploymentDetailsOccupation = "executive"
	WealthAndEmploymentDetailsOccupationFinanceAccounting  WealthAndEmploymentDetailsOccupation = "finance_accounting"
	WealthAndEmploymentDetailsOccupationFoodServices       WealthAndEmploymentDetailsOccupation = "food_services"
	WealthAndEmploymentDetailsOccupationGovernment         WealthAndEmploymentDetailsOccupation = "government"
	WealthAndEmploymentDetailsOccupationHealthcare         WealthAndEmploymentDetailsOccupation = "healthcare"
	WealthAndEmploymentDetailsOccupationLegalServices      WealthAndEmploymentDetailsOccupation = "legal_services"
	WealthAndEmploymentDetailsOccupationManufacturing      WealthAndEmploymentDetailsOccupation = "manufacturing"
	WealthAndEmploymentDetailsOccupationOther              WealthAndEmploymentDetailsOccupation = "other"
	WealthAndEmploymentDetailsOccupationSales              WealthAndEmploymentDetailsOccupation = "sales"
	WealthAndEmploymentDetailsOccupationScienceEngineering WealthAndEmploymentDetailsOccupation = "science_engineering"
	WealthAndEmploymentDetailsOccupationTechnology         WealthAndEmploymentDetailsOccupation = "technology"
)

func (r WealthAndEmploymentDetailsOccupation) IsKnown() bool {
	switch r {
	case WealthAndEmploymentDetailsOccupationConsulting, WealthAndEmploymentDetailsOccupationExecutive, WealthAndEmploymentDetailsOccupationFinanceAccounting, WealthAndEmploymentDetailsOccupationFoodServices, WealthAndEmploymentDetailsOccupationGovernment, WealthAndEmploymentDetailsOccupationHealthcare, WealthAndEmploymentDetailsOccupationLegalServices, WealthAndEmploymentDetailsOccupationManufacturing, WealthAndEmploymentDetailsOccupationOther, WealthAndEmploymentDetailsOccupationSales, WealthAndEmploymentDetailsOccupationScienceEngineering, WealthAndEmploymentDetailsOccupationTechnology:
		return true
	}
	return false
}

// The source of the individual's funds.
type WealthAndEmploymentDetailsSourceOfFunds string

const (
	WealthAndEmploymentDetailsSourceOfFundsAlimony            WealthAndEmploymentDetailsSourceOfFunds = "alimony"
	WealthAndEmploymentDetailsSourceOfFundsAnnuity            WealthAndEmploymentDetailsSourceOfFunds = "annuity"
	WealthAndEmploymentDetailsSourceOfFundsBusinessOwner      WealthAndEmploymentDetailsSourceOfFunds = "business_owner"
	WealthAndEmploymentDetailsSourceOfFundsGeneralEmployee    WealthAndEmploymentDetailsSourceOfFunds = "general_employee"
	WealthAndEmploymentDetailsSourceOfFundsGovernmentBenefits WealthAndEmploymentDetailsSourceOfFunds = "government_benefits"
	WealthAndEmploymentDetailsSourceOfFundsHomemaker          WealthAndEmploymentDetailsSourceOfFunds = "homemaker"
	WealthAndEmploymentDetailsSourceOfFundsInheritanceGift    WealthAndEmploymentDetailsSourceOfFunds = "inheritance_gift"
	WealthAndEmploymentDetailsSourceOfFundsInvestment         WealthAndEmploymentDetailsSourceOfFunds = "investment"
	WealthAndEmploymentDetailsSourceOfFundsLegalSettlement    WealthAndEmploymentDetailsSourceOfFunds = "legal_settlement"
	WealthAndEmploymentDetailsSourceOfFundsLottery            WealthAndEmploymentDetailsSourceOfFunds = "lottery"
	WealthAndEmploymentDetailsSourceOfFundsRealEstate         WealthAndEmploymentDetailsSourceOfFunds = "real_estate"
	WealthAndEmploymentDetailsSourceOfFundsRetired            WealthAndEmploymentDetailsSourceOfFunds = "retired"
	WealthAndEmploymentDetailsSourceOfFundsRetirement         WealthAndEmploymentDetailsSourceOfFunds = "retirement"
	WealthAndEmploymentDetailsSourceOfFundsSalary             WealthAndEmploymentDetailsSourceOfFunds = "salary"
	WealthAndEmploymentDetailsSourceOfFundsSelfEmployed       WealthAndEmploymentDetailsSourceOfFunds = "self_employed"
	WealthAndEmploymentDetailsSourceOfFundsSeniorExecutive    WealthAndEmploymentDetailsSourceOfFunds = "senior_executive"
	WealthAndEmploymentDetailsSourceOfFundsTrustIncome        WealthAndEmploymentDetailsSourceOfFunds = "trust_income"
)

func (r WealthAndEmploymentDetailsSourceOfFunds) IsKnown() bool {
	switch r {
	case WealthAndEmploymentDetailsSourceOfFundsAlimony, WealthAndEmploymentDetailsSourceOfFundsAnnuity, WealthAndEmploymentDetailsSourceOfFundsBusinessOwner, WealthAndEmploymentDetailsSourceOfFundsGeneralEmployee, WealthAndEmploymentDetailsSourceOfFundsGovernmentBenefits, WealthAndEmploymentDetailsSourceOfFundsHomemaker, WealthAndEmploymentDetailsSourceOfFundsInheritanceGift, WealthAndEmploymentDetailsSourceOfFundsInvestment, WealthAndEmploymentDetailsSourceOfFundsLegalSettlement, WealthAndEmploymentDetailsSourceOfFundsLottery, WealthAndEmploymentDetailsSourceOfFundsRealEstate, WealthAndEmploymentDetailsSourceOfFundsRetired, WealthAndEmploymentDetailsSourceOfFundsRetirement, WealthAndEmploymentDetailsSourceOfFundsSalary, WealthAndEmploymentDetailsSourceOfFundsSelfEmployed, WealthAndEmploymentDetailsSourceOfFundsSeniorExecutive, WealthAndEmploymentDetailsSourceOfFundsTrustIncome:
		return true
	}
	return false
}

// The source of the individual's wealth.
type WealthAndEmploymentDetailsWealthSource string

const (
	WealthAndEmploymentDetailsWealthSourceBusinessSale       WealthAndEmploymentDetailsWealthSource = "business_sale"
	WealthAndEmploymentDetailsWealthSourceFamilySupport      WealthAndEmploymentDetailsWealthSource = "family_support"
	WealthAndEmploymentDetailsWealthSourceGovernmentBenefits WealthAndEmploymentDetailsWealthSource = "government_benefits"
	WealthAndEmploymentDetailsWealthSourceInheritance        WealthAndEmploymentDetailsWealthSource = "inheritance"
	WealthAndEmploymentDetailsWealthSourceInvestments        WealthAndEmploymentDetailsWealthSource = "investments"
	WealthAndEmploymentDetailsWealthSourceOther              WealthAndEmploymentDetailsWealthSource = "other"
	WealthAndEmploymentDetailsWealthSourceRentalIncome       WealthAndEmploymentDetailsWealthSource = "rental_income"
	WealthAndEmploymentDetailsWealthSourceRetirement         WealthAndEmploymentDetailsWealthSource = "retirement"
	WealthAndEmploymentDetailsWealthSourceSalary             WealthAndEmploymentDetailsWealthSource = "salary"
	WealthAndEmploymentDetailsWealthSourceSelfEmployed       WealthAndEmploymentDetailsWealthSource = "self_employed"
)

func (r WealthAndEmploymentDetailsWealthSource) IsKnown() bool {
	switch r {
	case WealthAndEmploymentDetailsWealthSourceBusinessSale, WealthAndEmploymentDetailsWealthSourceFamilySupport, WealthAndEmploymentDetailsWealthSourceGovernmentBenefits, WealthAndEmploymentDetailsWealthSourceInheritance, WealthAndEmploymentDetailsWealthSourceInvestments, WealthAndEmploymentDetailsWealthSourceOther, WealthAndEmploymentDetailsWealthSourceRentalIncome, WealthAndEmploymentDetailsWealthSourceRetirement, WealthAndEmploymentDetailsWealthSourceSalary, WealthAndEmploymentDetailsWealthSourceSelfEmployed:
		return true
	}
	return false
}

type WealthAndEmploymentDetailsParam struct {
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
	EmploymentStatus param.Field[WealthAndEmploymentDetailsEmploymentStatus] `json:"employment_status,required"`
	// The country in which the individual's income is earned.
	IncomeCountry param.Field[string] `json:"income_country,required"`
	// The source of the individual's income.
	IncomeSource param.Field[WealthAndEmploymentDetailsIncomeSource] `json:"income_source,required"`
	// The state in which the individual's income is earned.
	IncomeState param.Field[string] `json:"income_state,required"`
	// The industry of the individual.
	Industry param.Field[WealthAndEmploymentDetailsIndustry] `json:"industry,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// The occupation of the individual.
	Occupation param.Field[WealthAndEmploymentDetailsOccupation] `json:"occupation,required"`
	// The source of the individual's funds.
	SourceOfFunds param.Field[WealthAndEmploymentDetailsSourceOfFunds] `json:"source_of_funds,required"`
	UpdatedAt     param.Field[time.Time]                               `json:"updated_at,required" format:"date-time"`
	// The source of the individual's wealth.
	WealthSource param.Field[WealthAndEmploymentDetailsWealthSource] `json:"wealth_source,required"`
}

func (r WealthAndEmploymentDetailsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParams struct {
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityType] `json:"legal_entity_type,required"`
	// A list of addresses for the entity.
	Addresses    param.Field[[]LegalEntityNewParamsAddress] `json:"addresses"`
	BankSettings param.Field[BankSettingsParam]             `json:"bank_settings"`
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
	Identifications param.Field[[]LegalEntityNewParamsIdentification] `json:"identifications"`
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
	Suffix                     param.Field[string]                          `json:"suffix"`
	WealthAndEmploymentDetails param.Field[WealthAndEmploymentDetailsParam] `json:"wealth_and_employment_details"`
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

type LegalEntityNewParamsAddress struct {
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
	AddressTypes param.Field[[]LegalEntityNewParamsAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                     `json:"line2"`
}

func (r LegalEntityNewParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsAddressesAddressType string

const (
	LegalEntityNewParamsAddressesAddressTypeBusiness    LegalEntityNewParamsAddressesAddressType = "business"
	LegalEntityNewParamsAddressesAddressTypeMailing     LegalEntityNewParamsAddressesAddressType = "mailing"
	LegalEntityNewParamsAddressesAddressTypeOther       LegalEntityNewParamsAddressesAddressType = "other"
	LegalEntityNewParamsAddressesAddressTypePoBox       LegalEntityNewParamsAddressesAddressType = "po_box"
	LegalEntityNewParamsAddressesAddressTypeResidential LegalEntityNewParamsAddressesAddressType = "residential"
)

func (r LegalEntityNewParamsAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsAddressesAddressTypeBusiness, LegalEntityNewParamsAddressesAddressTypeMailing, LegalEntityNewParamsAddressesAddressTypeOther, LegalEntityNewParamsAddressesAddressTypePoBox, LegalEntityNewParamsAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityNewParamsIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[LegalEntityNewParamsIdentificationsIDType] `json:"id_type,required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
	// The region in which the identifcation was issued.
	IssuingRegion param.Field[string] `json:"issuing_region"`
}

func (r LegalEntityNewParamsIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityNewParamsIdentificationsIDType string

const (
	LegalEntityNewParamsIdentificationsIDTypeArCuil         LegalEntityNewParamsIdentificationsIDType = "ar_cuil"
	LegalEntityNewParamsIdentificationsIDTypeArCuit         LegalEntityNewParamsIdentificationsIDType = "ar_cuit"
	LegalEntityNewParamsIdentificationsIDTypeBrCnpj         LegalEntityNewParamsIdentificationsIDType = "br_cnpj"
	LegalEntityNewParamsIdentificationsIDTypeBrCpf          LegalEntityNewParamsIdentificationsIDType = "br_cpf"
	LegalEntityNewParamsIdentificationsIDTypeClRun          LegalEntityNewParamsIdentificationsIDType = "cl_run"
	LegalEntityNewParamsIdentificationsIDTypeClRut          LegalEntityNewParamsIdentificationsIDType = "cl_rut"
	LegalEntityNewParamsIdentificationsIDTypeCoCedulas      LegalEntityNewParamsIdentificationsIDType = "co_cedulas"
	LegalEntityNewParamsIdentificationsIDTypeCoNit          LegalEntityNewParamsIdentificationsIDType = "co_nit"
	LegalEntityNewParamsIdentificationsIDTypeDriversLicense LegalEntityNewParamsIdentificationsIDType = "drivers_license"
	LegalEntityNewParamsIdentificationsIDTypeHnID           LegalEntityNewParamsIdentificationsIDType = "hn_id"
	LegalEntityNewParamsIdentificationsIDTypeHnRtn          LegalEntityNewParamsIdentificationsIDType = "hn_rtn"
	LegalEntityNewParamsIdentificationsIDTypeInLei          LegalEntityNewParamsIdentificationsIDType = "in_lei"
	LegalEntityNewParamsIdentificationsIDTypeKrBrn          LegalEntityNewParamsIdentificationsIDType = "kr_brn"
	LegalEntityNewParamsIdentificationsIDTypeKrCrn          LegalEntityNewParamsIdentificationsIDType = "kr_crn"
	LegalEntityNewParamsIdentificationsIDTypeKrRrn          LegalEntityNewParamsIdentificationsIDType = "kr_rrn"
	LegalEntityNewParamsIdentificationsIDTypePassport       LegalEntityNewParamsIdentificationsIDType = "passport"
	LegalEntityNewParamsIdentificationsIDTypeSaTin          LegalEntityNewParamsIdentificationsIDType = "sa_tin"
	LegalEntityNewParamsIdentificationsIDTypeSaVat          LegalEntityNewParamsIdentificationsIDType = "sa_vat"
	LegalEntityNewParamsIdentificationsIDTypeUsEin          LegalEntityNewParamsIdentificationsIDType = "us_ein"
	LegalEntityNewParamsIdentificationsIDTypeUsItin         LegalEntityNewParamsIdentificationsIDType = "us_itin"
	LegalEntityNewParamsIdentificationsIDTypeUsSsn          LegalEntityNewParamsIdentificationsIDType = "us_ssn"
	LegalEntityNewParamsIdentificationsIDTypeVnTin          LegalEntityNewParamsIdentificationsIDType = "vn_tin"
)

func (r LegalEntityNewParamsIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsIdentificationsIDTypeArCuil, LegalEntityNewParamsIdentificationsIDTypeArCuit, LegalEntityNewParamsIdentificationsIDTypeBrCnpj, LegalEntityNewParamsIdentificationsIDTypeBrCpf, LegalEntityNewParamsIdentificationsIDTypeClRun, LegalEntityNewParamsIdentificationsIDTypeClRut, LegalEntityNewParamsIdentificationsIDTypeCoCedulas, LegalEntityNewParamsIdentificationsIDTypeCoNit, LegalEntityNewParamsIdentificationsIDTypeDriversLicense, LegalEntityNewParamsIdentificationsIDTypeHnID, LegalEntityNewParamsIdentificationsIDTypeHnRtn, LegalEntityNewParamsIdentificationsIDTypeInLei, LegalEntityNewParamsIdentificationsIDTypeKrBrn, LegalEntityNewParamsIdentificationsIDTypeKrCrn, LegalEntityNewParamsIdentificationsIDTypeKrRrn, LegalEntityNewParamsIdentificationsIDTypePassport, LegalEntityNewParamsIdentificationsIDTypeSaTin, LegalEntityNewParamsIdentificationsIDTypeSaVat, LegalEntityNewParamsIdentificationsIDTypeUsEin, LegalEntityNewParamsIdentificationsIDTypeUsItin, LegalEntityNewParamsIdentificationsIDTypeUsSsn, LegalEntityNewParamsIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

type LegalEntityNewParamsLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]LegalEntityNewParamsLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity] `json:"child_legal_entity"`
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
	LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner LegalEntityNewParamsLegalEntityAssociationsRelationshipType = "beneficial_owner"
	LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson   LegalEntityNewParamsLegalEntityAssociationsRelationshipType = "control_person"
)

func (r LegalEntityNewParamsLegalEntityAssociationsRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson:
		return true
	}
	return false
}

// The child legal entity.
type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity struct {
	// A list of addresses for the entity.
	Addresses    param.Field[[]LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress] `json:"addresses"`
	BankSettings param.Field[BankSettingsParam]                                                    `json:"bank_settings"`
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
	Identifications param.Field[[]LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification] `json:"identifications"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications param.Field[[]shared.LegalEntityIndustryClassificationParam] `json:"industry_classifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName   param.Field[string]                                                                   `json:"middle_name"`
	PhoneNumbers param.Field[[]LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRating] `json:"risk_rating"`
	// An individual's suffix.
	Suffix                     param.Field[string]                          `json:"suffix"`
	WealthAndEmploymentDetails param.Field[WealthAndEmploymentDetailsParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress struct {
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
	AddressTypes param.Field[[]LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                                            `json:"line2"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType string

const (
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness    LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType = "business"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing     LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType = "mailing"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther       LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType = "other"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypePoBox       LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType = "po_box"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeResidential LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType = "residential"
)

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypePoBox, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
	// The region in which the identifcation was issued.
	IssuingRegion param.Field[string] `json:"issuing_region"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType string

const (
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil         LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuit         LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCnpj         LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCpf          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeClRun          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "cl_run"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeClRut          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "cl_rut"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoCedulas      LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoNit          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "co_nit"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeDriversLicense LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "drivers_license"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnID           LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "hn_id"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnRtn          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeInLei          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "in_lei"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeKrBrn          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "kr_brn"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeKrCrn          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "kr_crn"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeKrRrn          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "kr_rrn"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypePassport       LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "passport"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaTin          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "sa_tin"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaVat          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "sa_vat"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsEin          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "us_ein"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsItin         LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "us_itin"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsSsn          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "us_ssn"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeVnTin          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "vn_tin"
)

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuit, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCnpj, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCpf, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeClRun, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeClRut, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoCedulas, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoNit, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeDriversLicense, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnID, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnRtn, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeInLei, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeKrBrn, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeKrCrn, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeKrRrn, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypePassport, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaTin, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaVat, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsEin, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsItin, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsSsn, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

// The type of legal entity.
type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityType string

const (
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness   LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityType = "business"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeIndividual LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityType = "individual"
)

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure string

const (
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation        LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure = "corporation"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureLlc                LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure = "llc"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureNonProfit          LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure = "non_profit"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructurePartnership        LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure = "partnership"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureSoleProprietorship LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure = "sole_proprietorship"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureTrust              LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure = "trust"
)

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureLlc, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureNonProfit, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructurePartnership, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureSoleProprietorship, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The risk rating of the legal entity. One of low, medium, high.
type LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRating string

const (
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingLow    LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRating = "low"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingMedium LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRating = "medium"
	LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingHigh   LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRating = "high"
)

func (r LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRating) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingLow, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingMedium, LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingHigh:
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

type LegalEntityUpdateParams struct {
	// A list of addresses for the entity.
	Addresses    param.Field[[]LegalEntityUpdateParamsAddress] `json:"addresses"`
	BankSettings param.Field[BankSettingsParam]                `json:"bank_settings"`
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
	Identifications param.Field[[]LegalEntityUpdateParamsIdentification] `json:"identifications"`
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
	Suffix                     param.Field[string]                          `json:"suffix"`
	WealthAndEmploymentDetails param.Field[WealthAndEmploymentDetailsParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityUpdateParamsAddress struct {
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
	AddressTypes param.Field[[]LegalEntityUpdateParamsAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                        `json:"line2"`
}

func (r LegalEntityUpdateParamsAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityUpdateParamsAddressesAddressType string

const (
	LegalEntityUpdateParamsAddressesAddressTypeBusiness    LegalEntityUpdateParamsAddressesAddressType = "business"
	LegalEntityUpdateParamsAddressesAddressTypeMailing     LegalEntityUpdateParamsAddressesAddressType = "mailing"
	LegalEntityUpdateParamsAddressesAddressTypeOther       LegalEntityUpdateParamsAddressesAddressType = "other"
	LegalEntityUpdateParamsAddressesAddressTypePoBox       LegalEntityUpdateParamsAddressesAddressType = "po_box"
	LegalEntityUpdateParamsAddressesAddressTypeResidential LegalEntityUpdateParamsAddressesAddressType = "residential"
)

func (r LegalEntityUpdateParamsAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsAddressesAddressTypeBusiness, LegalEntityUpdateParamsAddressesAddressTypeMailing, LegalEntityUpdateParamsAddressesAddressTypeOther, LegalEntityUpdateParamsAddressesAddressTypePoBox, LegalEntityUpdateParamsAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityUpdateParamsIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[LegalEntityUpdateParamsIdentificationsIDType] `json:"id_type,required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
	// The region in which the identifcation was issued.
	IssuingRegion param.Field[string] `json:"issuing_region"`
}

func (r LegalEntityUpdateParamsIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityUpdateParamsIdentificationsIDType string

const (
	LegalEntityUpdateParamsIdentificationsIDTypeArCuil         LegalEntityUpdateParamsIdentificationsIDType = "ar_cuil"
	LegalEntityUpdateParamsIdentificationsIDTypeArCuit         LegalEntityUpdateParamsIdentificationsIDType = "ar_cuit"
	LegalEntityUpdateParamsIdentificationsIDTypeBrCnpj         LegalEntityUpdateParamsIdentificationsIDType = "br_cnpj"
	LegalEntityUpdateParamsIdentificationsIDTypeBrCpf          LegalEntityUpdateParamsIdentificationsIDType = "br_cpf"
	LegalEntityUpdateParamsIdentificationsIDTypeClRun          LegalEntityUpdateParamsIdentificationsIDType = "cl_run"
	LegalEntityUpdateParamsIdentificationsIDTypeClRut          LegalEntityUpdateParamsIdentificationsIDType = "cl_rut"
	LegalEntityUpdateParamsIdentificationsIDTypeCoCedulas      LegalEntityUpdateParamsIdentificationsIDType = "co_cedulas"
	LegalEntityUpdateParamsIdentificationsIDTypeCoNit          LegalEntityUpdateParamsIdentificationsIDType = "co_nit"
	LegalEntityUpdateParamsIdentificationsIDTypeDriversLicense LegalEntityUpdateParamsIdentificationsIDType = "drivers_license"
	LegalEntityUpdateParamsIdentificationsIDTypeHnID           LegalEntityUpdateParamsIdentificationsIDType = "hn_id"
	LegalEntityUpdateParamsIdentificationsIDTypeHnRtn          LegalEntityUpdateParamsIdentificationsIDType = "hn_rtn"
	LegalEntityUpdateParamsIdentificationsIDTypeInLei          LegalEntityUpdateParamsIdentificationsIDType = "in_lei"
	LegalEntityUpdateParamsIdentificationsIDTypeKrBrn          LegalEntityUpdateParamsIdentificationsIDType = "kr_brn"
	LegalEntityUpdateParamsIdentificationsIDTypeKrCrn          LegalEntityUpdateParamsIdentificationsIDType = "kr_crn"
	LegalEntityUpdateParamsIdentificationsIDTypeKrRrn          LegalEntityUpdateParamsIdentificationsIDType = "kr_rrn"
	LegalEntityUpdateParamsIdentificationsIDTypePassport       LegalEntityUpdateParamsIdentificationsIDType = "passport"
	LegalEntityUpdateParamsIdentificationsIDTypeSaTin          LegalEntityUpdateParamsIdentificationsIDType = "sa_tin"
	LegalEntityUpdateParamsIdentificationsIDTypeSaVat          LegalEntityUpdateParamsIdentificationsIDType = "sa_vat"
	LegalEntityUpdateParamsIdentificationsIDTypeUsEin          LegalEntityUpdateParamsIdentificationsIDType = "us_ein"
	LegalEntityUpdateParamsIdentificationsIDTypeUsItin         LegalEntityUpdateParamsIdentificationsIDType = "us_itin"
	LegalEntityUpdateParamsIdentificationsIDTypeUsSsn          LegalEntityUpdateParamsIdentificationsIDType = "us_ssn"
	LegalEntityUpdateParamsIdentificationsIDTypeVnTin          LegalEntityUpdateParamsIdentificationsIDType = "vn_tin"
)

func (r LegalEntityUpdateParamsIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityUpdateParamsIdentificationsIDTypeArCuil, LegalEntityUpdateParamsIdentificationsIDTypeArCuit, LegalEntityUpdateParamsIdentificationsIDTypeBrCnpj, LegalEntityUpdateParamsIdentificationsIDTypeBrCpf, LegalEntityUpdateParamsIdentificationsIDTypeClRun, LegalEntityUpdateParamsIdentificationsIDTypeClRut, LegalEntityUpdateParamsIdentificationsIDTypeCoCedulas, LegalEntityUpdateParamsIdentificationsIDTypeCoNit, LegalEntityUpdateParamsIdentificationsIDTypeDriversLicense, LegalEntityUpdateParamsIdentificationsIDTypeHnID, LegalEntityUpdateParamsIdentificationsIDTypeHnRtn, LegalEntityUpdateParamsIdentificationsIDTypeInLei, LegalEntityUpdateParamsIdentificationsIDTypeKrBrn, LegalEntityUpdateParamsIdentificationsIDTypeKrCrn, LegalEntityUpdateParamsIdentificationsIDTypeKrRrn, LegalEntityUpdateParamsIdentificationsIDTypePassport, LegalEntityUpdateParamsIdentificationsIDTypeSaTin, LegalEntityUpdateParamsIdentificationsIDTypeSaVat, LegalEntityUpdateParamsIdentificationsIDTypeUsEin, LegalEntityUpdateParamsIdentificationsIDTypeUsItin, LegalEntityUpdateParamsIdentificationsIDTypeUsSsn, LegalEntityUpdateParamsIdentificationsIDTypeVnTin:
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
