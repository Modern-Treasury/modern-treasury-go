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

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

func TestBulkRequestNewWithOptionalParams(t *testing.T) {
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
		Resources: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResource{moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequest(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequest{
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Subtype:              moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirectionCredit),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriorityHigh),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                moderntreasury.F(shared.CurrencyAed),
			EffectiveDate:           moderntreasury.F(time.Now()),
			Description:             moderntreasury.F("string"),
			StatementDescriptor:     moderntreasury.F("string"),
			RemittanceInformation:   moderntreasury.F("string"),
			Purpose:                 moderntreasury.F("string"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ChargeBearer:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearerShared),
			ForeignExchangeIndicator:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicatorFixedToVariable),
			ForeignExchangeContract:            moderntreasury.F("string"),
			NsfProtected:                       moderntreasury.F(true),
			OriginatingPartyName:               moderntreasury.F("string"),
			UltimateOriginatingPartyName:       moderntreasury.F("string"),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("string"),
			UltimateReceivingPartyName:         moderntreasury.F("string"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("string"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			ExpiresAt:                          moderntreasury.F(time.Now()),
			FallbackType:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackTypeACH),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("string"),
					Line2:      moderntreasury.F("string"),
					Locality:   moderntreasury.F("string"),
					Region:     moderntreasury.F("string"),
					PostalCode: moderntreasury.F("string"),
					Country:    moderntreasury.F("string"),
				}),
				Name: moderntreasury.F("string"),
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("string"),
				PartyIdentifier: moderntreasury.F("string"),
				LedgerAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccount{
					Name:             moderntreasury.F("string"),
					Description:      moderntreasury.F("string"),
					NormalBalance:    moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:         moderntreasury.F("string"),
					CurrencyExponent: moderntreasury.F(int64(0)),
					LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableTypeExternalAccount),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("string"),
				ContactDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			LedgerTransaction: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransaction{
				Description: moderntreasury.F("string"),
				Status:      moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatusArchived),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				EffectiveAt:   moderntreasury.F(time.Now()),
				EffectiveDate: moderntreasury.F(time.Now()),
				LedgerEntries: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerEntry{{
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
				ExternalID:     moderntreasury.F("string"),
				LedgerableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeCounterparty),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLineItem{{
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
			TransactionMonitoringEnabled: moderntreasury.F(true),
			Documents: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocument{{
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}, {
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}, {
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}}),
		}), moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequest(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequest{
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Subtype:              moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirectionCredit),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriorityHigh),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                moderntreasury.F(shared.CurrencyAed),
			EffectiveDate:           moderntreasury.F(time.Now()),
			Description:             moderntreasury.F("string"),
			StatementDescriptor:     moderntreasury.F("string"),
			RemittanceInformation:   moderntreasury.F("string"),
			Purpose:                 moderntreasury.F("string"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ChargeBearer:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearerShared),
			ForeignExchangeIndicator:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicatorFixedToVariable),
			ForeignExchangeContract:            moderntreasury.F("string"),
			NsfProtected:                       moderntreasury.F(true),
			OriginatingPartyName:               moderntreasury.F("string"),
			UltimateOriginatingPartyName:       moderntreasury.F("string"),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("string"),
			UltimateReceivingPartyName:         moderntreasury.F("string"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("string"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			ExpiresAt:                          moderntreasury.F(time.Now()),
			FallbackType:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackTypeACH),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("string"),
					Line2:      moderntreasury.F("string"),
					Locality:   moderntreasury.F("string"),
					Region:     moderntreasury.F("string"),
					PostalCode: moderntreasury.F("string"),
					Country:    moderntreasury.F("string"),
				}),
				Name: moderntreasury.F("string"),
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("string"),
				PartyIdentifier: moderntreasury.F("string"),
				LedgerAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccount{
					Name:             moderntreasury.F("string"),
					Description:      moderntreasury.F("string"),
					NormalBalance:    moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:         moderntreasury.F("string"),
					CurrencyExponent: moderntreasury.F(int64(0)),
					LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableTypeExternalAccount),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("string"),
				ContactDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			LedgerTransaction: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransaction{
				Description: moderntreasury.F("string"),
				Status:      moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatusArchived),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				EffectiveAt:   moderntreasury.F(time.Now()),
				EffectiveDate: moderntreasury.F(time.Now()),
				LedgerEntries: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerEntry{{
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
				ExternalID:     moderntreasury.F("string"),
				LedgerableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeCounterparty),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLineItem{{
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
			TransactionMonitoringEnabled: moderntreasury.F(true),
			Documents: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocument{{
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}, {
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}, {
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}}),
		}), moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequest(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequest{
			Type:                 moderntreasury.F(moderntreasury.PaymentOrderTypeACH),
			Subtype:              moderntreasury.F(moderntreasury.PaymentOrderSubtypeBacsNewInstruction),
			Amount:               moderntreasury.F(int64(0)),
			Direction:            moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDirectionCredit),
			Priority:             moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestPriorityHigh),
			OriginatingAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ReceivingAccountID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Accounting: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestAccounting{
				AccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ClassID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			AccountingCategoryID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccountingLedgerClassID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:                moderntreasury.F(shared.CurrencyAed),
			EffectiveDate:           moderntreasury.F(time.Now()),
			Description:             moderntreasury.F("string"),
			StatementDescriptor:     moderntreasury.F("string"),
			RemittanceInformation:   moderntreasury.F("string"),
			Purpose:                 moderntreasury.F("string"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ChargeBearer:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestChargeBearerShared),
			ForeignExchangeIndicator:           moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestForeignExchangeIndicatorFixedToVariable),
			ForeignExchangeContract:            moderntreasury.F("string"),
			NsfProtected:                       moderntreasury.F(true),
			OriginatingPartyName:               moderntreasury.F("string"),
			UltimateOriginatingPartyName:       moderntreasury.F("string"),
			UltimateOriginatingPartyIdentifier: moderntreasury.F("string"),
			UltimateReceivingPartyName:         moderntreasury.F("string"),
			UltimateReceivingPartyIdentifier:   moderntreasury.F("string"),
			SendRemittanceAdvice:               moderntreasury.F(true),
			ExpiresAt:                          moderntreasury.F(time.Now()),
			FallbackType:                       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestFallbackTypeACH),
			ReceivingAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccount{
				AccountType: moderntreasury.F(moderntreasury.ExternalAccountTypeCash),
				PartyType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyTypeBusiness),
				PartyAddress: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountPartyAddress{
					Line1:      moderntreasury.F("string"),
					Line2:      moderntreasury.F("string"),
					Locality:   moderntreasury.F("string"),
					Region:     moderntreasury.F("string"),
					PostalCode: moderntreasury.F("string"),
					Country:    moderntreasury.F("string"),
				}),
				Name: moderntreasury.F("string"),
				AccountDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetail{{
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}, {
					AccountNumber:     moderntreasury.F("string"),
					AccountNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountAccountDetailsAccountNumberTypeIban),
				}}),
				RoutingDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetail{{
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}, {
					RoutingNumber:     moderntreasury.F("string"),
					RoutingNumberType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsRoutingNumberTypeAba),
					PaymentType:       moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountRoutingDetailsPaymentTypeACH),
				}}),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				PartyName:       moderntreasury.F("string"),
				PartyIdentifier: moderntreasury.F("string"),
				LedgerAccount: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccount{
					Name:             moderntreasury.F("string"),
					Description:      moderntreasury.F("string"),
					NormalBalance:    moderntreasury.F(shared.TransactionDirectionCredit),
					LedgerID:         moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Currency:         moderntreasury.F("string"),
					CurrencyExponent: moderntreasury.F(int64(0)),
					LedgerableID:     moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					LedgerableType:   moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountLedgerAccountLedgerableTypeExternalAccount),
					Metadata: moderntreasury.F(map[string]string{
						"key":    "value",
						"foo":    "bar",
						"modern": "treasury",
					}),
				}),
				PlaidProcessorToken: moderntreasury.F("string"),
				ContactDetails: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetail{{
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}, {
					ContactIdentifier:     moderntreasury.F("string"),
					ContactIdentifierType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestReceivingAccountContactDetailsContactIdentifierTypeEmail),
				}}),
			}),
			LedgerTransaction: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransaction{
				Description: moderntreasury.F("string"),
				Status:      moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionStatusArchived),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
				EffectiveAt:   moderntreasury.F(time.Now()),
				EffectiveDate: moderntreasury.F(time.Now()),
				LedgerEntries: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerEntry{{
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
				ExternalID:     moderntreasury.F("string"),
				LedgerableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLedgerTransactionLedgerableTypeCounterparty),
				LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			LedgerTransactionID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LineItems: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestLineItem{{
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
			TransactionMonitoringEnabled: moderntreasury.F(true),
			Documents: moderntreasury.F([]moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocument{{
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}, {
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}, {
				DocumentableID:   moderntreasury.F("string"),
				DocumentableType: moderntreasury.F(moderntreasury.BulkRequestNewParamsResourcesPaymentOrderCreateRequestDocumentsDocumentableTypeCases),
				DocumentType:     moderntreasury.F("string"),
				File:             moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			}}),
		})}),
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
	_, err := client.BulkRequests.Get(context.TODO(), "string")
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
		AfterCursor: moderntreasury.F("string"),
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
