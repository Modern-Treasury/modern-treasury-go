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

func TestLedgerTransactionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerTransactions.New(context.TODO(), moderntreasury.LedgerTransactionNewParams{
		LedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionNewParamsLedgerEntry{{
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
		LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerableTypeExpectedPayment),
		Metadata: moderntreasury.F(map[string]string{
			"key":    "value",
			"foo":    "bar",
			"modern": "treasury",
		}),
		Status: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsStatusArchived),
	})
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionGet(t *testing.T) {
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
	_, err := client.LedgerTransactions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerTransactionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerTransactions.Update(
		context.TODO(),
		"id",
		moderntreasury.LedgerTransactionUpdateParams{
			Description: moderntreasury.F("description"),
			EffectiveAt: moderntreasury.F(time.Now()),
			LedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionUpdateParamsLedgerEntry{{
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
			LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerableTypeExpectedPayment),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Status: moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsStatusArchived),
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

func TestLedgerTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerTransactions.List(context.TODO(), moderntreasury.LedgerTransactionListParams{
		ID:          moderntreasury.F([]string{"string"}),
		AfterCursor: moderntreasury.F("after_cursor"),
		EffectiveAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		EffectiveDate: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		ExternalID:                moderntreasury.F("external_id"),
		LedgerAccountCategoryID:   moderntreasury.F("ledger_account_category_id"),
		LedgerAccountID:           moderntreasury.F("ledger_account_id"),
		LedgerAccountSettlementID: moderntreasury.F("ledger_account_settlement_id"),
		LedgerID:                  moderntreasury.F("ledger_id"),
		LedgerableID:              moderntreasury.F("ledgerable_id"),
		LedgerableType:            moderntreasury.F(moderntreasury.LedgerTransactionListParamsLedgerableTypeExpectedPayment),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		OrderBy: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderBy{
			CreatedAt:   moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByCreatedAtAsc),
			EffectiveAt: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByEffectiveAtAsc),
		}),
		PartiallyPostsLedgerTransactionID: moderntreasury.F("partially_posts_ledger_transaction_id"),
		PerPage:                           moderntreasury.F(int64(0)),
		PostedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		ReversesLedgerTransactionID: moderntreasury.F("reverses_ledger_transaction_id"),
		Status:                      moderntreasury.F(moderntreasury.LedgerTransactionListParamsStatusPending),
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

func TestLedgerTransactionNewPartialPostWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerTransactions.NewPartialPost(
		context.TODO(),
		"id",
		moderntreasury.LedgerTransactionNewPartialPostParams{
			PostedLedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionNewPartialPostParamsPostedLedgerEntry{{
				Amount:          moderntreasury.F(int64(0)),
				Direction:       moderntreasury.F(moderntreasury.LedgerTransactionNewPartialPostParamsPostedLedgerEntriesDirectionCredit),
				LedgerAccountID: moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Metadata: moderntreasury.F(map[string]string{
					"key":    "value",
					"foo":    "bar",
					"modern": "treasury",
				}),
			}}),
			Description: moderntreasury.F("description"),
			EffectiveAt: moderntreasury.F(time.Now()),
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

func TestLedgerTransactionNewReversalWithOptionalParams(t *testing.T) {
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
	_, err := client.LedgerTransactions.NewReversal(
		context.TODO(),
		"id",
		moderntreasury.LedgerTransactionNewReversalParams{
			Description:    moderntreasury.F("description"),
			EffectiveAt:    moderntreasury.F(time.Now()),
			ExternalID:     moderntreasury.F("external_id"),
			LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionNewReversalParamsLedgerableTypeExpectedPayment),
			Metadata: moderntreasury.F(map[string]string{
				"key":    "value",
				"foo":    "bar",
				"modern": "treasury",
			}),
			Status: moderntreasury.F(moderntreasury.LedgerTransactionNewReversalParamsStatusArchived),
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
