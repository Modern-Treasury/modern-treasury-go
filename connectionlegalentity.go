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
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// ConnectionLegalEntityService contains methods and other services that help with
// interacting with the Modern Treasury API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewConnectionLegalEntityService] method instead.
type ConnectionLegalEntityService struct {
	Options []option.RequestOption
}

// NewConnectionLegalEntityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectionLegalEntityService(opts ...option.RequestOption) (r *ConnectionLegalEntityService) {
	r = &ConnectionLegalEntityService{}
	r.Options = opts
	return
}

// Create a connection legal entity.
func (r *ConnectionLegalEntityService) New(ctx context.Context, body ConnectionLegalEntityNewParams, opts ...option.RequestOption) (res *ConnectionLegalEntity, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/connection_legal_entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single connection legal entity.
func (r *ConnectionLegalEntityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectionLegalEntity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/connection_legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a connection legal entity.
func (r *ConnectionLegalEntityService) Update(ctx context.Context, id string, body ConnectionLegalEntityUpdateParams, opts ...option.RequestOption) (res *ConnectionLegalEntity, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/connection_legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all connection legal entities.
func (r *ConnectionLegalEntityService) List(ctx context.Context, query ConnectionLegalEntityListParams, opts ...option.RequestOption) (res *pagination.Page[ConnectionLegalEntity], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/connection_legal_entities"
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

// Get a list of all connection legal entities.
func (r *ConnectionLegalEntityService) ListAutoPaging(ctx context.Context, query ConnectionLegalEntityListParams, opts ...option.RequestOption) *pagination.PageAutoPager[ConnectionLegalEntity] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type ConnectionLegalEntity struct {
	ID string `json:"id" format:"uuid"`
	// The ID of the connection.
	ConnectionID string    `json:"connection_id"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	DiscardedAt  time.Time `json:"discarded_at,nullable" format:"date-time"`
	// The ID of the legal entity.
	LegalEntityID string `json:"legal_entity_id"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode"`
	Object   string `json:"object"`
	// The status of the connection legal entity.
	Status    ConnectionLegalEntityStatus `json:"status"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      connectionLegalEntityJSON   `json:"-"`
}

// connectionLegalEntityJSON contains the JSON metadata for the struct
// [ConnectionLegalEntity]
type connectionLegalEntityJSON struct {
	ID            apijson.Field
	ConnectionID  apijson.Field
	CreatedAt     apijson.Field
	DiscardedAt   apijson.Field
	LegalEntityID apijson.Field
	LiveMode      apijson.Field
	Object        apijson.Field
	Status        apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ConnectionLegalEntity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionLegalEntityJSON) RawJSON() string {
	return r.raw
}

// The status of the connection legal entity.
type ConnectionLegalEntityStatus string

const (
	ConnectionLegalEntityStatusCompleted  ConnectionLegalEntityStatus = "completed"
	ConnectionLegalEntityStatusDenied     ConnectionLegalEntityStatus = "denied"
	ConnectionLegalEntityStatusFailed     ConnectionLegalEntityStatus = "failed"
	ConnectionLegalEntityStatusProcessing ConnectionLegalEntityStatus = "processing"
)

func (r ConnectionLegalEntityStatus) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityStatusCompleted, ConnectionLegalEntityStatusDenied, ConnectionLegalEntityStatusFailed, ConnectionLegalEntityStatusProcessing:
		return true
	}
	return false
}

type ConnectionLegalEntityNewParams struct {
	// The ID of the connection.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The legal entity.
	LegalEntity param.Field[ConnectionLegalEntityNewParamsLegalEntity] `json:"legal_entity"`
	// The ID of the legal entity.
	LegalEntityID param.Field[string] `json:"legal_entity_id"`
}

func (r ConnectionLegalEntityNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The legal entity.
type ConnectionLegalEntityNewParamsLegalEntity struct {
	// A list of addresses for the entity.
	Addresses param.Field[[]ConnectionLegalEntityNewParamsLegalEntityAddress] `json:"addresses"`
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
	Identifications param.Field[[]ConnectionLegalEntityNewParamsLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The legal entity associations and its child legal entities.
	LegalEntityAssociations param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation] `json:"legal_entity_associations"`
	// The type of legal entity.
	LegalEntityType param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                                      `json:"metadata"`
	PhoneNumbers param.Field[[]ConnectionLegalEntityNewParamsLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r ConnectionLegalEntityNewParamsLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityAddress struct {
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
	AddressTypes param.Field[[]ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                          `json:"line2"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType string

const (
	ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness    ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType = "business"
	ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing     ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType = "mailing"
	ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther       ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType = "other"
	ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypePoBox       ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType = "po_box"
	ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeResidential ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType = "residential"
)

func (r ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther, ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypePoBox, ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type ConnectionLegalEntityNewParamsLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType string

const (
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil    ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "ar_cuil"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuit    ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "ar_cuit"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeBrCnpj    ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "br_cnpj"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeBrCpf     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "br_cpf"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeClRut     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "cl_rut"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeCoCedulas ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "co_cedulas"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeCoNit     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "co_nit"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeHnID      ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "hn_id"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeHnRtn     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "hn_rtn"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeInLei     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "in_lei"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypePassport  ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "passport"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeSaTin     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "sa_tin"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeSaVat     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "sa_vat"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsEin     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "us_ein"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsItin    ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "us_itin"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsSsn     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "us_ssn"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeVnTin     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "vn_tin"
)

func (r ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuit, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeBrCnpj, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeBrCpf, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeClRut, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeCoCedulas, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeCoNit, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeHnID, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeHnRtn, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeInLei, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypePassport, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeSaTin, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeSaVat, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsEin, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsItin, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsSsn, ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity] `json:"child_legal_entity"`
	// The ID of the child legal entity.
	ChildLegalEntityID param.Field[string] `json:"child_legal_entity_id"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the child entity at the parent entity.
	Title param.Field[string] `json:"title"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the child entity relates to parent entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "beneficial_owner"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson   ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "control_person"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson:
		return true
	}
	return false
}

// The child legal entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity struct {
	// A list of addresses for the entity.
	Addresses param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress] `json:"addresses"`
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
	Identifications param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                                                                             `json:"metadata"`
	PhoneNumbers param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress struct {
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
	AddressTypes param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                                                                 `json:"line2"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType = "business"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType = "mailing"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther       ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType = "other"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypePoBox       ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType = "po_box"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeResidential ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType = "residential"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypePoBox, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeResidential:
		return true
	}
	return false
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "ar_cuil"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuit    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "ar_cuit"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCnpj    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "br_cnpj"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCpf     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "br_cpf"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeClRut     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "cl_rut"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoCedulas ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "co_cedulas"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoNit     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "co_nit"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnID      ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "hn_id"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnRtn     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "hn_rtn"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeInLei     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "in_lei"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypePassport  ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "passport"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaTin     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "sa_tin"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaVat     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "sa_vat"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsEin     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "us_ein"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsItin    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "us_itin"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsSsn     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "us_ssn"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeVnTin     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType = "vn_tin"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuit, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCnpj, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeBrCpf, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeClRut, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoCedulas, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeCoNit, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnID, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeHnRtn, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeInLei, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypePassport, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaTin, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeSaVat, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsEin, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsItin, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeUsSsn, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeVnTin:
		return true
	}
	return false
}

// The type of legal entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness   ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityType = "business"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeIndividual ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityType = "individual"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation        ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure = "corporation"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureLlc                ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure = "llc"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureNonProfit          ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure = "non_profit"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructurePartnership        ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure = "partnership"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureSoleProprietorship ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure = "sole_proprietorship"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureTrust              ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure = "trust"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureLlc, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureNonProfit, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructurePartnership, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureSoleProprietorship, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of legal entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness   ConnectionLegalEntityNewParamsLegalEntityLegalEntityType = "business"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeIndividual ConnectionLegalEntityNewParamsLegalEntityLegalEntityType = "individual"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness, ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type ConnectionLegalEntityNewParamsLegalEntityLegalStructure string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalStructureCorporation        ConnectionLegalEntityNewParamsLegalEntityLegalStructure = "corporation"
	ConnectionLegalEntityNewParamsLegalEntityLegalStructureLlc                ConnectionLegalEntityNewParamsLegalEntityLegalStructure = "llc"
	ConnectionLegalEntityNewParamsLegalEntityLegalStructureNonProfit          ConnectionLegalEntityNewParamsLegalEntityLegalStructure = "non_profit"
	ConnectionLegalEntityNewParamsLegalEntityLegalStructurePartnership        ConnectionLegalEntityNewParamsLegalEntityLegalStructure = "partnership"
	ConnectionLegalEntityNewParamsLegalEntityLegalStructureSoleProprietorship ConnectionLegalEntityNewParamsLegalEntityLegalStructure = "sole_proprietorship"
	ConnectionLegalEntityNewParamsLegalEntityLegalStructureTrust              ConnectionLegalEntityNewParamsLegalEntityLegalStructure = "trust"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalStructure) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalStructureCorporation, ConnectionLegalEntityNewParamsLegalEntityLegalStructureLlc, ConnectionLegalEntityNewParamsLegalEntityLegalStructureNonProfit, ConnectionLegalEntityNewParamsLegalEntityLegalStructurePartnership, ConnectionLegalEntityNewParamsLegalEntityLegalStructureSoleProprietorship, ConnectionLegalEntityNewParamsLegalEntityLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type ConnectionLegalEntityNewParamsLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityUpdateParams struct {
	// The status of the connection legal entity.
	Status param.Field[ConnectionLegalEntityUpdateParamsStatus] `json:"status"`
}

func (r ConnectionLegalEntityUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the connection legal entity.
type ConnectionLegalEntityUpdateParamsStatus string

const (
	ConnectionLegalEntityUpdateParamsStatusProcessing ConnectionLegalEntityUpdateParamsStatus = "processing"
)

func (r ConnectionLegalEntityUpdateParamsStatus) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityUpdateParamsStatusProcessing:
		return true
	}
	return false
}

type ConnectionLegalEntityListParams struct {
	AfterCursor   param.Field[string]                                `query:"after_cursor"`
	ConnectionID  param.Field[string]                                `query:"connection_id"`
	LegalEntityID param.Field[string]                                `query:"legal_entity_id"`
	PerPage       param.Field[int64]                                 `query:"per_page"`
	Status        param.Field[ConnectionLegalEntityListParamsStatus] `query:"status"`
}

// URLQuery serializes [ConnectionLegalEntityListParams]'s query parameters as
// `url.Values`.
func (r ConnectionLegalEntityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConnectionLegalEntityListParamsStatus string

const (
	ConnectionLegalEntityListParamsStatusCompleted  ConnectionLegalEntityListParamsStatus = "completed"
	ConnectionLegalEntityListParamsStatusDenied     ConnectionLegalEntityListParamsStatus = "denied"
	ConnectionLegalEntityListParamsStatusFailed     ConnectionLegalEntityListParamsStatus = "failed"
	ConnectionLegalEntityListParamsStatusProcessing ConnectionLegalEntityListParamsStatus = "processing"
)

func (r ConnectionLegalEntityListParamsStatus) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityListParamsStatusCompleted, ConnectionLegalEntityListParamsStatusDenied, ConnectionLegalEntityListParamsStatusFailed, ConnectionLegalEntityListParamsStatusProcessing:
		return true
	}
	return false
}
