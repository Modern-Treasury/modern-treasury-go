// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerEventHandlerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerEventHandlers.New(context.TODO(), moderntreasury.LedgerEventHandlerNewParams{
		LedgerTransactionTemplate: moderntreasury.F(moderntreasury.LedgerEventHandlerNewParamsLedgerTransactionTemplate{
			Description: moderntreasury.F("string"),
			EffectiveAt: moderntreasury.F("string"),
			LedgerEntries: moderntreasury.F([]moderntreasury.LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntry{{
				Amount:          moderntreasury.F("string"),
				Direction:       moderntreasury.F("string"),
				LedgerAccountID: moderntreasury.F("string"),
			}, {
				Amount:          moderntreasury.F("string"),
				Direction:       moderntreasury.F("string"),
				LedgerAccountID: moderntreasury.F("string"),
			}, {
				Amount:          moderntreasury.F("string"),
				Direction:       moderntreasury.F("string"),
				LedgerAccountID: moderntreasury.F("string"),
			}}),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		}),
		Name: moderntreasury.F("string"),
		Conditions: moderntreasury.F(moderntreasury.LedgerEventHandlerNewParamsConditions{
			Field:    moderntreasury.F("string"),
			Operator: moderntreasury.F("string"),
			Value:    moderntreasury.F("string"),
		}),
		Description: moderntreasury.F("string"),
		LedgerID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestLedgerEventHandlerGet(t *testing.T) {
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
	_, err := client.LedgerEventHandlers.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerEventHandlerListWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerEventHandlers.List(context.TODO(), moderntreasury.LedgerEventHandlerListParams{
		AfterCursor: moderntreasury.F("string"),
		CreatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
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

func TestLedgerEventHandlerDelete(t *testing.T) {
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
	_, err := client.LedgerEventHandlers.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
