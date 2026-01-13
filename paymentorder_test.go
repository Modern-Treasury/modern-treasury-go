// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

func TestPaymentOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentOrders.New(context.TODO(), moderntreasury.PaymentOrderNewParams{
		Amount:               moderntreasury.F(int64(0)),
		Direction:            moderntreasury.F(moderntreasury.PaymentOrderNewParamsDirectionCredit),
		OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
		Accounting: moderntreasury.F(moderntreasury.PaymentOrderNewParamsAccounting{
			AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ChargeBearer:            moderntreasury.F(moderntreasury.PaymentOrderNewParamsChargeBearerShared),
		Currency:                moderntreasury.F(shared.CurrencyAed),
		Description:             moderntreasury.F("description"),
		Documents: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsDocument{{
			File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			DocumentType:     moderntreasury.F("document_type"),
			DocumentableID:   moderntreasury.F("documentable_id"),
			DocumentableType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsDocumentsDocumentableTypeConnections),
		}}),
		EffectiveDate:            moderntreasury.F(time.Now()),
		ExpiresAt:                moderntreasury.F(time.Now()),
		FallbackType:             moderntreasury.F(moderntreasury.PaymentOrderNewParamsFallbackTypeACH),
		ForeignExchangeContract:  moderntreasury.F("foreign_exchange_contract"),
		ForeignExchangeIndicator: moderntreasury.F(moderntreasury.PaymentOrderNewParamsForeignExchangeIndicatorFixedToVariable),
		LedgerTransaction: moderntreasury.F(shared.LedgerTransactionCreateRequestParam{
			LedgerEntries: moderntreasury.F([]shared.LedgerEntryCreateRequestParam{{
				Amount:          moderntreasury.F(int64(0)),
				Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AvailableBalanceAmount: moderntreasury.F(map[string]int64{
					"foo": int64(0),
				}),
				EffectiveAt: moderntreasury.F(time.Now()),
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
		LineItems: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsLineItem{{
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
		Priority:             moderntreasury.F(moderntreasury.PaymentOrderNewParamsPriorityHigh),
		ProcessAfter:         moderntreasury.F(time.Now()),
		Purpose:              moderntreasury.F("purpose"),
		ReceivingAccount: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccount{
			AccountDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsReceivingAccountAccountDetail{{
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
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
			PartyType:           moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountPartyTypeBusiness),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
			RoutingDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetail{{
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}}),
		}),
		ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReconciliationStatus:               moderntreasury.F(moderntreasury.PaymentOrderNewParamsReconciliationStatusUnreconciled),
		RemittanceInformation:              moderntreasury.F("remittance_information"),
		SendRemittanceAdvice:               moderntreasury.F(true),
		StatementDescriptor:                moderntreasury.F("statement_descriptor"),
		Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
		TransactionMonitoringEnabled:       moderntreasury.F(true),
		UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
		UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
		UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
		UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentOrderGet(t *testing.T) {
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
	_, err := client.PaymentOrders.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentOrders.Update(
		context.TODO(),
		"id",
		moderntreasury.PaymentOrderUpdateParams{
			Accounting: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Amount:                   moderntreasury.F(int64(0)),
			ChargeBearer:             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsChargeBearerShared),
			CounterpartyID:           moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                 moderntreasury.F(shared.CurrencyAed),
			Description:              moderntreasury.F("description"),
			Direction:                moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsDirectionCredit),
			EffectiveDate:            moderntreasury.F(time.Now()),
			ExpiresAt:                moderntreasury.F(time.Now()),
			FallbackType:             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsFallbackTypeACH),
			ForeignExchangeContract:  moderntreasury.F("foreign_exchange_contract"),
			ForeignExchangeIndicator: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsForeignExchangeIndicatorFixedToVariable),
			LineItems: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsLineItem{{
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
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			OriginatingPartyName: moderntreasury.F("originating_party_name"),
			Priority:             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsPriorityHigh),
			ProcessAfter:         moderntreasury.F(time.Now()),
			Purpose:              moderntreasury.F("purpose"),
			ReceivingAccount: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccount{
				AccountDetails: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("account_number"),
					AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
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
				PartyType:           moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountPartyTypeBusiness),
				PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
				RoutingDetails: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("routing_number"),
					RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
			}),
			ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReconciliationStatus:               moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReconciliationStatusUnreconciled),
			RemittanceInformation:              moderntreasury.F("remittance_information"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			StatementDescriptor:                moderntreasury.F("statement_descriptor"),
			Status:                             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsStatusApproved),
			Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Type:                               moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
			UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
			UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
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

func TestPaymentOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentOrders.List(context.TODO(), moderntreasury.PaymentOrderListParams{
		AfterCursor:        moderntreasury.F("after_cursor"),
		CounterpartyID:     moderntreasury.F("counterparty_id"),
		CreatedAtEnd:       moderntreasury.F(time.Now()),
		CreatedAtStart:     moderntreasury.F(time.Now()),
		Direction:          moderntreasury.F(shared.TransactionDirectionCredit),
		EffectiveDateEnd:   moderntreasury.F(time.Now()),
		EffectiveDateStart: moderntreasury.F(time.Now()),
		ExternalID:         moderntreasury.F("external_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		OriginatingAccountID: moderntreasury.F("originating_account_id"),
		PerPage:              moderntreasury.F(int64(0)),
		Priority:             moderntreasury.F(moderntreasury.PaymentOrderListParamsPriorityHigh),
		ProcessAfterEnd:      moderntreasury.F(time.Now()),
		ProcessAfterStart:    moderntreasury.F(time.Now()),
		ReferenceNumber:      moderntreasury.F("reference_number"),
		Status:               moderntreasury.F(moderntreasury.PaymentOrderListParamsStatusApproved),
		TransactionID:        moderntreasury.F("transaction_id"),
		Type:                 moderntreasury.F(moderntreasury.PaymentOrderListParamsTypeACH),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentOrderNewAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentOrders.NewAsync(context.TODO(), moderntreasury.PaymentOrderNewAsyncParams{
		Amount:               moderntreasury.F(int64(0)),
		Direction:            moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsDirectionCredit),
		OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
		Accounting: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsAccounting{
			AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		AccountingCategoryID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AccountingLedgerClassID:  moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ChargeBearer:             moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsChargeBearerShared),
		Currency:                 moderntreasury.F(shared.CurrencyAed),
		Description:              moderntreasury.F("description"),
		EffectiveDate:            moderntreasury.F(time.Now()),
		ExpiresAt:                moderntreasury.F(time.Now()),
		FallbackType:             moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsFallbackTypeACH),
		ForeignExchangeContract:  moderntreasury.F("foreign_exchange_contract"),
		ForeignExchangeIndicator: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsForeignExchangeIndicatorFixedToVariable),
		LedgerTransaction: moderntreasury.F(shared.LedgerTransactionCreateRequestParam{
			LedgerEntries: moderntreasury.F([]shared.LedgerEntryCreateRequestParam{{
				Amount:          moderntreasury.F(int64(0)),
				Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
				LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				AvailableBalanceAmount: moderntreasury.F(map[string]int64{
					"foo": int64(0),
				}),
				EffectiveAt: moderntreasury.F(time.Now()),
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
		LineItems: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsLineItem{{
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
		Priority:             moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsPriorityHigh),
		ProcessAfter:         moderntreasury.F(time.Now()),
		Purpose:              moderntreasury.F("purpose"),
		ReceivingAccount: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccount{
			AccountDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountAccountDetail{{
				AccountNumber:     moderntreasury.F("account_number"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeAuNumber),
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
			PartyType:           moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountPartyTypeBusiness),
			PlaidProcessorToken: moderntreasury.F("plaid_processor_token"),
			RoutingDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetail{{
				RoutingNumber:     moderntreasury.F("routing_number"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}}),
		}),
		ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ReconciliationStatus:               moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReconciliationStatusUnreconciled),
		RemittanceInformation:              moderntreasury.F("remittance_information"),
		SendRemittanceAdvice:               moderntreasury.F(true),
		StatementDescriptor:                moderntreasury.F("statement_descriptor"),
		Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
		TransactionMonitoringEnabled:       moderntreasury.F(true),
		UltimateOriginatingPartyIdentifier: moderntreasury.F("ultimate_originating_party_identifier"),
		UltimateOriginatingPartyName:       moderntreasury.F("ultimate_originating_party_name"),
		UltimateReceivingPartyIdentifier:   moderntreasury.F("ultimate_receiving_party_identifier"),
		UltimateReceivingPartyName:         moderntreasury.F("ultimate_receiving_party_name"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
