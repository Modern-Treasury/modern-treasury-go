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
			AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddressesAddressType{moderntreasury.LegalEntityNewParamsAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsAddressesAddressTypeOther}),
			Line1:        moderntreasury.F("line1"),
			Line2:        moderntreasury.F("line2"),
			Locality:     moderntreasury.F("locality"),
			Region:       moderntreasury.F("region"),
			PostalCode:   moderntreasury.F("postal_code"),
			Country:      moderntreasury.F("country"),
		}, {
			AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddressesAddressType{moderntreasury.LegalEntityNewParamsAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsAddressesAddressTypeOther}),
			Line1:        moderntreasury.F("line1"),
			Line2:        moderntreasury.F("line2"),
			Locality:     moderntreasury.F("locality"),
			Region:       moderntreasury.F("region"),
			PostalCode:   moderntreasury.F("postal_code"),
			Country:      moderntreasury.F("country"),
		}, {
			AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddressesAddressType{moderntreasury.LegalEntityNewParamsAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsAddressesAddressTypeOther}),
			Line1:        moderntreasury.F("line1"),
			Line2:        moderntreasury.F("line2"),
			Locality:     moderntreasury.F("locality"),
			Region:       moderntreasury.F("region"),
			PostalCode:   moderntreasury.F("postal_code"),
			Country:      moderntreasury.F("country"),
		}}),
		BusinessName:         moderntreasury.F("business_name"),
		DateFormed:           moderntreasury.F(time.Now()),
		DateOfBirth:          moderntreasury.F(time.Now()),
		DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
		Email:                moderntreasury.F("email"),
		FirstName:            moderntreasury.F("first_name"),
		Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsIdentification{{
			IDNumber:       moderntreasury.F("id_number"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("issuing_country"),
		}, {
			IDNumber:       moderntreasury.F("id_number"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("issuing_country"),
		}, {
			IDNumber:       moderntreasury.F("id_number"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("issuing_country"),
		}}),
		LastName: moderntreasury.F("last_name"),
		LegalEntityAssociations: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociation{{
			RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson}),
			Title:               moderntreasury.F("title"),
			OwnershipPercentage: moderntreasury.F(int64(0)),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				RiskRating:           moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingLow),
				FirstName:            moderntreasury.F("first_name"),
				LastName:             moderntreasury.F("last_name"),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DateFormed:           moderntreasury.F(time.Now()),
				BusinessName:         moderntreasury.F("business_name"),
				DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
				LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("phone_number"),
				}, {
					PhoneNumber: moderntreasury.F("phone_number"),
				}, {
					PhoneNumber: moderntreasury.F("phone_number"),
				}}),
				Email:   moderntreasury.F("email"),
				Website: moderntreasury.F("website"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}}),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}, {
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}, {
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}}),
			}),
			ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
		}, {
			RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson}),
			Title:               moderntreasury.F("title"),
			OwnershipPercentage: moderntreasury.F(int64(0)),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				RiskRating:           moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingLow),
				FirstName:            moderntreasury.F("first_name"),
				LastName:             moderntreasury.F("last_name"),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DateFormed:           moderntreasury.F(time.Now()),
				BusinessName:         moderntreasury.F("business_name"),
				DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
				LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("phone_number"),
				}, {
					PhoneNumber: moderntreasury.F("phone_number"),
				}, {
					PhoneNumber: moderntreasury.F("phone_number"),
				}}),
				Email:   moderntreasury.F("email"),
				Website: moderntreasury.F("website"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}}),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}, {
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}, {
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}}),
			}),
			ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
		}, {
			RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson}),
			Title:               moderntreasury.F("title"),
			OwnershipPercentage: moderntreasury.F(int64(0)),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				RiskRating:           moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityRiskRatingLow),
				FirstName:            moderntreasury.F("first_name"),
				LastName:             moderntreasury.F("last_name"),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DateFormed:           moderntreasury.F(time.Now()),
				BusinessName:         moderntreasury.F("business_name"),
				DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
				LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("phone_number"),
				}, {
					PhoneNumber: moderntreasury.F("phone_number"),
				}, {
					PhoneNumber: moderntreasury.F("phone_number"),
				}}),
				Email:   moderntreasury.F("email"),
				Website: moderntreasury.F("website"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("line1"),
					Line2:        moderntreasury.F("line2"),
					Locality:     moderntreasury.F("locality"),
					Region:       moderntreasury.F("region"),
					PostalCode:   moderntreasury.F("postal_code"),
					Country:      moderntreasury.F("country"),
				}}),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}, {
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}, {
					IDNumber:       moderntreasury.F("id_number"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("issuing_country"),
				}}),
			}),
			ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
		}}),
		LegalStructure: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalStructureCorporation),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsPhoneNumber{{
			PhoneNumber: moderntreasury.F("phone_number"),
		}, {
			PhoneNumber: moderntreasury.F("phone_number"),
		}, {
			PhoneNumber: moderntreasury.F("phone_number"),
		}}),
		RiskRating: moderntreasury.F(moderntreasury.LegalEntityNewParamsRiskRatingLow),
		Website:    moderntreasury.F("website"),
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
			BusinessName:         moderntreasury.F("business_name"),
			DateFormed:           moderntreasury.F(time.Now()),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			Email:                moderntreasury.F("email"),
			FirstName:            moderntreasury.F("first_name"),
			LastName:             moderntreasury.F("last_name"),
			LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityUpdateParamsLegalStructureCorporation),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityUpdateParamsPhoneNumber{{
				PhoneNumber: moderntreasury.F("phone_number"),
			}, {
				PhoneNumber: moderntreasury.F("phone_number"),
			}, {
				PhoneNumber: moderntreasury.F("phone_number"),
			}}),
			RiskRating: moderntreasury.F(moderntreasury.LegalEntityUpdateParamsRiskRatingLow),
			Website:    moderntreasury.F("website"),
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
