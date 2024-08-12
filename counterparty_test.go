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

func TestCounterpartyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Counterparties.New(context.TODO(), moderntreasury.CounterpartyNewParams{
		Name: moderntreasury.F("name"),
		Accounting: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccounting{
			Type: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountingTypeCustomer),
		}),
		Accounts: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccount{{
			AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetail{{
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}, {
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}, {
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}}),
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetail{{
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}}),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Currency:                 moderntreasury.F("currency"),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:                     moderntreasury.F("name"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				CurrencyExponent:         moderntreasury.F(int64(0)),
				Description:              moderntreasury.F("description"),
				LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeCounterparty),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Name: moderntreasury.F("name"),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Country:    moderntreasury.F("country"),
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				PostalCode: moderntreasury.F("postal_code"),
				Region:     moderntreasury.F("region"),
			}),
			PartyIdentifier:     moderntreasury.F("party_identifier"),
			PartyName:           moderntreasury.F("party_name"),
			PartyType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
			RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetail{{
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}}),
		}, {
			AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetail{{
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}, {
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}, {
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}}),
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetail{{
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}}),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Currency:                 moderntreasury.F("currency"),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:                     moderntreasury.F("name"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				CurrencyExponent:         moderntreasury.F(int64(0)),
				Description:              moderntreasury.F("description"),
				LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeCounterparty),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Name: moderntreasury.F("name"),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Country:    moderntreasury.F("country"),
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				PostalCode: moderntreasury.F("postal_code"),
				Region:     moderntreasury.F("region"),
			}),
			PartyIdentifier:     moderntreasury.F("party_identifier"),
			PartyName:           moderntreasury.F("party_name"),
			PartyType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
			RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetail{{
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}}),
		}, {
			AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetail{{
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}, {
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}, {
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeAuNumber),
			}}),
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetail{{
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("contact_identifier"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}}),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Currency:                 moderntreasury.F("currency"),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Name:                     moderntreasury.F("name"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				CurrencyExponent:         moderntreasury.F(int64(0)),
				Description:              moderntreasury.F("description"),
				LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeCounterparty),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Name: moderntreasury.F("name"),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Country:    moderntreasury.F("country"),
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				PostalCode: moderntreasury.F("postal_code"),
				Region:     moderntreasury.F("region"),
			}),
			PartyIdentifier:     moderntreasury.F("party_identifier"),
			PartyName:           moderntreasury.F("party_name"),
			PartyType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
			RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetail{{
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}}),
		}}),
		Email:      moderntreasury.F("dev@stainlessapi.com"),
		LedgerType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLedgerTypeCustomer),
		LegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntity{
			LegalEntityType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityTypeBusiness),
			Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddress{{
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line2:        moderntreasury.F("line2"),
			}, {
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line2:        moderntreasury.F("line2"),
			}, {
				Country:      moderntreasury.F("country"),
				Line1:        moderntreasury.F("line1"),
				Locality:     moderntreasury.F("locality"),
				PostalCode:   moderntreasury.F("postal_code"),
				Region:       moderntreasury.F("region"),
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line2:        moderntreasury.F("line2"),
			}}),
			BusinessName:         moderntreasury.F("business_name"),
			DateFormed:           moderntreasury.F(time.Now()),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			Email:                moderntreasury.F("email"),
			FirstName:            moderntreasury.F("first_name"),
			Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}, {
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}, {
				IDNumber:       moderntreasury.F("id_number"),
				IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("issuing_country"),
			}}),
			LastName: moderntreasury.F("last_name"),
			LegalEntityAssociations: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociation{{
				RelationshipTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}, {
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}, {
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}}),
					BusinessName:         moderntreasury.F("business_name"),
					DateFormed:           moderntreasury.F(time.Now()),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					Email:                moderntreasury.F("email"),
					FirstName:            moderntreasury.F("first_name"),
					Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
					LastName:        moderntreasury.F("last_name"),
					LegalEntityType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					LegalStructure:  moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
						PhoneNumber: moderntreasury.F("phone_number"),
					}, {
						PhoneNumber: moderntreasury.F("phone_number"),
					}, {
						PhoneNumber: moderntreasury.F("phone_number"),
					}}),
					RiskRating: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					Website:    moderntreasury.F("website"),
				}),
				ChildLegalEntityID:  moderntreasury.F("child_legal_entity_id"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				Title:               moderntreasury.F("title"),
			}, {
				RelationshipTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}, {
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}, {
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}}),
					BusinessName:         moderntreasury.F("business_name"),
					DateFormed:           moderntreasury.F(time.Now()),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					Email:                moderntreasury.F("email"),
					FirstName:            moderntreasury.F("first_name"),
					Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
					LastName:        moderntreasury.F("last_name"),
					LegalEntityType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					LegalStructure:  moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
						PhoneNumber: moderntreasury.F("phone_number"),
					}, {
						PhoneNumber: moderntreasury.F("phone_number"),
					}, {
						PhoneNumber: moderntreasury.F("phone_number"),
					}}),
					RiskRating: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					Website:    moderntreasury.F("website"),
				}),
				ChildLegalEntityID:  moderntreasury.F("child_legal_entity_id"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				Title:               moderntreasury.F("title"),
			}, {
				RelationshipTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}, {
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}, {
						Country:      moderntreasury.F("country"),
						Line1:        moderntreasury.F("line1"),
						Locality:     moderntreasury.F("locality"),
						PostalCode:   moderntreasury.F("postal_code"),
						Region:       moderntreasury.F("region"),
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line2:        moderntreasury.F("line2"),
					}}),
					BusinessName:         moderntreasury.F("business_name"),
					DateFormed:           moderntreasury.F(time.Now()),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					Email:                moderntreasury.F("email"),
					FirstName:            moderntreasury.F("first_name"),
					Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}, {
						IDNumber:       moderntreasury.F("id_number"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("issuing_country"),
					}}),
					LastName:        moderntreasury.F("last_name"),
					LegalEntityType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					LegalStructure:  moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
						PhoneNumber: moderntreasury.F("phone_number"),
					}, {
						PhoneNumber: moderntreasury.F("phone_number"),
					}, {
						PhoneNumber: moderntreasury.F("phone_number"),
					}}),
					RiskRating: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					Website:    moderntreasury.F("website"),
				}),
				ChildLegalEntityID:  moderntreasury.F("child_legal_entity_id"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				Title:               moderntreasury.F("title"),
			}}),
			LegalStructure: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalStructureCorporation),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityPhoneNumber{{
				PhoneNumber: moderntreasury.F("phone_number"),
			}, {
				PhoneNumber: moderntreasury.F("phone_number"),
			}, {
				PhoneNumber: moderntreasury.F("phone_number"),
			}}),
			RiskRating: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityRiskRatingLow),
			Website:    moderntreasury.F("website"),
		}),
		LegalEntityID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		SendRemittanceAdvice: moderntreasury.F(true),
		TaxpayerIdentifier:   moderntreasury.F("taxpayer_identifier"),
		VerificationStatus:   moderntreasury.F(moderntreasury.CounterpartyNewParamsVerificationStatusDenied),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyGet(t *testing.T) {
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
	_, err := client.Counterparties.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Counterparties.Update(
		context.TODO(),
		"id",
		moderntreasury.CounterpartyUpdateParams{
			Email:         moderntreasury.F("dev@stainlessapi.com"),
			LegalEntityID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name:                 moderntreasury.F("name"),
			SendRemittanceAdvice: moderntreasury.F(true),
			TaxpayerIdentifier:   moderntreasury.F("taxpayer_identifier"),
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

func TestCounterpartyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Counterparties.List(context.TODO(), moderntreasury.CounterpartyListParams{
		AfterCursor:         moderntreasury.F("after_cursor"),
		CreatedAtLowerBound: moderntreasury.F(time.Now()),
		CreatedAtUpperBound: moderntreasury.F(time.Now()),
		Email:               moderntreasury.F("dev@stainlessapi.com"),
		LegalEntityID:       moderntreasury.F("legal_entity_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		Name:    moderntreasury.F("name"),
		PerPage: moderntreasury.F(int64(0)),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyDelete(t *testing.T) {
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
	err := client.Counterparties.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterpartyCollectAccountWithOptionalParams(t *testing.T) {
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
	_, err := client.Counterparties.CollectAccount(
		context.TODO(),
		"id",
		moderntreasury.CounterpartyCollectAccountParams{
			Direction:      moderntreasury.F(shared.TransactionDirectionCredit),
			CustomRedirect: moderntreasury.F("https://example.com"),
			Fields:         moderntreasury.F([]moderntreasury.CounterpartyCollectAccountParamsField{moderntreasury.CounterpartyCollectAccountParamsFieldName, moderntreasury.CounterpartyCollectAccountParamsFieldNameOnAccount, moderntreasury.CounterpartyCollectAccountParamsFieldTaxpayerIdentifier}),
			SendEmail:      moderntreasury.F(true),
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
