package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/core"
	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
)

func TestLedgerTransactionVersionListWithOptionalParams(t *testing.T) {
	c := moderntreasury.NewModernTreasury(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.LedgerTransactions.Versions.List(
		context.TODO(),
		"string",
		&requests.LedgerTransactionVersionListParams{AfterCursor: moderntreasury.F("string"), PerPage: moderntreasury.F(int64(0)), CreatedAt: moderntreasury.F(map[string]time.Time{"foo": time.Now()}), Version: moderntreasury.F(map[string]int64{"foo": int64(0)})},
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
