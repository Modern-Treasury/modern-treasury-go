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
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerEventHandlers.New(context.TODO(), moderntreasury.LedgerEventHandlerNewParams{
		LedgerTransactionTemplate: moderntreasury.F(moderntreasury.LedgerEventHandlerNewParamsLedgerTransactionTemplate{
			Description: moderntreasury.F("My Ledger Transaction Template Description"),
			EffectiveAt: moderntreasury.F("{{ledgerable_event.custom_data.effective_at}}"),
			Status:      moderntreasury.F("posted"),
			LedgerEntries: moderntreasury.F([]moderntreasury.LedgerEventHandlerNewParamsLedgerTransactionTemplateLedgerEntry{{
				Amount:          moderntreasury.F("amount"),
				Direction:       moderntreasury.F("direction"),
				LedgerAccountID: moderntreasury.F("ledger_account_id"),
			}, {
				Amount:          moderntreasury.F("amount"),
				Direction:       moderntreasury.F("direction"),
				LedgerAccountID: moderntreasury.F("ledger_account_id"),
			}, {
				Amount:          moderntreasury.F("amount"),
				Direction:       moderntreasury.F("direction"),
				LedgerAccountID: moderntreasury.F("ledger_account_id"),
			}}),
		}),
		Name: moderntreasury.F("name"),
		Conditions: moderntreasury.F(moderntreasury.LedgerEventHandlerNewParamsConditions{
			Field:    moderntreasury.F("ledgerable_event.name"),
			Operator: moderntreasury.F("equals"),
			Value:    moderntreasury.F("credit_card_swipe"),
		}),
		Description: moderntreasury.F("description"),
		LedgerID:    moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		Variables: moderntreasury.F(map[string]moderntreasury.LedgerEventHandlerVariableParam{
			"credit_account": {
				Type: moderntreasury.F("ledger_account"),
				Query: moderntreasury.F(moderntreasury.LedgerEventHandlerVariableQueryParam{
					Field:    moderntreasury.F("name"),
					Operator: moderntreasury.F("equals"),
					Value:    moderntreasury.F("my_credit_account"),
				}),
			},
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
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerEventHandlers.Get(context.TODO(), "id")
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
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerEventHandlers.List(context.TODO(), moderntreasury.LedgerEventHandlerListParams{
		AfterCursor: moderntreasury.F("after_cursor"),
		CreatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
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
		option.WithAPIKey("My API Key"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerEventHandlers.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
