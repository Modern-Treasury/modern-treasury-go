// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
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
		option.WithAPIKey("APIKey"),
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
		Description:             moderntreasury.F("string"),
		Documents: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsDocument{{
			DocumentableID:   moderntreasury.F("string"),
			DocumentableType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsDocumentsDocumentableTypeCases),
			DocumentType:     moderntreasury.F("string"),
			File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		}, {
			DocumentableID:   moderntreasury.F("string"),
			DocumentableType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsDocumentsDocumentableTypeCases),
			DocumentType:     moderntreasury.F("string"),
			File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		}, {
			DocumentableID:   moderntreasury.F("string"),
			DocumentableType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsDocumentsDocumentableTypeCases),
			DocumentType:     moderntreasury.F("string"),
			File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
		}}),
		EffectiveDate:            moderntreasury.F(time.Now()),
		ExpiresAt:                moderntreasury.F(time.Now()),
		FallbackType:             moderntreasury.F(moderntreasury.PaymentOrderNewParamsFallbackTypeACH),
		ForeignExchangeContract:  moderntreasury.F("string"),
		ForeignExchangeIndicator: moderntreasury.F(moderntreasury.PaymentOrderNewParamsForeignExchangeIndicatorFixedToVariable),
		LedgerTransaction: moderntreasury.F(moderntreasury.PaymentOrderNewParamsLedgerTransaction{
			Description: moderntreasury.F("string"),
			Status:      moderntreasury.F(moderntreasury.PaymentOrderNewParamsLedgerTransactionStatusArchived),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			EffectiveAt:   moderntreasury.F(time.Now()),
			EffectiveDate: moderntreasury.F(time.Now()),
			LedgerEntries: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsLedgerTransactionLedgerEntry{{
				Amount:          moderntreasury.F(int64(0)),
				Direction:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirectionCredit),
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
				Direction:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirectionCredit),
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
				Direction:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsLedgerTransactionLedgerEntriesDirectionCredit),
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
			ExternalID:     moderntreasury.F("string"),
			LedgerableType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsLedgerTransactionLedgerableTypeCounterparty),
			LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		LineItems: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsLineItem{{
			Amount: moderntreasury.F(int64(0)),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Description:          moderntreasury.F("string"),
			AccountingCategoryID: moderntreasury.F("string"),
		}, {
			Amount: moderntreasury.F(int64(0)),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Description:          moderntreasury.F("string"),
			AccountingCategoryID: moderntreasury.F("string"),
		}, {
			Amount: moderntreasury.F(int64(0)),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Description:          moderntreasury.F("string"),
			AccountingCategoryID: moderntreasury.F("string"),
		}}),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		NsfProtected:         moderntreasury.F(true),
		OriginatingPartyName: moderntreasury.F("string"),
		Priority:             moderntreasury.F(moderntreasury.PaymentOrderNewParamsPriorityHigh),
		Purpose:              moderntreasury.F("string"),
		ReceivingAccount: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccount{
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountPartyAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			Name: moderntreasury.F("string"),
			AccountDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsReceivingAccountAccountDetail{{
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
			}}),
			RoutingDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetail{{
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("string"),
			PartyIdentifier: moderntreasury.F("string"),
			LedgerAccount: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountLedgerAccount{
				Name:             moderntreasury.F("string"),
				Description:      moderntreasury.F("string"),
				NormalBalance:    moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountLedgerAccountNormalBalanceCredit),
				LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:         moderntreasury.F("string"),
				CurrencyExponent: moderntreasury.F(int64(0)),
				LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:   moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountLedgerAccountLedgerableTypeExternalAccount),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			PlaidProcessorToken: moderntreasury.F("string"),
			ContactDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewParamsReceivingAccountContactDetail{{
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderNewParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
			}}),
		}),
		ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RemittanceInformation:              moderntreasury.F("string"),
		SendRemittanceAdvice:               moderntreasury.F(true),
		StatementDescriptor:                moderntreasury.F("string"),
		Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
		TransactionMonitoringEnabled:       moderntreasury.F(true),
		UltimateOriginatingPartyIdentifier: moderntreasury.F("string"),
		UltimateOriginatingPartyName:       moderntreasury.F("string"),
		UltimateReceivingPartyIdentifier:   moderntreasury.F("string"),
		UltimateReceivingPartyName:         moderntreasury.F("string"),
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
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.PaymentOrders.Get(context.TODO(), "string")
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
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.PaymentOrders.Update(
		context.TODO(),
		"string",
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
			Description:              moderntreasury.F("string"),
			Direction:                moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsDirectionCredit),
			EffectiveDate:            moderntreasury.F(time.Now()),
			ExpiresAt:                moderntreasury.F(time.Now()),
			FallbackType:             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsFallbackTypeACH),
			ForeignExchangeContract:  moderntreasury.F("string"),
			ForeignExchangeIndicator: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsForeignExchangeIndicatorFixedToVariable),
			LineItems: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsLineItem{{
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("string"),
				AccountingCategoryID: moderntreasury.F("string"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("string"),
				AccountingCategoryID: moderntreasury.F("string"),
			}, {
				Amount: moderntreasury.F(int64(0)),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				Description:          moderntreasury.F("string"),
				AccountingCategoryID: moderntreasury.F("string"),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			NsfProtected:         moderntreasury.F(true),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			OriginatingPartyName: moderntreasury.F("string"),
			Priority:             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsPriorityHigh),
			Purpose:              moderntreasury.F("string"),
			ReceivingAccount: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("string"),
					Line2:      moderntreasury.F("string"),
					Locality:   moderntreasury.F("string"),
					Region:     moderntreasury.F("string"),
					PostalCode: moderntreasury.F("string"),
					Country:    moderntreasury.F("string"),
				}),
				Name: moderntreasury.F("string"),
				AccountDetails: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("string"),
				PartyIdentifier: moderntreasury.F("string"),
				LedgerAccount: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountLedgerAccount{
					Name:             moderntreasury.F("string"),
					Description:      moderntreasury.F("string"),
					NormalBalance:    moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountLedgerAccountNormalBalanceCredit),
					LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:         moderntreasury.F("string"),
					CurrencyExponent: moderntreasury.F(int64(0)),
					LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:   moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountLedgerAccountLedgerableTypeExternalAccount),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("string"),
				ContactDetails: moderntreasury.F([]moderntreasury.PaymentOrderUpdateParamsReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RemittanceInformation:              moderntreasury.F("string"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			StatementDescriptor:                moderntreasury.F("string"),
			Status:                             moderntreasury.F(moderntreasury.PaymentOrderUpdateParamsStatusApproved),
			Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Type:                               moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("string"),
			UltimateOriginatingPartyName:       moderntreasury.F("string"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("string"),
			UltimateReceivingPartyName:         moderntreasury.F("string"),
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
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.PaymentOrders.List(context.TODO(), moderntreasury.PaymentOrderListParams{
		AfterCursor:        moderntreasury.F("string"),
		CounterpartyID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Direction:          moderntreasury.F(moderntreasury.PaymentOrderListParamsDirectionCredit),
		EffectiveDateEnd:   moderntreasury.F(time.Now()),
		EffectiveDateStart: moderntreasury.F(time.Now()),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PerPage:              moderntreasury.F(int64(0)),
		Priority:             moderntreasury.F(moderntreasury.PaymentOrderListParamsPriorityHigh),
		ReferenceNumber:      moderntreasury.F("string"),
		Status:               moderntreasury.F(moderntreasury.PaymentOrderListParamsStatusApproved),
		TransactionID:        moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		option.WithAPIKey("APIKey"),
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
		Description:              moderntreasury.F("string"),
		EffectiveDate:            moderntreasury.F(time.Now()),
		ExpiresAt:                moderntreasury.F(time.Now()),
		FallbackType:             moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsFallbackTypeACH),
		ForeignExchangeContract:  moderntreasury.F("string"),
		ForeignExchangeIndicator: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsForeignExchangeIndicatorFixedToVariable),
		LedgerTransaction: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsLedgerTransaction{
			Description: moderntreasury.F("string"),
			Status:      moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsLedgerTransactionStatusArchived),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			EffectiveAt:   moderntreasury.F(time.Now()),
			EffectiveDate: moderntreasury.F(time.Now()),
			LedgerEntries: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntry{{
				Amount:          moderntreasury.F(int64(0)),
				Direction:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirectionCredit),
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
				Direction:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirectionCredit),
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
				Direction:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsLedgerTransactionLedgerEntriesDirectionCredit),
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
			ExternalID:     moderntreasury.F("string"),
			LedgerableType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsLedgerTransactionLedgerableTypeCounterparty),
			LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}),
		LineItems: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsLineItem{{
			Amount: moderntreasury.F(int64(0)),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Description:          moderntreasury.F("string"),
			AccountingCategoryID: moderntreasury.F("string"),
		}, {
			Amount: moderntreasury.F(int64(0)),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Description:          moderntreasury.F("string"),
			AccountingCategoryID: moderntreasury.F("string"),
		}, {
			Amount: moderntreasury.F(int64(0)),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Description:          moderntreasury.F("string"),
			AccountingCategoryID: moderntreasury.F("string"),
		}}),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		NsfProtected:         moderntreasury.F(true),
		OriginatingPartyName: moderntreasury.F("string"),
		Priority:             moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsPriorityHigh),
		Purpose:              moderntreasury.F("string"),
		ReceivingAccount: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccount{
			AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
			PartyType:   moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountPartyTypeBusiness),
			PartyAddress: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountPartyAddress{
				Line1:      moderntreasury.F("string"),
				Line2:      moderntreasury.F("string"),
				Locality:   moderntreasury.F("string"),
				Region:     moderntreasury.F("string"),
				PostalCode: moderntreasury.F("string"),
				Country:    moderntreasury.F("string"),
			}),
			Name: moderntreasury.F("string"),
			AccountDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountAccountDetail{{
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
			}, {
				AccountNumber:     moderntreasury.F("string"),
				AccountNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountAccountDetailsAccountNumberTypeIban),
			}}),
			RoutingDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetail{{
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}, {
				RoutingNumber:     moderntreasury.F("string"),
				RoutingNumberType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsRoutingNumberTypeAba),
				PaymentType:       moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountRoutingDetailsPaymentTypeACH),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			PartyName:       moderntreasury.F("string"),
			PartyIdentifier: moderntreasury.F("string"),
			LedgerAccount: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountLedgerAccount{
				Name:             moderntreasury.F("string"),
				Description:      moderntreasury.F("string"),
				NormalBalance:    moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountNormalBalanceCredit),
				LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:         moderntreasury.F("string"),
				CurrencyExponent: moderntreasury.F(int64(0)),
				LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				LedgerableType:   moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountLedgerAccountLedgerableTypeExternalAccount),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}),
			PlaidProcessorToken: moderntreasury.F("string"),
			ContactDetails: moderntreasury.F([]moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountContactDetail{{
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
			}, {
				ContactIdentifier:     moderntreasury.F("string"),
				ContactIdentifierType: moderntreasury.F(moderntreasury.PaymentOrderNewAsyncParamsReceivingAccountContactDetailsContactIdentifierTypeEmail),
			}}),
		}),
		ReceivingAccountID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		RemittanceInformation:              moderntreasury.F("string"),
		SendRemittanceAdvice:               moderntreasury.F(true),
		StatementDescriptor:                moderntreasury.F("string"),
		Subtype:                            moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
		TransactionMonitoringEnabled:       moderntreasury.F(true),
		UltimateOriginatingPartyIdentifier: moderntreasury.F("string"),
		UltimateOriginatingPartyName:       moderntreasury.F("string"),
		UltimateReceivingPartyIdentifier:   moderntreasury.F("string"),
		UltimateReceivingPartyName:         moderntreasury.F("string"),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
