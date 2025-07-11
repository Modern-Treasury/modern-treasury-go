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

func TestPaymentOrderReversalNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentOrders.Reversals.New(
		context.TODO(),
		"payment_order_id",
		moderntreasury.PaymentOrderReversalNewParams{
			Reason: moderntreasury.F(moderntreasury.PaymentOrderReversalNewParamsReasonDuplicate),
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

func TestPaymentOrderReversalGet(t *testing.T) {
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
	_, err := client.PaymentOrders.Reversals.Get(
		context.TODO(),
		"payment_order_id",
		"reversal_id",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentOrderReversalListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentOrders.Reversals.List(
		context.TODO(),
		"payment_order_id",
		moderntreasury.PaymentOrderReversalListParams{
			AfterCursor: moderntreasury.F("after_cursor"),
			PerPage:     moderntreasury.F(int64(0)),
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
