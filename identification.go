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
	IdentificationIDTypeAdNrt                IdentificationIDType = "ad_nrt"
	IdentificationIDTypeAeEid                IdentificationIDType = "ae_eid"
	IdentificationIDTypeAeTrn                IdentificationIDType = "ae_trn"
	IdentificationIDTypeAgTin                IdentificationIDType = "ag_tin"
	IdentificationIDTypeAITin                IdentificationIDType = "ai_tin"
	IdentificationIDTypeAlNid                IdentificationIDType = "al_nid"
	IdentificationIDTypeAlNipt               IdentificationIDType = "al_nipt"
	IdentificationIDTypeAmTin                IdentificationIDType = "am_tin"
	IdentificationIDTypeAoNif                IdentificationIDType = "ao_nif"
	IdentificationIDTypeArCuil               IdentificationIDType = "ar_cuil"
	IdentificationIDTypeArCuit               IdentificationIDType = "ar_cuit"
	IdentificationIDTypeAtAtin               IdentificationIDType = "at_atin"
	IdentificationIDTypeAtVat                IdentificationIDType = "at_vat"
	IdentificationIDTypeAuAbn                IdentificationIDType = "au_abn"
	IdentificationIDTypeAuTfn                IdentificationIDType = "au_tfn"
	IdentificationIDTypeAwTin                IdentificationIDType = "aw_tin"
	IdentificationIDTypeAzPin                IdentificationIDType = "az_pin"
	IdentificationIDTypeBbTin                IdentificationIDType = "bb_tin"
	IdentificationIDTypeBdTin                IdentificationIDType = "bd_tin"
	IdentificationIDTypeBeEnt                IdentificationIDType = "be_ent"
	IdentificationIDTypeBeNrn                IdentificationIDType = "be_nrn"
	IdentificationIDTypeBfIfu                IdentificationIDType = "bf_ifu"
	IdentificationIDTypeBgEgn                IdentificationIDType = "bg_egn"
	IdentificationIDTypeBhCpr                IdentificationIDType = "bh_cpr"
	IdentificationIDTypeBhVat                IdentificationIDType = "bh_vat"
	IdentificationIDTypeBjIfu                IdentificationIDType = "bj_ifu"
	IdentificationIDTypeBoNit                IdentificationIDType = "bo_nit"
	IdentificationIDTypeBrCnpj               IdentificationIDType = "br_cnpj"
	IdentificationIDTypeBrCpf                IdentificationIDType = "br_cpf"
	IdentificationIDTypeBsTin                IdentificationIDType = "bs_tin"
	IdentificationIDTypeBtBin                IdentificationIDType = "bt_bin"
	IdentificationIDTypeBwTin                IdentificationIDType = "bw_tin"
	IdentificationIDTypeBzTin                IdentificationIDType = "bz_tin"
	IdentificationIDTypeCaBn                 IdentificationIDType = "ca_bn"
	IdentificationIDTypeCaSin                IdentificationIDType = "ca_sin"
	IdentificationIDTypeChAhv                IdentificationIDType = "ch_ahv"
	IdentificationIDTypeChUid                IdentificationIDType = "ch_uid"
	IdentificationIDTypeCiNcc                IdentificationIDType = "ci_ncc"
	IdentificationIDTypeClRun                IdentificationIDType = "cl_run"
	IdentificationIDTypeClRut                IdentificationIDType = "cl_rut"
	IdentificationIDTypeCmNiu                IdentificationIDType = "cm_niu"
	IdentificationIDTypeCoCedulas            IdentificationIDType = "co_cedulas"
	IdentificationIDTypeCoNit                IdentificationIDType = "co_nit"
	IdentificationIDTypeCrCpf                IdentificationIDType = "cr_cpf"
	IdentificationIDTypeCwCrib               IdentificationIDType = "cw_crib"
	IdentificationIDTypeCyTin                IdentificationIDType = "cy_tin"
	IdentificationIDTypeCzIco                IdentificationIDType = "cz_ico"
	IdentificationIDTypeCzRc                 IdentificationIDType = "cz_rc"
	IdentificationIDTypeDeStid               IdentificationIDType = "de_stid"
	IdentificationIDTypeDeStnr               IdentificationIDType = "de_stnr"
	IdentificationIDTypeDeVat                IdentificationIDType = "de_vat"
	IdentificationIDTypeDkCpr                IdentificationIDType = "dk_cpr"
	IdentificationIDTypeDkCvr                IdentificationIDType = "dk_cvr"
	IdentificationIDTypeDmTin                IdentificationIDType = "dm_tin"
	IdentificationIDTypeDoCedula             IdentificationIDType = "do_cedula"
	IdentificationIDTypeDoRnc                IdentificationIDType = "do_rnc"
	IdentificationIDTypeDriversLicense       IdentificationIDType = "drivers_license"
	IdentificationIDTypeEcRuc                IdentificationIDType = "ec_ruc"
	IdentificationIDTypeEeIk                 IdentificationIDType = "ee_ik"
	IdentificationIDTypeEeRk                 IdentificationIDType = "ee_rk"
	IdentificationIDTypeEsNie                IdentificationIDType = "es_nie"
	IdentificationIDTypeEsNif                IdentificationIDType = "es_nif"
	IdentificationIDTypeFiHetu               IdentificationIDType = "fi_hetu"
	IdentificationIDTypeFiYtj                IdentificationIDType = "fi_ytj"
	IdentificationIDTypeFjTin                IdentificationIDType = "fj_tin"
	IdentificationIDTypeFoPtal               IdentificationIDType = "fo_ptal"
	IdentificationIDTypeFrNif                IdentificationIDType = "fr_nif"
	IdentificationIDTypeFrSiren              IdentificationIDType = "fr_siren"
	IdentificationIDTypeFrVat                IdentificationIDType = "fr_vat"
	IdentificationIDTypeGBNino               IdentificationIDType = "gb_nino"
	IdentificationIDTypeGBUtr                IdentificationIDType = "gb_utr"
	IdentificationIDTypeGBVat                IdentificationIDType = "gb_vat"
	IdentificationIDTypeGdTin                IdentificationIDType = "gd_tin"
	IdentificationIDTypeGeIc                 IdentificationIDType = "ge_ic"
	IdentificationIDTypeGePn                 IdentificationIDType = "ge_pn"
	IdentificationIDTypeGeTin                IdentificationIDType = "ge_tin"
	IdentificationIDTypeGenericInternational IdentificationIDType = "generic_international"
	IdentificationIDTypeGgSin                IdentificationIDType = "gg_sin"
	IdentificationIDTypeGhPin                IdentificationIDType = "gh_pin"
	IdentificationIDTypeGhTin                IdentificationIDType = "gh_tin"
	IdentificationIDTypeGiTrn                IdentificationIDType = "gi_trn"
	IdentificationIDTypeGlCpr                IdentificationIDType = "gl_cpr"
	IdentificationIDTypeGlGer                IdentificationIDType = "gl_ger"
	IdentificationIDTypeGmTin                IdentificationIDType = "gm_tin"
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
	case IdentificationIDTypeAdNrt, IdentificationIDTypeAeEid, IdentificationIDTypeAeTrn, IdentificationIDTypeAgTin, IdentificationIDTypeAITin, IdentificationIDTypeAlNid, IdentificationIDTypeAlNipt, IdentificationIDTypeAmTin, IdentificationIDTypeAoNif, IdentificationIDTypeArCuil, IdentificationIDTypeArCuit, IdentificationIDTypeAtAtin, IdentificationIDTypeAtVat, IdentificationIDTypeAuAbn, IdentificationIDTypeAuTfn, IdentificationIDTypeAwTin, IdentificationIDTypeAzPin, IdentificationIDTypeBbTin, IdentificationIDTypeBdTin, IdentificationIDTypeBeEnt, IdentificationIDTypeBeNrn, IdentificationIDTypeBfIfu, IdentificationIDTypeBgEgn, IdentificationIDTypeBhCpr, IdentificationIDTypeBhVat, IdentificationIDTypeBjIfu, IdentificationIDTypeBoNit, IdentificationIDTypeBrCnpj, IdentificationIDTypeBrCpf, IdentificationIDTypeBsTin, IdentificationIDTypeBtBin, IdentificationIDTypeBwTin, IdentificationIDTypeBzTin, IdentificationIDTypeCaBn, IdentificationIDTypeCaSin, IdentificationIDTypeChAhv, IdentificationIDTypeChUid, IdentificationIDTypeCiNcc, IdentificationIDTypeClRun, IdentificationIDTypeClRut, IdentificationIDTypeCmNiu, IdentificationIDTypeCoCedulas, IdentificationIDTypeCoNit, IdentificationIDTypeCrCpf, IdentificationIDTypeCwCrib, IdentificationIDTypeCyTin, IdentificationIDTypeCzIco, IdentificationIDTypeCzRc, IdentificationIDTypeDeStid, IdentificationIDTypeDeStnr, IdentificationIDTypeDeVat, IdentificationIDTypeDkCpr, IdentificationIDTypeDkCvr, IdentificationIDTypeDmTin, IdentificationIDTypeDoCedula, IdentificationIDTypeDoRnc, IdentificationIDTypeDriversLicense, IdentificationIDTypeEcRuc, IdentificationIDTypeEeIk, IdentificationIDTypeEeRk, IdentificationIDTypeEsNie, IdentificationIDTypeEsNif, IdentificationIDTypeFiHetu, IdentificationIDTypeFiYtj, IdentificationIDTypeFjTin, IdentificationIDTypeFoPtal, IdentificationIDTypeFrNif, IdentificationIDTypeFrSiren, IdentificationIDTypeFrVat, IdentificationIDTypeGBNino, IdentificationIDTypeGBUtr, IdentificationIDTypeGBVat, IdentificationIDTypeGdTin, IdentificationIDTypeGeIc, IdentificationIDTypeGePn, IdentificationIDTypeGeTin, IdentificationIDTypeGenericInternational, IdentificationIDTypeGgSin, IdentificationIDTypeGhPin, IdentificationIDTypeGhTin, IdentificationIDTypeGiTrn, IdentificationIDTypeGlCpr, IdentificationIDTypeGlGer, IdentificationIDTypeGmTin, IdentificationIDTypeGrVat, IdentificationIDTypeHkBrn, IdentificationIDTypeHkHkid, IdentificationIDTypeHnID, IdentificationIDTypeHnRtn, IdentificationIDTypeHrOib, IdentificationIDTypeHuAdj, IdentificationIDTypeHuAnum, IdentificationIDTypeIePps, IdentificationIDTypeIeTrn, IdentificationIDTypeInLei, IdentificationIDTypeIsKnt, IdentificationIDTypeItCf, IdentificationIDTypeItPiva, IdentificationIDTypeJpHb, IdentificationIDTypeJpMn, IdentificationIDTypeKrBrn, IdentificationIDTypeKrCrn, IdentificationIDTypeKrRrn, IdentificationIDTypeLiPeid, IdentificationIDTypeLtAk, IdentificationIDTypeLtJak, IdentificationIDTypeLuMtc, IdentificationIDTypeLuVat, IdentificationIDTypeLvPk, IdentificationIDTypeLvRn, IdentificationIDTypeMtTin, IdentificationIDTypeMtVat, IdentificationIDTypeMxCurp, IdentificationIDTypeMxIne, IdentificationIDTypeMxRfc, IdentificationIDTypeNationalID, IdentificationIDTypeNlBsn, IdentificationIDTypeNlBtw, IdentificationIDTypeNlRsin, IdentificationIDTypeNoFdn, IdentificationIDTypeNoMva, IdentificationIDTypeNoOrgnr, IdentificationIDTypeNzIrd, IdentificationIDTypePassport, IdentificationIDTypePlNip, IdentificationIDTypePlPesel, IdentificationIDTypePtNif, IdentificationIDTypeRoCnp, IdentificationIDTypeRoCui, IdentificationIDTypeSaTin, IdentificationIDTypeSaVat, IdentificationIDTypeSeOrgnr, IdentificationIDTypeSePnmr, IdentificationIDTypeSgFin, IdentificationIDTypeSgNric, IdentificationIDTypeSgUen, IdentificationIDTypeSiDav, IdentificationIDTypeSiTin, IdentificationIDTypeSkIco, IdentificationIDTypeSkRc, IdentificationIDTypeUsEin, IdentificationIDTypeUsItin, IdentificationIDTypeUsSsn, IdentificationIDTypeUyRut, IdentificationIDTypeVnTin:
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
	IdentificationNewParamsIDTypeAdNrt                IdentificationNewParamsIDType = "ad_nrt"
	IdentificationNewParamsIDTypeAeEid                IdentificationNewParamsIDType = "ae_eid"
	IdentificationNewParamsIDTypeAeTrn                IdentificationNewParamsIDType = "ae_trn"
	IdentificationNewParamsIDTypeAgTin                IdentificationNewParamsIDType = "ag_tin"
	IdentificationNewParamsIDTypeAITin                IdentificationNewParamsIDType = "ai_tin"
	IdentificationNewParamsIDTypeAlNid                IdentificationNewParamsIDType = "al_nid"
	IdentificationNewParamsIDTypeAlNipt               IdentificationNewParamsIDType = "al_nipt"
	IdentificationNewParamsIDTypeAmTin                IdentificationNewParamsIDType = "am_tin"
	IdentificationNewParamsIDTypeAoNif                IdentificationNewParamsIDType = "ao_nif"
	IdentificationNewParamsIDTypeArCuil               IdentificationNewParamsIDType = "ar_cuil"
	IdentificationNewParamsIDTypeArCuit               IdentificationNewParamsIDType = "ar_cuit"
	IdentificationNewParamsIDTypeAtAtin               IdentificationNewParamsIDType = "at_atin"
	IdentificationNewParamsIDTypeAtVat                IdentificationNewParamsIDType = "at_vat"
	IdentificationNewParamsIDTypeAuAbn                IdentificationNewParamsIDType = "au_abn"
	IdentificationNewParamsIDTypeAuTfn                IdentificationNewParamsIDType = "au_tfn"
	IdentificationNewParamsIDTypeAwTin                IdentificationNewParamsIDType = "aw_tin"
	IdentificationNewParamsIDTypeAzPin                IdentificationNewParamsIDType = "az_pin"
	IdentificationNewParamsIDTypeBbTin                IdentificationNewParamsIDType = "bb_tin"
	IdentificationNewParamsIDTypeBdTin                IdentificationNewParamsIDType = "bd_tin"
	IdentificationNewParamsIDTypeBeEnt                IdentificationNewParamsIDType = "be_ent"
	IdentificationNewParamsIDTypeBeNrn                IdentificationNewParamsIDType = "be_nrn"
	IdentificationNewParamsIDTypeBfIfu                IdentificationNewParamsIDType = "bf_ifu"
	IdentificationNewParamsIDTypeBgEgn                IdentificationNewParamsIDType = "bg_egn"
	IdentificationNewParamsIDTypeBhCpr                IdentificationNewParamsIDType = "bh_cpr"
	IdentificationNewParamsIDTypeBhVat                IdentificationNewParamsIDType = "bh_vat"
	IdentificationNewParamsIDTypeBjIfu                IdentificationNewParamsIDType = "bj_ifu"
	IdentificationNewParamsIDTypeBoNit                IdentificationNewParamsIDType = "bo_nit"
	IdentificationNewParamsIDTypeBrCnpj               IdentificationNewParamsIDType = "br_cnpj"
	IdentificationNewParamsIDTypeBrCpf                IdentificationNewParamsIDType = "br_cpf"
	IdentificationNewParamsIDTypeBsTin                IdentificationNewParamsIDType = "bs_tin"
	IdentificationNewParamsIDTypeBtBin                IdentificationNewParamsIDType = "bt_bin"
	IdentificationNewParamsIDTypeBwTin                IdentificationNewParamsIDType = "bw_tin"
	IdentificationNewParamsIDTypeBzTin                IdentificationNewParamsIDType = "bz_tin"
	IdentificationNewParamsIDTypeCaBn                 IdentificationNewParamsIDType = "ca_bn"
	IdentificationNewParamsIDTypeCaSin                IdentificationNewParamsIDType = "ca_sin"
	IdentificationNewParamsIDTypeChAhv                IdentificationNewParamsIDType = "ch_ahv"
	IdentificationNewParamsIDTypeChUid                IdentificationNewParamsIDType = "ch_uid"
	IdentificationNewParamsIDTypeCiNcc                IdentificationNewParamsIDType = "ci_ncc"
	IdentificationNewParamsIDTypeClRun                IdentificationNewParamsIDType = "cl_run"
	IdentificationNewParamsIDTypeClRut                IdentificationNewParamsIDType = "cl_rut"
	IdentificationNewParamsIDTypeCmNiu                IdentificationNewParamsIDType = "cm_niu"
	IdentificationNewParamsIDTypeCoCedulas            IdentificationNewParamsIDType = "co_cedulas"
	IdentificationNewParamsIDTypeCoNit                IdentificationNewParamsIDType = "co_nit"
	IdentificationNewParamsIDTypeCrCpf                IdentificationNewParamsIDType = "cr_cpf"
	IdentificationNewParamsIDTypeCwCrib               IdentificationNewParamsIDType = "cw_crib"
	IdentificationNewParamsIDTypeCyTin                IdentificationNewParamsIDType = "cy_tin"
	IdentificationNewParamsIDTypeCzIco                IdentificationNewParamsIDType = "cz_ico"
	IdentificationNewParamsIDTypeCzRc                 IdentificationNewParamsIDType = "cz_rc"
	IdentificationNewParamsIDTypeDeStid               IdentificationNewParamsIDType = "de_stid"
	IdentificationNewParamsIDTypeDeStnr               IdentificationNewParamsIDType = "de_stnr"
	IdentificationNewParamsIDTypeDeVat                IdentificationNewParamsIDType = "de_vat"
	IdentificationNewParamsIDTypeDkCpr                IdentificationNewParamsIDType = "dk_cpr"
	IdentificationNewParamsIDTypeDkCvr                IdentificationNewParamsIDType = "dk_cvr"
	IdentificationNewParamsIDTypeDmTin                IdentificationNewParamsIDType = "dm_tin"
	IdentificationNewParamsIDTypeDoCedula             IdentificationNewParamsIDType = "do_cedula"
	IdentificationNewParamsIDTypeDoRnc                IdentificationNewParamsIDType = "do_rnc"
	IdentificationNewParamsIDTypeDriversLicense       IdentificationNewParamsIDType = "drivers_license"
	IdentificationNewParamsIDTypeEcRuc                IdentificationNewParamsIDType = "ec_ruc"
	IdentificationNewParamsIDTypeEeIk                 IdentificationNewParamsIDType = "ee_ik"
	IdentificationNewParamsIDTypeEeRk                 IdentificationNewParamsIDType = "ee_rk"
	IdentificationNewParamsIDTypeEsNie                IdentificationNewParamsIDType = "es_nie"
	IdentificationNewParamsIDTypeEsNif                IdentificationNewParamsIDType = "es_nif"
	IdentificationNewParamsIDTypeFiHetu               IdentificationNewParamsIDType = "fi_hetu"
	IdentificationNewParamsIDTypeFiYtj                IdentificationNewParamsIDType = "fi_ytj"
	IdentificationNewParamsIDTypeFjTin                IdentificationNewParamsIDType = "fj_tin"
	IdentificationNewParamsIDTypeFoPtal               IdentificationNewParamsIDType = "fo_ptal"
	IdentificationNewParamsIDTypeFrNif                IdentificationNewParamsIDType = "fr_nif"
	IdentificationNewParamsIDTypeFrSiren              IdentificationNewParamsIDType = "fr_siren"
	IdentificationNewParamsIDTypeFrVat                IdentificationNewParamsIDType = "fr_vat"
	IdentificationNewParamsIDTypeGBNino               IdentificationNewParamsIDType = "gb_nino"
	IdentificationNewParamsIDTypeGBUtr                IdentificationNewParamsIDType = "gb_utr"
	IdentificationNewParamsIDTypeGBVat                IdentificationNewParamsIDType = "gb_vat"
	IdentificationNewParamsIDTypeGdTin                IdentificationNewParamsIDType = "gd_tin"
	IdentificationNewParamsIDTypeGeIc                 IdentificationNewParamsIDType = "ge_ic"
	IdentificationNewParamsIDTypeGePn                 IdentificationNewParamsIDType = "ge_pn"
	IdentificationNewParamsIDTypeGeTin                IdentificationNewParamsIDType = "ge_tin"
	IdentificationNewParamsIDTypeGenericInternational IdentificationNewParamsIDType = "generic_international"
	IdentificationNewParamsIDTypeGgSin                IdentificationNewParamsIDType = "gg_sin"
	IdentificationNewParamsIDTypeGhPin                IdentificationNewParamsIDType = "gh_pin"
	IdentificationNewParamsIDTypeGhTin                IdentificationNewParamsIDType = "gh_tin"
	IdentificationNewParamsIDTypeGiTrn                IdentificationNewParamsIDType = "gi_trn"
	IdentificationNewParamsIDTypeGlCpr                IdentificationNewParamsIDType = "gl_cpr"
	IdentificationNewParamsIDTypeGlGer                IdentificationNewParamsIDType = "gl_ger"
	IdentificationNewParamsIDTypeGmTin                IdentificationNewParamsIDType = "gm_tin"
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
	case IdentificationNewParamsIDTypeAdNrt, IdentificationNewParamsIDTypeAeEid, IdentificationNewParamsIDTypeAeTrn, IdentificationNewParamsIDTypeAgTin, IdentificationNewParamsIDTypeAITin, IdentificationNewParamsIDTypeAlNid, IdentificationNewParamsIDTypeAlNipt, IdentificationNewParamsIDTypeAmTin, IdentificationNewParamsIDTypeAoNif, IdentificationNewParamsIDTypeArCuil, IdentificationNewParamsIDTypeArCuit, IdentificationNewParamsIDTypeAtAtin, IdentificationNewParamsIDTypeAtVat, IdentificationNewParamsIDTypeAuAbn, IdentificationNewParamsIDTypeAuTfn, IdentificationNewParamsIDTypeAwTin, IdentificationNewParamsIDTypeAzPin, IdentificationNewParamsIDTypeBbTin, IdentificationNewParamsIDTypeBdTin, IdentificationNewParamsIDTypeBeEnt, IdentificationNewParamsIDTypeBeNrn, IdentificationNewParamsIDTypeBfIfu, IdentificationNewParamsIDTypeBgEgn, IdentificationNewParamsIDTypeBhCpr, IdentificationNewParamsIDTypeBhVat, IdentificationNewParamsIDTypeBjIfu, IdentificationNewParamsIDTypeBoNit, IdentificationNewParamsIDTypeBrCnpj, IdentificationNewParamsIDTypeBrCpf, IdentificationNewParamsIDTypeBsTin, IdentificationNewParamsIDTypeBtBin, IdentificationNewParamsIDTypeBwTin, IdentificationNewParamsIDTypeBzTin, IdentificationNewParamsIDTypeCaBn, IdentificationNewParamsIDTypeCaSin, IdentificationNewParamsIDTypeChAhv, IdentificationNewParamsIDTypeChUid, IdentificationNewParamsIDTypeCiNcc, IdentificationNewParamsIDTypeClRun, IdentificationNewParamsIDTypeClRut, IdentificationNewParamsIDTypeCmNiu, IdentificationNewParamsIDTypeCoCedulas, IdentificationNewParamsIDTypeCoNit, IdentificationNewParamsIDTypeCrCpf, IdentificationNewParamsIDTypeCwCrib, IdentificationNewParamsIDTypeCyTin, IdentificationNewParamsIDTypeCzIco, IdentificationNewParamsIDTypeCzRc, IdentificationNewParamsIDTypeDeStid, IdentificationNewParamsIDTypeDeStnr, IdentificationNewParamsIDTypeDeVat, IdentificationNewParamsIDTypeDkCpr, IdentificationNewParamsIDTypeDkCvr, IdentificationNewParamsIDTypeDmTin, IdentificationNewParamsIDTypeDoCedula, IdentificationNewParamsIDTypeDoRnc, IdentificationNewParamsIDTypeDriversLicense, IdentificationNewParamsIDTypeEcRuc, IdentificationNewParamsIDTypeEeIk, IdentificationNewParamsIDTypeEeRk, IdentificationNewParamsIDTypeEsNie, IdentificationNewParamsIDTypeEsNif, IdentificationNewParamsIDTypeFiHetu, IdentificationNewParamsIDTypeFiYtj, IdentificationNewParamsIDTypeFjTin, IdentificationNewParamsIDTypeFoPtal, IdentificationNewParamsIDTypeFrNif, IdentificationNewParamsIDTypeFrSiren, IdentificationNewParamsIDTypeFrVat, IdentificationNewParamsIDTypeGBNino, IdentificationNewParamsIDTypeGBUtr, IdentificationNewParamsIDTypeGBVat, IdentificationNewParamsIDTypeGdTin, IdentificationNewParamsIDTypeGeIc, IdentificationNewParamsIDTypeGePn, IdentificationNewParamsIDTypeGeTin, IdentificationNewParamsIDTypeGenericInternational, IdentificationNewParamsIDTypeGgSin, IdentificationNewParamsIDTypeGhPin, IdentificationNewParamsIDTypeGhTin, IdentificationNewParamsIDTypeGiTrn, IdentificationNewParamsIDTypeGlCpr, IdentificationNewParamsIDTypeGlGer, IdentificationNewParamsIDTypeGmTin, IdentificationNewParamsIDTypeGrVat, IdentificationNewParamsIDTypeHkBrn, IdentificationNewParamsIDTypeHkHkid, IdentificationNewParamsIDTypeHnID, IdentificationNewParamsIDTypeHnRtn, IdentificationNewParamsIDTypeHrOib, IdentificationNewParamsIDTypeHuAdj, IdentificationNewParamsIDTypeHuAnum, IdentificationNewParamsIDTypeIePps, IdentificationNewParamsIDTypeIeTrn, IdentificationNewParamsIDTypeInLei, IdentificationNewParamsIDTypeIsKnt, IdentificationNewParamsIDTypeItCf, IdentificationNewParamsIDTypeItPiva, IdentificationNewParamsIDTypeJpHb, IdentificationNewParamsIDTypeJpMn, IdentificationNewParamsIDTypeKrBrn, IdentificationNewParamsIDTypeKrCrn, IdentificationNewParamsIDTypeKrRrn, IdentificationNewParamsIDTypeLiPeid, IdentificationNewParamsIDTypeLtAk, IdentificationNewParamsIDTypeLtJak, IdentificationNewParamsIDTypeLuMtc, IdentificationNewParamsIDTypeLuVat, IdentificationNewParamsIDTypeLvPk, IdentificationNewParamsIDTypeLvRn, IdentificationNewParamsIDTypeMtTin, IdentificationNewParamsIDTypeMtVat, IdentificationNewParamsIDTypeMxCurp, IdentificationNewParamsIDTypeMxIne, IdentificationNewParamsIDTypeMxRfc, IdentificationNewParamsIDTypeNationalID, IdentificationNewParamsIDTypeNlBsn, IdentificationNewParamsIDTypeNlBtw, IdentificationNewParamsIDTypeNlRsin, IdentificationNewParamsIDTypeNoFdn, IdentificationNewParamsIDTypeNoMva, IdentificationNewParamsIDTypeNoOrgnr, IdentificationNewParamsIDTypeNzIrd, IdentificationNewParamsIDTypePassport, IdentificationNewParamsIDTypePlNip, IdentificationNewParamsIDTypePlPesel, IdentificationNewParamsIDTypePtNif, IdentificationNewParamsIDTypeRoCnp, IdentificationNewParamsIDTypeRoCui, IdentificationNewParamsIDTypeSaTin, IdentificationNewParamsIDTypeSaVat, IdentificationNewParamsIDTypeSeOrgnr, IdentificationNewParamsIDTypeSePnmr, IdentificationNewParamsIDTypeSgFin, IdentificationNewParamsIDTypeSgNric, IdentificationNewParamsIDTypeSgUen, IdentificationNewParamsIDTypeSiDav, IdentificationNewParamsIDTypeSiTin, IdentificationNewParamsIDTypeSkIco, IdentificationNewParamsIDTypeSkRc, IdentificationNewParamsIDTypeUsEin, IdentificationNewParamsIDTypeUsItin, IdentificationNewParamsIDTypeUsSsn, IdentificationNewParamsIDTypeUyRut, IdentificationNewParamsIDTypeVnTin:
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
	IdentificationUpdateParamsIDTypeAdNrt                IdentificationUpdateParamsIDType = "ad_nrt"
	IdentificationUpdateParamsIDTypeAeEid                IdentificationUpdateParamsIDType = "ae_eid"
	IdentificationUpdateParamsIDTypeAeTrn                IdentificationUpdateParamsIDType = "ae_trn"
	IdentificationUpdateParamsIDTypeAgTin                IdentificationUpdateParamsIDType = "ag_tin"
	IdentificationUpdateParamsIDTypeAITin                IdentificationUpdateParamsIDType = "ai_tin"
	IdentificationUpdateParamsIDTypeAlNid                IdentificationUpdateParamsIDType = "al_nid"
	IdentificationUpdateParamsIDTypeAlNipt               IdentificationUpdateParamsIDType = "al_nipt"
	IdentificationUpdateParamsIDTypeAmTin                IdentificationUpdateParamsIDType = "am_tin"
	IdentificationUpdateParamsIDTypeAoNif                IdentificationUpdateParamsIDType = "ao_nif"
	IdentificationUpdateParamsIDTypeArCuil               IdentificationUpdateParamsIDType = "ar_cuil"
	IdentificationUpdateParamsIDTypeArCuit               IdentificationUpdateParamsIDType = "ar_cuit"
	IdentificationUpdateParamsIDTypeAtAtin               IdentificationUpdateParamsIDType = "at_atin"
	IdentificationUpdateParamsIDTypeAtVat                IdentificationUpdateParamsIDType = "at_vat"
	IdentificationUpdateParamsIDTypeAuAbn                IdentificationUpdateParamsIDType = "au_abn"
	IdentificationUpdateParamsIDTypeAuTfn                IdentificationUpdateParamsIDType = "au_tfn"
	IdentificationUpdateParamsIDTypeAwTin                IdentificationUpdateParamsIDType = "aw_tin"
	IdentificationUpdateParamsIDTypeAzPin                IdentificationUpdateParamsIDType = "az_pin"
	IdentificationUpdateParamsIDTypeBbTin                IdentificationUpdateParamsIDType = "bb_tin"
	IdentificationUpdateParamsIDTypeBdTin                IdentificationUpdateParamsIDType = "bd_tin"
	IdentificationUpdateParamsIDTypeBeEnt                IdentificationUpdateParamsIDType = "be_ent"
	IdentificationUpdateParamsIDTypeBeNrn                IdentificationUpdateParamsIDType = "be_nrn"
	IdentificationUpdateParamsIDTypeBfIfu                IdentificationUpdateParamsIDType = "bf_ifu"
	IdentificationUpdateParamsIDTypeBgEgn                IdentificationUpdateParamsIDType = "bg_egn"
	IdentificationUpdateParamsIDTypeBhCpr                IdentificationUpdateParamsIDType = "bh_cpr"
	IdentificationUpdateParamsIDTypeBhVat                IdentificationUpdateParamsIDType = "bh_vat"
	IdentificationUpdateParamsIDTypeBjIfu                IdentificationUpdateParamsIDType = "bj_ifu"
	IdentificationUpdateParamsIDTypeBoNit                IdentificationUpdateParamsIDType = "bo_nit"
	IdentificationUpdateParamsIDTypeBrCnpj               IdentificationUpdateParamsIDType = "br_cnpj"
	IdentificationUpdateParamsIDTypeBrCpf                IdentificationUpdateParamsIDType = "br_cpf"
	IdentificationUpdateParamsIDTypeBsTin                IdentificationUpdateParamsIDType = "bs_tin"
	IdentificationUpdateParamsIDTypeBtBin                IdentificationUpdateParamsIDType = "bt_bin"
	IdentificationUpdateParamsIDTypeBwTin                IdentificationUpdateParamsIDType = "bw_tin"
	IdentificationUpdateParamsIDTypeBzTin                IdentificationUpdateParamsIDType = "bz_tin"
	IdentificationUpdateParamsIDTypeCaBn                 IdentificationUpdateParamsIDType = "ca_bn"
	IdentificationUpdateParamsIDTypeCaSin                IdentificationUpdateParamsIDType = "ca_sin"
	IdentificationUpdateParamsIDTypeChAhv                IdentificationUpdateParamsIDType = "ch_ahv"
	IdentificationUpdateParamsIDTypeChUid                IdentificationUpdateParamsIDType = "ch_uid"
	IdentificationUpdateParamsIDTypeCiNcc                IdentificationUpdateParamsIDType = "ci_ncc"
	IdentificationUpdateParamsIDTypeClRun                IdentificationUpdateParamsIDType = "cl_run"
	IdentificationUpdateParamsIDTypeClRut                IdentificationUpdateParamsIDType = "cl_rut"
	IdentificationUpdateParamsIDTypeCmNiu                IdentificationUpdateParamsIDType = "cm_niu"
	IdentificationUpdateParamsIDTypeCoCedulas            IdentificationUpdateParamsIDType = "co_cedulas"
	IdentificationUpdateParamsIDTypeCoNit                IdentificationUpdateParamsIDType = "co_nit"
	IdentificationUpdateParamsIDTypeCrCpf                IdentificationUpdateParamsIDType = "cr_cpf"
	IdentificationUpdateParamsIDTypeCwCrib               IdentificationUpdateParamsIDType = "cw_crib"
	IdentificationUpdateParamsIDTypeCyTin                IdentificationUpdateParamsIDType = "cy_tin"
	IdentificationUpdateParamsIDTypeCzIco                IdentificationUpdateParamsIDType = "cz_ico"
	IdentificationUpdateParamsIDTypeCzRc                 IdentificationUpdateParamsIDType = "cz_rc"
	IdentificationUpdateParamsIDTypeDeStid               IdentificationUpdateParamsIDType = "de_stid"
	IdentificationUpdateParamsIDTypeDeStnr               IdentificationUpdateParamsIDType = "de_stnr"
	IdentificationUpdateParamsIDTypeDeVat                IdentificationUpdateParamsIDType = "de_vat"
	IdentificationUpdateParamsIDTypeDkCpr                IdentificationUpdateParamsIDType = "dk_cpr"
	IdentificationUpdateParamsIDTypeDkCvr                IdentificationUpdateParamsIDType = "dk_cvr"
	IdentificationUpdateParamsIDTypeDmTin                IdentificationUpdateParamsIDType = "dm_tin"
	IdentificationUpdateParamsIDTypeDoCedula             IdentificationUpdateParamsIDType = "do_cedula"
	IdentificationUpdateParamsIDTypeDoRnc                IdentificationUpdateParamsIDType = "do_rnc"
	IdentificationUpdateParamsIDTypeDriversLicense       IdentificationUpdateParamsIDType = "drivers_license"
	IdentificationUpdateParamsIDTypeEcRuc                IdentificationUpdateParamsIDType = "ec_ruc"
	IdentificationUpdateParamsIDTypeEeIk                 IdentificationUpdateParamsIDType = "ee_ik"
	IdentificationUpdateParamsIDTypeEeRk                 IdentificationUpdateParamsIDType = "ee_rk"
	IdentificationUpdateParamsIDTypeEsNie                IdentificationUpdateParamsIDType = "es_nie"
	IdentificationUpdateParamsIDTypeEsNif                IdentificationUpdateParamsIDType = "es_nif"
	IdentificationUpdateParamsIDTypeFiHetu               IdentificationUpdateParamsIDType = "fi_hetu"
	IdentificationUpdateParamsIDTypeFiYtj                IdentificationUpdateParamsIDType = "fi_ytj"
	IdentificationUpdateParamsIDTypeFjTin                IdentificationUpdateParamsIDType = "fj_tin"
	IdentificationUpdateParamsIDTypeFoPtal               IdentificationUpdateParamsIDType = "fo_ptal"
	IdentificationUpdateParamsIDTypeFrNif                IdentificationUpdateParamsIDType = "fr_nif"
	IdentificationUpdateParamsIDTypeFrSiren              IdentificationUpdateParamsIDType = "fr_siren"
	IdentificationUpdateParamsIDTypeFrVat                IdentificationUpdateParamsIDType = "fr_vat"
	IdentificationUpdateParamsIDTypeGBNino               IdentificationUpdateParamsIDType = "gb_nino"
	IdentificationUpdateParamsIDTypeGBUtr                IdentificationUpdateParamsIDType = "gb_utr"
	IdentificationUpdateParamsIDTypeGBVat                IdentificationUpdateParamsIDType = "gb_vat"
	IdentificationUpdateParamsIDTypeGdTin                IdentificationUpdateParamsIDType = "gd_tin"
	IdentificationUpdateParamsIDTypeGeIc                 IdentificationUpdateParamsIDType = "ge_ic"
	IdentificationUpdateParamsIDTypeGePn                 IdentificationUpdateParamsIDType = "ge_pn"
	IdentificationUpdateParamsIDTypeGeTin                IdentificationUpdateParamsIDType = "ge_tin"
	IdentificationUpdateParamsIDTypeGenericInternational IdentificationUpdateParamsIDType = "generic_international"
	IdentificationUpdateParamsIDTypeGgSin                IdentificationUpdateParamsIDType = "gg_sin"
	IdentificationUpdateParamsIDTypeGhPin                IdentificationUpdateParamsIDType = "gh_pin"
	IdentificationUpdateParamsIDTypeGhTin                IdentificationUpdateParamsIDType = "gh_tin"
	IdentificationUpdateParamsIDTypeGiTrn                IdentificationUpdateParamsIDType = "gi_trn"
	IdentificationUpdateParamsIDTypeGlCpr                IdentificationUpdateParamsIDType = "gl_cpr"
	IdentificationUpdateParamsIDTypeGlGer                IdentificationUpdateParamsIDType = "gl_ger"
	IdentificationUpdateParamsIDTypeGmTin                IdentificationUpdateParamsIDType = "gm_tin"
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
	case IdentificationUpdateParamsIDTypeAdNrt, IdentificationUpdateParamsIDTypeAeEid, IdentificationUpdateParamsIDTypeAeTrn, IdentificationUpdateParamsIDTypeAgTin, IdentificationUpdateParamsIDTypeAITin, IdentificationUpdateParamsIDTypeAlNid, IdentificationUpdateParamsIDTypeAlNipt, IdentificationUpdateParamsIDTypeAmTin, IdentificationUpdateParamsIDTypeAoNif, IdentificationUpdateParamsIDTypeArCuil, IdentificationUpdateParamsIDTypeArCuit, IdentificationUpdateParamsIDTypeAtAtin, IdentificationUpdateParamsIDTypeAtVat, IdentificationUpdateParamsIDTypeAuAbn, IdentificationUpdateParamsIDTypeAuTfn, IdentificationUpdateParamsIDTypeAwTin, IdentificationUpdateParamsIDTypeAzPin, IdentificationUpdateParamsIDTypeBbTin, IdentificationUpdateParamsIDTypeBdTin, IdentificationUpdateParamsIDTypeBeEnt, IdentificationUpdateParamsIDTypeBeNrn, IdentificationUpdateParamsIDTypeBfIfu, IdentificationUpdateParamsIDTypeBgEgn, IdentificationUpdateParamsIDTypeBhCpr, IdentificationUpdateParamsIDTypeBhVat, IdentificationUpdateParamsIDTypeBjIfu, IdentificationUpdateParamsIDTypeBoNit, IdentificationUpdateParamsIDTypeBrCnpj, IdentificationUpdateParamsIDTypeBrCpf, IdentificationUpdateParamsIDTypeBsTin, IdentificationUpdateParamsIDTypeBtBin, IdentificationUpdateParamsIDTypeBwTin, IdentificationUpdateParamsIDTypeBzTin, IdentificationUpdateParamsIDTypeCaBn, IdentificationUpdateParamsIDTypeCaSin, IdentificationUpdateParamsIDTypeChAhv, IdentificationUpdateParamsIDTypeChUid, IdentificationUpdateParamsIDTypeCiNcc, IdentificationUpdateParamsIDTypeClRun, IdentificationUpdateParamsIDTypeClRut, IdentificationUpdateParamsIDTypeCmNiu, IdentificationUpdateParamsIDTypeCoCedulas, IdentificationUpdateParamsIDTypeCoNit, IdentificationUpdateParamsIDTypeCrCpf, IdentificationUpdateParamsIDTypeCwCrib, IdentificationUpdateParamsIDTypeCyTin, IdentificationUpdateParamsIDTypeCzIco, IdentificationUpdateParamsIDTypeCzRc, IdentificationUpdateParamsIDTypeDeStid, IdentificationUpdateParamsIDTypeDeStnr, IdentificationUpdateParamsIDTypeDeVat, IdentificationUpdateParamsIDTypeDkCpr, IdentificationUpdateParamsIDTypeDkCvr, IdentificationUpdateParamsIDTypeDmTin, IdentificationUpdateParamsIDTypeDoCedula, IdentificationUpdateParamsIDTypeDoRnc, IdentificationUpdateParamsIDTypeDriversLicense, IdentificationUpdateParamsIDTypeEcRuc, IdentificationUpdateParamsIDTypeEeIk, IdentificationUpdateParamsIDTypeEeRk, IdentificationUpdateParamsIDTypeEsNie, IdentificationUpdateParamsIDTypeEsNif, IdentificationUpdateParamsIDTypeFiHetu, IdentificationUpdateParamsIDTypeFiYtj, IdentificationUpdateParamsIDTypeFjTin, IdentificationUpdateParamsIDTypeFoPtal, IdentificationUpdateParamsIDTypeFrNif, IdentificationUpdateParamsIDTypeFrSiren, IdentificationUpdateParamsIDTypeFrVat, IdentificationUpdateParamsIDTypeGBNino, IdentificationUpdateParamsIDTypeGBUtr, IdentificationUpdateParamsIDTypeGBVat, IdentificationUpdateParamsIDTypeGdTin, IdentificationUpdateParamsIDTypeGeIc, IdentificationUpdateParamsIDTypeGePn, IdentificationUpdateParamsIDTypeGeTin, IdentificationUpdateParamsIDTypeGenericInternational, IdentificationUpdateParamsIDTypeGgSin, IdentificationUpdateParamsIDTypeGhPin, IdentificationUpdateParamsIDTypeGhTin, IdentificationUpdateParamsIDTypeGiTrn, IdentificationUpdateParamsIDTypeGlCpr, IdentificationUpdateParamsIDTypeGlGer, IdentificationUpdateParamsIDTypeGmTin, IdentificationUpdateParamsIDTypeGrVat, IdentificationUpdateParamsIDTypeHkBrn, IdentificationUpdateParamsIDTypeHkHkid, IdentificationUpdateParamsIDTypeHnID, IdentificationUpdateParamsIDTypeHnRtn, IdentificationUpdateParamsIDTypeHrOib, IdentificationUpdateParamsIDTypeHuAdj, IdentificationUpdateParamsIDTypeHuAnum, IdentificationUpdateParamsIDTypeIePps, IdentificationUpdateParamsIDTypeIeTrn, IdentificationUpdateParamsIDTypeInLei, IdentificationUpdateParamsIDTypeIsKnt, IdentificationUpdateParamsIDTypeItCf, IdentificationUpdateParamsIDTypeItPiva, IdentificationUpdateParamsIDTypeJpHb, IdentificationUpdateParamsIDTypeJpMn, IdentificationUpdateParamsIDTypeKrBrn, IdentificationUpdateParamsIDTypeKrCrn, IdentificationUpdateParamsIDTypeKrRrn, IdentificationUpdateParamsIDTypeLiPeid, IdentificationUpdateParamsIDTypeLtAk, IdentificationUpdateParamsIDTypeLtJak, IdentificationUpdateParamsIDTypeLuMtc, IdentificationUpdateParamsIDTypeLuVat, IdentificationUpdateParamsIDTypeLvPk, IdentificationUpdateParamsIDTypeLvRn, IdentificationUpdateParamsIDTypeMtTin, IdentificationUpdateParamsIDTypeMtVat, IdentificationUpdateParamsIDTypeMxCurp, IdentificationUpdateParamsIDTypeMxIne, IdentificationUpdateParamsIDTypeMxRfc, IdentificationUpdateParamsIDTypeNationalID, IdentificationUpdateParamsIDTypeNlBsn, IdentificationUpdateParamsIDTypeNlBtw, IdentificationUpdateParamsIDTypeNlRsin, IdentificationUpdateParamsIDTypeNoFdn, IdentificationUpdateParamsIDTypeNoMva, IdentificationUpdateParamsIDTypeNoOrgnr, IdentificationUpdateParamsIDTypeNzIrd, IdentificationUpdateParamsIDTypePassport, IdentificationUpdateParamsIDTypePlNip, IdentificationUpdateParamsIDTypePlPesel, IdentificationUpdateParamsIDTypePtNif, IdentificationUpdateParamsIDTypeRoCnp, IdentificationUpdateParamsIDTypeRoCui, IdentificationUpdateParamsIDTypeSaTin, IdentificationUpdateParamsIDTypeSaVat, IdentificationUpdateParamsIDTypeSeOrgnr, IdentificationUpdateParamsIDTypeSePnmr, IdentificationUpdateParamsIDTypeSgFin, IdentificationUpdateParamsIDTypeSgNric, IdentificationUpdateParamsIDTypeSgUen, IdentificationUpdateParamsIDTypeSiDav, IdentificationUpdateParamsIDTypeSiTin, IdentificationUpdateParamsIDTypeSkIco, IdentificationUpdateParamsIDTypeSkRc, IdentificationUpdateParamsIDTypeUsEin, IdentificationUpdateParamsIDTypeUsItin, IdentificationUpdateParamsIDTypeUsSsn, IdentificationUpdateParamsIDTypeUyRut, IdentificationUpdateParamsIDTypeVnTin:
		return true
	}
	return false
}
