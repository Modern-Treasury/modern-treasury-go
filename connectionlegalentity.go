// File generated from our OpenAPI spec by Stainless.

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
func (r *ConnectionLegalEntityService) List(ctx context.Context, query ConnectionLegalEntityListParams, opts ...option.RequestOption) (res *shared.Page[ConnectionLegalEntity], err error) {
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
func (r *ConnectionLegalEntityService) ListAutoPaging(ctx context.Context, query ConnectionLegalEntityListParams, opts ...option.RequestOption) *shared.PageAutoPager[ConnectionLegalEntity] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
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
	// The legal entity associations and its associated legal entities.
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
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeClNut     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "cl_nut"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeCoCedulas ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "co_cedulas"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeCoNit     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "co_nit"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeHnID      ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "hn_id"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeHnRtn     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "hn_rtn"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypePassport  ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "passport"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsEin     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "us_ein"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsItin    ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "us_itin"
	ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeUsSsn     ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDType = "us_ssn"
)

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The associated legal entity.
	AssociatedLegalEntity param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntity] `json:"associated_legal_entity"`
	// The ID of the associated legal entity.
	AssociatedLegalEntityID param.Field[string] `json:"associated_legal_entity_id"`
	// The associated entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the associated entity at the associator entity.
	Title param.Field[string] `json:"title"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the associated entity relates to associator
// entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "beneficial_owner"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson   ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "control_person"
)

// The associated legal entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntity struct {
	// A list of addresses for the entity.
	Addresses param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddress] `json:"addresses"`
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
	Identifications param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentification] `json:"identifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata     param.Field[map[string]string]                                                                                  `json:"metadata"`
	PhoneNumbers param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityPhoneNumber] `json:"phone_numbers"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddress struct {
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
	AddressTypes param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType] `json:"address_types"`
	Line2        param.Field[string]                                                                                                      `json:"line2"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeBusiness    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "business"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeMailing     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "mailing"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeOther       ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "other"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypePoBox       ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "po_box"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressTypeResidential ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityAddressesAddressType = "residential"
)

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentification struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType] `json:"id_type,required"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentification) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeArCuil    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "ar_cuil"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeArCuit    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "ar_cuit"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeBrCnpj    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "br_cnpj"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeBrCpf     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "br_cpf"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeClNut     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "cl_nut"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeCoCedulas ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "co_cedulas"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeCoNit     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "co_nit"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeHnID      ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "hn_id"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeHnRtn     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "hn_rtn"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypePassport  ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "passport"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsEin     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "us_ein"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsItin    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "us_itin"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDTypeUsSsn     ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityIdentificationsIDType = "us_ssn"
)

// The type of legal entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalEntityType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalEntityTypeBusiness   ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalEntityType = "business"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalEntityTypeIndividual ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalEntityType = "individual"
)

// The business's legal structure.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructureCorporation        ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "corporation"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructureLlc                ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "llc"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructureNonProfit          ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "non_profit"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructurePartnership        ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "partnership"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructureSoleProprietorship ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "sole_proprietorship"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructureTrust              ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityLegalStructure = "trust"
)

// A list of phone numbers in E.164 format.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityPhoneNumber struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsAssociatedLegalEntityPhoneNumber) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of legal entity.
type ConnectionLegalEntityNewParamsLegalEntityLegalEntityType string

const (
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness   ConnectionLegalEntityNewParamsLegalEntityLegalEntityType = "business"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeIndividual ConnectionLegalEntityNewParamsLegalEntityLegalEntityType = "individual"
)

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
