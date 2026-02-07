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

func TestLedgerAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerAccounts.New(context.TODO(), moderntreasury.LedgerAccountNewParams{
		LedgerAccountCreateRequest: shared.LedgerAccountCreateRequestParam{
			Currency:                 moderntreasury.F("currency"),
			LedgerID:                 moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                     moderntreasury.F("name"),
			NormalBalance:            moderntreasury.F(shared.TransactionDirectionCredit),
			CurrencyExponent:         moderntreasury.F(int64(0)),
			Description:              moderntreasury.F("description"),
			ExternalID:               moderntreasury.F("external_id"),
			LedgerAccountCategoryIDs: moderntreasury.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			LedgerableID:             moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType:           moderntreasury.F(shared.LedgerAccountCreateRequestLedgerableTypeCounterparty),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
		},
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerAccountGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerAccounts.Get(
		context.TODO(),
		"id",
		moderntreasury.LedgerAccountGetParams{
			Balances: moderntreasury.F(moderntreasury.LedgerAccountGetParamsBalances{
				AsOfDate:              moderntreasury.F(time.Now()),
				AsOfLockVersion:       moderntreasury.F(int64(0)),
				EffectiveAt:           moderntreasury.F(time.Now()),
				EffectiveAtLowerBound: moderntreasury.F(time.Now()),
				EffectiveAtUpperBound: moderntreasury.F(time.Now()),
			}),
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

func TestLedgerAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerAccounts.Update(
		context.TODO(),
		"id",
		moderntreasury.LedgerAccountUpdateParams{
			Description: moderntreasury.F("description"),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Name: moderntreasury.F("name"),
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

func TestLedgerAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerAccounts.List(context.TODO(), moderntreasury.LedgerAccountListParams{
		ID:          moderntreasury.F([]string{"string"}),
		AfterCursor: moderntreasury.F("after_cursor"),
		AvailableBalanceAmount: moderntreasury.F(moderntreasury.LedgerAccountListParamsAvailableBalanceAmount{
			Eq:    moderntreasury.F(int64(0)),
			Gt:    moderntreasury.F(int64(0)),
			Gte:   moderntreasury.F(int64(0)),
			Lt:    moderntreasury.F(int64(0)),
			Lte:   moderntreasury.F(int64(0)),
			NotEq: moderntreasury.F(int64(0)),
		}),
		Balances: moderntreasury.F(moderntreasury.LedgerAccountListParamsBalances{
			AsOfDate:              moderntreasury.F(time.Now()),
			EffectiveAt:           moderntreasury.F(time.Now()),
			EffectiveAtLowerBound: moderntreasury.F(time.Now()),
			EffectiveAtUpperBound: moderntreasury.F(time.Now()),
		}),
		CreatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		Currency:                moderntreasury.F("currency"),
		ExternalID:              moderntreasury.F("external_id"),
		LedgerAccountCategoryID: moderntreasury.F("ledger_account_category_id"),
		LedgerID:                moderntreasury.F("ledger_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		Name:          moderntreasury.F([]string{"string"}),
		NormalBalance: moderntreasury.F(shared.TransactionDirectionCredit),
		PendingBalanceAmount: moderntreasury.F(moderntreasury.LedgerAccountListParamsPendingBalanceAmount{
			Eq:    moderntreasury.F(int64(0)),
			Gt:    moderntreasury.F(int64(0)),
			Gte:   moderntreasury.F(int64(0)),
			Lt:    moderntreasury.F(int64(0)),
			Lte:   moderntreasury.F(int64(0)),
			NotEq: moderntreasury.F(int64(0)),
		}),
		PerPage: moderntreasury.F(int64(0)),
		PostedBalanceAmount: moderntreasury.F(moderntreasury.LedgerAccountListParamsPostedBalanceAmount{
			Eq:    moderntreasury.F(int64(0)),
			Gt:    moderntreasury.F(int64(0)),
			Gte:   moderntreasury.F(int64(0)),
			Lt:    moderntreasury.F(int64(0)),
			Lte:   moderntreasury.F(int64(0)),
			NotEq: moderntreasury.F(int64(0)),
		}),
		UpdatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
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

func TestLedgerAccountDelete(t *testing.T) {
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
	_, err := client.LedgerAccounts.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
