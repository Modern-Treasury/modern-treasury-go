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
			Direction:       moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit),
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
			Direction:       moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit),
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
			Direction:       moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerEntriesDirectionCredit),
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
		Description:    moderntreasury.F("string"),
		EffectiveAt:    moderntreasury.F(time.Now()),
		EffectiveDate:  moderntreasury.F(time.Now()),
		ExternalID:     moderntreasury.F("string"),
		LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionNewParamsLedgerableTypeCounterparty),
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
	_, err := client.LedgerTransactions.Get(context.TODO(), "string")
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
		"string",
		moderntreasury.LedgerTransactionUpdateParams{
			Description: moderntreasury.F("string"),
			EffectiveAt: moderntreasury.F(time.Now()),
			LedgerEntries: moderntreasury.F([]moderntreasury.LedgerTransactionUpdateParamsLedgerEntry{{
				Amount:          moderntreasury.F(int64(0)),
				Direction:       moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit),
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
				Direction:       moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit),
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
				Direction:       moderntreasury.F(moderntreasury.LedgerTransactionUpdateParamsLedgerEntriesDirectionCredit),
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
		ID:          moderntreasury.F([]string{"string", "string", "string"}),
		AfterCursor: moderntreasury.F("string"),
		EffectiveAt: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		EffectiveDate: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		ExternalID:              moderntreasury.F("string"),
		LedgerAccountCategoryID: moderntreasury.F("string"),
		LedgerAccountID:         moderntreasury.F("string"),
		LedgerAccountPayoutID:   moderntreasury.F("string"),
		LedgerID:                moderntreasury.F("string"),
		LedgerableID:            moderntreasury.F("string"),
		LedgerableType:          moderntreasury.F(moderntreasury.LedgerTransactionListParamsLedgerableTypeCounterparty),
		Metadata: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		OrderBy: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderBy{
			CreatedAt:   moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByCreatedAtAsc),
			EffectiveAt: moderntreasury.F(moderntreasury.LedgerTransactionListParamsOrderByEffectiveAtAsc),
		}),
		PerPage: moderntreasury.F(int64(0)),
		PostedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		ReversesLedgerTransactionID: moderntreasury.F("string"),
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
		"string",
		moderntreasury.LedgerTransactionNewReversalParams{
			Description:    moderntreasury.F("string"),
			EffectiveAt:    moderntreasury.F(time.Now()),
			ExternalID:     moderntreasury.F("string"),
			LedgerableID:   moderntreasury.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LedgerableType: moderntreasury.F(moderntreasury.LedgerTransactionNewReversalParamsLedgerableTypeCounterparty),
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
