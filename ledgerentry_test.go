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
		"string",
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
		ID:              moderntreasury.F([]string{"string", "string", "string"}),
		AfterCursor:     moderntreasury.F("string"),
		AsOfLockVersion: moderntreasury.F(int64(0)),
		Direction:       moderntreasury.F(moderntreasury.LedgerEntryListParamsDirectionCredit),
		EffectiveAt: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		EffectiveDate: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		LedgerAccountCategoryID: moderntreasury.F("string"),
		LedgerAccountID:         moderntreasury.F("string"),
		LedgerAccountLockVersion: moderntreasury.F(map[string]int64{
			"foo": int64(0),
		}),
		LedgerAccountPayoutID:    moderntreasury.F("string"),
		LedgerAccountStatementID: moderntreasury.F("string"),
		LedgerTransactionID:      moderntreasury.F("string"),
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
