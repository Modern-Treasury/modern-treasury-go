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

func TestLegalEntityAssociationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LegalEntityAssociations.New(context.TODO(), moderntreasury.LegalEntityAssociationNewParams{
		ParentLegalEntityID: moderntreasury.F("parent_legal_entity_id"),
		RelationshipTypes:   moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsRelationshipType{moderntreasury.LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityAssociationNewParamsRelationshipTypeControlPerson}),
		ChildLegalEntity: moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntity{
			Addresses: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddress{{
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther}),
				Line2:        moderntreasury.F("line2"),
			}, {
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther}),
				Line2:        moderntreasury.F("line2"),
			}, {
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther}),
				Line2:        moderntreasury.F("line2"),
			}}),
			BusinessName:         moderntreasury.F("business_name"),
			DateFormed:           moderntreasury.F(time.Now()),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			Email:                moderntreasury.F("email"),
			FirstName:            moderntreasury.F("first_name"),
			Identifications: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}, {
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}, {
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}}),
			LastName:        moderntreasury.F("last_name"),
			LegalEntityType: moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityLegalEntityTypeBusiness),
			LegalStructure:  moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityLegalStructureCorporation),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityPhoneNumber{{
				PhoneNumber: moderntreasury.F("phone_number"),
			}, {
				PhoneNumber: moderntreasury.F("phone_number"),
			}, {
				PhoneNumber: moderntreasury.F("phone_number"),
			}}),
			RiskRating: moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityRiskRatingLow),
			Website:    moderntreasury.F("website"),
		}),
		ChildLegalEntityID:  moderntreasury.F("child_legal_entity_id"),
		OwnershipPercentage: moderntreasury.F(int64(0)),
		Title:               moderntreasury.F("title"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
