// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"net/http"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// LegalEntityAssociationService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLegalEntityAssociationService] method instead.
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
	opts = append(r.Options[:], opts...)
	path := "api/legal_entity_associations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LegalEntityAssociation struct {
	ID string `json:"id" format:"uuid"`
	// The child legal entity.
	ChildLegalEntity LegalEntityAssociationChildLegalEntity `json:"child_legal_entity"`
	CreatedAt        time.Time                              `json:"created_at" format:"date-time"`
	DiscardedAt      time.Time                              `json:"discarded_at,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode"`
	Object   string `json:"object"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage int64 `json:"ownership_percentage,nullable"`
	// The ID of the parent legal entity. This must be a business or joint legal
	// entity.
	ParentLegalEntityID string                                   `json:"parent_legal_entity_id"`
	RelationshipTypes   []LegalEntityAssociationRelationshipType `json:"relationship_types"`
	// The job title of the child entity at the parent entity.
	Title     string                     `json:"title,nullable"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
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

// The child legal entity.
type LegalEntityAssociationChildLegalEntity struct {
	ID string `json:"id" format:"uuid"`
	// A list of addresses for the entity.
	Addresses []LegalEntityAssociationChildLegalEntityAddress `json:"addresses"`
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
	Identifications []LegalEntityAssociationChildLegalEntityIdentification `json:"identifications"`
	// An individual's last name.
	LastName string `json:"last_name,nullable"`
	// The type of legal entity.
	LegalEntityType LegalEntityAssociationChildLegalEntityLegalEntityType `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure LegalEntityAssociationChildLegalEntityLegalStructure `json:"legal_structure,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     map[string]string                                   `json:"metadata"`
	Object       string                                              `json:"object"`
	PhoneNumbers []LegalEntityAssociationChildLegalEntityPhoneNumber `json:"phone_numbers"`
	UpdatedAt    time.Time                                           `json:"updated_at" format:"date-time"`
	// The entity's primary website URL.
	Website string                                     `json:"website,nullable"`
	JSON    legalEntityAssociationChildLegalEntityJSON `json:"-"`
}

// legalEntityAssociationChildLegalEntityJSON contains the JSON metadata for the
// struct [LegalEntityAssociationChildLegalEntity]
type legalEntityAssociationChildLegalEntityJSON struct {
	ID                   apijson.Field
	Addresses            apijson.Field
	BusinessName         apijson.Field
	CreatedAt            apijson.Field
	DateFormed           apijson.Field
	DateOfBirth          apijson.Field
	DiscardedAt          apijson.Field
	DoingBusinessAsNames apijson.Field
	Email                apijson.Field
	FirstName            apijson.Field
	Identifications      apijson.Field
	LastName             apijson.Field
	LegalEntityType      apijson.Field
	LegalStructure       apijson.Field
	LiveMode             apijson.Field
	Metadata             apijson.Field
	Object               apijson.Field
	PhoneNumbers         apijson.Field
	UpdatedAt            apijson.Field
	Website              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LegalEntityAssociationChildLegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityAssociationChildLegalEntityJSON) RawJSON() string {
	return r.raw
}

type LegalEntityAssociationChildLegalEntityAddress struct {
	ID string `json:"id,required" format:"uuid"`
	// The types of this address.
	AddressTypes []LegalEntityAssociationChildLegalEntityAddressesAddressType `json:"address_types,required"`
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
	Region    string                                            `json:"region,required,nullable"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityAssociationChildLegalEntityAddressJSON `json:"-"`
}

// legalEntityAssociationChildLegalEntityAddressJSON contains the JSON metadata for
// the struct [LegalEntityAssociationChildLegalEntityAddress]
type legalEntityAssociationChildLegalEntityAddressJSON struct {
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

func (r *LegalEntityAssociationChildLegalEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityAssociationChildLegalEntityAddressJSON) RawJSON() string {
	return r.raw
}

type LegalEntityAssociationChildLegalEntityAddressesAddressType string

const (
	LegalEntityAssociationChildLegalEntityAddressesAddressTypeBusiness    LegalEntityAssociationChildLegalEntityAddressesAddressType = "business"
	LegalEntityAssociationChildLegalEntityAddressesAddressTypeMailing     LegalEntityAssociationChildLegalEntityAddressesAddressType = "mailing"
	LegalEntityAssociationChildLegalEntityAddressesAddressTypeOther       LegalEntityAssociationChildLegalEntityAddressesAddressType = "other"
	LegalEntityAssociationChildLegalEntityAddressesAddressTypePoBox       LegalEntityAssociationChildLegalEntityAddressesAddressType = "po_box"
	LegalEntityAssociationChildLegalEntityAddressesAddressTypeResidential LegalEntityAssociationChildLegalEntityAddressesAddressType = "residential"
)

func (r LegalEntityAssociationChildLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationChildLegalEntityAddressesAddressTypeBusiness, LegalEntityAssociationChildLegalEntityAddressesAddressTypeMailing, LegalEntityAssociationChildLegalEntityAddressesAddressTypeOther, LegalEntityAssociationChildLegalEntityAddressesAddressTypePoBox, LegalEntityAssociationChildLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityAssociationChildLegalEntityIdentification struct {
	ID          string    `json:"id,required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The type of ID number.
	IDType LegalEntityAssociationChildLegalEntityIdentificationsIDType `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                                                     `json:"live_mode,required"`
	Object    string                                                   `json:"object,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityAssociationChildLegalEntityIdentificationJSON `json:"-"`
}

// legalEntityAssociationChildLegalEntityIdentificationJSON contains the JSON
// metadata for the struct [LegalEntityAssociationChildLegalEntityIdentification]
type legalEntityAssociationChildLegalEntityIdentificationJSON struct {
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

func (r *LegalEntityAssociationChildLegalEntityIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityAssociationChildLegalEntityIdentificationJSON) RawJSON() string {
	return r.raw
}

// The type of ID number.
type LegalEntityAssociationChildLegalEntityIdentificationsIDType string

const (
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeArCuil    LegalEntityAssociationChildLegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeArCuit    LegalEntityAssociationChildLegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeBrCnpj    LegalEntityAssociationChildLegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeBrCpf     LegalEntityAssociationChildLegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeClRut     LegalEntityAssociationChildLegalEntityIdentificationsIDType = "cl_rut"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeCoCedulas LegalEntityAssociationChildLegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeCoNit     LegalEntityAssociationChildLegalEntityIdentificationsIDType = "co_nit"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeHnID      LegalEntityAssociationChildLegalEntityIdentificationsIDType = "hn_id"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeHnRtn     LegalEntityAssociationChildLegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypePassport  LegalEntityAssociationChildLegalEntityIdentificationsIDType = "passport"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeUsEin     LegalEntityAssociationChildLegalEntityIdentificationsIDType = "us_ein"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeUsItin    LegalEntityAssociationChildLegalEntityIdentificationsIDType = "us_itin"
	LegalEntityAssociationChildLegalEntityIdentificationsIDTypeUsSsn     LegalEntityAssociationChildLegalEntityIdentificationsIDType = "us_ssn"
)

func (r LegalEntityAssociationChildLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationChildLegalEntityIdentificationsIDTypeArCuil, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeArCuit, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeBrCnpj, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeBrCpf, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeClRut, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeCoCedulas, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeCoNit, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeHnID, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeHnRtn, LegalEntityAssociationChildLegalEntityIdentificationsIDTypePassport, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeUsEin, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeUsItin, LegalEntityAssociationChildLegalEntityIdentificationsIDTypeUsSsn:
		return true
	}
	return false
}

// The type of legal entity.
type LegalEntityAssociationChildLegalEntityLegalEntityType string

const (
	LegalEntityAssociationChildLegalEntityLegalEntityTypeBusiness   LegalEntityAssociationChildLegalEntityLegalEntityType = "business"
	LegalEntityAssociationChildLegalEntityLegalEntityTypeIndividual LegalEntityAssociationChildLegalEntityLegalEntityType = "individual"
	LegalEntityAssociationChildLegalEntityLegalEntityTypeJoint      LegalEntityAssociationChildLegalEntityLegalEntityType = "joint"
)

func (r LegalEntityAssociationChildLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationChildLegalEntityLegalEntityTypeBusiness, LegalEntityAssociationChildLegalEntityLegalEntityTypeIndividual, LegalEntityAssociationChildLegalEntityLegalEntityTypeJoint:
		return true
	}
	return false
}

// The business's legal structure.
type LegalEntityAssociationChildLegalEntityLegalStructure string

const (
	LegalEntityAssociationChildLegalEntityLegalStructureCorporation        LegalEntityAssociationChildLegalEntityLegalStructure = "corporation"
	LegalEntityAssociationChildLegalEntityLegalStructureLlc                LegalEntityAssociationChildLegalEntityLegalStructure = "llc"
	LegalEntityAssociationChildLegalEntityLegalStructureNonProfit          LegalEntityAssociationChildLegalEntityLegalStructure = "non_profit"
	LegalEntityAssociationChildLegalEntityLegalStructurePartnership        LegalEntityAssociationChildLegalEntityLegalStructure = "partnership"
	LegalEntityAssociationChildLegalEntityLegalStructureSoleProprietorship LegalEntityAssociationChildLegalEntityLegalStructure = "sole_proprietorship"
	LegalEntityAssociationChildLegalEntityLegalStructureTrust              LegalEntityAssociationChildLegalEntityLegalStructure = "trust"
)

func (r LegalEntityAssociationChildLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityAssociationChildLegalEntityLegalStructureCorporation, LegalEntityAssociationChildLegalEntityLegalStructureLlc, LegalEntityAssociationChildLegalEntityLegalStructureNonProfit, LegalEntityAssociationChildLegalEntityLegalStructurePartnership, LegalEntityAssociationChildLegalEntityLegalStructureSoleProprietorship, LegalEntityAssociationChildLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityAssociationChildLegalEntityPhoneNumber struct {
	PhoneNumber string                                                `json:"phone_number"`
	JSON        legalEntityAssociationChildLegalEntityPhoneNumberJSON `json:"-"`
}

// legalEntityAssociationChildLegalEntityPhoneNumberJSON contains the JSON metadata
// for the struct [LegalEntityAssociationChildLegalEntityPhoneNumber]
type legalEntityAssociationChildLegalEntityPhoneNumberJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalEntityAssociationChildLegalEntityPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityAssociationChildLegalEntityPhoneNumberJSON) RawJSON() string {
	return r.raw
}

// A list of relationship types for how the child entity relates to parent entity.
type LegalEntityAssociationRelationshipType string

const (
	LegalEntityAssociationRelationshipTypeBeneficialOwner LegalEntityAssociationRelationshipType = "beneficial_owner"
	LegalEntityAssociationRelationshipTypeControlPerson   LegalEntityAssociationRelationshipType = "control_person"
)

func (r LegalEntityAssociationRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationRelationshipTypeBeneficialOwner, LegalEntityAssociationRelationshipTypeControlPerson:
		return true
	}
	return false
}

type LegalEntityAssociationNewParams struct {
	RelationshipTypes param.Field[[]LegalEntityAssociationNewParamsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[LegalEntityAssociationNewParamsChildLegalEntity] `json:"child_legal_entity"`
	// The ID of the child legal entity.
	ChildLegalEntityID param.Field[string] `json:"child_legal_entity_id"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The ID of the parent legal entity. This must be a business or joint legal
	// entity.
	ParentLegalEntityID param.Field[string] `json:"parent_legal_entity_id"`
	// The job title of the child entity at the parent entity.
	Title param.Field[string] `json:"title"`
}

func (r LegalEntityAssociationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the child entity relates to parent entity.
type LegalEntityAssociationNewParamsRelationshipType string

const (
	LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner LegalEntityAssociationNewParamsRelationshipType = "beneficial_owner"
	LegalEntityAssociationNewParamsRelationshipTypeControlPerson   LegalEntityAssociationNewParamsRelationshipType = "control_person"
)

func (r LegalEntityAssociationNewParamsRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner, LegalEntityAssociationNewParamsRelationshipTypeControlPerson:
		return true
	}
	return false
}

// The child legal entity.
type LegalEntityAssociationNewParamsChildLegalEntity struct {
	// A list of addresses for the entity.
	Addresses param.Field[[]LegalEntityAssociationNewParamsChildLegalEntityAddress] `json:"addresses"`
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
	Identifications param.Field[[]LegalEntityAssociationNewParamsChildLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityAssociationNewParamsChildLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[LegalEntityAssociationNewParamsChildLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                                            `json:"metadata"`
	PhoneNumbers param.Field[[]LegalEntityAssociationNewParamsChildLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityAssociationNewParamsChildLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityAssociationNewParamsChildLegalEntityAddress struct {
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
	AddressTypes param.Field[[]LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                                `json:"line2"`
}

func (r LegalEntityAssociationNewParamsChildLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType string

const (
	LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness    LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType = "business"
	LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing     LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType = "mailing"
	LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther       LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType = "other"
	LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypePoBox       LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType = "po_box"
	LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeResidential LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType = "residential"
)

func (r LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther, LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypePoBox, LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityAssociationNewParamsChildLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r LegalEntityAssociationNewParamsChildLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType string

const (
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuil    LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuit    LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeBrCnpj    LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeBrCpf     LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeClRut     LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "cl_rut"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeCoCedulas LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeCoNit     LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "co_nit"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeHnID      LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "hn_id"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeHnRtn     LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypePassport  LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "passport"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeUsEin     LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "us_ein"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeUsItin    LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "us_itin"
	LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeUsSsn     LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType = "us_ssn"
)

func (r LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuil, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuit, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeBrCnpj, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeBrCpf, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeClRut, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeCoCedulas, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeCoNit, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeHnID, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeHnRtn, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypePassport, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeUsEin, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeUsItin, LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeUsSsn:
		return true
	}
	return false
}

// The type of legal entity.
type LegalEntityAssociationNewParamsChildLegalEntityLegalEntityType string

const (
	LegalEntityAssociationNewParamsChildLegalEntityLegalEntityTypeBusiness   LegalEntityAssociationNewParamsChildLegalEntityLegalEntityType = "business"
	LegalEntityAssociationNewParamsChildLegalEntityLegalEntityTypeIndividual LegalEntityAssociationNewParamsChildLegalEntityLegalEntityType = "individual"
)

func (r LegalEntityAssociationNewParamsChildLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsChildLegalEntityLegalEntityTypeBusiness, LegalEntityAssociationNewParamsChildLegalEntityLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type LegalEntityAssociationNewParamsChildLegalEntityLegalStructure string

const (
	LegalEntityAssociationNewParamsChildLegalEntityLegalStructureCorporation        LegalEntityAssociationNewParamsChildLegalEntityLegalStructure = "corporation"
	LegalEntityAssociationNewParamsChildLegalEntityLegalStructureLlc                LegalEntityAssociationNewParamsChildLegalEntityLegalStructure = "llc"
	LegalEntityAssociationNewParamsChildLegalEntityLegalStructureNonProfit          LegalEntityAssociationNewParamsChildLegalEntityLegalStructure = "non_profit"
	LegalEntityAssociationNewParamsChildLegalEntityLegalStructurePartnership        LegalEntityAssociationNewParamsChildLegalEntityLegalStructure = "partnership"
	LegalEntityAssociationNewParamsChildLegalEntityLegalStructureSoleProprietorship LegalEntityAssociationNewParamsChildLegalEntityLegalStructure = "sole_proprietorship"
	LegalEntityAssociationNewParamsChildLegalEntityLegalStructureTrust              LegalEntityAssociationNewParamsChildLegalEntityLegalStructure = "trust"
)

func (r LegalEntityAssociationNewParamsChildLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsChildLegalEntityLegalStructureCorporation, LegalEntityAssociationNewParamsChildLegalEntityLegalStructureLlc, LegalEntityAssociationNewParamsChildLegalEntityLegalStructureNonProfit, LegalEntityAssociationNewParamsChildLegalEntityLegalStructurePartnership, LegalEntityAssociationNewParamsChildLegalEntityLegalStructureSoleProprietorship, LegalEntityAssociationNewParamsChildLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type LegalEntityAssociationNewParamsChildLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityAssociationNewParamsChildLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
