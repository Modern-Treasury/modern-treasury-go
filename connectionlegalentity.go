// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// ConnectionLegalEntityService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionLegalEntityService] method instead.
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
	opts = slices.Concat(r.Options, opts)
	path := "api/connection_legal_entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details on a single connection legal entity.
func (r *ConnectionLegalEntityService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConnectionLegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/connection_legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a connection legal entity.
func (r *ConnectionLegalEntityService) Update(ctx context.Context, id string, body ConnectionLegalEntityUpdateParams, opts ...option.RequestOption) (res *ConnectionLegalEntity, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/connection_legal_entities/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a list of all connection legal entities.
func (r *ConnectionLegalEntityService) List(ctx context.Context, query ConnectionLegalEntityListParams, opts ...option.RequestOption) (res *pagination.Page[ConnectionLegalEntity], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	ID string `json:"id,required" format:"uuid"`
	// The ID of the connection.
	ConnectionID string    `json:"connection_id,required"`
	CreatedAt    time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt  time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The ID of the legal entity.
	LegalEntityID string `json:"legal_entity_id,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The status of the connection legal entity.
	Status    ConnectionLegalEntityStatus `json:"status,required"`
	UpdatedAt time.Time                   `json:"updated_at,required" format:"date-time"`
	// The ID of the legal entity at the vendor.
	VendorID string                    `json:"vendor_id,required"`
	JSON     connectionLegalEntityJSON `json:"-"`
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
	VendorID      apijson.Field
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
	Addresses    param.Field[[]shared.LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[shared.LegalEntityBankSettingsParam]           `json:"bank_settings"`
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
	LegalEntityAssociations param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation] `json:"legal_entity_associations"`
	// The type of legal entity.
	LegalEntityType param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[ConnectionLegalEntityNewParamsLegalEntityLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName   param.Field[string]                                                 `json:"middle_name"`
	PhoneNumbers param.Field[[]ConnectionLegalEntityNewParamsLegalEntityPhoneNumber] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[ConnectionLegalEntityNewParamsLegalEntityRiskRating] `json:"risk_rating"`
	// An individual's suffix.
	Suffix                     param.Field[string]                                        `json:"suffix"`
	WealthAndEmploymentDetails param.Field[shared.LegalEntityWealthEmploymentDetailParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r ConnectionLegalEntityNewParamsLegalEntity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation struct {
	RelationshipTypes param.Field[[]ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[shared.ChildLegalEntityCreateParam] `json:"child_legal_entity"`
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
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeAuthorizedSigner ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "authorized_signer"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner  ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "beneficial_owner"
	ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson    ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType = "control_person"
)

func (r ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeAuthorizedSigner, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson:
		return true
	}
	return false
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

// The risk rating of the legal entity. One of low, medium, high.
type ConnectionLegalEntityNewParamsLegalEntityRiskRating string

const (
	ConnectionLegalEntityNewParamsLegalEntityRiskRatingLow    ConnectionLegalEntityNewParamsLegalEntityRiskRating = "low"
	ConnectionLegalEntityNewParamsLegalEntityRiskRatingMedium ConnectionLegalEntityNewParamsLegalEntityRiskRating = "medium"
	ConnectionLegalEntityNewParamsLegalEntityRiskRatingHigh   ConnectionLegalEntityNewParamsLegalEntityRiskRating = "high"
)

func (r ConnectionLegalEntityNewParamsLegalEntityRiskRating) IsKnown() bool {
	switch r {
	case ConnectionLegalEntityNewParamsLegalEntityRiskRatingLow, ConnectionLegalEntityNewParamsLegalEntityRiskRatingMedium, ConnectionLegalEntityNewParamsLegalEntityRiskRatingHigh:
		return true
	}
	return false
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
