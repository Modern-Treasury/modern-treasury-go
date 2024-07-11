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
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				Region:     moderntreasury.F("region"),
				PostalCode: moderntreasury.F("postal_code"),
				Country:    moderntreasury.F("country"),
			}),
			Name: moderntreasury.F("name"),
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
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("party_name"),
			PartyIdentifier: moderntreasury.F("party_identifier"),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Name:                     moderntreasury.F("name"),
				Description:              moderntreasury.F("description"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:                 moderntreasury.F("currency"),
				CurrencyExponent:         moderntreasury.F(int64(0)),
				LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeCounterparty),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
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
		}, {
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				Region:     moderntreasury.F("region"),
				PostalCode: moderntreasury.F("postal_code"),
				Country:    moderntreasury.F("country"),
			}),
			Name: moderntreasury.F("name"),
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
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("party_name"),
			PartyIdentifier: moderntreasury.F("party_identifier"),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Name:                     moderntreasury.F("name"),
				Description:              moderntreasury.F("description"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:                 moderntreasury.F("currency"),
				CurrencyExponent:         moderntreasury.F(int64(0)),
				LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeCounterparty),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
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
		}, {
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Line1:      moderntreasury.F("line1"),
				Line2:      moderntreasury.F("line2"),
				Locality:   moderntreasury.F("locality"),
				Region:     moderntreasury.F("region"),
				PostalCode: moderntreasury.F("postal_code"),
				Country:    moderntreasury.F("country"),
			}),
			Name: moderntreasury.F("name"),
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
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("party_name"),
			PartyIdentifier: moderntreasury.F("party_identifier"),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Name:                     moderntreasury.F("name"),
				Description:              moderntreasury.F("description"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:                 moderntreasury.F("currency"),
				CurrencyExponent:         moderntreasury.F(int64(0)),
				LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
				LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:           moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccountLedgerableTypeCounterparty),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
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
		}}),
		Email:      moderntreasury.F("dev@stainlessapi.com"),
		LedgerType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLedgerTypeCustomer),
		LegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntity{
			LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityTypeBusiness),
			RiskRating:           moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityRiskRatingLow),
			FirstName:            moderntreasury.F("first_name"),
			LastName:             moderntreasury.F("last_name"),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DateFormed:           moderntreasury.F(time.Now()),
			BusinessName:         moderntreasury.F("business_name"),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalStructureCorporation),
			PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityPhoneNumber{{
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
			Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddress{{
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("line1"),
				Line2:        moderntreasury.F("line2"),
				Locality:     moderntreasury.F("locality"),
				Region:       moderntreasury.F("region"),
				PostalCode:   moderntreasury.F("postal_code"),
				Country:      moderntreasury.F("country"),
			}}),
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
			LegalEntityAssociations: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociation{{
				RelationshipTypes:   moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("title"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("first_name"),
					LastName:             moderntreasury.F("last_name"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("business_name"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}}),
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
				}),
				ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("title"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("first_name"),
					LastName:             moderntreasury.F("last_name"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("business_name"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}}),
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
				}),
				ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("title"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					RiskRating:           moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityRiskRatingLow),
					FirstName:            moderntreasury.F("first_name"),
					LastName:             moderntreasury.F("last_name"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("business_name"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("line1"),
						Line2:        moderntreasury.F("line2"),
						Locality:     moderntreasury.F("locality"),
						Region:       moderntreasury.F("region"),
						PostalCode:   moderntreasury.F("postal_code"),
						Country:      moderntreasury.F("country"),
					}}),
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
				}),
				ChildLegalEntityID: moderntreasury.F("child_legal_entity_id"),
			}}),
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
