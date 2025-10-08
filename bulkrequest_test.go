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
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestDirectionCredit),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ChargeBearer:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestChargeBearerShared),
			Currency:                 moderntreasury.F(shared.CurrencyAed),
			Description:              moderntreasury.F("description"),
			EffectiveDate:            moderntreasury.F(time.Now()),
			ExpiresAt:                moderntreasury.F(time.Now()),
			FallbackType:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestFallbackTypeACH),
			ForeignExchangeContract:  moderntreasury.F("foreign_exchange_contract"),
			ForeignExchangeIndicator: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestForeignExchangeIndicatorFixedToVariable),
			LedgerTransaction: moderntreasury.F(shared.LedgerTransactionCreateRequestParam{
				LedgerEntries: moderntreasury.F([]shared.LedgerEntryCreateRequestParam{{
					Amount:          moderntreasury.F(int64(0)),
					Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					AvailableBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					LockVersion: moderntreasury.F(int64(0)),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
					PendingBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					PostedBalanceAmount: moderntreasury.F(map[string]int64{
						"foo": int64(0),
					}),
					ShowResultingLedgerAccountBalances: moderntreasury.F(true),
				}}),
				Description:    moderntreasury.F("description"),
				EffectiveAt:    moderntreasury.F(time.Now()),
				EffectiveDate:  moderntreasury.F(time.Now()),
				ExternalID:     moderntreasury.F("external_id"),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType: moderntreasury.F(shared.LedgerTransactionCreateRequestLedgerableTypeExpectedPayment),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Status: moderntreasury.F(shared.LedgerTransactionCreateRequestStatusArchived),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestLineItem{{
				Amount:               moderntreasury.F(int64(0)),
				AccountingCategoryID: moderntreasury.F("accounting_category_id"),
				Description:          moderntreasury.F("description"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			NsfProtected:         moderntreasury.F(true),
			OriginatingPartyName: moderntreasury.F("originating_party_name"),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestPriorityHigh),
			ProcessAfter:         moderntreasury.F(time.Now()),
			Purpose:              moderntreasury.F("purpose"),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccount{
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
				}}),
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeBaseWallet),
				ContactDetails: moderntreasury.F([]moderntreasury.ContactDetailCreateRequestParam{{
					ContactIdentifier:     moderntreasury.F("contact_identifier"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.ContactDetailCreateRequestContactIdentifierTypeEmail),
				}}),
				ExternalID: moderntreasury.F("external_id"),
				LedgerAccount: moderntreasury.F(shared.LedgerAccountCreateRequestParam{
					Currency:                 moderntreasury.F("currency"),
					LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Name:                     moderntreasury.F("name"),
					NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
					CurrencyExponent:         moderntreasury.F(int64(0)),
					Description:              moderntreasury.F("description"),
					LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
					LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:           moderntreasury.F(shared.LedgerAccountCreateRequestLedgerableTypeCounterparty),
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
				PartyAddress: moderntreasury.F(shared.AddressRequestParam{
					Country:    moderntreasury.F("country"),
					Line1:      moderntreasury.F("line1"),
					Line2:      moderntreasury.F("line2"),
					Locality:   moderntreasury.F("locality"),
					PostalCode: moderntreasury.F("postal_code"),
					Region:     moderntreasury.F("region"),
				}),
				PartyIdentifier:     moderntreasury.F("party_identifier"),
				PartyName:           moderntreasury.F("party_name"),
				PartyType:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountPartyTypeBusiness),
				PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
			}),
			ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReconciliationStatus:               moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderAsyncCreateRequestReconciliationStatusReconciled),
			RemittanceInformation:              moderntreasury.F("remittance_information"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			StatementDescriptor:                moderntreasury.F("statement_descriptor"),
			Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			TransactionMonitoringEnabled:       moderntreasury.F(true),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
			UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
			UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
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
