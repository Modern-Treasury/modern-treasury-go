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

func TestExpectedPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ExpectedPayments.New(context.TODO(), moderntreasury.ExpectedPaymentNewParams{
		AmountLowerBound:  moderntreasury.F(int64(0)),
		AmountUpperBound:  moderntreasury.F(int64(0)),
		CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Currency:          moderntreasury.F(shared.CurrencyAed),
		DateLowerBound:    moderntreasury.F(time.Now()),
		DateUpperBound:    moderntreasury.F(time.Now()),
		Description:       moderntreasury.F("description"),
		Direction:         moderntreasury.F(moderntreasury.ExpectedPaymentNewParamsDirectionCredit),
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
		LineItems: moderntreasury.F([]moderntreasury.ExpectedPaymentNewParamsLineItem{{
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
		ReconciliationFilters: moderntreasury.F[any](map[string]interface{}{}),
		ReconciliationGroups:  moderntreasury.F[any](map[string]interface{}{}),
		ReconciliationRuleVariables: moderntreasury.F([]moderntreasury.ReconciliationRuleParam{{
			AmountLowerBound:  moderntreasury.F(int64(0)),
			AmountUpperBound:  moderntreasury.F(int64(0)),
			Direction:         moderntreasury.F(moderntreasury.ReconciliationRuleDirectionCredit),
			InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:          moderntreasury.F(shared.CurrencyAed),
			CustomIdentifiers: moderntreasury.F(map[string]string{
				"foo": "string",
			}),
			DateLowerBound: moderntreasury.F(time.Now()),
			DateUpperBound: moderntreasury.F(time.Now()),
			Type:           moderntreasury.F(moderntreasury.ReconciliationRuleTypeACH),
		}}),
		RemittanceInformation: moderntreasury.F("remittance_information"),
		StatementDescriptor:   moderntreasury.F("statement_descriptor"),
		Type:                  moderntreasury.F(moderntreasury.ExpectedPaymentTypeACH),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentGet(t *testing.T) {
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
	_, err := client.ExpectedPayments.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ExpectedPayments.Update(
		context.TODO(),
		"id",
		moderntreasury.ExpectedPaymentUpdateParams{
			AmountLowerBound:  moderntreasury.F(int64(0)),
			AmountUpperBound:  moderntreasury.F(int64(0)),
			CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:          moderntreasury.F(shared.CurrencyAed),
			DateLowerBound:    moderntreasury.F(time.Now()),
			DateUpperBound:    moderntreasury.F(time.Now()),
			Description:       moderntreasury.F("description"),
			Direction:         moderntreasury.F(moderntreasury.ExpectedPaymentUpdateParamsDirectionCredit),
			InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ReconciliationFilters: moderntreasury.F[any](map[string]interface{}{}),
			ReconciliationGroups:  moderntreasury.F[any](map[string]interface{}{}),
			ReconciliationRuleVariables: moderntreasury.F([]moderntreasury.ReconciliationRuleParam{{
				AmountLowerBound:  moderntreasury.F(int64(0)),
				AmountUpperBound:  moderntreasury.F(int64(0)),
				Direction:         moderntreasury.F(moderntreasury.ReconciliationRuleDirectionCredit),
				InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Currency:          moderntreasury.F(shared.CurrencyAed),
				CustomIdentifiers: moderntreasury.F(map[string]string{
					"foo": "string",
				}),
				DateLowerBound: moderntreasury.F(time.Now()),
				DateUpperBound: moderntreasury.F(time.Now()),
				Type:           moderntreasury.F(moderntreasury.ReconciliationRuleTypeACH),
			}}),
			RemittanceInformation: moderntreasury.F("remittance_information"),
			StatementDescriptor:   moderntreasury.F("statement_descriptor"),
			Status:                moderntreasury.F(moderntreasury.ExpectedPaymentUpdateParamsStatusReconciled),
			Type:                  moderntreasury.F(moderntreasury.ExpectedPaymentTypeACH),
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

func TestExpectedPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExpectedPayments.List(context.TODO(), moderntreasury.ExpectedPaymentListParams{
		AfterCursor:         moderntreasury.F("after_cursor"),
		CounterpartyID:      moderntreasury.F("counterparty_id"),
		CreatedAtLowerBound: moderntreasury.F(time.Now()),
		CreatedAtUpperBound: moderntreasury.F(time.Now()),
		Direction:           moderntreasury.F(shared.TransactionDirectionCredit),
		InternalAccountID:   moderntreasury.F("internal_account_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage:             moderntreasury.F(int64(0)),
		Status:              moderntreasury.F(moderntreasury.ExpectedPaymentListParamsStatusArchived),
		Type:                moderntreasury.F(moderntreasury.ExpectedPaymentListParamsTypeACH),
		UpdatedAtLowerBound: moderntreasury.F(time.Now()),
		UpdatedAtUpperBound: moderntreasury.F(time.Now()),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExpectedPaymentDelete(t *testing.T) {
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
	_, err := client.ExpectedPayments.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
