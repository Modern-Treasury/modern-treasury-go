// File generated from our OpenAPI spec by Stainless.

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
		RelationshipTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsRelationshipType{moderntreasury.LegalEntityAssociationNewParamsRelationshipTypeBeneficialOwner, moderntreasury.LegalEntityAssociationNewParamsRelationshipTypeControlPerson}),
		AssociatedLegalEntity: moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntity{
			LegalEntityType:      moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityLegalEntityTypeBusiness),
			FirstName:            moderntreasury.F("string"),
			LastName:             moderntreasury.F("string"),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DateFormed:           moderntreasury.F(time.Now()),
			BusinessName:         moderntreasury.F("string"),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			LegalStructure:       moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityLegalStructureCorporation),
			PhoneNumbers: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityPhoneNumber{{
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
			Addresses: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddress{{
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressType{moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeBusiness, moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeMailing, moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}}),
			Identifications: moderntreasury.F([]moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}, {
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}, {
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.LegalEntityAssociationNewParamsAssociatedLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}}),
		}),
		AssociatedLegalEntityID: moderntreasury.F("string"),
		AssociatorLegalEntityID: moderntreasury.F("string"),
		OwnershipPercentage:     moderntreasury.F(int64(0)),
		Title:                   moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
