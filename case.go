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
)

// CaseService contains methods and other services that help with interacting with
// the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCaseService] method instead.
type CaseService struct {
	Options []option.RequestOption
}

// NewCaseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCaseService(opts ...option.RequestOption) (r *CaseService) {
	r = &CaseService{}
	r.Options = opts
	return
}

// Get details on a single case.
func (r *CaseService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Case, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/cases/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a list of cases.
func (r *CaseService) List(ctx context.Context, query CaseListParams, opts ...option.RequestOption) (res *pagination.Page[Case], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/cases"
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

// Get a list of cases.
func (r *CaseService) ListAutoPaging(ctx context.Context, query CaseListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Case] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type Case struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The pending actions requested to resolve the case.
	RequestedActions []CaseRequestedAction `json:"requested_actions" api:"required"`
	// The requested actions that have been resolved.
	ResolvedActions []CaseResolvedAction `json:"resolved_actions" api:"required"`
	// The status of the case.
	Status CaseStatus `json:"status" api:"required"`
	// The ID of the object the case is about.
	SubjectID string `json:"subject_id" api:"required" format:"uuid"`
	// The type of the object the case is about.
	SubjectType string    `json:"subject_type" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	JSON        caseJSON  `json:"-"`
}

// caseJSON contains the JSON metadata for the struct [Case]
type caseJSON struct {
	ID               apijson.Field
	CreatedAt        apijson.Field
	LiveMode         apijson.Field
	Object           apijson.Field
	RequestedActions apijson.Field
	ResolvedActions  apijson.Field
	Status           apijson.Field
	SubjectID        apijson.Field
	SubjectType      apijson.Field
	UpdatedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Case) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseJSON) RawJSON() string {
	return r.raw
}

type CaseRequestedAction struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The category of the requested action.
	Category  CaseRequestedActionsCategory `json:"category" api:"required"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	// The field that needs to be corrected or provided, if any.
	Field CaseRequestedActionsField `json:"field" api:"required,nullable"`
	// Instructions on how to resolve the requested action.
	Instructions string `json:"instructions" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The reasons the action was requested.
	Reasons   []string                `json:"reasons" api:"required"`
	UpdatedAt time.Time               `json:"updated_at" api:"required" format:"date-time"`
	JSON      caseRequestedActionJSON `json:"-"`
}

// caseRequestedActionJSON contains the JSON metadata for the struct
// [CaseRequestedAction]
type caseRequestedActionJSON struct {
	ID           apijson.Field
	Category     apijson.Field
	CreatedAt    apijson.Field
	Field        apijson.Field
	Instructions apijson.Field
	LiveMode     apijson.Field
	Object       apijson.Field
	Reasons      apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CaseRequestedAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseRequestedActionJSON) RawJSON() string {
	return r.raw
}

// The category of the requested action.
type CaseRequestedActionsCategory string

const (
	CaseRequestedActionsCategoryOnboardingArticlesOfIncorporationFailure      CaseRequestedActionsCategory = "onboarding_articles_of_incorporation_failure"
	CaseRequestedActionsCategoryOnboardingBusinessRegistryVerificationFailure CaseRequestedActionsCategory = "onboarding_business_registry_verification_failure"
	CaseRequestedActionsCategoryOnboardingDatabaseFailure                     CaseRequestedActionsCategory = "onboarding_database_failure"
	CaseRequestedActionsCategoryOnboardingProofOfAddressFailure               CaseRequestedActionsCategory = "onboarding_proof_of_address_failure"
	CaseRequestedActionsCategoryOnboardingSsnCheckFailure                     CaseRequestedActionsCategory = "onboarding_ssn_check_failure"
	CaseRequestedActionsCategoryOnboardingTinCheckFailure                     CaseRequestedActionsCategory = "onboarding_tin_check_failure"
)

func (r CaseRequestedActionsCategory) IsKnown() bool {
	switch r {
	case CaseRequestedActionsCategoryOnboardingArticlesOfIncorporationFailure, CaseRequestedActionsCategoryOnboardingBusinessRegistryVerificationFailure, CaseRequestedActionsCategoryOnboardingDatabaseFailure, CaseRequestedActionsCategoryOnboardingProofOfAddressFailure, CaseRequestedActionsCategoryOnboardingSsnCheckFailure, CaseRequestedActionsCategoryOnboardingTinCheckFailure:
		return true
	}
	return false
}

// The field that needs to be corrected or provided, if any.
type CaseRequestedActionsField string

const (
	CaseRequestedActionsFieldArticlesOfIncorporation CaseRequestedActionsField = "articles_of_incorporation"
	CaseRequestedActionsFieldEinLetter               CaseRequestedActionsField = "ein_letter"
	CaseRequestedActionsFieldLegalEntityDetails      CaseRequestedActionsField = "legal_entity_details"
	CaseRequestedActionsFieldProofOfAddress          CaseRequestedActionsField = "proof_of_address"
)

func (r CaseRequestedActionsField) IsKnown() bool {
	switch r {
	case CaseRequestedActionsFieldArticlesOfIncorporation, CaseRequestedActionsFieldEinLetter, CaseRequestedActionsFieldLegalEntityDetails, CaseRequestedActionsFieldProofOfAddress:
		return true
	}
	return false
}

type CaseResolvedAction struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The category of the requested action.
	Category  CaseResolvedActionsCategory `json:"category" api:"required"`
	CreatedAt time.Time                   `json:"created_at" api:"required" format:"date-time"`
	// The field that needs to be corrected or provided, if any.
	Field CaseResolvedActionsField `json:"field" api:"required,nullable"`
	// Instructions on how to resolve the requested action.
	Instructions string `json:"instructions" api:"required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode" api:"required"`
	Object   string `json:"object" api:"required"`
	// The reasons the action was requested.
	Reasons   []string               `json:"reasons" api:"required"`
	UpdatedAt time.Time              `json:"updated_at" api:"required" format:"date-time"`
	JSON      caseResolvedActionJSON `json:"-"`
}

// caseResolvedActionJSON contains the JSON metadata for the struct
// [CaseResolvedAction]
type caseResolvedActionJSON struct {
	ID           apijson.Field
	Category     apijson.Field
	CreatedAt    apijson.Field
	Field        apijson.Field
	Instructions apijson.Field
	LiveMode     apijson.Field
	Object       apijson.Field
	Reasons      apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CaseResolvedAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caseResolvedActionJSON) RawJSON() string {
	return r.raw
}

// The category of the requested action.
type CaseResolvedActionsCategory string

const (
	CaseResolvedActionsCategoryOnboardingArticlesOfIncorporationFailure      CaseResolvedActionsCategory = "onboarding_articles_of_incorporation_failure"
	CaseResolvedActionsCategoryOnboardingBusinessRegistryVerificationFailure CaseResolvedActionsCategory = "onboarding_business_registry_verification_failure"
	CaseResolvedActionsCategoryOnboardingDatabaseFailure                     CaseResolvedActionsCategory = "onboarding_database_failure"
	CaseResolvedActionsCategoryOnboardingProofOfAddressFailure               CaseResolvedActionsCategory = "onboarding_proof_of_address_failure"
	CaseResolvedActionsCategoryOnboardingSsnCheckFailure                     CaseResolvedActionsCategory = "onboarding_ssn_check_failure"
	CaseResolvedActionsCategoryOnboardingTinCheckFailure                     CaseResolvedActionsCategory = "onboarding_tin_check_failure"
)

func (r CaseResolvedActionsCategory) IsKnown() bool {
	switch r {
	case CaseResolvedActionsCategoryOnboardingArticlesOfIncorporationFailure, CaseResolvedActionsCategoryOnboardingBusinessRegistryVerificationFailure, CaseResolvedActionsCategoryOnboardingDatabaseFailure, CaseResolvedActionsCategoryOnboardingProofOfAddressFailure, CaseResolvedActionsCategoryOnboardingSsnCheckFailure, CaseResolvedActionsCategoryOnboardingTinCheckFailure:
		return true
	}
	return false
}

// The field that needs to be corrected or provided, if any.
type CaseResolvedActionsField string

const (
	CaseResolvedActionsFieldArticlesOfIncorporation CaseResolvedActionsField = "articles_of_incorporation"
	CaseResolvedActionsFieldEinLetter               CaseResolvedActionsField = "ein_letter"
	CaseResolvedActionsFieldLegalEntityDetails      CaseResolvedActionsField = "legal_entity_details"
	CaseResolvedActionsFieldProofOfAddress          CaseResolvedActionsField = "proof_of_address"
)

func (r CaseResolvedActionsField) IsKnown() bool {
	switch r {
	case CaseResolvedActionsFieldArticlesOfIncorporation, CaseResolvedActionsFieldEinLetter, CaseResolvedActionsFieldLegalEntityDetails, CaseResolvedActionsFieldProofOfAddress:
		return true
	}
	return false
}

// The status of the case.
type CaseStatus string

const (
	CaseStatusOpen     CaseStatus = "open"
	CaseStatusResolved CaseStatus = "resolved"
)

func (r CaseStatus) IsKnown() bool {
	switch r {
	case CaseStatusOpen, CaseStatusResolved:
		return true
	}
	return false
}

type CaseListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	PerPage     param.Field[int64]  `query:"per_page"`
	// The status of the case.
	Status param.Field[CaseListParamsStatus] `query:"status"`
	// The ID of the object the case is about.
	SubjectID param.Field[string] `query:"subject_id"`
	// The type of the object the case is about.
	SubjectType param.Field[CaseListParamsSubjectType] `query:"subject_type"`
}

// URLQuery serializes [CaseListParams]'s query parameters as `url.Values`.
func (r CaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The status of the case.
type CaseListParamsStatus string

const (
	CaseListParamsStatusOpen     CaseListParamsStatus = "open"
	CaseListParamsStatusResolved CaseListParamsStatus = "resolved"
)

func (r CaseListParamsStatus) IsKnown() bool {
	switch r {
	case CaseListParamsStatusOpen, CaseListParamsStatusResolved:
		return true
	}
	return false
}

// The type of the object the case is about.
type CaseListParamsSubjectType string

const (
	CaseListParamsSubjectTypeLegalEntity CaseListParamsSubjectType = "legal_entity"
)

func (r CaseListParamsSubjectType) IsKnown() bool {
	switch r {
	case CaseListParamsSubjectTypeLegalEntity:
		return true
	}
	return false
}
