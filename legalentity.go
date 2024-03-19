// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// LegalEntityService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the [NewLegalEntityService]
// method instead.
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
	path := fmt.Sprintf("api/legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a legal entity.
func (r *LegalEntityService) Update(ctx context.Context, id string, body LegalEntityUpdateParams, opts ...option.RequestOption) (res *LegalEntity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all legal entities.
func (r *LegalEntityService) List(ctx context.Context, query LegalEntityListParams, opts ...option.RequestOption) (res *shared.Page[LegalEntity], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *LegalEntityService) ListAutoPaging(ctx context.Context, query LegalEntityListParams, opts ...option.RequestOption) *shared.PageAutoPager[LegalEntity] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type LegalEntity struct {
	ID string `json:"id" format:"uuid"`
	// A list of addresses for the entity.
	Addresses []LegalEntityAddress `json:"addresses"`
	// The business's legal business name.
	BusinessName string    `json:"business_name,nullable"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed time.Time `json:"date_formed,nullable" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          time.Time `json:"date_of_birth,nullable" format:"date"`
	DiscardedAt          time.Time `json:"discarded_at,nullable" format:"date-time"`
	DoingBusinessAsNames []string  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email string `json:"email,nullable"`
	// An individual's first name.
	FirstName string `json:"first_name,nullable"`
	// A list of identifications for the legal entity.
	Identifications []LegalEntityIdentification `json:"identifications"`
	// An individual's last name.
	LastName string `json:"last_name,nullable"`
	// The legal entity associations and its associated legal entities.
	LegalEntityAssociations []LegalEntityAssociation `json:"legal_entity_associations,nullable"`
	// The type of legal entity.
	LegalEntityType LegalEntityLegalEntityType `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure LegalEntityLegalStructure `json:"legal_structure,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     map[string]string        `json:"metadata"`
	Object       string                   `json:"object"`
	PhoneNumbers []LegalEntityPhoneNumber `json:"phone_numbers"`
	UpdatedAt    time.Time                `json:"updated_at" format:"date-time"`
	// The entity's primary website URL.
	Website string          `json:"website,nullable"`
	JSON    legalEntityJSON `json:"-"`
}

// legalEntityJSON contains the JSON metadata for the struct [LegalEntity]
type legalEntityJSON struct {
	ID                      apijson.Field
	Addresses               apijson.Field
	BusinessName            apijson.Field
	CreatedAt               apijson.Field
	DateFormed              apijson.Field
	DateOfBirth             apijson.Field
	DiscardedAt             apijson.Field
	DoingBusinessAsNames    apijson.Field
	Email                   apijson.Field
	FirstName               apijson.Field
	Identifications         apijson.Field
	LastName                apijson.Field
	LegalEntityAssociations apijson.Field
	LegalEntityType         apijson.Field
	LegalStructure          apijson.Field
	LiveMode                apijson.Field
	Metadata                apijson.Field
	Object                  apijson.Field
	PhoneNumbers            apijson.Field
	UpdatedAt               apijson.Field
	Website                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
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
	// The type of ID number.
	IDType LegalEntityIdentificationsIDType `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country,required,nullable"`
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
	IDType         apijson.Field
	IssuingCountry apijson.Field
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
	LegalEntityIdentificationsIDTypeArCuil    LegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityIdentificationsIDTypeArCuit    LegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityIdentificationsIDTypeBrCnpj    LegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityIdentificationsIDTypeBrCpf     LegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityIdentificationsIDTypeClNut     LegalEntityIdentificationsIDType = "cl_nut"
	LegalEntityIdentificationsIDTypeCoCedulas LegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityIdentificationsIDTypeCoNit     LegalEntityIdentificationsIDType = "co_nit"
	LegalEntityIdentificationsIDTypeHnID      LegalEntityIdentificationsIDType = "hn_id"
	LegalEntityIdentificationsIDTypeHnRtn     LegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityIdentificationsIDTypePassport  LegalEntityIdentificationsIDType = "passport"
	LegalEntityIdentificationsIDTypeUsEin     LegalEntityIdentificationsIDType = "us_ein"
	LegalEntityIdentificationsIDTypeUsItin    LegalEntityIdentificationsIDType = "us_itin"
	LegalEntityIdentificationsIDTypeUsSsn     LegalEntityIdentificationsIDType = "us_ssn"
)

func (r LegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityIdentificationsIDTypeArCuil, LegalEntityIdentificationsIDTypeArCuit, LegalEntityIdentificationsIDTypeBrCnpj, LegalEntityIdentificationsIDTypeBrCpf, LegalEntityIdentificationsIDTypeClNut, LegalEntityIdentificationsIDTypeCoCedulas, LegalEntityIdentificationsIDTypeCoNit, LegalEntityIdentificationsIDTypeHnID, LegalEntityIdentificationsIDTypeHnRtn, LegalEntityIdentificationsIDTypePassport, LegalEntityIdentificationsIDTypeUsEin, LegalEntityIdentificationsIDTypeUsItin, LegalEntityIdentificationsIDTypeUsSsn:
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

type LegalEntityNewParams struct {
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityType] `json:"legal_entity_type,required"`
	// A list of addresses for the entity.
	Addresses param.Field[[]LegalEntityNewParamsAddress] `json:"addresses"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
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
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The legal entity associations and its associated legal entities.
	LegalEntityAssociations param.Field[[]LegalEntityNewParamsLegalEntityAssociation] `json:"legal_entity_associations"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityNewParamsLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                 `json:"metadata"`
	PhoneNumbers param.Field[[]LegalEntityNewParamsPhoneNumber] `json:"phone_numbers"`
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
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r LegalEntityNewParamsIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityNewParamsIdentificationsIDType string

const (
	LegalEntityNewParamsIdentificationsIDTypeArCuil    LegalEntityNewParamsIdentificationsIDType = "ar_cuil"
	LegalEntityNewParamsIdentificationsIDTypeArCuit    LegalEntityNewParamsIdentificationsIDType = "ar_cuit"
	LegalEntityNewParamsIdentificationsIDTypeBrCnpj    LegalEntityNewParamsIdentificationsIDType = "br_cnpj"
	LegalEntityNewParamsIdentificationsIDTypeBrCpf     LegalEntityNewParamsIdentificationsIDType = "br_cpf"
	LegalEntityNewParamsIdentificationsIDTypeClNut     LegalEntityNewParamsIdentificationsIDType = "cl_nut"
	LegalEntityNewParamsIdentificationsIDTypeCoCedulas LegalEntityNewParamsIdentificationsIDType = "co_cedulas"
	LegalEntityNewParamsIdentificationsIDTypeCoNit     LegalEntityNewParamsIdentificationsIDType = "co_nit"
	LegalEntityNewParamsIdentificationsIDTypeHnID      LegalEntityNewParamsIdentificationsIDType = "hn_id"
	LegalEntityNewParamsIdentificationsIDTypeHnRtn     LegalEntityNewParamsIdentificationsIDType = "hn_rtn"
	LegalEntityNewParamsIdentificationsIDTypePassport  LegalEntityNewParamsIdentificationsIDType = "passport"
	LegalEntityNewParamsIdentificationsIDTypeUsEin     LegalEntityNewParamsIdentificationsIDType = "us_ein"
	LegalEntityNewParamsIdentificationsIDTypeUsItin    LegalEntityNewParamsIdentificationsIDType = "us_itin"
	LegalEntityNewParamsIdentificationsIDTypeUsSsn     LegalEntityNewParamsIdentificationsIDType = "us_ssn"
)

func (r LegalEntityNewParamsIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsIdentificationsIDTypeArCuil, LegalEntityNewParamsIdentificationsIDTypeArCuit, LegalEntityNewParamsIdentificationsIDTypeBrCnpj, LegalEntityNewParamsIdentificationsIDTypeBrCpf, LegalEntityNewParamsIdentificationsIDTypeClNut, LegalEntityNewParamsIdentificationsIDTypeCoCedulas, LegalEntityNewParamsIdentificationsIDTypeCoNit, LegalEntityNewParamsIdentificationsIDTypeHnID, LegalEntityNewParamsIdentificationsIDTypeHnRtn, LegalEntityNewParamsIdentificationsIDTypePassport, LegalEntityNewParamsIdentificationsIDTypeUsEin, LegalEntityNewParamsIdentificationsIDTypeUsItin, LegalEntityNewParamsIdentificationsIDTypeUsSsn:
		return true
	}
	return false
}

type LegalEntityNewParamsLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]LegalEntityNewParamsLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The associated legal entity.
	AssociatedLegalEntity param.Field[LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntity] `json:"associated_legal_entity"`
	// The ID of the associated legal entity.
	AssociatedLegalEntityID param.Field[string] `json:"associated_legal_entity_id"`
	// The associated entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the associated entity at the associator entity.
	Title param.Field[string] `json:"title"`
}

func (r LegalEntityNewParamsLegalEntityAssociation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the associated entity relates to associator
// entity.
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

// The associated legal entity.
type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntity struct {
	// A list of addresses for the entity.
	Addresses param.Field[[]LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddress] `json:"addresses"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
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
	Identifications param.Field[[]LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                                                             `json:"metadata"`
	PhoneNumbers param.Field[[]LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddress struct {
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
	AddressTypes param.Field[[]LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                                                 `json:"line2"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType string

const (
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeBusiness    LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "business"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeMailing     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "mailing"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeOther       LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "other"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypePoBox       LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "po_box"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeResidential LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "residential"
)

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeBusiness, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeMailing, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeOther, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypePoBox, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType string

const (
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeArCuil    LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeArCuit    LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeBrCnpj    LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeBrCpf     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeClNut     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "cl_nut"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeCoCedulas LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeCoNit     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "co_nit"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeHnID      LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "hn_id"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeHnRtn     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypePassport  LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "passport"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsEin     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "us_ein"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsItin    LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "us_itin"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsSsn     LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "us_ssn"
)

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeArCuil, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeArCuit, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeBrCnpj, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeBrCpf, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeClNut, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeCoCedulas, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeCoNit, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeHnID, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeHnRtn, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypePassport, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsEin, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsItin, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsSsn:
		return true
	}
	return false
}

// The type of legal entity.
type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityType string

const (
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityTypeBusiness   LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityType = "business"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityTypeIndividual LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityType = "individual"
)

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityTypeBusiness, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure string

const (
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureCorporation        LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "corporation"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureLlc                LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "llc"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureNonProfit          LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "non_profit"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructurePartnership        LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "partnership"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureSoleProprietorship LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "sole_proprietorship"
	LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureTrust              LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "trust"
)

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureCorporation, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureLlc, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureNonProfit, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructurePartnership, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureSoleProprietorship, LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityNewParamsLegalEntityAssociationsAssociatedLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type LegalEntityUpdateParams struct {
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed param.Field[time.Time] `json:"date_formed" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityUpdateParamsLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                    `json:"metadata"`
	PhoneNumbers param.Field[[]LegalEntityUpdateParamsPhoneNumber] `json:"phone_numbers"`
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
