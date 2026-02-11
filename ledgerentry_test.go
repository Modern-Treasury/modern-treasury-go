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

func TestLedgerEntryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerEntries.Get(
		context.TODO(),
		"id",
		moderntreasury.LedgerEntryGetParams{
			ShowBalances: moderntreasury.F(true),
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

func TestLedgerEntryUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerEntries.Update(
		context.TODO(),
		"id",
		moderntreasury.LedgerEntryUpdateParams{
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
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

func TestLedgerEntryListWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerEntries.List(context.TODO(), moderntreasury.LedgerEntryListParams{
		ID:          moderntreasury.F([]string{"string"}),
		AfterCursor: moderntreasury.F("after_cursor"),
		Amount: moderntreasury.F(moderntreasury.LedgerEntryListParamsAmount{
			Eq:  moderntreasury.F(int64(0)),
			Gt:  moderntreasury.F(int64(0)),
			Gte: moderntreasury.F(int64(0)),
			Lt:  moderntreasury.F(int64(0)),
			Lte: moderntreasury.F(int64(0)),
		}),
		AsOfLockVersion: moderntreasury.F(int64(0)),
		Direction:       moderntreasury.F(shared.TransactionDirectionCredit),
		EffectiveAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		EffectiveDate: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		LedgerAccountCategoryID: moderntreasury.F("ledger_account_category_id"),
		LedgerAccountID:         moderntreasury.F("ledger_account_id"),
		LedgerAccountLockVersion: moderntreasury.F(map[string]int64{
			"foo": int64(0),
		}),
		LedgerAccountPayoutID:     moderntreasury.F("ledger_account_payout_id"),
		LedgerAccountSettlementID: moderntreasury.F("ledger_account_settlement_id"),
		LedgerAccountStatementID:  moderntreasury.F("ledger_account_statement_id"),
		LedgerTransactionID:       moderntreasury.F("ledger_transaction_id"),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		OrderBy: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderBy{
			CreatedAt:   moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderByCreatedAtAsc),
			EffectiveAt: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderByEffectiveAtAsc),
		}),
		PerPage:      moderntreasury.F(int64(0)),
		ShowBalances: moderntreasury.F(true),
		ShowDeleted:  moderntreasury.F(true),
		Status:       moderntreasury.F(moderntreasury.LedgerEntryListParamsStatusPending),
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
