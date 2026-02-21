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
			Addresses: moderntreasury.F([]shared.LegalEntityAddressCreateRequestParam{{
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]shared.LegalEntityAddressCreateRequestAddressType{shared.LegalEntityAddressCreateRequestAddressTypeBusiness}),
				Line2:        moderntreasury.F("line2"),
			}}),
			BankSettings: moderntreasury.F(shared.LegalEntityBankSettingsParam{
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
			BusinessDescription:    moderntreasury.F("business_description"),
			BusinessName:           moderntreasury.F("business_name"),
			CitizenshipCountry:     moderntreasury.F("citizenship_country"),
			ConnectionID:           moderntreasury.F("connection_id"),
			CountryOfIncorporation: moderntreasury.F("country_of_incorporation"),
			DateFormed:             moderntreasury.F(time.Now()),
			DateOfBirth:            moderntreasury.F(time.Now()),
			DoingBusinessAsNames:   moderntreasury.F([]string{"string"}),
			Email:                  moderntreasury.F("email"),
			ExpectedActivityVolume: moderntreasury.F(int64(0)),
			ExternalID:             moderntreasury.F("external_id"),
			FirstName:              moderntreasury.F("first_name"),
			Identifications: moderntreasury.F([]shared.IdentificationCreateRequestParam{{
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(shared.IdentificationCreateRequestIDTypeArCuil),
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
			IntendedUse: moderntreasury.F("intended_use"),
			LastName:    moderntreasury.F("last_name"),
			LegalEntityAssociations: moderntreasury.F([]shared.LegalEntityAssociationInlineCreateParam{{
				RelationshipTypes: moderntreasury.F([]shared.LegalEntityAssociationInlineCreateRelationshipType{shared.LegalEntityAssociationInlineCreateRelationshipTypeAuthorizedSigner}),
				ChildLegalEntity: moderntreasury.F(shared.ChildLegalEntityCreateParam{
					Addresses: moderntreasury.F([]shared.LegalEntityAddressCreateRequestParam{{
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]shared.LegalEntityAddressCreateRequestAddressType{shared.LegalEntityAddressCreateRequestAddressTypeBusiness}),
						Line2:        moderntreasury.F("line2"),
					}}),
					BankSettings: moderntreasury.F(shared.LegalEntityBankSettingsParam{
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
					BusinessDescription:    moderntreasury.F("business_description"),
					BusinessName:           moderntreasury.F("business_name"),
					CitizenshipCountry:     moderntreasury.F("citizenship_country"),
					ConnectionID:           moderntreasury.F("connection_id"),
					CountryOfIncorporation: moderntreasury.F("country_of_incorporation"),
					DateFormed:             moderntreasury.F(time.Now()),
					DateOfBirth:            moderntreasury.F(time.Now()),
					DoingBusinessAsNames:   moderntreasury.F([]string{"string"}),
					Email:                  moderntreasury.F("email"),
					ExpectedActivityVolume: moderntreasury.F(int64(0)),
					ExternalID:             moderntreasury.F("external_id"),
					FirstName:              moderntreasury.F("first_name"),
					Identifications: moderntreasury.F([]shared.IdentificationCreateRequestParam{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(shared.IdentificationCreateRequestIDTypeArCuil),
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
					IntendedUse:             moderntreasury.F("intended_use"),
					LastName:                moderntreasury.F("last_name"),
					LegalEntityAssociations: moderntreasury.F([]shared.LegalEntityAssociationInlineCreateParam{}),
					LegalEntityType:         moderntreasury.F(shared.ChildLegalEntityCreateLegalEntityTypeBusiness),
					LegalStructure:          moderntreasury.F(shared.ChildLegalEntityCreateLegalStructureCorporation),
					ListedExchange:          moderntreasury.F("listed_exchange"),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
					MiddleName:             moderntreasury.F("middle_name"),
					OperatingJurisdictions: moderntreasury.F([]string{"string"}),
					PhoneNumbers: moderntreasury.F([]shared.ChildLegalEntityCreatePhoneNumberParam{{
						PhoneNumber: moderntreasury.F("phone_number"),
					}}),
					PoliticallyExposedPerson: moderntreasury.F(true),
					PreferredName:            moderntreasury.F("preferred_name"),
					Prefix:                   moderntreasury.F("prefix"),
					PrimarySocialMediaSites:  moderntreasury.F([]string{"string"}),
					Regulators: moderntreasury.F([]shared.ChildLegalEntityCreateRegulatorParam{{
						Jurisdiction:       moderntreasury.F("jurisdiction"),
						Name:               moderntreasury.F("name"),
						RegistrationNumber: moderntreasury.F("registration_number"),
					}}),
					RiskRating: moderntreasury.F(shared.ChildLegalEntityCreateRiskRatingLow),
					Status:     moderntreasury.F(shared.ChildLegalEntityCreateStatusActive),
					Suffix:     moderntreasury.F("suffix"),
					ThirdPartyVerification: moderntreasury.F(shared.ChildLegalEntityCreateThirdPartyVerificationParam{
						Vendor:               moderntreasury.F(shared.ChildLegalEntityCreateThirdPartyVerificationVendorPersona),
						VendorVerificationID: moderntreasury.F("vendor_verification_id"),
					}),
					TickerSymbol: moderntreasury.F("ticker_symbol"),
					WealthAndEmploymentDetails: moderntreasury.F(shared.LegalEntityWealthEmploymentDetailParam{
						ID:               moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
						AnnualIncome:     moderntreasury.F(int64(0)),
						CreatedAt:        moderntreasury.F(time.Now()),
						DiscardedAt:      moderntreasury.F(time.Now()),
						EmployerCountry:  moderntreasury.F("employer_country"),
						EmployerName:     moderntreasury.F("employer_name"),
						EmployerState:    moderntreasury.F("employer_state"),
						EmploymentStatus: moderntreasury.F(shared.LegalEntityWealthEmploymentDetailEmploymentStatusEmployed),
						IncomeCountry:    moderntreasury.F("income_country"),
						IncomeSource:     moderntreasury.F(shared.LegalEntityWealthEmploymentDetailIncomeSourceFamilySupport),
						IncomeState:      moderntreasury.F("income_state"),
						Industry:         moderntreasury.F(shared.LegalEntityWealthEmploymentDetailIndustryAccounting),
						LiveMode:         moderntreasury.F(true),
						Object:           moderntreasury.F("object"),
						Occupation:       moderntreasury.F(shared.LegalEntityWealthEmploymentDetailOccupationConsulting),
						SourceOfFunds:    moderntreasury.F(shared.LegalEntityWealthEmploymentDetailSourceOfFundsAlimony),
						UpdatedAt:        moderntreasury.F(time.Now()),
						WealthSource:     moderntreasury.F(shared.LegalEntityWealthEmploymentDetailWealthSourceBusinessSale),
					}),
					Website: moderntreasury.F("website"),
				}),
				ChildLegalEntityID:  moderntreasury.F("child_legal_entity_id"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				Title:               moderntreasury.F("title"),
			}}),
			LegalEntityType: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness),
			LegalStructure:  moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalStructureCorporation),
			ListedExchange:  moderntreasury.F("listed_exchange"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			MiddleName:             moderntreasury.F("middle_name"),
			OperatingJurisdictions: moderntreasury.F([]string{"string"}),
			PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityPhoneNumber{{
				PhoneNumber: moderntreasury.F("phone_number"),
			}}),
			PoliticallyExposedPerson: moderntreasury.F(true),
			PreferredName:            moderntreasury.F("preferred_name"),
			Prefix:                   moderntreasury.F("prefix"),
			PrimarySocialMediaSites:  moderntreasury.F([]string{"string"}),
			Regulators: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityRegulator{{
				Jurisdiction:       moderntreasury.F("jurisdiction"),
				Name:               moderntreasury.F("name"),
				RegistrationNumber: moderntreasury.F("registration_number"),
			}}),
			RiskRating: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityRiskRatingLow),
			Status:     moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityStatusActive),
			Suffix:     moderntreasury.F("suffix"),
			ThirdPartyVerification: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityThirdPartyVerification{
				Vendor:               moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityThirdPartyVerificationVendorPersona),
				VendorVerificationID: moderntreasury.F("vendor_verification_id"),
			}),
			TickerSymbol: moderntreasury.F("ticker_symbol"),
			WealthAndEmploymentDetails: moderntreasury.F(shared.LegalEntityWealthEmploymentDetailParam{
				ID:               moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AnnualIncome:     moderntreasury.F(int64(0)),
				CreatedAt:        moderntreasury.F(time.Now()),
				DiscardedAt:      moderntreasury.F(time.Now()),
				EmployerCountry:  moderntreasury.F("employer_country"),
				EmployerName:     moderntreasury.F("employer_name"),
				EmployerState:    moderntreasury.F("employer_state"),
				EmploymentStatus: moderntreasury.F(shared.LegalEntityWealthEmploymentDetailEmploymentStatusEmployed),
				IncomeCountry:    moderntreasury.F("income_country"),
				IncomeSource:     moderntreasury.F(shared.LegalEntityWealthEmploymentDetailIncomeSourceFamilySupport),
				IncomeState:      moderntreasury.F("income_state"),
				Industry:         moderntreasury.F(shared.LegalEntityWealthEmploymentDetailIndustryAccounting),
				LiveMode:         moderntreasury.F(true),
				Object:           moderntreasury.F("object"),
				Occupation:       moderntreasury.F(shared.LegalEntityWealthEmploymentDetailOccupationConsulting),
				SourceOfFunds:    moderntreasury.F(shared.LegalEntityWealthEmploymentDetailSourceOfFundsAlimony),
				UpdatedAt:        moderntreasury.F(time.Now()),
				WealthSource:     moderntreasury.F(shared.LegalEntityWealthEmploymentDetailWealthSourceBusinessSale),
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
		Status:        moderntreasury.F(moderntreasury.ConnectionLegalEntityListParamsStatusClosed),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
