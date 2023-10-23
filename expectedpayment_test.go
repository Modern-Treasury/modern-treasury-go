// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
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
		Direction:         moderntreasury.F(shared.TransactionDirectionCredit),
		InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Currency:          moderntreasury.F(shared.CurrencyAed),
		DateLowerBound:    moderntreasury.F(time.Now()),
		DateUpperBound:    moderntreasury.F(time.Now()),
		Description:       moderntreasury.F("string"),
		LineItems: moderntreasury.F([]moderntreasury.ExpectedPaymentNewParamsLineItem{{
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
		ReconciliationFilters: moderntreasury.F[any](map[string]interface{}{}),
		ReconciliationGroups:  moderntreasury.F[any](map[string]interface{}{}),
		RemittanceInformation: moderntreasury.F("string"),
		StatementDescriptor:   moderntreasury.F("string"),
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
	_, err := client.ExpectedPayments.Get(context.TODO(), "string")
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
		"string",
		moderntreasury.ExpectedPaymentUpdateParams{
			AmountLowerBound:  moderntreasury.F(int64(0)),
			AmountUpperBound:  moderntreasury.F(int64(0)),
			CounterpartyID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Currency:          moderntreasury.F(shared.CurrencyAed),
			DateLowerBound:    moderntreasury.F(time.Now()),
			DateUpperBound:    moderntreasury.F(time.Now()),
			Description:       moderntreasury.F("string"),
			Direction:         moderntreasury.F(shared.TransactionDirectionCredit),
			InternalAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			ReconciliationFilters: moderntreasury.F[any](map[string]interface{}{}),
			ReconciliationGroups:  moderntreasury.F[any](map[string]interface{}{}),
			RemittanceInformation: moderntreasury.F("string"),
			StatementDescriptor:   moderntreasury.F("string"),
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
		AfterCursor:         moderntreasury.F("string"),
		CounterpartyID:      moderntreasury.F("string"),
		CreatedAtLowerBound: moderntreasury.F(time.Now()),
		CreatedAtUpperBound: moderntreasury.F(time.Now()),
		Direction:           moderntreasury.F(shared.TransactionDirectionCredit),
		InternalAccountID:   moderntreasury.F("string"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		PerPage: moderntreasury.F(int64(0)),
		Status:  moderntreasury.F(moderntreasury.ExpectedPaymentListParamsStatusArchived),
		Type:    moderntreasury.F(moderntreasury.ExpectedPaymentListParamsTypeACH),
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
	_, err := client.ExpectedPayments.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
