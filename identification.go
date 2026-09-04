// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// IdentificationService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentificationService] method instead.
type IdentificationService struct {
	Options []option.RequestOption
}

// NewIdentificationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentificationService(opts ...option.RequestOption) (r *IdentificationService) {
	r = &IdentificationService{}
	r.Options = opts
	return
}

// Create an Identification for a Legal Entity.
func (r *IdentificationService) New(ctx context.Context, body IdentificationNewParams, opts ...option.RequestOption) (res *Identification, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/identifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an existing Identification.
func (r *IdentificationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Identification, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/identifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing Identification.
func (r *IdentificationService) Update(ctx context.Context, id string, body IdentificationUpdateParams, opts ...option.RequestOption) (res *Identification, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/identifications/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type Identification struct {
	ID          string     `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time  `json:"created_at" api:"required" format:"date-time"`
	DiscardedAt time.Time  `json:"discarded_at" api:"required,nullable" format:"date-time"`
	Documents   []Document `json:"documents" api:"required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate time.Time `json:"expiration_date" api:"required,nullable" format:"date"`
	// The type of ID number.
	IDType IdentificationIDType `json:"id_type" api:"required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country" api:"required,nullable"`
	// The region in which the identifcation was issued.
	IssuingRegion string `json:"issuing_region" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool               `json:"live_mode" api:"required"`
	Object    string             `json:"object" api:"required"`
	UpdatedAt time.Time          `json:"updated_at" api:"required" format:"date-time"`
	JSON      identificationJSON `json:"-"`
}

// identificationJSON contains the JSON metadata for the struct [Identification]
type identificationJSON struct {
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

func (r *Identification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identificationJSON) RawJSON() string {
	return r.raw
}

// The type of ID number.
type IdentificationIDType string

const (
	IdentificationIDTypeArCuil               IdentificationIDType = "ar_cuil"
	IdentificationIDTypeArCuit               IdentificationIDType = "ar_cuit"
	IdentificationIDTypeAtAtin               IdentificationIDType = "at_atin"
	IdentificationIDTypeAtVat                IdentificationIDType = "at_vat"
	IdentificationIDTypeAuAbn                IdentificationIDType = "au_abn"
	IdentificationIDTypeAuTfn                IdentificationIDType = "au_tfn"
	IdentificationIDTypeBeEnt                IdentificationIDType = "be_ent"
	IdentificationIDTypeBeNrn                IdentificationIDType = "be_nrn"
	IdentificationIDTypeBrCnpj               IdentificationIDType = "br_cnpj"
	IdentificationIDTypeBrCpf                IdentificationIDType = "br_cpf"
	IdentificationIDTypeCaBn                 IdentificationIDType = "ca_bn"
	IdentificationIDTypeCaSin                IdentificationIDType = "ca_sin"
	IdentificationIDTypeChAhv                IdentificationIDType = "ch_ahv"
	IdentificationIDTypeChUid                IdentificationIDType = "ch_uid"
	IdentificationIDTypeClRun                IdentificationIDType = "cl_run"
	IdentificationIDTypeClRut                IdentificationIDType = "cl_rut"
	IdentificationIDTypeCoCedulas            IdentificationIDType = "co_cedulas"
	IdentificationIDTypeCoNit                IdentificationIDType = "co_nit"
	IdentificationIDTypeCyTin                IdentificationIDType = "cy_tin"
	IdentificationIDTypeCzIco                IdentificationIDType = "cz_ico"
	IdentificationIDTypeCzRc                 IdentificationIDType = "cz_rc"
	IdentificationIDTypeDeStid               IdentificationIDType = "de_stid"
	IdentificationIDTypeDeStnr               IdentificationIDType = "de_stnr"
	IdentificationIDTypeDeVat                IdentificationIDType = "de_vat"
	IdentificationIDTypeDkCpr                IdentificationIDType = "dk_cpr"
	IdentificationIDTypeDkCvr                IdentificationIDType = "dk_cvr"
	IdentificationIDTypeDriversLicense       IdentificationIDType = "drivers_license"
	IdentificationIDTypeEeIk                 IdentificationIDType = "ee_ik"
	IdentificationIDTypeEeRk                 IdentificationIDType = "ee_rk"
	IdentificationIDTypeEsNie                IdentificationIDType = "es_nie"
	IdentificationIDTypeEsNif                IdentificationIDType = "es_nif"
	IdentificationIDTypeFiHetu               IdentificationIDType = "fi_hetu"
	IdentificationIDTypeFiYtj                IdentificationIDType = "fi_ytj"
	IdentificationIDTypeFrNif                IdentificationIDType = "fr_nif"
	IdentificationIDTypeFrSiren              IdentificationIDType = "fr_siren"
	IdentificationIDTypeFrVat                IdentificationIDType = "fr_vat"
	IdentificationIDTypeGBNino               IdentificationIDType = "gb_nino"
	IdentificationIDTypeGBUtr                IdentificationIDType = "gb_utr"
	IdentificationIDTypeGBVat                IdentificationIDType = "gb_vat"
	IdentificationIDTypeGenericInternational IdentificationIDType = "generic_international"
	IdentificationIDTypeGrVat                IdentificationIDType = "gr_vat"
	IdentificationIDTypeHkBrn                IdentificationIDType = "hk_brn"
	IdentificationIDTypeHkHkid               IdentificationIDType = "hk_hkid"
	IdentificationIDTypeHnID                 IdentificationIDType = "hn_id"
	IdentificationIDTypeHnRtn                IdentificationIDType = "hn_rtn"
	IdentificationIDTypeHrOib                IdentificationIDType = "hr_oib"
	IdentificationIDTypeHuAdj                IdentificationIDType = "hu_adj"
	IdentificationIDTypeHuAnum               IdentificationIDType = "hu_anum"
	IdentificationIDTypeIePps                IdentificationIDType = "ie_pps"
	IdentificationIDTypeIeTrn                IdentificationIDType = "ie_trn"
	IdentificationIDTypeInLei                IdentificationIDType = "in_lei"
	IdentificationIDTypeIsKnt                IdentificationIDType = "is_knt"
	IdentificationIDTypeItCf                 IdentificationIDType = "it_cf"
	IdentificationIDTypeItPiva               IdentificationIDType = "it_piva"
	IdentificationIDTypeJpHb                 IdentificationIDType = "jp_hb"
	IdentificationIDTypeJpMn                 IdentificationIDType = "jp_mn"
	IdentificationIDTypeKrBrn                IdentificationIDType = "kr_brn"
	IdentificationIDTypeKrCrn                IdentificationIDType = "kr_crn"
	IdentificationIDTypeKrRrn                IdentificationIDType = "kr_rrn"
	IdentificationIDTypeLiPeid               IdentificationIDType = "li_peid"
	IdentificationIDTypeLtAk                 IdentificationIDType = "lt_ak"
	IdentificationIDTypeLtJak                IdentificationIDType = "lt_jak"
	IdentificationIDTypeLuMtc                IdentificationIDType = "lu_mtc"
	IdentificationIDTypeLuVat                IdentificationIDType = "lu_vat"
	IdentificationIDTypeLvPk                 IdentificationIDType = "lv_pk"
	IdentificationIDTypeLvRn                 IdentificationIDType = "lv_rn"
	IdentificationIDTypeMtTin                IdentificationIDType = "mt_tin"
	IdentificationIDTypeMtVat                IdentificationIDType = "mt_vat"
	IdentificationIDTypeMxCurp               IdentificationIDType = "mx_curp"
	IdentificationIDTypeMxIne                IdentificationIDType = "mx_ine"
	IdentificationIDTypeMxRfc                IdentificationIDType = "mx_rfc"
	IdentificationIDTypeNationalID           IdentificationIDType = "national_id"
	IdentificationIDTypeNlBsn                IdentificationIDType = "nl_bsn"
	IdentificationIDTypeNlBtw                IdentificationIDType = "nl_btw"
	IdentificationIDTypeNlRsin               IdentificationIDType = "nl_rsin"
	IdentificationIDTypeNoFdn                IdentificationIDType = "no_fdn"
	IdentificationIDTypeNoMva                IdentificationIDType = "no_mva"
	IdentificationIDTypeNoOrgnr              IdentificationIDType = "no_orgnr"
	IdentificationIDTypeNzIrd                IdentificationIDType = "nz_ird"
	IdentificationIDTypePassport             IdentificationIDType = "passport"
	IdentificationIDTypePlNip                IdentificationIDType = "pl_nip"
	IdentificationIDTypePlPesel              IdentificationIDType = "pl_pesel"
	IdentificationIDTypePtNif                IdentificationIDType = "pt_nif"
	IdentificationIDTypeRoCnp                IdentificationIDType = "ro_cnp"
	IdentificationIDTypeRoCui                IdentificationIDType = "ro_cui"
	IdentificationIDTypeSaTin                IdentificationIDType = "sa_tin"
	IdentificationIDTypeSaVat                IdentificationIDType = "sa_vat"
	IdentificationIDTypeSeOrgnr              IdentificationIDType = "se_orgnr"
	IdentificationIDTypeSePnmr               IdentificationIDType = "se_pnmr"
	IdentificationIDTypeSgFin                IdentificationIDType = "sg_fin"
	IdentificationIDTypeSgNric               IdentificationIDType = "sg_nric"
	IdentificationIDTypeSgUen                IdentificationIDType = "sg_uen"
	IdentificationIDTypeSiDav                IdentificationIDType = "si_dav"
	IdentificationIDTypeSiTin                IdentificationIDType = "si_tin"
	IdentificationIDTypeSkIco                IdentificationIDType = "sk_ico"
	IdentificationIDTypeSkRc                 IdentificationIDType = "sk_rc"
	IdentificationIDTypeUsEin                IdentificationIDType = "us_ein"
	IdentificationIDTypeUsItin               IdentificationIDType = "us_itin"
	IdentificationIDTypeUsSsn                IdentificationIDType = "us_ssn"
	IdentificationIDTypeUyRut                IdentificationIDType = "uy_rut"
	IdentificationIDTypeVnTin                IdentificationIDType = "vn_tin"
)

func (r IdentificationIDType) IsKnown() bool {
	switch r {
	case IdentificationIDTypeArCuil, IdentificationIDTypeArCuit, IdentificationIDTypeAtAtin, IdentificationIDTypeAtVat, IdentificationIDTypeAuAbn, IdentificationIDTypeAuTfn, IdentificationIDTypeBeEnt, IdentificationIDTypeBeNrn, IdentificationIDTypeBrCnpj, IdentificationIDTypeBrCpf, IdentificationIDTypeCaBn, IdentificationIDTypeCaSin, IdentificationIDTypeChAhv, IdentificationIDTypeChUid, IdentificationIDTypeClRun, IdentificationIDTypeClRut, IdentificationIDTypeCoCedulas, IdentificationIDTypeCoNit, IdentificationIDTypeCyTin, IdentificationIDTypeCzIco, IdentificationIDTypeCzRc, IdentificationIDTypeDeStid, IdentificationIDTypeDeStnr, IdentificationIDTypeDeVat, IdentificationIDTypeDkCpr, IdentificationIDTypeDkCvr, IdentificationIDTypeDriversLicense, IdentificationIDTypeEeIk, IdentificationIDTypeEeRk, IdentificationIDTypeEsNie, IdentificationIDTypeEsNif, IdentificationIDTypeFiHetu, IdentificationIDTypeFiYtj, IdentificationIDTypeFrNif, IdentificationIDTypeFrSiren, IdentificationIDTypeFrVat, IdentificationIDTypeGBNino, IdentificationIDTypeGBUtr, IdentificationIDTypeGBVat, IdentificationIDTypeGenericInternational, IdentificationIDTypeGrVat, IdentificationIDTypeHkBrn, IdentificationIDTypeHkHkid, IdentificationIDTypeHnID, IdentificationIDTypeHnRtn, IdentificationIDTypeHrOib, IdentificationIDTypeHuAdj, IdentificationIDTypeHuAnum, IdentificationIDTypeIePps, IdentificationIDTypeIeTrn, IdentificationIDTypeInLei, IdentificationIDTypeIsKnt, IdentificationIDTypeItCf, IdentificationIDTypeItPiva, IdentificationIDTypeJpHb, IdentificationIDTypeJpMn, IdentificationIDTypeKrBrn, IdentificationIDTypeKrCrn, IdentificationIDTypeKrRrn, IdentificationIDTypeLiPeid, IdentificationIDTypeLtAk, IdentificationIDTypeLtJak, IdentificationIDTypeLuMtc, IdentificationIDTypeLuVat, IdentificationIDTypeLvPk, IdentificationIDTypeLvRn, IdentificationIDTypeMtTin, IdentificationIDTypeMtVat, IdentificationIDTypeMxCurp, IdentificationIDTypeMxIne, IdentificationIDTypeMxRfc, IdentificationIDTypeNationalID, IdentificationIDTypeNlBsn, IdentificationIDTypeNlBtw, IdentificationIDTypeNlRsin, IdentificationIDTypeNoFdn, IdentificationIDTypeNoMva, IdentificationIDTypeNoOrgnr, IdentificationIDTypeNzIrd, IdentificationIDTypePassport, IdentificationIDTypePlNip, IdentificationIDTypePlPesel, IdentificationIDTypePtNif, IdentificationIDTypeRoCnp, IdentificationIDTypeRoCui, IdentificationIDTypeSaTin, IdentificationIDTypeSaVat, IdentificationIDTypeSeOrgnr, IdentificationIDTypeSePnmr, IdentificationIDTypeSgFin, IdentificationIDTypeSgNric, IdentificationIDTypeSgUen, IdentificationIDTypeSiDav, IdentificationIDTypeSiTin, IdentificationIDTypeSkIco, IdentificationIDTypeSkRc, IdentificationIDTypeUsEin, IdentificationIDTypeUsItin, IdentificationIDTypeUsSsn, IdentificationIDTypeUyRut, IdentificationIDTypeVnTin:
		return true
	}
	return false
}

type IdentificationNewParams struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number" api:"required"`
	// The type of ID number.
	IDType param.Field[IdentificationNewParamsIDType] `json:"id_type" api:"required"`
	// The ID of the Legal Entity the identification belongs to.
	LegalEntityID param.Field[string] `json:"legal_entity_id" api:"required"`
	// A list of documents to attach to the identification.
	Documents param.Field[[]IdentificationNewParamsDocument] `json:"documents"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
	// The region in which the identifcation was issued.
	IssuingRegion param.Field[string] `json:"issuing_region"`
}

func (r IdentificationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type IdentificationNewParamsIDType string

const (
	IdentificationNewParamsIDTypeArCuil               IdentificationNewParamsIDType = "ar_cuil"
	IdentificationNewParamsIDTypeArCuit               IdentificationNewParamsIDType = "ar_cuit"
	IdentificationNewParamsIDTypeAtAtin               IdentificationNewParamsIDType = "at_atin"
	IdentificationNewParamsIDTypeAtVat                IdentificationNewParamsIDType = "at_vat"
	IdentificationNewParamsIDTypeAuAbn                IdentificationNewParamsIDType = "au_abn"
	IdentificationNewParamsIDTypeAuTfn                IdentificationNewParamsIDType = "au_tfn"
	IdentificationNewParamsIDTypeBeEnt                IdentificationNewParamsIDType = "be_ent"
	IdentificationNewParamsIDTypeBeNrn                IdentificationNewParamsIDType = "be_nrn"
	IdentificationNewParamsIDTypeBrCnpj               IdentificationNewParamsIDType = "br_cnpj"
	IdentificationNewParamsIDTypeBrCpf                IdentificationNewParamsIDType = "br_cpf"
	IdentificationNewParamsIDTypeCaBn                 IdentificationNewParamsIDType = "ca_bn"
	IdentificationNewParamsIDTypeCaSin                IdentificationNewParamsIDType = "ca_sin"
	IdentificationNewParamsIDTypeChAhv                IdentificationNewParamsIDType = "ch_ahv"
	IdentificationNewParamsIDTypeChUid                IdentificationNewParamsIDType = "ch_uid"
	IdentificationNewParamsIDTypeClRun                IdentificationNewParamsIDType = "cl_run"
	IdentificationNewParamsIDTypeClRut                IdentificationNewParamsIDType = "cl_rut"
	IdentificationNewParamsIDTypeCoCedulas            IdentificationNewParamsIDType = "co_cedulas"
	IdentificationNewParamsIDTypeCoNit                IdentificationNewParamsIDType = "co_nit"
	IdentificationNewParamsIDTypeCyTin                IdentificationNewParamsIDType = "cy_tin"
	IdentificationNewParamsIDTypeCzIco                IdentificationNewParamsIDType = "cz_ico"
	IdentificationNewParamsIDTypeCzRc                 IdentificationNewParamsIDType = "cz_rc"
	IdentificationNewParamsIDTypeDeStid               IdentificationNewParamsIDType = "de_stid"
	IdentificationNewParamsIDTypeDeStnr               IdentificationNewParamsIDType = "de_stnr"
	IdentificationNewParamsIDTypeDeVat                IdentificationNewParamsIDType = "de_vat"
	IdentificationNewParamsIDTypeDkCpr                IdentificationNewParamsIDType = "dk_cpr"
	IdentificationNewParamsIDTypeDkCvr                IdentificationNewParamsIDType = "dk_cvr"
	IdentificationNewParamsIDTypeDriversLicense       IdentificationNewParamsIDType = "drivers_license"
	IdentificationNewParamsIDTypeEeIk                 IdentificationNewParamsIDType = "ee_ik"
	IdentificationNewParamsIDTypeEeRk                 IdentificationNewParamsIDType = "ee_rk"
	IdentificationNewParamsIDTypeEsNie                IdentificationNewParamsIDType = "es_nie"
	IdentificationNewParamsIDTypeEsNif                IdentificationNewParamsIDType = "es_nif"
	IdentificationNewParamsIDTypeFiHetu               IdentificationNewParamsIDType = "fi_hetu"
	IdentificationNewParamsIDTypeFiYtj                IdentificationNewParamsIDType = "fi_ytj"
	IdentificationNewParamsIDTypeFrNif                IdentificationNewParamsIDType = "fr_nif"
	IdentificationNewParamsIDTypeFrSiren              IdentificationNewParamsIDType = "fr_siren"
	IdentificationNewParamsIDTypeFrVat                IdentificationNewParamsIDType = "fr_vat"
	IdentificationNewParamsIDTypeGBNino               IdentificationNewParamsIDType = "gb_nino"
	IdentificationNewParamsIDTypeGBUtr                IdentificationNewParamsIDType = "gb_utr"
	IdentificationNewParamsIDTypeGBVat                IdentificationNewParamsIDType = "gb_vat"
	IdentificationNewParamsIDTypeGenericInternational IdentificationNewParamsIDType = "generic_international"
	IdentificationNewParamsIDTypeGrVat                IdentificationNewParamsIDType = "gr_vat"
	IdentificationNewParamsIDTypeHkBrn                IdentificationNewParamsIDType = "hk_brn"
	IdentificationNewParamsIDTypeHkHkid               IdentificationNewParamsIDType = "hk_hkid"
	IdentificationNewParamsIDTypeHnID                 IdentificationNewParamsIDType = "hn_id"
	IdentificationNewParamsIDTypeHnRtn                IdentificationNewParamsIDType = "hn_rtn"
	IdentificationNewParamsIDTypeHrOib                IdentificationNewParamsIDType = "hr_oib"
	IdentificationNewParamsIDTypeHuAdj                IdentificationNewParamsIDType = "hu_adj"
	IdentificationNewParamsIDTypeHuAnum               IdentificationNewParamsIDType = "hu_anum"
	IdentificationNewParamsIDTypeIePps                IdentificationNewParamsIDType = "ie_pps"
	IdentificationNewParamsIDTypeIeTrn                IdentificationNewParamsIDType = "ie_trn"
	IdentificationNewParamsIDTypeInLei                IdentificationNewParamsIDType = "in_lei"
	IdentificationNewParamsIDTypeIsKnt                IdentificationNewParamsIDType = "is_knt"
	IdentificationNewParamsIDTypeItCf                 IdentificationNewParamsIDType = "it_cf"
	IdentificationNewParamsIDTypeItPiva               IdentificationNewParamsIDType = "it_piva"
	IdentificationNewParamsIDTypeJpHb                 IdentificationNewParamsIDType = "jp_hb"
	IdentificationNewParamsIDTypeJpMn                 IdentificationNewParamsIDType = "jp_mn"
	IdentificationNewParamsIDTypeKrBrn                IdentificationNewParamsIDType = "kr_brn"
	IdentificationNewParamsIDTypeKrCrn                IdentificationNewParamsIDType = "kr_crn"
	IdentificationNewParamsIDTypeKrRrn                IdentificationNewParamsIDType = "kr_rrn"
	IdentificationNewParamsIDTypeLiPeid               IdentificationNewParamsIDType = "li_peid"
	IdentificationNewParamsIDTypeLtAk                 IdentificationNewParamsIDType = "lt_ak"
	IdentificationNewParamsIDTypeLtJak                IdentificationNewParamsIDType = "lt_jak"
	IdentificationNewParamsIDTypeLuMtc                IdentificationNewParamsIDType = "lu_mtc"
	IdentificationNewParamsIDTypeLuVat                IdentificationNewParamsIDType = "lu_vat"
	IdentificationNewParamsIDTypeLvPk                 IdentificationNewParamsIDType = "lv_pk"
	IdentificationNewParamsIDTypeLvRn                 IdentificationNewParamsIDType = "lv_rn"
	IdentificationNewParamsIDTypeMtTin                IdentificationNewParamsIDType = "mt_tin"
	IdentificationNewParamsIDTypeMtVat                IdentificationNewParamsIDType = "mt_vat"
	IdentificationNewParamsIDTypeMxCurp               IdentificationNewParamsIDType = "mx_curp"
	IdentificationNewParamsIDTypeMxIne                IdentificationNewParamsIDType = "mx_ine"
	IdentificationNewParamsIDTypeMxRfc                IdentificationNewParamsIDType = "mx_rfc"
	IdentificationNewParamsIDTypeNationalID           IdentificationNewParamsIDType = "national_id"
	IdentificationNewParamsIDTypeNlBsn                IdentificationNewParamsIDType = "nl_bsn"
	IdentificationNewParamsIDTypeNlBtw                IdentificationNewParamsIDType = "nl_btw"
	IdentificationNewParamsIDTypeNlRsin               IdentificationNewParamsIDType = "nl_rsin"
	IdentificationNewParamsIDTypeNoFdn                IdentificationNewParamsIDType = "no_fdn"
	IdentificationNewParamsIDTypeNoMva                IdentificationNewParamsIDType = "no_mva"
	IdentificationNewParamsIDTypeNoOrgnr              IdentificationNewParamsIDType = "no_orgnr"
	IdentificationNewParamsIDTypeNzIrd                IdentificationNewParamsIDType = "nz_ird"
	IdentificationNewParamsIDTypePassport             IdentificationNewParamsIDType = "passport"
	IdentificationNewParamsIDTypePlNip                IdentificationNewParamsIDType = "pl_nip"
	IdentificationNewParamsIDTypePlPesel              IdentificationNewParamsIDType = "pl_pesel"
	IdentificationNewParamsIDTypePtNif                IdentificationNewParamsIDType = "pt_nif"
	IdentificationNewParamsIDTypeRoCnp                IdentificationNewParamsIDType = "ro_cnp"
	IdentificationNewParamsIDTypeRoCui                IdentificationNewParamsIDType = "ro_cui"
	IdentificationNewParamsIDTypeSaTin                IdentificationNewParamsIDType = "sa_tin"
	IdentificationNewParamsIDTypeSaVat                IdentificationNewParamsIDType = "sa_vat"
	IdentificationNewParamsIDTypeSeOrgnr              IdentificationNewParamsIDType = "se_orgnr"
	IdentificationNewParamsIDTypeSePnmr               IdentificationNewParamsIDType = "se_pnmr"
	IdentificationNewParamsIDTypeSgFin                IdentificationNewParamsIDType = "sg_fin"
	IdentificationNewParamsIDTypeSgNric               IdentificationNewParamsIDType = "sg_nric"
	IdentificationNewParamsIDTypeSgUen                IdentificationNewParamsIDType = "sg_uen"
	IdentificationNewParamsIDTypeSiDav                IdentificationNewParamsIDType = "si_dav"
	IdentificationNewParamsIDTypeSiTin                IdentificationNewParamsIDType = "si_tin"
	IdentificationNewParamsIDTypeSkIco                IdentificationNewParamsIDType = "sk_ico"
	IdentificationNewParamsIDTypeSkRc                 IdentificationNewParamsIDType = "sk_rc"
	IdentificationNewParamsIDTypeUsEin                IdentificationNewParamsIDType = "us_ein"
	IdentificationNewParamsIDTypeUsItin               IdentificationNewParamsIDType = "us_itin"
	IdentificationNewParamsIDTypeUsSsn                IdentificationNewParamsIDType = "us_ssn"
	IdentificationNewParamsIDTypeUyRut                IdentificationNewParamsIDType = "uy_rut"
	IdentificationNewParamsIDTypeVnTin                IdentificationNewParamsIDType = "vn_tin"
)

func (r IdentificationNewParamsIDType) IsKnown() bool {
	switch r {
	case IdentificationNewParamsIDTypeArCuil, IdentificationNewParamsIDTypeArCuit, IdentificationNewParamsIDTypeAtAtin, IdentificationNewParamsIDTypeAtVat, IdentificationNewParamsIDTypeAuAbn, IdentificationNewParamsIDTypeAuTfn, IdentificationNewParamsIDTypeBeEnt, IdentificationNewParamsIDTypeBeNrn, IdentificationNewParamsIDTypeBrCnpj, IdentificationNewParamsIDTypeBrCpf, IdentificationNewParamsIDTypeCaBn, IdentificationNewParamsIDTypeCaSin, IdentificationNewParamsIDTypeChAhv, IdentificationNewParamsIDTypeChUid, IdentificationNewParamsIDTypeClRun, IdentificationNewParamsIDTypeClRut, IdentificationNewParamsIDTypeCoCedulas, IdentificationNewParamsIDTypeCoNit, IdentificationNewParamsIDTypeCyTin, IdentificationNewParamsIDTypeCzIco, IdentificationNewParamsIDTypeCzRc, IdentificationNewParamsIDTypeDeStid, IdentificationNewParamsIDTypeDeStnr, IdentificationNewParamsIDTypeDeVat, IdentificationNewParamsIDTypeDkCpr, IdentificationNewParamsIDTypeDkCvr, IdentificationNewParamsIDTypeDriversLicense, IdentificationNewParamsIDTypeEeIk, IdentificationNewParamsIDTypeEeRk, IdentificationNewParamsIDTypeEsNie, IdentificationNewParamsIDTypeEsNif, IdentificationNewParamsIDTypeFiHetu, IdentificationNewParamsIDTypeFiYtj, IdentificationNewParamsIDTypeFrNif, IdentificationNewParamsIDTypeFrSiren, IdentificationNewParamsIDTypeFrVat, IdentificationNewParamsIDTypeGBNino, IdentificationNewParamsIDTypeGBUtr, IdentificationNewParamsIDTypeGBVat, IdentificationNewParamsIDTypeGenericInternational, IdentificationNewParamsIDTypeGrVat, IdentificationNewParamsIDTypeHkBrn, IdentificationNewParamsIDTypeHkHkid, IdentificationNewParamsIDTypeHnID, IdentificationNewParamsIDTypeHnRtn, IdentificationNewParamsIDTypeHrOib, IdentificationNewParamsIDTypeHuAdj, IdentificationNewParamsIDTypeHuAnum, IdentificationNewParamsIDTypeIePps, IdentificationNewParamsIDTypeIeTrn, IdentificationNewParamsIDTypeInLei, IdentificationNewParamsIDTypeIsKnt, IdentificationNewParamsIDTypeItCf, IdentificationNewParamsIDTypeItPiva, IdentificationNewParamsIDTypeJpHb, IdentificationNewParamsIDTypeJpMn, IdentificationNewParamsIDTypeKrBrn, IdentificationNewParamsIDTypeKrCrn, IdentificationNewParamsIDTypeKrRrn, IdentificationNewParamsIDTypeLiPeid, IdentificationNewParamsIDTypeLtAk, IdentificationNewParamsIDTypeLtJak, IdentificationNewParamsIDTypeLuMtc, IdentificationNewParamsIDTypeLuVat, IdentificationNewParamsIDTypeLvPk, IdentificationNewParamsIDTypeLvRn, IdentificationNewParamsIDTypeMtTin, IdentificationNewParamsIDTypeMtVat, IdentificationNewParamsIDTypeMxCurp, IdentificationNewParamsIDTypeMxIne, IdentificationNewParamsIDTypeMxRfc, IdentificationNewParamsIDTypeNationalID, IdentificationNewParamsIDTypeNlBsn, IdentificationNewParamsIDTypeNlBtw, IdentificationNewParamsIDTypeNlRsin, IdentificationNewParamsIDTypeNoFdn, IdentificationNewParamsIDTypeNoMva, IdentificationNewParamsIDTypeNoOrgnr, IdentificationNewParamsIDTypeNzIrd, IdentificationNewParamsIDTypePassport, IdentificationNewParamsIDTypePlNip, IdentificationNewParamsIDTypePlPesel, IdentificationNewParamsIDTypePtNif, IdentificationNewParamsIDTypeRoCnp, IdentificationNewParamsIDTypeRoCui, IdentificationNewParamsIDTypeSaTin, IdentificationNewParamsIDTypeSaVat, IdentificationNewParamsIDTypeSeOrgnr, IdentificationNewParamsIDTypeSePnmr, IdentificationNewParamsIDTypeSgFin, IdentificationNewParamsIDTypeSgNric, IdentificationNewParamsIDTypeSgUen, IdentificationNewParamsIDTypeSiDav, IdentificationNewParamsIDTypeSiTin, IdentificationNewParamsIDTypeSkIco, IdentificationNewParamsIDTypeSkRc, IdentificationNewParamsIDTypeUsEin, IdentificationNewParamsIDTypeUsItin, IdentificationNewParamsIDTypeUsSsn, IdentificationNewParamsIDTypeUyRut, IdentificationNewParamsIDTypeVnTin:
		return true
	}
	return false
}

type IdentificationNewParamsDocument struct {
	// A category given to the document, can be `null`.
	DocumentType param.Field[IdentificationNewParamsDocumentsDocumentType] `json:"document_type" api:"required"`
	// Base64-encoded file content for the document.
	FileData param.Field[string] `json:"file_data" api:"required"`
	// The original filename of the document.
	Filename param.Field[string] `json:"filename"`
}

func (r IdentificationNewParamsDocument) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category given to the document, can be `null`.
type IdentificationNewParamsDocumentsDocumentType string

const (
	IdentificationNewParamsDocumentsDocumentTypeArticlesOfIncorporation   IdentificationNewParamsDocumentsDocumentType = "articles_of_incorporation"
	IdentificationNewParamsDocumentsDocumentTypeCertificateOfGoodStanding IdentificationNewParamsDocumentsDocumentType = "certificate_of_good_standing"
	IdentificationNewParamsDocumentsDocumentTypeEinLetter                 IdentificationNewParamsDocumentsDocumentType = "ein_letter"
	IdentificationNewParamsDocumentsDocumentTypeGeneric                   IdentificationNewParamsDocumentsDocumentType = "generic"
	IdentificationNewParamsDocumentsDocumentTypeIdentificationBack        IdentificationNewParamsDocumentsDocumentType = "identification_back"
	IdentificationNewParamsDocumentsDocumentTypeIdentificationFront       IdentificationNewParamsDocumentsDocumentType = "identification_front"
	IdentificationNewParamsDocumentsDocumentTypeProofOfAddress            IdentificationNewParamsDocumentsDocumentType = "proof_of_address"
)

func (r IdentificationNewParamsDocumentsDocumentType) IsKnown() bool {
	switch r {
	case IdentificationNewParamsDocumentsDocumentTypeArticlesOfIncorporation, IdentificationNewParamsDocumentsDocumentTypeCertificateOfGoodStanding, IdentificationNewParamsDocumentsDocumentTypeEinLetter, IdentificationNewParamsDocumentsDocumentTypeGeneric, IdentificationNewParamsDocumentsDocumentTypeIdentificationBack, IdentificationNewParamsDocumentsDocumentTypeIdentificationFront, IdentificationNewParamsDocumentsDocumentTypeProofOfAddress:
		return true
	}
	return false
}

type IdentificationUpdateParams struct {
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number"`
	// The type of ID number.
	IDType param.Field[IdentificationUpdateParamsIDType] `json:"id_type"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
	// The region in which the identifcation was issued.
	IssuingRegion param.Field[string] `json:"issuing_region"`
}

func (r IdentificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type IdentificationUpdateParamsIDType string

const (
	IdentificationUpdateParamsIDTypeArCuil               IdentificationUpdateParamsIDType = "ar_cuil"
	IdentificationUpdateParamsIDTypeArCuit               IdentificationUpdateParamsIDType = "ar_cuit"
	IdentificationUpdateParamsIDTypeAtAtin               IdentificationUpdateParamsIDType = "at_atin"
	IdentificationUpdateParamsIDTypeAtVat                IdentificationUpdateParamsIDType = "at_vat"
	IdentificationUpdateParamsIDTypeAuAbn                IdentificationUpdateParamsIDType = "au_abn"
	IdentificationUpdateParamsIDTypeAuTfn                IdentificationUpdateParamsIDType = "au_tfn"
	IdentificationUpdateParamsIDTypeBeEnt                IdentificationUpdateParamsIDType = "be_ent"
	IdentificationUpdateParamsIDTypeBeNrn                IdentificationUpdateParamsIDType = "be_nrn"
	IdentificationUpdateParamsIDTypeBrCnpj               IdentificationUpdateParamsIDType = "br_cnpj"
	IdentificationUpdateParamsIDTypeBrCpf                IdentificationUpdateParamsIDType = "br_cpf"
	IdentificationUpdateParamsIDTypeCaBn                 IdentificationUpdateParamsIDType = "ca_bn"
	IdentificationUpdateParamsIDTypeCaSin                IdentificationUpdateParamsIDType = "ca_sin"
	IdentificationUpdateParamsIDTypeChAhv                IdentificationUpdateParamsIDType = "ch_ahv"
	IdentificationUpdateParamsIDTypeChUid                IdentificationUpdateParamsIDType = "ch_uid"
	IdentificationUpdateParamsIDTypeClRun                IdentificationUpdateParamsIDType = "cl_run"
	IdentificationUpdateParamsIDTypeClRut                IdentificationUpdateParamsIDType = "cl_rut"
	IdentificationUpdateParamsIDTypeCoCedulas            IdentificationUpdateParamsIDType = "co_cedulas"
	IdentificationUpdateParamsIDTypeCoNit                IdentificationUpdateParamsIDType = "co_nit"
	IdentificationUpdateParamsIDTypeCyTin                IdentificationUpdateParamsIDType = "cy_tin"
	IdentificationUpdateParamsIDTypeCzIco                IdentificationUpdateParamsIDType = "cz_ico"
	IdentificationUpdateParamsIDTypeCzRc                 IdentificationUpdateParamsIDType = "cz_rc"
	IdentificationUpdateParamsIDTypeDeStid               IdentificationUpdateParamsIDType = "de_stid"
	IdentificationUpdateParamsIDTypeDeStnr               IdentificationUpdateParamsIDType = "de_stnr"
	IdentificationUpdateParamsIDTypeDeVat                IdentificationUpdateParamsIDType = "de_vat"
	IdentificationUpdateParamsIDTypeDkCpr                IdentificationUpdateParamsIDType = "dk_cpr"
	IdentificationUpdateParamsIDTypeDkCvr                IdentificationUpdateParamsIDType = "dk_cvr"
	IdentificationUpdateParamsIDTypeDriversLicense       IdentificationUpdateParamsIDType = "drivers_license"
	IdentificationUpdateParamsIDTypeEeIk                 IdentificationUpdateParamsIDType = "ee_ik"
	IdentificationUpdateParamsIDTypeEeRk                 IdentificationUpdateParamsIDType = "ee_rk"
	IdentificationUpdateParamsIDTypeEsNie                IdentificationUpdateParamsIDType = "es_nie"
	IdentificationUpdateParamsIDTypeEsNif                IdentificationUpdateParamsIDType = "es_nif"
	IdentificationUpdateParamsIDTypeFiHetu               IdentificationUpdateParamsIDType = "fi_hetu"
	IdentificationUpdateParamsIDTypeFiYtj                IdentificationUpdateParamsIDType = "fi_ytj"
	IdentificationUpdateParamsIDTypeFrNif                IdentificationUpdateParamsIDType = "fr_nif"
	IdentificationUpdateParamsIDTypeFrSiren              IdentificationUpdateParamsIDType = "fr_siren"
	IdentificationUpdateParamsIDTypeFrVat                IdentificationUpdateParamsIDType = "fr_vat"
	IdentificationUpdateParamsIDTypeGBNino               IdentificationUpdateParamsIDType = "gb_nino"
	IdentificationUpdateParamsIDTypeGBUtr                IdentificationUpdateParamsIDType = "gb_utr"
	IdentificationUpdateParamsIDTypeGBVat                IdentificationUpdateParamsIDType = "gb_vat"
	IdentificationUpdateParamsIDTypeGenericInternational IdentificationUpdateParamsIDType = "generic_international"
	IdentificationUpdateParamsIDTypeGrVat                IdentificationUpdateParamsIDType = "gr_vat"
	IdentificationUpdateParamsIDTypeHkBrn                IdentificationUpdateParamsIDType = "hk_brn"
	IdentificationUpdateParamsIDTypeHkHkid               IdentificationUpdateParamsIDType = "hk_hkid"
	IdentificationUpdateParamsIDTypeHnID                 IdentificationUpdateParamsIDType = "hn_id"
	IdentificationUpdateParamsIDTypeHnRtn                IdentificationUpdateParamsIDType = "hn_rtn"
	IdentificationUpdateParamsIDTypeHrOib                IdentificationUpdateParamsIDType = "hr_oib"
	IdentificationUpdateParamsIDTypeHuAdj                IdentificationUpdateParamsIDType = "hu_adj"
	IdentificationUpdateParamsIDTypeHuAnum               IdentificationUpdateParamsIDType = "hu_anum"
	IdentificationUpdateParamsIDTypeIePps                IdentificationUpdateParamsIDType = "ie_pps"
	IdentificationUpdateParamsIDTypeIeTrn                IdentificationUpdateParamsIDType = "ie_trn"
	IdentificationUpdateParamsIDTypeInLei                IdentificationUpdateParamsIDType = "in_lei"
	IdentificationUpdateParamsIDTypeIsKnt                IdentificationUpdateParamsIDType = "is_knt"
	IdentificationUpdateParamsIDTypeItCf                 IdentificationUpdateParamsIDType = "it_cf"
	IdentificationUpdateParamsIDTypeItPiva               IdentificationUpdateParamsIDType = "it_piva"
	IdentificationUpdateParamsIDTypeJpHb                 IdentificationUpdateParamsIDType = "jp_hb"
	IdentificationUpdateParamsIDTypeJpMn                 IdentificationUpdateParamsIDType = "jp_mn"
	IdentificationUpdateParamsIDTypeKrBrn                IdentificationUpdateParamsIDType = "kr_brn"
	IdentificationUpdateParamsIDTypeKrCrn                IdentificationUpdateParamsIDType = "kr_crn"
	IdentificationUpdateParamsIDTypeKrRrn                IdentificationUpdateParamsIDType = "kr_rrn"
	IdentificationUpdateParamsIDTypeLiPeid               IdentificationUpdateParamsIDType = "li_peid"
	IdentificationUpdateParamsIDTypeLtAk                 IdentificationUpdateParamsIDType = "lt_ak"
	IdentificationUpdateParamsIDTypeLtJak                IdentificationUpdateParamsIDType = "lt_jak"
	IdentificationUpdateParamsIDTypeLuMtc                IdentificationUpdateParamsIDType = "lu_mtc"
	IdentificationUpdateParamsIDTypeLuVat                IdentificationUpdateParamsIDType = "lu_vat"
	IdentificationUpdateParamsIDTypeLvPk                 IdentificationUpdateParamsIDType = "lv_pk"
	IdentificationUpdateParamsIDTypeLvRn                 IdentificationUpdateParamsIDType = "lv_rn"
	IdentificationUpdateParamsIDTypeMtTin                IdentificationUpdateParamsIDType = "mt_tin"
	IdentificationUpdateParamsIDTypeMtVat                IdentificationUpdateParamsIDType = "mt_vat"
	IdentificationUpdateParamsIDTypeMxCurp               IdentificationUpdateParamsIDType = "mx_curp"
	IdentificationUpdateParamsIDTypeMxIne                IdentificationUpdateParamsIDType = "mx_ine"
	IdentificationUpdateParamsIDTypeMxRfc                IdentificationUpdateParamsIDType = "mx_rfc"
	IdentificationUpdateParamsIDTypeNationalID           IdentificationUpdateParamsIDType = "national_id"
	IdentificationUpdateParamsIDTypeNlBsn                IdentificationUpdateParamsIDType = "nl_bsn"
	IdentificationUpdateParamsIDTypeNlBtw                IdentificationUpdateParamsIDType = "nl_btw"
	IdentificationUpdateParamsIDTypeNlRsin               IdentificationUpdateParamsIDType = "nl_rsin"
	IdentificationUpdateParamsIDTypeNoFdn                IdentificationUpdateParamsIDType = "no_fdn"
	IdentificationUpdateParamsIDTypeNoMva                IdentificationUpdateParamsIDType = "no_mva"
	IdentificationUpdateParamsIDTypeNoOrgnr              IdentificationUpdateParamsIDType = "no_orgnr"
	IdentificationUpdateParamsIDTypeNzIrd                IdentificationUpdateParamsIDType = "nz_ird"
	IdentificationUpdateParamsIDTypePassport             IdentificationUpdateParamsIDType = "passport"
	IdentificationUpdateParamsIDTypePlNip                IdentificationUpdateParamsIDType = "pl_nip"
	IdentificationUpdateParamsIDTypePlPesel              IdentificationUpdateParamsIDType = "pl_pesel"
	IdentificationUpdateParamsIDTypePtNif                IdentificationUpdateParamsIDType = "pt_nif"
	IdentificationUpdateParamsIDTypeRoCnp                IdentificationUpdateParamsIDType = "ro_cnp"
	IdentificationUpdateParamsIDTypeRoCui                IdentificationUpdateParamsIDType = "ro_cui"
	IdentificationUpdateParamsIDTypeSaTin                IdentificationUpdateParamsIDType = "sa_tin"
	IdentificationUpdateParamsIDTypeSaVat                IdentificationUpdateParamsIDType = "sa_vat"
	IdentificationUpdateParamsIDTypeSeOrgnr              IdentificationUpdateParamsIDType = "se_orgnr"
	IdentificationUpdateParamsIDTypeSePnmr               IdentificationUpdateParamsIDType = "se_pnmr"
	IdentificationUpdateParamsIDTypeSgFin                IdentificationUpdateParamsIDType = "sg_fin"
	IdentificationUpdateParamsIDTypeSgNric               IdentificationUpdateParamsIDType = "sg_nric"
	IdentificationUpdateParamsIDTypeSgUen                IdentificationUpdateParamsIDType = "sg_uen"
	IdentificationUpdateParamsIDTypeSiDav                IdentificationUpdateParamsIDType = "si_dav"
	IdentificationUpdateParamsIDTypeSiTin                IdentificationUpdateParamsIDType = "si_tin"
	IdentificationUpdateParamsIDTypeSkIco                IdentificationUpdateParamsIDType = "sk_ico"
	IdentificationUpdateParamsIDTypeSkRc                 IdentificationUpdateParamsIDType = "sk_rc"
	IdentificationUpdateParamsIDTypeUsEin                IdentificationUpdateParamsIDType = "us_ein"
	IdentificationUpdateParamsIDTypeUsItin               IdentificationUpdateParamsIDType = "us_itin"
	IdentificationUpdateParamsIDTypeUsSsn                IdentificationUpdateParamsIDType = "us_ssn"
	IdentificationUpdateParamsIDTypeUyRut                IdentificationUpdateParamsIDType = "uy_rut"
	IdentificationUpdateParamsIDTypeVnTin                IdentificationUpdateParamsIDType = "vn_tin"
)

func (r IdentificationUpdateParamsIDType) IsKnown() bool {
	switch r {
	case IdentificationUpdateParamsIDTypeArCuil, IdentificationUpdateParamsIDTypeArCuit, IdentificationUpdateParamsIDTypeAtAtin, IdentificationUpdateParamsIDTypeAtVat, IdentificationUpdateParamsIDTypeAuAbn, IdentificationUpdateParamsIDTypeAuTfn, IdentificationUpdateParamsIDTypeBeEnt, IdentificationUpdateParamsIDTypeBeNrn, IdentificationUpdateParamsIDTypeBrCnpj, IdentificationUpdateParamsIDTypeBrCpf, IdentificationUpdateParamsIDTypeCaBn, IdentificationUpdateParamsIDTypeCaSin, IdentificationUpdateParamsIDTypeChAhv, IdentificationUpdateParamsIDTypeChUid, IdentificationUpdateParamsIDTypeClRun, IdentificationUpdateParamsIDTypeClRut, IdentificationUpdateParamsIDTypeCoCedulas, IdentificationUpdateParamsIDTypeCoNit, IdentificationUpdateParamsIDTypeCyTin, IdentificationUpdateParamsIDTypeCzIco, IdentificationUpdateParamsIDTypeCzRc, IdentificationUpdateParamsIDTypeDeStid, IdentificationUpdateParamsIDTypeDeStnr, IdentificationUpdateParamsIDTypeDeVat, IdentificationUpdateParamsIDTypeDkCpr, IdentificationUpdateParamsIDTypeDkCvr, IdentificationUpdateParamsIDTypeDriversLicense, IdentificationUpdateParamsIDTypeEeIk, IdentificationUpdateParamsIDTypeEeRk, IdentificationUpdateParamsIDTypeEsNie, IdentificationUpdateParamsIDTypeEsNif, IdentificationUpdateParamsIDTypeFiHetu, IdentificationUpdateParamsIDTypeFiYtj, IdentificationUpdateParamsIDTypeFrNif, IdentificationUpdateParamsIDTypeFrSiren, IdentificationUpdateParamsIDTypeFrVat, IdentificationUpdateParamsIDTypeGBNino, IdentificationUpdateParamsIDTypeGBUtr, IdentificationUpdateParamsIDTypeGBVat, IdentificationUpdateParamsIDTypeGenericInternational, IdentificationUpdateParamsIDTypeGrVat, IdentificationUpdateParamsIDTypeHkBrn, IdentificationUpdateParamsIDTypeHkHkid, IdentificationUpdateParamsIDTypeHnID, IdentificationUpdateParamsIDTypeHnRtn, IdentificationUpdateParamsIDTypeHrOib, IdentificationUpdateParamsIDTypeHuAdj, IdentificationUpdateParamsIDTypeHuAnum, IdentificationUpdateParamsIDTypeIePps, IdentificationUpdateParamsIDTypeIeTrn, IdentificationUpdateParamsIDTypeInLei, IdentificationUpdateParamsIDTypeIsKnt, IdentificationUpdateParamsIDTypeItCf, IdentificationUpdateParamsIDTypeItPiva, IdentificationUpdateParamsIDTypeJpHb, IdentificationUpdateParamsIDTypeJpMn, IdentificationUpdateParamsIDTypeKrBrn, IdentificationUpdateParamsIDTypeKrCrn, IdentificationUpdateParamsIDTypeKrRrn, IdentificationUpdateParamsIDTypeLiPeid, IdentificationUpdateParamsIDTypeLtAk, IdentificationUpdateParamsIDTypeLtJak, IdentificationUpdateParamsIDTypeLuMtc, IdentificationUpdateParamsIDTypeLuVat, IdentificationUpdateParamsIDTypeLvPk, IdentificationUpdateParamsIDTypeLvRn, IdentificationUpdateParamsIDTypeMtTin, IdentificationUpdateParamsIDTypeMtVat, IdentificationUpdateParamsIDTypeMxCurp, IdentificationUpdateParamsIDTypeMxIne, IdentificationUpdateParamsIDTypeMxRfc, IdentificationUpdateParamsIDTypeNationalID, IdentificationUpdateParamsIDTypeNlBsn, IdentificationUpdateParamsIDTypeNlBtw, IdentificationUpdateParamsIDTypeNlRsin, IdentificationUpdateParamsIDTypeNoFdn, IdentificationUpdateParamsIDTypeNoMva, IdentificationUpdateParamsIDTypeNoOrgnr, IdentificationUpdateParamsIDTypeNzIrd, IdentificationUpdateParamsIDTypePassport, IdentificationUpdateParamsIDTypePlNip, IdentificationUpdateParamsIDTypePlPesel, IdentificationUpdateParamsIDTypePtNif, IdentificationUpdateParamsIDTypeRoCnp, IdentificationUpdateParamsIDTypeRoCui, IdentificationUpdateParamsIDTypeSaTin, IdentificationUpdateParamsIDTypeSaVat, IdentificationUpdateParamsIDTypeSeOrgnr, IdentificationUpdateParamsIDTypeSePnmr, IdentificationUpdateParamsIDTypeSgFin, IdentificationUpdateParamsIDTypeSgNric, IdentificationUpdateParamsIDTypeSgUen, IdentificationUpdateParamsIDTypeSiDav, IdentificationUpdateParamsIDTypeSiTin, IdentificationUpdateParamsIDTypeSkIco, IdentificationUpdateParamsIDTypeSkRc, IdentificationUpdateParamsIDTypeUsEin, IdentificationUpdateParamsIDTypeUsItin, IdentificationUpdateParamsIDTypeUsSsn, IdentificationUpdateParamsIDTypeUyRut, IdentificationUpdateParamsIDTypeVnTin:
		return true
	}
	return false
}
