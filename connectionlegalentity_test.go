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
			LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityTypeBusiness),
			RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityRiskRatingLow),
			FirstName:            moderntreasury.F("first_name"),
			LastName:             moderntreasury.F("last_name"),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DateFormed:           moderntreasury.F(time.Now()),
			BusinessName:         moderntreasury.F("business_name"),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalStructureCorporation),
			PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityPhoneNumber{{
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
			Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddress{{
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}}),
			Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}, {
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}, {
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}}),
			LegalEntityAssociations: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociation{{
				RelationshipTypes:   moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("title"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("first_name"),
					LastName:             moderntreasury.F("last_name"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("business_name"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("title"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("first_name"),
					LastName:             moderntreasury.F("last_name"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("business_name"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("title"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("first_name"),
					LastName:             moderntreasury.F("last_name"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("business_name"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.ConnectionLegalEntityNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
			}}),
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
