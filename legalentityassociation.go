// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
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

type ChildLegalEntity struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// A list of addresses for the entity.
	Addresses    []ChildLegalEntityAddress      `json:"addresses" api:"required"`
	BankSettings shared.LegalEntityBankSettings `json:"bank_settings" api:"required,nullable"`
	// A description of the business.
	BusinessDescription string `json:"business_description" api:"required,nullable"`
	// Legal designation associated with the business.
	BusinessDesignation ChildLegalEntityBusinessDesignation `json:"business_designation" api:"required,nullable"`
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
	Identifications []ChildLegalEntityIdentification `json:"identifications" api:"required"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications []shared.LegalEntityIndustryClassification `json:"industry_classifications" api:"required"`
	// A description of the intended use of the legal entity.
	IntendedUse string `json:"intended_use" api:"required,nullable"`
	// An individual's last name.
	LastName string `json:"last_name" api:"required,nullable"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations []LegalEntityAssociation `json:"legal_entity_associations" api:"required,nullable"`
	// The type of legal entity.
	LegalEntityType ChildLegalEntityLegalEntityType `json:"legal_entity_type" api:"required"`
	// The business's legal structure.
	LegalStructure ChildLegalEntityLegalStructure `json:"legal_structure" api:"required,nullable"`
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
	OperatingJurisdictions []string                      `json:"operating_jurisdictions" api:"required"`
	PhoneNumbers           []ChildLegalEntityPhoneNumber `json:"phone_numbers" api:"required"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson bool `json:"politically_exposed_person" api:"required,nullable"`
	// An individual's preferred name.
	PreferredName string `json:"preferred_name" api:"required,nullable"`
	// An individual's prefix.
	Prefix string `json:"prefix" api:"required,nullable"`
	// A list of primary social media URLs for the business.
	PrimarySocialMediaSites []string `json:"primary_social_media_sites" api:"required"`
	// Array of regulatory bodies overseeing this institution.
	Regulators []ChildLegalEntityRegulator `json:"regulators" api:"required,nullable"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating ChildLegalEntityRiskRating `json:"risk_rating" api:"required,nullable"`
	// The UUID of the parent legal entity in the service provider tree.
	ServiceProviderLegalEntityID string `json:"service_provider_legal_entity_id" api:"required,nullable" format:"uuid"`
	// The activation status of the legal entity. One of pending, active, suspended, or
	// denied.
	Status ChildLegalEntityStatus `json:"status" api:"required,nullable"`
	// An individual's suffix.
	Suffix string `json:"suffix" api:"required,nullable"`
	// Acceptance of terms of use by the legal entity.
	TermsOfUse ChildLegalEntityTermsOfUse `json:"terms_of_use" api:"required,nullable"`
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
	Website string               `json:"website" api:"required,nullable"`
	JSON    childLegalEntityJSON `json:"-"`
}

// childLegalEntityJSON contains the JSON metadata for the struct
// [ChildLegalEntity]
type childLegalEntityJSON struct {
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
	LegalEntityAssociations      apijson.Field
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
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ChildLegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityJSON) RawJSON() string {
	return r.raw
}

type ChildLegalEntityAddress struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The types of this address.
	AddressTypes []ChildLegalEntityAddressesAddressType `json:"address_types" api:"required"`
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
	Region    string                      `json:"region" api:"required,nullable"`
	UpdatedAt time.Time                   `json:"updated_at" api:"required" format:"date-time"`
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
	Primary      apijson.Field
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
	ChildLegalEntityAddressesAddressTypeBusiness           ChildLegalEntityAddressesAddressType = "business"
	ChildLegalEntityAddressesAddressTypeBusinessPhysical   ChildLegalEntityAddressesAddressType = "business_physical"
	ChildLegalEntityAddressesAddressTypeBusinessRegistered ChildLegalEntityAddressesAddressType = "business_registered"
	ChildLegalEntityAddressesAddressTypeMailing            ChildLegalEntityAddressesAddressType = "mailing"
	ChildLegalEntityAddressesAddressTypeOther              ChildLegalEntityAddressesAddressType = "other"
	ChildLegalEntityAddressesAddressTypePoBox              ChildLegalEntityAddressesAddressType = "po_box"
	ChildLegalEntityAddressesAddressTypeResidential        ChildLegalEntityAddressesAddressType = "residential"
)

func (r ChildLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case ChildLegalEntityAddressesAddressTypeBusiness, ChildLegalEntityAddressesAddressTypeBusinessPhysical, ChildLegalEntityAddressesAddressTypeBusinessRegistered, ChildLegalEntityAddressesAddressTypeMailing, ChildLegalEntityAddressesAddressTypeOther, ChildLegalEntityAddressesAddressTypePoBox, ChildLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

// Legal designation associated with the business.
type ChildLegalEntityBusinessDesignation string

const (
	ChildLegalEntityBusinessDesignationExemptFinancialInstitution ChildLegalEntityBusinessDesignation = "exempt_financial_institution"
	ChildLegalEntityBusinessDesignationNonOperatingBusiness       ChildLegalEntityBusinessDesignation = "non_operating_business"
)

func (r ChildLegalEntityBusinessDesignation) IsKnown() bool {
	switch r {
	case ChildLegalEntityBusinessDesignationExemptFinancialInstitution, ChildLegalEntityBusinessDesignationNonOperatingBusiness:
		return true
	}
	return false
}

type ChildLegalEntityIdentification struct {
	ID          string     `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time  `json:"created_at" api:"required" format:"date-time"`
	DiscardedAt time.Time  `json:"discarded_at" api:"required,nullable" format:"date-time"`
	Documents   []Document `json:"documents" api:"required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate time.Time `json:"expiration_date" api:"required,nullable" format:"date"`
	// The type of ID number.
	IDType ChildLegalEntityIdentificationsIDType `json:"id_type" api:"required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country" api:"required,nullable"`
	// The region in which the identifcation was issued.
	IssuingRegion string `json:"issuing_region" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                               `json:"live_mode" api:"required"`
	Object    string                             `json:"object" api:"required"`
	UpdatedAt time.Time                          `json:"updated_at" api:"required" format:"date-time"`
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
	ChildLegalEntityIdentificationsIDTypeArCuil               ChildLegalEntityIdentificationsIDType = "ar_cuil"
	ChildLegalEntityIdentificationsIDTypeArCuit               ChildLegalEntityIdentificationsIDType = "ar_cuit"
	ChildLegalEntityIdentificationsIDTypeAtAtin               ChildLegalEntityIdentificationsIDType = "at_atin"
	ChildLegalEntityIdentificationsIDTypeAtVat                ChildLegalEntityIdentificationsIDType = "at_vat"
	ChildLegalEntityIdentificationsIDTypeAuAbn                ChildLegalEntityIdentificationsIDType = "au_abn"
	ChildLegalEntityIdentificationsIDTypeAuTfn                ChildLegalEntityIdentificationsIDType = "au_tfn"
	ChildLegalEntityIdentificationsIDTypeBeEnt                ChildLegalEntityIdentificationsIDType = "be_ent"
	ChildLegalEntityIdentificationsIDTypeBeNrn                ChildLegalEntityIdentificationsIDType = "be_nrn"
	ChildLegalEntityIdentificationsIDTypeBrCnpj               ChildLegalEntityIdentificationsIDType = "br_cnpj"
	ChildLegalEntityIdentificationsIDTypeBrCpf                ChildLegalEntityIdentificationsIDType = "br_cpf"
	ChildLegalEntityIdentificationsIDTypeCaBn                 ChildLegalEntityIdentificationsIDType = "ca_bn"
	ChildLegalEntityIdentificationsIDTypeCaSin                ChildLegalEntityIdentificationsIDType = "ca_sin"
	ChildLegalEntityIdentificationsIDTypeChAhv                ChildLegalEntityIdentificationsIDType = "ch_ahv"
	ChildLegalEntityIdentificationsIDTypeChUid                ChildLegalEntityIdentificationsIDType = "ch_uid"
	ChildLegalEntityIdentificationsIDTypeClRun                ChildLegalEntityIdentificationsIDType = "cl_run"
	ChildLegalEntityIdentificationsIDTypeClRut                ChildLegalEntityIdentificationsIDType = "cl_rut"
	ChildLegalEntityIdentificationsIDTypeCoCedulas            ChildLegalEntityIdentificationsIDType = "co_cedulas"
	ChildLegalEntityIdentificationsIDTypeCoNit                ChildLegalEntityIdentificationsIDType = "co_nit"
	ChildLegalEntityIdentificationsIDTypeCyTin                ChildLegalEntityIdentificationsIDType = "cy_tin"
	ChildLegalEntityIdentificationsIDTypeCzIco                ChildLegalEntityIdentificationsIDType = "cz_ico"
	ChildLegalEntityIdentificationsIDTypeCzRc                 ChildLegalEntityIdentificationsIDType = "cz_rc"
	ChildLegalEntityIdentificationsIDTypeDeStid               ChildLegalEntityIdentificationsIDType = "de_stid"
	ChildLegalEntityIdentificationsIDTypeDeStnr               ChildLegalEntityIdentificationsIDType = "de_stnr"
	ChildLegalEntityIdentificationsIDTypeDeVat                ChildLegalEntityIdentificationsIDType = "de_vat"
	ChildLegalEntityIdentificationsIDTypeDkCpr                ChildLegalEntityIdentificationsIDType = "dk_cpr"
	ChildLegalEntityIdentificationsIDTypeDkCvr                ChildLegalEntityIdentificationsIDType = "dk_cvr"
	ChildLegalEntityIdentificationsIDTypeDriversLicense       ChildLegalEntityIdentificationsIDType = "drivers_license"
	ChildLegalEntityIdentificationsIDTypeEeIk                 ChildLegalEntityIdentificationsIDType = "ee_ik"
	ChildLegalEntityIdentificationsIDTypeEeRk                 ChildLegalEntityIdentificationsIDType = "ee_rk"
	ChildLegalEntityIdentificationsIDTypeEsNie                ChildLegalEntityIdentificationsIDType = "es_nie"
	ChildLegalEntityIdentificationsIDTypeEsNif                ChildLegalEntityIdentificationsIDType = "es_nif"
	ChildLegalEntityIdentificationsIDTypeFiHetu               ChildLegalEntityIdentificationsIDType = "fi_hetu"
	ChildLegalEntityIdentificationsIDTypeFiYtj                ChildLegalEntityIdentificationsIDType = "fi_ytj"
	ChildLegalEntityIdentificationsIDTypeFrNif                ChildLegalEntityIdentificationsIDType = "fr_nif"
	ChildLegalEntityIdentificationsIDTypeFrSiren              ChildLegalEntityIdentificationsIDType = "fr_siren"
	ChildLegalEntityIdentificationsIDTypeFrVat                ChildLegalEntityIdentificationsIDType = "fr_vat"
	ChildLegalEntityIdentificationsIDTypeGBNino               ChildLegalEntityIdentificationsIDType = "gb_nino"
	ChildLegalEntityIdentificationsIDTypeGBUtr                ChildLegalEntityIdentificationsIDType = "gb_utr"
	ChildLegalEntityIdentificationsIDTypeGBVat                ChildLegalEntityIdentificationsIDType = "gb_vat"
	ChildLegalEntityIdentificationsIDTypeGenericInternational ChildLegalEntityIdentificationsIDType = "generic_international"
	ChildLegalEntityIdentificationsIDTypeGrVat                ChildLegalEntityIdentificationsIDType = "gr_vat"
	ChildLegalEntityIdentificationsIDTypeHkBrn                ChildLegalEntityIdentificationsIDType = "hk_brn"
	ChildLegalEntityIdentificationsIDTypeHkHkid               ChildLegalEntityIdentificationsIDType = "hk_hkid"
	ChildLegalEntityIdentificationsIDTypeHnID                 ChildLegalEntityIdentificationsIDType = "hn_id"
	ChildLegalEntityIdentificationsIDTypeHnRtn                ChildLegalEntityIdentificationsIDType = "hn_rtn"
	ChildLegalEntityIdentificationsIDTypeHrOib                ChildLegalEntityIdentificationsIDType = "hr_oib"
	ChildLegalEntityIdentificationsIDTypeHuAdj                ChildLegalEntityIdentificationsIDType = "hu_adj"
	ChildLegalEntityIdentificationsIDTypeHuAnum               ChildLegalEntityIdentificationsIDType = "hu_anum"
	ChildLegalEntityIdentificationsIDTypeIePps                ChildLegalEntityIdentificationsIDType = "ie_pps"
	ChildLegalEntityIdentificationsIDTypeIeTrn                ChildLegalEntityIdentificationsIDType = "ie_trn"
	ChildLegalEntityIdentificationsIDTypeInLei                ChildLegalEntityIdentificationsIDType = "in_lei"
	ChildLegalEntityIdentificationsIDTypeIsKnt                ChildLegalEntityIdentificationsIDType = "is_knt"
	ChildLegalEntityIdentificationsIDTypeItCf                 ChildLegalEntityIdentificationsIDType = "it_cf"
	ChildLegalEntityIdentificationsIDTypeItPiva               ChildLegalEntityIdentificationsIDType = "it_piva"
	ChildLegalEntityIdentificationsIDTypeJpHb                 ChildLegalEntityIdentificationsIDType = "jp_hb"
	ChildLegalEntityIdentificationsIDTypeJpMn                 ChildLegalEntityIdentificationsIDType = "jp_mn"
	ChildLegalEntityIdentificationsIDTypeKrBrn                ChildLegalEntityIdentificationsIDType = "kr_brn"
	ChildLegalEntityIdentificationsIDTypeKrCrn                ChildLegalEntityIdentificationsIDType = "kr_crn"
	ChildLegalEntityIdentificationsIDTypeKrRrn                ChildLegalEntityIdentificationsIDType = "kr_rrn"
	ChildLegalEntityIdentificationsIDTypeLiPeid               ChildLegalEntityIdentificationsIDType = "li_peid"
	ChildLegalEntityIdentificationsIDTypeLtAk                 ChildLegalEntityIdentificationsIDType = "lt_ak"
	ChildLegalEntityIdentificationsIDTypeLtJak                ChildLegalEntityIdentificationsIDType = "lt_jak"
	ChildLegalEntityIdentificationsIDTypeLuMtc                ChildLegalEntityIdentificationsIDType = "lu_mtc"
	ChildLegalEntityIdentificationsIDTypeLuVat                ChildLegalEntityIdentificationsIDType = "lu_vat"
	ChildLegalEntityIdentificationsIDTypeLvPk                 ChildLegalEntityIdentificationsIDType = "lv_pk"
	ChildLegalEntityIdentificationsIDTypeLvRn                 ChildLegalEntityIdentificationsIDType = "lv_rn"
	ChildLegalEntityIdentificationsIDTypeMtTin                ChildLegalEntityIdentificationsIDType = "mt_tin"
	ChildLegalEntityIdentificationsIDTypeMtVat                ChildLegalEntityIdentificationsIDType = "mt_vat"
	ChildLegalEntityIdentificationsIDTypeMxCurp               ChildLegalEntityIdentificationsIDType = "mx_curp"
	ChildLegalEntityIdentificationsIDTypeMxIne                ChildLegalEntityIdentificationsIDType = "mx_ine"
	ChildLegalEntityIdentificationsIDTypeMxRfc                ChildLegalEntityIdentificationsIDType = "mx_rfc"
	ChildLegalEntityIdentificationsIDTypeNationalID           ChildLegalEntityIdentificationsIDType = "national_id"
	ChildLegalEntityIdentificationsIDTypeNlBsn                ChildLegalEntityIdentificationsIDType = "nl_bsn"
	ChildLegalEntityIdentificationsIDTypeNlBtw                ChildLegalEntityIdentificationsIDType = "nl_btw"
	ChildLegalEntityIdentificationsIDTypeNlRsin               ChildLegalEntityIdentificationsIDType = "nl_rsin"
	ChildLegalEntityIdentificationsIDTypeNoFdn                ChildLegalEntityIdentificationsIDType = "no_fdn"
	ChildLegalEntityIdentificationsIDTypeNoMva                ChildLegalEntityIdentificationsIDType = "no_mva"
	ChildLegalEntityIdentificationsIDTypeNoOrgnr              ChildLegalEntityIdentificationsIDType = "no_orgnr"
	ChildLegalEntityIdentificationsIDTypeNzIrd                ChildLegalEntityIdentificationsIDType = "nz_ird"
	ChildLegalEntityIdentificationsIDTypePassport             ChildLegalEntityIdentificationsIDType = "passport"
	ChildLegalEntityIdentificationsIDTypePlNip                ChildLegalEntityIdentificationsIDType = "pl_nip"
	ChildLegalEntityIdentificationsIDTypePlPesel              ChildLegalEntityIdentificationsIDType = "pl_pesel"
	ChildLegalEntityIdentificationsIDTypePtNif                ChildLegalEntityIdentificationsIDType = "pt_nif"
	ChildLegalEntityIdentificationsIDTypeRoCnp                ChildLegalEntityIdentificationsIDType = "ro_cnp"
	ChildLegalEntityIdentificationsIDTypeRoCui                ChildLegalEntityIdentificationsIDType = "ro_cui"
	ChildLegalEntityIdentificationsIDTypeSaTin                ChildLegalEntityIdentificationsIDType = "sa_tin"
	ChildLegalEntityIdentificationsIDTypeSaVat                ChildLegalEntityIdentificationsIDType = "sa_vat"
	ChildLegalEntityIdentificationsIDTypeSeOrgnr              ChildLegalEntityIdentificationsIDType = "se_orgnr"
	ChildLegalEntityIdentificationsIDTypeSePnmr               ChildLegalEntityIdentificationsIDType = "se_pnmr"
	ChildLegalEntityIdentificationsIDTypeSgFin                ChildLegalEntityIdentificationsIDType = "sg_fin"
	ChildLegalEntityIdentificationsIDTypeSgNric               ChildLegalEntityIdentificationsIDType = "sg_nric"
	ChildLegalEntityIdentificationsIDTypeSgUen                ChildLegalEntityIdentificationsIDType = "sg_uen"
	ChildLegalEntityIdentificationsIDTypeSiDav                ChildLegalEntityIdentificationsIDType = "si_dav"
	ChildLegalEntityIdentificationsIDTypeSiTin                ChildLegalEntityIdentificationsIDType = "si_tin"
	ChildLegalEntityIdentificationsIDTypeSkIco                ChildLegalEntityIdentificationsIDType = "sk_ico"
	ChildLegalEntityIdentificationsIDTypeSkRc                 ChildLegalEntityIdentificationsIDType = "sk_rc"
	ChildLegalEntityIdentificationsIDTypeUsEin                ChildLegalEntityIdentificationsIDType = "us_ein"
	ChildLegalEntityIdentificationsIDTypeUsItin               ChildLegalEntityIdentificationsIDType = "us_itin"
	ChildLegalEntityIdentificationsIDTypeUsSsn                ChildLegalEntityIdentificationsIDType = "us_ssn"
	ChildLegalEntityIdentificationsIDTypeUyRut                ChildLegalEntityIdentificationsIDType = "uy_rut"
	ChildLegalEntityIdentificationsIDTypeVnTin                ChildLegalEntityIdentificationsIDType = "vn_tin"
)

func (r ChildLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case ChildLegalEntityIdentificationsIDTypeArCuil, ChildLegalEntityIdentificationsIDTypeArCuit, ChildLegalEntityIdentificationsIDTypeAtAtin, ChildLegalEntityIdentificationsIDTypeAtVat, ChildLegalEntityIdentificationsIDTypeAuAbn, ChildLegalEntityIdentificationsIDTypeAuTfn, ChildLegalEntityIdentificationsIDTypeBeEnt, ChildLegalEntityIdentificationsIDTypeBeNrn, ChildLegalEntityIdentificationsIDTypeBrCnpj, ChildLegalEntityIdentificationsIDTypeBrCpf, ChildLegalEntityIdentificationsIDTypeCaBn, ChildLegalEntityIdentificationsIDTypeCaSin, ChildLegalEntityIdentificationsIDTypeChAhv, ChildLegalEntityIdentificationsIDTypeChUid, ChildLegalEntityIdentificationsIDTypeClRun, ChildLegalEntityIdentificationsIDTypeClRut, ChildLegalEntityIdentificationsIDTypeCoCedulas, ChildLegalEntityIdentificationsIDTypeCoNit, ChildLegalEntityIdentificationsIDTypeCyTin, ChildLegalEntityIdentificationsIDTypeCzIco, ChildLegalEntityIdentificationsIDTypeCzRc, ChildLegalEntityIdentificationsIDTypeDeStid, ChildLegalEntityIdentificationsIDTypeDeStnr, ChildLegalEntityIdentificationsIDTypeDeVat, ChildLegalEntityIdentificationsIDTypeDkCpr, ChildLegalEntityIdentificationsIDTypeDkCvr, ChildLegalEntityIdentificationsIDTypeDriversLicense, ChildLegalEntityIdentificationsIDTypeEeIk, ChildLegalEntityIdentificationsIDTypeEeRk, ChildLegalEntityIdentificationsIDTypeEsNie, ChildLegalEntityIdentificationsIDTypeEsNif, ChildLegalEntityIdentificationsIDTypeFiHetu, ChildLegalEntityIdentificationsIDTypeFiYtj, ChildLegalEntityIdentificationsIDTypeFrNif, ChildLegalEntityIdentificationsIDTypeFrSiren, ChildLegalEntityIdentificationsIDTypeFrVat, ChildLegalEntityIdentificationsIDTypeGBNino, ChildLegalEntityIdentificationsIDTypeGBUtr, ChildLegalEntityIdentificationsIDTypeGBVat, ChildLegalEntityIdentificationsIDTypeGenericInternational, ChildLegalEntityIdentificationsIDTypeGrVat, ChildLegalEntityIdentificationsIDTypeHkBrn, ChildLegalEntityIdentificationsIDTypeHkHkid, ChildLegalEntityIdentificationsIDTypeHnID, ChildLegalEntityIdentificationsIDTypeHnRtn, ChildLegalEntityIdentificationsIDTypeHrOib, ChildLegalEntityIdentificationsIDTypeHuAdj, ChildLegalEntityIdentificationsIDTypeHuAnum, ChildLegalEntityIdentificationsIDTypeIePps, ChildLegalEntityIdentificationsIDTypeIeTrn, ChildLegalEntityIdentificationsIDTypeInLei, ChildLegalEntityIdentificationsIDTypeIsKnt, ChildLegalEntityIdentificationsIDTypeItCf, ChildLegalEntityIdentificationsIDTypeItPiva, ChildLegalEntityIdentificationsIDTypeJpHb, ChildLegalEntityIdentificationsIDTypeJpMn, ChildLegalEntityIdentificationsIDTypeKrBrn, ChildLegalEntityIdentificationsIDTypeKrCrn, ChildLegalEntityIdentificationsIDTypeKrRrn, ChildLegalEntityIdentificationsIDTypeLiPeid, ChildLegalEntityIdentificationsIDTypeLtAk, ChildLegalEntityIdentificationsIDTypeLtJak, ChildLegalEntityIdentificationsIDTypeLuMtc, ChildLegalEntityIdentificationsIDTypeLuVat, ChildLegalEntityIdentificationsIDTypeLvPk, ChildLegalEntityIdentificationsIDTypeLvRn, ChildLegalEntityIdentificationsIDTypeMtTin, ChildLegalEntityIdentificationsIDTypeMtVat, ChildLegalEntityIdentificationsIDTypeMxCurp, ChildLegalEntityIdentificationsIDTypeMxIne, ChildLegalEntityIdentificationsIDTypeMxRfc, ChildLegalEntityIdentificationsIDTypeNationalID, ChildLegalEntityIdentificationsIDTypeNlBsn, ChildLegalEntityIdentificationsIDTypeNlBtw, ChildLegalEntityIdentificationsIDTypeNlRsin, ChildLegalEntityIdentificationsIDTypeNoFdn, ChildLegalEntityIdentificationsIDTypeNoMva, ChildLegalEntityIdentificationsIDTypeNoOrgnr, ChildLegalEntityIdentificationsIDTypeNzIrd, ChildLegalEntityIdentificationsIDTypePassport, ChildLegalEntityIdentificationsIDTypePlNip, ChildLegalEntityIdentificationsIDTypePlPesel, ChildLegalEntityIdentificationsIDTypePtNif, ChildLegalEntityIdentificationsIDTypeRoCnp, ChildLegalEntityIdentificationsIDTypeRoCui, ChildLegalEntityIdentificationsIDTypeSaTin, ChildLegalEntityIdentificationsIDTypeSaVat, ChildLegalEntityIdentificationsIDTypeSeOrgnr, ChildLegalEntityIdentificationsIDTypeSePnmr, ChildLegalEntityIdentificationsIDTypeSgFin, ChildLegalEntityIdentificationsIDTypeSgNric, ChildLegalEntityIdentificationsIDTypeSgUen, ChildLegalEntityIdentificationsIDTypeSiDav, ChildLegalEntityIdentificationsIDTypeSiTin, ChildLegalEntityIdentificationsIDTypeSkIco, ChildLegalEntityIdentificationsIDTypeSkRc, ChildLegalEntityIdentificationsIDTypeUsEin, ChildLegalEntityIdentificationsIDTypeUsItin, ChildLegalEntityIdentificationsIDTypeUsSsn, ChildLegalEntityIdentificationsIDTypeUyRut, ChildLegalEntityIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

// The type of legal entity.
type ChildLegalEntityLegalEntityType string

const (
	ChildLegalEntityLegalEntityTypeBusiness   ChildLegalEntityLegalEntityType = "business"
	ChildLegalEntityLegalEntityTypeIndividual ChildLegalEntityLegalEntityType = "individual"
)

func (r ChildLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case ChildLegalEntityLegalEntityTypeBusiness, ChildLegalEntityLegalEntityTypeIndividual:
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
	// A phone number in E.164 format. This format is strictly validated: include a
	// leading + and country code, followed by digits only (no spaces or dashes), e.g.
	// +12025551234.
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
	Jurisdiction string `json:"jurisdiction" api:"required"`
	// Full name of the regulatory body.
	Name string `json:"name" api:"required"`
	// Registration or identification number with the regulator.
	RegistrationNumber string                        `json:"registration_number" api:"required"`
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

// The activation status of the legal entity. One of pending, active, suspended, or
// denied.
type ChildLegalEntityStatus string

const (
	ChildLegalEntityStatusActive    ChildLegalEntityStatus = "active"
	ChildLegalEntityStatusDenied    ChildLegalEntityStatus = "denied"
	ChildLegalEntityStatusPending   ChildLegalEntityStatus = "pending"
	ChildLegalEntityStatusSuspended ChildLegalEntityStatus = "suspended"
)

func (r ChildLegalEntityStatus) IsKnown() bool {
	switch r {
	case ChildLegalEntityStatusActive, ChildLegalEntityStatusDenied, ChildLegalEntityStatusPending, ChildLegalEntityStatusSuspended:
		return true
	}
	return false
}

// Acceptance of terms of use by the legal entity.
type ChildLegalEntityTermsOfUse struct {
	// The ISO 8601 timestamp indicating when the terms of use were accepted.
	AcceptedAt time.Time `json:"accepted_at" format:"date-time"`
	// The IP address from which the terms of use were accepted. Supports both IPv4 and
	// IPv6 formats.
	IPAddress string                         `json:"ip_address"`
	JSON      childLegalEntityTermsOfUseJSON `json:"-"`
}

// childLegalEntityTermsOfUseJSON contains the JSON metadata for the struct
// [ChildLegalEntityTermsOfUse]
type childLegalEntityTermsOfUseJSON struct {
	AcceptedAt  apijson.Field
	IPAddress   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChildLegalEntityTermsOfUse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r childLegalEntityTermsOfUseJSON) RawJSON() string {
	return r.raw
}

type LegalEntityAssociation struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The child legal entity.
	ChildLegalEntity *ChildLegalEntity `json:"child_legal_entity" api:"required"`
	CreatedAt        time.Time         `json:"created_at" api:"required" format:"date-time"`
	DiscardedAt      time.Time         `json:"discarded_at" api:"required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage int64 `json:"ownership_percentage" api:"required,nullable"`
	// The ID of the parent legal entity. This must be a business legal entity.
	ParentLegalEntityID string                                   `json:"parent_legal_entity_id" api:"required"`
	RelationshipTypes   []LegalEntityAssociationRelationshipType `json:"relationship_types" api:"required"`
	// The job title of the child entity at the parent entity.
	Title     string                     `json:"title" api:"required,nullable"`
	UpdatedAt time.Time                  `json:"updated_at" api:"required" format:"date-time"`
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
