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
			Line1:        moderntreasury.F("string"),
			Line2:        moderntreasury.F("string"),
			Locality:     moderntreasury.F("string"),
			Region:       moderntreasury.F("string"),
			PostalCode:   moderntreasury.F("string"),
			Country:      moderntreasury.F("string"),
		}, {
			AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddressesAddressType{moderntreasury.LegalEntityNewParamsAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsAddressesAddressTypeOther}),
			Line1:        moderntreasury.F("string"),
			Line2:        moderntreasury.F("string"),
			Locality:     moderntreasury.F("string"),
			Region:       moderntreasury.F("string"),
			PostalCode:   moderntreasury.F("string"),
			Country:      moderntreasury.F("string"),
		}, {
			AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsAddressesAddressType{moderntreasury.LegalEntityNewParamsAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsAddressesAddressTypeOther}),
			Line1:        moderntreasury.F("string"),
			Line2:        moderntreasury.F("string"),
			Locality:     moderntreasury.F("string"),
			Region:       moderntreasury.F("string"),
			PostalCode:   moderntreasury.F("string"),
			Country:      moderntreasury.F("string"),
		}}),
		BusinessName:         moderntreasury.F("string"),
		DateFormed:           moderntreasury.F(time.Now()),
		DateOfBirth:          moderntreasury.F(time.Now()),
		DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
		Email:                moderntreasury.F("string"),
		FirstName:            moderntreasury.F("string"),
		Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsIdentification{{
			IDNumber:       moderntreasury.F("string"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("string"),
		}, {
			IDNumber:       moderntreasury.F("string"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("string"),
		}, {
			IDNumber:       moderntreasury.F("string"),
			IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsIdentificationsIDTypeArCuil),
			IssuingCountry: moderntreasury.F("string"),
		}}),
		LastName: moderntreasury.F("string"),
		LegalEntityAssociations: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociation{{
			RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson}),
			Title:               moderntreasury.F("string"),
			OwnershipPercentage: moderntreasury.F(int64(0)),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				FirstName:            moderntreasury.F("string"),
				LastName:             moderntreasury.F("string"),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DateFormed:           moderntreasury.F(time.Now()),
				BusinessName:         moderntreasury.F("string"),
				DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
				LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("string"),
				}, {
					PhoneNumber: moderntreasury.F("string"),
				}, {
					PhoneNumber: moderntreasury.F("string"),
				}}),
				Email:   moderntreasury.F("string"),
				Website: moderntreasury.F("string"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}}),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}, {
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}, {
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}}),
			}),
			ChildLegalEntityID: moderntreasury.F("string"),
		}, {
			RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson}),
			Title:               moderntreasury.F("string"),
			OwnershipPercentage: moderntreasury.F(int64(0)),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				FirstName:            moderntreasury.F("string"),
				LastName:             moderntreasury.F("string"),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DateFormed:           moderntreasury.F(time.Now()),
				BusinessName:         moderntreasury.F("string"),
				DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
				LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("string"),
				}, {
					PhoneNumber: moderntreasury.F("string"),
				}, {
					PhoneNumber: moderntreasury.F("string"),
				}}),
				Email:   moderntreasury.F("string"),
				Website: moderntreasury.F("string"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}}),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}, {
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}, {
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}}),
			}),
			ChildLegalEntityID: moderntreasury.F("string"),
		}, {
			RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsRelationshipTypeControlPerson}),
			Title:               moderntreasury.F("string"),
			OwnershipPercentage: moderntreasury.F(int64(0)),
			ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntity{
				LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
				FirstName:            moderntreasury.F("string"),
				LastName:             moderntreasury.F("string"),
				DateOfBirth:          moderntreasury.F(time.Now()),
				DateFormed:           moderntreasury.F(time.Now()),
				BusinessName:         moderntreasury.F("string"),
				DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
				LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
				PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityPhoneNumber{{
					PhoneNumber: moderntreasury.F("string"),
				}, {
					PhoneNumber: moderntreasury.F("string"),
				}, {
					PhoneNumber: moderntreasury.F("string"),
				}}),
				Email:   moderntreasury.F("string"),
				Website: moderntreasury.F("string"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Addresses: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddress{{
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}, {
					AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
					Line1:        moderntreasury.F("string"),
					Line2:        moderntreasury.F("string"),
					Locality:     moderntreasury.F("string"),
					Region:       moderntreasury.F("string"),
					PostalCode:   moderntreasury.F("string"),
					Country:      moderntreasury.F("string"),
				}}),
				Identifications: moderntreasury.F([]moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentification{{
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}, {
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}, {
					IDNumber:       moderntreasury.F("string"),
					IDType:         moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
					IssuingCountry: moderntreasury.F("string"),
				}}),
			}),
			ChildLegalEntityID: moderntreasury.F("string"),
		}}),
		LegalStructure: moderntreasury.F(moderntreasury.LegalEntityNewParamsLegalStructureCorporation),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityNewParamsPhoneNumber{{
			PhoneNumber: moderntreasury.F("string"),
		}, {
			PhoneNumber: moderntreasury.F("string"),
		}, {
			PhoneNumber: moderntreasury.F("string"),
		}}),
		Website: moderntreasury.F("string"),
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
	_, err := client.LegalEntities.Get(context.TODO(), "string")
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
		"string",
		moderntreasury.LegalEntityUpdateParams{
			BusinessName:         moderntreasury.F("string"),
			DateFormed:           moderntreasury.F(time.Now()),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			Email:                moderntreasury.F("string"),
			FirstName:            moderntreasury.F("string"),
			LastName:             moderntreasury.F("string"),
			LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityUpdateParamsLegalStructureCorporation),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityUpdateParamsPhoneNumber{{
				PhoneNumber: moderntreasury.F("string"),
			}, {
				PhoneNumber: moderntreasury.F("string"),
			}, {
				PhoneNumber: moderntreasury.F("string"),
			}}),
			Website: moderntreasury.F("string"),
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
		AfterCursor:     moderntreasury.F("string"),
		LegalEntityType: moderntreasury.F(moderntreasury.LegalEntityListParamsLegalEntityTypeBusiness),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage:     moderntreasury.F(int64(0)),
		ShowDeleted: moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
