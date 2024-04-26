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
		Name: moderntreasury.F("string"),
		Accounting: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccounting{
			Type: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountingTypeCustomer),
		}),
		Accounts: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccount{{
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			Name: moderntreasury.F("string"),
			AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetail{{
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}}),
			RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetail{{
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("string"),
			PartyIdentifier: moderntreasury.F("string"),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Name:                     moderntreasury.F("string"),
				Description:              moderntreasury.F("string"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:                 moderntreasury.F("string"),
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
			PlaidProcessorToken: moderntreasury.F("string"),
			ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetail{{
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}}),
		}, {
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			Name: moderntreasury.F("string"),
			AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetail{{
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}}),
			RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetail{{
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("string"),
			PartyIdentifier: moderntreasury.F("string"),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Name:                     moderntreasury.F("string"),
				Description:              moderntreasury.F("string"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:                 moderntreasury.F("string"),
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
			PlaidProcessorToken: moderntreasury.F("string"),
			ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetail{{
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}}),
		}, {
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsPartyAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			Name: moderntreasury.F("string"),
			AccountDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsAccountDetail{{
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsAccountDetailsAccountNumberTypeIban),
			}}),
			RoutingDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsRoutingDetail{{
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsRoutingDetailsPaymentTypeACH),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("string"),
			PartyIdentifier: moderntreasury.F("string"),
			LedgerAccount: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsLedgerAccount{
				Name:                     moderntreasury.F("string"),
				Description:              moderntreasury.F("string"),
				NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:                 moderntreasury.F("string"),
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
			PlaidProcessorToken: moderntreasury.F("string"),
			ContactDetails: moderntreasury.F([]moderntreasury.CounterpartyNewParamsAccountsContactDetail{{
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.CounterpartyNewParamsAccountsContactDetailsContactIdentifierTypeEmail),
			}}),
		}}),
		Email:      moderntreasury.F("dev@stainlessapi.com"),
		LedgerType: moderntreasury.F(moderntreasury.CounterpartyNewParamsLedgerTypeCustomer),
		LegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntity{
			LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityTypeBusiness),
			FirstName:            moderntreasury.F("string"),
			LastName:             moderntreasury.F("string"),
			DateOfBirth:          moderntreasury.F(time.Now()),
			DateFormed:           moderntreasury.F(time.Now()),
			BusinessName:         moderntreasury.F("string"),
			DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
			LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalStructureCorporation),
			PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityPhoneNumber{{
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
			Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddress{{
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}, {
				AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityAddressesAddressTypeOther}),
				Line1:        moderntreasury.F("string"),
				Line2:        moderntreasury.F("string"),
				Locality:     moderntreasury.F("string"),
				Region:       moderntreasury.F("string"),
				PostalCode:   moderntreasury.F("string"),
				Country:      moderntreasury.F("string"),
			}}),
			Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityIdentification{{
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}, {
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}, {
				IDNumber:       moderntreasury.F("string"),
				IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityIdentificationsIDTypeArCuil),
				IssuingCountry: moderntreasury.F("string"),
			}}),
			LegalEntityAssociations: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociation{{
				RelationshipTypes:   moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("string"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					FirstName:            moderntreasury.F("string"),
					LastName:             moderntreasury.F("string"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("string"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("string"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("string"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					FirstName:            moderntreasury.F("string"),
					LastName:             moderntreasury.F("string"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("string"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("string"),
			}, {
				RelationshipTypes:   moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeBeneficialOwner, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsRelationshipTypeControlPerson}),
				Title:               moderntreasury.F("string"),
				OwnershipPercentage: moderntreasury.F(int64(0)),
				ChildLegalEntity: moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntity{
					LegalEntityType:      moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalEntityTypeBusiness),
					FirstName:            moderntreasury.F("string"),
					LastName:             moderntreasury.F("string"),
					DateOfBirth:          moderntreasury.F(time.Now()),
					DateFormed:           moderntreasury.F(time.Now()),
					BusinessName:         moderntreasury.F("string"),
					DoingBusinessAsNames: moderntreasury.F([]string{"string", "string", "string"}),
					LegalStructure:       moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityLegalStructureCorporation),
					PhoneNumbers: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityPhoneNumber{{
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
					Addresses: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddress{{
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}, {
						AddressTypes: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressType{moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeBusiness, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeMailing, moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityAddressesAddressTypeOther}),
						Line1:        moderntreasury.F("string"),
						Line2:        moderntreasury.F("string"),
						Locality:     moderntreasury.F("string"),
						Region:       moderntreasury.F("string"),
						PostalCode:   moderntreasury.F("string"),
						Country:      moderntreasury.F("string"),
					}}),
					Identifications: moderntreasury.F([]moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentification{{
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}, {
						IDNumber:       moderntreasury.F("string"),
						IDType:         moderntreasury.F(moderntreasury.CounterpartyNewParamsLegalEntityLegalEntityAssociationsChildLegalEntityIdentificationsIDTypeArCuil),
						IssuingCountry: moderntreasury.F("string"),
					}}),
				}),
				ChildLegalEntityID: moderntreasury.F("string"),
			}}),
		}),
		LegalEntityID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		SendRemittanceAdvice: moderntreasury.F(true),
		TaxpayerIdentifier:   moderntreasury.F("string"),
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
	_, err := client.Counterparties.Get(context.TODO(), "string")
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
		"string",
		moderntreasury.CounterpartyUpdateParams{
			Email:         moderntreasury.F("dev@stainlessapi.com"),
			LegalEntityID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			Name:                 moderntreasury.F("string"),
			SendRemittanceAdvice: moderntreasury.F(true),
			TaxpayerIdentifier:   moderntreasury.F("string"),
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
		AfterCursor:         moderntreasury.F("string"),
		CreatedAtLowerBound: moderntreasury.F(time.Now()),
		CreatedAtUpperBound: moderntreasury.F(time.Now()),
		Email:               moderntreasury.F("dev@stainlessapi.com"),
		LegalEntityID:       moderntreasury.F("string"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		Name:    moderntreasury.F("string"),
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
	err := client.Counterparties.Delete(context.TODO(), "string")
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
		"string",
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
