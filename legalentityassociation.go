// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
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

// create legal_entity_association
func (r *LegalEntityAssociationService) New(ctx context.Context, body LegalEntityAssociationNewParams, opts ...option.RequestOption) (res *LegalEntityAssociation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/legal_entity_associations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LegalEntityAssociation struct {
	ID string `json:"id,required" format:"uuid"`
	// The child legal entity.
	ChildLegalEntity interface{} `json:"child_legal_entity,required"`
	CreatedAt        time.Time   `json:"created_at,required" format:"date-time"`
	DiscardedAt      time.Time   `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage int64 `json:"ownership_percentage,required,nullable"`
	// The ID of the parent legal entity. This must be a business or joint legal
	// entity.
	ParentLegalEntityID string                                   `json:"parent_legal_entity_id,required"`
	RelationshipTypes   []LegalEntityAssociationRelationshipType `json:"relationship_types,required"`
	// The job title of the child entity at the parent entity.
	Title     string                     `json:"title,required,nullable"`
	UpdatedAt time.Time                  `json:"updated_at,required" format:"date-time"`
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

type LegalEntityAssociationNewParams struct {
	// The ID of the parent legal entity. This must be a business or joint legal
	// entity.
	ParentLegalEntityID param.Field[string]                                            `json:"parent_legal_entity_id,required"`
	RelationshipTypes   param.Field[[]LegalEntityAssociationNewParamsRelationshipType] `json:"relationship_types,required"`
	// The child legal entity.
	ChildLegalEntity param.Field[shared.ChildLegalEntityCreateParam] `json:"child_legal_entity"`
	// The ID of the child legal entity.
	ChildLegalEntityID param.Field[string] `json:"child_legal_entity_id"`
	// The child entity's ownership percentage iff they are a beneficial owner.
	OwnershipPercentage param.Field[int64] `json:"ownership_percentage"`
	// The job title of the child entity at the parent entity.
	Title param.Field[string] `json:"title"`
}

func (r LegalEntityAssociationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of relationship types for how the child entity relates to parent entity.
type LegalEntityAssociationNewParamsRelationshipType string

const (
	LegalEntityAssociationNewParamsRelationshipTypeAuthorizedSigner LegalEntityAssociationNewParamsRelationshipType = "authorized_signer"
	LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner  LegalEntityAssociationNewParamsRelationshipType = "beneficial_owner"
	LegalEntityAssociationNewParamsRelationshipTypeControlPerson    LegalEntityAssociationNewParamsRelationshipType = "control_person"
)

func (r LegalEntityAssociationNewParamsRelationshipType) IsKnown() bool {
	switch r {
	case LegalEntityAssociationNewParamsRelationshipTypeAuthorizedSigner, LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner, LegalEntityAssociationNewParamsRelationshipTypeControlPerson:
		return true
	}
	return false
}
