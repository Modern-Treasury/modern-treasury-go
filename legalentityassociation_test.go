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
			LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityLegalEntityTypeBusiness),
			RiskRating:           moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityRiskRatingLow),
			FirstName:            moderntreasury.F("first_name"),
			LastName:             moderntreasury.F("last_name"),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DateFormed:           moderntreasury.F(time.Now()),
			BusinessName:         moderntreasury.F("business_name"),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityLegalStructureCorporation),
			PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityPhoneNumber{{
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
			Addresses: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddress{{
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsChildLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}}),
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
