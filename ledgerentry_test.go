// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"context"
	"errors"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestLedgerEntryGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.LedgerEntries.Get(context.TODO(), "string")
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLedgerEntryListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.LedgerEntries.List(context.TODO(), moderntreasury.LedgerEntryListParams{
		AfterCursor:     moderntreasury.F("string"),
		AsOfLockVersion: moderntreasury.F(int64(0)),
		Direction:       moderntreasury.F(moderntreasury.LedgerEntryListParamsDirectionCredit),
		EffectiveAt: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		EffectiveDate: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		ID: moderntreasury.F(map[string]string{
			"foo": "string",
		}),
		LedgerAccountCategoryID: moderntreasury.F("string"),
		LedgerAccountID:         moderntreasury.F("string"),
		LedgerAccountLockVersion: moderntreasury.F(map[string]int64{
			"foo": int64(0),
		}),
		LedgerTransactionID: moderntreasury.F("string"),
		OrderBy: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderBy{
			CreatedAt:   moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderByCreatedAtAsc),
			EffectiveAt: moderntreasury.F(moderntreasury.LedgerEntryListParamsOrderByEffectiveAtAsc),
		}),
		PerPage:     moderntreasury.F(int64(0)),
		ShowDeleted: moderntreasury.F(true),
		Status:      moderntreasury.F(moderntreasury.LedgerEntryListParamsStatusPending),
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
