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

func TestBulkRequestNewWithOptionalParams(t *testing.T) {
	t.Skip("Multipart documents aren't constructed properly yet")
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
	_, err := client.BulkRequests.New(context.TODO(), moderntreasury.BulkRequestNewParams{
		ActionType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsActionTypeCreate),
		ResourceType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourceTypePaymentOrder),
		Resources: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourceUnion{moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest{
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Subtype:              moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionCredit),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityHigh),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                moderntreasury.F(shared.CurrencyAed),
			EffectiveDate:           moderntreasury.F(time.Now()),
			Description:             moderntreasury.F("description"),
			StatementDescriptor:     moderntreasury.F("statement_descriptor"),
			RemittanceInformation:   moderntreasury.F("remittance_information"),
			ProcessAfter:            moderntreasury.F(time.Now()),
			Purpose:                 moderntreasury.F("purpose"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ChargeBearer:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerShared),
			ForeignExchangeIndicator:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorFixedToVariable),
			ForeignExchangeContract:            moderntreasury.F("foreign_exchange_contract"),
			NsfProtected:                       moderntreasury.F(true),
			OriginatingPartyName:               moderntreasury.F("originating_party_name"),
			UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
			UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			ExpiresAt:                          moderntreasury.F(time.Now()),
			FallbackType:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackTypeACH),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("line1"),
					Line2:      moderntreasury.F("line2"),
					Locality:   moderntreasury.F("locality"),
					Region:     moderntreasury.F("region"),
					PostalCode: moderntreasury.F("postal_code"),
					Country:    moderntreasury.F("country"),
				}),
				Name: moderntreasury.F("name"),
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}, {
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}, {
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("party_name"),
				PartyIdentifier: moderntreasury.F("party_identifier"),
				LedgerAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccount{
					Name:                     moderntreasury.F("name"),
					Description:              moderntreasury.F("description"),
					NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:                 moderntreasury.F("currency"),
					CurrencyExponent:         moderntreasury.F(int64(0)),
					LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
					LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeCounterparty),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
				ContactDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			LedgerTransaction: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransaction{
				Description: moderntreasury.F("description"),
				Status:      moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusArchived),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				EffectiveAt:   moderntreasury.F(time.Now()),
				EffectiveDate: moderntreasury.F(time.Now()),
				LedgerEntries: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerEntry{{
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}, {
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}, {
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}}),
				ExternalID:     moderntreasury.F("external_id"),
				LedgerableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeExpectedPayment),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem{{
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}}),
			TransactionMonitoringEnabled: moderntreasury.F(true),
		}, moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest{
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Subtype:              moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionCredit),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityHigh),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                moderntreasury.F(shared.CurrencyAed),
			EffectiveDate:           moderntreasury.F(time.Now()),
			Description:             moderntreasury.F("description"),
			StatementDescriptor:     moderntreasury.F("statement_descriptor"),
			RemittanceInformation:   moderntreasury.F("remittance_information"),
			ProcessAfter:            moderntreasury.F(time.Now()),
			Purpose:                 moderntreasury.F("purpose"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ChargeBearer:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerShared),
			ForeignExchangeIndicator:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorFixedToVariable),
			ForeignExchangeContract:            moderntreasury.F("foreign_exchange_contract"),
			NsfProtected:                       moderntreasury.F(true),
			OriginatingPartyName:               moderntreasury.F("originating_party_name"),
			UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
			UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			ExpiresAt:                          moderntreasury.F(time.Now()),
			FallbackType:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackTypeACH),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("line1"),
					Line2:      moderntreasury.F("line2"),
					Locality:   moderntreasury.F("locality"),
					Region:     moderntreasury.F("region"),
					PostalCode: moderntreasury.F("postal_code"),
					Country:    moderntreasury.F("country"),
				}),
				Name: moderntreasury.F("name"),
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}, {
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}, {
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("party_name"),
				PartyIdentifier: moderntreasury.F("party_identifier"),
				LedgerAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccount{
					Name:                     moderntreasury.F("name"),
					Description:              moderntreasury.F("description"),
					NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:                 moderntreasury.F("currency"),
					CurrencyExponent:         moderntreasury.F(int64(0)),
					LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
					LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeCounterparty),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
				ContactDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			LedgerTransaction: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransaction{
				Description: moderntreasury.F("description"),
				Status:      moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusArchived),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				EffectiveAt:   moderntreasury.F(time.Now()),
				EffectiveDate: moderntreasury.F(time.Now()),
				LedgerEntries: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerEntry{{
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}, {
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}, {
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}}),
				ExternalID:     moderntreasury.F("external_id"),
				LedgerableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeExpectedPayment),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem{{
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}}),
			TransactionMonitoringEnabled: moderntreasury.F(true),
		}, moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequest{
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Subtype:              moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionCredit),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityHigh),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                moderntreasury.F(shared.CurrencyAed),
			EffectiveDate:           moderntreasury.F(time.Now()),
			Description:             moderntreasury.F("description"),
			StatementDescriptor:     moderntreasury.F("statement_descriptor"),
			RemittanceInformation:   moderntreasury.F("remittance_information"),
			ProcessAfter:            moderntreasury.F(time.Now()),
			Purpose:                 moderntreasury.F("purpose"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ChargeBearer:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerShared),
			ForeignExchangeIndicator:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorFixedToVariable),
			ForeignExchangeContract:            moderntreasury.F("foreign_exchange_contract"),
			NsfProtected:                       moderntreasury.F(true),
			OriginatingPartyName:               moderntreasury.F("originating_party_name"),
			UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
			UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			ExpiresAt:                          moderntreasury.F(time.Now()),
			FallbackType:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackTypeACH),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("line1"),
					Line2:      moderntreasury.F("line2"),
					Locality:   moderntreasury.F("locality"),
					Region:     moderntreasury.F("region"),
					PostalCode: moderntreasury.F("postal_code"),
					Country:    moderntreasury.F("country"),
				}),
				Name: moderntreasury.F("name"),
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}, {
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}, {
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("party_name"),
				PartyIdentifier: moderntreasury.F("party_identifier"),
				LedgerAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccount{
					Name:                     moderntreasury.F("name"),
					Description:              moderntreasury.F("description"),
					NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:                 moderntreasury.F("currency"),
					CurrencyExponent:         moderntreasury.F(int64(0)),
					LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
					LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountLedgerAccountLedgerableTypeCounterparty),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
				ContactDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			LedgerTransaction: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransaction{
				Description: moderntreasury.F("description"),
				Status:      moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionStatusArchived),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				EffectiveAt:   moderntreasury.F(time.Now()),
				EffectiveDate: moderntreasury.F(time.Now()),
				LedgerEntries: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerEntry{{
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}, {
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}, {
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LockVersion:     moderntreasury.F(int64(0)),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}}),
				ExternalID:     moderntreasury.F("external_id"),
				LedgerableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLedgerTransactionLedgerableTypeExpectedPayment),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem{{
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("description"),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
			}}),
			TransactionMonitoringEnabled: moderntreasury.F(true),
		}}),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBulkRequestGet(t *testing.T) {
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
	_, err := client.BulkRequests.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBulkRequestListWithOptionalParams(t *testing.T) {
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
	_, err := client.BulkRequests.List(context.TODO(), moderntreasury.BulkRequestListParams{
		ActionType:  moderntreasury.F(moderntreasury.BulkRequestListParamsActionTypeCreate),
		AfterCursor: moderntreasury.F("after_cursor"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage:      moderntreasury.F(int64(0)),
		ResourceType: moderntreasury.F(moderntreasury.BulkRequestListParamsResourceTypePaymentOrder),
		Status:       moderntreasury.F(moderntreasury.BulkRequestListParamsStatusPending),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
