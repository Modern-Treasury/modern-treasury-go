// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestConnectionLegalEntityNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.ConnectionLegalEntities.New(context.TODO(), moderntreasury.ConnectionLegalEntityNewParams{
		ConnectionID: moderntreasury.F("connection_id"),
		LegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntity{
			Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddress{{
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness}),
				Line2:        moderntreasury.F("line2"),
			}}),
			BankSettings: moderntreasury.F(moderntreasury.BankSettingsParam{
				ID:                          moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				BackupWithholdingPercentage: moderntreasury.F(int64(0)),
				CreatedAt:                   moderntreasury.F(time.Now()),
				DiscardedAt:                 moderntreasury.F(time.Now()),
				EnableBackupWithholding:     moderntreasury.F(true),
				LiveMode:                    moderntreasury.F(true),
				Object:                      moderntreasury.F("object"),
				PrivacyOptOut:               moderntreasury.F(true),
				RegulationO:                 moderntreasury.F(true),
				UpdatedAt:                   moderntreasury.F(time.Now()),
			}),
			BusinessName:         moderntreasury.F("business_name"),
			CitizenshipCountry:   moderntreasury.F("citizenship_country"),
			DateFormed:           moderntreasury.F(time.Now()),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DoingBusinessAsNames: moderntreasury.F([]string{"string"}),
			Email:                moderntreasury.F("email"),
			FirstName:            moderntreasury.F("first_name"),
			Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}}),
			LastName: moderntreasury.F("last_name"),
			LegalEntityAssociations: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation{{
				RelationshipTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner}),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness}),
						Line2:        moderntreasury.F("line2"),
					}}),
					BankSettings: moderntreasury.F(moderntreasury.BankSettingsParam{
						ID:                          moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
						BackupWithholdingPercentage: moderntreasury.F(int64(0)),
						CreatedAt:                   moderntreasury.F(time.Now()),
						DiscardedAt:                 moderntreasury.F(time.Now()),
						EnableBackupWithholding:     moderntreasury.F(true),
						LiveMode:                    moderntreasury.F(true),
						Object:                      moderntreasury.F("object"),
						PrivacyOptOut:               moderntreasury.F(true),
						RegulationO:                 moderntreasury.F(true),
						UpdatedAt:                   moderntreasury.F(time.Now()),
					}),
					BusinessName:         moderntreasury.F("business_name"),
					CitizenshipCountry:   moderntreasury.F("citizenship_country"),
					DateFormed:           moderntreasury.F(time.Now()),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DoingBusinessAsNames: moderntreasury.F([]string{"string"}),
					Email:                moderntreasury.F("email"),
					FirstName:            moderntreasury.F("first_name"),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
					LastName:        moderntreasury.F("last_name"),
					LegalEntityType: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					LegalStructure:  moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
					MiddleName: moderntreasury.F("middle_name"),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
						PhoneNumber: moderntreasury.F("phone_number"),
					}}),
					PoliticallyExposedPerson: moderntreasury.F(true),
					PreferredName:            moderntreasury.F("preferred_name"),
					Prefix:                   moderntreasury.F("prefix"),
					RiskRating:               moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					Suffix:                   moderntreasury.F("suffix"),
					WealthAndEmploymentDetails: moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsParam{
						ID:               moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
						AnnualIncome:     moderntreasury.F(int64(0)),
						CreatedAt:        moderntreasury.F(time.Now()),
						DiscardedAt:      moderntreasury.F(time.Now()),
						EmployerCountry:  moderntreasury.F("employer_country"),
						EmployerName:     moderntreasury.F("employer_name"),
						EmployerState:    moderntreasury.F("employer_state"),
						EmploymentStatus: moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsEmploymentStatusEmployed),
						IncomeCountry:    moderntreasury.F("income_country"),
						IncomeSource:     moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsIncomeSourceFamilySupport),
						IncomeState:      moderntreasury.F("income_state"),
						Industry:         moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsIndustryAccounting),
						LiveMode:         moderntreasury.F(true),
						Object:           moderntreasury.F("object"),
						Occupation:       moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsOccupationConsulting),
						SourceOfFunds:    moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsSourceOfFundsAlimony),
						UpdatedAt:        moderntreasury.F(time.Now()),
						WealthSource:     moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsWealthSourceBusinessSale),
					}),
					Website: moderntreasury.F("website"),
				}),
				ChildLegalEntityID:  moderntreasury.F("child_legal_entity_id"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				Title:               moderntreasury.F("title"),
			}}),
			LegalEntityType: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness),
			LegalStructure:  moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalStructureCorporation),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			MiddleName: moderntreasury.F("middle_name"),
			PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityPhoneNumber{{
				PhoneNumber: moderntreasury.F("phone_number"),
			}}),
			PoliticallyExposedPerson: moderntreasury.F(true),
			PreferredName:            moderntreasury.F("preferred_name"),
			Prefix:                   moderntreasury.F("prefix"),
			RiskRating:               moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityRiskRatingLow),
			Suffix:                   moderntreasury.F("suffix"),
			WealthAndEmploymentDetails: moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsParam{
				ID:               moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AnnualIncome:     moderntreasury.F(int64(0)),
				CreatedAt:        moderntreasury.F(time.Now()),
				DiscardedAt:      moderntreasury.F(time.Now()),
				EmployerCountry:  moderntreasury.F("employer_country"),
				EmployerName:     moderntreasury.F("employer_name"),
				EmployerState:    moderntreasury.F("employer_state"),
				EmploymentStatus: moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsEmploymentStatusEmployed),
				IncomeCountry:    moderntreasury.F("income_country"),
				IncomeSource:     moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsIncomeSourceFamilySupport),
				IncomeState:      moderntreasury.F("income_state"),
				Industry:         moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsIndustryAccounting),
				LiveMode:         moderntreasury.F(true),
				Object:           moderntreasury.F("object"),
				Occupation:       moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsOccupationConsulting),
				SourceOfFunds:    moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsSourceOfFundsAlimony),
				UpdatedAt:        moderntreasury.F(time.Now()),
				WealthSource:     moderntreasury.F(moderntreasury.WealthAndEmploymentDetailsWealthSourceBusinessSale),
			}),
			Website: moderntreasury.F("website"),
		}),
		LegalEntityID: moderntreasury.F("legal_entity_id"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionLegalEntityGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.ConnectionLegalEntities.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionLegalEntityUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.ConnectionLegalEntities.Update(
		context.TODO(),
		"id",
		moderntreasury.ConnectionLegalEntityUpdateParams{
			Status: moderntreasury.F(moderntreasury.ConnectionLegalEntityUpdateParamsStatusProcessing),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectionLegalEntityListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.ConnectionLegalEntities.List(context.TODO(), moderntreasury.ConnectionLegalEntityListParams{
		AfterCursor:   moderntreasury.F("after_cursor"),
		ConnectionID:  moderntreasury.F("connection_id"),
		LegalEntityID: moderntreasury.F("legal_entity_id"),
		PerPage:       moderntreasury.F(int64(0)),
		Status:        moderntreasury.F(moderntreasury.ConnectionLegalEntityListParamsStatusCompleted),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
