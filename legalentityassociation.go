// File generated from our OpenAPI spec by Stainless.

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
	// The associated legal entity.
	AssociatedLegalEntity LegalEntityAssociationAssociatedLegalEntity `json:"associated_legal_entity"`
	// The ID of the associator legal entity. This must be a business or joint legal
	// entity.
	AssociatorLegalEntityID string    `json:"associator_legal_entity_id"`
	CreatedAt               time.Time `json:"created_at" format:"date-time"`
	DiscardedAt             time.Time `json:"discarded_at,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode"`
	Object   string `json:"object"`
	// The associated entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage int64                                    `json:"ownership_percentage,nullable"`
	RelationshipTypes   []LegalEntityAssociationRelationshipType `json:"relationship_types"`
	// The job title of the associated entity at the associator entity.
	Title     string                     `json:"title,nullable"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      legalEntityAssociationJSON `json:"-"`
}

// legalEntityAssociationJSON contains the JSON metadata for the struct
// [LegalEntityAssociation]
type legalEntityAssociationJSON struct {
	ID                      apijson.Field
	AssociatedLegalEntity   apijson.Field
	AssociatorLegalEntityID apijson.Field
	CreatedAt               apijson.Field
	DiscardedAt             apijson.Field
	LiveMode                apijson.Field
	Object                  apijson.Field
	OwnershipPercentage     apijson.Field
	RelationshipTypes       apijson.Field
	Title                   apijson.Field
	UpdatedAt               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LegalEntityAssociation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The associated legal entity.
type LegalEntityAssociationAssociatedLegalEntity struct {
	ID string `json:"id" format:"uuid"`
	// A list of addresses for the entity.
	Addresses []LegalEntityAssociationAssociatedLegalEntityAddress `json:"addresses"`
	// The business's legal business name.
	BusinessName string    `json:"business_name,nullable"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// An individual's data of birth (YYYY-MM-DD).
	DateOfBirth          time.Time `json:"date_of_birth,nullable" format:"date"`
	DiscardedAt          time.Time `json:"discarded_at,nullable" format:"date-time"`
	DoingBusinessAsNames []string  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email string `json:"email,nullable"`
	// An individual's first name.
	FirstName string `json:"first_name,nullable"`
	// A list of identifications for the legal entity.
	Identifications []LegalEntityAssociationAssociatedLegalEntityIdentification `json:"identifications"`
	// An individual's last name.
	LastName string `json:"last_name,nullable"`
	// The type of legal entity.
	LegalEntityType LegalEntityAssociationAssociatedLegalEntityLegalEntityType `json:"legal_entity_type"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     map[string]string                                        `json:"metadata"`
	Object       string                                                   `json:"object"`
	PhoneNumbers []LegalEntityAssociationAssociatedLegalEntityPhoneNumber `json:"phone_numbers"`
	UpdatedAt    time.Time                                                `json:"updated_at" format:"date-time"`
	// The entity's primary website URL.
	Website string                                          `json:"website,nullable"`
	JSON    legalEntityAssociationAssociatedLegalEntityJSON `json:"-"`
}

// legalEntityAssociationAssociatedLegalEntityJSON contains the JSON metadata for
// the struct [LegalEntityAssociationAssociatedLegalEntity]
type legalEntityAssociationAssociatedLegalEntityJSON struct {
	ID                   apijson.Field
	Addresses            apijson.Field
	BusinessName         apijson.Field
	CreatedAt            apijson.Field
	DateOfBirth          apijson.Field
	DiscardedAt          apijson.Field
	DoingBusinessAsNames apijson.Field
	Email                apijson.Field
	FirstName            apijson.Field
	Identifications      apijson.Field
	LastName             apijson.Field
	LegalEntityType      apijson.Field
	LiveMode             apijson.Field
	Metadata             apijson.Field
	Object               apijson.Field
	PhoneNumbers         apijson.Field
	UpdatedAt            apijson.Field
	Website              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LegalEntityAssociationAssociatedLegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LegalEntityAssociationAssociatedLegalEntityAddress struct {
	ID string `json:"id,required" format:"uuid"`
	// The types of this address.
	AddressTypes []string `json:"address_types,required"`
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
	Region    string                                                 `json:"region,required,nullable"`
	UpdatedAt time.Time                                              `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityAssociationAssociatedLegalEntityAddressJSON `json:"-"`
}

// legalEntityAssociationAssociatedLegalEntityAddressJSON contains the JSON
// metadata for the struct [LegalEntityAssociationAssociatedLegalEntityAddress]
type legalEntityAssociationAssociatedLegalEntityAddressJSON struct {
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

func (r *LegalEntityAssociationAssociatedLegalEntityAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LegalEntityAssociationAssociatedLegalEntityIdentification struct {
	ID          string    `json:"id,required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The type of ID number.
	IDType LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry string `json:"issuing_country,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                                                          `json:"live_mode,required"`
	Object    string                                                        `json:"object,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityAssociationAssociatedLegalEntityIdentificationJSON `json:"-"`
}

// legalEntityAssociationAssociatedLegalEntityIdentificationJSON contains the JSON
// metadata for the struct
// [LegalEntityAssociationAssociatedLegalEntityIdentification]
type legalEntityAssociationAssociatedLegalEntityIdentificationJSON struct {
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

func (r *LegalEntityAssociationAssociatedLegalEntityIdentification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of ID number.
type LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType string

const (
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeArCuil    LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeArCuit    LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeBrCnpj    LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeBrCpf     LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeClNut     LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "cl_nut"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeCoCedulas LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeCoNit     LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "co_nit"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeHnID      LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "hn_id"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeHnRtn     LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypePassport  LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "passport"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeUsEin     LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "us_ein"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeUsItin    LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "us_itin"
	LegalEntityAssociationAssociatedLegalEntityIdentificationsIDTypeUsSsn     LegalEntityAssociationAssociatedLegalEntityIdentificationsIDType = "us_ssn"
)

// The type of legal entity.
type LegalEntityAssociationAssociatedLegalEntityLegalEntityType string

const (
	LegalEntityAssociationAssociatedLegalEntityLegalEntityTypeBusiness   LegalEntityAssociationAssociatedLegalEntityLegalEntityType = "business"
	LegalEntityAssociationAssociatedLegalEntityLegalEntityTypeIndividual LegalEntityAssociationAssociatedLegalEntityLegalEntityType = "individual"
	LegalEntityAssociationAssociatedLegalEntityLegalEntityTypeJoint      LegalEntityAssociationAssociatedLegalEntityLegalEntityType = "joint"
)

// A list of phone numbers in E.164 format.
type LegalEntityAssociationAssociatedLegalEntityPhoneNumber struct {
	PhoneNumber string                                                     `json:"phone_number"`
	JSON        legalEntityAssociationAssociatedLegalEntityPhoneNumberJSON `json:"-"`
}

// legalEntityAssociationAssociatedLegalEntityPhoneNumberJSON contains the JSON
// metadata for the struct [LegalEntityAssociationAssociatedLegalEntityPhoneNumber]
type legalEntityAssociationAssociatedLegalEntityPhoneNumberJSON struct {
	PhoneNumber apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegalEntityAssociationAssociatedLegalEntityPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of relationship types for how the associated entity relates to associator
// entity.
type LegalEntityAssociationRelationshipType string

const (
	LegalEntityAssociationRelationshipTypeBeneficialOwner LegalEntityAssociationRelationshipType = "beneficial_owner"
	LegalEntityAssociationRelationshipTypeControlPerson   LegalEntityAssociationRelationshipType = "control_person"
)

type LegalEntityAssociationNewParams struct {
	RelationshipTypes param.Field[[]LegalEntityAssociationNewParamsRelationshipType] `json:"relationship_types,required"`
	// The associated legal entity.
	AssociatedLegalEntity param.Field[LegalEntityAssociationNewParamsAssociatedLegalEntity] `json:"associated_legal_entity"`
	// The ID of the associated legal entity.
	AssociatedLegalEntityID param.Field[string] `json:"associated_legal_entity_id"`
	// The ID of the associator legal entity. This must be a business or joint legal
	// entity.
	AssociatorLegalEntityID param.Field[string] `json:"associator_legal_entity_id"`
	// The associated entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the associated entity at the associator entity.
	Title param.Field[string] `json:"title"`
}

func (r LegalEntityAssociationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the associated entity relates to associator
// entity.
type LegalEntityAssociationNewParamsRelationshipType string

const (
	LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner LegalEntityAssociationNewParamsRelationshipType = "beneficial_owner"
	LegalEntityAssociationNewParamsRelationshipTypeControlPerson   LegalEntityAssociationNewParamsRelationshipType = "control_person"
)

// The associated legal entity.
type LegalEntityAssociationNewParamsAssociatedLegalEntity struct {
	// A list of addresses for the entity.
	Addresses param.Field[[]LegalEntityAssociationNewParamsAssociatedLegalEntityAddress] `json:"addresses"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// An individual's data of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// A list of identifications for the legal entity.
	Identifications param.Field[[]LegalEntityAssociationNewParamsAssociatedLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                                                 `json:"metadata"`
	PhoneNumbers param.Field[[]LegalEntityAssociationNewParamsAssociatedLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r LegalEntityAssociationNewParamsAssociatedLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityAssociationNewParamsAssociatedLegalEntityAddress struct {
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
	AddressTypes param.Field[[]string] `json:"address_types"`
	Line2        param.Field[string]   `json:"line2"`
}

func (r LegalEntityAssociationNewParamsAssociatedLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityAssociationNewParamsAssociatedLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r LegalEntityAssociationNewParamsAssociatedLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType string

const (
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeArCuil    LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "ar_cuil"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeArCuit    LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "ar_cuit"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeBrCnpj    LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "br_cnpj"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeBrCpf     LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "br_cpf"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeClNut     LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "cl_nut"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeCoCedulas LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "co_cedulas"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeCoNit     LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "co_nit"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeHnID      LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "hn_id"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeHnRtn     LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "hn_rtn"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypePassport  LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "passport"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeUsEin     LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "us_ein"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeUsItin    LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "us_itin"
	LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeUsSsn     LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDType = "us_ssn"
)

// The type of legal entity.
type LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityType string

const (
	LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityTypeBusiness   LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityType = "business"
	LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityTypeIndividual LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityType = "individual"
)

// A list of phone numbers in E.164 format.
type LegalEntityAssociationNewParamsAssociatedLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r LegalEntityAssociationNewParamsAssociatedLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
