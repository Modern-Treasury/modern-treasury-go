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
		ConnectionID: moderntreasury.F("string"),
		LegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntity{
			LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness),
			RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityRiskRatingLow),
			FirstName:            moderntreasury.F("string"),
			LastName:             moderntreasury.F("string"),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DateFormed:           moderntreasury.F(time.Now()),
			BusinessName:         moderntreasury.F("string"),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalStructureCorporation),
			PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityPhoneNumber{{
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
			Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddress{{
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}}),
			Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}, {
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}, {
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}}),
			LegalEntityAssociations: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation{{
				RelationshipTypes:   moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("string"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("string"),
					LastName:             moderntreasury.F("string"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("string"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("string"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("string"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("string"),
					LastName:             moderntreasury.F("string"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("string"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("string"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("string"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("string"),
					LastName:             moderntreasury.F("string"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("string"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("string"),
			}}),
		}),
		LegalEntityID: moderntreasury.F("string"),
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
	_, err := client.ConnectionLegalEntities.Get(context.TODO(), "string")
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
		"string",
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
		AfterCursor:   moderntreasury.F("string"),
		ConnectionID:  moderntreasury.F("string"),
		LegalEntityID: moderntreasury.F("string"),
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
