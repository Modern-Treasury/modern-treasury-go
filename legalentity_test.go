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
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

func TestLegalEntityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LegalEntities.New(context.TODO(), moderntreasury.LegalEntityNewParams{
		LegalEntityType: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityTypeBusiness),
		Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddress{{
			Country:      moderntreasury.F("country"),
			Line1:        moderntreasury.F("line1"),
			Locality:     moderntreasury.F("locality"),
			PostalCode:   moderntreasury.F("postal_code"),
			Region:       moderntreasury.F("region"),
			AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddressesAddressType{moderntreasury.LegalEntityNewParamsAddressesAddressTypeBusiness}),
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
		BusinessName:       moderntreasury.F("business_name"),
		CitizenshipCountry: moderntreasury.F("citizenship_country"),
		ComplianceDetails: moderntreasury.F(shared.LegalEntityComplianceDetailParam{
			ID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CreatedAt:      moderntreasury.F(time.Now()),
			DiscardedAt:    moderntreasury.F(time.Now()),
			Issuer:         moderntreasury.F("issuer"),
			LiveMode:       moderntreasury.F(true),
			Object:         moderntreasury.F("object"),
			TokenExpiresAt: moderntreasury.F(time.Now()),
			TokenIssuedAt:  moderntreasury.F(time.Now()),
			TokenURL:       moderntreasury.F("token_url"),
			UpdatedAt:      moderntreasury.F(time.Now()),
			Validated:      moderntreasury.F(true),
			ValidatedAt:    moderntreasury.F(time.Now()),
		}),
		DateFormed:           moderntreasury.F(time.Now()),
		DateOfBirth:          moderntreasury.F(time.Now()),
		DoingBusinessAsNames: moderntreasury.F([]string{"string"}),
		Email:                moderntreasury.F("email"),
		FirstName:            moderntreasury.F("first_name"),
		Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsIdentification{{
			IDNumber:       moderntreasury.F("id_number"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			ExpirationDate: moderntreasury.F(time.Now()),
			IssuingCountry: moderntreasury.F("issuing_country"),
			IssuingRegion:  moderntreasury.F("issuing_region"),
		}}),
		IndustryClassifications: moderntreasury.F([]shared.LegalEntityIndustryClassificationParam{{
			ID:                  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ClassificationCodes: moderntreasury.F([]string{"string"}),
			ClassificationType:  moderntreasury.F(shared.LegalEntityIndustryClassificationClassificationTypeAnzsic),
			CreatedAt:           moderntreasury.F(time.Now()),
			DiscardedAt:         moderntreasury.F(time.Now()),
			LiveMode:            moderntreasury.F(true),
			Object:              moderntreasury.F("object"),
			UpdatedAt:           moderntreasury.F(time.Now()),
		}}),
		LastName: moderntreasury.F("last_name"),
		LegalEntityAssociations: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociation{{
			RelationshipTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner}),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					Country:      moderntreasury.F("country"),
					Line1:        moderntreasury.F("line1"),
					Locality:     moderntreasury.F("locality"),
					PostalCode:   moderntreasury.F("postal_code"),
					Region:       moderntreasury.F("region"),
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness}),
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
				BusinessName:       moderntreasury.F("business_name"),
				CitizenshipCountry: moderntreasury.F("citizenship_country"),
				ComplianceDetails: moderntreasury.F(shared.LegalEntityComplianceDetailParam{
					ID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					CreatedAt:      moderntreasury.F(time.Now()),
					DiscardedAt:    moderntreasury.F(time.Now()),
					Issuer:         moderntreasury.F("issuer"),
					LiveMode:       moderntreasury.F(true),
					Object:         moderntreasury.F("object"),
					TokenExpiresAt: moderntreasury.F(time.Now()),
					TokenIssuedAt:  moderntreasury.F(time.Now()),
					TokenURL:       moderntreasury.F("token_url"),
					UpdatedAt:      moderntreasury.F(time.Now()),
					Validated:      moderntreasury.F(true),
					ValidatedAt:    moderntreasury.F(time.Now()),
				}),
				DateFormed:           moderntreasury.F(time.Now()),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DoingBusinessAsNames: moderntreasury.F([]string{"string"}),
				Email:                moderntreasury.F("email"),
				FirstName:            moderntreasury.F("first_name"),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					ExpirationDate: moderntreasury.F(time.Now()),
					IssuingCountry: moderntreasury.F("issuing_country"),
					IssuingRegion:  moderntreasury.F("issuing_region"),
				}}),
				IndustryClassifications: moderntreasury.F([]shared.LegalEntityIndustryClassificationParam{{
					ID:                  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					ClassificationCodes: moderntreasury.F([]string{"string"}),
					ClassificationType:  moderntreasury.F(shared.LegalEntityIndustryClassificationClassificationTypeAnzsic),
					CreatedAt:           moderntreasury.F(time.Now()),
					DiscardedAt:         moderntreasury.F(time.Now()),
					LiveMode:            moderntreasury.F(true),
					Object:              moderntreasury.F("object"),
					UpdatedAt:           moderntreasury.F(time.Now()),
				}}),
				LastName:        moderntreasury.F("last_name"),
				LegalEntityType: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				LegalStructure:  moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				MiddleName: moderntreasury.F("middle_name"),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("phone_number"),
				}}),
				PoliticallyExposedPerson: moderntreasury.F(true),
				PreferredName:            moderntreasury.F("preferred_name"),
				Prefix:                   moderntreasury.F("prefix"),
				RiskRating:               moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingLow),
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
		LegalStructure: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalStructureCorporation),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		MiddleName: moderntreasury.F("middle_name"),
		PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsPhoneNumber{{
			PhoneNumber: moderntreasury.F("phone_number"),
		}}),
		PoliticallyExposedPerson: moderntreasury.F(true),
		PreferredName:            moderntreasury.F("preferred_name"),
		Prefix:                   moderntreasury.F("prefix"),
		RiskRating:               moderntreasury.F(moderntreasury.LegalEntityNewParamsRiskRatingLow),
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
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalEntityGet(t *testing.T) {
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
	_, err := client.LegalEntities.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLegalEntityUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LegalEntities.Update(
		context.TODO(),
		"id",
		moderntreasury.LegalEntityUpdateParams{
			Addresses: moderntreasury.F([]moderntreasury.LegalEntityUpdateParamsAddress{{
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityUpdateParamsAddressesAddressType{moderntreasury.LegalEntityUpdateParamsAddressesAddressTypeBusiness}),
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
			BusinessName:       moderntreasury.F("business_name"),
			CitizenshipCountry: moderntreasury.F("citizenship_country"),
			ComplianceDetails: moderntreasury.F(shared.LegalEntityComplianceDetailParam{
				ID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CreatedAt:      moderntreasury.F(time.Now()),
				DiscardedAt:    moderntreasury.F(time.Now()),
				Issuer:         moderntreasury.F("issuer"),
				LiveMode:       moderntreasury.F(true),
				Object:         moderntreasury.F("object"),
				TokenExpiresAt: moderntreasury.F(time.Now()),
				TokenIssuedAt:  moderntreasury.F(time.Now()),
				TokenURL:       moderntreasury.F("token_url"),
				UpdatedAt:      moderntreasury.F(time.Now()),
				Validated:      moderntreasury.F(true),
				ValidatedAt:    moderntreasury.F(time.Now()),
			}),
			DateFormed:           moderntreasury.F(time.Now()),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DoingBusinessAsNames: moderntreasury.F([]string{"string"}),
			Email:                moderntreasury.F("email"),
			FirstName:            moderntreasury.F("first_name"),
			Identifications: moderntreasury.F([]moderntreasury.LegalEntityUpdateParamsIdentification{{
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityUpdateParamsIdentificationsIDTypeArCuil),
				ExpirationDate: moderntreasury.F(time.Now()),
				IssuingCountry: moderntreasury.F("issuing_country"),
				IssuingRegion:  moderntreasury.F("issuing_region"),
			}}),
			IndustryClassifications: moderntreasury.F([]shared.LegalEntityIndustryClassificationParam{{
				ID:                  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassificationCodes: moderntreasury.F([]string{"string"}),
				ClassificationType:  moderntreasury.F(shared.LegalEntityIndustryClassificationClassificationTypeAnzsic),
				CreatedAt:           moderntreasury.F(time.Now()),
				DiscardedAt:         moderntreasury.F(time.Now()),
				LiveMode:            moderntreasury.F(true),
				Object:              moderntreasury.F("object"),
				UpdatedAt:           moderntreasury.F(time.Now()),
			}}),
			LastName:       moderntreasury.F("last_name"),
			LegalStructure: moderntreasury.F(moderntreasury.LegalEntityUpdateParamsLegalStructureCorporation),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			MiddleName: moderntreasury.F("middle_name"),
			PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityUpdateParamsPhoneNumber{{
				PhoneNumber: moderntreasury.F("phone_number"),
			}}),
			PoliticallyExposedPerson: moderntreasury.F(true),
			PreferredName:            moderntreasury.F("preferred_name"),
			Prefix:                   moderntreasury.F("prefix"),
			RiskRating:               moderntreasury.F(moderntreasury.LegalEntityUpdateParamsRiskRatingLow),
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

func TestLegalEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.LegalEntities.List(context.TODO(), moderntreasury.LegalEntityListParams{
		AfterCursor:     moderntreasury.F("after_cursor"),
		LegalEntityType: moderntreasury.F(moderntreasury.LegalEntityListParamsLegalEntityTypeBusiness),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage:     moderntreasury.F(int64(0)),
		ShowDeleted: moderntreasury.F("show_deleted"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
