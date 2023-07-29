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

func TestLedgerTransactionVersionListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := moderntreasury.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
	)
	_, err := client.LedgerTransactions.Versions.List(context.TODO(), moderntreasury.LedgerTransactionVersionListParams{
		AfterCursor: moderntreasury.F("string"),
		CreatedAt: moderntreasury.F(map[string]time.Time{
			"foo": time.Now(),
		}),
		LedgerAccountStatementID: moderntreasury.F("string"),
		LedgerTransactionID:      moderntreasury.F("string"),
		PerPage:                  moderntreasury.F(int64(0)),
		Version: moderntreasury.F(map[string]int64{
			"foo": int64(0),
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
